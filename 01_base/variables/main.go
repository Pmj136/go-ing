package variables

import (
	"fmt"
	"strconv"
)

func Type1() {
	var num int
	fmt.Println("int 默认是: " + strconv.Itoa(num))
	var a string
	fmt.Println("字符串默认是：" + a)
	var flag bool
	fmt.Println("bool 默认是：" + strconv.FormatBool(flag))
}
func Type2() {
	var name string = "彭明久"
	fmt.Println(name)
}

func Type3() {
	//自动判断类型
	var name = "彭明久"
	fmt.Println(name)
}

func Type4() {
	//如果已经声明过了 再使用 := 就会编译出错
	/*var name := "彭明久"*/

	/*var age int
	age := 1*/

	//正确的做法
	age := 1
	fmt.Println(age)

	name := "彭明久"
	fmt.Println(name)
}

func Type5() {
	//类型相同的多个变量
	var year, month int = 2021, 12
	fmt.Println(strconv.Itoa(year) + "," + strconv.Itoa(month))

	x, y, z := 1, 2, "aaaaaaa"
	fmt.Print(x)
	fmt.Print(y)
	fmt.Println(z)
}

// Type6 值类型和引用类型
func Type6() {
	a := 1
	fmt.Println(&a) //取内存地址
}
