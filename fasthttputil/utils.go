package fasthttputil

import (
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
	"github.com/zc310/headers"
	"strings"
)

func Ok(ctx *fasthttp.RequestCtx) {
	fmt.Fprint(ctx, "ok\n")
}

// GetArgs   Post->Get->Header->Cookie
func GetArgs(ctx *fasthttp.RequestCtx, key string) []byte {
	b := ctx.Request.PostArgs().Peek(key)
	if len(b) == 0 {
		b = ctx.QueryArgs().Peek(key)
		if len(b) == 0 {
			b = ctx.Request.Header.Peek(key)
			if len(b) == 0 {
				b = ctx.Request.Header.Cookie(key)
			}
		}
	}
	return b
}

// StripPrefix returns a handler that serves HTTP requests
// by removing the given prefix from the request URL's Path
// and invoking the handler h. StripPrefix handles a
// request for a path that doesn't begin with prefix by
// replying with an HTTP 404 not found error.
func StripPrefix(prefix string, h fasthttp.RequestHandler) fasthttp.RequestHandler {
	if prefix == "" {
		return h
	}
	return func(ctx *fasthttp.RequestCtx) {
		if p := strings.TrimPrefix(string(ctx.Path()), prefix); len(p) < len(ctx.Path()) {
			ctx.URI().SetPath(p)
			h(ctx)
		} else {
			ctx.Error(fasthttp.StatusMessage(fasthttp.StatusNotFound), fasthttp.StatusNotFound)
		}
	}
}

func JSON(ctx *fasthttp.RequestCtx, v interface{}) error {
	ctx.Response.Header.Set(headers.ContentType, "application/json; charset=UTF-8")
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}
	_, err = ctx.Write(b)
	return err
}
