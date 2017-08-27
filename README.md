# pushsvr
极光推送服务demo

## 创建极光推送账号
![创建账号](https://raywangblog.files.wordpress.com/2017/08/jgpush-account.png)

## 创建应用并配置应用
![应用信息](https://raywangblog.files.wordpress.com/2017/08/jgpush-app-info.png)

## 推送配置
![推送配置信息](https://raywangblog.files.wordpress.com/2017/08/jgpush-app-push-settings.png)

## 编写Go服务端，对外提供push调用服务
![coding](https://raywangblog.files.wordpress.com/2017/08/jgpush-coding.png)

## 按照极光推送官网文档说明
> 将配置的应用下载下来，导入到Android Studio
![导入应用](https://raywangblog.files.wordpress.com/2017/08/import-module.png)

> 将应用打包到Android设备
![打包应用到设备](https://raywangblog.files.wordpress.com/2017/08/install-app-to-android-device.png)

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
![请求](https://raywangblog.files.wordpress.com/2017/08/jgpush-request.png)
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
## Android移动端即时收到推送的通知消息
![Android收到推送消息](https://raywangblog.files.wordpress.com/2017/08/jgpush-success.png)
