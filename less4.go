package main

import "fmt"

//Go语言循环语句和条件语句
func main() {
	
	simpleCondition()
}

func simpleCondition() {
	
	var a int = 4
	var b int = 3

	if a > 3 && b < 5 {
		fmt.Println("a大于3，b小于5")
	}else{
		fmt.Println("a不大于3或b不小于5")
	}
}

func middleCondition() {
	
	var a = 2
	var b = 4

	select{
		case a = b:
			fmt.Println("a < b")
	}
}

func simpleLoop() {
	for i := 0; i < 10; i++ {
		fmt.Println("%d",i)
	}
}

func middleLoop() {
	var index = 0

	for {
		index ++
		if index > 100 {
				break
			}	
		fmt.Println("%d",index)
	}
}


