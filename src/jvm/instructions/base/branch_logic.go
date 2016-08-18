package base

import (
	"jvm/rtda"
)

/*
 * 分支逻辑, 跳转
 */
func Branch(frame *rtda.Frame, offset int) {
	pc := frame.Thread().PC()
	nextPC := pc + offset
	frame.SetNextPC(nextPC)
}
