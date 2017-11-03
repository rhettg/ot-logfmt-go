package ot_logfmt

import (
	"testing"

	"github.com/opentracing/opentracing-go/log"
)

func BenchmarkLifecycle(b *testing.B) {
	t := NewTracer()

	for n := 0; n < b.N; n++ {
		s := t.StartSpan("benchmark")
		s.SetTag("foo", "bar")
		s.SetTag("count", n)
		s.SetTag("error", false)
		s.LogFields(log.String("event", "finish"), log.String("message", "I'm doing some benchmarking"))
		s.Finish()
	}
}
