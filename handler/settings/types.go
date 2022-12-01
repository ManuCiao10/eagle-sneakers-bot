package settings

type Settings struct {
	AuthKey        string `json:"key"`
	DiscordWebhook string `json:"webhook"`
	TwoCaptchaKey  string `json:"2captcha_key"`
	AnticaptchaKey string `json:"anticaptcha_key"`
	CapmonsterKey  string `json:"capmonster_key"`
	Solver         string `json:"solver"`
	Delay          Delay  `json:"delay"`
}

type Delay struct {
	Retry   string `json:"retry"`
	Timeout string `json:"timeout"`
}

type Env struct {
	API_TOKEN      string `json:"API_TOKEN"`
	AUTH_DOWNLOAD  string `json:"AUTH_DOWNLOAD"`
	ACC_DOWLOAD    string `json:"ACC_DOWLOAD"`
	DISCORD_APP_ID string `json:"DISCORD_APP_ID"`
}
