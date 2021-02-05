package main

import (
	"fmt"
	"time"
)

func sexyCount(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "is sexy - ", i)
		time.Sleep(time.Second)
	}
}

/*

 */
func isSexy(person string, c chan string) {
	time.Sleep(time.Second * 5)
	// send message through channel
	c <- "sexy " + person
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
	people := [2]string{"silva", "nico"}

	// declare channel
	// - write the data type which channle will contain
	c := make(chan string)
	for _, person := range people {
		// go sexyCount(person)
		go isSexy(person, c)
	}

	fmt.Println("Waiting . . .")
	// wait for 1 message by channel (blocking operation)
	resultOne := <-c
	resultTwo := <-c
	fmt.Println("Received: ", resultOne)
	fmt.Println("Received: ", resultTwo)
	// time.Sleep(time.Second * 10)
}
