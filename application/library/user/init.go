package user

import (
	"github.com/coscms/webcore/model"
	"github.com/coscms/webcore/registry/dashboard"
)

func init() {
	model.RegisterSafeItem(`webauthn`, `免密登录`, model.SafeItemInfo{
		Step: 1, ConfigTitle: `免密登录`, ConfigRoute: `webauthn`,
	})
	d := dashboard.Default.Backend.GetOrNewExtend(`login`)
	d.GlobalFooters.Add(-1, &dashboard.Tmplx{
		Tmpl: `webauthn/login/footer`,
	})
	d.GetOrNewGroupedButtons(`bottom`).Add(-1, &dashboard.Tmplx{
		Tmpl: `webauthn/login/button_bottom`,
	})
}
