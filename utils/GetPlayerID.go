package utils

import (
	"fmt"
	"regexp"
	"strings"
	"time"
	sgcon "xsgsd/conf"
)

func GetPlayerID() {
	data := fmt.Sprintf("%s/zone%s/sanguo2.jsp?userName%s&pwd=%s&channel=xmw&entertype=2&uuid=%s&sid=%s&isapp=",
		sgcon.LoginURL, sgcon.LoginSid, sgcon.UserName, sgcon.LoginToken, sgcon.LoginUUID, sgcon.LoginSid)

	GetServer(data)
	time.Sleep(time.Second * 4)
	GetServer(data)
	time.Sleep(time.Second * 4)
	data = fmt.Sprintf("%s?userName=%s&serverType=build&eventID=18",
		sgcon.APIServer, sgcon.UserName)

	GetServer(data)
	time.Sleep(time.Second * 4)

	data = fmt.Sprintf("serverType=build&eventID=1&type=0&userName=%s&pwd=%s&device=&d=web&uuid=%s&entertype=2&channel=xmw&networkFashion=&area=&model=&manufacturer=&pDeivce=&imei=&imsi=&p=&v=3.2.0&androidid=&token=%s&playerID=0",
		sgcon.UserName, sgcon.LoginToken, sgcon.LoginUUID, sgcon.LoginToken)

	ret := Server(data)
	time.Sleep(time.Second * 8)
	r, _ := regexp.Compile(`"playerID":"([0-9])+"`)
	x1 := r.FindString(ret)
	x2 := strings.Replace(strings.Split(x1, ":")[1], `"`, "", -1)
	sgcon.PlayerID = x2

	// all
	fmt.Println("登录成功，开始执行任务...请等待")
}
