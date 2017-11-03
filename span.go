package ot_logfmt

import (
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
)

type span struct {
	tracer        *tracer
	operationName string
}

func (s *span) Finish() {}

func (s *span) FinishWithOptions(opts opentracing.FinishOptions) {}

func (s *span) Context() opentracing.SpanContext {
	return defaultNoopSpanContext
}

func (s *span) SetOperationName(operationName string) opentracing.Span {
	s.operationName = operationName
	return s
}

func (s *span) SetTag(key string, value interface{}) opentracing.Span {
	return s
}

func (s *span) LogFields(fields ...log.Field) {

}

func (s *span) LogKV(alternatingKeyValues ...interface{}) {

}

func (s *span) SetBaggageItem(restrictedKey, value string) opentracing.Span {
	return s
}

func (s *span) BaggageItem(restrictedKey string) string {
	return ""
}

func (s *span) Tracer() opentracing.Tracer {
	return s.tracer
}

// Deprecated in the standard, not implemented here.
func (s *span) LogEvent(event string) {
	panic("not implemented")
}

func (s *span) LogEventWithPayload(event string, payload interface{}) {
	panic("not implemented")
}

func (s *span) Log(data opentracing.LogData) {
	panic("not implemented")
}
