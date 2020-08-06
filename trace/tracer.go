package trace

import (
	"fmt"
	"io"
)

// Tracer はコード内での出来事を記録できるオブジェクトを表すインタフェースです。
type Tracer struct {
	out io.Writer
}

// type Tracer interface {
// 	Trace(...interface{})
// }

// type tracer struct {
// 	out io.Writer
// }

// New はTracerを返します。
func New(w io.Writer) *Tracer {
	return &Tracer{out: w}
}

// func New(w io.Writer) Tracer {
// 	return &tracer{out: w}
// }

// Trace はコード内での出来事を記録します。
func (t *Tracer) Trace(a ...interface{}) {
	if t == nil || t.out == nil {
		return
	}
	fmt.Fprintln(t.out, a...)
}

// func (t *tracer) Trace(a ...interface{}) {
// 	t.out.Write([]byte(fmt.Sprint(a...)))
// 	t.out.Write([]byte("\n"))
// }
