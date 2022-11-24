package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/eagle/deadstock"
	// "github.com/eagle/eagle/handler/auth"
	"github.com/eagle/eagle/handler/version"
	"github.com/eagle/utils"
	"github.com/fatih/color"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	//add check updates
	err := godotenv.Load("config/.env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func Change_id(ID_object string) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://" + os.Getenv("USERNAME") + ":" + os.Getenv("PASSWORD") + "@cluster0.8azzuqv.mongodb.net/?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	col := client.Database("autentichation").Collection("autentichation")
	id_obj, _ := primitive.ObjectIDFromHex(ID_object)
	res, err := col.UpdateOne(
		ctx,
		bson.M{"_id": id_obj},
		bson.D{
			{Key: "$set", Value: bson.D{{Key: "id", Value: utils.Gen_id()}}},
		},
	)
	_ = res
	if err != nil {
		color.HiMagenta("[ " + time.Now().Format("15:04:05.000000") + " ]" + "OPEN TICKET ERROR TO UPDATE KEY")
		os.Exit(1)
	}
	color.HiMagenta("[ " + time.Now().Format("15:04:05.000000") + " ]" + " ID UPDATED RESTART BOT")
	os.Exit(1)
}

func CheckId(id_database string, ID_object string) {
	if id_database == "" {
		Change_id(ID_object)
	}
	if utils.Gen_id() == id_database {
		color.HiMagenta("[" + time.Now().Format("15:04:05.000000") + "]" + " ID VALID")
	} else {
		color.HiMagenta("[ " + time.Now().Format("15:04:05.000000") + " ]" + " ID NOT VALID")
		answer := utils.SelectMode("[ " + time.Now().Format("15:04:05.000000") + " ]" + " DO YOU WANT TO RESET IT? (Y/N): ")
		if answer == "y" || answer == "Y" || answer == "yes" || answer == "YES" {
			Change_id(ID_object)
			color.HiMagenta("[ " + time.Now().Format("15:04:05.000000") + " ]" + " ID UPDATED")
			os.Exit(1)
		} else {
			color.HiMagenta("[ " + time.Now().Format("15:04:05.000000") + " ]" + " ID NOT UPDATED")
			os.Exit(1)
		}
	}

}

func Read_json() bool {
	color.Red("[" + time.Now().Format("15:04:05.000000") + "]" + " CHECKING KEY...")

	var key string
	content, err := os.ReadFile("eagle/bin/release/setting.json")
	if err != nil {
		log.Fatal(err)
	}

	var payload map[string]interface{}
	err = json.Unmarshal(content, &payload)
	if err != nil {
		log.Fatal(err)
	}
	key = payload["key"].(string)
	uuid := payload["uuid"].(string)
	return Read_database(key, uuid) // HE'S is gonna return true or false
}

func Read_database(key string, uuid string) bool {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://" + os.Getenv("USERNAME") + ":" + os.Getenv("PASSWORD") + "@cluster0.8azzuqv.mongodb.net/?retryWrites=true&w=majority"))
	if err != nil {
		deadstock.Print_err("CONNECTION ERROR D")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	col := client.Database("autentichation").Collection("autentichation")
	cursor, err := col.Find(context.TODO(), bson.D{})
	if err != nil {
		fmt.Println("Finding all documents ERROR:", err)
		defer cursor.Close(ctx)
	} else {
		// iterate over docs using Next()
		for cursor.Next(ctx) {
			// Declare a result BSON object
			var result bson.M
			err := cursor.Decode(&result)
			if err != nil {
				fmt.Println("cursor.Next() error:", err)
				os.Exit(1)
				// If there are no cursor.Decode errors
			} else {
				if result[uuid] == key {
					// fmt.Println("\nresult type:", reflect.TypeOf(result))
					ID_object := result["_id"].(primitive.ObjectID).Hex()
					CheckId(result["id"].(string), ID_object)
					return true
				}
			}
		}
	}
	return false
}

func main() {
	// auth.Initialize()
	if !Read_json() {
		color.Red("KEY NOT VALID")
		os.Exit(1)
	}

	utils.Logo()
	utils.Site_list()
	mode := utils.SelectMode("[Eagle " + version.Version + "]" + "[" + time.Now().Format("15:04:05.000000") + "]" + " PLEASE SELECT SITE:")
	if mode == "1" {
		print("GAMESTOP")
	} else if mode == "2" {
		deadstock.Menu_deadstock()
	} else if mode == "3" {
		print("DADSTOCK")
	} else {
		color.HiMagenta("[" + time.Now().Format("15:04:05.000000") + "] " + "INVALID CHOICE!")
	}
}

/*

---------BOT--------------------
// Add loader for all the info( profiles, proxies, key)
 improve the inizialization of ALL the data
 rich_presence
 create an executable file in golang
 change id security
 check how to compile go module and create the CONSOLE APP
 ADD APII
 SET CONSOLE LOG
 Function to generate all the file necessary to set up csv etc..
 Auto-updates
 Dashboard
 monitor
 client
 modules
 Scrape PID + put them encrypted
 Sniffer tipo proxyman, fiddler
 RANDOM Name + Surname
 Add MQT MONITOR MODE

---------ERROR_HANDLING----------------------

---------OTHERS-----------------------------
 Function to close the bot from remote

---------GUIDE-------------------------------
 1.TO find your ID you must activate Delepoer Mode. Goig in Setings-->advances-->Developer Mode
 after right-click to your profile picture and select Copy ID

---------WEB_SITES-------------------------------
 monitor
 early info / pid endpoint
 modulo



---------Structure-------------------------------
1. The architecture does not depend on the existence of some library of feature laden software.
2. Testable. The business rules can be tested without the UI, Database, Web Server, or any other external element.



func setConsoleTitle(title string) (int, error) {
	handle, err := syscall.LoadLibrary("Kernel32.dll")
	if err != nil {
		return 0, err
	}
	defer syscall.FreeLibrary(handle)
	proc, err := syscall.GetProcAddress(handle, "SetConsoleTitleW")
	if err != nil {
		return 0, err
	}
	r, _, err := syscall.Syscall(proc, 1, uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(title))), 0, 0)
	return int(r), err
}

func updateTitle() {
	_, _ = setConsoleTitle(fmt.Sprintf("HellasAIO ｜ Carts: %d ｜ Checkouts: %d ｜ Failures: %d", carts, checkouts, failures))
}
*/
