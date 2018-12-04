package main 

import "fmt"

//Go语言 数组和指针
func main() {
	//数组初始化
	var arr = [3]int{1,2,5}
	var arr1 = [...]int{1,4,5,6}
	//数组访问
	var value1 = arr1[1]//访问数组第二位的数据

	var ip *int //声明指针

	ip = &arr //存储数组arr的地址

	fmt.println("arr变量地址 %x\n",&arr) //获取数组arr地址

	fmt.println("指针变量地址 %x\n",ip) //获取指针变量的地址

	fmt.println("*ip变量值 %d\n",*ip) //使用指针访问值

	//当一个指针被定义后灭有分配到任何变量时，他的值是nil
	var ptr *int

	if ptr != nil { } //ptr不是空指针
	if ptr == nil { } //ptr是空指针
}