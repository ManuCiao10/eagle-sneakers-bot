package thebrokenarm

type Product struct {
	Pid         string
	Size        string
	Email       string
	profile     string
	method      string
	Card_Number string
	Month       string
	Year        string
	CVV         string
	Proxy_List  string
}

type Info struct {
	Profile_name string
	First_name   string
	Last_name    string
	Phone        string
	Address      string
	Address_2    string
	House_Number string
	City         string
	State        string
	ZIP          string
	Country      string
}

var profile []Info
var list []Product
