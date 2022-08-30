package user

import (
	"net/url"

	"github.com/admpub/log"
	cw "github.com/coscms/webauthn"
	"github.com/coscms/webauthn/static"
	"github.com/duo-labs/webauthn/webauthn"
	"github.com/webx-top/echo"
	"github.com/webx-top/echo/handler/embed"
	"github.com/webx-top/echo/param"

	"github.com/admpub/nging/v4/application/cmd/bootconfig"
	"github.com/admpub/nging/v4/application/library/common"
	"github.com/admpub/nging/v4/application/library/config"
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
	icon := backendURL + `/public/assets/backend/images/logo.png`
	u, _ := url.Parse(backendURL)
	cfg := &webauthn.Config{
		RPDisplayName: bootconfig.SoftwareName, // Display Name for your site
		RPID:          u.Host,                  // Generally the domain name for your site
		RPOrigin:      backendURL,              // The origin URL for WebAuthn requests
		RPIcon:        icon,                    // Optional icon URL for your site
	}
	log.Debugf(`webauthn.Config: %+v`, cfg)
	return cfg
}

func Register(r echo.RouteRegister) {
	global.RegisterRoute(r)
	fs := embed.NewFileSystems()
	fs.Register(static.JS)
	g := r.Group(`/webauthn`)
	g.Get(`/static/*`, embed.File(fs))
}
