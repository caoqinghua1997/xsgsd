package utils

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"regexp"
	"strings"
	sgcon "xsgsd/conf"
)

func GameServer() {
	//fmt.Println("get  Server list 。。。。。。。。。。。。。。。。。。。。")
	//data := fmt.Sprintf("channel=xmw&time=&access_token=%s&username=%s&uuid=%s&isapp=", sgcon.LoginToken, sgcon.UserName, sgcon.LoginUUID)
	//fmt.Println(data)

	data := fmt.Sprintf("http://xsgsd.xmwan.com/server.php?channel=xmw&time=&access_token=%s&username=%s&uuid=%s&isapp=", sgcon.LoginToken, sgcon.UserName, sgcon.LoginUUID)
	sgcon.APIServer = "http://xsgsd.xmwan.com/server.php"

	ret := GetServer(data)
	//fmt.Println(ret)
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(ret))
	if err != nil {
		log.Fatalln(err)
	}
	dom.Find(".serverboxtit a").Each(func(i int, selection *goquery.Selection) {

		url, _ := selection.Attr("href")
		sgcon.StartUrl = url
		//fmt.Println("获取到的url: ", url)
		r := regexp.MustCompile(`(http://.*?)/sanguo2.jsp.*?userName=(.*?)&pwd=(.*?)&.*?uuid=(.*?)&.*?sid=(.*?)&.*?sname=(.*?)&.*?`)
		matchs := r.FindAllStringSubmatch(url, -1)
		sgcon.LoginURL = matchs[0][1]
		sgcon.APIServer = sgcon.LoginURL + "/http.do"

		sgcon.LoginSid = matchs[0][5]
		sgcon.LoginSname = matchs[0][6]
		data := fmt.Sprintf("%s/sanguo2.jsp?userName=%s&pwd=%s&channel=xmw&entertype=2&uuid=%s&sid=%s",
			sgcon.LoginURL, sgcon.UserName, sgcon.LoginToken, sgcon.LoginUUID, sgcon.LoginSid)
		sgcon.LoginURL = data
	})

}
