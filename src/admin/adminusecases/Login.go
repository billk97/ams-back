package adminusecases

import (
	"ams-back/src/admin"
	"ams-back/src/amserr"
	"ams-back/src/dtos"
)

func Login(dto dtos.CredentialDto) *amserr.ApiError {
	user, err := admin.FindAdminByUsername(dto.Username)
	if err != nil || user == nil {
		return err
	}
	return nil
}
