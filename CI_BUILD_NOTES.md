# CI Build Notes

## Successfully Tested Locally

The build has been successfully tested locally with the following steps:

### 1. Environment
- Go 1.19+ (tested with Go 1.24.5)
- All dependencies resolved

### 2. Build Steps Executed
```bash
# Install dependencies
go get -v -t -d ./...

# Install pkger
go install -v github.com/markbates/pkger/cmd/pkger@v0.17.1

# Generate pkged.go (IMPORTANT: Use -mod=mod to bypass vendor)
GO111MODULE=on go run -mod=mod github.com/markbates/pkger/cmd/pkger -o bundled

# Build the CLI tool
GO111MODULE=on go build ./cli/gocomply_fedramp

# Run tests (excluding examples)
go test $(go list ./... | grep -v /examples)
```

### 3. Key Fixes Applied
- Updated GitHub Actions workflows to use Go 1.19+
- Fixed pkger installation to use `go install` (required for Go 1.17+)
- Resolved type conflicts in new R5 Balance files
- Removed interfering `bundled/pkger.go` file

### 4. Important Note About bundled/pkged.go
The `bundled/pkged.go` file is a large (5.9MB) generated file that contains embedded templates and catalogs. It is:
- Already tracked in the repository
- Too large to push in some cases
- Must be regenerated locally after pulling changes

To regenerate:
```bash
GO111MODULE=on go run -mod=mod github.com/markbates/pkger/cmd/pkger -o bundled
```

### 5. CI Should Now Pass
With the updated workflows and fixes, the CI should successfully:
1. Use Go 1.19+ (not 1.14.15)
2. Install pkger properly
3. Build without type conflicts
4. Pass all tests 