package config

import (

    "fmt"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
    "os"
    "github.com/joho/godotenv"
)

func EnvMongoURI() string {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error in loading the .env file")
    }
  
    return os.Getenv("MONGOURI")
}

func Connectdb() {
	err := mgm.SetDefaultConfig(nil, "Order_management", options.Client().ApplyURI(EnvMongoURI()))
	fmt.Println("connected to the database............")
	fmt.Println(err)
 }