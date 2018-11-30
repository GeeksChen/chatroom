package main 

import "fmt"

type Student struct {
	name string
	age int
	weight float
	score float
}

func main() {
	
	

	

}
//创建结构体
func createStudentFunction() {
 	
 	//创建一个新结构体
	fmt.Println(Student{"Geeks_Chen",19,65,88.0})

	//也可以使用key => value格式
	fmt.Println(Student{name:"Geeks_Chen",age:18,weight:69,score:98})

	//忽略字段为0或者空
	fmt.Println(Student{name:"Geeks_Chen",score:100})
}

//读取结构体的属性
func readStudentFunction() {
 	
 	fmt.Println(Student{"Geeks_Chen",19,65,88.0})

 	fmt.Println(Student.name)
}





