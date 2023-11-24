package service

import (
	"cciicc/types"

	qrcode "github.com/skip2/go-qrcode"
)

func GenerateQRPNG(space_id string) {
	var url string
	if string(types.URL_ADDESS[len(types.URL_ADDESS)-1]) != "/" { //상수 URL_ADDRESS가 '/'로 끝나지 아니할 때
		url = types.URL_ADDESS + "/"
	} else { // '/'로 끝날 떄
		url = types.URL_ADDESS
	}
	url = url + space_id
	filepath := "wwwfiles/assets/space_qr/" + space_id + ".png"
	QR_err := qrcode.WriteFile(url, qrcode.Highest, 4096, filepath)
	ErrHandler(QR_err, "QR_err")

}
