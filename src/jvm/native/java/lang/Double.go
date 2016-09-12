package lang

import (
	"jvm/native"
	"jvm/rtda"
	"math"
)

func init() {
	native.Register("java/lang/Double", "doubleToRawLongBits", "(D)J", doubleToRawLongBits)
	native.Register("java/lang/Double", "longBitsToDouble", "(J)D", longBitsToDouble)
}

func doubleToRawLongBits(frame *rtda.Frame) {
	value := frame.LocalVars().GetDouble(0)
	bits := math.Float64bits(value)
	frame.OperandStack().PushLong(int64(bits))
}

func longBitsToDouble(frame *rtda.Frame) {
	bits := frame.LocalVars().GetLong(0)
	value := math.Float64frombits(uint64(bits))
	frame.OperandStack().PushDouble(value)
}
