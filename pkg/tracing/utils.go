package tracing

import "github.com/opentracing/opentracing-go"

func TraceErr(span opentracing.Span, err error) {
	span.SetTag("error", true)
	span.LogKV("error_code", err.Error())
}

func InjectTextMapCarrier(spanCtx opentracing.SpanContext) (opentracing.TextMapCarrier, error) {
	m := make(opentracing.TextMapCarrier)
	if err := opentracing.GlobalTracer().Inject(spanCtx, opentracing.TextMap, m); err != nil {
		return nil, err
	}
	return m, nil
}

func ExtractTextMapCarrier(spanCtx opentracing.SpanContext) opentracing.TextMapCarrier {
	textMapCarrier, err := InjectTextMapCarrier(spanCtx)
	if err != nil {
		return make(opentracing.TextMapCarrier)
	}
	return textMapCarrier
}
