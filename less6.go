package main

import "fmt"

//全局变量
var b int 20

func main() {
	//局部变量
	var a int = 10
	//全局变量和局部变量同名时，局部变量优先级高
	var b int = 30
	fmt.prinfln(a,b)
	
	//形式参数做局部变量处理
	var c = sum(a,b)
}

func sum(a,b int) int{
	return a+b
}