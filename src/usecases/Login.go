package usecases

import (
	"ams-back/src/dtos"
	"ams-back/src/repos"
	"ams-back/src/utils"
	"fmt"

	"github.com/golang-jwt/jwt/v4"
)

func Login(dto dtos.CredentialDto) (*dtos.AccessDto, *utils.ApiError) {
	user, err := repos.FindAdminByUsername(dto.Username)
	if err != nil || user == nil {
		return nil, err
	}

	if !user.CheckPasswordHash(dto.Password) {
		return nil, utils.NewApiError(
			"INVALED_CREDENTIALS",
			nil,
			fmt.Sprintf("Credentials are invalid!"),
		)
	}
	access := dtos.AccessDto{}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{})
	tokenString, singErr := jwtToken.SignedString([]byte("secret"))
	if singErr != nil {
		return nil, utils.NewApiError(
			"SIGN_FAILED",
			singErr,
			fmt.Sprintf("Failed to sing jwt"),
		)
	}
	access.Token = tokenString
	return &access, nil
}
