package main

import (
	// "fmt"
	"log"

	// "net/http"
	// "os"
	"regexp"

	// "strconv"
	//"strings"
	
	// "github.com/line/line-bot-sdk-go/linebot"

	// "bytes"

	// "io/ioutil"

	// "image/jpeg"

    // "crypto/md5"
    // "encoding/hex"

    // "encoding/json"
    // "github.com/bitly/go-simplejson"
)

func lo(text string) (string) {
	//https://gitter.im/kkdai/LineBotTemplate?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge：也可以透過 string.Contains("我要找的字", 原始字串) 來判斷
	print_string := text
	text = real_num(text)

	// bible_json_string := ""
	// bible_text_string := ""

	// if GetMD5Hash(text) == "5fed8bfd031fa698e125567b128d1024" {
	// 	print_string = "5fed8bfd031fa698e125567b128d1024"
	// 	return print_string
	// }

	switch GetMD5Hash(text){
		case "6a5401cd1d5e30c52a79ee61b34262dc","5fed8bfd031fa698e125567b128d1024":
			print_string = "5fed8bfd031fa698e125567b128d1024"
			return print_string
		case "546eabd81e99ec08b8b0301af80310d9","9b7340292bee67f5a81c49f023f9f867":
			print_string = "546eabd81e99ec08b8b0301af80310d9"
			return print_string
		default:
	}


	// reg_bible_plus := regexp.MustCompile("^(.文)(.文)(聖經|聖書|Bible|bible|ｂｉｂｌｅ|Ｂｉｂｌｅ|客家聖經|台語聖經巴克禮漢羅|台語聖經巴克禮全羅|台語聖經馬雅各全羅|台語聖經馬雅各漢羅|台語聖經|閩南語聖經|台語聖經全羅|全民台語聖經全羅|台語聖經漢羅|全民台語聖經漢羅|中文聖經|中文聖經和合本修訂版|Rcuv|rcuv|ｒｃｕｖ|Ｒｃｕｖ|文言文聖經|深文理和合本|中文聖經新譯本|ncv|Ncv|Ｎｃｖ|ｎｃｖ|聖經中文譯本修訂版|tcv|TCV|Ｔｃｖ|ＴＣＶ|日文聖經|日本語聖書|JP bible|JP Bible|Jp bible|韓文聖經|KR bible|Kr Bible|Kr bible|英文聖經|英語聖書|Kjv|kjv|Ｋｊｖ|ｋｊｖ|Eng bible|ENG Bible|English bible|BBE|Bbe|bbe|ＢＢＥ|Ｂｂｅ|ｂｂｅ|英文聖經BBE|英文聖經WEB|WEB|Web|web|ＷＥＢ|Ｗｅｂ|ｗｅｂ|英文聖經ASV|ASV|Asv|asv|ＡＳＶ|Ａｓｖ|ａｓｖ|英文聖經Darby|darby|DARBY|Ｄａｒｂｙ|ＤＡＲＢＹ|ｄａｒｂｙ|英文聖經ERV|erv|ERV|Erv|ＥＲＶ|Ｅｒｖ|ｅｒｖ|希臘聖經|lxx|LXX|Lxx|ＬＸＸ|Ｌｘｘ|ｌｘｘ|馬索拉聖經|bhs|Bhs|BHS|ＢＨＳ|Ｂｈｓ|ｂｈｓ|越南聖經|俄文聖經|多國聖經|多語聖經|多語言聖經|多國語聖經|多國語言聖經|allbible|all bible|All bible|All Bible|總和聖經|綜合聖經|研究聖經|聖經研究|多版聖經|多版本言聖經|Allbible)(\\s|　|:|;|：|；|-|－)([\u0590-\u05ff\u0370—\u03ff\u1f00-\u1fff\u2c80-\u2cff\u0400-\u04ff\u0500-\u052f\uff21-\uff3a\uff41-\uff5a\uff10-\uff19\u30a0-\u30ff\u3040-\u30ff\u31f0-\u31ff\u4e00-\u9fff\u1100-\u11ff\u3130—\u318f\uac00-\ud7af\ua960-\ua97f_a-zA-Z0-9]*)\\D*([0-9.]{0,})(\\s|　|:|;|：|；){0,}\\D*([0-9.\\-－～\\~]{0,})")
	// log.Print("--抓取分析觀察--")
	// log.Print("$1 = 觸發關鍵字(X文) = " + reg_bible_plus.ReplaceAllString(text, "$1"))
	// log.Print("$2 = 觸發關鍵字(X文) = " + reg_bible_plus.ReplaceAllString(text, "$2"))
	// log.Print("$3 = 觸發關鍵字 = " + reg_bible_plus.ReplaceAllString(text, "$3"))
	// log.Print("$4 = 分割符 = " + reg_bible_plus.ReplaceAllString(text, "$4"))
	// log.Print("$5 = 第一主題 = " + reg_bible_plus.ReplaceAllString(text, "$5"))
	// log.Print("$6 = 章 = " + reg_bible_plus.ReplaceAllString(text, "$6"))
	// log.Print("$7 = 章節分割符 = " + reg_bible_plus.ReplaceAllString(text, "$7"))
	// log.Print("$8 = 節 = " + reg_bible_plus.ReplaceAllString(text, "$8"))
	// log.Print("--抓取分析結束--")

	// chap_string := reg_bible_plus.ReplaceAllString(text, "$6")	//章
	// sec_string := reg_bible_plus.ReplaceAllString(text, "$8")	//節
	// sec_string = strings.Replace(sec_string,`－`, "-", -1)
	// sec_string = strings.Replace(sec_string,`～`, "-", -1)
	// sec_string = strings.Replace(sec_string,`~`, "-", -1)
	// sec_string = strings.Replace(sec_string,` ~ `, "-", -1)
	// bible_short_name := ""

	// switch reg.ReplaceAllString(text, "$1"){
	// 	case "中文",:
	// 		return
	// }


	//優先處理這邊
	//如果處理完成，就轉移
	//https://coggle.it/diagram/WIgDvVUIaygPvZsf

	//2017.01.03+
	//reg := regexp.MustCompile("^(懶|聖經|Bible)(\\s|　|:|;|：|；)([\u4e00-\u9fa5_a-zA-Z0-9]*)\\D*([0-9.]{1,})") //fmt.Printf("%q\n", reg.FindAllString(text, -1))
	//2017.01.04+	https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/07.3.md
	//reg := regexp.MustCompile("(聖經|聖書|Bible|bible)(\\s|　|:|;|：|；|-|－)([\uff21-\uff3a\uff41-\uff5a\uff10-\uff19\u30a0-\u30ff\u3040-\u309f\u4e00-\u9fd5_a-zA-Z0-9]*)\\D*([0-9.]{0,})(\\s|　|:|;|：|；){0,}\\D*([0-9.]{0,})") //fmt.Printf("%q\n", reg.FindAllString(text, -1))
	//2017.01.05+
	// reg := regexp.MustCompile("(聖經|聖書|Bible|bible|日文聖經|日本語聖書|JP bible|JP Bible|Jp bible|韓文聖經|KR bible|Kr Bible|Kr bible|英文聖經|英語聖書|Eng bible|ENG Bible|English bible|越南聖經|俄文聖經|多國聖經|多語聖經|多語言聖經|多國語聖經|多國語言聖經|allbible|all bible|All bible|All Bible)(\\s|　|:|;|：|；|-|－)([\uff21-\uff3a\uff41-\uff5a\uff10-\uff19\u30a0-\u30ff\u3040-\u309f\u4e00-\u9fd5_a-zA-Z0-9]*)\\D*([0-9.]{0,})(\\s|　|:|;|：|；){0,}\\D*([0-9.]{0,})") //fmt.Printf("%q\n", reg.FindAllString(text, -1))
	//2017.01.06+ //https://regexper.com/#%5E(%E8%81%96%E7%B6%93%7C%E8%81%96%E6%9B%B8%7CBible%7Cbible%7C%EF%BD%82%EF%BD%89%EF%BD%82%EF%BD%8C%EF%BD%85%7C%EF%BC%A2%EF%BD%89%EF%BD%82%EF%BD%8C%EF%BD%85%7C%E5%8F%B0%E8%AA%9E%E8%81%96%E7%B6%93%E5%B7%B4%E5%85%8B%E7%A6%AE%E6%BC%A2%E7%BE%85%7C%E5%8F%B0%E8%AA%9E%E8%81%96%E7%B6%93%7C%E9%96%A9%E5%8D%97%E8%AA%9E%E8%81%96%E7%B6%93%7C%E5%8F%B0%E8%AA%9E%E8%81%96%E7%B6%93%E5%85%A8%E7%BE%85%7C%E5%85%A8%E6%B0%91%E5%8F%B0%E8%AA%9E%E8%81%96%E7%B6%93%E5%85%A8%E7%BE%85%7C%E5%8F%B0%E8%AA%9E%E8%81%96%E7%B6%93%E6%BC%A2%E7%BE%85%7C%E5%85%A8%E6%B0%91%E5%8F%B0%E8%AA%9E%E8%81%96%E7%B6%93%E6%BC%A2%E7%BE%85%7C%E4%B8%AD%E6%96%87%E8%81%96%E7%B6%93%7C%E4%B8%AD%E6%96%87%E8%81%96%E7%B6%93%E5%92%8C%E5%90%88%E6%9C%AC%E4%BF%AE%E8%A8%82%E7%89%88%7CRcuv%7Crcuv%7C%EF%BD%92%EF%BD%83%EF%BD%95%EF%BD%96%7C%EF%BC%B2%EF%BD%83%EF%BD%95%EF%BD%96%7C%E6%96%87%E8%A8%80%E6%96%87%E8%81%96%E7%B6%93%7C%E6%B7%B1%E6%96%87%E7%90%86%E5%92%8C%E5%90%88%E6%9C%AC%7C%E4%B8%AD%E6%96%87%E8%81%96%E7%B6%93%E6%96%B0%E8%AD%AF%E6%9C%AC%7Cncv%7CNcv%7C%EF%BC%AE%EF%BD%83%EF%BD%96%7C%EF%BD%8E%EF%BD%83%EF%BD%96%7C%E8%81%96%E7%B6%93%E4%B8%AD%E6%96%87%E8%AD%AF%E6%9C%AC%E4%BF%AE%E8%A8%82%E7%89%88%7Ctcv%7CTCV%7C%EF%BC%B4%EF%BD%83%EF%BD%96%7C%EF%BC%B4%EF%BC%A3%EF%BC%B6%7C%E6%97%A5%E6%96%87%E8%81%96%E7%B6%93%7C%E6%97%A5%E6%9C%AC%E8%AA%9E%E8%81%96%E6%9B%B8%7CJP%20bible%7CJP%20Bible%7CJp%20bible%7C%E9%9F%93%E6%96%87%E8%81%96%E7%B6%93%7CKR%20bible%7CKr%20Bible%7CKr%20bible%7C%E8%8B%B1%E6%96%87%E8%81%96%E7%B6%93%7C%E8%8B%B1%E8%AA%9E%E8%81%96%E6%9B%B8%7CEng%20bible%7CENG%20Bible%7CEnglish%20bible%7C%E8%B6%8A%E5%8D%97%E8%81%96%E7%B6%93%7C%E4%BF%84%E6%96%87%E8%81%96%E7%B6%93%7C%E5%A4%9A%E5%9C%8B%E8%81%96%E7%B6%93%7C%E5%A4%9A%E8%AA%9E%E8%81%96%E7%B6%93%7C%E5%A4%9A%E8%AA%9E%E8%A8%80%E8%81%96%E7%B6%93%7C%E5%A4%9A%E5%9C%8B%E8%AA%9E%E8%81%96%E7%B6%93%7C%E5%A4%9A%E5%9C%8B%E8%AA%9E%E8%A8%80%E8%81%96%E7%B6%93%7Callbible%7Call%20bible%7CAll%20bible%7CAll%20Bible%7C%E7%B8%BD%E5%92%8C%E8%81%96%E7%B6%93%7C%E7%B6%9C%E5%90%88%E8%81%96%E7%B6%93%7C%E7%A0%94%E7%A9%B6%E8%81%96%E7%B6%93%7C%E8%81%96%E7%B6%93%E7%A0%94%E7%A9%B6%7C%E5%A4%9A%E7%89%88%E8%81%96%E7%B6%93%7C%E5%A4%9A%E7%89%88%E6%9C%AC%E8%A8%80%E8%81%96%E7%B6%93%7CAllbible)(%5Cs%7C%E3%80%80%7C%3A%7C%3B%7C%EF%BC%9A%7C%EF%BC%9B%7C-%7C%EF%BC%8D)(%5B%5Cuff21-%5Cuff3a%5Cuff41-%5Cuff5a%5Cuff10-%5Cuff19%5Cu30a0-%5Cu30ff%5Cu3040-%5Cu309f%5Cu4e00-%5Cu9fd5_a-zA-Z0-9%5D*)%5CD*(%5B0-9.%5D%7B0%2C%7D)(%5Cs%7C%E3%80%80%7C%3A%7C%3B%7C%EF%BC%9A%7C%EF%BC%9B)%7B0%2C%7D%5CD*(%5B0-9.%5C-%EF%BC%8D%EF%BD%9E%5C~%5D%7B0%2C%7D)
	//reg := regexp.MustCompile("^(聖經|聖書|Bible|bible|ｂｉｂｌｅ|Ｂｉｂｌｅ|台語聖經巴克禮漢羅|台語聖經|閩南語聖經|台語聖經全羅|全民台語聖經全羅|台語聖經漢羅|全民台語聖經漢羅|中文聖經|中文聖經和合本修訂版|Rcuv|rcuv|ｒｃｕｖ|Ｒｃｕｖ|文言文聖經|深文理和合本|中文聖經新譯本|ncv|Ncv|Ｎｃｖ|ｎｃｖ|聖經中文譯本修訂版|tcv|TCV|Ｔｃｖ|ＴＣＶ|日文聖經|日本語聖書|JP bible|JP Bible|Jp bible|韓文聖經|KR bible|Kr Bible|Kr bible|英文聖經|英語聖書|Eng bible|ENG Bible|English bible|越南聖經|俄文聖經|多國聖經|多語聖經|多語言聖經|多國語聖經|多國語言聖經|allbible|all bible|All bible|All Bible|總和聖經|綜合聖經|研究聖經|聖經研究|多版聖經|多版本言聖經|Allbible)(\\s|　|:|;|：|；|-|－)([\uff21-\uff3a\uff41-\uff5a\uff10-\uff19\u30a0-\u30ff\u3040-\u309f\u4e00-\u9fd5_a-zA-Z0-9]*)\\D*([0-9.]{0,})(\\s|　|:|;|：|；){0,}\\D*([0-9.\\-－～\\~]{0,})")
	//2017.01.11+  https://34e.cc/552 //\u0400-\u04ff\u0500-\u052f=俄文 https://unicode-table.com/cn/blocks/cyrillic-supplement/  \u0370—\u03ff\u1f00-\u1fff\u2c80-\u2cff=希臘 \u0590-\u05ff=希伯來文 \u1100-\u11ff\u3130—\u318f\uac00-\ud7af\ua960-\ua97f=韓文  00C0-00FF=德法(http://www.programmer-club.com.tw/ShowSameTitleN/general/4309.html)
	//reg := regexp.MustCompile("^(聖經|聖書|Bible|bible|ｂｉｂｌｅ|Ｂｉｂｌｅ|多版本聖經|客家聖經|台語聖經巴克禮漢羅|台語聖經巴克禮全羅|台語聖經馬雅各全羅|台語聖經馬雅各漢羅|台語聖經|閩南語聖經|台語聖經全羅|全民台語聖經全羅|台語聖經漢羅|全民台語聖經漢羅|中文聖經|中文聖經和合本修訂版|Rcuv|rcuv|ｒｃｕｖ|Ｒｃｕｖ|文言文聖經|深文理和合本|中文聖經新譯本|ncv|Ncv|Ｎｃｖ|ｎｃｖ|聖經中文譯本修訂版|tcv|TCV|Ｔｃｖ|ＴＣＶ|日文聖經|日本語聖書|JP bible|JP Bible|Jp bible|韓文聖經|KR bible|Kr Bible|Kr bible|英文聖經|英語聖書|Kjv|kjv|Ｋｊｖ|ｋｊｖ|Eng bible|ENG Bible|English bible|BBE|Bbe|bbe|ＢＢＥ|Ｂｂｅ|ｂｂｅ|英文聖經BBE|英文聖經WEB|WEB|Web|web|ＷＥＢ|Ｗｅｂ|ｗｅｂ|英文聖經ASV|ASV|Asv|asv|ＡＳＶ|Ａｓｖ|ａｓｖ|英文聖經Darby|darby|DARBY|Ｄａｒｂｙ|ＤＡＲＢＹ|ｄａｒｂｙ|英文聖經ERV|erv|ERV|Erv|ＥＲＶ|Ｅｒｖ|ｅｒｖ|希臘聖經|lxx|LXX|Lxx|ＬＸＸ|Ｌｘｘ|ｌｘｘ|馬索拉聖經|bhs|Bhs|BHS|ＢＨＳ|Ｂｈｓ|ｂｈｓ|越南聖經|俄文聖經|多國聖經|多語聖經|多語言聖經|多國語聖經|多國語言聖經|allbible|all bible|All bible|All Bible|總和聖經|綜合聖經|研究聖經|聖經研究|多版聖經|多版本言聖經|Allbible)(\\s|　|:|;|：|；|-|－)([\u0590-\u05ff\u0370—\u03ff\u1f00-\u1fff\u2c80-\u2cff\u0400-\u04ff\u0500-\u052f\uff21-\uff3a\uff41-\uff5a\uff10-\uff19\u30a0-\u30ff\u3040-\u30ff\u31f0-\u31ff\u4e00-\u9fff\u1100-\u11ff\u3130—\u318f\uac00-\ud7af\ua960-\ua97f_a-zA-Z0-9]*)\\D*([0-9.]{0,})(\\s|　|:|;|：|；){0,}\\D*([0-9.\\-－～\\~]{0,})")
	reg := regexp.MustCompile("^(♪)")
	log.Print("--抓取分析觀察--")
	log.Print("$1 = 觸發關鍵字 = " + reg.ReplaceAllString(text, "$1"))
	// log.Print("$2 = 分割符 = " + reg.ReplaceAllString(text, "$2"))
	// log.Print("$3 = 第一主題 = " + reg.ReplaceAllString(text, "$3"))
	// log.Print("$4 = 章 = " + reg.ReplaceAllString(text, "$4"))
	// log.Print("$5 = 章節分割符 = " + reg.ReplaceAllString(text, "$5"))
	// log.Print("$6 = 節 = " + reg.ReplaceAllString(text, "$6"))
	log.Print("--抓取分析結束--")
	
	// chap_string := reg.ReplaceAllString(text, "$4")	//章
	// sec_string := reg.ReplaceAllString(text, "$6")	//節
	// sec_string = strings.Replace(sec_string,`－`, "-", -1)
	// sec_string = strings.Replace(sec_string,`～`, "-", -1)
	// sec_string = strings.Replace(sec_string,`~`, "-", -1)
	// sec_string = strings.Replace(sec_string,` ~ `, "-", -1)
	// bible_short_name := ""

	//2017.02.14+ IFTTT 也讓他執行命令
	// text = strings.Replace(text,`【IFTTT】 `, "", -1)
	// log.Print("text 執行過濾 IFTTT 結束後 = " + text)

	switch reg.ReplaceAllString(text, "$1"){
	case "♪":
		print_string = "♪"
	// case "下週預告":
	// 	print_string = "下週預告"
	// case "新眼光","每日經文":
	// 	print_string = "新眼光"
	// case "家庭禮拜":
	// 	print_string = "家庭禮拜"
	// case "行事曆":
	// 	print_string = "行事曆"
	// case "查詢可用簡寫":
	// 	print_string = "查詢可用簡寫"
	// case "圖書查詢","圖書館","館藏":
	// 	print_string = "圖書查詢"
	// case "轉傳":
	// 	print_string = "轉傳"
	// case "新約列表":
	// 	print_string = "新約列表"
	// case "舊約列表":
	// 	print_string = "舊約列表"
	// case "聚會時間":
	// 	print_string = "聚會時間"
	// case "週報","周報","最新訊息","本周資訊","本週資訊","本週週報":
	// 	print_string = "週報"
	// case "聯絡資訊","教會電話","牧師手機":
	// 	print_string = "聯絡資訊"
	// case "教會地圖","交通資訊","教會住址","單位地圖","找教會":
	// 	print_string = "地圖"
	// case "機器人88":
	// 	print_string = "機器人88"
	// case "網站資訊","教會網站","教會資訊","官方網站","臉書","FB","ＦＢ","Fb","Ｆｂ","fb","ｆｂ","FACEBOOK","ＦＡＣＥＢＯＯＫ","Facebook","Ｆａｃｅｂｏｏｋ","facebook","ｆａｃｅｂｏｏｋ","Youtube":
	// 	print_string = "網站資訊"		
	// case "主選單","選單","簡介","教學","help","Help","Ｈｅｌｐ","ｈｅｌｐ","ＨＥＬＰ","HELP":
	// 	print_string = "選單"
	// case "test","測試":
	// 	print_string = "測試"
	// case "bot","機器人","目錄","教會目錄","清單","索引","ｉｎｄｅｘ","index","Index","介紹","教會介紹","info","Info","ｉｎｆｏ":
	// 	print_string = "簡介"
	// case "開發者","admin","Admin","ａｄｍｉｎ","意見回饋":
	// 	print_string = "開發者"
	//case "多國聖經","多語聖經","多語言聖經","多國語聖經","多國語言聖經","allbible","all bible","All bible","All Bible":
	// case "模子~~~~":
	// 	log.Print(reg.ReplaceAllString(text, "$3"))
	// 	switch reg.ReplaceAllString(text, "$3") {
	// 		default:
	// 			print_string = "聖經"
	// 			//print_string = "你是要找 " +  reg.ReplaceAllString(text, "$3") + " 對嗎？\n對不起，我還沒學呢...\n"
	// 	}
	default:
		// if reply_mode!="" {
		// 	print_string = ""//"HI～\n如有其他建議或想討論，請對我輸入「開發者」進行聯絡。"
		// } else {
  //           print_string = "" //安靜模式
		// }
		print_string = ""
	}
	log.Print("Return 前的 print_string = " + print_string)
	return print_string//, chap_string, sec_string, bible_short_name
}


