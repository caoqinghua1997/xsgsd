package modules

import (
	"fmt"
	sgcon "xsgsd/conf"
	"xsgsd/utils"
)

type wjInfo struct {
	EventID    int    `json:"eventID"`
	ServerType string `json:"serverType"`
	JunshiID   int    `json:"junshiID"`
	IsOpen     int    `json:"isOpen"`
	Block      struct {
		Wj []struct {
			FightValue string      `json:"fightValue"`
			FightState interface{} `json:"fightState"`
			ArmID      int         `json:"armID"`
			Loyalty    int         `json:"loyalty"`
			Arm        string      `json:"arm"`
			Name       string      `json:"name"`
			ID         string      `json:"id"`
			State      string      `json:"state"`
			Status     string      `json:"status"`
		} `json:"wj"`
	} `json:"block"`
}

func GetWjInfo() {
	var data = fmt.Sprintf(`serverType=junshi&eventID=1&token=%s&playerID=%s`, sgcon.LoginToken, sgcon.PlayerID)
	wjinfo := utils.Server(data)
	fmt.Println(wjinfo)

}
