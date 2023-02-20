package conf

import (
	"github.com/spf13/viper"
	"log"
	"net/http"
	"net/http/cookiejar"
)

var (
	UserName       string
	LoginTokenName string
	LoginTime      string
	LoginToken     string
	LoginUUID      string
	LoginURL       string
	LoginSid       string
	LoginSname     string
	Password       string
	StartUrl       string
	UserUUID       string
	APIServer      string
	PlayerID       string
	Cookie         string
	UserVIP        string
	GifType        string
	Client         *http.Client
	Wj             map[string]map[string]string
	Friends        = make(map[string]string, 10)
)

func init() {

	viper.SetConfigName("conf.yaml")
	viper.SetConfigFile("conf/conf.yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	// set client
	jar, err := cookiejar.New(nil)
	if err != nil {
		panic(err)
	}
	Client = &http.Client{
		Jar: jar,
	}

	//
	parseLoginConfInfo()
}

func parseLoginConfInfo() {
	UserName = viper.GetString("login.username")
	Password = viper.GetString("login.password")
}
