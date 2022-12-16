package profile

type Profile struct {
	ID           string `json:"name"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	MobileNumber string `json:"mobileNumber"`
	Address      string `json:"address"`
	Address2     string `json:"address2"`
	HouseNumber  string `json:"houseNumber"`
	City         string `json:"city"`
	State        string `json:"state"`
	Zip          string `json:"zip"`
	Country      string `json:"country"`
}

