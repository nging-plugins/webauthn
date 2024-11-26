package user

import (
	"net/url"

	"github.com/admpub/log"
	cw "github.com/coscms/webauthn"
	"github.com/coscms/webauthn/static"
	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/webx-top/com"
	"github.com/webx-top/echo"
	"github.com/webx-top/echo/handler/embed"
	"github.com/webx-top/echo/param"

	"github.com/coscms/webcore/cmd/bootconfig"
	"github.com/coscms/webcore/library/common"
	"github.com/coscms/webcore/library/config"
	"github.com/coscms/webcore/library/httpserver"
)

var global = cw.New(handle, initWebAuthn)

func Init(cfg *webauthn.Config) error {
	log.Debugf(`webauthn.Config: %+v`, cfg)
	return global.Init(cfg)
}

// init webauthn
func initWebAuthn(ctx echo.Context) *webauthn.Config {
	backendURL := common.BackendURL(ctx)
	if len(backendURL) == 0 {
		backendURL = `http://localhost:` + param.AsString(config.FromCLI().Port)
	}
	u, _ := url.Parse(backendURL)
	cfg := &webauthn.Config{
		RPDisplayName: bootconfig.SoftwareName, // Display Name for your site
		RPID:          com.SplitHost(u.Host),   // Generally the domain name for your site
		RPOrigins:     []string{backendURL},    // The origin URL for WebAuthn requests
	}
	log.Debugf(`webauthn.Config: %+v`, cfg)
	return cfg
}

func RegisterBackend(r echo.RouteRegister) {
	g := r.Group(`/webauthn`)
	httpserver.SetGroupMetaPermissionPublic(g) // 登录用户 group
	global.RegisterRoute(r)
	fs := embed.NewFileSystems()
	fs.Register(static.JS)
	g.Get(`/static/*`, embed.File(fs))
}

func RegisterLogin(r echo.RouteRegister) {
	g := r.Group(`/webauthn`)
	httpserver.SetGroupMetaPermissionGuest(g) // 匿名 group
	global.RegisterRouteForLogin(r)
	fs := embed.NewFileSystems()
	fs.Register(static.JS)
	g.Get(`/static/*`, embed.File(fs))
}
