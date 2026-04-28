package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/google/go-github/v60/github"
	"github.com/joho/godotenv"
)

func main() {

	_ = godotenv.Load()
	logger := log.New(os.Stdout, "[MARROW-PRO] ", log.LstdFlags)

	// Define CLI Flags for a commercial feel
	// Usage: go run main.go -repos="repo1,repo2" -org="Marrow-Stack"
	repoList := flag.String("repos", "", "Comma-separated list of repositories to migrate")
	targetOrg := flag.String("org", "Marrow-Stack", "The destination GitHub Organization")
	flag.Parse()

	if *repoList == "" {
		logger.Fatal("Missing required flag: -repos. Example: -repos=\"repo1,repo2\"")
	}

	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		logger.Fatal("Environment variable GITHUB_TOKEN is not set")
	}

	ctx := context.Background()
	client := github.NewClient(nil).WithAuthToken(token)

	user, _, err := client.Users.Get(ctx, "")
	if err != nil {
		logger.Fatalf("Authentication failed: %v", err)
	}
	logger.Printf("Authenticated as %s. Target Org: %s", *user.Login, *targetOrg)


	repos := strings.Split(*repoList, ",")
	var wg sync.WaitGroup

	fmt.Printf("\n🚀 Starting migration for %d repositories...\n\n", len(repos))

	for _, rName := range repos {
		rName = strings.TrimSpace(rName) 
		if rName == "" {
			continue
		}

		wg.Add(1) 

		
		go func(name string) {
			defer wg.Done() 
			migrate(ctx, client, *user.Login, name, *targetOrg, logger)
		}(rName)
	}

	
	wg.Wait()
	fmt.Println("\n🏁 Batch migration process complete. Check your Org dashboard.")
}

func migrate(ctx context.Context, client *github.Client, owner, repo, org string, logger *log.Logger) {
	req := github.TransferRequest{NewOwner: org}
	_, _, err := client.Repositories.Transfer(ctx, owner, repo, req)

	if err != nil {
		
		if strings.Contains(err.Error(), "job scheduled") {
			logger.Printf("✅ QUEUED: %s", repo)
			return
		}
		logger.Printf("❌ FAILED: %s | Error: %v", repo, err)
		return
	}

	logger.Printf("✨ INSTANT SUCCESS: %s", repo)
}