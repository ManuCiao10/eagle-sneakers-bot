package version

type Update struct {
	Data []Data `json:"data"`
}

type Data struct {
	Version string `json:"name"`
	ID      string `json:"id"`
}
