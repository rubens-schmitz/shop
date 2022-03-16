package category

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/rubens-schmitz/shop/types"
	"github.com/rubens-schmitz/shop/util"
)

func parseParams(r *http.Request) (types.GetCategoriesParams, error) {
	limit, err := util.GetIntParam(r, "limit")
	if err != nil {
		return types.GetCategoriesParams{}, err
	}
	offset, err := util.GetIntParam(r, "offset")
	if err != nil {
		return types.GetCategoriesParams{}, err
	}
	title := util.GetStringParam(r, "title")
	params := types.GetCategoriesParams{Limit: limit, Offset: offset,
		Title: title}
	return params, nil
}

func queryRows(params types.GetCategoriesParams) *sql.Rows {
	var query string
	var rows *sql.Rows
	var err error
	if params.Limit == 0 {
		query = `select id, title from category where deleted = false`
		rows, err = util.DB.Query(query)
	} else {
		query = `select id, title from category where deleted = false
				 and lower(title) like lower('%' || $1 || '%')
				 limit $2 offset $3`
		rows, err = util.DB.Query(query, params.Title,
			params.Limit, params.Offset)
	}
	if err != nil {
		log.Fatal(err)
	}
	return rows
}

func makeCategories(rows *sql.Rows) []types.GetCategoryResponse {
	categories := make([]types.GetCategoryResponse, 0)
	for rows.Next() {
		category := new(types.GetCategoryResponse)
		err := rows.Scan(&category.Id, &category.Title)
		if err != nil {
			log.Fatal(err)
		}
		categories = append(categories, *category)
	}
	return categories
}

func GetCategoriesHandler(w http.ResponseWriter, r *http.Request) {
	params, err := parseParams(r)
	if err != nil {
		util.WriteAsJSON(w, types.SuccessResponse{Success: false,
			Msg: err.Error()})
		return
	}
	rows := queryRows(params)
	defer rows.Close()
	categories := makeCategories(rows)
	util.WriteAsJSON(w, categories)
}
