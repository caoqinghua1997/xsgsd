package modules

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"math/rand"
	"strings"
	"time"
	sgcon "xsgsd/conf"
	"xsgsd/utils"
)

type friensInfo struct {
	Block struct {
		Friends []struct {
			CityID   string      `json:"cityID"`
			Union    string      `json:"union"`
			PlayerID string      `json:"playerID"`
			VipLevel int         `json:"vipLevel"`
			Zhujun   interface{} `json:"zhujun"`
			Fight    string      `json:"fight"`
			IsOnline int         `json:"isOnline"`
			Zuobiao  string      `json:"zuobiao"`
			UnionID  string      `json:"unionID"`
			Name     string      `json:"name"`
			Level    int         `json:"level"`
		} `json:"friends"`
	} `json:"block"`
}

// articleID "1010000000000005550"
type giftStruct struct {
	Block struct {
		SendCount       int `json:"sendCount"`
		SendMaxCount    int `json:"sendMaxCount"`
		IsChouJiang     int `json:"isChouJiang"`
		UseArticleCount int `json:"useArticleCount"`
		Daoju           struct {
			ArticleID string `json:"articleID"`
			HasDaoju  int    `json:"hasDaoju"`
		} `json:"daoju"`
		Friends []struct {
			Fid  string `json:"fid"`
			Name string `json:"name"`
		} `json:"friends"`
	} `json:"block"`
}

func GetFriendsID() {
	var data = fmt.Sprintf("serverType=haoyou&eventID=1&token=%s&playerID=%s", sgcon.LoginToken, sgcon.PlayerID)
	ret := utils.Server(data)
	m := friensInfo{}
	err := json.Unmarshal([]byte(ret), &m)
	if err != nil {
		fmt.Println("Umarshal failed:", err)
		return
	}
	for _, v := range m.Block.Friends {
		sgcon.Friends[v.Name] = v.PlayerID
	}
}

func SendGift() {

	// 判断是否还有小礼袋
	var data = fmt.Sprintf("serverType=liwu&eventID=1&token=%s&playerID=%s", sgcon.LoginToken, sgcon.PlayerID)
	ret := utils.Server(data)
	m := giftStruct{}
	err := json.Unmarshal([]byte(ret), &m)
	if err != nil {
		fmt.Println("Umarshal failed:", err)
		return
	}
	time.Sleep(time.Second * 1)
	if m.Block.Daoju.HasDaoju == 1 && m.Block.SendMaxCount == 5 {
		// 使用小礼袋
		var djData = fmt.Sprintf("serverType=useArticle&eventID=2&operateName=addSendCount&articleID=%s&token=%s&playerID=%s", m.Block.Daoju.ArticleID, sgcon.LoginToken, sgcon.PlayerID)
		ret := utils.Server(djData)
		if strings.Contains(ret, "使用小礼袋成功") {
			fmt.Println("小礼袋使用成功")
		} else {
			fmt.Println("小礼袋使用失败")
		}
	} else {
		fmt.Println("已经使用过了，不再重复使用小礼袋")
	}
	time.Sleep(time.Second * 1)
	ret = utils.Server(data)
	m = giftStruct{}
	err = json.Unmarshal([]byte(ret), &m)
	if err != nil {
		fmt.Println("Umarshal failed:", err)
		return
	}

	// 判断是否有固定的好友位
	giftFriends := viper.Get("gift.friends")
	var temp = make([]string, 0, 10)
	for _, v := range giftFriends.([]interface{}) {
		temp = append(temp, v.(string))
	}
	var gdFriendTemp string
	if len(temp) != 0 {
		for _, f := range temp {
			for _, v := range m.Block.Friends {
				if strings.Contains(v.Name, string(f)) {
					gdFriendTemp += v.Fid + ","
				}
			}
		}
		gdFriendTemp := gdFriendTemp[:len(gdFriendTemp)-1]
		fmt.Println("赠送的固定好友： ", gdFriendTemp)
		// 赠送
		sendGifToFriends(gdFriendTemp)
	} else {
		fmt.Println("没有固定好友位")
	}
	time.Sleep(time.Second * 1)
	// 重新请求
	ret = utils.Server(data)
	m = giftStruct{}
	err = json.Unmarshal([]byte(ret), &m)
	if err != nil {
		fmt.Println("Umarshal failed:", err)
		return
	}

	var randFriendsID string
	// 判断是否随机
	if sgcon.GifType == "random" {
		randnum := genRandomNumber(0, len(m.Block.Friends), m.Block.SendMaxCount-m.Block.SendCount)
		for _, v := range randnum {
			fmt.Printf("%s   -->   %s\n", m.Block.Friends[v].Name, m.Block.Friends[v].Fid)
			randFriendsID += m.Block.Friends[v].Fid + ","
		}
		randFriendsID = randFriendsID[:len(randFriendsID)-1]
		sendGifToFriends(randFriendsID)
		fmt.Println("随机好友礼物赠送完成...")
	}

	// 获取赠送奖励
	//data = fmt.Sprintf("serverType=liwu&eventID=4&token=%s&playerID=%s", sgcon.Password, sgcon.playerID)
	//utils.Server(data)

}

func sendGifToFriends(friendsID string) {
	var data = fmt.Sprintf("serverType=liwu&eventID=3&aid=102&fids=%s&token=%s&playerID=%s", friendsID, sgcon.LoginToken, sgcon.PlayerID)
	fmt.Println("好友的ID：", data)
	utils.Server(data)

}

// 生成count个[start,end)结束的不重复的随机数
func genRandomNumber(start int, end int, count int) []int {
	//范围检查
	if end < start || (end-start) < count {
		return nil
	}

	//存放结果的slice
	nums := make([]int, 0)
	//随机数生成器，加入时间戳保证每次生成的随机数不一样
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for len(nums) < count {
		//生成随机数
		num := r.Intn((end - start)) + start

		//查重
		exist := false
		for _, v := range nums {
			if v == num {
				exist = true
				break
			}
		}

		if !exist {
			nums = append(nums, num)
		}
	}

	return nums
}
