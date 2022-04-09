/*
author:鹄思鹄想_bit森
warning:注意事件，无则可以删去
*/

/*
重构go语言包,影响调试运行，必要时可删去
*/
package main

import (
	"fmt"
)

func main() {
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d\t", j, i, i*j)
		}
		fmt.Println()
	}
}
