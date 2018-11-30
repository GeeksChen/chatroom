package main

import "fmt"

func main() {
	
	// var num1 int = 6
	// var num2 int = 8
	// var result int

	var value1 string = "Geek"
	var value2 string = "Chen"

	// result = maxFunction(num1,num2)
	// fmt.Println("max value ",result,"\n")

	a, b := moreValueFunction(value1,value2)

	fmt.Println(a,b)
}

func maxFunction(num1 int,num2 int) int {
	var result int

	if num1 > num2 {
		result = num1
	}else{
		result = num2
	}
	return result
}

func moreValueFunction(value1 string,value2 string) (string,string) {
	return value1,value2
}
