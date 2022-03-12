package main

import (
	"log"
	"net/http"
	"regexp"

	"github.com/rubens-schmitz/shop/cart"
	"github.com/rubens-schmitz/shop/category"
	"github.com/rubens-schmitz/shop/deal"
	"github.com/rubens-schmitz/shop/item"
	"github.com/rubens-schmitz/shop/product"
)

var FS = http.FileServer(http.Dir("static"))

func logRequest(w http.ResponseWriter, r *http.Request) {
	urlValues := r.URL.Query()
	cartId := cart.GetCartId(w, r)
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
			category.GetCategoriesHandler(w, r)
			return
		}
		match, err = regexp.Match("/api/category", path)
		if err != nil {
			log.Fatal(err)
		}
		if match {
			switch r.Method {
			case "POST":
				category.PostCategoryHandler(w, r)
			case "GET":
				category.GetCategoryHandler(w, r)
			case "PUT":
				category.PutCategoryHandler(w, r)
			case "DELETE":
				category.DeleteCategoryHandler(w, r)
			}
			return
		}

		match, err = regexp.Match("/api/cart", path)
		if err != nil {
			log.Fatal(err)
		}
		if match {
			switch r.Method {
			case "GET":
				cart.GetCartHandler(w, r)
			}
			return
		}

		match, err = regexp.Match("/api/items", path)
		if err != nil {
			log.Fatal(err)
		}
		if match {
			item.GetItemsHandler(w, r)
			return
		}
		match, err = regexp.Match("/api/item", path)
		if err != nil {
			log.Fatal(err)
		}
		if match {
			switch r.Method {
			case "POST":
				item.PostItemHandler(w, r)
			case "PUT":
				item.PutItemHandler(w, r)
			case "DELETE":
				item.DeleteItemHandler(w, r)
			}
			return
		}

		match, err = regexp.Match("/api/products", path)
		if err != nil {
			log.Fatal(err)
		}
		if match {
			product.GetProductsHandler(w, r)
			return
		}
		match, err = regexp.Match("/api/product", path)
		if err != nil {
			log.Fatal(err)
		}
		if match {
			switch r.Method {
			case "POST":
				product.PostProductHandler(w, r)
			case "GET":
				product.GetProductHandler(w, r)
			case "PUT":
				product.PutProductHandler(w, r)
			case "DELETE":
				product.DeleteProductHandler(w, r)
			}
			return
		}

		match, err = regexp.Match("/api/deals", path)
		if err != nil {
			log.Fatal(err)
		}
		if match {
			deal.GetDealsHandler(w, r)
			return
		}
		match, err = regexp.Match("/api/deal", path)
		if err != nil {
			log.Fatal(err)
		}
		if match {
			switch r.Method {
			case "POST":
				deal.PostDealHandler(w, r)
			case "GET":
				deal.GetDealHandler(w, r)
			case "DELETE":
				deal.DeleteDealHandler(w, r)
			}
			return
		}

		FS.ServeHTTP(w, r)
	})
}
