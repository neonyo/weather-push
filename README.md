# weather-push
微信推送
添加配置信息config.json

TianXingKey ：https://www.tianapi.com/天行彩虹屁接口key

loveDate ： 恋爱时间格式：YYYY-mm-dd

birthday: 生日时间格式：YYYY-mm-dd

appId： 微信公众号appID

appSecret： 微信公众号appSecret

toUser：消息接收人


templateID：模版ID

//天气API，https://tianqiapi.com/index/doc

tqAppId：天气appId

tqAppSecret:天气 AppSecret

city：获取天气城市

设置好相关配置执行下面代码

go mod vendor

go run main

看看是否接收成功

成功后编译打包：GOOS=linux GOARCH=amd64 go build -o main main.go

然后把编译后的 main 文件和config文件 打包成zip文件
上传至腾讯云函数即可，设置触发时间为每天x点

微信模板消息：<br>
{{date.DATA}}<br>
地区：{{region.DATA}}<br>
天气：{{weather.DATA}}<br>
最高温度：{{maxTemp.DATA}}<br>
最低温度：{{minTemp.DATA}}<br>
风向：{{wind_dir.DATA}}<br>
今天是我们恋爱的第{{love_day.DATA}}天<br>
{{birthday.DATA}}<br>
{{note_ch.DATA}}
