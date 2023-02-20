package modules

import (
	"encoding/json"
	"fmt"
	"time"
	sgcon "xsgsd/conf"
	"xsgsd/utils"
)

type vip struct {
	Block struct {
		VipLevel int `json:"vipLevel"`
	} `json:"block"`
}

func GetDaysGift() {
	// login gift
	var loginGift = fmt.Sprintf("serverType=jinnang&eventID=2&token=%s&playerID=%s", sgcon.LoginToken, sgcon.PlayerID)
	for i := 0; i < 2; i++ {
		utils.Server(loginGift)
		time.Sleep(time.Second * 1)
		fmt.Println("每日登录奖励领取完成...")
	}

	// vip
	var isvip = fmt.Sprintf("serverType=vip&eventID=1&token=%s&playerID=%s", sgcon.LoginToken, sgcon.PlayerID)
	ret := utils.Server(isvip)
	m := vip{}
	err := json.Unmarshal([]byte(ret), &m)
	if err != nil {
		fmt.Println("Umarshal failed:", err)
		return
	}
	if m.Block.VipLevel >= 1 {
		var vipData = fmt.Sprintf("serverType=vip&eventID=3&token=%s&playerID=%s", sgcon.Password, sgcon.PlayerID)
		utils.Server(vipData)
		fmt.Println("VIP奖励领取完成")
	} else {
		fmt.Println("不是VIP哦")
	}

	// 好友互访
	friendID := 1
	var data = fmt.Sprintf("serverType=hufang&eventID=1&haoyouID=%s&token=%s&playerID=%s", friendID, sgcon.LoginToken, sgcon.PlayerID)
	utils.Server(data)
	for i := 0; i < 3; i++ {
		var data = fmt.Sprintf("serverType=hufang&eventID=5&hufangID=%s&token=%s&playerID=%s", i, sgcon.Password, sgcon.PlayerID)
		utils.Server(data)
	}
	fmt.Println("好友互访任务完成")

	// 赠送好友礼物

}
