package main

import "fmt"

func main() {

	var v_a int = 10
	var v_b int = 8
	var v_c bool = true
	var v_d bool = false
	var result int = 0

	arithmeticTestFunction(v_a,v_b,result)
	
	relationTestFunction(v_a,v_b)

	logicTestFunction(v_c,v_d)

	valuationTestFunction(v_a,v_b)

}

//关系运算符
func relationTestFunction(v_a int,v_b int) {

	if v_a == v_b {
		fmt.Printf("%d = %d\n",v_a,v_b)
	}else{
		fmt.Printf("%d != %d\n",v_a,v_b)
	}

	if v_a > v_b {
		fmt.Printf("%d > %d",v_a,v_b)
	}else{
		fmt.Printf("%d < %d",v_a,v_b)
	}
	
}
//逻辑运算符
func logicTestFunction(v_c bool,v_d bool) {
	
	if v_c && v_d {
		fmt.Printf("两个都是真")
	}else{
		fmt.Printf("两个不都是真")		
	}

	// if v_c || v_d {
	// 	fmt.Printf("至少一个真")
	// }

	// if !v_c {
	// 	fmt.Printf("与v_c相反")
	// }
}

//算术运算符
func arithmeticTestFunction(v_a int,v_b int,result int) {
	
	//加法
	result = v_a + v_b
	fmt.Printf("v_a + v_b = %d\n",result)
	//减法
	result = v_a - v_b
	fmt.Printf("v_a - v_b = %d\n",result)
	//乘法
	result = v_a * v_b
	fmt.Printf("v_a * v_b = %d\n",result)
	//除法
	result = v_a / v_b
	fmt.Printf("v_a / v_b = %d\n",result)
	//取余
	result = v_a % v_b
	fmt.Printf("%d\n",result)
	//自增
	v_a++
	fmt.Printf("v_a++ = %d\n",v_a)
	//自减
	v_b--
	fmt.Printf("v_b-- = %d",v_b)
}

func valuationTestFunction(v_a int,v_b int) {
	
	v_a = v_b

	v_a += v_b

	v_a -= v_b

	v_a *= v_b

	v_a /= v_b

	v_a %= v_b
	
}



