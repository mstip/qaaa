# QAAA
qaaa tool

### Commands
```
build
===
go build bin/qaaa.exe cmd/main.go

test all
===
go test **/*

dev
===
go run cmd/main.go

dev watch
===
https://github.com/cosmtrek/air
air
```


### Data
```
Project
- name
- []Testsuite

Testsuite
- name
- before
- beforeeach
- after
- aftereach
- []Test

Test
- name
- command
- check

Project
- name: backend
- []Testsuite
    1. Name: User
        - [] Test
            1. Name: register
               command: curl post register newuser
               check: status code 200
            2. Name: login
               command: curl post login newuser
               check: status code 200
```

### Todo
Data refactoren, with related ids ?
deep checks mehr asserts für json
web checks
ui für tests