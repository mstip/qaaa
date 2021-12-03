# QAAA
qaaa tool

### Commands
```
build
===
go build bin/qaaa.exe main.go

test all
===
go test ./... -v

dev
===
go run main.go

dev watch
===
https://github.com/cosmtrek/air
air

go releaser
===
https://goreleaser.com/
go install github.com/goreleaser/goreleaser@latest
local: goreleaser release --snapshot --rm-dist
new release:
export GITHUB_TOKEN="YOUR_GH_TOKEN"
git tag -a v0.1.0 -m "First release"
git push origin v0.1.0
goreleaser release --rm-dist
```


### Todo
- ui für projekte crud
- ui für suite crud
- ui für task crud
- task create view mit viel hilfe und js und simulation etc
- web tasks
- cmd tasks
- manual tasks
- deep checks mehr asserts für json
- project und suite variablen
- links, iframes, externe views/details (jira, confluence, bitbucket, github etc...)
- use https://github.com/thedevsaddam/gojsonq to check json
- templates function refactorn - auto discover etc
- controller injection