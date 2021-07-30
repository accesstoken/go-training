package log

import (
	"fmt"
	"io"
)

type Logger struct {
	Writer io.Writer
}

func (receiver Logger) Log(a ...interface{}) {
	fmt.Fprintln(receiver.Writer,  a...)
}

func (receiver Logger) Begin() {
	receiver.Log("BEGIN")
}

func (receiver Logger) End() {
	receiver.Log("END")
}
