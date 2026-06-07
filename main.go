package main

import (
	"fmt"
)

var specialSymbols = []rune{'!', '@', '#', '$', '%', '^', '&', '?', '*', '(', ')', '/'}
var DIGITS = []rune{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
var hardnesses = make(map[string][2]int, 3)
var letters = []rune("qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM")

func main() {
	hardnesses["easy"] = [2]int{7, 2}
	hardnesses["medium"] = [2]int{15, 4}
	hardnesses["hard"] = [2]int{23, 6}
	one, two, three, four := flagtry()
	fmt.Println(generatePassword(&one, &two, &three, four))
}
