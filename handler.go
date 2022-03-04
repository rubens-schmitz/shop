package main

import (
	"log"
	"net/http"
	"regexp"
)

func handler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		verifyCartId(w, r)
		logRequest(r)
		path := []byte(r.URL.Path)

		match, err := regexp.Match("/api/categories", path)
		if err != nil {
			log.Fatal(err)
		}
		if match {
			getCategories(w, r)
			return
		}
		match, err = regexp.Match("/api/category", path)
		if err != nil {
			log.Fatal(err)
		}
		if match {
			switch r.Method {
			case "POST":
				postCategory(w, r)
			case "GET":
				getCategory(w, r)
			case "PUT":
				putCategory(w, r)
			case "DELETE":
				deleteCategory(w, r)
			}
			return
		}

		match, err = regexp.Match("/api/items", path)
		if err != nil {
			log.Fatal(err)
		}
		if match {
			getItems(w, r)
			return
		}
		match, err = regexp.Match("/api/item", path)
		if err != nil {
			log.Fatal(err)
		}
		if match {
			switch r.Method {
			case "POST":
				postItem(w, r)
			case "PUT":
				putItem(w, r)
			case "DELETE":
				deleteItem(w, r)
			}
			return
		}

		match, err = regexp.Match("/api/products", path)
		if err != nil {
			log.Fatal(err)
		}
		if match {
			getProducts(w, r)
			return
		}
		match, err = regexp.Match("/api/product", path)
		if err != nil {
			log.Fatal(err)
		}
		if match {
			switch r.Method {
			case "POST":
				postProduct(w, r)
			case "GET":
				getProduct(w, r)
			case "PUT":
				putProduct(w, r)
			case "DELETE":
				deleteProduct(w, r)
			}
			return
		}

		FS.ServeHTTP(w, r)
	})
}
