package auth

import (
	"BackendDeveloperVK-testTask/pkg/auth/jwt"
	"BackendDeveloperVK-testTask/pkg/utils"
)

func InitAuth() {
	jwt.SecretKey = []byte(utils.CONFIG.Auth.SecretKey)
	jwt.Salt = utils.CONFIG.Auth.Salt

	initRegistrationHandlers()
	initLoginHandlers()
}
