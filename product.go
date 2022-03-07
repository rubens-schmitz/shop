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

type GetProductsParams struct {
	Limit      int
	Offset     int
	CategoryId int
	Title      string
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

func getIntParam(r *http.Request, name string) (int, error) {
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

func getStringParam(r *http.Request, name string) string {
	arr := r.URL.Query()[""]
	val := ""
	if len(arr) != 0 {
		val = arr[0]
	}
	return val
}

func parseParams(r *http.Request) (GetProductsParams, error) {
	limit, err := getIntParam(r, "limit")
	offset, err := getIntParam(r, "offset")
	categoryId, err := getIntParam(r, "categoryId")
	title := getStringParam(r, "title")
	params := &GetProductsParams{Limit: limit, Offset: offset,
		CategoryId: categoryId, Title: title}
	return *params, err
}

func getRows(params GetProductsParams) *sql.Rows {
	var query string
	var rows *sql.Rows
	var err error
	if params.CategoryId == 0 {
		query = `select * from product where lower(title) 
				 like lower('%' || $1 || '%') limit $2 offset $3`
		rows, err = DB.Query(query, params.Title, params.Limit, params.Offset)
	} else {
		query = `select * from product where categoryId = $1 and lower(title) 
				 like lower('%' || $2 || '%') limit $3 offset $4`
		rows, err = DB.Query(query, params.CategoryId, params.Title,
			params.Limit, params.Offset)
	}
	if err != nil {
		log.Fatal(err)
	}
	return rows
}

func reallyGetProducts(rows *sql.Rows) []GetProductResponse {
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
	return products
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	params, err := parseParams(r)
	if err != nil {
		writeAsJSON(w, &ErrorResponse{
			Ok: false, Error: err.Error()})
		return
	}
	rows := getRows(params)
	defer rows.Close()
	products := reallyGetProducts(rows)
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
