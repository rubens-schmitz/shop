package main

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"
)

type PostItemRequest struct {
	Id        int `json:"id"`
	ProductId int `json:"productId"`
	CartId    int `json:"cartId"`
	Quantity  int `json:"quantity"`
}

type GetItemResponse struct {
	Id        int      `json:"id"`
	ProductId int      `json:"productId"`
	Title     string   `json:"title"`
	Price     float32  `json:"price"`
	Pictures  []string `json:"pictures"`
	Quantity  int      `json:"quantity"`
}

func postItem(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		log.Fatal(err)
	}
	productId, err := strconv.Atoi(r.FormValue("productId"))
	if err != nil {
		log.Fatal(err)
	}
	cartId := getCartId(r)
	item := &PostItemRequest{ProductId: productId, CartId: cartId}
	query := `select id, quantity from item where productId=$1 and cartId=$2`
	row := DB.QueryRow(query, productId, cartId)
	err = row.Scan(&item.Id, &item.Quantity)
	if err != nil {
		if err == sql.ErrNoRows {
			query = "insert into item (productId, cartId) values ($1, $2)"
			_, err = DB.Exec(query, productId, cartId)
			if err != nil {
				log.Println(err)
				writeAsJSON(w, &ErrorResponse{Ok: false, Error: err.Error()})
				return
			}
		} else {
			log.Fatal(err)
		}
	}
	item.Quantity += 1
	query = "update item set quantity = $2 where id = $1"
	_, err = DB.Exec(query, item.Id, item.Quantity)
	if err != nil {
		log.Fatal(err)
	}
	writeAsJSON(w, &ErrorResponse{Ok: true, Error: ""})
}

func getItems(w http.ResponseWriter, r *http.Request) {
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

	cartId := strconv.Itoa(getCartId(r))

	var query string
	var rows *sql.Rows
	if categoryId == 0 {
		query = `select item.id, product.id, product.title, product.price,
				 item.quantity from product inner join item on
				 product.id = item.productId where item.cartId = $1
				 and lower(product.title) like lower('%' || $2 || '%')
				 limit 2 offset $3`
		rows, err = DB.Query(query, cartId, title, offset)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		query = `select item.id, product.id, product.title, product.price,
				 item.quantity from product inner join item on
				 product.id = item.productId where categoryId = $1 and 
				 item.cartId = $2 and lower(product.title) like 
				 lower('%' || $3 || '%') limit 2 offset $4`
		rows, err = DB.Query(query, categoryId, cartId, title, offset)
		if err != nil {
			log.Fatal(err)
		}
	}
	defer rows.Close()

	items := make([]GetItemResponse, 0)
	for rows.Next() {
		item := new(GetItemResponse)
		err := rows.Scan(&item.Id, &item.ProductId, &item.Title, &item.Price,
			&item.Quantity)
		if err != nil {
			log.Fatal(err)
		}
		item.Pictures = getPictures(item.ProductId)
		items = append(items, *item)
	}
	writeAsJSON(w, items)
}

func putItem(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		log.Fatal(err)
	}
	id := r.FormValue("id")
	quantity, err := strconv.Atoi(r.FormValue("quantity"))
	if err != nil {
		log.Fatal(err)
	}
	if quantity <= 0 {
		writeAsJSON(w, &ErrorResponse{
			Ok: false, Error: "Parameter 'quantity' is less than or equal zero."})
		return
	}
	query := `update item set quantity = $1 where id = $2`
	_, err = DB.Exec(query, quantity, id)
	if err != nil {
		log.Fatal(err)
	}
	writeAsJSON(w, &ErrorResponse{Ok: true, Error: ""})
}

func deleteItem(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		log.Fatal(err)
	}
	id := r.FormValue("id")
	query := "delete from item where id=$1"
	_, err = DB.Exec(query, id)
	if err != nil {
		log.Fatal(err)
	}
	writeAsJSON(w, &ErrorResponse{Ok: true, Error: ""})
}
