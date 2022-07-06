package dtos

type CreateInvitationDTO struct {
	ConnectionId  string      `json:"connection_id"`
	Invitation    interface{} `json:"invitation"`
	InvitationUrl string      `json:"invitation_url"`
}
