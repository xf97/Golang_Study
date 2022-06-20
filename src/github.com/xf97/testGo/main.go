package main

import "fmt"

/* var name string
var age int
var is_ok bool */

var (
	name  string
	age   int
	is_ok bool
)


//常量
const (
	ERROR_NONE = 1	 //常量必须初始化
	ERROR_NON = 2
	ERROR_MULTIPLE = 3
	ERROR_SINGLE	//若不初始化，则值与上一行值相同
)

func foo() (int, string) {
	return 10, "Quit"
}

func main() {
	name = "xiaofeng"
	fmt.Println("Hello World")
	is_ok = true
	age = 25
	//var company string = "Ant Group"
	var company = "Ant Group"
	var bigNum = 1e10 //类型自动推导auto，注意基础类型还是尽量指明其类型，防止推导的类型不够大
	fmt.Printf("name:%s, age:%d\n", name, age)
	fmt.Println(is_ok)
	fmt.Println(company)
	fmt.Println(bigNum)
	smallNum := 1e-10 //短变量声明，只能在函数里面用
	fmt.Println(smallNum)
	age, _ = foo() //_是匿名变量，不会分配内存，不会占用命名空间，所以不存在重复声明
	fmt.Println(age)
	fmt.Println(ERROR_NON, ERROR_NONE, ERROR_MULTIPLE, ERROR_SINGLE)
}
