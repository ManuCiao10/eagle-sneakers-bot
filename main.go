package main

import (
	"log"
	"strings"
	"time"

	"github.com/eagle/deadstock"
	"github.com/eagle/eaglebot/handler/auth"
	"github.com/fatih/color"

	// "github.com/eagle/eaglebot/handler/console"
	"github.com/eagle/eaglebot/handler/loading"
	"github.com/eagle/eaglebot/handler/utils"
	"github.com/eagle/eaglebot/handler/version"
	"github.com/joho/godotenv"
	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	// "go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	//add check updates
	err := godotenv.Load("config/.env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	loading.Initialize()
	auth.Initialize()

	// rich_presence.Initialize()
	utils.Banner()
	username := strings.ToUpper(auth.Auth.Integrations.Discord.Username)
	color.Red("WELCOME BACK TO EAGLE " + username)
	println("\n")
	utils.Site_list()
	// utils.Check_update()
	// utils.Menu()
	mode := utils.SelectMode("[Eagle " + version.Version + "]" + "[" + time.Now().Format("15:04:05.000000") + "]" + " PLEASE SELECT SITE:")
	if mode == "1" {
		print("GAMESTOP")
	} else if mode == "2" {
		deadstock.Menu_deadstock()
	} else if mode == "3" {
		print("DADSTOCK")
	} else {
		utils.ConsolePrint("INVALID OPTION!", "red")
	}
}
