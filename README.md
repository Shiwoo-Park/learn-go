# Learning Go with nomadcoder

[Lecture Link](https://nomadcoders.co/go-for-beginners/lobby)

## What I learned about Go

- Project structure

```
~/go/src/github.com/username/myproj/
  main.go
  mypackage1/
    mypackage1.go
    util.go
  mypackage2/
    mypackage2.go
    types.go
```

- Code convention
  - Use `snake_case` to file name (ex) `some_util.go`
  - Use `PascalCase` or `camelCase` to all kinds of element declaration in code
- Public & Private rule
  - It totally depends on how we `name` the element
  - `PascalCase`: public (= exported) / `camelCase`: private
  - This rule applies to the name of `variable`, `function`, `struct`, `key of struct`
- Organizing Go modules
  - Go doesn't have package manager like pip in python or npm in Node
  - So you need to organize all external package in some custom directory
  - You can organize dir like: `~/go/src/github.com/USER_ID/REPO`
  - You can use `go mod` command and refer to [this](https://blog.golang.org/using-go-modules) document
- Project entrypoint
  - Go specifically looking for `main()` under `package main` in `main.go` file for entry point
  - main.go exists only for projects which want to be compiled.
- The MAIN reason we learn about Go is
  - when you want to resolve repeating independent jobs fast
  - when you need concurrency with robust code and easy implementation
- Go-routine
  - `Go-routine` is a function or method which executes independently and simultaneously in connection with any other Goroutines present in your program
  - It allows us to execute program concurrently
  - How to use it? Just put `go` in front of the function call
  - Control flow just skip the go routine ()
  - You can use it if you want to execute the funciton asyncronously
  - Go-routine can alive only while main function(=process) is running
  - Main function do not wait for go routine
- Channel
  - `Go Channel` is a means through which different goroutines communicate.
  - Think of them as pipes through which you can connect with different concurrent goroutines.
  - We need to define the type of data send through the channel
  - Communication method between go-routine and main()
  - Communication method between go-routine and another go-routine
  - We can send a message to channel in inside of function `c <- sendingMsg`
  - And We can receive a messsage send from channel by blocking operation `receivedMsg := <-c`
  - If we use blocking operation when there's no message in channel, the main function experience deadlock and we can't have any clue about it. SO BECAREFUL WHEN YOU USE IT !!!
  - If we want to use some func as just producer (`for only send msg`) we can declare channel argument using `chan<-`

## ETC

- Go [built-in packages](https://golang.org/pkg/) are very huge and has a lot of useful functions.
- Go [Documentation](https://golang.org/doc/)
- Web minimal framework [Echo](https://echo.labstack.com/)
- Web full-stack framework [buffalo](https://gobuffalo.io/en/)

## Commands

```bash
// run main func
go run main.go

// go module manage
go mod init
go mod tidy
```
