package main

import (
	"log"
	"os"
	"time"

	"github.com/opentracing/opentracing-go"
	tlog "github.com/opentracing/opentracing-go/log"
)

func main() {
	opentracing.InitGlobalTracer(opentracing.NoopTracer{})

	idx := 0
	for {
		log.Printf("One Loop")
		s := opentracing.StartSpan("loop")

		s.SetTag("pid", os.Getpid())
		s.SetTag("ppid", os.Getppid())

		time.Sleep(1 * time.Second)
		s.LogFields(tlog.String("event", "finish"), tlog.String("message", "One Loop"), tlog.Int("idx", idx))
		s.Finish()

		time.Sleep(1 * time.Second)
		idx += 1
	}
}
