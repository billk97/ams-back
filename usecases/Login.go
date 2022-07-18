package usecases

import (
	"ams-back/dtos"
	"ams-back/repos"
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v4"
)

func Login(dto dtos.CredentialDto) (*dtos.AccessDto, error) {
	user, err := repos.FindAdminByUsername(dto.Username)
	if err != nil || user == nil {
		return nil, err
	}

	if !user.CheckPasswordHash(dto.Password) {
		return nil, errors.New(fmt.Sprintf("Credentials are invalid!"))
	}
	access := dtos.AccessDto{}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{})
	//todo get from .env
	tokenString, singErr := jwtToken.SignedString([]byte("secret"))
	if singErr != nil {
		return nil, singErr
	}
	access.Token = tokenString
	return &access, nil
}
