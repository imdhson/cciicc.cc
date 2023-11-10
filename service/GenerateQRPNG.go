package service

import qrcode "github.com/skip2/go-qrcode"

func GenerateQRPNG(space_id string) {
	url := "http://localhost/guest/" + space_id
	filepath := "wwwfiles/assets/space_qr/" + space_id + ".png"
	QR_err := qrcode.WriteFile(url, qrcode.Highest, 4096, filepath)
	ErrHandler(QR_err, "QR_err")

}
