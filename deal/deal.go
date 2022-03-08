package deal

import (
	"image"
	"image/png"
	"log"
	"net/http"
	"os"

	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"
)

func GetDeal(w http.ResponseWriter, r *http.Request)    {}
func DeleteDeal(w http.ResponseWriter, r *http.Request) {}

func readQrcode() {
	// open and decode image file
	file, _ := os.Open("qrcode.png")
	img, _, _ := image.Decode(file)

	// prepare BinaryBitmap
	bmp, _ := gozxing.NewBinaryBitmapFromImage(img)

	// decode image
	qrReader := qrcode.NewQRCodeReader()
	result, _ := qrReader.Decode(bmp, nil)

	log.Println(result)
}

func writeQrcode() {
	// Generate a barcode image (*BitMatrix)
	enc := qrcode.NewQRCodeWriter()
	img, _ := enc.Encode("Hello, Gophers!",
		gozxing.BarcodeFormat_QR_CODE, 264, 264, nil)

	file, _ := os.Create("qrcode.png")
	defer file.Close()

	// *BitMatrix implements the image.Image interface,
	// so it is able to be passed to png.Encode directly.
	_ = png.Encode(file, img)
}
