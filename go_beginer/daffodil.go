/*
author:鹄思鹄想_bit森
warning:注意事件，无则可以删去
*/
package main

/*
重构go语言包,影响调试运行，必要时可删去
*/

import (
	"fmt"
)

// 判断n是不是水仙花数
func isDaffodil(n int) bool {
	if n < 100 || n > 999 {
		return false
	}

	first := n % 10
	second := (n / 10) % 10
	third := (n / 100) % 10
	if n == first*first*first+second*second*second+third*third*third {
		return true
	}
	return false
}

func main() {
	for i := 100; i < 1000; i++ {
		if isDaffodil(i) == true {
			fmt.Printf("%d is daffodil\n", i)
		}
	}
}
