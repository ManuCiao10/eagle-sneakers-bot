package thebrokenarm_monitor

import (
	"github.com/eagle/handler/account"
	"github.com/eagle/handler/quicktask"
)

// with quicktask read the accounts.csv file and login
var (
	INITIALIZE quicktask.TaskState = "initialize"
	LOGIN      quicktask.TaskState = "login"
	SESSION    quicktask.TaskState = "session"
	GUEST      quicktask.TaskState = "guest"
	CHECKOUT   quicktask.TaskState = "checkout"
)

var siteIdMap = map[string]int{
	"thebrokenarm": 0,
}

var TBAInternalQuick = struct {
	Account    account.Account
	Cookies    string
	ProductID  string
}{}

var dataResponse Response

var payPal Token

type Token struct {
	Success bool   `json:"success"`
	Token   string `json:"token"`
}

type Response struct {
	Success  bool        `json:"success"`
	Quantity interface{} `json:"quantity"`
	Cart     Cart        `json:"cart"`
}
type Cart struct {
	Products []Products `json:"products"`
}
type Products struct {
	AddToCartURL string     `json:"add_to_cart_url"`
	ID           string     `json:"id"`
	Attributes   Attributes `json:"attributes"`

	Images []Images `json:"images"`

	RegularPriceAmount string `json:"regular_price_amount"`
}

type Attributes struct {
	Taille string `json:"Taille"`
}

type Images struct {
	Medium Medium `json:"medium"`
}
type Medium struct {
	URL string `json:"url"`
}
