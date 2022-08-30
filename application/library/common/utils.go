package common

import (
	ua "github.com/admpub/useragent"
)

const (
	SessionKeyUnbindToken = `webauthn.unbind.token`
)

func GetOS(userAgent string) string {
	infoUA := ua.Parse(userAgent)
	return infoUA.OS
}
