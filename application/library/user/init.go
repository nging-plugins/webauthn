package user

import (
	"github.com/admpub/nging/v4/application/model"
)

func init() {
	model.SafeItems.Add(`webauthn`, `免密登录`)
}
