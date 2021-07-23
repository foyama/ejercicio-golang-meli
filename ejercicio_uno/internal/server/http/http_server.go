package http

import (
	"io/ioutil"
	"log"
	"net/http"
)

func DoGet(url string) (body []byte, err error) {
	resp, err := http.Get("https://api.coingecko.com/api/v3/coins/bitcoin")
	if err != nil {
		log.Fatalln(err)
	}
	return ioutil.ReadAll(resp.Body)
}
