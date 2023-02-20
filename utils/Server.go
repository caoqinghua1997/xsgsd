package utils

import (
	"io"
	"log"
	"net/http"
	"strings"
	sgcon "xsgsd/conf"
)

func Server(formData string) string {
	var data = strings.NewReader(formData)
	req, err := http.NewRequest("POST", sgcon.APIServer, data)
	if err != nil {
		log.Fatal(err)
	}
	return httpClient(req)
}

func GetServer(formdata string) string {
	req, err := http.NewRequest("GET", formdata, nil)
	if err != nil {
		log.Fatal(err)
	}
	return httpClient(req)
}

func httpClient(req *http.Request) string {

	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.66 Safari/537.36")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Origin", sgcon.APIServer)
	req.Header.Set("Referer", sgcon.StartUrl)
	req.Header.Set("Cookie", sgcon.Cookie)
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")

	resp, err := sgcon.Client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return string(bodyText)
}
