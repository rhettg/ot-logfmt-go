package ot_logfmt

import (
	"bytes"
	"log"

	"github.com/go-logfmt/logfmt"
	opentracing "github.com/opentracing/opentracing-go"
	tlog "github.com/opentracing/opentracing-go/log"
)

type span struct {
	tracer        *tracer
	tagEnc        *logfmt.Encoder
	tagBuf        *bytes.Buffer
	operationName string
}

func newSpan(t *tracer, operationName string) *span {
	tagBuf := &bytes.Buffer{}
	return &span{
		tracer:        t,
		tagEnc:        logfmt.NewEncoder(tagBuf),
		tagBuf:        tagBuf,
		operationName: operationName}
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
	err := s.tagEnc.EncodeKeyval(key, value)
	if err != nil {
		// TODO: Deal with our failures
		panic(err)
	}
	return s
}

func (s *span) LogFields(fields ...tlog.Field) {
	buf := bytes.Buffer{}
	enc := logfmt.NewEncoder(&buf)

	for _, f := range fields {
		err := enc.EncodeKeyval(f.Key(), f.Value())
		if err != nil {
			// TODO: Deal with our failures
			panic(err)
		}
	}

	buf.Write([]byte{' '})
	r := bytes.NewReader(s.tagBuf.Bytes())
	buf.ReadFrom(r)

	log.Println(buf.String())
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
