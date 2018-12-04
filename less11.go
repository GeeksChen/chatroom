package main 

import "fmt"

//Go 语言Map

//Map是一种无序键值对的集合，Map最重要的一点是通过key来快速检索数据，key类似于索引，指向数据的值
func main() {
	
	var countryMap map[string]string //创建集合
	countryMap = make(map[string]string)	
	
	//插入数据
	countryMap["chian"] = "北京"
	countryMap["japan"] = "东京"

	//循环输出数据
	for country := range countryMap {
		fmt.Println(country,"首都是：",countryMap[country])
	}

	//查询某个元素是否存在
	captial, ok := countryMap["hanguo"]
	if ok {
		fmt.Println("hanguo 首都是：",captial)
	}else{
		fmt.Println("该国家未收录")
	}

	//delete函数用于删除集合中的元素
	fruitMap := map[string]string{"first":"apple","second":"pear","third":"bannar"}

	fmt.Println("原始Map")

	for fruit := range fruitMap {
		fmt.Println(fruit,"水果是：",fruitMap[fruit])
	}

	//删除元素 通过key值删除
	delete(fruitMap,"first")

	fmt.Println("删除apple")

	for fruit := range fruitMap {
		fmt.Println(fruit,"水果是：",fruitMap[fruit])
	}

}