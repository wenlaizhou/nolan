package server

import (
	"encoding/base64"
	"fmt"
	"github.com/wenlaizhou/middleware"
)

func checkAuth(conf middleware.Config, token string) bool {
	privateKey := conf.UnsafeDefault("private.key", "")

	if len(privateKey) <= 0 { // 没有key配置，则不进行校验
		return true
	}

	if len(token) <= 0 {
		return false
	}

	privateKeyBearer := base64.StdEncoding.EncodeToString([]byte(privateKey))
	return fmt.Sprintf("Bearer %v", privateKeyBearer) == token
}
