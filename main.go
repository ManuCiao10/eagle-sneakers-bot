package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/denisbrodbeck/machineid"
	"github.com/joho/godotenv"
)

func init() {
	//add check updates
	err := godotenv.Load("config/.env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func gen_id() string {
	id, err := machineid.ProtectedID("myAppName")
	if err != nil {
		log.Fatal(err)
	}
	return id
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
			{Key: "$set", Value: bson.D{{Key: "id", Value: gen_id()}}},
		},
	)
	_ = res
	if err != nil {
		color.HiMagenta("[ " + time.Now().Format("15:04:05.000000") + " ]" + "OPEN TICKET ERROR TO UPDATE KEY")
		os.Exit(1)
	}
	color.HiMagenta("[ " + time.Now().Format("15:04:05.000000") + " ]" + " ID UPDATED")
}

func CheckId(id_database string, ID_object string) {
	if gen_id() == id_database {
		color.HiMagenta("[ " + time.Now().Format("15:04:05.000000") + " ]" + " ID VALID")
	} else {
		color.HiMagenta("[ " + time.Now().Format("15:04:05.000000") + " ]" + " ID NOT VALID")
		answer := SelectMode("[ " + time.Now().Format("15:04:05.000000") + " ]" + " DO YOU WANT TO RESET IT? (Y/N): ")
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

func SelectMode(label string) string {
	var s string
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Fprint(os.Stderr, label+" ")
		s, _ = r.ReadString('\n')
		if s != "" {
			break
		}
	}
	return strings.TrimSpace(s)
}

func Read_json() bool {
	var key string
	content, err := os.ReadFile("./setting.json")
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
	color.Red("[ " + time.Now().Format("15:04:05.000000") + " ]" + " CHECKING KEY...")
	if !Read_json() {
		color.Red("KEY NOT VALID")
		os.Exit(1)
	}
	content, err := os.ReadFile("config/logo.txt")
	if err != nil {
		log.Fatal(err)
	}
	color.Red(string(content))

	color.Red("[ Eagle 0.0.2 ]" + "[ " + time.Now().Format("15:04:05.000000") + " ]" + " 1. GAMESTOP")
	color.Red("[ Eagle 0.0.2 ]" + "[ " + time.Now().Format("15:04:05.000000") + " ]" + " 2. UNIEURO")
	color.Red("[ Eagle 0.0.2 ]" + "[ " + time.Now().Format("15:04:05.000000") + " ]" + " 3. DADSTOCK")
	println("\n")

	mode := SelectMode("[ Eagle 0.0.2 ]" + "[ " + time.Now().Format("15:04:05.000000") + " ]" + " PLEASE SELECT SITE:")
	if mode == "1" {
		print("GAMESTOP")
	} else if mode == "2" {
		print("UNIEURO")
	} else if mode == "3" {
		print("DADSTOCK")
	} else {
		print("INVALID CHOICE!\n") //ADD RESTART BOT
	}
	// fmt.Println(mode)
}

// clean code adding utils.go
// Add Dashboard
// Add monitor
// Add client
// Add sql database

//------------------------------------------------------------------//
// ADD a guide to get the uuid
