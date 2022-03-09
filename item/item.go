package item

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/rubens-schmitz/shop/util"
)

type PostItemRequest struct {
	Id        int `json:"id"`
	ProductId int `json:"productId"`
	CartId    int `json:"cartId"`
	Quantity  int `json:"quantity"`
}

func PostItem(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		log.Fatal(err)
	}
	productId, err := strconv.Atoi(r.FormValue("productId"))
	if err != nil {
		log.Fatal(err)
	}
	cartId := util.GetCartId(w, r)
	item := &PostItemRequest{ProductId: productId, CartId: cartId}
	query := `select id, quantity from item where productId = $1 and cartId = $2`
	row := util.DB.QueryRow(query, productId, cartId)
	err = row.Scan(&item.Id, &item.Quantity)
	if err != nil {
		if err == sql.ErrNoRows {
			query = `insert into item (productId, cartId) values ($1, $2)`
			_, err = util.DB.Exec(query, productId, cartId)
			if err != nil {
				log.Println(err)
				util.WriteAsJSON(w, &util.ErrorResponse{Ok: false, Error: err.Error()})
				return
			}
		} else {
			log.Fatal(err)
		}
	}
	item.Quantity += 1
	query = `update item set quantity = $2 where id = $1`
	_, err = util.DB.Exec(query, item.Id, item.Quantity)
	if err != nil {
		log.Fatal(err)
	}
	util.WriteAsJSON(w, &util.ErrorResponse{Ok: true, Error: ""})
}

func PutItem(w http.ResponseWriter, r *http.Request) {
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
		util.WriteAsJSON(w, &util.ErrorResponse{
			Ok: false, Error: "Parameter 'quantity' is less than or equal zero."})
		return
	}
	query := `update item set quantity = $1 where id = $2`
	_, err = util.DB.Exec(query, quantity, id)
	if err != nil {
		log.Fatal(err)
	}
	util.WriteAsJSON(w, &util.ErrorResponse{Ok: true, Error: ""})
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		log.Fatal(err)
	}
	id := r.FormValue("id")
	query := `delete from item where id = $1`
	_, err = util.DB.Exec(query, id)
	if err != nil {
		log.Fatal(err)
	}
	util.WriteAsJSON(w, &util.ErrorResponse{Ok: true, Error: ""})
}
