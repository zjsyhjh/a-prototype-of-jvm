package rtda

import (
	"jvm/rtda/heap"
)

/*
 * 一个可以同时容纳一个int值和一个引用值的结构体
 */
type Slot struct {
	num       int32
	reference *heap.Object
}
