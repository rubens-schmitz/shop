package product

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/rubens-schmitz/shop/types"
	"github.com/rubens-schmitz/shop/util"
)

func PostProductHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		log.Fatal(err)
	}
	title := r.FormValue("title")
	price := r.FormValue("price")
	categoryId := r.FormValue("categoryId")
	query := `insert into product (title, price, categoryId)
	          values ($1, $2, $3) returning id`
	row := util.DB.QueryRow(query, title, price, categoryId)
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
		_, err = util.DB.Exec(query, id, bytes)
	}
	util.WriteAsJSON(w, util.ErrorResponse{Ok: true, Error: ""})
}

func extractPicture(r *http.Request, n int) []byte {
	name := fmt.Sprintf("pictures%v", n)
	picture := r.FormValue(name)
	return []byte(picture)
}

func GetProductHandler(w http.ResponseWriter, r *http.Request) {
	id, err := util.GetIntParam(r, "id")
	if err != nil {
		log.Fatal(err)
	}
	pictures := util.GetPictures(id)
	product := types.GetProductResponse{Id: id, Pictures: pictures}
	query := "select title, price, categoryId from product where id=$1"
	row := util.DB.QueryRow(query, id)
	err = row.Scan(&product.Title, &product.Price, &product.CategoryId)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println(err)
			return
		}
		log.Fatal(err)
	}
	util.WriteAsJSON(w, product)
}

func PutProductHandler(w http.ResponseWriter, r *http.Request) {
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
	_, err = util.DB.Exec(query, title, price, categoryId, id)
	query = "delete from picture where productId = $1"
	_, err = util.DB.Exec(query, id)
	pictures, err := strconv.Atoi(r.FormValue("pictures"))
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < pictures; i++ {
		bytes := extractPicture(r, i)
		query = "insert into picture (productId, bytes) values ($1, $2)"
		_, err = util.DB.Exec(query, id, bytes)
	}
	util.WriteAsJSON(w, util.ErrorResponse{Ok: true, Error: ""})
}

func DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		log.Fatal(err)
	}
	id := r.FormValue("id")
	query := `delete from product where id = $1`
	_, err = util.DB.Exec(query, id)
	if err != nil {
		log.Fatal(err)
	}
	util.WriteAsJSON(w, util.ErrorResponse{Ok: true, Error: ""})
}
