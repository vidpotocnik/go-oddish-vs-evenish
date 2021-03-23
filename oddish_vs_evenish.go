package main

import (
	"fmt"
	"strconv"
)

/**
Create a function that determines whether a number is Oddish or Evenish.
A number is Oddish if the sum of all of its digits is odd, and a number is Evenish if the sum of all of its digits is even.
If a number is Oddish, return "Oddish". Otherwise, return "Evenish".
For example, oddishOrEvenish(121) should return "Evenish", since 1 + 2 + 1 = 4. oddishOrEvenish(41) should return "Oddish", since 4 + 1 = 5.
*/


func main() {
	fmt.Println(checkOddOrEven(40))
}

func checkOddOrEven(num int) string {
	convert := strconv.Itoa(num)
	result := 0
	var output string
	if len(convert) > 1 {
		for _, digit := range convert {
			convInt, _ := strconv.Atoi(string(digit))
			result += convInt
		}
	} else {
		result = num
	}

	if result % 2 == 0 {
		output = "Oddish"
	} else {
		output = "Evenish"
	}

	return output
}
