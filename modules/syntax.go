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
func PlayWithPointer() {
	fmt.Println("=== PlayWithPointer !!! ===")

	// * : look through the address
	// & : get memory address
	a := 2
	b := a // copy value
	a = 10
	c := 5
	fmt.Println("a, b, c:", a, b, c)

	// shows memory address
	fmt.Println("&a, &b, &c:", &a, &b, &c)

	// follow memory address
	a_addr := &a
	fmt.Println("*a_addr:", *a_addr)
	a = 20
	fmt.Println("*a_addr:", *a_addr)
	*a_addr = 30
	fmt.Println("a:", a)
}

func PlayWithArray() {
	// Array (fixed size)
	names := [5]string{"silva", "sun", "matt"}
	names[3] = "david"
	fmt.Println("names:", names)

	// Mutable Array
	grades := []int{87, 66, 99, 45}
	grades = append(grades, 77)
	fmt.Println("grades:", grades)
}

func PlayWithMap() {
	mymap := map[string]string{"city": "jeonju", "name": "nico", "age": "25"}
	mymap["height"] = "177cm"
	fmt.Println(mymap)

	for k, v := range mymap {
		fmt.Println(k, ": ", v)
	}
}
