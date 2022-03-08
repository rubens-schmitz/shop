package deal

import (
	"encoding/base64"
	"image/png"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"
	"github.com/rubens-schmitz/shop/util"
	"github.com/sethvargo/go-password/password"
)

type PostDealResponse struct {
	Qrcode string `json:"qrcode"`
}

func makeQRCode(code string) string {
	enc := qrcode.NewQRCodeWriter()
	img, err := enc.Encode(code, gozxing.BarcodeFormat_QR_CODE, 256, 256, nil)
	if err != nil {
		log.Fatal(err)
	}
	file, err := os.CreateTemp("", "")
	defer file.Close()
	err = png.Encode(file, img)
	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadFile(file.Name())
	if err != nil {
		log.Fatal(err)
	}
	qrcode := "data:image/png;base64,"
	qrcode += base64.StdEncoding.EncodeToString(data)
	return qrcode
}

func PostDeal(w http.ResponseWriter, r *http.Request) {
	code, err := password.Generate(64, 10, 10, false, false)
	if err != nil {
		log.Fatal(err)
	}
	datestamp := time.Now().String()
	cartId := util.GetCartId(w, r)

	query := `insert into deal (code, datestamp, cartId) values ($1, $2, $3)`
	_, err = util.DB.Exec(query, code, datestamp, cartId)
	if err != nil {
		log.Fatal(err)
	}

	qrcode := makeQRCode(code)
	res := &PostDealResponse{Qrcode: qrcode}
	util.AddNewCartIdCookie(w, r)
	util.WriteAsJSON(w, res)
}
