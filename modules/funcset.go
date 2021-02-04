package modules

import (
	"fmt"
	"strings"
)

// Multiply documentation comes here for this function
func Multiply(a, b int) int {
	return a * b
}

/*
LenAndUpper desc title

multi-lined comment looks like this.
it has multiple return value
*/
func LenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}
func LenAndUpperV2(name string) (nameLength int, nameUpper string) {
	// post function hook
	defer fmt.Println("After LenAndUpperV2 finished!!!")

	// initialized with declaration
	nameLength = 10
	nameUpper = strings.ToUpper(name)

	// automatically return = NAKED return
	return
}
func RepeatMe(words ...string) {
	// how to get multiple argumets as a list
	fmt.Println(words)
}
func SuperAdd(numbers ...int) int {
	sum := 0
	for _, number := range numbers {
		fmt.Println("number: ", number)
		sum += number
	}

	sum = 0
	for index, number := range numbers {
		fmt.Println(index, "th: ", number)
		sum += number
	}

	sum = 0
	for i := 0; i < len(numbers); i++ {
		fmt.Printf("numbers[%d] = %d\n", i, numbers[i])
		sum += numbers[i]
	}

	return sum
}

func CanIDrink(age int) bool {

	// we can declare variable in "if" clause and use it
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	} else if age > 60 {
		return false
	} else {
		return true
	}
}

func GetScoreByGrade(grade string) int {
	switch upperGrade := strings.ToUpper(grade); upperGrade {
	case "S":
		return 100
	case "A":
		return 90
	case "B":
		return 80
	case "C":
		return 70
	}
	return 50
}
