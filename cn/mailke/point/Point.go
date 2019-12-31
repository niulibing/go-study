package main

import "fmt"

func main() {

	var a = 11
	//1.通过在变量名前加&来获取变量的地址
	//2.*i出现在"="左边表示指针声明，*i出现在"="右边表示取指针指向的值
	i := &a
	fmt.Printf("a的值为【 %v】\r\n", a)
	fmt.Printf("i的值为【 %v】\r\n", *i)

	//3.结构体指针访问结构体字段仍使用"."点操作

	type User struct {
		name string
		age  int
	}
	user := User{name: "牛娃娃", age: 26}
	u := &user
	fmt.Printf("用户的名称为[%v]\r\n", u.name)
	fmt.Printf("用户的年龄为[%v]\r\n", u.age)

}
