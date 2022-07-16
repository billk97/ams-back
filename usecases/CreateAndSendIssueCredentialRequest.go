package usecases

import (
	"ams-back/dtos"
	"ams-back/models"
	"ams-back/repos"
	"errors"
	"fmt"
	"time"
)

func CreateAndSendIssueCredentialRequest(employeeId int) (*dtos.IssueCredentialDTO, error) {
	employee, err := repos.FindEmployeeById(employeeId)
	if err != nil {
		return nil, err
	}
	if len(employee.Permission) <= 0 {
		return nil, errors.New("employee doesn't have any permissions")
	}
	return populateCredential(employee)
}

func populateCredential(employee *models.Employee) (*dtos.IssueCredentialDTO, error) {
	var permissions []string
	for i, s := range employee.Permission {
		fmt.Println(i, s)
		permissions = append(permissions, s.Alias)
	}
	filter := dtos.Filter{
		LdProof: dtos.LdProof{
			CredentialBody: dtos.CredentialBody{
				Context: []string{
					"https://www.w3.org/2018/credentials/v1",
					"https://www.w3.org/2018/credentials/examples/v1",
				},
				Type: []string{
					"VerifiableCredential",
				},
				Issuer:       "did:sov:GHZXFFQdytHVVXywsQaukB", // todo get from ledger
				IssuanceDate: time.Now().UTC(),
				CredentialSubject: dtos.CredentialSubject{
					Id:         "did:sov:FyEkpqHm8NBGmWqdf4DPbn", // fetch the DID
					GivenName:  employee.FirstName,
					FamilyName: employee.LastName,
					JobTitle:   employee.JobTitle,
					Email:      employee.Email,
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
