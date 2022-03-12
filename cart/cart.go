package cart

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/rubens-schmitz/shop/types"
	"github.com/rubens-schmitz/shop/util"
)

func GetCartHandler(w http.ResponseWriter, r *http.Request) {
	id := GetCartId(w, r)
	query := `select price, quantity from cart where id = $1`
	row := util.DB.QueryRow(query, id)
	cart := new(types.GetCartResponse)
	err := row.Scan(&cart.Price, &cart.Quantity)
	if err != nil {
		log.Fatal(err)
	}
	util.WriteAsJSON(w, cart)
}

func AddNewCartIdCookie(w http.ResponseWriter, r *http.Request) {
	datestamp := time.Now()
	query := `insert into cart (price, quantity, datestamp)
			  values ($1, $2, $3) returning id`
	row := util.DB.QueryRow(query, 0, 0, datestamp)
	var id string
	err := row.Scan(&id)
	if err != nil {
		log.Fatal(err)
	}
	cookie := &http.Cookie{Name: "cartId", Value: id, Path: "/"}
	http.SetCookie(w, cookie)
	r.AddCookie(cookie)
}

func GetCartId(w http.ResponseWriter, r *http.Request) int {
	cookie, err := r.Cookie("cartId")
	if err != nil {
		AddNewCartIdCookie(w, r)
		cookie, err = r.Cookie("cartId")
		if err != nil {
			log.Fatal(err)
		}
	}
	id, err := strconv.Atoi(cookie.Value)
	if err != nil {
		log.Fatal(err)
	}
	return id
}

func Update(cartId int) {
	query := `select product.price,	item.quantity from product inner join item 
			  on product.id = item.productId where item.cartId = $1`
	rows, err := util.DB.Query(query, cartId)
	if err != nil {
		log.Fatal(err)
	}
	var totalPrice float32 = 0
	var totalQuantity int = 0
	for rows.Next() {
		var price float32
		var quantity int
		err := rows.Scan(&price, &quantity)
		if err != nil {
			log.Fatal(err)
		}
		totalPrice += price * float32(quantity)
		totalQuantity += quantity
	}
	datestamp := util.ShortDatestamp(time.Now().String())
	query = `update cart set price = $1, quantity = $2,
			 datestamp = $3 where id = $4`
	_, err = util.DB.Exec(query, totalPrice, totalQuantity, datestamp, cartId)
	if err != nil {
		log.Fatal(err)
	}
}
