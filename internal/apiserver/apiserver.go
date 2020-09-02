package apiserver

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/omekov/exclusive/internal/apiserver/handlers"
	"github.com/omekov/exclusive/internal/apiserver/stores"
)

var (
	MONGO_URI             string = ""
	PORT                  string = "9090"
	MONGO_COLLECTION_NAME string = "example"
)

func initENV() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	MONGO_URI = os.Getenv("MONGO_URI")
	PORT = os.Getenv("PORT")
	MONGO_COLLECTION_NAME = os.Getenv("MONGO_COLLECTION_NAME")
}

func connection() {

}

// Run ...
func Run() {
	initENV()
	store := stores.Store{}
	store.ConfigMongo(MONGO_URI)
	router := handlers.Router()
	log.Fatal(http.ListenAndServe(":"+PORT, router))
}
