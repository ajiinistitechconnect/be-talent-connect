package authenticator

import (
	"github.com/alwinihza/talent-connect-be/model"
	"github.com/golang-jwt/jwt"
)

type MyClaims struct {
	jwt.StandardClaims
	model.TokenModel
	AccessUUID string
}
