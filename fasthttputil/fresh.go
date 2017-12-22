package fasthttputil

import (
	"errors"
	"github.com/valyala/fasthttp"
	"github.com/zc310/headers"
	"strings"
	"time"
)

// IsFresh check whether cache can be used in this fasthttp RequestCtx
func IsFresh(ctx *fasthttp.RequestCtx) bool {
	respCacheControl := ParseCacheControl(string(ctx.Response.Header.Peek(headers.CacheControl)))
	reqCacheControl := ParseCacheControl(string(ctx.Request.Header.Peek(headers.CacheControl)))
	if _, ok := reqCacheControl[headers.NoCache]; ok {
		return false
	}
	if _, ok := respCacheControl[headers.NoCache]; ok {
		return false
	}
	if _, ok := reqCacheControl[headers.OnlyIfCached]; ok {
		return true
	}

	date, err := ResponseDate(ctx)
	if err != nil {
		return false
	}
	currentAge := time.Since(date)

	var lifetime time.Duration
	var zeroDuration time.Duration

	// If a response includes both an Expires header and a max-age directive,
	// the max-age directive overrides the Expires header, even if the Expires header is more restrictive.
	if maxAge, ok := respCacheControl[headers.MaxAge]; ok {
		lifetime, err = time.ParseDuration(maxAge + "s")
		if err != nil {
			lifetime = zeroDuration
		}
	} else {
		expiresHeader := string(ctx.Response.Header.Peek(headers.Expires))
		if expiresHeader != "" {
			expires, err := time.Parse(time.RFC1123, expiresHeader)
			if err != nil {
				lifetime = zeroDuration
			} else {
				lifetime = expires.Sub(date)
			}
		}
	}

	if maxAge, ok := reqCacheControl[headers.MaxAge]; ok {
		// the client is willing to accept a response whose age is no greater than the specified time in seconds
		lifetime, err = time.ParseDuration(maxAge + "s")
		if err != nil {
			lifetime = zeroDuration
		}
	}
	if minfresh, ok := reqCacheControl[headers.MinFresh]; ok {
		//  the client wants a response that will still be fresh for at least the specified number of seconds.
		minfreshDuration, err := time.ParseDuration(minfresh + "s")
		if err == nil {
			currentAge = time.Duration(currentAge + minfreshDuration)
		}
	}

	if maxstale, ok := reqCacheControl[headers.MaxStale]; ok {
		// Indicates that the client is willing to accept a response that has exceeded its expiration time.
		// If max-stale is assigned a value, then the client is willing to accept a response that has exceeded
		// its expiration time by no more than the specified number of seconds.
		// If no value is assigned to max-stale, then the client is willing to accept a stale response of any age.
		//
		// Responses served only because of a max-stale value are supposed to have a Warning header added to them,
		// but that seems like a  hassle, and is it actually useful? If so, then there needs to be a different
		// return-value available here.
		if maxstale == "" {
			return true
		}
		maxstaleDuration, err := time.ParseDuration(maxstale + "s")
		if err == nil {
			currentAge = time.Duration(currentAge - maxstaleDuration)
		}
	}

	if lifetime > currentAge {
		return true
	}

	return false
}

type CacheControl map[string]string

func ParseCacheControl(headers string) CacheControl {
	cc := CacheControl{}

	for _, part := range strings.Split(headers, ",") {
		part = strings.Trim(part, " ")
		if part == "" {
			continue
		}
		if strings.ContainsRune(part, '=') {
			keyval := strings.Split(part, "=")
			cc[strings.Trim(keyval[0], " ")] = strings.Trim(keyval[1], ",")
		} else {
			cc[part] = ""
		}
	}
	return cc
}

// Date parses and returns the value of the Date header.
func ResponseDate(ctx *fasthttp.RequestCtx) (date time.Time, err error) {
	dateHeader := string(ctx.Response.Header.Peek(headers.Date))
	if dateHeader == "" {
		err = errors.New("no Date header")
		return
	}

	return time.Parse(time.RFC1123, dateHeader)
}
func GetResponseAge(ctx *fasthttp.RequestCtx, def time.Duration) (age time.Duration, cache bool) {
	respCacheControl := ParseCacheControl(string(ctx.Response.Header.Peek(headers.CacheControl)))
	if _, ok := respCacheControl[headers.NoCache]; ok {
		return def, false
	}

	date, err := ResponseDate(ctx)
	if err != nil {
		return def, true
	}
	if maxAge, ok := respCacheControl[headers.MaxAge]; ok {
		age, err = time.ParseDuration(maxAge + "s")
		if err != nil {
			return def, true
		}
		return
	} else {
		expiresHeader := string(ctx.Response.Header.Peek(headers.Expires))
		if expiresHeader != "" {
			expires, err := time.Parse(time.RFC1123, expiresHeader)
			if err != nil {
				return def, true
			} else {
				return expires.Sub(date), true
			}
		}
	}
	return def, true
}
