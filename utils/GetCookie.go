package utils

import (
	"log"
	"net/http"
	"strings"
	sgcon "xsgsd/conf"
)

func GetCookie() {
	req, err := http.NewRequest("GET", sgcon.LoginURL, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36")
	resp, err := sgcon.Client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	getcookie := resp.Header.Get("set-cookie")
	sgcon.Cookie = strings.Replace(getcookie, "; path=/", "", -1)
}
