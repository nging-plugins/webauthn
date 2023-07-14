package user

import (
	"github.com/admpub/nging/v5/application/model"
)

func init() {
	model.RegisterSafeItem(`webauthn`, `免密登录`, model.SafeItemInfo{
		Step: 1, ConfigTitle: `免密登录`, ConfigRoute: `webauthn`,
	})
}
