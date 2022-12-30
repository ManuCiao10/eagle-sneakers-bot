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
}
