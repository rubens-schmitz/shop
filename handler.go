package main

import (
	"log"
	"net/http"
	"regexp"

	"github.com/rubens-schmitz/shop/category"
	"github.com/rubens-schmitz/shop/deal"
	"github.com/rubens-schmitz/shop/item"
	"github.com/rubens-schmitz/shop/product"
	"github.com/rubens-schmitz/shop/util"
)

var FS = http.FileServer(http.Dir("static"))

func logRequest(w http.ResponseWriter, r *http.Request) {
	urlValues := r.URL.Query()
	cartId := util.GetCartId(w, r)
	if len(urlValues) == 0 {
		log.Printf("%v %v %v\n", cartId, r.Method, r.URL.Path)
	} else {
		log.Printf("%v %v %v %v\n", cartId, r.Method, r.URL.Path, urlValues)
	}
}

func handler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logRequest(w, r)
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

		match, err = regexp.Match("/api/deals", path)
		if err != nil {
			log.Fatal(err)
		}
		if match {
			deal.GetDeals(w, r)
			return
		}
		match, err = regexp.Match("/api/deal", path)
		if err != nil {
			log.Fatal(err)
		}
		if match {
			switch r.Method {
			case "POST":
				deal.PostDeal(w, r)
			case "GET":
				deal.GetDeal(w, r)
			case "DELETE":
				deal.DeleteDeal(w, r)
			}
			return
		}

		FS.ServeHTTP(w, r)
	})
}
