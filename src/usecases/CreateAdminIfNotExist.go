package usecases

import (
	"ams-back/src/database"
	"ams-back/src/models"
	"ams-back/src/repos"
	"crypto/rand"
	b64 "encoding/base64"
	"fmt"
	"log"
)

func CreateSuperAdminIfNotExists() {
	existingAdmin, err := repos.FindAdminByUsername("superAdmin")
	if err != nil {
		log.Printf("Error In creation of user superAdmin")
		return
	}
	if existingAdmin != nil && err == nil {
		log.Printf("Admin already exists")
		return
	}
	superAdmin := models.Admin{Username: "superAdmin"}
	b := make([]byte, 50)
	_, randErr := rand.Read(b)
	if randErr != nil {
		log.Printf("Error In password generation")
		return
	}
	sEnc := b64.StdEncoding.EncodeToString(b)
	fmt.Printf("SuperAdmin password: %s", string(sEnc))
	superAdmin.GeneratePasswordHash(string(sEnc))
	db := database.GetDb()
	db.Save(superAdmin)
}
