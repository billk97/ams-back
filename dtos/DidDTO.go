package dtos

type DidDTO struct {
	Did     string `json:"did"`
	Verkey  string `json:"verkey"`
	Posture string `json:"posture"`
	KeyType string `json:"key_type"`
	Method  string `json:"method"`
}

type WrapperDidDTO struct {
	Result DidDTO `json:"result"`
}
