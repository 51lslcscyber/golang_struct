/*
author:鹄思鹄想_bit森
Time:2022/4/20 23:18
warning:注意事件，无则可以删去
*/
package main

import "fmt"

/*
重构go语言包,影响调试运行，必要时可删去
我们可以看下结果依次是
0,1,2,4,8,16,32,64,128,256,512,1024
但到了 1024 之后,就变成了1024,1280,1696,2304
每次都是扩容了四分之一左右
*/
func main() {
	arr := make([]int, 0)
	for i := 0; i < 2000; i++ {
		fmt.Println("len 为", len(arr), "cap 为", cap(arr))
		arr = append(arr, i)
	}
}
