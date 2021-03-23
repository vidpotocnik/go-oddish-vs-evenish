package main

import (
	"fmt"
	"strconv"
)


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
