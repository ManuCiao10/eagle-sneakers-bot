package logs

import "time"

var (
	colorReset = "\033[0m"

	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorPurple = "\033[35m"
	colorCyan   = "\033[36m"
	colorWhite  = "\033[37m"
	img         = "https://media.discordapp.net/attachments/1013517214906859540/1039155134556536894/01IHswd8_400x400.jpeg"
)

type CheckoutLogRequest struct {
	TaskStart   time.Time `json:"-"`
	TaskEnd     time.Time `json:"-"`
	Price       float64   `json:"price"`
	ProductName string    `json:"product_name"`
	ProductMSKU string    `json:"product_msku"`
	Mode        string    `json:"mode"`
	CheckoutMs  int       `json:"checkout_ms"`
	Size        string    `json:"size"`
	Status      string    `json:"status"`
	Website     string    `json:"website"`
	ImageUrl    string    `json:"image_url"`
	AllowPublic bool      `json:"allow_public"`

	PayPal string `json:"paypal"`
}

type LogtailData struct {
	AuthKey     string `json:"auth_key"`
	SiteName    string `json:"site_name"`
	TaskType    string `json:"task_type"`
	TaskMode    string `json:"task_mode"`
	TaskProduct string `json:"task_product"`
	Message     string `json:"message"`
	Count       int    `json:"count"`
}

// discord webhook
type Author struct {
	Name    string `json:"name"`
	URL     string `json:"url"`
	IconURL string `json:"icon_url"`
}

type Fields struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline,omitempty"`
}
type Thumbnail struct {
	URL string `json:"url"`
}

type Image struct {
	URL string `json:"url"`
}

type Footer struct {
	Text    string `json:"text"`
	IconURL string `json:"icon_url"`
}
type Embeds struct {
	Author      Author    `json:"author"`
	Title       string    `json:"title"`
	URL         string    `json:"url"`
	Description string    `json:"description"`
	Color       int       `json:"color"`
	Fields      []Fields  `json:"fields"`
	Thumbnail   Thumbnail `json:"thumbnail"`
	Image       Image     `json:"image"`
	Footer      Footer    `json:"footer"`
}
type Top struct {
	Username  string   `json:"username"`
	AvatarURL string   `json:"avatar_url"`
	Content   string   `json:"content"`
	Embeds    []Embeds `json:"embeds"`
}

type MonitorDetected struct {
	Pid          string
	Size         string
	TaskQuantity int
	Proxy        string
	TaskFile     string
	Delay        int
	Store        string
}
