package trace

import (
	"fmt"
	"io"
)

// Traverはコード内での出来事を記録できるオブジェクトを表すインターフェースです
type Tracer interface {
	Trace(...interface{})
}

type tracer struct {
	out io.Writer
}

func (t *tracer) Trace(a ...interface{}) {
	t.out.Write([]byte(fmt.Sprint(a...)))
	t.out.Write([]byte("\n"))
}

type nilTracer struct{}

func (t *nilTracer) Trace(a ...interface{}) {
}

func Off() Tracer {
	return &nilTracer{}
}

func New(w io.Writer) Tracer {
	return &tracer{out: w}
}
