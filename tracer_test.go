package ot_logfmt

import (
	"os"
	"testing"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
)

func TestNoopInterfaces(t *testing.T) {
	tr := NewTracer()

	var i struct{}
	err := tr.Inject(defaultNoopSpanContext, i, i)
	if err != nil {
		t.Fatal("Should not have error", err)
	}

	ctx, err := tr.Extract(i, i)
	if err != nil {
		t.Fatal("Should not have error")
	}

	if ctx != defaultNoopSpanContext {
		t.Fatal("Should be default span context")
	}
}

func TestGlobalTracer(t *testing.T) {
	opentracing.InitGlobalTracer(NewTracer())

	idx := 0
	s := opentracing.StartSpan("loop")

	s.SetTag("foo", "bar")
	s.SetTag("pid", os.Getpid())
	s.SetTag("ppid", os.Getppid())

	s.LogFields(log.String("event", "finish"), log.String("message", "One Loop"), log.Int("idx", idx))
	s.Finish()

	idx += 1
}

func BenchmarkSpanStart(b *testing.B) {
	t := NewTracer()

	for n := 0; n < b.N; n++ {
		t.StartSpan("benchmark")
	}
}
