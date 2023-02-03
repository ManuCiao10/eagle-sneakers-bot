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
	"webhook": "https://discord.com/api/webhooks/1039416403834441768/sjA3RKtRY2H3v-RZUOTPvA3RTSP9WAm2ndkWdcPhQRUw2EE97_C5tfcx9dmV9qVjKvZj",
			
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

	ProxiesTemplateDEV = []byte(`
	185.91.205.3:5733:hgj3x3cas2:0ef2uixpcu
	185.91.205.175:5874:hgj3x3cas2:0ef2uixpcu
	185.91.204.143:5446:hgj3x3cas2:0ef2uixpcu
	185.91.204.135:6175:hgj3x3cas2:0ef2uixpcu
	185.91.207.12:5221:hgj3x3cas2:0ef2uixpcu
	185.91.206.21:6956:hgj3x3cas2:0ef2uixpcu
	185.91.205.234:5194:hgj3x3cas2:0ef2uixpcu
	185.91.207.197:5460:hgj3x3cas2:0ef2uixpcu
	185.91.204.5:5022:hgj3x3cas2:0ef2uixpcu
	185.91.206.69:7516:hgj3x3cas2:0ef2uixpcu
	185.91.205.40:5736:hgj3x3cas2:0ef2uixpcu
	185.91.204.182:6318:hgj3x3cas2:0ef2uixpcu
	185.91.206.181:6871:hgj3x3cas2:0ef2uixpcu
	185.91.206.115:6924:hgj3x3cas2:0ef2uixpcu
	185.91.205.212:5452:hgj3x3cas2:0ef2uixpcu
	185.91.207.248:6701:hgj3x3cas2:0ef2uixpcu
	185.91.204.172:7204:hgj3x3cas2:0ef2uixpcu
	185.91.207.159:7432:hgj3x3cas2:0ef2uixpcu
	185.91.206.165:7445:hgj3x3cas2:0ef2uixpcu
	185.91.207.39:5820:hgj3x3cas2:0ef2uixpcu
	185.91.207.95:6908:hgj3x3cas2:0ef2uixpcu
	185.91.206.185:5032:hgj3x3cas2:0ef2uixpcu
	185.91.204.249:7231:hgj3x3cas2:0ef2uixpcu
	185.91.205.45:5360:hgj3x3cas2:0ef2uixpcu
	185.91.205.243:5262:hgj3x3cas2:0ef2uixpcuf 
`)

	CsvTemplateAccount = []byte(`
	EMAIL,PASSWORD`)

	// CsvTemplateMQT = []byte(`Site,Tasks Quantity,Profiles,Accounts (guest/accounts),Email,Proxylist,Payment Method,Credit Card,Other`)
)
