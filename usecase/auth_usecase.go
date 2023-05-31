package usecase

import (
	"fmt"

	"github.com/alwinihza/talent-connect-be/delivery/api/request"
	"github.com/alwinihza/talent-connect-be/model"
	"github.com/alwinihza/talent-connect-be/utils"
)

type AuthUsecase interface {
	Login(payload model.UserCredentials) (*model.User, error)
	LoginGmail() (model.User, error)
	ChangePassword(email string, requestData request.ChangePassword) error
	ForgetPassword() error
}

type authUsecase struct {
	user UserUsecase
}

// LoginGmail implements AuthUsecase
func (*authUsecase) LoginGmail() (model.User, error) {
	panic("unimplemented")
}

// changePassword implements AuthUsecase
func (a *authUsecase) ChangePassword(email string, requestData request.ChangePassword) error {
	user, err := a.user.SearchEmail(email)
	if err != nil {
		return err
	}
	if !utils.ComparePassword(user.Password, []byte(requestData.CurrentPassword)) {
		return fmt.Errorf("Password not valid")
	}
	newPassword, err := utils.SaltPassword([]byte(requestData.NewPassword))
	if err != nil {
		return err
	}
	user.Password = newPassword
	return a.user.UpdateData(user)
}

// forgetPassword implements AuthUsecase
func (*authUsecase) ForgetPassword() error {
	// add key to redis (different DB)
	panic("unimplemented")
}

// verifyLogin implements AuthUsecase
func (a *authUsecase) Login(payload model.UserCredentials) (*model.User, error) {
	user, err := a.user.SearchEmail(payload.Email)
	if err != nil {
		return nil, err
	}
	if !utils.ComparePassword(user.Password, []byte(payload.Password)) {
		return nil, fmt.Errorf("Email/Password invalid")
	}
	return user, nil
}

func NewAuthUsecase(user UserUsecase) AuthUsecase {
	return &authUsecase{
		user: user,
	}
}
