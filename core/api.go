package core

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
)

type Api struct {
	Config *viper.Viper
}

func NewApi(config *viper.Viper) Api {
	return Api{
		config,
	}
}

func (t Api) GetCaiHongPi() (string, error) {
	var (
		err     error
		url     string
		rspBody []byte
	)
	key := t.Config.GetString("TianXingKey")
	url = fmt.Sprintf("http://api.tianapi.com/caihongpi/index?key=%s", key)
	rspBody, err = t.Curl(url)
	var result CaiHongPi
	if err = json.Unmarshal(rspBody, &result); err != nil {
		return "", err
	}
	if len(result.NewsList) > 0 {
		return result.NewsList[0].Content, nil
	}
	return "", nil
}

func (t Api) TianQi() (result TianQi, err error) {
	var rspBody []byte
	appId := t.Config.GetString("tqAppId")
	appSecret := t.Config.GetString("tqAppSecret")
	city := t.Config.GetString("city")
	url := fmt.Sprintf("https://v0.yiketianqi.com/api?unescape=1&version=v61&appid=%s&appsecret=%s&city=%s", appId, appSecret, city)
	rspBody, err = t.Curl(url)
	if err = json.Unmarshal(rspBody, &result); err != nil {
		return
	}
	return
}

func (t Api) Curl(url string) (rspBody []byte, err error) {
	var (
		httpReq *http.Request
		httpRsp *http.Response
	)
	httpReq, err = http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	httpRsp, err = http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRsp.Body.Close()
	rspBody, err = ioutil.ReadAll(httpRsp.Body)
	if err != nil {
		return nil, err
	}
	return rspBody, nil
}
