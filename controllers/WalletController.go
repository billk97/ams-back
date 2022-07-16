package controllers

import (
	"ams-back/usecases"
	"ams-back/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

func CreateWalletController(r *gin.Engine) {
	api := r.Group("api/wallet")
	{
		api.GET("did/public", getPublicDid)
	}
}

func getAriesWalletUrl() string {
	url := fmt.Sprintf("%s/wallet", utils.Config.Aries)
	return url
}

func getPublicDid(c *gin.Context) {

	publicDids, err := usecases.GetAriesPublicDid()
	if err != nil {
		apiErr := utils.NewApiError("NOT_FOUND", err, "")
		apiErr.Enhance(c)
		c.JSON(400, apiErr)
		return
	}
	c.JSON(200, &publicDids)
}
