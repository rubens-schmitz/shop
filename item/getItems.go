package item

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/rubens-schmitz/shop/util"
)

type GetItemResponse struct {
	Id        int      `json:"id"`
	ProductId int      `json:"productId"`
	Title     string   `json:"title"`
	Price     float32  `json:"price"`
	Pictures  []string `json:"pictures"`
	Quantity  int      `json:"quantity"`
}

type GetItemsParams struct {
	Title      string
	Offset     int
	Limit      int
	CategoryId int
	CartId     int
}

func getParams(w http.ResponseWriter, r *http.Request) GetItemsParams {
	title := util.GetStringParam(r, "title")
	limit, err := util.GetIntParam(r, "limit")
	if err != nil {
		log.Fatal(err)
	}
	offset, err := util.GetIntParam(r, "offset")
	if err != nil {
		log.Fatal(err)
	}
	categoryId, err := util.GetIntParam(r, "categoryId")
	if err != nil {
		log.Fatal(err)
	}
	cartId := util.GetCartId(w, r)
	return GetItemsParams{Title: title, Limit: limit, Offset: offset,
		CategoryId: categoryId, CartId: cartId}
}

func getRows(params GetItemsParams) *sql.Rows {
	var query string
	var rows *sql.Rows
	var err error
	if params.CategoryId == 0 {
		query = `select item.id, product.id, product.title, product.price,
				 item.quantity from product inner join item on
				 product.id = item.productId where item.cartId = $1
				 and lower(product.title) like lower('%' || $2 || '%')
				 order by title asc limit $3 offset $4`
		rows, err = util.DB.Query(query, params.CartId, params.Title,
			params.Limit, params.Offset)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		query = `select item.id, product.id, product.title, product.price,
				 item.quantity from product inner join item on
				 product.id = item.productId where item.cartId = $1
				 and lower(product.title) like lower('%' || $2 || '%') 
				 and categoryId = $3
				 order by title asc limit $4 offset $5`
		rows, err = util.DB.Query(query, params.CartId, params.Title,
			params.CategoryId, params.Limit, params.Offset)
		if err != nil {
			log.Fatal(err)
		}
	}
	return rows
}

func reallyGetItems(rows *sql.Rows) []GetItemResponse {
	items := make([]GetItemResponse, 0)
	for rows.Next() {
		item := new(GetItemResponse)
		err := rows.Scan(&item.Id, &item.ProductId, &item.Title, &item.Price,
			&item.Quantity)
		if err != nil {
			log.Fatal(err)
		}
		item.Pictures = util.GetPictures(item.ProductId)
		items = append(items, *item)
	}
	return items
}

func GetItems(w http.ResponseWriter, r *http.Request) {
	params := getParams(w, r)
	rows := getRows(params)
	defer rows.Close()
	items := reallyGetItems(rows)
	util.WriteAsJSON(w, items)
}
