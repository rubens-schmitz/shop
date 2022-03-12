package item

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/rubens-schmitz/shop/cart"
	"github.com/rubens-schmitz/shop/types"
	"github.com/rubens-schmitz/shop/util"
)

func parseParams(w http.ResponseWriter, r *http.Request) types.GetItemsParams {
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
	cartId, err := util.GetIntParam(r, "cartId")
	if err != nil {
		log.Fatal(err)
	}
	if cartId == 0 {
		cartId = cart.GetCartId(w, r)
	}
	return types.GetItemsParams{Title: title, Limit: limit, Offset: offset,
		CategoryId: categoryId, CartId: cartId}
}

func queryRows(params types.GetItemsParams) *sql.Rows {
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

func makeItems(rows *sql.Rows) []types.GetItemResponse {
	items := make([]types.GetItemResponse, 0)
	for rows.Next() {
		item := new(types.GetItemResponse)
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

func GetItems(params types.GetItemsParams) []types.GetItemResponse {
	rows := queryRows(params)
	defer rows.Close()
	items := makeItems(rows)
	return items
}

func GetItemsHandler(w http.ResponseWriter, r *http.Request) {
	params := parseParams(w, r)
	items := GetItems(params)
	util.WriteAsJSON(w, items)
}
