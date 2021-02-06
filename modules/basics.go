package modules

import "fmt"

func examplePrivate() {
	fmt.Println("examplePrivate")
}

func ExamplePublic() {
	fmt.Println("ExamplePublic")
}

func ExampleConstant() {
	const name string = "silva"
	fmt.Println(name)
}

func ExampleVariable() {
	var vname string = "shiwoo"
	vname = "gogogo"

	vname2 := "guess type for me"
	vname2 = "hey!!!"
	fmt.Println(vname + " / " + vname2)
}

func ExampleList() {
	a := []int{1, 2, 3, 4}
	b := []int{5, 6, 7, 8}

	// we can unpack array by ...
	b = append(b, a...)
	fmt.Println(b)
}
