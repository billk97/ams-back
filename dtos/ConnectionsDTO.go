package dtos

type ConnectionDTO struct {
	Rfc23State         string `json:"rfc23_state"`
	CreatedAt          string `json:"created_at"`
	InvitationMode     string `json:"invitation_mode"`
	State              string `json:"state"`
	ConnectionProtocol string `json:"connection_protocol"`
	InvitationKey      string `json:"invitation_key"`
	Accept             string `json:"accept"`
	UpdatedAt          string `json:"updated_at"`
	TheirRole          string `json:"their_role"`
	RoutingState       string `json:"routing_state"`
	ConnectionId       string `json:"connection_id"`
}
