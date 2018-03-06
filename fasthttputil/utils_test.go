package fasthttputil

import (
	"bufio"
	"bytes"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
	"testing"
)

func TestGetParameters(t *testing.T) {
	var ctx fasthttp.RequestCtx

	s := "GET /?a=a HTTP/1.1\nHost: aaa.com\n\n"
	br := bufio.NewReader(bytes.NewBufferString(s))
	if err := ctx.Request.Read(br); err != nil {
		t.Fatalf("cannot read request: %s", err)
	}
	assert.Equal(t, []byte("a"), GetArgs(&ctx, "a"))
}

func TestOk(t *testing.T) {
	var ctx fasthttp.RequestCtx
	Ok(&ctx)
	assert.Equal(t, "ok\n", string(ctx.Response.Body()))
}
