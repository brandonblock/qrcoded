package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fmt.Println("Hello QR Code")

	qrcode := GenerateQRCode("555-2368") //"wishful thinking" function
	ioutil.WriteFile("qrcode.png", qrcode, 0644)
}

func GenerateQRCode(code string) []byte {
	return nil
}
