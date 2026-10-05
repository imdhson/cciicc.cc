package service

import (
	"os"

	qrcode "github.com/skip2/go-qrcode"
)

func GenerateQRPNG(space_id string) {
	baseURL := os.Getenv("BASE_URL")

	var url string
	if baseURL[len(baseURL)-1] != '/' { //상수 BASE_URL가 '/'로 끝나지 아니할 때
		url = baseURL + "/"
	} else { // '/'로 끝날 떄
		url = baseURL
	}
	url = url + space_id
	filepath := "wwwfiles/assets/space_qr/" + space_id + ".png"
	QR_err := qrcode.WriteFile(url, qrcode.Highest, 4096, filepath)
	ErrHandler(QR_err, "QR_err")

}
