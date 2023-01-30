package qr

import (
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"gopartsrv/public/consts"
	"image/png"
	"os"
)

//根据url生成二维码地址
func CreateQr(url string) string {
	picName := consts.Uuid() + ".png"
	qrCode, _ := qr.Encode(url, qr.M, qr.Auto)
	qrCode, _ = barcode.Scale(qrCode, 256, 256)
	file, _ := os.Create(picName)
	defer file.Close()
	png.Encode(file, qrCode)
	return picName
}
