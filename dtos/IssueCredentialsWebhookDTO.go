package dtos

import "time"

type IssueCredentialWebhookDTO struct {
	State        string      `json:"state"`
	UpdatedAt    time.Time   `json:"updated_at"`
	ThreadID     string      `json:"thread_id"`
	CredProposal interface{} `json:"cred_proposal"`
	CredOffer    interface{} `json:"cred_offer"`
	CredExID     string      `json:"cred_ex_id"`
	ByFormat     struct {
		CredProposal struct {
			LdProof struct {
				Credential interface{} `json:"credential"`
				Options    Options     `json:"options"`
			} `json:"ld_proof"`
		} `json:"cred_proposal"`
		CredOffer struct {
			LdProof struct {
				Credential interface{} `json:"credential"`
				Options    Options     `json:"options"`
			} `json:"ld_proof"`
		} `json:"cred_offer"`
	} `json:"by_format"`
	AutoOffer    bool      `json:"auto_offer"`
	AutoIssue    bool      `json:"auto_issue"`
	AutoRemove   bool      `json:"auto_remove"`
	Trace        bool      `json:"trace"`
	ConnectionID string    `json:"connection_id"`
	Initiator    string    `json:"initiator"`
	Role         string    `json:"role"`
	CreatedAt    time.Time `json:"created_at"`
}

type Options struct {
	ProofType string `json:"proofType"`
}
