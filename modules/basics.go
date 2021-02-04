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
