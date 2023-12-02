package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

const BOT_HOST = "https://a1.fanbook.mobi/api"

// const BOT_TOKEN = "8eabb0350429aa292f1e4ad60a2448a267ed8142c3be0581272e80240f5df893"
const BOT_TOKEN = "88892d2897c2bb61fe26f77272bd7bc65d3b86e0a90000ee7a9a8d8f3d68c71d2ab6726620e9d29fd83e82d225cc190c"

const GET_ME = BOT_HOST + "/bot" + "/bot" + BOT_TOKEN + "/getMe"
const SEND_TXT_MSG = BOT_HOST + "/bot" + "/bot" + BOT_TOKEN + "/sendMessage"

type KPRsp struct {
	OK     bool        `json:"ok"`
	Result KPRspResult `json:"result"`
}

type KPRspResult struct {
	Id                      int    `json:"id"`
	IsBot                   bool   `json:"is_Bot"`
	FirstName               string `json:"first_Name"`
	LastName                string `json:"last_name"`
	Username                string `json:"username"`
	Avatar                  string `json:"avatar"`
	UserToken               string `json:"user_token"`
	OwnerId                 int    `json:"owner_id"`
	CanJoinGroups           bool   `json:"can_join_groups"`
	CanReadAllGroupMessages bool   `json:"can_read_all_group_messages"`
	SupportsInlineQueries   bool   `json:"supports_inline_queries"`
}

func main() {
	rsp, err := http.Get(GET_ME)
	if err != nil {
		log.Println("request of getMe error: ", err.Error())
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println("close response's body error: ", err.Error())
		}
	}(rsp.Body)
	log.Println(rsp.Header)
	log.Println(rsp.Status)

	body, err := ioutil.ReadAll(rsp.Body)
	//body, err := io.Copy(os.Stdout, rsp.Body)
	if err != nil {
		log.Println("read body error: ", err.Error())
		return
	}

	log.Println("get info: ", string(body))

	var jsonData KPRsp
	jsonErr := json.Unmarshal(body, &jsonData)
	if jsonErr != nil {
		log.Println("json error: ", jsonErr.Error())
		return
	}

	log.Printf("json data: %v\n", jsonData)
	//fmt.Printf("json data: %v\n", jsonData)

	// 发消息
	HttpPost(SEND_TXT_MSG)
}

func HttpPost(urlP string) []byte {
	// 参数
	data := make(map[string]interface{})
	// data["chat_id"] = 534932844435587072
	data["chat_id"] = 380272446231543810
	data["text"] = "test123 20231201"
	data["selective"] = false
	data["disable_web_page_preview"] = false
	data["disable_notification"] = true

	// 序列化
	bytesData, _ := json.Marshal(data)

	resp, err := http.Post(urlP, "application/json", bytes.NewReader(bytesData))
	if err != nil {
		fmt.Println(err)
		return nil
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
		}
	}(resp.Body)

	body, _ := ioutil.ReadAll(resp.Body)

	var jsonData KPRsp
	jsonErr := json.Unmarshal(body, &jsonData)
	if jsonErr != nil {
		log.Println("json error: ", jsonErr.Error())
		return nil
	}

	log.Printf("json data: %v\n", jsonData)

	return body
}
