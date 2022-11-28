package auth

type AuthResponse struct {
	Integrations Integrations `json:"integrations"`
	Metadata     Metadata     `json:"metadata"`
}
type Discord struct {
	ID            string `json:"id"`
	Username      string `json:"username"`
	Discriminator string `json:"discriminator"`
}
type Integrations struct {
	Discord Discord `json:"discord"`
}

type Metadata struct {
	HWID string `json:"hwid"`
}
