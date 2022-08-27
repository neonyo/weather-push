package core

type CaiHongPi struct {
	Code     int        `json:"code"`
	Msg      string     `json:"msg"`
	NewsList []NewsList `json:"newslist"`
}
type NewsList struct {
	Content string `json:"content"`
}

type TianQi struct {
	Cityid        string `json:"cityid"`
	Date          string `json:"date"`
	Week          string `json:"week"`
	UpdateTime    string `json:"update_time"`
	City          string `json:"city"`
	CityEn        string `json:"cityEn"`
	Country       string `json:"country"`
	CountryEn     string `json:"countryEn"`
	Wea           string `json:"wea"`
	WeaImg        string `json:"wea_img"`
	Tem           string `json:"tem"`
	Tem1          string `json:"tem1"`
	Tem2          string `json:"tem2"`
	Win           string `json:"win"`
	WinSpeed      string `json:"win_speed"`
	WinMeter      string `json:"win_meter"`
	Humidity      string `json:"humidity"`
	Visibility    string `json:"visibility"`
	Pressure      string `json:"pressure"`
	Air           string `json:"air"`
	AirPm25       string `json:"air_pm25"`
	AirLevel      string `json:"air_level"`
	AirTips       string `json:"air_tips"`
	Alarm         Alarm  `json:"alarm"`
	WinSpeedDay   string `json:"win_speed_day"`
	WinSpeedNight string `json:"win_speed_night"`
	Aqi           Aqi    `json:"aqi"`
}
type Alarm struct {
	AlarmType    string `json:"alarm_type"`
	AlarmLevel   string `json:"alarm_level"`
	AlarmContent string `json:"alarm_content"`
}
type Aqi struct {
	UpdateTime string `json:"update_time"`
	Cityid     string `json:"cityid"`
	City       string `json:"city"`
	CityEn     string `json:"cityEn"`
	Country    string `json:"country"`
	CountryEn  string `json:"countryEn"`
	Air        string `json:"air"`
	AirLevel   string `json:"air_level"`
	AirTips    string `json:"air_tips"`
	Pm25       string `json:"pm25"`
	Pm25Desc   string `json:"pm25_desc"`
	Pm10       string `json:"pm10"`
	Pm10Desc   string `json:"pm10_desc"`
	O3         string `json:"o3"`
	O3Desc     string `json:"o3_desc"`
	No2        string `json:"no2"`
	No2Desc    string `json:"no2_desc"`
	So2        string `json:"so2"`
	So2Desc    string `json:"so2_desc"`
	Co         string `json:"co"`
	CoDesc     string `json:"co_desc"`
	Kouzhao    string `json:"kouzhao"`
	Yundong    string `json:"yundong"`
	Waichu     string `json:"waichu"`
	Kaichuang  string `json:"kaichuang"`
	Jinghuaqi  string `json:"jinghuaqi"`
}
