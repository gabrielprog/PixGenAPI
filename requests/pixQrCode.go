package requests

import (
	"encoding/json"

	"github.com/go-resty/resty/v2"

	"log"
	"os"
)

var URLBASE = os.Getenv("PUSHINPAY_URL")
var token = os.Getenv("PUSHINPAY_TOKEN")
var headers = map[string]string{
	"Content-Type":  "application/json",
	"Authorization": "Bearer " + token,
	"Accept":        "application/json",
}

type ValueInCents interface {
	int | int8 | int16 | int32 | int64
}

type IdTransaction interface {
	string
}
type ResponsePixQrCode struct {
	Qrode        string `json:"qr_code"`
	Base64QrCode string `json:"qr_code_base64"`
	Status       string `json:"status"`
	Message      string `json:"message"`
}

type ResponseStatusPixQrCode struct {
	responsePixQrCode ResponsePixQrCode
}

func GeneratePixQrCode[T ValueInCents](valueOfPix *T) *ResponsePixQrCode {
	if *valueOfPix < 50 {
		*valueOfPix = 50
	}

	client := resty.New()
	rawJson := map[string]interface{}{
		"value": *valueOfPix,
	}
	responseInJson := &ResponsePixQrCode{}
	response, err := client.R().EnableTrace().SetHeaders(headers).SetBody(rawJson).Post(URLBASE + "/api/pix/cashIn")
	if err != nil {
		return responseInJson
	}
	if err := json.Unmarshal(response.Body(), responseInJson); err != nil {
		log.Fatal("Error unmarshalling response: ", err)
	}
	return responseInJson
}

func IsApprovedPayment[T IdTransaction](id T) bool {
	client := resty.New()
	responseInJson := &ResponsePixQrCode{}
	response, err := client.R().EnableTrace().SetHeaders(headers).Get(URLBASE + "/api/transactions/" + string(id))
	if err != nil {
		return false
	}
	if err := json.Unmarshal(response.Body(), responseInJson); err != nil || responseInJson.Status != "paid" {
		return false
	}
	return true
}
