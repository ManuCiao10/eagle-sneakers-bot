package settings

type Settings struct {
	AuthKey        string `json:"key"`
	DiscordWebhook string `json:"discord_webhook"`

	TwoCaptcha  string `json:"TwoCaptcha"`
	CapMonster  string `json:"CapMonster"`
	AntiCaptcha string `json:"AntiCaptcha"`

	Solver string `json:"Solver"`
	Delay  Delay  `json:"Delay"`
}

type Delay struct {
	Retry int `json:"retry"`
}
