package main

import (
	"fmt"

	"github.com/shiwoo-park/learngo/modules"
)

func main() {
	fmt.Println("START LEAN GO !!!!!")

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
}
