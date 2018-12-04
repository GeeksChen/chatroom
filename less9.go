package main 

import "fmt"

//Go 语言的切片(动态数组)
func main() {
	
	//创建切片1
	slice1 := [] int {1,2,3,4,5} //slice1 切片名称 int 切片类型 {1,2,3,4,5}表示切片的值
	//创建切片2
	slice2 := arr[:] //初始化slice2是数组arr的引用
	//创建切片3
	slice3 := arr[startIndex : endIndex] //将arr中从下标startIndex到endIndex-1下的元素创建为一个新切片
	//创建切片4
	slice4 := slice3[startIndex:endIndex] //根据切片slice3创建slice4
	//创建切片5
	slcie5 := make([]int,len,cap) //通过内置函数创建切片

	//获取切片长度
	len()
	//获取切片容量
	cap()

	//切片截取
	slice1[1,3]

	//切片拷贝和扩容
	append(slice1,9)
	slice6 := make([]int,len(slice1)*2,cap(slice1)*2)
	copy(slice6,slice1)// 拷贝slice1的内容到slice6中


}








