package util

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type ErrorResponse struct {
	Ok    bool   `json:"id"`
	Error string `json:"error"`
}

var DB *sql.DB

func WriteAsJSON(w http.ResponseWriter, v any) {
	data, err := json.Marshal(v)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(data)
	if err != nil {
		log.Fatal(err)
	}
}

func VerifyCartId(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("cartId")
	if err != nil {
		query := "insert into cart default values returning id"
		row := DB.QueryRow(query)
		var id string
		err := row.Scan(&id)
		if err != nil {
			log.Fatal(err)
		}
		cookie := http.Cookie{Name: "cartId", Value: id, Path: "/"}
		http.SetCookie(w, &cookie)
		r.AddCookie(&cookie)
	}
}

func GetCartId(r *http.Request) int {
	cookie, err := r.Cookie("cartId")
	if err != nil {
		log.Fatal(err)
	}
	id, err := strconv.Atoi(cookie.Value)
	if err != nil {
		log.Fatal(err)
	}
	return id
}

func GetPictures(productId int) []string {
	pictures := make([]string, 0)
	query := "select id, bytes from picture where productId = $1"
	rows, err := DB.Query(query, productId)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var id int
		var bytes []byte
		err := rows.Scan(&id, &bytes)
		if err != nil {
			log.Fatal(err)
		}
		pictures = append(pictures, string(bytes))
	}
	return pictures
}

func GetIntParam(r *http.Request, name string) (int, error) {
	arr := r.URL.Query()[name]
	val := 0
	var err error
	if len(arr) != 0 {
		val, err = strconv.Atoi(arr[0])
		if err != nil {
			log.Fatal(err)
		}
		if val < 0 {
			s := fmt.Sprintf("Parameter '%v' is less than zero.", name)
			return 0, errors.New(s)
		}
	}
	return val, nil
}

func GetStringParam(r *http.Request, name string) string {
	arr := r.URL.Query()[""]
	val := ""
	if len(arr) != 0 {
		val = arr[0]
	}
	return val
}
