package keyspot

import (
	"net/http"
	"json"
)

url := "https://database-driver-ifhogzjzbq-uc.a.run.app"

func Keyspot(accessKey string) map[string]interface, err {
	resp, err := http.Get(url + "/" + accessKey)
}