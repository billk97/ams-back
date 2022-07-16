package usecases

import (
	"ams-back/utils"
	"fmt"
	"net/http"
)

func DeleteAllCredentialsOffersWithState(state string) ([]string, error) {
	credOffers, err := GetAllCredentialOffers()
	if err != nil {
		return nil, err
	}
	var credOffersToDelete []string
	for _, v := range credOffers.Results {
		if v.CredExRecord.State == state {
			credOffersToDelete = append(credOffersToDelete, v.CredExRecord.CredExID)
		}
	}

	for _, v := range credOffersToDelete {
		go deleteCredentialRecord(v)
	}

	return nil, nil
}

func deleteCredentialRecord(id string) error {
	url := fmt.Sprintf("%s/issue-credential-2.0/records/%s", utils.Config.Aries, id)
	fmt.Println("concurency")
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}
	_, err = http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	return nil
}
