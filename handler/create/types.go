package create

var (
	JsonTemplate = []byte(`{
  "key": "INSERT_YOUR_KEY",
  "webhook": "INSERT_YOUR_WEBHOOK",
		  
  "2captcha_key": "INSERT_YOUR_2CAPTCHA_KEY",
  "anticaptcha_key": "INSERT_YOUR_ANTICAPTCHA_KEY",
  "capmonster_key": "INSERT_YOUR_CAPMONSTER_KEY",
		  
  "solver": "SELECT_YOUR_SOLVER",
		  
  "delay": {
    "retry": "DELAY",
    "timeout": "DELAY"
  }
}`)

	JsonTemplateDEV = []byte(`{
	"key": "EAGLE-LD9W-CJ3K-NAO7-KFOV",
	"webhook": "INSERT_YOUR_WEBHOOK",
			
	"2captcha_key": "INSERT_YOUR_2CAPTCHA_KEY",
	"anticaptcha_key": "INSERT_YOUR_ANTICAPTCHA_KEY",
	"capmonster_key": "INSERT_YOUR_CAPMONSTER_KEY",
			
	"solver": "SELECT_YOUR_SOLVER",
			
	"delay": {
	  "retry": "DELAY",
	  "timeout": "DELAY"
	}
  }`)

	CsvTemplate     = []byte(`Profile Name,First Name,Last Name,Mobile Number,Address,Address 2,House Number,City,State,ZIP,Country,Billing First Name,Billing Last Name,Billing Mobile Number,Billing Address,Billing Address 2,Billing Address 3,Billing House Number,Billing City,Billing State,Billing ZIP,Billing Country`)
	CsvTemplateTask = []byte(`Url / PID,Size,E-mail,Profile,Payment Method,Card Number,Month,Year,CVV,Proxy List`)
)
