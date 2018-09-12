package utils

import (
	"github.com/satori/go.uuid"
	qrcode "github.com/skip2/go-qrcode"
)

func GenerateAuthTokenAndQRCodePair() (string, []byte, error) {
	key := uuid.Must(uuid.NewV4()).String()
	qrCodeData, err := qrcode.Encode(key, qrcode.Medium, 256)
	if err != nil {
		return "", nil, err
	}
	return key, qrCodeData, nil
}
