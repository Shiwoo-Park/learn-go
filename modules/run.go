package modules

import (
	"fmt"
	"log"
)

func LearnBasics() {

	// Public/Private, Constant, Variable
	ExamplePublic()
	ExampleConstant()
	ExampleVariable()
	ExampleList()

	// Function - part1
	fmt.Println(Multiply(23, 56))
	totLen, upperName := LenAndUpper("Silva is CooL GuY")
	fmt.Println(upperName, "(", totLen, ")")

	totLen2, _ := LenAndUpper("Great JOb")
	fmt.Println(totLen2)

	RepeatMe("I", "like", "you", "so", "much")

	// Function - part2
	totLen2, upperName2 := LenAndUpperV2("Silva is CooL GuY")
	fmt.Println(upperName2, "(", totLen2, ")")

	// For loop
	total := SuperAdd(11, 12, 13, 14, 15, 16, 17, 18)
	fmt.Println("Total:", total)

	// String Formatting
	StrFormat()

	// Conditional
	fmt.Println("canIDrink(18):", CanIDrink(18))

	// Switch
	fmt.Println("GetScoreByGrade('s'):", GetScoreByGrade("s"))
	fmt.Println("GetScoreByGrade('a'):", GetScoreByGrade("a"))
	fmt.Println("GetScoreByGrade('c'):", GetScoreByGrade("c"))
	fmt.Println("GetScoreByGrade('k'):", GetScoreByGrade("k"))

	// Low-level programming using Pointer
	PlayWithPointer()

	// Array
	PlayWithArray()

	// Map
	PlayWithMap()

	// Struct
	PlayWithStruct()
}

func DicTypeTest() {
	mydic := Dictionary{"first": "First word"}

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
