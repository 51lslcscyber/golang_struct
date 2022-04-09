/*
author:鹄思鹄想_bit森
*/

/*
重构go语言包,影响调试运行，必要时可删去
*/
package main

import (
	"fmt"
)

// 判断n是不是质数
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}

	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	// 求100以内的所有质数
	for i := 2; i <= 100; i++ {
		if isPrime(i) == true {
			fmt.Printf("%d is prime\n", i)
		}
	}
}
