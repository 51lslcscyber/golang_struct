/*
author:鹄思鹄想_bit森
Time:2022/5/11 22:46
warning:注意事件，无则可以删去
*/
package main

import (
	_ "fmt"
)

type PolicyType int32

const (
	Policy_MIN PolicyType = 0
	Policy_MAX PolicyType = 1
	Policy_MID PolicyType = 2
	Policy_AVG PolicyType = 3
)
