package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type GetProductResponse struct {
	Id         int      `json:"id"`
	Title      string   `json:"title"`
	Price      float32  `json:"price"`
	CategoryId int      `json:"categoryId"`
	Pictures   []string `json:"pictures"`
}

var errNoProduct = errors.New("The requested product does not exist.")

func postProduct(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		log.Fatal(err)
	}
	title := r.FormValue("title")
	price := r.FormValue("price")
	categoryId := r.FormValue("categoryId")
	query := `insert into product (title, price, categoryId)
	          values ($1, $2, $3) returning id`
	row := DB.QueryRow(query, title, price, categoryId)
	var id int
	err = row.Scan(&id)
	if err != nil {
		log.Fatal(err)
	}
	pictures, err := strconv.Atoi(r.FormValue("pictures"))
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < pictures; i++ {
		bytes := extractPicture(r, i)
		query = "insert into picture (productId, bytes) values ($1, $2)"
		_, err = DB.Exec(query, id, bytes)
	}
	writeAsJSON(w, &ErrorResponse{Ok: true, Error: ""})
}

func extractPicture(r *http.Request, n int) []byte {
	name := fmt.Sprintf("pictures%v", n)
	picture := r.FormValue(name)
	return []byte(picture)
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	titles := r.URL.Query()["title"]
	title := ""
	if len(titles) != 0 {
		title = titles[0]
	}

	offsets := r.URL.Query()["offset"]
	offset := 0
	var err error
	if len(offsets) != 0 {
		offset, err = strconv.Atoi(offsets[0])
		if err != nil {
			log.Fatal(err)
		}
		if offset < 0 {
			writeAsJSON(w, &ErrorResponse{
				Ok: false, Error: "Parameter 'offset' is less than zero."})
			return
		}
	}

	categories := r.URL.Query()["categoryId"]
	categoryId := 0
	if len(categories) != 0 {
		categoryId, err = strconv.Atoi(categories[0])
		if err != nil {
			log.Fatal(err)
		}
		if offset < 0 {
			writeAsJSON(w, &ErrorResponse{
				Ok: false, Error: "Parameter 'categoryId' is less than zero."})
			return
		}
	}

	var query string
	var rows *sql.Rows
	if categoryId == 0 {
		query = `select * from product where lower(title) 
				 like lower('%' || $1 || '%') limit 2 offset $2`
		rows, err = DB.Query(query, title, offset)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		query = `select * from product where categoryId = $1 and lower(title) 
				 like lower('%' || $2 || '%') limit 2 offset $3`
		rows, err = DB.Query(query, categoryId, title, offset)
		if err != nil {
			log.Fatal(err)
		}
	}
	defer rows.Close()

	products := make([]GetProductResponse, 0)
	for rows.Next() {
		product := new(GetProductResponse)
		err := rows.Scan(&product.Id, &product.Title,
			&product.Price, &product.CategoryId)
		if err != nil {
			log.Fatal(err)
		}
		product.Pictures = getPictures(product.Id)
		products = append(products, *product)
	}
	writeAsJSON(w, products)
}

func getPictures(productId int) []string {
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

func getProduct(w http.ResponseWriter, r *http.Request) {
	ids, ok := r.URL.Query()["id"]
	if !ok || len(ids[0]) < 1 {
		writeAsJSON(w, &ErrorResponse{
			Ok: false, Error: "Url Param 'id' is missing"})
		return
	}
	idStr := ids[0]
	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println(err)
		return
	}
	pictures := getPictures(idInt)
	product := &GetProductResponse{Id: idInt, Pictures: pictures}
	query := "select title, price, categoryId from product where id=$1"
	row := DB.QueryRow(query, idStr)
	err = row.Scan(&product.Title, &product.Price, &product.CategoryId)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println(err)
			return
		}
		log.Fatal(err)
	}
	writeAsJSON(w, product)
}

func putProduct(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		log.Fatal(err)
	}
	id := r.FormValue("id")
	title := r.FormValue("title")
	price := r.FormValue("price")
	categoryId := r.FormValue("categoryId")

	query := `update product set title = $1, price = $2, categoryId = $3
	          where id = $4`
	_, err = DB.Exec(query, title, price, categoryId, id)

	query = "delete from picture where productId = $1"
	_, err = DB.Exec(query, id)

	pictures, err := strconv.Atoi(r.FormValue("pictures"))
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < pictures; i++ {
		bytes := extractPicture(r, i)
		query = "insert into picture (productId, bytes) values ($1, $2)"
		_, err = DB.Exec(query, id, bytes)
	}

	writeAsJSON(w, &ErrorResponse{Ok: true, Error: ""})
}

func deleteProduct(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		log.Fatal(err)
	}
	id := r.FormValue("id")
	query := "delete from product where id=$1"
	_, err = DB.Exec(query, id)
	if err != nil {
		log.Fatal(err)
	}
	writeAsJSON(w, &ErrorResponse{Ok: true, Error: ""})
}

func makeProduct(id int) (*GetProductResponse, error) {
	idStr := strconv.Itoa(id)
	pictures := getPictures(id)
	product := &GetProductResponse{Id: id, Pictures: pictures}
	query := "select title, price, categoryId from product where id=$1"
	row := DB.QueryRow(query, idStr)
	err := row.Scan(&product.Title, &product.Price, &product.CategoryId)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println(err)
			return nil, errNoProduct
		}
		log.Fatal(err)
	}
	return product, nil
}
