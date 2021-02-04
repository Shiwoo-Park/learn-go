package modules

import "fmt"

// No constructor(__init__) function provided to struct
type person struct {
	name    string
	age     int
	favFood []string
}

func PlayWithStruct() {
	favFood := []string{"kimchi", "bulgogi"}

	silva := person{"silva", 29, favFood}
	fmt.Println("silva: ", silva)

	// this is better to declaration
	mimi := person{name: "mimi", age: 13, favFood: favFood}
	fmt.Println("mimi: ", mimi)
}
