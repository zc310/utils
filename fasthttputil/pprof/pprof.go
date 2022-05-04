package pprof

import (
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"

	"net/http/pprof"
	"strings"
)

var (
	cmdline = fasthttpadaptor.NewFastHTTPHandlerFunc(pprof.Cmdline)
	profile = fasthttpadaptor.NewFastHTTPHandlerFunc(pprof.Profile)
	symbol  = fasthttpadaptor.NewFastHTTPHandlerFunc(pprof.Symbol)
	trace   = fasthttpadaptor.NewFastHTTPHandlerFunc(pprof.Trace)
	index   = fasthttpadaptor.NewFastHTTPHandlerFunc(pprof.Index)
)

func Handler(ctx *fasthttp.RequestCtx) {
	if strings.HasSuffix(string(ctx.Path()), "/pprof/cmdline") {
		cmdline(ctx)
	} else if strings.HasSuffix(string(ctx.Path()), "/pprof/profile") {
		profile(ctx)
	} else if strings.HasSuffix(string(ctx.Path()), "/pprof/symbol") {
		symbol(ctx)
	} else if strings.HasSuffix(string(ctx.Path()), "/pprof/trace") {
		trace(ctx)
	} else {
		ctx.Response.Header.SetContentType("text/html")
		index(ctx)
	}
}
