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
)

type ErrorResponse struct {
	Ok    bool   `json:"id"`
	Error string `json:"error"`
}

var DB *sql.DB
var FS = http.FileServer(http.Dir("static"))

func logRequest(r *http.Request) {
	urlValues := r.URL.Query()
	cartId := getCartId(r)
	if len(urlValues) == 0 {
		log.Printf("%v %v %v\n", cartId, r.Method, r.URL.Path)
	} else {
		log.Printf("%v %v %v %v\n", cartId, r.Method, r.URL.Path, urlValues)
	}
}

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
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	err = DB.Ping()
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
