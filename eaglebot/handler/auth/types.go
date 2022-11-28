package auth

// type AuthRequest struct {
// 	Key  string `json:"licenseKey"`
// 	HWID string `json:"HWID"`
// }

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
