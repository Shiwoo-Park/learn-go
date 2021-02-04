package modules

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func StrFormat() {
	fmt.Println("=== DIVERSE STRING FORMATTING !!! ===")
	p := point{1, 2}

	fmt.Print("%v = ")
	fmt.Printf("%v\n", p)

	fmt.Print("%+v = ")
	fmt.Printf("%+v\n", p)

	fmt.Print("%#v = ")
	fmt.Printf("%#v\n", p)

	fmt.Print("%T = ")
	fmt.Printf("%T\n", p)

	fmt.Print("%t = ")
	fmt.Printf("%t\n", true)

	fmt.Print("%d = ")
	fmt.Printf("%d\n", 123)

	fmt.Print("%b = ")
	fmt.Printf("%b\n", 14)

	fmt.Print("%c = ")
	fmt.Printf("%c\n", 33)

	fmt.Print("%x = ")
	fmt.Printf("%x\n", 456)

	fmt.Print("%f = ")
	fmt.Printf("%f\n", 78.9)

	fmt.Print("%e = ")
	fmt.Printf("%e\n", 123400000.0)

	fmt.Print("%E = ")
	fmt.Printf("%E\n", 123400000.0)

	fmt.Print("%s = ")
	fmt.Printf("%s\n", "\"string\"")

	fmt.Print("%q = ")
	fmt.Printf("%q\n", "\"string\"")

	fmt.Print("%x = ")
	fmt.Printf("%x\n", "hex this")

	fmt.Print("%p = ")
	fmt.Printf("%p\n", &p)

	fmt.Print("|%6d|%6d| = ")
	fmt.Printf("|%6d|%6d|\n", 12, 345)

	fmt.Print("|%6.2f|%6.2f| = ")
	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)

	fmt.Print("|%-6.2f|%-6.2f| = ")
	fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)

	fmt.Print("|%6s|%6s| = ")
	fmt.Printf("|%6s|%6s|\n", "foo", "b")

	fmt.Print("|%-6s|%-6s| = ")
	fmt.Printf("|%-6s|%-6s|\n", "foo", "b")

	s := fmt.Sprintf("a %s", "string")
	fmt.Println(s)

	fmt.Fprintf(os.Stderr, "an %s\n", "error")
}
