package main

import (
	"fmt"
	"log"

	"github.com/shiwoo-park/learngo/banking"
	"github.com/shiwoo-park/learngo/modules"
)

func learnBasics() {
	// Public/Private, Constant, Variable
	modules.ExamplePublic()
	modules.ExampleConstant()
	modules.ExampleVariable()

	// Function - part1
	fmt.Println(modules.Multiply(23, 56))
	totLen, upperName := modules.LenAndUpper("Silva is CooL GuY")
	fmt.Println(upperName, "(", totLen, ")")

	totLen2, _ := modules.LenAndUpper("Great JOb")
	fmt.Println(totLen2)

	modules.RepeatMe("I", "like", "you", "so", "much")

	// Function - part2
	totLen2, upperName2 := modules.LenAndUpperV2("Silva is CooL GuY")
	fmt.Println(upperName2, "(", totLen2, ")")

	// For loop
	total := modules.SuperAdd(11, 12, 13, 14, 15, 16, 17, 18)
	fmt.Println("Total:", total)

	// String Formatting
	modules.StrFormat()

	// Conditional
	fmt.Println("canIDrink(18):", modules.CanIDrink(18))

	// Switch
	fmt.Println("GetScoreByGrade('s'):", modules.GetScoreByGrade("s"))
	fmt.Println("GetScoreByGrade('a'):", modules.GetScoreByGrade("a"))
	fmt.Println("GetScoreByGrade('c'):", modules.GetScoreByGrade("c"))
	fmt.Println("GetScoreByGrade('k'):", modules.GetScoreByGrade("k"))

	// Low-level programming using Pointer
	modules.PlayWithPointer()

	// Array
	modules.PlayWithArray()

	// Map
	modules.PlayWithMap()

	// Struct
	modules.PlayWithStruct()
}

func bankingProject() {
	silvaAccount := banking.NewAccount("silva")
	silvaAccount.Balance()

	silvaAccount.Deposit(10)
	// this would automatically call String()
	fmt.Println(silvaAccount)

	silvaAccount.Withdraw(5)
	fmt.Println(silvaAccount)

	err := silvaAccount.Withdraw(20)
	if err != nil {
		log.Fatalln(err)
	}

}

func dicTest() {
	mydic := banking.Dictionary{"first": "First word"}

	// == Search ==
	// success
	val1, _ := mydic.Search("first")
	fmt.Println("Found !!!", val1)

	// error
	_, err2 := mydic.Search("second")
	if err2 != nil {
		log.Println(err2)
	}

	mydic.Add("second", "This is second")
	mydic.Add("third", "This is third")
	err := mydic.Add("first", "new first")
	if err != nil {
		log.Println(err)
	}

	mydic.Update("first", "new new first")
	mydic.Delete("third")
	mydic.Delete("forth") // ignore

	fmt.Println(mydic)
}

func main() {
	fmt.Println("START LEAN GO !!!!!")
	// learnBasics()
	// bankingProject()
	dicTest()
}
