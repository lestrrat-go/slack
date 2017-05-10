package slack

import (
	"context"
	"fmt"
	"io"
	"strings"
)

type nilLogger struct{}

func (_ nilLogger) Debugf(_ context.Context, _ string, _ ...interface{}) {}
func (_ nilLogger) Infof(_ context.Context, _ string, _ ...interface{})  {}

type traceLogger struct {
	dst io.Writer
}

func (l traceLogger) Debugf(_ context.Context, f string, args ...interface{}) {
	if !strings.HasSuffix(f, "\n") {
		f = f + "\n"
	}
	fmt.Fprintf(l.dst, f, args...)
}

func (l traceLogger) Infof(_ context.Context, f string, args ...interface{}) {
	if !strings.HasSuffix(f, "\n") {
		f = f + "\n"
	}
	fmt.Fprintf(l.dst, f, args...)
}
