# Learning Go with nomadcoder

[Lecture Link](https://nomadcoders.co/go-for-beginners/lectures/1499)

## What I learned about Go

- Use `snake_case` to file name (ex) `some_util.go`
- Use `PascalCase` or `camelCase` to all kinds of element declaration in code
- Public & Private rule
  * It totally depends on how we `name` the element
  * `PascalCase`: public (= exported) / `camelCase`: private
  * This rule applies to the name of `variable`, `function`, `struct`, `key of struct`
- Go doesn't have package manager like pip in python or npm in Node
  * So you need to organize all external package in some custom directory
  * you can organize dir like: `~/go/src/github.com/USER_ID/REPO`
- Project entrypoint
  * Go specifically looking for `main()` under `package main` in `main.go` file for entry point
  * main.go exists only for projects which want to be compiled.


## ETC
- Go built-in package is very huge and has a lot of useful functions. [take look!](https://golang.org/pkg/)


## Commands

```bash
go run main.go
```