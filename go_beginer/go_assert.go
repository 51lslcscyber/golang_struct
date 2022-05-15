/*
author:鹄思鹄想_bit森
Time:2022/5/15 12:20
warning:注意事件，无则可以删去
*/
package main

import (
	_ "fmt"
)

type PolicyType int32

const (
	PolicyMin PolicyType = 0
	PolicyMax PolicyType = 1
	PolicyMid PolicyType = 2
	PolicyAvg PolicyType = 3
)
