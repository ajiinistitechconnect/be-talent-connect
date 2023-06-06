package response

import "github.com/alwinihza/talent-connect-be/model"

type LoginResponse struct {
	AccessToken string
	TokenModel  model.TokenModel
}
