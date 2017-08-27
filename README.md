# pushsvr
极光推送服务demo

## 请求地址
> post请求
```
http://www.ray-xyz.com:6060/push
```
> 参数
kind => 类型: "notif"为通知，"msg"为消息
content => 推送的内容
```
{
	"kind": "notif",
	"content": "Hi, push server!!"
}
```
## 成功时返回
```
{"success":"true","msg":"Successfully pushed data.","data":"{\"sendno\":\"0\",\"msg_id\":\"1973830861\"}"}
```
## 后台打印
```
Push server runs on port 6060.2017/08/28 03:27:39 params =>  {notif Hi, push server!!}
2017/08/28 03:27:39 kind =>  notif
2017/08/28 03:27:39 content =>  Hi, push server!!
2017/08/28 03:27:39 >>> Payload detail info:
{"platform":["android","ios"],"audience":"all","notification":{"alert":"ray_alert","android":{"alert":"Hi, push server!!"},"ios":{"alert":"Hi, push server!!"}},"options":{"apns_production":false}}
```
