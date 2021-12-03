# QAAA
[![Build and Test](https://github.com/mstip/qaaa/actions/workflows/build_and_test.yml/badge.svg)](https://github.com/mstip/qaaa/actions/workflows/build_and_test.yml)

qaaa tool

# Commands
## build
```
go build bin/qaaa.exe main.go
```
## test all
```
go test -v ./...
```
## dev
```
go run main.go
```
## dev watch
```
install: https://github.com/cosmtrek/air
air
```
## release
```
https://goreleaser.com/
go install github.com/goreleaser/goreleaser@latest
local: goreleaser release --snapshot --rm-dist
new release:
export GITHUB_TOKEN="YOUR_GH_TOKEN"
git tag -a v0.1.0 -m "First release"
git push origin v0.1.0
goreleaser release --rm-dist
```