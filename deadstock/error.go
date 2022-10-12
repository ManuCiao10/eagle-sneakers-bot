package deadstock

import (
	"os"
	"time"

	"github.com/fatih/color"
)

//--------ERROR---------//

func Check_product(list []Product) {
	if len(list) == 0 {
		color.Red("[ Eagle 0.0.2 ]" + "[ " + time.Now().Format("15:04:05.000000") + " ] " + "NO PRODUCT DETECTED")
		os.Exit(0)
	}
	if len(list[0].Pid) == 0 {
		color.Red("[ Eagle 0.0.2 ]" + "[ " + time.Now().Format("15:04:05.000000") + " ] " + "NO PID DETECTED")
		os.Exit(0)
	}
	if len(list[0].Size) == 0 {
		color.Red("[ Eagle 0.0.2 ]" + "[ " + time.Now().Format("15:04:05.000000") + " ] " + "NO SIZE DETECTED")
		os.Exit(0)
	}
	if len(list[0].Email) == 0 {
		color.Red("[ Eagle 0.0.2 ]" + "[ " + time.Now().Format("15:04:05.000000") + " ] " + "NO PRODUCT DETECTED")
		os.Exit(0)
	}
	if len(list[0].profile) == 0 {
		color.Red("[ Eagle 0.0.2 ]" + "[ " + time.Now().Format("15:04:05.000000") + " ] " + "NO PROFILE DETECTED")
		os.Exit(0)
	}
	if len(list[0].method) == 0 {
		color.Red("[ Eagle 0.0.2 ]" + "[ " + time.Now().Format("15:04:05.000000") + " ] " + "NO METHOD DETECTED")
		os.Exit(0)
	}
	if len(list[0].Card_Number) == 0 {
		color.Red("[ Eagle 0.0.2 ]" + "[ " + time.Now().Format("15:04:05.000000") + " ] " + "NO CARD_NUMBER DETECTED")
		os.Exit(0)
	}
	if len(list[0].Month) == 0 {
		color.Red("[ Eagle 0.0.2 ]" + "[ " + time.Now().Format("15:04:05.000000") + " ] " + "NO MONTH DETECTED")
		os.Exit(0)
	}
	if len(list[0].Year) == 0 {
		color.Red("[ Eagle 0.0.2 ]" + "[ " + time.Now().Format("15:04:05.000000") + " ] " + "NO YEAR DETECTED")
		os.Exit(0)
	}
	if len(list[0].CVV) == 0 {
		color.Red("[ Eagle 0.0.2 ]" + "[ " + time.Now().Format("15:04:05.000000") + " ] " + "NO CVV DETECTED")
		os.Exit(0)
	}
	if len(list[0].Proxy_List) == 0 {
		color.Red("[ Eagle 0.0.2 ]" + "[ " + time.Now().Format("15:04:05.000000") + " ] " + "NO PROXY DETECTED")
		os.Exit(0)
	}
}

func Check_profile(profile []Info) {
	if len(profile) == 0 {
		color.Red("[ Eagle 0.0.2 ]" + "[ " + time.Now().Format("15:04:05.000000") + " ] " + "PROFILE NOT FOUND")
		os.Exit(1)
	}
	if len(profile[0].Profile_name) == 0 {
		color.Red("[ Eagle 0.0.2 ]" + "[ " + time.Now().Format("15:04:05.000000") + " ] " + "PROFILE NOT FOUND")
		os.Exit(1)
	}
	if len(profile[0].First_name) == 0 {
		color.Red("[ Eagle 0.0.2 ]" + "[ " + time.Now().Format("15:04:05.000000") + " ] " + "NAME NOT FOUND")
		os.Exit(1)
	}
	if len(profile[0].Last_name) == 0 {
		color.Red("[ Eagle 0.0.2 ]" + "[ " + time.Now().Format("15:04:05.000000") + " ] " + "LAST_NAME NOT FOUND")
		os.Exit(1)
	}
	if len(profile[0].Phone) == 0 {
		color.Red("[ Eagle 0.0.2 ]" + "[ " + time.Now().Format("15:04:05.000000") + " ] " + "PHONE NOT FOUND")
		os.Exit(1)
	}
	if len(profile[0].Address) == 0 {
		color.Red("[ Eagle 0.0.2 ]" + "[ " + time.Now().Format("15:04:05.000000") + " ] " + "ADDRESS NOT FOUND")
		os.Exit(1)
	}
	if len(profile[0].House_Number) == 0 {
		color.Red("[ Eagle 0.0.2 ]" + "[ " + time.Now().Format("15:04:05.000000") + " ] " + "HOUSE_NUMBER NOT FOUND")
		os.Exit(1)
	}
	if len(profile[0].City) == 0 {
		color.Red("[ Eagle 0.0.2 ]" + "[ " + time.Now().Format("15:04:05.000000") + " ] " + "CITY NOT FOUND")
		os.Exit(1)
	}
	if len(profile[0].State) == 0 {
		color.Red("[ Eagle 0.0.2 ]" + "[ " + time.Now().Format("15:04:05.000000") + " ] " + "STATE NOT FOUND")
		os.Exit(1)
	}
	if len(profile[0].ZIP) == 0 {
		color.Red("[ Eagle 0.0.2 ]" + "[ " + time.Now().Format("15:04:05.000000") + " ] " + "ZIP NOT FOUND")
		os.Exit(1)
	}
	if len(profile[0].Country) == 0 {
		color.Red("[ Eagle 0.0.2 ]" + "[ " + time.Now().Format("15:04:05.000000") + " ] " + "COUNTRY NOT FOUND")
		os.Exit(1)
	}

}