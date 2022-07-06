/*
author:鹄思鹄想_bit森
Time:2022/7/5 23:52
warning:注意事件，无则可以删去
*/

/*
重构go语言包,影响调试运行，必要时可删去
*/
//拼接字符串
package main

import "fmt"

func main() {

	//创建和初始化字符串

	//使用var关键字
	var str1 string
	str1 = "Welcome!"

	var str2 string
	str2 = "(https://github.com/51lslcscyber.com)"

	//连接字符串
	//使用+运算符
	fmt.Println("新字符串 1: ", str1+str2)

	//创建和初始化字符串
	//使用简写声明
	str3 := "Geeks"
	str4 := "Geeks"

	//连接字符串
	//使用+运算符
	result := str3 + "for" + str4

	fmt.Println("新字符串 2: ", result)

}
