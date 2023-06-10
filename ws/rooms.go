package ws

type Room struct {
	ID          string             `json:"id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Clients     map[string]*Client `json:"clients"`
}
