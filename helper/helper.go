package helper

import (
	"io"
	"log"
	"net/http"
	"strings"
)

func URLParam(reqpath string, url string) bool {
	urls := strings.Split(url, ":")
	prefix := reqpath[:strings.LastIndex(reqpath, "/")+1]
	return prefix == urls[0]
}

func GetParam(r *http.Request) string {
	return r.URL.Path[strings.LastIndex(r.URL.Path, "/")+1:]
}

func GetIPaddress() string {

	resp, err := http.Get("https://icanhazip.com/")

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}
	return string(body)
}
