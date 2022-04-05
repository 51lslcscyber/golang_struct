/*
author:鹄思鹄想_bit森
warning:注意事件，无则可以删去
Time:2022/4/5 10:30
*/

/*
重构go语言包,影响调试运行，必要时可删去
*/
package main

import "fmt"

type boby struct {
	id   int
	name string
	age  int
}

var b boby

func main() {
	b.id = 0
	b.name = "minmin"
	b.age = 10
	fmt.Printf("The boby's name is %s", b.name)
}
