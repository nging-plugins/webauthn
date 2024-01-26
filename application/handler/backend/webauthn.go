package backend

import (
	"strings"

	"github.com/webx-top/db"
	"github.com/webx-top/echo"
	"github.com/webx-top/echo/code"
	"github.com/webx-top/echo/param"

	"github.com/admpub/nging/v5/application/handler"
	"github.com/admpub/nging/v5/application/library/common"
	"github.com/admpub/nging/v5/application/model"
)

var handlers = echo.HandlerFuncs{
	`setName`: setName,
}

var checkers = map[string]func(*string) bool{
	`id`: func(pk *string) bool {
		return param.AsUint64(*pk) > 0
	},
	`field`: func(field *string) bool {
		return *field == `name`
	},
	`value`: func(value *string) bool {
		*value = strings.TrimSpace(*value)
		return len(*value) > 0
	},
}

func setName(ctx echo.Context) error {
	m := model.NewUserU2F(ctx)
	user := handler.User(ctx)
	data := ctx.Data()
	pk, field, value, err := common.GetEditableJSFormData(ctx, checkers)
	if err != nil {
		return ctx.JSON(data.SetError(err))
	}
	id := ctx.Atop(pk).Uint64()
	err = m.Get(nil, db.And(
		db.Cond{`id`: id},
		db.Cond{`uid`: user.Id},
		db.Cond{`type`: `webauthn`},
		db.Cond{`step`: 1},
	))
	if err != nil {
		return ctx.JSON(data.SetError(err))
	}
	err = m.UpdateField(nil, field, value, `id`, id)
	if err != nil {
		return ctx.JSON(data.SetError(err))
	}
	return ctx.JSON(data.SetInfo(ctx.T(`修改成功`), code.Success.Int()))
}

func WebAuthn(ctx echo.Context) error {
	user := handler.User(ctx)
	if user == nil {
		return ctx.NewError(code.Unauthenticated, `请先登录`)
	}
	if ctx.IsPost() {
		return handlers.Call(ctx, ctx.Form(`op`))
	}
	m := model.NewUserU2F(ctx)
	err := m.ListPageByType(user.Id, `webauthn`, 1)
	if err != nil {
		return err
	}
	ctx.Set(`listData`, m.Objects())
	ctx.Set(`activeSafeItem`, `webauthn`)
	ctx.Set(`safeItems`, model.SafeItems.Slice())
	return ctx.Render(`webauthn/user`, handler.Err(ctx, err))
}
