package httputil

import "net/http"

var UserAgent = "Mozilla/5.0 (Windows NT 6.3; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/47.0.2526.27 Safari/537.36"

func init() {
	http.DefaultTransport = &UserAgentTransport{http.DefaultTransport}
}

type UserAgentTransport struct {
	rt http.RoundTripper
}

func (p UserAgentTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	r.Header.Set("User-Agent", UserAgent)
	return p.rt.RoundTrip(r)
}
