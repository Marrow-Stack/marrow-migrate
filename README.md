# Marrow-Migrate

**A simple Go-based CLI tool to bulk migrate GitHub repositories from one user or organization to another in a single command.**

Part of the [Marrow Stack](https://marrowstack.dev) ecosystem.



## Screenshot

<img width="810" height="493" alt="Screenshot from 2026-04-29 00-53-32" src="https://github.com/user-attachments/assets/93559ead-79ec-41aa-8e18-3fbc1eaed685" />

## Overview

Transferring multiple GitHub repositories manually can be repetitive and slow, especially when moving projects between personal accounts and organizations.

`marrow-migrate` helps automate that process from the terminal so you can transfer many repositories in one workflow instead of handling each repository one by one.

## Why use marrow-migrate?

- Bulk transfer repositories in a single command
- Reduce repetitive manual work in the GitHub UI
- Preview migrations safely with dry-run mode
- Use a lightweight compiled Go binary
- Run the tool across Linux, macOS, and Windows

## Features

- Migrate repositories from user to organization
- Migrate repositories from organization to organization
- Migrate repositories from user to user
- Transfer all repositories or only selected repositories
- Preview planned actions with `--dry-run`
- Use a GitHub token from an environment variable or CLI flag

## Installation

### Option 1: Install with Go

```bash
go install github.com/Marrow-Stack/marrow-migrate@latest
```

### Option 2: Build from source

```bash
git clone https://github.com/Marrow-Stack/marrow-migrate.git
cd marrow-migrate
go build -o marrow-migrate main.go
sudo mv marrow-migrate /usr/local/bin/
```

## Requirements

Before running the tool, make sure you have:

- Go 1.21 or later if installing from source
- A GitHub Personal Access Token
- Permission to transfer repositories from the source account
- Permission to create or receive repositories in the target account or organization

## Authentication

`marrow-migrate` requires a GitHub Personal Access Token.

Recommended scopes:

- `repo` — full control of private repositories
- `admin:org` — required for organization transfers

Recommended setup:

```bash
export GITHUB_TOKEN=ghp_your_actual_token_here
```

You can also provide the token directly with the `--token` flag.

## Usage

### Migrate all repositories from one organization to another

```bash
marrow-migrate --source-org old-org --target-org new-org
```

### Migrate repositories from a personal account to an organization

```bash
marrow-migrate --source-user samarth --target-org marrow-stack
```

### Preview a migration without making changes

```bash
marrow-migrate --source-org old-org --target-org new-org --dry-run
```

### Migrate only selected repositories

```bash
marrow-migrate --source-org old-org --target-org new-org --repos repo1,repo2,api-service
```

## Flags

| Flag | Description |
|------|-------------|
| `--source-org` | Source GitHub organization |
| `--source-user` | Source GitHub username |
| `--target-org` | Target GitHub organization |
| `--target-user` | Target GitHub username |
| `--repos` | Comma-separated list of specific repository names |
| `--dry-run` | Preview the migration without executing it |
| `--token` | GitHub token, optional if `GITHUB_TOKEN` is set |

## Rules

- Provide either `--source-org` or `--source-user`
- Provide either `--target-org` or `--target-user`
- Use `--repos` when only specific repositories should be transferred
- Use `--dry-run` first if you want to validate the operation before execution

## Example workflows

### Organization to organization

```bash
marrow-migrate \
  --source-org old-org \
  --target-org new-org
```

### User to organization with dry run

```bash
marrow-migrate \
  --source-user samarth \
  --target-org marrow-stack \
  --dry-run
```

### User to user for selected repositories

```bash
marrow-migrate \
  --source-user old-user \
  --target-user new-user \
  --repos repo1,repo2
```

## Current status

`marrow-migrate` is in early development and is being built as part of the Marrow Stack proof-of-work journey.

## Roadmap

- Improve error handling and logging
- Add progress indicators
- Support migrating releases and tags
- Add interactive mode
- Expand documentation and usage examples

## License

This project is licensed under the MIT License. See the `LICENSE` file for details.

## Author

Built by **Samarth**.

Solo builder of [Marrow Stack](https://marrowstack.dev) — production-ready Next.js + Supabase code blocks.


