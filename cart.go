package main

import (
	"log"
	"net/http"
	"strconv"
)

func verifyCartId(w http.ResponseWriter, r *http.Request) {
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

func getCartId(r *http.Request) int {
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
