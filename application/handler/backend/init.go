package backend

import "github.com/webx-top/echo"

func Register(r echo.RouteRegister) {
	g := r.Group(`/user`)
	g.Route(`POST,GET`, `/webauthn`, WebAuthn)
}
