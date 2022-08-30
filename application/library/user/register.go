package user

import (
	cw "github.com/coscms/webauthn"
	"github.com/coscms/webauthn/static"
	"github.com/duo-labs/webauthn/webauthn"
	"github.com/webx-top/echo"
	"github.com/webx-top/echo/handler/embed"
)

var global = cw.New(handle)

func Init(cfg *webauthn.Config) error {
	return global.Init(cfg)
}

func Register(r echo.RouteRegister) {
	global.RegisterRoute(r)
	fs := embed.NewFileSystems()
	fs.Register(static.JS)
	g := r.Group(`/webauthn`)
	g.Get(`/static/*`, embed.File(fs))
}
