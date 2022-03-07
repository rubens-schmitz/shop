package category

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/rubens-schmitz/shop/util"
)

type GetCategoriesParams struct {
	Limit  int
	Offset int
	Title  string
}

func parseParams(r *http.Request) (GetCategoriesParams, error) {
	limit, err := util.GetIntParam(r, "limit")
	if err != nil {
		return GetCategoriesParams{}, err
	}
	offset, err := util.GetIntParam(r, "offset")
	if err != nil {
		return GetCategoriesParams{}, err
	}
	title := util.GetStringParam(r, "title")
	params := GetCategoriesParams{Limit: limit, Offset: offset, Title: title}
	return params, nil
}

func getRows(params GetCategoriesParams) *sql.Rows {
	var query string
	var rows *sql.Rows
	var err error
	if params.Limit == 0 {
		query = `select * from category`
		rows, err = util.DB.Query(query)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		query = `select * from category where lower(title) 
				 like lower('%' || $1 || '%') limit $2 offset $3`
		rows, err = util.DB.Query(query, params.Title,
			params.Limit, params.Offset)
		if err != nil {
			log.Fatal(err)
		}
	}
	return rows
}

func reallyGetCategories(rows *sql.Rows) []GetCategoryResponse {
	categories := make([]GetCategoryResponse, 0)
	for rows.Next() {
		category := new(GetCategoryResponse)
		err := rows.Scan(&category.Id, &category.Title)
		if err != nil {
			log.Fatal(err)
		}
		categories = append(categories, *category)
	}
	return categories
}

func GetCategories(w http.ResponseWriter, r *http.Request) {
	params, err := parseParams(r)
	if err != nil {
		util.WriteAsJSON(w, &util.ErrorResponse{Ok: false, Error: err.Error()})
		return
	}
	rows := getRows(params)
	defer rows.Close()
	categories := reallyGetCategories(rows)
	util.WriteAsJSON(w, categories)
}
