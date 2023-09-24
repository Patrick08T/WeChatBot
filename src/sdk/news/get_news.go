package news

import (
	"bot/src/utils/http"
	"bot/src/utils/log"
	"encoding/json"
	"fmt"
)

var url = "https://v2.api-m.com/api/weibohot"

type Data struct {
	Index int    `json:"index"`
	Title string `json:"title"`
	Hot   string `json:"hot"`
	Url   string `json:"url"`
}

func (data *Data) ToString() string {
	return fmt.Sprintf("%d %s", data.Index, data.Url)
}

type Response struct {
	Code int     `json:"code"`
	Msg  string  `json:"msg"`
	Data []*Data `json:"data"`
}

func (resp *Response) ToString() (result string) {
	length := len(resp.Data)
	if length == 0 {
		return
	}
	for i := 0; i < length-1; i++ {
		result += resp.Data[i].ToString() + "\n"
	}
	d := resp.Data[length-1]
	if d == nil {
		log.ERRORF("Response ToString fail.")
	}
	result += d.ToString()
	return
}

func GetNews() (resp *Response) {
	res, err := http.GetWithHeaderAndBody(url, nil, nil, nil)
	if err != nil {
		log.ERROR("GetNews %s", err.Error())
		return
	}
	log.INFOF("GetNews result=%v", res)
	resp = &Response{
		Data: []*Data{},
	}
	err = json.Unmarshal([]byte(res), resp)
	if err != nil {
		log.ERRORF("GetNews %v", err)
		return
	}
	log.INFOF("GetNews json:%v", resp)
	return
}
