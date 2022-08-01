package usecases

import (
	"ams-back/dtos"
	"ams-back/models"
	"ams-back/repos"
	"ams-back/utils"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func CreateAndSendIssueCredentialRequest(employeeId int) (*dtos.CredExRecord, error) {
	employee, err := repos.FindEmployeeById(employeeId)
	if err != nil {
		return nil, err
	}
	if len(employee.Permission) <= 0 {
		return nil, errors.New("employee doesn't have any permissions")
	}
	cred, err := populateCredential(employee)
	if err != nil {
		return nil, err
	}
	resp, err := sendCredentialOffer(cred)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func sendCredentialOffer(cred *dtos.IssueCredentialDTO) (*dtos.CredExRecord, error) {
	url := fmt.Sprintf("%s/issue-credential-2.0/send", utils.Config.Aries)
	fmt.Println("url: " + url)
	body, err := json.Marshal(cred)
	if err != nil {
		return nil, err
	}
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var dto dtos.CredExRecord
	err = json.Unmarshal(respBody, &dto)
	if err != nil {
		return nil, err
	}
	return &dto, nil
}

func populateCredential(employee *models.Employee) (*dtos.IssueCredentialDTO, error) {
	var permissions []string
	for _, s := range employee.Permission {
		permissions = append(permissions, s.Alias)
	}
	didDTO, err := GetAriesPublicDid()
	if err != nil {
		return nil, err
	}
	connection, err := GetConnectionDetails(employee.DidConnectionId)
	if err != nil {
		return nil, err
	}

	filter := dtos.Filter{
		LdProof: dtos.LdProof{
			CredentialBody: dtos.CredentialBody{
				Context: []string{
					"https://www.w3.org/2018/credentials/v1",
					"https://api.alphacorp.vsk.gr/contexts/rooms/v1",
					"https://api.alphacorp.vsk.gr/contexts/alphacorp-employee",
				},
				Type: []string{
					"VerifiableCredential",
					"RoomsCredential",
					"AlphacorpCredential",
				},
				Issuer:       fmt.Sprintf("did:%s:%s", didDTO.Method, didDTO.Did),
				IssuanceDate: time.Now().UTC(),
				CredentialSubject: dtos.CredentialSubject{
					CredentialType: []string{"AlphacorpCredential"},
					Id:             fmt.Sprintf("did:sov:%s", connection.TheirDid),
					GivenName:      employee.FirstName,
					FamilyName:     employee.LastName,
					JobTitle:       employee.JobTitle,
					Email:          employee.Email,
				},
				Rooms: permissions,
			},

			Options: dtos.Options{
				ProofType: "Ed25519Signature2018",
			},
		},
	}
	dto := dtos.IssueCredentialDTO{
		ConnectionID: employee.DidConnectionId,
		Filter:       filter,
	}
	return &dto, nil
}
