package usecases

import (
	"ams-back/dtos"
	"ams-back/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetAllCredentialOffers() (*dtos.CredentialExchangeRecordDTO, error) {
	url := fmt.Sprintf("%s/issue-credential-2.0/records", utils.Config.Aries)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var credExchRecord dtos.CredentialExchangeRecordDTO
	err = json.Unmarshal(body, &credExchRecord)
	if err != nil {
		return nil, err
	}
	return &credExchRecord, nil
}
