package api

import (
	"fmt"
	"log"
	"os"

	"github.com/ibnanjung/job2go_todo/api/controllers"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var server = controllers.Server{}

//Run method
func Run() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatal("error get env file")
	} else {
		fmt.Println("got the env")
	}
	server.Initialize(os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	server.Run(":8080")

}
