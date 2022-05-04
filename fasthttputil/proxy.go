package fasthttputil

import (
	"fmt"
	"net/http"
	"net/url"
	"sync/atomic"
	"time"

	"github.com/valyala/fasthttp"
)

var (
	XForwardedFor  = http.CanonicalHeaderKey("X-Forwarded-For")
	XForwardedHost = http.CanonicalHeaderKey("X-Forwarded-Host")
	XRealIP        = http.CanonicalHeaderKey("X-Real-IP")
)

type Target struct {
	Name     string `json:"name,omitempty"`
	URL      string `json:"url"`
	MaxConns int    `json:"max_conns"`
	Status   string `json:"status"`
	Host     string `json:"host"`
}
type Proxy struct {
	client *fasthttp.HostClient
	status string

	MaxConns int32
	Conns    int32
	Name     string
	Host     []byte
	Fails    int32
}

func NewProxyClient(to *Target) (*Proxy, error) {
	var status string
	status = to.Status
	url, err := url.Parse(to.Status)
	if err != nil {
		status = to.URL
	}
	if status == "" {
		status = fmt.Sprintf("http://%s", to.Host)
	}
	url, err = url.Parse(to.URL)
	if err != nil {
		return nil, err
	}
	if to.Name == "" {
		to.Name = to.URL
	}
	return &Proxy{client: &fasthttp.HostClient{Addr: url.Host,
		MaxConns:            to.MaxConns,
		MaxIdleConnDuration: time.Minute * 2,
	}, status: status, Name: to.Name, Host: []byte(to.Host)}, err
}

func (p *Proxy) Handler(ctx *fasthttp.RequestCtx) (err error) {
	err = p.Do(&ctx.Request, &ctx.Response)
	if err != nil {
		atomic.AddInt32(&p.Fails, 1)
	}
	return
}
func (p *Proxy) Do(req *fasthttp.Request, resp *fasthttp.Response) (err error) {
	req.Header.Del("Connection")
	req.Header.SetBytesV(XForwardedHost, req.Host())
	if len(p.Host) > 0 {
		req.SetHostBytes(p.Host)
	}
	err = p.client.Do(req, resp)
	resp.Header.Del("Connection")
	return
}

// Full checks whether the upstream host has reached its maximum connections
func (p Proxy) Full() bool {
	return p.MaxConns > 0 && atomic.LoadInt32(&p.Conns) >= p.MaxConns
}
func (p *Proxy) Down() bool {
	return atomic.LoadInt32(&p.Fails) > 0
}
func (p Proxy) Available() bool {
	return !p.Down() && !p.Full()
}
func (p Proxy) StatusOK() bool {
	a, _, err := p.client.Get(nil, p.status)
	return err == nil && a == 200
}
func (p *Proxy) HealthCheck() {
	if !p.StatusOK() {
		atomic.AddInt32(&p.Fails, 1)
	}
}
