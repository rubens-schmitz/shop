package deal

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/rubens-schmitz/shop/access"
	"github.com/rubens-schmitz/shop/types"
	"github.com/rubens-schmitz/shop/util"
)

func parseParams(r *http.Request) (types.GetDealsParams, error) {
	limit, err := util.GetIntParam(r, "limit")
	if err != nil {
		return types.GetDealsParams{}, err
	}
	offset, err := util.GetIntParam(r, "offset")
	if err != nil {
		return types.GetDealsParams{}, err
	}
	params := types.GetDealsParams{Limit: limit, Offset: offset}
	return params, nil
}

func queryRows(params types.GetDealsParams) *sql.Rows {
	var query string
	var rows *sql.Rows
	var err error
	query = `select deal.id, cart.price, cart.quantity, cart.datestamp
			 from deal inner join cart on cart.id = deal.cartId
			 order by cart.datestamp desc limit $1 offset $2`
	rows, err = util.DB.Query(query, params.Limit, params.Offset)
	if err != nil {
		log.Fatal(err)
	}
	return rows
}

func makeDeals(rows *sql.Rows) []types.GetDealResponse {
	deals := make([]types.GetDealResponse, 0)
	for rows.Next() {
		deal := new(types.GetDealResponse)
		var datestamp string
		err := rows.Scan(&deal.Id, &deal.Price, &deal.Quantity, &datestamp)
		deal.Datestamp = util.ShortDatestamp(datestamp)
		if err != nil {
			log.Fatal(err)
		}
		deals = append(deals, *deal)
	}
	return deals
}

func GetDealsHandler(w http.ResponseWriter, r *http.Request) {
	if !access.IsAdmin(r) {
		return
	}
	params, err := parseParams(r)
	if err != nil {
		log.Fatal(err)
	}
	rows := queryRows(params)
	defer rows.Close()
	deals := makeDeals(rows)
	util.WriteAsJSON(w, deals)
}
