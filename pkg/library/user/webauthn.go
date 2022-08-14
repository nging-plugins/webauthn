package user

import (
	"encoding/json"

	cw "github.com/coscms/webauthn"
	"github.com/duo-labs/webauthn/webauthn"
	"github.com/webx-top/com"
	"github.com/webx-top/db"
	"github.com/webx-top/echo"
	"github.com/webx-top/echo/code"

	"github.com/admpub/nging/v4/application/dbschema"
	"github.com/admpub/nging/v4/application/handler"
	"github.com/admpub/nging/v4/application/model"
)

var handle cw.UserHandler = &UserHandle{}

type UserHandle struct {
}

func (u *UserHandle) GetUser(ctx echo.Context, username string, opType cw.Type, stage cw.Stage) (webauthn.User, error) {
	if opType == cw.TypeRegister {
		user := handler.User(ctx)
		if user == nil {
			return nil, ctx.NewError(code.Unauthenticated, `请先登录`)
		}
	}
	userM := model.NewUser(ctx)
	err := userM.Get(func(r db.Result) db.Result {
		return r.Select(`id`, `username`, `avatar`, `disabled`)
	}, `username`, username)
	if err != nil {
		return nil, err
	}
	if userM.Disabled == `Y` {
		err = ctx.NewError(code.UserDisabled, `该用户已被禁用`).SetZone(`disabled`)
		return nil, err
	}
	user := &cw.User{
		ID:          uint64(userM.Id),
		Name:        userM.Username,
		DisplayName: userM.Username,
		Icon:        userM.Avatar,
	}
	u2f := dbschema.NewNgingUserU2f(ctx)
	_, err = u2f.ListByOffset(nil, nil, 0, -1, db.And(
		db.Cond{`uid`: userM.Id},
		db.Cond{`type`: `webauthn`},
		db.Cond{`step`: 1},
	))
	if err != nil {
		return nil, err
	}
	u2fList := u2f.Objects()
	user.Credentials = make([]webauthn.Credential, len(u2fList))
	for index, row := range u2fList {
		cred := webauthn.Credential{}
		err = json.Unmarshal([]byte(row.Extra), &cred)
		if err != nil {
			return nil, err
		}
		user.Credentials[index] = cred
	}
	return user, nil
}

func (u *UserHandle) Register(ctx echo.Context, user webauthn.User, cred *webauthn.Credential) error {
	userM := model.NewUser(ctx)
	err := userM.Get(func(r db.Result) db.Result {
		return r.Select(`id`, `disabled`)
	}, `username`, user.WebAuthnName())
	if err != nil {
		return err
	}
	if userM.Disabled == `Y` {
		err = ctx.NewError(code.UserDisabled, `该用户已被禁用`).SetZone(`disabled`)
		return err
	}
	u2fM := model.NewUserU2F(ctx)
	u2fM.Uid = userM.Id
	u2fM.Token = com.ByteMd5(cred.ID)
	b, err := json.Marshal(cred)
	if err != nil {
		return err
	}
	u2fM.Extra = string(b)
	u2fM.Type = `webauthn`
	u2fM.Step = 1
	_, err = u2fM.Add()
	return err
}

func (u *UserHandle) Login(ctx echo.Context, user webauthn.User, cred *webauthn.Credential) error {
	userM := model.NewUser(ctx)
	err := userM.Get(nil, `username`, user.WebAuthnName())
	if err != nil {
		return err
	}
	userM.SetSession()
	return err
}
