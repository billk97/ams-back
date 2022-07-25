package dtos

type AlphaCorpEmployeeDTO struct {
	AlphaCorpEmployeeContext AlphaCorpEmployeeContext `json:"@context"`
}

type AlphaCorpEmployeeContext struct {
	Version           float64     `json:"@version"`
	Protected         bool        `json:"@protected"`
	Emp               string      `json:"emp"`
	Schema            string      `json:"schema"`
	GivenName         string      `json:"givenName"`
	FamilyName        string      `json:"familyName"`
	JobTitle          string      `json:"jobTitle"`
	Email             string      `json:"email"`
	Id                string      `json:"id"`
	Type              string      `json:"type"`
	AtId              string      `json:"@id"`
	AlphaCorpEmployee interface{} `json:"AlphacorpCredential"`
}

type SubCred struct {
	AtId    string        `json:"@id"`
	Context CustomContext `json:"@context"`
}

type CustomContext struct {
	Version    float64 `json:"@version"`
	Protected  bool    `json:"@protected"`
	Id         string  `json:"id"`
	Type       string  `json:"type"`
	Schema     string  `json:"schema"`
	GivenName  string  `json:"givenName"`
	FamilyName string  `json:"familyName"`
	JobTitle   string  `json:"jobTitle"`
	Email      string  `json:"email"`
}

func NewPersonContext() AlphaCorpEmployeeDTO {
	custom := CustomContext{
		1.1,
		true,
		"@id",
		"@type",
		"http://schema.org/",
		"schema:givenName",
		"schema:familyName",
		"schema:jobTitle",
		"schema:email",
	}
	sub := SubCred{
		"https://api.alphacorp.vsk.gr/contexts/alphacorp-employee#AlphacorpCredential",
		custom,
	}
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
		sub,
	}
	dto := AlphaCorpEmployeeDTO{
		context,
	}
	return dto
}
