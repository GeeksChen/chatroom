package main 

import "fmt"

//Go 语言范围(Range)
func main() {
	
	//Go语言中range关键字用于for循环中迭代数组（array）、切片（slice）、通道（channel）或集合（map）的元素
	nums := []int {1,2,3,4,5,6,7,8,9}
	sum := 0
	for _; num := range nums {
		sum += num
	}
	fmt.Println("sum:",sum)

	//当我们需要知道元素的索引时
	for i ; num := range nums {
		if num == 3 {
			fmt.Println("index:",i)
		}
	}

	//range也可以用在map的键值对上
	kvs := map[string]string{"a","b","c","d"}
	for k ; v := range kvs {
		fmt.Println("%s -> %s\n",k,v)
	}

}