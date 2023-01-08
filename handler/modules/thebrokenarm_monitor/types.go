package thebrokenarm_monitor

import (
	"github.com/eagle/handler/account"
	"github.com/eagle/handler/quicktask"
)

// with quicktask read the accounts.csv file and login
var (
	INITIALIZE quicktask.TaskState = "initialize"
	LOGIN      quicktask.TaskState = "login"
)

var siteIdMap = map[string]int{
	"thebrokenarm": 0,
}

var TBAInternalQuick = struct {
	Account account.Account
}{}
