package api

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/abdghn/widi-group-backend/api/controllers"

	"github.com/joho/godotenv"
)

var server = controllers.Server{}

func Run() {

	var err error
	err = godotenv.Load()
	if err != nil {
		ioutil.WriteFile(".env", []byte(""), 0755)
		// log.Fatalf("Error getting env, not comming through %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}
	_, err = os.Stat("files")

	if os.IsNotExist(err) {
		errDir := os.MkdirAll("files", 0755)
		if errDir != nil {
			log.Fatal(err)
		}

	}

	server.Initialize(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))
	// auto migration
	// seed.Load(server.DB)

	// server.Run(os.Getenv("PORT"))
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	server.Run(":" + port)

}
