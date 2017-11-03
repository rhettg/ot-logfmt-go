package ot_logfmt

import (
	opentracing "github.com/opentracing/opentracing-go"
)

type noopSpanContext struct{}

func (n noopSpanContext) ForeachBaggageItem(handler func(k, v string) bool) {}

var defaultNoopSpanContext = noopSpanContext{}

type tracer struct{}

func (t *tracer) StartSpan(operationName string, opts ...opentracing.StartSpanOption) opentracing.Span {
	// TODO: propogate baggage

	return newSpan(t, operationName)
}

func (t *tracer) Inject(sm opentracing.SpanContext, format interface{}, carrier interface{}) error {
	// No-op
	return nil
}

func (t *tracer) Extract(format interface{}, carrier interface{}) (opentracing.SpanContext, error) {
	// No-op
	return defaultNoopSpanContext, nil
}

func NewTracer() opentracing.Tracer {
	return &tracer{}
}
