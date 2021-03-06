package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	cp "github.com/otiai10/copy"

	"github.com/rubens-schmitz/shop/util"
)

func build() {
	err := os.Chdir("client")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Building client")
	cmd := exec.Command("npm", "run", "build")
	err = cmd.Run()
	if err != nil {
		log.Fatal("Build failed")
	}

	err = os.Chdir("..")
	if err != nil {
		log.Fatal(err)
	}

	err = cp.Copy("client/build", "static")
	if err != nil {
		log.Fatal(err)
	}
}

func connect() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	user := os.Getenv("DBUSER")
	password := os.Getenv("DBPASSWORD")
	dbname := os.Getenv("DBNAME")
	connStr := fmt.Sprintf("user=%v password=%v dbname=%v", user, password, dbname)
	util.DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	err = util.DB.Ping()
	if err != nil {
		log.Fatal(err)
	}
}

func serve() {
	port := flag.String("p", "3000", "port to serve on")
	flag.Parse()
	log.Printf("Serving on localhost:%s\n", *port)
	log.Fatal(http.ListenAndServe(":"+*port, handler()))
}

func main() {
	log.SetFlags(log.Lshortfile)
	build()
	connect()
	serve()
}
