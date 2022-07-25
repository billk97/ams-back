package dtos

type AlphaCorpEmployeeDTO struct {
	AlphaCorpEmployeeContext AlphaCorpEmployeeContext `json:"@context"`
}

type AlphaCorpEmployeeContext struct {
	Version    float64 `json:"@version"`
	Protected  bool    `json:"@protected"`
	Emp        string  `json:"emp"`
	Schema     string  `json:"schema"`
	GivenName  string  `json:"givenName"`
	FamilyName string  `json:"familyName"`
	JobTitle   string  `json:"jobTitle"`
	Email      string  `json:"email"`
	Id         string  `json:"id"`
	Type       string  `json:"type"`
	atId       string  `json:"@id"`
}

func NewPersonContext() AlphaCorpEmployeeDTO {
	context := AlphaCorpEmployeeContext{
		1.1,
		true,
		"https://api.alphacorp.vsk.gr/contexts/alphacorp-employee#",
		"http://schema.org/",
		"schema:givenName",
		"schema:familyName",
		"schema:jobTitle",
		"schema:email",
		"@id",
		"@type",
		"https://api.alphacorp.vsk.gr/contexts/alphacorp-employee#AlphacorpCredential",
	}
	dto := AlphaCorpEmployeeDTO{
		context,
	}
	return dto
}
