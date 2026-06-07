package main

import (
	"flag"
	"fmt"
)

func flagtry() (bool, bool, string, int) {

	digits := flag.Bool("digits", true, "Presence of digits in password")
	specialSymbols := flag.Bool("symbols", true, "Presence of special signs in password")
	hardness := flag.String("hardness", "medium", "Hardness of password: easy, medium, hard")
	flag.Parse()

	if *digits == true {
		fmt.Println("Enter the digit: ")
		var n int
		fmt.Scan(&n)
		return *digits, *specialSymbols, *hardness, n
	}

	fmt.Println("Digits: ", *digits)
	fmt.Println("Special signs: ", *specialSymbols)
	fmt.Println("Hardness of password: ", *hardness)

	return *digits, *specialSymbols, *hardness, 0
}
