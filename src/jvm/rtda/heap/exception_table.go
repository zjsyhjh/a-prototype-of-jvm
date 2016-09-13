package heap

import (
	"jvm/classfile"
)

type ExceptionHandler struct {
	startPc   int
	endPc     int
	handlerPc int
	catchType *ClassRef
}

type ExceptionTable []*ExceptionHandler

func newExceptionTable(entries []*classfile.ExceptionTableEntry, cp *ConstantPool) ExceptionTable {
	table := make([]*ExceptionHandler, len(entries))
	for i, entry := range entries {
		table[i] = &ExceptionHandler{
			startPc:   int(entry.StartPc()),
			endPc:     int(entry.EndPc()),
			handlerPc: int(entry.HandlePc()),
			catchType: getCatchType(uint(entry.CatchType()), cp),
		}
	}
	return table
}

/*
 * 从常量池中查找类符号引用
 */
func getCatchType(index uint, cp *ConstantPool) *ClassRef {
	if index == 0 {
		// catch all
		return nil
	}
	return cp.GetConstant(index).(*ClassRef)
}

/*
 * startPc是try块的第一条指令，endPc是try块的下一条指令
 */
func (self ExceptionTable) findExceptionHandler(exClass *Class, pc int) *ExceptionHandler {
	for _, handler := range self {
		if pc >= handler.startPc && pc < handler.endPc {
			if handler.catchType == nil {
				return handler
			}

			catchClass := handler.catchType.ResolvedClass()
			if catchClass == exClass || catchClass.IsSuperClassOf(exClass) {
				return handler
			}
		}
	}
	return nil
}
