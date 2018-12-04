package main 

import "fmt"

//Go 递归函数
func recursion() {
	recursion() //函数调用自身
}
// 阶乘
func Factorial(n uint64)(result uint64) {
	if n > 0 {
		result = n * Factorial(n-1)
		return result
	}
	return 1
}

//斐波那楔数列
func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-2) + fibonacci(n-1)
}

func main() {
	
	recursion()

	i := 15
	fmt.Println("%d 的阶乘 %d\n",Factorial(uint64(i)) //15的阶乘

}