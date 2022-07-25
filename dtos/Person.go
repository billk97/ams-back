package dtos

type AlphaCorpEmployeeDTO struct {
	AlphaCorpEmployeeContext AlphaCorpEmployeeContext `json:"@context"`
}

type AlphaCorpEmployeeContext struct {
	Version     float64 `json:"@version"`
	Protected   bool    `json:"@protected"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Identifier  string  `json:"identifier"`
	//Emp               string      `json:"emp"`
	//Schema            string      `json:"schema"`
	//GivenName         string      `json:"givenName"`
	//FamilyName        string      `json:"familyName"`
	//JobTitle          string      `json:"jobTitle"`
	//Email             string      `json:"email"`
	//Id                string      `json:"id"`
	//Type              string      `json:"type"`
	//AtId              string      `json:"@id"`
	AlphaCorpEmployee AlphaCorpEmployee `json:"AlphaCorpEmployee"`
	Person            string            `json:"Person"`
}

type AlphaCorpEmployee struct {
	AtId    string        `json:"@id"`
	Context CustomContext `json:"@context"`
}

type CustomContext struct {
	Version    float64 `json:"@version"`
	Protected  bool    `json:"@protected"`
	Id         string  `json:"id"`
	Type       string  `json:"type"`
	Ctzn       string  `json:"ctzn"`
	Schema     string  `json:"schema"`
	Xsd        string  `json:"xsd"`
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
		"https://api.alphacorp.vsk.gr/contexts/alphacorp-employee#",
		"http://schema.org/",
		"http://www.w3.org/2001/XMLSchema#",
		"schema:givenName",
		"schema:familyName",
		"schema:jobTitle",
		"schema:email",
	}
	sub := AlphaCorpEmployee{
		"https://api.alphacorp.vsk.gr/contexts/alphacorp-employee#AlphacorpCredential",
		custom,
	}
	context := AlphaCorpEmployeeContext{
		1.1,
		true,
		"http://schema.org/name",
		"http://schema.org/description",
		"http://schema.org/identifier",
		sub,
		"http://schema.org/Person",
	}
	dto := AlphaCorpEmployeeDTO{
		context,
	}
	return dto
}
