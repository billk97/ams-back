package dtos

import "time"

type CredentialExchangeRecordDTO struct {
	Results []CredentialExchangeRecord `json:"results"`
}

type CredentialExchangeRecord struct {
	CredExRecord CredExRecord `json:"cred_ex_record"`
	Indy         string       `json:"indy"`
	LdProof      interface{}  `json:"ld_proof"`
}

type CredExRecord struct {
	AutoIssue    bool        `json:"auto_issue"`
	AutoOffer    bool        `json:"auto_offer"`
	AutoRemove   bool        `json:"auto_remove"`
	ByFormat     interface{} `json:"by_format"`
	ConnectionId string      `json:"connection_id"`
	CreatedAt    time.Time   `json:"created_at"`
	CredExID     string      `json:"cred_ex_id"`
	CredOffer    interface{} `json:"cred_offer"`
	CredProposal interface{} `json:"cred_proposal"`
	Initiator    string      `json:"initiator"`
	Role         string      `json:"role"`
	State        string      `json:"state"`
	ThreadID     string      `json:"thread_id"`
	Trace        bool        `json:"trace"`
	UpdatedAt    time.Time   `json:"updated_at"`
}
