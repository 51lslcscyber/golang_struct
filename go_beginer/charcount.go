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

// 分别统计一个字符串中的英文字符个数, 数字个数, 空格个数, 其他字符个数
func statisticChars(str string) (charCount, numCount, spcaeCount, otherCount int) {
	charArr := []rune(str)
	for i := 0; i < len(charArr); i++ {
		switch {
		case charArr[i] >= 'a' && charArr[i] <= 'z' || charArr[i] >= 'A' && charArr[i] <= 'Z':
			charCount++
		case charArr[i] >= '0' && charArr[i] <= '9':
			numCount++
		case charArr[i] == ' ':
			spcaeCount++
		default:
			otherCount++
		}
	}
	return
}

func main() {
	str := "abcdefg   23456hh99  uwq你好##"
	charCount, numCount, spcaeCount, otherCount := statisticChars(str)
	fmt.Printf("charCount:%d  numCount:%d  spcaeCount:%d  otherCount:%d", charCount, numCount, spcaeCount, otherCount)
}
