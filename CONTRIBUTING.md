# Contributing
## Getting Started
### 0. Prerequisites

- [Golang](https://golang.org)
- [GitHub CLI](https://cli.github.com/)
- [GoReleaser](https://goreleaser.com/)

### 1. Clone this repo.
```bash
$ gh repo clone spesnova/dotkeeper
$ cd dotkeeper
```

### 2. Install dependencies
```bash
$ go mod tidy
```

### 3. Set up GitHub token
Create a GitHub token with the `repo` scope.

Then, store it in `.github_token`.

```bash
$ touch .github_token
```

## Tasks
### Running the app
```bash
$ go run main.go
```

### Testing
```bash
$ go test ./...
```

### Versioning
This CLI follows [Semantic Versioning](https://semver.org/).

- Update the version in `internal/version/version.go`
- Update the version in `README.md`

### Building binaries
goreleaser is used to build binaries.

```bash
$ goreleaser build --snapshot --clean
```

#### Building binaries manually
```bash
# Linux AMD
GOOS=linux GOARCH=amd64 go build -o bin/dotkeeper-linux-amd64

# Linux ARM
GOOS=linux GOARCH=arm64 go build -o bin/dotkeeper-linux-arm64

# Intel Mac
GOOS=darwin GOARCH=amd64 go build -o bin/dotkeeper-mac-amd64

# ARM Mac
GOOS=darwin GOARCH=arm64 go build -o bin/dotkeeper-mac-arm64
```

### Releasing
First, check if the goreleaser configuration file is valid.

```bash
$ goreleaser check
```

Then, create a new release.
```bash
$ git tag -a v0.1.0 -m "First release"
$ git push origin v0.1.0
```

Then, release the binaries to GitHub.

```bash
$ goreleaser release --clean
```
