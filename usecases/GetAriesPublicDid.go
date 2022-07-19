package usecases

import (
	"ams-back/dtos"
	"ams-back/utils"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func GetAriesPublicDid() (*dtos.DidDTO, error) {
	var publicDid dtos.DidDTO
	publicWalletUrl := utils.Config.Aries + "/wallet/did/public"
	response, err := http.Get(publicWalletUrl)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	respDTO := dtos.WrapperDidDTO{}
	err = json.Unmarshal(body, &respDTO)
	publicDid = respDTO.Result
	if err != nil {
		return nil, err
	}
	return &publicDid, nil
}
