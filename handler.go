package main

import (
	"log"
	"net/http"
	"regexp"

	"github.com/rubens-schmitz/shop/category"
	"github.com/rubens-schmitz/shop/item"
	"github.com/rubens-schmitz/shop/product"
	"github.com/rubens-schmitz/shop/util"
)

var FS = http.FileServer(http.Dir("static"))

func handler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		util.VerifyCartId(w, r)
		logRequest(r)
		path := []byte(r.URL.Path)

		match, err := regexp.Match("/api/categories", path)
		if err != nil {
			log.Fatal(err)
		}
		if match {
			category.GetCategories(w, r)
			return
		}
		match, err = regexp.Match("/api/category", path)
		if err != nil {
			log.Fatal(err)
		}
		if match {
			switch r.Method {
			case "POST":
				category.PostCategory(w, r)
			case "GET":
				category.GetCategory(w, r)
			case "PUT":
				category.PutCategory(w, r)
			case "DELETE":
				category.DeleteCategory(w, r)
			}
			return
		}

		match, err = regexp.Match("/api/items", path)
		if err != nil {
			log.Fatal(err)
		}
		if match {
			item.GetItems(w, r)
			return
		}
		match, err = regexp.Match("/api/item", path)
		if err != nil {
			log.Fatal(err)
		}
		if match {
			switch r.Method {
			case "POST":
				item.PostItem(w, r)
			case "PUT":
				item.PutItem(w, r)
			case "DELETE":
				item.DeleteItem(w, r)
			}
			return
		}

		match, err = regexp.Match("/api/products", path)
		if err != nil {
			log.Fatal(err)
		}
		if match {
			product.GetProducts(w, r)
			return
		}
		match, err = regexp.Match("/api/product", path)
		if err != nil {
			log.Fatal(err)
		}
		if match {
			switch r.Method {
			case "POST":
				product.PostProduct(w, r)
			case "GET":
				product.GetProduct(w, r)
			case "PUT":
				product.PutProduct(w, r)
			case "DELETE":
				product.DeleteProduct(w, r)
			}
			return
		}

		FS.ServeHTTP(w, r)
	})
}
