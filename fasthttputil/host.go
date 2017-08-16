package fasthttputil

import (
	"fmt"
	"strings"

	"github.com/valyala/fasthttp"
)

// HostSwitch is the host-handler map
type HostSwitch map[string]fasthttp.RequestHandler

// Handler for processing incoming requests.
func (p HostSwitch) Handler(ctx *fasthttp.RequestCtx) {
	h, ok := p[string(ctx.Host())]
	if ok {
		h(ctx)
		return
	}
	if h, ok = p["*"]; ok {
		h(ctx)
	} else {
		ctx.Error(fmt.Sprintf("%s Forbidden", string(ctx.Host())), 403)
	}
}

// Add add host Handler
func (p HostSwitch) Add(host string, h fasthttp.RequestHandler) {
	for _, s := range strings.Split(host, ",") {
		p[s] = h
	}
}
