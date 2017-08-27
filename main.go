// @Author: Raywang
// @Date: 2017-08-28
// 极光JPush 推送服务
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	jpc "github.com/ylywyn/jpush-api-go-client"
)

// appKey和Master secret到极光推送后台应用中查找
const (
	// 应用唯一标识
	appKey = "e23387da6e8637eaeaa2c3e5"
	// Master secret
	secret = "af76f91befb027abd9755bff"
)

var pusher *jpc.PushClient
var port = 6060

type result struct {
	Success bool        `json:"success,string"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
}

// 支持平台设置
func addPlatforms() *jpc.Platform {
	var pf jpc.Platform
	pf.Add(jpc.ANDROID)
	pf.Add(jpc.IOS)
	//pf.All()
	return &pf
}

// 设置接受推送推送内容的客户端
func setAudience() *jpc.Audience {
	var ad jpc.Audience
	// s := []string{"1", "2", "3"}
	// ad.SetTag(s)
	// ad.SetAlias(s)
	// ad.SetID(s)
	// 这里设置为推送给所有应用用户，应根据实际情况推送给特定用户或用户群
	ad.All()
	return &ad
}

// 初始化通知
func setNotice(payload *jpc.PayLoad, content string) {
	var notice jpc.Notice
	notice.SetAlert("ray_alert")
	notice.SetAndroidNotice(&jpc.AndroidNotice{Alert: content})
	notice.SetIOSNotice(&jpc.IOSNotice{Alert: content})
	payload.SetNotice(&notice)
}

// 初始化消息
func setMessage(payload *jpc.PayLoad, content string) {
	var msg jpc.Message
	msg.Title = "Message from push server"
	msg.Content = content
	payload.SetMessage(&msg)
}

func createPayload(kind, content string) []byte {
	payload := jpc.NewPushPayLoad()
	payload.SetPlatform(addPlatforms())
	payload.SetAudience(setAudience())
	if kind == "notif" {
		setNotice(payload, content)
	}
	if kind == "msg" {
		setMessage(payload, content)
	}
	bytes, _ := payload.ToBytes()
	log.Printf(">>> Payload detail info:\n%s\r\n", string(bytes))
	return bytes
}

type reqparams struct {
	Kind    string
	Content string
}

func push(w http.ResponseWriter, r *http.Request) {
	var params reqparams
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		log.Println("Parse parameters error. => ", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Parse parameters error."))
		return
	}
	log.Println("params => ", params)
	// kind为`notif`时推送通知，为`msg`时推送消息
	kind := params.Kind
	// 要推送的内容
	content := params.Content
	log.Println("kind => ", kind)
	log.Println("content => ", content)
	// 推送
	str, err := pusher.Send(createPayload(kind, content))
	if err != nil {
		log.Println("Push error. => ", err.Error())
		data, _ := json.Marshal(result{
			Success: false,
			Msg:     err.Error(),
		})
		w.Write(data)
		return
	}
	data, _ := json.Marshal(result{
		Success: true,
		Msg:     "Successfully pushed data.",
		Data:    str,
	})
	w.Write(data)
}

func main() {
	pusher = jpc.NewPushClient(secret, appKey)
	http.HandleFunc("/push", func(w http.ResponseWriter, r *http.Request) {
		push(w, r)
	})
	fmt.Printf("Push server runs on port %d.", port)
	if err := http.ListenAndServe(":"+strconv.Itoa(port), nil); err != nil {
		log.Println("Init push server failed. err => ", err.Error())
	}
}
