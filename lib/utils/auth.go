package utils

import (
	"net"

	"github.com/satori/go.uuid"
	qrcode "github.com/skip2/go-qrcode"
	"github.com/vence722/convert"
)

func GenerateQRCodeAndAuthToken(port int) ([]byte, string, error) {
	authToken := uuid.Must(uuid.NewV4()).String()
	key := "http://" + getLocalIP() + ":" + convert.Int2Str(port) + "/control?token=" + authToken
	qrCodeData, err := qrcode.Encode(key, qrcode.Medium, 256)
	if err != nil {
		return nil, "", err
	}
	return qrCodeData, authToken, nil
}

func getLocalIP() string {
	var ip net.IP
	ifaces, _ := net.Interfaces()
	for _, i := range ifaces {
		addrs, _ := i.Addrs()
		for _, addr := range addrs {
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip.IsGlobalUnicast() {
				return ip.String()
			}
		}
	}
	return ip.String()
}
