package dtos

type ContextDto struct {
	Context Context `json:"@context"`
}

type Context struct {
	Version      float64 `json:"@version"`
	Protected    bool    `json:"@protected"`
	Rom          string  `json:"room"`
	Capabilities Rooms   `json:"rooms"`
}

type Rooms struct {
	ID        string `json:"@id"`
	Container string `json:"@container"`
}

func NewContext() ContextDto {
	capabilities := Rooms{
		"room:rooms",
		"@set",
	}
	context := Context{
		1.1,
		true,
		"https://api.alphacorp.vsk.gr/contexts/rooms#",
		capabilities,
	}
	dto := ContextDto{
		context,
	}
	return dto
}
