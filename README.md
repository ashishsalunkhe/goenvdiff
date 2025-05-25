# goenvdiff

Git-aware `.env` drift detector for platform engineers, SREs, and developers.

**goenvdiff** identifies differences in `.env` configuration files across Git branches or commits. It compares key-value pairs and shows what’s added, removed, or changed — useful during code reviews, DevOps audits, or deployment validations.

## Features

* Git-aware comparison of `.env` files across branches or commits
* Shows added, removed, and modified keys
* Color-coded terminal output
* JSON output support via `--json` for automation
* Simple and extensible CLI using Cobra

## Installation & Environment Setup

### Requirements

* Go 1.20 or higher
* Git CLI installed
* VS Code or any text editor

### Setup Locally

```bash
git clone https://github.com/yourusername/goenvdiff.git
cd goenvdiff
go mod tidy
go build -o goenvdiff
```

### Optional: Install Globally

```bash
go install github.com/yourusername/goenvdiff@latest
```

## Usage

```bash
goenvdiff --from main --to feature/login --path .env
```

### Sample Output

```
+ API_KEY added (abc123)
- DEBUG removed (was true)
~ PORT changed from 8080 to 9090
```

### JSON Mode

```bash
goenvdiff --from main --to feature/login --path .env --json
```

```json
[
  { "Key": "API_KEY", "OldValue": "", "NewValue": "abc123", "Type": 0 },
  { "Key": "DEBUG", "OldValue": "true", "NewValue": "", "Type": 1 },
  { "Key": "PORT", "OldValue": "8080", "NewValue": "9090", "Type": 2 }
]
```

## Project Structure

```
.
├── main.go               # Entry point; wires up CLI
├── cmd/
│   └── root.go           # Cobra CLI setup and flag handling
├── internal/
│   ├── diff.go           # Core diffing logic
│   ├── git.go            # Reads `.env` from Git refs (branches/commits)
│   └── env.go            # Placeholder for env validation or merging
├── go.mod                # Go module metadata
├── .gitignore
├── README.md             # Project documentation
└── testdata/             # (Optional) Test fixtures for `.env` files
```

## Architecture Overview

```
+------------+         +------------------+         +---------------+
| Git Commit |  --->   | Read .env file   |  --->   | Parse KeyVals |
+------------+         +------------------+         +---------------+
       |                                                  |
       |                                                  v
       |                                      +--------------------------+
       +---> another Git ref --->            | Diff Key-Value Pairs      |
                                             |  - Added / Removed / Mod  |
                                             +--------------------------+
                                                             |
                                                             v
                                      +------------------------------------+
                                      | Print Output or Export as JSON     |
                                      +------------------------------------+
```

## Running Tests

```bash
go test ./internal/...
```

## Steps to Reproduce

1. Create `.env` files on two Git branches:

```bash
echo "PORT=8080\nDEBUG=true" > .env
git checkout -b main
git add .env && git commit -m "main env"

git checkout -b feature/login
echo "PORT=9090\nAPI_KEY=xyz" > .env
git commit -am "feature env"
```

2. Run the tool:

```bash
goenvdiff --from main --to feature/login --path .env
```

## Roadmap

* [ ] Support multi-file diffs (`.env.production`, `.env.development`)
* [ ] HTML/Markdown export
* [ ] Git pre-commit integration
* [ ] Homebrew install formula

## License

MIT — feel free to use, extend, and contribute.

## Contributing

Feel free to fork, file issues, or submit PRs. Run `go fmt ./...` before committing.
