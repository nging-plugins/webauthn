package common

import (
	"github.com/admpub/useragent"
)

func GetOS(userAgent string) string {
	infoUA := useragent.Parse(userAgent)
	return infoUA.OS
}
