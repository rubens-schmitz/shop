package deal

import (
	"log"
	"net/http"
	"regexp"
	"time"

	"github.com/rubens-schmitz/shop/item"
	"github.com/rubens-schmitz/shop/util"
	"github.com/sethvargo/go-password/password"
)

type PostDealResponse struct {
	Qrcode string `json:"qrcode"`
}

type GetDealResponse struct {
	Id        int     `json:"id"`
	Quantity  int     `json:"quantity"`
	Price     float32 `json:"price"`
	Datestamp string  `json:"datestamp"`
	CartId    int     `json:"cartId"`
}

func shortDatestamp(datestamp string) string {
	r, err := regexp.Compile("([0-9]+-[0-9]+-[0-9]+ [0-9]+:[0-9]+:[0-9]+)")
	if err != nil {
		log.Fatal(err)
	}
	return r.FindString(datestamp)
}

func computePriceAndQuantity(items []item.GetItemResponse) (float32, int) {
	var price float32 = 0.0
	var quantity int = 0
	for i := 0; i < len(items); i++ {
		price += items[i].Price * float32(items[i].Quantity)
		quantity += items[i].Quantity
	}
	return price, quantity
}

func PostDealHandler(w http.ResponseWriter, r *http.Request) {
	code, err := password.Generate(64, 10, 10, false, false)
	if err != nil {
		log.Fatal(err)
	}
	datestamp := time.Now()
	cartId := util.GetCartId(w, r)
	query := `insert into deal (code, datestamp, cartId) values ($1, $2, $3)`
	_, err = util.DB.Exec(query, code, datestamp, cartId)
	if err != nil {
		log.Fatal(err)
	}
	qrcode := util.EncodeQRCode(code)
	res := PostDealResponse{Qrcode: qrcode}
	util.AddNewCartIdCookie(w, r)
	util.WriteAsJSON(w, res)
}

func GetDealHandler(w http.ResponseWriter, r *http.Request) {
	id, err := util.GetIntParam(r, "id")
	if err != nil {
		log.Fatal(err)
	}
	query := `select datestamp, cartId from deal where id = $1`
	row := util.DB.QueryRow(query, id)
	if err != nil {
		log.Fatal(err)
	}
	deal := GetDealResponse{Id: id}
	var datestamp string
	row.Scan(&datestamp, &deal.CartId)
	if err != nil {
		log.Fatal(err)
	}
	deal.Datestamp = shortDatestamp(datestamp)
	params := item.GetItemsParams{Title: "", Offset: 0, Limit: 16,
		CategoryId: 0, CartId: deal.CartId}
	items := item.GetItems(params)
	deal.Price, deal.Quantity = computePriceAndQuantity(items)
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
