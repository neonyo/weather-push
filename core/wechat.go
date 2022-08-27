package core

import (
	"fmt"
	"github.com/golang-module/carbon/v2"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	offConfig "github.com/silenceper/wechat/v2/officialaccount/config"
	"github.com/silenceper/wechat/v2/officialaccount/message"
	"github.com/spf13/viper"
)

type Wechat struct {
	Config *viper.Viper
}

func NewWeChat(config *viper.Viper) Wechat {
	return Wechat{
		config,
	}
}

func (w Wechat) Send() {
	var err error
	wc := wechat.NewWechat()

	memory := cache.NewMemory()
	cfg := &offConfig.Config{
		AppID:     w.Config.GetString("appId"),
		AppSecret: w.Config.GetString("appSecret"),
		Cache:     memory,
	}
	oa := wc.GetOfficialAccount(cfg)
	bd := oa.GetTemplate()
	var msgTpl *message.TemplateMessage
	msgTpl, err = w.GetTpl()
	_, err = bd.Send(msgTpl)
	if err != nil {
		fmt.Printf("Unmarshal fail, err:%v", err)
	}
}

func (w Wechat) GetTpl() (msgTpl *message.TemplateMessage, err error) {
	msgTpl = new(message.TemplateMessage)
	msgTpl.ToUser = w.Config.GetString("toUser")
	msgTpl.TemplateID = w.Config.GetString("templateID")
	msgTpl.Data = make(map[string]*message.TemplateDataItem)

	api := NewApi(w.Config)
	var tq TianQi
	tq, err = api.TianQi()
	loveDate := w.Config.GetString("loveDate")
	birthday := w.Config.GetString("birthday")
	loveDay := carbon.Parse(loveDate).DiffInDays(carbon.Now())
	date := carbon.Now().Format("Y年m月d") + " " + carbon.Now().ToWeekString()
	md := carbon.Parse(birthday).Format("-m-d")
	var nextBirthday string
	var year string
	var birthdayMsg string
	thisYearBirthday := fmt.Sprint(carbon.Now().Year()) + md
	if carbon.Now().Gt(carbon.Parse(thisYearBirthday)) {
		year = fmt.Sprint(carbon.Now().Year() + 1)
		nextBirthday = carbon.Parse(birthday).Format(year + md)
		birthdayMsg = "距离你的生日还有" + fmt.Sprint(carbon.Now().DiffInDays(carbon.Parse(nextBirthday))) + "天"
	} else if carbon.Now().Eq(carbon.Parse(thisYearBirthday)) {
		birthdayMsg = "今天是你的生日"
	} else {
		year = fmt.Sprint(carbon.Now().Year())
		nextBirthday = carbon.Parse(birthday).Format(year + md)
		birthdayMsg = "距离你的生日还有" + fmt.Sprint(carbon.Now().DiffInDays(carbon.Parse(nextBirthday))) + "天"
	}
	//
	msgTpl.Data["date"] = &message.TemplateDataItem{
		Value: date,
		Color: "#434343",
	}
	msgTpl.Data["region"] = &message.TemplateDataItem{
		Value: tq.City,
		Color: "#ff00ff",
	}
	msgTpl.Data["weather"] = &message.TemplateDataItem{
		Value: tq.Wea,
		Color: "#3c78d8",
	}
	msgTpl.Data["maxTemp"] = &message.TemplateDataItem{
		Value: tq.Tem1,
		Color: "#dd7e6b",
	}
	msgTpl.Data["minTemp"] = &message.TemplateDataItem{
		Value: tq.Tem2,
		Color: "#dd7e6b",
	}
	msgTpl.Data["wind_dir"] = &message.TemplateDataItem{
		Value: tq.Win + tq.WinSpeed,
		Color: "#9900ff",
	}
	msgTpl.Data["love_day"] = &message.TemplateDataItem{
		Value: fmt.Sprint(loveDay),
		Color: "#6fa8dc",
	}
	msgTpl.Data["birthday"] = &message.TemplateDataItem{
		Value: birthdayMsg,
		Color: "#ea9999",
	}
	var chp string
	chp, err = api.GetCaiHongPi()
	if err != nil {
		fmt.Printf("Unmarshal fail, err:%v", err)
	}
	msgTpl.Data["note_ch"] = &message.TemplateDataItem{
		Value: chp,
		Color: "#FF0000",
	}
	return msgTpl, nil
}
