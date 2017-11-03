# Opentracing Logging Implementation

[Opentracing](http://opentracing.io) compatible logging implementation.

## Usage

Configure ot-logfmt as your tracer:

    opentracing.InitGlobalTracer(
            // tracing impl specific:
            ot_logfmt.NewTracer()
        )

Trace and log against the tracer interface:

    span := opentracing.StartSpan("request")
    span.SetTag("http.method", "GET")

    span.LogFields(log.String("event", "finish"), log.String("message", "We have done a request"))

This will have outputed to stdout:

    ts=1509745762970 event=finish message="We have done a request" http.method=GET operation=request

