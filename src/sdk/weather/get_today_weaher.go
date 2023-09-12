package weather

import (
	"WeChatBot/src/utils/http"
	"WeChatBot/src/utils/log"
	"encoding/json"
	"fmt"
)

var Appid = "27751869"
var AppSecret = "07fhbBb2"
var Url = "https://v0.yiketianqi.com/free/day"

// GetTodayWeather 获取当日天气
type Request struct {
	Appid     string `json:"appid"`
	Appsecret string `json:"appsecret"`
	Cityid    string `json:"cityid"`
	City      string `json:"city"`
	Ip        string `json:"ip"`
	Callback  string `json:"callback"`
	Vue       string `json:"vue"`
	Unescape  int    `json:"unescape"`
}
type Response struct {
	Nums       int    `json:"nums"`
	Cityid     string `json:"cityid"`
	City       string `json:"city"`
	UpdateTime string `json:"update_time"`
	Wea        string `json:"wea"`
	WeaImg     string `json:"wea_img"`
	Tem        string `json:"tem"`
	TemDay     string `json:"tem_day"`
	TemNight   string `json:"tem_night"`
	Win        string `json:"win"`
	WinSpeed   string `json:"win_speed"`
	WinMeter   string `json:"win_meter"`
	Air        string `json:"air"`
	Humidity   string `json:"humidity"`
}

func (res *Response) ToString() string {
	return fmt.Sprintf("%s %s %s°-%s° %s%s 空气质量:%s 湿度:%s", res.City, res.Wea, res.TemNight, res.TemDay, res.Win, res.WinSpeed, res.Air, res.Humidity)
}

func GetTodayWeather() (resp *Response) {
	req := &Request{}
	reqStr, err := json.Marshal(req)
	if err != nil {
		log.ERROR("GetTodayWeather marshal %s", err.Error())
		return
	}

	res, err := http.GetWithHeaderAndBody(Url, reqStr, nil, map[string]string{
		"appid":     Appid,
		"appsecret": AppSecret,
		"cityid":    "CN101021500",
		"city":      "",
		"ip":        "",
		"callback":  "",
		"vue":       "",
		"unescape":  "1",
	})
	if err != nil {
		log.ERROR("GetTodayWeather %s", err.Error())
		return
	}
	log.INFOF("GetTodayWeather result=%v", res)
	resp = &Response{}
	err = json.Unmarshal([]byte(res), resp)
	if err != nil {
		log.ERRORF("GetTodayWeather %v", err)
		return
	}
	return
}
