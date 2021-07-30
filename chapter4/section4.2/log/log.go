package log

import (
	"fmt"
	"io"
)

type Logger struct {
	Writer io.Writer
}

func (receiver Logger) Logf(format string, a ...interface{}) {
	fmt.Fprintf(receiver.Writer,  format + "\n", a...)
}

func (receiver Logger) Log(a ...interface{}) {
	var s string = fmt.Sprint(a...)
	receiver.Logf("%s", s)
}


func (receiver Logger) Begin() {
	receiver.Log("BEGIN")
}

func (receiver Logger) End() {
	receiver.Log("END")
}
