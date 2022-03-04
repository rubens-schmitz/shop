package main

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"
)

type GetCategoryResponse struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

func postCategory(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		log.Fatal(err)
	}
	title := r.FormValue("title")
	query := "insert into category (title) values ($1)"
	_, err = DB.Exec(query, title)
	if err != nil {
		log.Fatal(err)
	}
	writeAsJSON(w, &ErrorResponse{Ok: true, Error: ""})
}

func getCategories(w http.ResponseWriter, r *http.Request) {
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

	limits := r.URL.Query()["limit"]
	limit := 0
	if len(limits) != 0 {
		limit, err = strconv.Atoi(limits[0])
		if err != nil {
			log.Fatal(err)
		}
		if limit < 0 {
			writeAsJSON(w, &ErrorResponse{
				Ok: false, Error: "Parameter 'limit' is less than zero."})
			return
		}
	}

	var query string
	var rows *sql.Rows
	if limit == 0 {
		query = `select * from category`
		rows, err = DB.Query(query)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		query = `select * from category where lower(title) 
				 like lower('%' || $1 || '%') limit $2 offset $3`
		rows, err = DB.Query(query, title, limit, offset)
		if err != nil {
			log.Fatal(err)
		}
	}
	defer rows.Close()

	categories := make([]GetCategoryResponse, 0)
	for rows.Next() {
		category := new(GetCategoryResponse)
		err := rows.Scan(&category.Id, &category.Title)
		if err != nil {
			log.Fatal(err)
		}
		categories = append(categories, *category)
	}
	writeAsJSON(w, categories)
}

func getCategory(w http.ResponseWriter, r *http.Request) {
	ids := r.URL.Query()["id"]
	if len(ids) < 1 {
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
	category := &GetCategoryResponse{Id: idInt}
	query := "select title from category where id = $1"
	row := DB.QueryRow(query, idStr)
	err = row.Scan(&category.Title)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println(err)
			return
		}
		log.Fatal(err)
	}
	writeAsJSON(w, category)
}

func putCategory(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		log.Fatal(err)
	}
	id := r.FormValue("id")
	title := r.FormValue("title")
	query := "update category set title = $1 where id = $2"
	_, err = DB.Exec(query, title, id)
	if err != nil {
		log.Fatal(err)
	}
	writeAsJSON(w, &ErrorResponse{Ok: true, Error: ""})
}

func deleteCategory(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		log.Fatal(err)
	}
	id := r.FormValue("id")
	query := "delete from category where id=$1"
	_, err = DB.Exec(query, id)
	if err != nil {
		log.Fatal(err)
	}
	writeAsJSON(w, &ErrorResponse{Ok: true, Error: ""})
}
