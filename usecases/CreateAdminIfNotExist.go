package usecases

import (
	database "ams-back/database"
	models "ams-back/models"
	repos "ams-back/repos"
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
	b := make([]byte, 20)
	_, randErr := rand.Read(b)
	if randErr != nil {
		log.Printf("Error In password generation")
		return
	}
	sEnc := b64.StdEncoding.EncodeToString(b)
	fmt.Printf("SuperAdmin password: %s", string(sEnc))
	superAdmin.GeneratePasswordHash(string(sEnc))
	db := database.GetDb()
	db.Save(&superAdmin)
}
