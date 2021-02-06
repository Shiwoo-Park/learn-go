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
- The MAIN reason we learn about Go is
  * when you want to resolve repeating independent jobs fast
  * when you need concurrency with robust code and easy implementation
- Go-routine
  * put `go` in front of the function call
  * Control flow just skip the go routine ()
  * You can use it if you want to execute the funciton asyncronously
  * Go-routine can alive only while main function(=process) is running
  * Main function do not wait for go routine
- Channel
  * We need to define the type of data send through the channel
  * Communication method between go-routine and main()
  * Communication method between go-routine and another go-routine
  * We can send a message to channel in inside of function `c <- sendingMsg`
  * And We can receive a messsage send from channel by blocking operation `receivedMsg := <-c`
  * If we use blocking operation when there's no message in channel, the main function experience deadlock and we can't have any clue about it. SO BECAREFUL WHEN YOU USE IT !!!
  * If we want to use some func as just producer (`for only send msg`) we can declare channel argument using `chan<-`


## ETC
- Go [built-in packages](https://golang.org/pkg/) are very huge and has a lot of useful functions.
- Go [Documentation](https://golang.org/doc/)


## Commands

```bash
go run main.go
```