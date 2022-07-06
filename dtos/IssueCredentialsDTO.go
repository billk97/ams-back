package dtos

import "time"

type IssueCredentialDTO struct {
	ConnectionID string `json:"connection_id"`
	Filter       Filter `json:"filter"`
}

type Filter struct {
	LdProof LdProof `json:"ld_proof"`
}

type LdProof struct {
	Credential struct {
		Context           []string  `json:"@context"`
		Type              []string  `json:"type"`
		Issuer            string    `json:"issuer"`
		IssuanceDate      time.Time `json:"issuanceDate"`
		CredentialSubject struct {
			Degree struct {
				Type string `json:"type"`
				Name string `json:"name"`
			} `json:"degree"`
			College string `json:"college"`
		} `json:"credentialSubject"`
	} `json:"credential"`
	Options Options `json:"options"`
}
