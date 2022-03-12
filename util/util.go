package util

import (
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"image/png"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"

	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"
)

type ErrorResponse struct {
	Ok    bool   `json:"id"`
	Error string `json:"error"`
}

var DB *sql.DB

func WriteAsJSON(w http.ResponseWriter, v any) {
	data, err := json.Marshal(v)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(data)
	if err != nil {
		log.Fatal(err)
	}
}

func ShortDatestamp(datestamp string) string {
	r, err := regexp.Compile("([0-9]+-[0-9]+-[0-9]+ [0-9]+:[0-9]+:[0-9]+)")
	if err != nil {
		log.Fatal(err)
	}
	return r.FindString(datestamp)
}

func GetPictures(productId int) []string {
	pictures := make([]string, 0)
	query := "select id, bytes from picture where productId = $1"
	rows, err := DB.Query(query, productId)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var id int
		var bytes []byte
		err := rows.Scan(&id, &bytes)
		if err != nil {
			log.Fatal(err)
		}
		pictures = append(pictures, string(bytes))
	}
	return pictures
}

func GetIntParam(r *http.Request, name string) (int, error) {
	arr := r.URL.Query()[name]
	val := 0
	var err error
	if len(arr) != 0 {
		val, err = strconv.Atoi(arr[0])
		if err != nil {
			log.Fatal(err)
		}
		if val < 0 {
			s := fmt.Sprintf("Parameter '%v' is less than zero.", name)
			return 0, errors.New(s)
		}
	}
	return val, nil
}

func GetStringParam(r *http.Request, name string) string {
	arr := r.URL.Query()[""]
	val := ""
	if len(arr) != 0 {
		val = arr[0]
	}
	return val
}

func EncodeQRCode(code string) string {
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

// func DecodeQRCode(qrcode string) {
// 	// open and decode image file
// 	file, _ := os.Open("qrcode.png")
// 	img, _, _ := image.Decode(file)

// 	// prepare BinaryBitmap
// 	bmp, _ := gozxing.NewBinaryBitmapFromImage(img)

// 	// decode image
// 	qrReader := qrcode.NewQRCodeReader()
// 	result, _ := qrReader.Decode(bmp, nil)

// 	log.Println(result)
// }
