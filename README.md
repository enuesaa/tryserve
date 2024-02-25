# tryup
Instant web server for development.

## Usage
```console
$ tryup --help
Instant web server

USAGE:
  tryup <path> [global options]

GLOBAL OPTIONS:
  --watch        run watch mode (default: false)
  --help, -h     show help
  --version, -v  print the version
```

## Development Plan
- [runapp] watch mode
- [serve] index
- [serve] custom logger

### Planning Usecase
```bash
tryup .       # this serve static content
tryup main.go # this run `go run main.go` internally
tryup main.go --watch . # this run `go run main.go` and also, do hot reload

# or 
tryup go run main.go
tryup pnpm dev
```
