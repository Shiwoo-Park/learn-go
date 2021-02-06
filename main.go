package main

import (
	"fmt"

	"github.com/shiwoo-park/learngo/jobscrapper"
)

type BigStruct struct {
	Data [10]byte
}

func main() {
	fmt.Println("START LEAN GO !!!!!")

	// #1 ---- Theory ----
	// modules.LearnBasics()
	// modules.StrFormat()

	// #2 ---- Bank & Dict Project ----
	// banking.BankingProject()
	// modules.DicTypeTest()

	// #3 ---- URL Checker & Go Routine ----
	// urlcheck.CheckV1()
	// modules.ExperienceChannelAndGoRoutine()
	// urlcheck.CheckV2()

	// #4 ---- Job Scrapper ----
	jobscrapper.ExampleScrape()
}
