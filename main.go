package main

import (
	"log"
	"os"
	"strconv"
)

func main() {
	a := App{}

	port, err := strconv.Atoi(os.Getenv("APP_DB_PORT"))
	if err != nil {
		log.Println("Port not well specified. Using 5432")
		//log.Fatal(err)
		port = 5432
	}
	a.Initialize(
		os.Getenv("APP_DB_HOST"),
		port,
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"))

	a.Run("localhost:8080")
}
