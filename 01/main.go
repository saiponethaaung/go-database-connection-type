package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"test/driver"
	"test/handler"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../.env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	sql := driver.ConnectSQL()

	mux := handler.DefaultMux(sql)

	hostname := ":" + os.Getenv("PORT")

	if os.Getenv("ENV") == "development" {
		fmt.Println("Serving at http://" + os.Getenv("HOST") + hostname)
	}

	http.ListenAndServe(hostname, mux)
}
