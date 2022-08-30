package webauthn

import (
	"github.com/admpub/nging/v4/application/library/module"
	"github.com/admpub/nging/v4/application/library/route"

	"github.com/nging-plugins/webauthn/application/handler/backend"
	//"github.com/nging-plugins/webauthn/application/library/customer"
	"github.com/nging-plugins/webauthn/application/library/user"
)

const ID = `webauthn`

var Module = module.Module{
	TemplatePath: map[string]string{
		ID: `webauthn/template/backend`,
	},
	AssetsPath: []string{},
	//Navigate: ,
	Route: func(r *route.Collection) {
		user.Register(r.Backend.Echo().Group(`/user`))
		backend.Register(r.Backend.Echo())
		//customer.Register(r.Frontend.Echo())
	},
	DBSchemaVer: 0.0000,
}
