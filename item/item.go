package item

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/rubens-schmitz/shop/cart"
	"github.com/rubens-schmitz/shop/types"
	"github.com/rubens-schmitz/shop/util"
)

func PostItemHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		log.Fatal(err)
	}
	productId, err := strconv.Atoi(r.FormValue("productId"))
	if err != nil {
		log.Fatal(err)
	}
	cartId := cart.GetCartId(w, r)
	item := types.PostItemRequest{ProductId: productId, CartId: cartId}
	query := `select id, quantity from item where productId = $1 and cartId = $2`
	row := util.DB.QueryRow(query, productId, cartId)
	err = row.Scan(&item.Id, &item.Quantity)
	if err != nil {
		if err == sql.ErrNoRows {
			query = `insert into item (productId, cartId) values ($1, $2)`
			_, err = util.DB.Exec(query, productId, cartId)
			if err != nil {
				log.Println(err)
				util.WriteAsJSON(w, types.SuccessResponse{Success: false,
					Msg: err.Error()})
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
	cart.Update(cartId)
	util.WriteAsJSON(w, types.SuccessResponse{Success: true,
		Msg: "Item added"})
}

func PutItemHandler(w http.ResponseWriter, r *http.Request) {
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
		util.WriteAsJSON(w, types.SuccessResponse{Success: false,
			Msg: "Parameter 'quantity' is less than or equal zero."})
		return
	}
	query := `update item set quantity = $1 where id = $2`
	_, err = util.DB.Exec(query, quantity, id)
	if err != nil {
		log.Fatal(err)
	}
	cart.Update(cart.GetCartId(w, r))
	util.WriteAsJSON(w, types.SuccessResponse{Success: true,
		Msg: "Item changed"})
}

func DeleteItemHandler(w http.ResponseWriter, r *http.Request) {
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
	cart.Update(cart.GetCartId(w, r))
	util.WriteAsJSON(w, types.SuccessResponse{Success: true,
		Msg: "Item deleted"})
}
