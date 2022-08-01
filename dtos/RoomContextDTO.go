package dtos

type RoomContextDto struct {
	RoomContext RoomContext `json:"@context"`
}

type RoomContext struct {
	Version         float64         `json:"@version"`
	Protected       bool            `json:"@protected"`
	Name            string          `json:"name"`
	Description     string          `json:"description"`
	Identifier      string          `json:"identifier"`
	RoomCredentials RoomCredentials `json:"RoomCredential"`
	//Rooms           Rooms           `json:"rooms"`
}

type RoomCredentials struct {
	ID string `json:"@id"`
	//Context   RoomCredentialsContext `json:"@context"`
	Container string `json:"@container"`
}

type RoomCredentialsContext struct {
	Version   float64 `json:"@version"`
	Protected bool    `json:"@protected"`
	Id        string  `json:"id"`
	Type      string  `json:"type"`
}

func NewRoomContext() RoomContextDto {
	//rooms := Rooms{
	//	"room:rooms",
	//	"@set",
	//}
	context := RoomContext{
		1.1,
		true,
		"http://schema.org/name",
		"http://schema.org/description",
		"http://schema.org/identifier",
		RoomCredentials{
			"https://api.alphacorp.vsk.gr/contexts/rooms#RoomCredentials",
			"@set",
		},
	}
	dto := RoomContextDto{
		context,
	}
	return dto
}
