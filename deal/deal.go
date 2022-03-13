package deal

import (
	"log"
	"net/http"

	"github.com/rubens-schmitz/shop/access"
	"github.com/rubens-schmitz/shop/cart"
	"github.com/rubens-schmitz/shop/types"
	"github.com/rubens-schmitz/shop/util"
)

func PostDealHandler(w http.ResponseWriter, r *http.Request) {
	accessId := access.CreateAccess("deal")
	cartId := cart.GetCartId(w, r)
	query := `insert into deal (accessId, cartId) values ($1, $2)`
	_, err := util.DB.Exec(query, accessId, cartId)
	if err != nil {
		log.Fatal(err)
	}
	qrcode := util.EncodeQRCode(access.GetCode(accessId))
	res := types.PostDealResponse{Qrcode: qrcode}
	cart.AddNewCartIdCookie(w, r)
	util.WriteAsJSON(w, res)
}

func GetDealHandler(w http.ResponseWriter, r *http.Request) {
	id, err := util.GetIntParam(r, "id")
	if err != nil {
		log.Fatal(err)
	}
	query := `select cart.price, cart.quantity, cart.datestamp, deal.cartId
			  from deal inner join cart on cart.id = deal.cartId
			  where deal.id = $1`
	row := util.DB.QueryRow(query, id)
	if err != nil {
		log.Fatal(err)
	}
	deal := types.GetDealResponse{Id: id}
	var datestamp string
	row.Scan(&deal.Price, &deal.Quantity, &datestamp, &deal.CartId)
	deal.Datestamp = util.ShortDatestamp(datestamp)
	if err != nil {
		log.Fatal(err)
	}
	util.WriteAsJSON(w, deal)
}

func DeleteDealHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		log.Fatal(err)
	}
	id := r.FormValue("id")
	query := "delete from deal where id = $1"
	_, err = util.DB.Exec(query, id)
	if err != nil {
		log.Fatal(err)
	}
	util.WriteAsJSON(w, util.ErrorResponse{Ok: true, Error: ""})
}
