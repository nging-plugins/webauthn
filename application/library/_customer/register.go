package customer

import (
	"net/url"

	"github.com/admpub/log"
	cw "github.com/coscms/webauthn"
	"github.com/coscms/webauthn/static"
	"github.com/duo-labs/webauthn/webauthn"
	"github.com/webx-top/com"
	"github.com/webx-top/echo"
	"github.com/webx-top/echo/handler/embed"
	"github.com/webx-top/echo/param"

	"github.com/admpub/nging/v4/application/cmd/bootconfig"
	"github.com/admpub/nging/v4/application/library/config"
	"github.com/admpub/webx/application/library/top"
)

var global = cw.New(handle)

func Init(cfg *webauthn.Config) error {
	return global.Init(cfg)
}

// init webauthn
func initWebAuthn(ctx echo.Context) *webauthn.Config {
	siteURL := top.SiteURL(ctx)
	if len(backendURL) == 0 {
		siteURL = `http://localhost:` + param.AsString(config.FromCLI().Port)
	}
	icon := siteURL + `/public/assets/backend/images/logo.png`
	u, _ := url.Parse(siteURL)
	cfg := &webauthn.Config{
		RPDisplayName: bootconfig.SoftwareName, // Display Name for your site
		RPID:          com.SplitHost(u.Host),   // Generally the domain name for your site
		RPOrigin:      siteURL,                 // The origin URL for WebAuthn requests
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
