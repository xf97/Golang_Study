package main

import(
	"fmt"
	"math"
)

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
	ERROR_NONE     = 1 //常量必须初始化
	ERROR_NON      = 2
	ERROR_MULTIPLE = 3
	ERROR_SINGLE   //若不初始化，则值与上一行值相同
)

//用常量和iota声明枚举
const (
	TYPE_1 = iota //0，iota是常量计数器，每新增一个常量就+1，常用于声明枚举
	TYPE_2        //1
	_             //使用匿名变量跳过某些值
	TYPE_3        //2
)

const (
	ERROR_1 = iota //会被const再出现的时候，被重置为0，iota相当于行索引
	ERROR_2 = 100  //特殊值
	ERROR_3 = iota //2
)

//定义数量级
const (
	_  = iota //跳过0
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
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
	fmt.Println(TYPE_1, TYPE_2, TYPE_3)
	fmt.Println(ERROR_1, ERROR_2, ERROR_3)
	fmt.Println(KB, MB, GB, TB, PB)
	//可以用uint8表示bytes
	//%d表示10进制，%b表示二进制，%o八进制，%x十六进制
	//无法直接定义二进制数，八进制以0开头，十六进制以0x开头
	var a int = 10
	fmt.Printf("%d %b\n", a, a)
	b := 077
	c := 0xfe
	fmt.Printf("%o %x\n", b, c)
	//可以在format中直接转换进制
	//默认go语言的小数使用float64
	//%T查看变量类型，%v打印值，%s是字符串
	fmt.Printf("%T %T %T\n", name, bigNum, smallNum)
	fmt.Println(math.MaxFloat32, math.MaxFloat64)
	fmt.Printf("%.2f\n", math.Pi)
	//复数
	var c1 complex64 = 1 + 2i
	fmt.Println(c1)
	//bool的默认值为false
}
