package product

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/rubens-schmitz/shop/types"
	"github.com/rubens-schmitz/shop/util"
)

func parseParams(r *http.Request) (types.GetProductsParams, error) {
	limit, err := util.GetIntParam(r, "limit")
	if err != nil {
		return types.GetProductsParams{}, err
	}
	offset, err := util.GetIntParam(r, "offset")
	if err != nil {
		return types.GetProductsParams{}, err
	}
	categoryId, err := util.GetIntParam(r, "categoryId")
	if err != nil {
		return types.GetProductsParams{}, err
	}
	title := util.GetStringParam(r, "title")
	params := types.GetProductsParams{Limit: limit, Offset: offset,
		CategoryId: categoryId, Title: title}
	return params, nil
}

func queryRows(params types.GetProductsParams) *sql.Rows {
	var query string
	var rows *sql.Rows
	var err error
	if params.CategoryId == 0 {
		query = `select id, title, price, categoryId from product
				 where lower(title) like lower('%' || $1 || '%')
				 order by title asc limit $2 offset $3`
		rows, err = util.DB.Query(query, params.Title,
			params.Limit, params.Offset)
	} else {
		query = `select id, title, price, categoryId from product
				 where lower(title) like lower('%' || $1 || '%')
				 and categoryId = $2 
				 order by title asc limit $3 offset $4`
		rows, err = util.DB.Query(query, params.Title, params.CategoryId,
			params.Limit, params.Offset)
	}
	if err != nil {
		log.Fatal(err)
	}
	return rows
}

func makeProducts(rows *sql.Rows) []types.GetProductResponse {
	products := make([]types.GetProductResponse, 0)
	for rows.Next() {
		product := new(types.GetProductResponse)
		err := rows.Scan(&product.Id, &product.Title,
			&product.Price, &product.CategoryId)
		if err != nil {
			log.Fatal(err)
		}
		product.Pictures = util.GetPictures(product.Id)
		products = append(products, *product)
	}
	return products
}

func GetProductsHandler(w http.ResponseWriter, r *http.Request) {
	params, err := parseParams(r)
	if err != nil {
		util.WriteAsJSON(w, util.ErrorResponse{Ok: false, Error: err.Error()})
		return
	}
	rows := queryRows(params)
	defer rows.Close()
	products := makeProducts(rows)
	util.WriteAsJSON(w, products)
}
