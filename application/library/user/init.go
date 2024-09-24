package user

import (
	"github.com/coscms/webcore/library/dashboard"
	"github.com/coscms/webcore/library/httpserver"
	"github.com/coscms/webcore/model"
)

func init() {
	model.RegisterSafeItem(`webauthn`, `免密登录`, model.SafeItemInfo{
		Step: 1, ConfigTitle: `免密登录`, ConfigRoute: `webauthn`,
	})
	d := httpserver.Backend.Dashboard.GetOrNewExtend(`login`)
	d.GlobalFooters.Add(-1, &dashboard.Tmplx{
		Tmpl: `webauthn/login/footer`,
	})
	d.GetOrNewGroupedButtons(`bottom`).Add(-1, &dashboard.Tmplx{
		Tmpl: `webauthn/login/button_bottom`,
	})
}
