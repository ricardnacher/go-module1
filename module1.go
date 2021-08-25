package module1

import "fmt"

func Greeting() {
	fmt.Println("hello! from module1 pakage")
	if false {
		fmt.Println("never went here")
	}
}

func add(x, y int) int {
	return x + y // Noncompliant
	z := x + y   // dead code
}
