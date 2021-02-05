package main

import (
	"fmt"
	"time"
)

func sexyCount(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "is sexy")
		time.Sleep(time.Second)
	}
}

func main() {
	fmt.Println("START LEAN GO !!!!!")

	// #1 ---- Theory ----
	// modules.LearnBasics()

	// #2 ---- Bank & Dict Project ----
	// banking.BankingProject()
	// modules.DicTest()

	// #3 ---- URL Checker & Go Routine ----
	// urlcheck.CheckV1()
	sexyCount("silva")
	sexyCount("nico")
}
