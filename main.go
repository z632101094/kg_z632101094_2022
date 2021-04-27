package main

import (
	"fmt"
	"os"
)

func main() {
	nums := os.Args[1:] //get the input
	result := ""
	for i := 0; i < len(nums); i++ { //iterate the input array
		result = result + "," + translate(nums[i])
	}
	if len(result) > 0 {
		fmt.Println(result[1:])
	} else {
		fmt.Println() //if input is null, it can avoid exception
	}
}

func translate(x string) string {
	result := ""
	for _, char := range x {
		switch char {
		case '0':
			result += "Zero"
		case '1':
			result += "One"
		case '2':
			result += "Two"
		case '3':
			result += "Three"
		case '4':
			result += "Four"
		case '5':
			result += "Five"
		case '6':
			result += "Six"
		case '7':
			result += "Seven"
		case '8':
			result += "Eight"
		case '9':
			result += "Nine"
		}
	}
	return result
}
