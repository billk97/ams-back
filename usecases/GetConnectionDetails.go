package usecases

import (
	"ams-back/dtos"
	"ams-back/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetConnectionDetails(connectionId string) (*dtos.ConnectionDTO, error) {
	url := fmt.Sprintf("%s/connections/%s", utils.Config.Aries, connectionId)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var connection dtos.ConnectionDTO
	jsonErr := json.Unmarshal(body, &connection)
	if jsonErr != nil {
		return nil, jsonErr
	}
	return &connection, nil
}
