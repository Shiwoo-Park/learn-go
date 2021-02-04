package modules

import (
	"fmt"
	"strings"
)

func Multiply(a, b int) int {
	return a * b
}
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
	fmt.Println(words)
}
