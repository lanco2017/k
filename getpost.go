package main

import (
	// "fmt"
	"log"

	"net/http"
	// "os"
	"regexp"

	// "strconv"
	"strings"
	
	// "github.com/line/line-bot-sdk-go/linebot"

	"bytes"

	"io/ioutil"

	// "image/jpeg"

    // "crypto/md5"
    // "encoding/hex"

    // "encoding/json"
    // "github.com/bitly/go-simplejson"
)

func HttpGET_(url string) string {
	//https://internal-api.ifttt.com/maker
	log.Print("已經進來 GET")
	log.Print("url = " + url)

	// url := "https://hooks.zapier.com/hooks/catch/132196/txma4i/"

	req, err := http.NewRequest(
		"GET",
		url,
		nil,
	)
	if err != nil {
		log.Print(err)
		return ""
	}

	// Content-Type 設定
	//req.Header.Set("Accept", "application/vnd.tosslab.jandi-v2+json")
	//req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Print(err)		
		return ""
	}
	defer resp.Body.Close()

	log.Print(err)

	//http://cepave.com/http-restful-api-with-golang/
    log.Print("HttpGET_response Status = ")
    log.Print(resp.Status)
    log.Print("HttpGET_response Headers = ")
    log.Print(resp.Header)
    rebody, _ := ioutil.ReadAll(resp.Body)
    log.Print("response Body = " + string(rebody))
	//http://cepave.com/http-restful-api-with-golang/

	return string(rebody) //return err
}

func HttpPost_JSON(mode, jsonStr, url  string) string {
	log.Print("已經進來 any_JSON POST by " + mode)

	if mode=="LINE"{
		url = "https://notify-api.line.me/api/notify"
	}

	req, err := http.NewRequest(
		"POST",
		url,
		bytes.NewBuffer([]byte(jsonStr)),
	)
	if err != nil {
		log.Print(err)
		//return err.Error()
		return ""
	}

	// Content-Type 設定                                                          其實是因為有這個的關係所以才不能 function 化
	// if mode == "JANDI" {
	// 	req.Header.Set("Accept", "application/vnd.tosslab.jandi-v2+json")
	// 	req.Header.Set("Content-Type", "application/json")
	// }
	switch mode {
		case "JANDI":
			req.Header.Set("Accept", "application/vnd.tosslab.jandi-v2+json")
			req.Header.Set("Content-Type", "application/json")
		case "IFTTT","ZAPIER":
			req.Header.Set("Content-Type", "application/json")
		case "LINE":
			req.Header.Set("Authorization", "Bearer XXX")
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Print(err)		
		//return err.Error()
		return ""
	}
	defer resp.Body.Close()

	log.Print(err)

	//http://cepave.com/http-restful-api-with-golang/
    log.Print("HttpPost_JSON_response Status = ")
    log.Print(resp.Status)
    log.Print("HttpPost_JSON_response Headers = ")
    log.Print(resp.Header)
    rebody, _ := ioutil.ReadAll(resp.Body)
    log.Print("response Body = " +string(rebody))
	//http://cepave.com/http-restful-api-with-golang/

	return string(rebody)
}

func HttpPost_Zapier(body , title_text, this_id, codename string) {
	body = strings.Replace(body,"\n", `\n`, -1)
	title_text = strings.Replace(title_text,"\n", `\n`, -1)
	this_id = strings.Replace(this_id,"\n", `\n`, -1)
	codename = strings.Replace(codename,"\n", `\n`, -1)
	//https://internal-api.ifttt.com/maker
	log.Print("已經進來 Zapier POST")
	log.Print("body = " + body)
	log.Print("title_text = " + title_text)
	log.Print("this_id = " + this_id)

	url := "https://hooks.zapier.com/hooks/catch/132196/txma4i/"
	jsonStr := `{
		"value1":"` + body + `",
		"value2": "` + title_text + `",
		"value3": "` + this_id + `",
		"value4": "` + codename + `"
	}`

	post_string_1 := HttpPost_JSON("ZAPIER", jsonStr, url)
	if post_string_1 =="" {
		log.Print("執行 post_string_1 " + `HttpPost_JSON("ZAPIER", jsonStr, url) 出錯！！！！！！` + " 無回應資料")
	}else{
		log.Print("執行 post_string_1 " + `HttpPost_JSON("ZAPIER", jsonStr, url)` + " 後 = " + post_string_1)
	}	

	demo_string := HttpPost_JSON("ZAPIER", jsonStr,"https://hooks.zapier.com/hooks/catch/1817473/homfl8/")
	if demo_string =="" {
		log.Print("執行 demo_string " + `HttpPost_JSON("ZAPIER", jsonStr, url) 出錯！！！！！！` + " 無回應資料")
	}else{
		log.Print("執行 demo_string " + `HttpPost_JSON("ZAPIER", jsonStr, url)` + " 後 = " + demo_string)
	}		


	// req, err := http.NewRequest(
	// 	"POST",
	// 	url,
	// 	bytes.NewBuffer([]byte(jsonStr)),
	// )
	// if err != nil {
	// 	log.Print(err)
	// 	return err
	// }

	// // Content-Type 設定
	// //req.Header.Set("Accept", "application/vnd.tosslab.jandi-v2+json")
	// req.Header.Set("Content-Type", "application/json")

	// client := &http.Client{}
	// resp, err := client.Do(req)
	// if err != nil {
	// 	log.Print(err)		
	// 	return err
	// }
	// defer resp.Body.Close()

	// log.Print(err)

	// //http://cepave.com/http-restful-api-with-golang/
 //    // log.Print("HttpPost_Zapier_response Status = ")
 //    // log.Print(resp.Status)
 //    // log.Print("HttpPost_Zapier_response Headers = ")
 //    // log.Print(resp.Header)
 //    // rebody, _ := ioutil.ReadAll(resp.Body)
 //    // log.Print("response Body = " +string(rebody))
	// //http://cepave.com/http-restful-api-with-golang/

	// return err
}

//func //HttpPost_IFTTT(body , title_text, this_id string) error {
func HttpPost_IFTTT_for_kkcpct(body , title_text, this_id string) {
	body = strings.Replace(body,"\n", `\n<br>`, -1)
	title_text = strings.Replace(title_text,"\n", `\n`, -1)
	this_id = strings.Replace(this_id,"\n", `\n`, -1)
	//https://internal-api.ifttt.com/maker
	log.Print("已經進來 IFTTT POST")
	log.Print("body = " + body)
	log.Print("title_text = " + title_text)
	log.Print("this_id = " + this_id)

	url := "https://maker.ifttt.com/trigger/kk_linebot/with/key/WJCRNxQhGJuzPd-sUDext"
	jsonStr := `{
		"value1":"` + body + `",
		"value2": "` + title_text + `",
		"value3": "` + this_id + `"
	}`

	// jsonStr := `{
	// 	"message":"` + body + `"
	// }`

	post_string_1 := HttpPost_JSON("IFTTT", jsonStr, url)
	if post_string_1 =="" {
		log.Print("執行 post_string_1 " + `HttpPost_JSON("IFTTT", jsonStr, url) 出錯！！！！！！` + " 無回應資料")
	}else{
		log.Print("執行 post_string_1 " + `HttpPost_JSON("IFTTT", jsonStr, url)` + " 後 = " + post_string_1)
	}
}

//func //HttpPost_IFTTT(body , title_text, this_id string) error {
func HttpPost_IFTTT(body , title_text, this_id string) {
	body = strings.Replace(body,"\n", `\n`, -1)
	title_text = strings.Replace(title_text,"\n", `\n`, -1)
	this_id = strings.Replace(this_id,"\n", `\n`, -1)
	//https://internal-api.ifttt.com/maker
	log.Print("已經進來 IFTTT POST")
	log.Print("body = " + body)
	log.Print("title_text = " + title_text)
	log.Print("this_id = " + this_id)

	url := "https://maker.ifttt.com/trigger/linebotexample/with/key/WJCRNxQhGJuzPd-sUDext"
	jsonStr := `{
		"value1":"` + body + `",
		"value2": "` + title_text + `",
		"value3": "` + this_id + `"
	}`

	post_string_1 := HttpPost_JSON("IFTTT", jsonStr, url)
	if post_string_1 =="" {
		log.Print("執行 post_string_1 " + `HttpPost_JSON("IFTTT", jsonStr, url) 出錯！！！！！！` + " 無回應資料")
	}else{
		log.Print("執行 post_string_1 " + `HttpPost_JSON("IFTTT", jsonStr, url)` + " 後 = " + post_string_1)
	}	

	// req, err := http.NewRequest(
	// 	"POST",
	// 	url,
	// 	bytes.NewBuffer([]byte(jsonStr)),
	// )
	// if err != nil {
	// 	log.Print(err)
	// 	return err
	// }

	// // Content-Type 設定
	// //req.Header.Set("Accept", "application/vnd.tosslab.jandi-v2+json")
	// req.Header.Set("Content-Type", "application/json")

	// client := &http.Client{}
	// resp, err := client.Do(req)
	// if err != nil {
	// 	log.Print(err)		
	// 	return err
	// }
	// defer resp.Body.Close()

	// log.Print(err)

	// //http://cepave.com/http-restful-api-with-golang/
 //    log.Print("//HttpPost_IFTTT_response Status = ")
 //    log.Print(resp.Status)
 //    log.Print("//HttpPost_IFTTT_response Headers = ")
 //    log.Print(resp.Header)
 //    rebody, _ := ioutil.ReadAll(resp.Body)
 //    log.Print("response Body = " +string(rebody))
	// //http://cepave.com/http-restful-api-with-golang/

	// return err
}

func send_to_JANDI(text, target_item, user_talk, userImageUrl, userStatus, target_id_code string) (string) {
	//reg := regexp.MustCompile("^(給教會)(\\s|　|:|;|：|；|-|－)(.*)")
	reg := regexp.MustCompile("^(給教會)(\\s|　|:|;|：|；|-|－)(.{0,}\\n{0,}.{0,}\\n{0,}.{0,}\\n{0,}.{0,}\\n{0,}.{0,}\\n{0,}.{0,}\\n{0,}.{0,}\\n{0,}.{0,}\\n{0,}.{0,}\\n{0,}.{0,}\\n{0,}.{0,}\\n{0,}.{0,}\\n{0,}.{0,}\\n{0,}.{0,}\\n{0,}.{0,}\\n{0,}.{0,}\\n{0,}.{0,}\\n{0,}.{0,}\\n{0,}.{0,}\\n{0,}.{0,}\\n{0,}.{0,}\\n{0,}.{0,}\\n{0,}.{0,}\\n{0,})")
	log.Print("--抓取分析 給教會 觀察--")
	log.Print("$1 = 觸發關鍵字 = " + reg.ReplaceAllString(text, "$1"))
	log.Print("$2 = 分割符 = " + reg.ReplaceAllString(text, "$2"))
	log.Print("$3 = 第一主題 = " + reg.ReplaceAllString(text, "$3"))
	log.Print("--抓取分析結束--")

	if (reg.ReplaceAllString(text, "$1")=="給教會") && (reg.ReplaceAllString(text, "$3")!="") && (reg.ReplaceAllString(text, "$3")!="給教會") {
		get_text := reg.ReplaceAllString(text, "$3")
		//HttpPost_JANDI(target_item + " [" + user_talk + "](" + userImageUrl + ")：" + get_text + `\n` + userStatus, "orange" , "LINE 同步零時差通知：【" + target_item + "】" + user_talk + " 給教會",target_id_code)
		HttpPost_IFTTT_for_kkcpct("某 " + target_item + " 透過「給教會」發送的傳話"  + "：" + `\n<br>` + get_text , get_text ,"")
		//HttpPost_IFTTT(target_item + " " + user_talk + "：" + get_text + `\n<br>` + userImageUrl + `\n<br>` + userStatus, "LINE 同步零時差通知：【" + target_item + "】" + user_talk + " 給教會",target_id_code)
		//HttpPost_Zapier(target_item + " [" + user_talk + "](" + userImageUrl + ")：" + get_text + `\n` + userStatus, "LINE 同步零時差通知：【" + target_item + "】" + user_talk + " 給教會" ,target_id_code,user_talk)
		//傳成功就會傳成功，方便執行後面的機器人動作
		text = "已經傳送給教會"
	}

	return text
}

//2017.01.08 改 分離格式 與 POST 動作

func HttpPost_JANDI(body, connectColor, title, code string) {
	body = strings.Replace(body,"\n", `\n`, -1)
	title = strings.Replace(title,"\n", `\n`, -1)
	code = strings.Replace(code,"\n", `\n`, -1)

	log.Print("已經進來 JANDI POST")
	log.Print("body = " + body)
	log.Print("connectColor = " + connectColor)
	log.Print("title = " + title)
	log.Print("code = " + code)

	jsonStr := `{
		"body":"` + body + `",
		"connectColor":"` + connectColor + `",
		"connectInfo" : [{
				"title" : "` + title + `",
				"description" : "這是來自 LINE BOT 的通風報信",
				"imageUrl": "https://line.me/R/ti/p/@rgp6918s"
		},{
				"title" : "參考數據",
				"description" : "` + code + `"
		}]
	}`

	//url := "https://wh.jandi.com/connect-api/webhook/12797246/78c9e40acac43d634e321a9c306815c3"//"https://wh.jandi.com/connect-api/webhook/12797246/a170893b87f09b238a9a85919efc6389"
	url := "https://wh.jandi.com/connect-api/webhook/11691684/a481ad19001738278c7c034475d10af7"//"https://wh.jandi.com/connect-api/webhook/12797246/a170893b87f09b238a9a85919efc6389"
	post_string_1 := HttpPost_JSON("JANDI", jsonStr, url)
	if post_string_1 =="" {
		log.Print("執行 post_string_1 " + `HttpPost_JSON("JANDI", jsonStr, url) 出錯！！！！！！` + " 無回應資料")
	}else{
		log.Print("執行 post_string_1 " + `HttpPost_JSON("JANDI", jsonStr, url)` + " 後 = " + post_string_1)
	}

	// url = "https://wh.jandi.com/connect-api/webhook/12797246/f662337353f2cfc58443068b5b69923d"
	// post_string_2 := HttpPost_JSON("JANDI", jsonStr, url)
	// if post_string_2 =="" {
	// 	log.Print("執行 post_string_2 " + `HttpPost_JSON("JANDI", jsonStr, url)` + " 後無回應資料")
	// }else{
	// 	log.Print("執行 post_string_2 " + `HttpPost_JSON("JANDI", jsonStr, url)` + " 後 = " + post_string_2)
	// }
}

//原
// func //HttpPost_JANDI(body, connectColor, title, code string) error {
// 	body = strings.Replace(body,"\n", `\n`, -1)
// 	title = strings.Replace(title,"\n", `\n`, -1)
// 	code = strings.Replace(code,"\n", `\n`, -1)

// 	log.Print("已經進來 JANDI POST")
// 	log.Print("body = " + body)
// 	log.Print("connectColor = " + connectColor)
// 	log.Print("title = " + title)
// 	log.Print("code = " + code)

// 	url := "https://wh.jandi.com/connect-api/webhook/12797246/78c9e40acac43d634e321a9c306815c3"
// 	jsonStr := `{
// 		"body":"` + body + `",
// 		"connectColor":"` + connectColor + `",
// 		"connectInfo" : [{
// 				"title" : "` + title + `",
// 				"description" : "這是來自 LINE BOT 的通風報信",
// 				"imageUrl": "https://line.me/R/ti/p/@rgp6918s"
// 		},{
// 				"title" : "參考數據",
// 				"description" : "` + code + `"
// 		}]
// 	}`

// 	req, err := http.NewRequest(
// 		"POST",
// 		url,
// 		bytes.NewBuffer([]byte(jsonStr)),
// 	)
// 	if err != nil {
// 		log.Print(err)
// 		return err
// 	}

// 	// Content-Type 設定
// 	req.Header.Set("Accept", "application/vnd.tosslab.jandi-v2+json")
// 	req.Header.Set("Content-Type", "application/json")

// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		log.Print(err)		
// 		return err
// 	}
// 	defer resp.Body.Close()

// 	log.Print(err)

// 	//http://cepave.com/http-restful-api-with-golang/
//     log.Print("//HttpPost_JANDI_response Status = ")
//     log.Print(resp.Status)
//     log.Print("//HttpPost_JANDI_response Headers = ")
//     log.Print(resp.Header)
//     rebody, _ := ioutil.ReadAll(resp.Body)
//     log.Print("response Body = " +string(rebody))
// 	//http://cepave.com/http-restful-api-with-golang/

// 	return err
// }

