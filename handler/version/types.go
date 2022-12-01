package version

type Update struct {
	Info []Info `json:"data"`
}

type Info struct {
	Version string   `json:"name"`
	Files   []string `json:"files"`
}
