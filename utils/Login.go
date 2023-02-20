package utils

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"strings"
	sgcon "xsgsd/conf"
)

func Login() {
	sgcon.APIServer = "http://sg.xmwan.com/xmw/dologin/from/mobile"
	login := fmt.Sprintf("username=%s&password=%s&channel=xmw&geetest_challenge2=&geetest_validate2=&geetest_seccode2=", sgcon.UserName, sgcon.Password)
	ret := Server(login)
	// get tokenName
	/*
		<form style='display:none;' id='form1' name='form1' method='post' action='http://xsgsd.xmwan.com/indexauth.php'>
		              <input name='channel' type='text' value='xmw' />
		              <input name='time' type='text' value='1676621738' />
		              <input name='tokenName' type='text' value='3d5391ef63bea27b399a1cde208dc406230cabb3'/>
		            </form>
		            <script type='text/javascript'>function load_submit(){document.form1.submit()}load_submit();</script>
	*/
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(ret))
	if err != nil {
		log.Fatalln(err)
	}

	// get tokenName
	dom.Find("input[name=tokenName]").Each(func(i int, selection *goquery.Selection) {
		token, _ := selection.Attr("value")
		sgcon.LoginTokenName = token
	})

	// get time
	dom.Find("input[name=time]").Each(func(i int, selection *goquery.Selection) {
		logintime, _ := selection.Attr("value")
		sgcon.LoginTime = logintime
	})

	// get
	sgcon.APIServer = "http://xsgsd.xmwan.com/indexauth.php"
	login = fmt.Sprintf("channel=xmw&time=%s&tokenName=%s", sgcon.LoginTime, sgcon.LoginTokenName)
	ret = Server(login)
	dom, err = goquery.NewDocumentFromReader(strings.NewReader(ret))
	if err != nil {
		log.Fatalln(err)
	}
	dom.Find("input[name=token]").Each(func(i int, selection *goquery.Selection) {
		token, _ := selection.Attr("value")
		sgcon.LoginToken = token
	})

	// get uuid and username
	login = fmt.Sprintf(`geetest_challenge2=&geetest_validate2=&geetest_seccode2=&token=%s`, sgcon.LoginToken)
	sgcon.APIServer = "http://sg.xmwan.com/xmw/dologin2/from/mobile/channel/xmw"
	ret = Server(login)
	dom, err = goquery.NewDocumentFromReader(strings.NewReader(ret))
	if err != nil {
		log.Fatalln(err)
	}
	dom.Find("input[name=username]").Each(func(i int, selection *goquery.Selection) {
		username, _ := selection.Attr("value")
		sgcon.UserName = username
	})
	dom.Find("input[name=uuid]").Each(func(i int, selection *goquery.Selection) {
		uuid, _ := selection.Attr("value")
		sgcon.LoginUUID = uuid
	})

	// print

}
