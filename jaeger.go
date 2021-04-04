package jaeger

import (
	"github.com/opentracing/opentracing-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
)

// Tracer
var Tracer opentracing.Tracer

// NewJaegerClient 新建 Jaeger 客户端
func NewJaegerClient(config *Config) {
	// jaeger 设置项
	cfg := jaegercfg.Configuration{
		ServiceName: config.Name,
		Disabled:    !config.Enable,
		Sampler: &jaegercfg.SamplerConfig{
			Type:  "probabilistic", // 自定义采样
			Param: config.Sampler,  // 设置采样率
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:           false,
			LocalAgentHostPort: config.Endpoint,
		},
	}

	tracer, _, err := cfg.NewTracer(
		jaegercfg.Logger(jaegerlog.StdLogger),
		jaegercfg.MaxTagValueLength(512),
	)

	if err == nil {
		Tracer = tracer
		opentracing.SetGlobalTracer(tracer)
	}
}
