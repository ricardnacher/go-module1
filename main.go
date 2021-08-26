package main

import (
	"fmt"
)

func main() {
	fmt.Println("program1")
	NeverGetInIf()

}

func NeverGetInIf() {
	if false {
		fmt.Println("program1")
	}
}
