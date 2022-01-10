package main

import "fmt"

map romanNumeral[string]int := {}

func romanToInteger(input string) int {

	return 0
}

func main() {
	fmt.Println(romanToInteger("III"))     // 3
	fmt.Println(romanToInteger("LVIII"))   // 58
	fmt.Println(romanToInteger("MCMXCIV")) // 1994
}

// ***ALGORITHM***
//  - create a hash {romanNumeral: value }
//  - create an integer variable to hold total
//  - loop over the input string
//  - always check the next letter to see if its value is greater or less than the current
//     - if less than, subtract from the current sum
//     - if greater than or equal, add to the sum
