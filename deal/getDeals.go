package deal

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/rubens-schmitz/shop/util"
)

type GetDealResponse struct {
	Id        int    `json:"id"`
	Code      string `json:"code"`
	Datestamp string `json:"datestamp"`
	CartId    int    `json:"cartId"`
}

func getRows() *sql.Rows {
	var query string
	var rows *sql.Rows
	var err error
	query = `select id, code, datestamp, cartId from deal`
	rows, err = util.DB.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	return rows
}

func reallyGetDeals(rows *sql.Rows) []GetDealResponse {
	deals := make([]GetDealResponse, 0)
	for rows.Next() {
		deal := new(GetDealResponse)
		err := rows.Scan(&deal.Id, &deal.Code, &deal.Datestamp, &deal.CartId)
		if err != nil {
			log.Fatal(err)
		}
		deals = append(deals, *deal)
	}
	return deals
}

func GetDeals(w http.ResponseWriter, r *http.Request) {
	rows := getRows()
	defer rows.Close()
	deals := reallyGetDeals(rows)
	util.WriteAsJSON(w, deals)
}
