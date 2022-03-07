package category

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/rubens-schmitz/shop/util"
)

type GetCategoryResponse struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

func PostCategory(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		log.Fatal(err)
	}
	title := r.FormValue("title")
	query := "insert into category (title) values ($1)"
	_, err = util.DB.Exec(query, title)
	if err != nil {
		log.Fatal(err)
	}
	util.WriteAsJSON(w, &util.ErrorResponse{Ok: true, Error: ""})
}

func GetCategory(w http.ResponseWriter, r *http.Request) {
	id, err := util.GetIntParam(r, "id")
	if err != nil {
		log.Fatal(err)
	}
	category := &GetCategoryResponse{Id: id}
	query := "select title from category where id = $1"
	row := util.DB.QueryRow(query, id)
	err = row.Scan(&category.Title)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println(err)
			return
		}
		log.Fatal(err)
	}
	util.WriteAsJSON(w, category)
}

func PutCategory(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		log.Fatal(err)
	}
	id := r.FormValue("id")
	title := r.FormValue("title")
	query := "update category set title = $1 where id = $2"
	_, err = util.DB.Exec(query, title, id)
	if err != nil {
		log.Fatal(err)
	}
	util.WriteAsJSON(w, &util.ErrorResponse{Ok: true, Error: ""})
}

func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		log.Fatal(err)
	}
	id := r.FormValue("id")
	query := "delete from category where id=$1"
	_, err = util.DB.Exec(query, id)
	if err != nil {
		log.Fatal(err)
	}
	util.WriteAsJSON(w, &util.ErrorResponse{Ok: true, Error: ""})
}
