package main

import (
	"fmt"
)

//定义接口
type Phone interface{
	call()
}

//定义结构体
type IPhone struct{

}

func (iphone iphone )call() {
	fmt.Println("i am iphone")
}

//定义结构体
type Summing struct{

}

func (summing Summing) call() {
	fmt.Println("i am Summing")
}


func main() {
	
	var phone Phone
	phone = new(IPhone)
	phone.call()
}