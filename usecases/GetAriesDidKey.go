package usecases

import (
	"ams-back/dtos"
	"ams-back/utils"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func GetAriesDidKey() (*dtos.DidDTO, error) {
	var dids []dtos.DidDTO
	privateWalletUrl := utils.Config.Aries + "/wallet/did"
	response, err := http.Get(privateWalletUrl)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	respDTO := dtos.WrapperPrivateDidDTO{}
	err = json.Unmarshal(body, &respDTO)
	dids = respDTO.Results
	if err != nil {
		return nil, err
	}
	var did dtos.DidDTO
	for _, d := range dids {
		if d.Method == "key" {
			did = d
		}
	}
	return &did, nil
}
