package customer

import (
	"encoding/json"

	cw "github.com/coscms/webauthn"
	"github.com/duo-labs/webauthn/webauthn"
	"github.com/webx-top/com"
	"github.com/webx-top/db"
	"github.com/webx-top/echo"
	"github.com/webx-top/echo/code"

	"github.com/admpub/webx/application/dbschema"
	modelCustomer "github.com/admpub/webx/application/model/official/customer"
)

var handle cw.UserHandler = &CustomerHandle{}

type CustomerHandle struct {
}

func (u *CustomerHandle) GetUser(ctx echo.Context, username string, opType cw.Type, stage cw.Stage) (webauthn.User, error) {
	m := modelCustomer.NewCustomer(ctx)
	err := m.Get(func(r db.Result) db.Result {
		return r.Select(`id`, `name`, `avatar`, `disabled`)
	}, `name`, username)
	if err != nil {
		return nil, err
	}
	if m.Disabled == `Y` {
		err = ctx.NewError(code.UserDisabled, `该用户已被禁用`).SetZone(`disabled`)
		return nil, err
	}
	user := &cw.User{
		ID:          m.Id,
		Name:        m.Name,
		DisplayName: m.Name,
		Icon:        m.Avatar,
	}
	u2f := dbschema.NewOfficialCustomerU2f(ctx)
	_, err = u2f.ListByOffset(nil, nil, 0, -1, db.And(
		db.Cond{`uid`: m.Id},
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

func (u *CustomerHandle) Register(ctx echo.Context, user webauthn.User, cred *webauthn.Credential) error {
	m := modelCustomer.NewCustomer(ctx)
	err := m.Get(func(r db.Result) db.Result {
		return r.Select(`id`, `disabled`)
	}, `username`, user.WebAuthnName())
	if err != nil {
		return err
	}
	if m.Disabled == `Y` {
		err = ctx.NewError(code.UserDisabled, `该用户已被禁用`).SetZone(`disabled`)
		return err
	}
	u2fM := modelCustomer.NewU2F(ctx)
	u2fM.Uid = m.Id
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

func (u *CustomerHandle) Login(ctx echo.Context, user webauthn.User, cred *webauthn.Credential) error {
	m := modelCustomer.NewCustomer(ctx)
	err := m.Get(nil, `name`, user.WebAuthnName())
	if err != nil {
		return err
	}
	m.SetSession()
	return err
}
