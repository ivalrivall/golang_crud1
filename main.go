package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ivalrivall/golang_crud1/router"
	"github.com/ivalrivall/golang_crud1/seeds"
	"github.com/joho/godotenv"
)

func main() {
	r := router.Router()
	// fs := http.FileServer(http.Dir("build"))
	// http.Handle("/", fs)
	fmt.Println("Starting server on the port 8080...")
	handleArgs()
	log.Fatal(http.ListenAndServe(":8080", r))
}

func handleArgs() {
	flag.Parse()
	args := flag.Args()

	if len(args) >= 1 {
		switch args[0] {
		case "seed":
			err := godotenv.Load(".env")

			if err != nil {
				log.Fatalf("Error loading .env file")
			}

			// Open the connection
			db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))

			if err != nil {
				panic(err)
			}

			seeds.Execute(db, args[1:]...)
			os.Exit(0)
		}
	}
}
