package main

import (
	"fmt"

	"github.com/shiwoo-park/learngo/modules"
)

func main() {
	fmt.Println("START LEAN GO !!!!!")

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
