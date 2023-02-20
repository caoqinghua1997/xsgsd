package main

import (
	"fmt"
	"xsgsd/modules"
	"xsgsd/utils"
)

func main() {

	fmt.Println("获取帐号信息中，请稍后...")
	// login

	utils.Login()
	//// servers
	//fmt.Println("get game server")
	utils.GameServer()
	utils.GetCookie()
	//
	utils.GetPlayerID()
	// get info
	modules.GetWjInfo()
}
