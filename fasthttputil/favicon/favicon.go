package favicon

import (
	"github.com/disintegration/letteravatar"
	"github.com/valyala/fasthttp"
	"image"
	"image/png"
)

// Handler favicon.ico Handler
func Handler(ctx *fasthttp.RequestCtx) {
	img, err := letteravatar.Draw(32, 'F', nil)
	if err != nil {
		img = image.NewGray(image.Rect(0, 0, 32, 32))
	}
	png.Encode(ctx.Response.BodyWriter(), img)
}
