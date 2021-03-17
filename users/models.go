package users

type Location struct {
	country string
	city    string
	address string
}

type User struct {
	ID             string
	Name           string
	LastName       string
	Email          string
	CommercialName string
	Segment        string
	IDType         string `json:"id_type,omitempty"`
	IDNumber       string `json:"id_number,omitempty"`
	Location       Location
	Proffesionals  []string
}
