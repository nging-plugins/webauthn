package webauthn

import (
	"github.com/admpub/nging/v4/application/library/module"
	"github.com/admpub/nging/v4/application/library/route"

	"github.com/nging-plugins/webauthn/pkg/library/customer"
	"github.com/nging-plugins/webauthn/pkg/library/user"
)

const ID = `webauthn`

var Module = module.Module{
	TemplatePath: map[string]string{
		ID: `webauthn/template/backend`,
	},
	AssetsPath: []string{},
	//Navigate: ,
	Route: func(r *route.Collection) {
		user.Register(r.Backend.Echo())
		customer.Register(r.Frontend.Echo())
	},
	DBSchemaVer: 0.0000,
}
