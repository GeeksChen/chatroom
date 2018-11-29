package main

// import "fmt"

func main() {

	constTestFunction()
	
	varTestFunction()
}

func constTestFunction() {

	const c_a string = "This is a const value"	
	const c_b int = 10
	const c_c bool = false

	println(c_a,c_b,c_c);
}

//变量由数字，字母，下划线组成，数字不能开头；通过变量名访问变量；使用var修饰
func varTestFunction() {
	var v_a string = "Geeks_Chen is Geek"
	var v_b int = 3
	var v_c bool = true

	println(v_a,v_b,v_c)
}