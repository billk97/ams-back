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
	CredentialBody CredentialBody `json:"credential"`
	Options        Options        `json:"options"`
}

type CredentialBody struct {
	Context           []string          `json:"@context"`
	Type              []string          `json:"type"`
	Issuer            string            `json:"issuer"`
	IssuanceDate      time.Time         `json:"issuanceDate"`
	CredentialSubject CredentialSubject `json:"credentialSubject"`
	Rooms             []string          `json:"rooms"`
}

// this is a must have property
type CredentialSubject struct {
	Name string `json:"name"`
}
