package create

var (
	JsonTemplate = []byte(`{
  "key": "INSERT_YOUR_KEY",
  "webhook": "INSERT_YOUR_WEBHOOK",
		  
  "2captcha_key": "INSERT_YOUR_2CAPTCHA_KEY",
  "anticaptcha_key": "INSERT_YOUR_ANTICAPTCHA_KEY",
  "capmonster_key": "INSERT_YOUR_CAPMONSTER_KEY",
		  
  "solver": "capmonster_key",
		  
  "delay": {
    "retry": "2000",
    "timeout": "2000"
  },

  "task_shutdown": "3"
}`)

	JsonTemplateDEV = []byte(`{
	"key": "EAGLE-LD9W-CJ3K-NAO7-KFOV",
	"webhook": "",
			
	"2captcha_key": "INSERT_YOUR_2CAPTCHA_KEY",
	"anticaptcha_key": "INSERT_YOUR_ANTICAPTCHA_KEY",
	"capmonster_key": "INSERT_YOUR_CAPMONSTER_KEY",
			
	"solver": "SELECT_YOUR_SOLVER",
			
	"delay": {
	  "retry": "2000",
	  "timeout": "2000"
	},

	"task_shutdown": "3"
  }`)

	CsvTemplate     = []byte(`Profile Name,First Name,Last Name,Mobile Number,Address,Address 2,House Number,City,State,ZIP,Country`)
	CsvTemplateTask = []byte(`Mode,Url / PID,Size,E-mail,Profile,Payment Method,Card Number,Month,Year,CVV,Proxy List`)

	ProxiesTemplateDEV = []byte(`INSERT_YOUR_PROXY_LIST_HERE`)

	CsvTemplateAccount = []byte(`EMAIL,PASSWORD`)

	// CsvTemplateMQT = []byte(`Site,Tasks Quantity,Profiles,Accounts (guest/accounts),Email,Proxylist,Payment Method,Credit Card,Other`)
)
