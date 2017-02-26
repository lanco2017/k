package main

import (
	// "fmt"
	"log"

	// "net/http"
	// "os"
	"regexp"

	// "strconv"
	"strings"
	
	// "github.com/line/line-bot-sdk-go/linebot"

	// "bytes"

	// "io/ioutil"

	// "image/jpeg"

    // "crypto/md5"
    // "encoding/hex"

    // "encoding/json"
    // "github.com/bitly/go-simplejson"
)

func bible(text string,user_msgid string,reply_mode string) (string, string, string, string) {
	//https://gitter.im/kkdai/LineBotTemplate?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge：也可以透過 string.Contains("我要找的字", 原始字串) 來判斷
	print_string := text
	text = real_num(text)

	// bible_json_string := ""
	// bible_text_string := ""

	if GetMD5Hash(text) == "c38b3100b02ef42411a99b7975e4ff47" {
		print_string = "c38b3100b02ef42411a99b7975e4ff47"
		return print_string,"","",""
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
	reg := regexp.MustCompile("^(聖經|聖書|Bible|bible|ｂｉｂｌｅ|Ｂｉｂｌｅ|多版本聖經|客家聖經|台語聖經巴克禮漢羅|台語聖經巴克禮全羅|台語聖經馬雅各全羅|台語聖經馬雅各漢羅|台語聖經|閩南語聖經|台語聖經全羅|全民台語聖經全羅|台語聖經漢羅|全民台語聖經漢羅|中文聖經|中文聖經和合本修訂版|Rcuv|rcuv|ｒｃｕｖ|Ｒｃｕｖ|文言文聖經|深文理和合本|中文聖經新譯本|ncv|Ncv|Ｎｃｖ|ｎｃｖ|聖經中文譯本修訂版|tcv|TCV|Ｔｃｖ|ＴＣＶ|日文聖經|日本語聖書|JP bible|JP Bible|Jp bible|韓文聖經|KR bible|Kr Bible|Kr bible|英文聖經|英語聖書|Kjv|kjv|Ｋｊｖ|ｋｊｖ|Eng bible|ENG Bible|English bible|BBE|Bbe|bbe|ＢＢＥ|Ｂｂｅ|ｂｂｅ|英文聖經BBE|英文聖經WEB|WEB|Web|web|ＷＥＢ|Ｗｅｂ|ｗｅｂ|英文聖經ASV|ASV|Asv|asv|ＡＳＶ|Ａｓｖ|ａｓｖ|英文聖經Darby|darby|DARBY|Ｄａｒｂｙ|ＤＡＲＢＹ|ｄａｒｂｙ|英文聖經ERV|erv|ERV|Erv|ＥＲＶ|Ｅｒｖ|ｅｒｖ|希臘聖經|lxx|LXX|Lxx|ＬＸＸ|Ｌｘｘ|ｌｘｘ|馬索拉聖經|bhs|Bhs|BHS|ＢＨＳ|Ｂｈｓ|ｂｈｓ|越南聖經|俄文聖經|多國聖經|多語聖經|多語言聖經|多國語聖經|多國語言聖經|allbible|all bible|All bible|All Bible|總和聖經|綜合聖經|研究聖經|聖經研究|多版聖經|多版本言聖經|Allbible)(\\s|　|:|;|：|；|-|－)([\u0590-\u05ff\u0370—\u03ff\u1f00-\u1fff\u2c80-\u2cff\u0400-\u04ff\u0500-\u052f\uff21-\uff3a\uff41-\uff5a\uff10-\uff19\u30a0-\u30ff\u3040-\u30ff\u31f0-\u31ff\u4e00-\u9fff\u1100-\u11ff\u3130—\u318f\uac00-\ud7af\ua960-\ua97f_a-zA-Z0-9]*)\\D*([0-9.]{0,})(\\s|　|:|;|：|；){0,}\\D*([0-9.\\-－～\\~]{0,})")
	log.Print("--抓取分析觀察--")
	log.Print("$1 = 觸發關鍵字 = " + reg.ReplaceAllString(text, "$1"))
	log.Print("$2 = 分割符 = " + reg.ReplaceAllString(text, "$2"))
	log.Print("$3 = 第一主題 = " + reg.ReplaceAllString(text, "$3"))
	log.Print("$4 = 章 = " + reg.ReplaceAllString(text, "$4"))
	log.Print("$5 = 章節分割符 = " + reg.ReplaceAllString(text, "$5"))
	log.Print("$6 = 節 = " + reg.ReplaceAllString(text, "$6"))
	log.Print("--抓取分析結束--")
	
	chap_string := reg.ReplaceAllString(text, "$4")	//章
	sec_string := reg.ReplaceAllString(text, "$6")	//節
	sec_string = strings.Replace(sec_string,`－`, "-", -1)
	sec_string = strings.Replace(sec_string,`～`, "-", -1)
	sec_string = strings.Replace(sec_string,`~`, "-", -1)
	sec_string = strings.Replace(sec_string,` ~ `, "-", -1)
	bible_short_name := ""

	//2017.02.14+ IFTTT 也讓他執行命令
	text = strings.Replace(text,`【IFTTT】 `, "", -1)
	log.Print("text 執行過濾 IFTTT 結束後 = " + text)

	switch reg.ReplaceAllString(text, "$1"){
	case "下週預告":
		print_string = "下週預告"
	case "新眼光","每日經文":
		print_string = "新眼光"
	case "家庭禮拜":
		print_string = "家庭禮拜"
	case "行事曆":
		print_string = "行事曆"
	case "查詢可用簡寫":
		print_string = "查詢可用簡寫"
	case "圖書查詢","圖書館","館藏":
		print_string = "圖書查詢"
	case "轉傳":
		print_string = "轉傳"
	case "新約列表":
		print_string = "新約列表"
	case "舊約列表":
		print_string = "舊約列表"
	case "聚會時間":
		print_string = "聚會時間"
	case "週報","周報","最新訊息","本周資訊","本週資訊","本週週報":
		print_string = "週報"
	case "聯絡資訊","教會電話","牧師手機":
		print_string = "聯絡資訊"
	case "教會地圖","交通資訊","教會住址","單位地圖","找教會":
		print_string = "地圖"
	case "機器人88":
		print_string = "機器人88"
	case "網站資訊","教會網站","教會資訊","官方網站","臉書","FB","ＦＢ","Fb","Ｆｂ","fb","ｆｂ","FACEBOOK","ＦＡＣＥＢＯＯＫ","Facebook","Ｆａｃｅｂｏｏｋ","facebook","ｆａｃｅｂｏｏｋ","Youtube":
		print_string = "網站資訊"		
	case "主選單","選單","簡介","教學","help","Help","Ｈｅｌｐ","ｈｅｌｐ","ＨＥＬＰ","HELP":
		print_string = "選單"
	case "test","測試":
		print_string = "測試"
	case "bot","機器人","目錄","教會目錄","清單","索引","ｉｎｄｅｘ","index","Index","介紹","教會介紹","info","Info","ｉｎｆｏ":
		print_string = "簡介"
	case "開發者","admin","Admin","ａｄｍｉｎ","意見回饋":
		print_string = "開發者"
	//case "多國聖經","多語聖經","多語言聖經","多國語聖經","多國語言聖經","allbible","all bible","All bible","All Bible":
	// case "模子~~~~":
	// 	log.Print(reg.ReplaceAllString(text, "$3"))
	// 	switch reg.ReplaceAllString(text, "$3") {
	// 		default:
	// 			print_string = "聖經"
	// 			//print_string = "你是要找 " +  reg.ReplaceAllString(text, "$3") + " 對嗎？\n對不起，我還沒學呢...\n"
	// 	}
	case "總和聖經","綜合聖經","研究聖經","聖經研究","多版聖經","多版本聖經","Allbible":
		log.Print(reg.ReplaceAllString(text, "$3"))
		switch reg.ReplaceAllString(text, "$3") {
			case "Gen","Genesis","創","創世","創世紀","創世記","Ge","ge","gen","창세기","Sáng-thế Ký","Бытие":
				bible_short_name = "創"
				switch chap_string {
					case "":
						print_string = "創世記" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_var_string("創","創世記", chap_string, "1")
							default:
								print_string = Bible_print_all_var_string("創","創世記", chap_string, sec_string)
						}
				}
			case "ex","Ex","Exodus","埃及","出","出埃及","出埃及記","출애굽기","エジプト","出エジプト","出エジプト記","Xuất Ê-díp-tô Ký","Исход":
				bible_short_name = "出"
				switch chap_string {
					case "":
						print_string = "出埃及記" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_var_string("出","出埃及記", chap_string, "1")
							default:
								print_string = Bible_print_all_var_string("出","出埃及記", chap_string, sec_string)
						}
				}
			case "Lev","Leviticus","利","利未","利未記","Le","le","Левит","Lê-vi Ký","レビ記","レビ","레위기":
				bible_short_name = "利"
				switch chap_string {
					case "":
						print_string = "利未記" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_var_string("利","利未記", chap_string, "1")
							default:
								print_string = Bible_print_all_var_string("利","利未記", chap_string, sec_string)
						}
				}
			case "Num","Numbers","民","民數","民數記","Nu","nu","민수기","民数","民数記","Dân-số Ký","Числа":
				bible_short_name = "民"
				switch chap_string {
					case "":
						print_string = "民數記" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_var_string("民","民數記", chap_string, "1")
							default:
								print_string = Bible_print_all_var_string("民","民數記", chap_string, sec_string)
						}
				}
			case "Deut","Deuteronomy","申","申命","申命記","De","de","신명기","Phục-truyền Luật-lệ Ký","Второзаконие":
				bible_short_name = "申"
				switch chap_string {
					case "":
						print_string = "申命記" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_var_string("申","申命記", chap_string, "1")
							default:
								print_string = Bible_print_all_var_string("申","申命記", chap_string, sec_string)
						}
				}
			case "Josh","Joshua","約書亞","約書亞記","Jos","jos","여호수아","ヨシュア記","ヨシュア","Giô-suê","Книга Иисуса Навина":
				bible_short_name = "書"
				switch chap_string {
					case "":
						print_string = "約書亞記" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_var_string("書","約書亞記", chap_string, "1")
							default:
								print_string = Bible_print_all_var_string("書","約書亞記", chap_string, sec_string)
						}
				}
			case "Judg","Judges","士","士師","士師記","Jud","jud","jdg","Jdg","Книга Судей израилевых","Các Quan Xét","사사기":
				bible_short_name = "士"
				switch chap_string {
					case "":
						print_string = "士師記" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_var_string("士","士師記", chap_string, "1")
							default:
								print_string = Bible_print_all_var_string("士","士師記", chap_string, sec_string)
						}
				}
			case "Ruth","路得","路得記","Ru","ru","Rut","rut","룻기","ルツ","ルツ記","Ru-tơ","Книга Руфи":
				bible_short_name = "得"
				switch chap_string {
					case "":
						print_string = "路得記" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_var_string("得","路得記", chap_string, "1")
							default:
								print_string = Bible_print_all_var_string("得","路得記", chap_string, sec_string)
						}
				}
			case "1 Sam","First Samuel","撒上","撒母耳記上","1Sa","1sa","サムエル記上","サムエル上","サム上","사무엘상","1 Sa-mu-ên","Первая книга Царств":
				bible_short_name = "撒上"
				switch chap_string {
					case "":
						print_string = "撒母耳記上" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_var_string("撒上","撒母耳記上", chap_string, "1")
							default:
								print_string = Bible_print_all_var_string("撒上","撒母耳記上", chap_string, sec_string)
						}
				}
			case "2 Sam","Second Samuel","撒下","撒母耳記下","2Sa","2sa","사무엘하","サムエル記下","サムエル下","サム下","2 Sa-mu-ên","Вторая книга Царств":
				bible_short_name = "撒下"
				switch chap_string {
					case "":
						print_string = "撒母耳記下" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_var_string("撒下","撒母耳記下", chap_string, "1")
							default:
								print_string = Bible_print_all_var_string("撒下","撒母耳記下", chap_string, sec_string)
						}
				}
			case "1 Kin","First Kings","王上","列王上","列王紀上","列王記上","1Ki","1ki","열왕기상","Третья книга Царств","1 Các Vua":
				bible_short_name = "王上"
				switch chap_string {
					case "":
						print_string = "列王紀上" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_var_string("王上","列王紀上", chap_string, "1")
							default:
								print_string = Bible_print_all_var_string("王上","列王紀上", chap_string, sec_string)
						}
				}
			case "2 Kin","Second Kings","王下","列王下","列王記下","列王紀下","2Ki","2ki","열왕기하","2 Các Vua","Четвертая книга Царств":
				bible_short_name = "王下"
				switch chap_string {
					case "":
						print_string = "列王紀下" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_var_string("王下","列王紀下", chap_string, "1")
							default:
								print_string = Bible_print_all_var_string("王下","列王紀下", chap_string, sec_string)
						}
				}
			case "1 Chr","First Chronicles","歷上","代上","歷代志上","歷代上","1Ch","1ch","Первая книга Паралипоменон","1 Sử-ký","歴上","歴代上","歴代志上","역대상":
				bible_short_name = "代上"
				switch chap_string {
					case "":
						print_string = "歷代志上" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_var_string("代上","歷代志上", chap_string, "1")
							default:
								print_string = Bible_print_all_var_string("代上","歷代志上", chap_string, sec_string)
						}
				}
			case "2 Chr","Second Chronicles","代下","歷下","歷代下","歷代志下","2Ch","2ch","역대하","歴代志下","歴代下","歴下","2 Sử-ký","Вторая книга Паралипоменон":
				bible_short_name = "代下"
				switch chap_string {
					case "":
						print_string = "歷代志下" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_var_string("代下","歷代志下", chap_string, "1")
							default:
								print_string = Bible_print_all_var_string("代下","歷代志下", chap_string, sec_string)
						}
				}
			case "Ezra","拉","以斯拉","以斯拉記","Ezr","ezr","Первая книга Ездры","Ê-xơ-ra","エズラ","エズラ記","에스라":
				bible_short_name = "拉"
				switch chap_string {
					case "":
						print_string = "以斯拉記" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_var_string("拉","以斯拉記", chap_string, "1")
							default:
								print_string = Bible_print_all_var_string("拉","以斯拉記", chap_string, sec_string)
						}
				}
			case "Neh","Nehemiah","尼","尼希米","尼希米記","Ne","ne","느헤미야","ネヘミヤ書","ネヘミヤ","Nê-hê-mi","Книга Неемии":
				bible_short_name = "尼"
				switch chap_string {
					case "":
						print_string = "尼希米記" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_var_string("尼","尼希米記", chap_string, "1")
							default:
								print_string = Bible_print_all_var_string("尼","尼希米記", chap_string, sec_string)
						}
				}
			case "Esth","Esther","斯","以斯帖","以斯帖記","Es","est","Есфирь","Ê-xơ-tê","エステル","エステル記","에스더":
				bible_short_name = "斯"
				switch chap_string {
					case "":
						print_string = "以斯帖記" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_var_string("斯","以斯帖記", chap_string, "1")
							default:
								print_string = Bible_print_all_var_string("斯","以斯帖記", chap_string, sec_string)
						}
				}
			case "Job","job","伯","約伯","約伯記","Книга Иова","Gióp","ヨブ","ヨブ記","욥기":
				bible_short_name = "伯"
				switch chap_string {
					case "":
						print_string = "約伯記" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_var_string("伯","約伯記", chap_string, "1")
							default:
								print_string = Bible_print_all_var_string("伯","約伯記", chap_string, sec_string)
						}
				}
			case "Ps","Psalms","詩","詩篇","ps","시편","Thi-thiên","Псалтирь":
				bible_short_name = "詩"
				switch chap_string {
					case "":
						print_string = "詩篇" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_var_string("詩","詩篇", chap_string, "1")
							default:
								print_string = Bible_print_all_var_string("詩","詩篇", chap_string, sec_string)
						}
				}
			case "Prov","Proverbs","箴","箴言","Pr","pr","Притчи Соломона","Châm-ngôn","잠언":
				bible_short_name = "箴"
				switch chap_string {
					case "":
						print_string = "箴言" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_var_string("箴","箴言", chap_string, "1")
							default:
								print_string = Bible_print_all_var_string("箴","箴言", chap_string, sec_string)
						}
				}
			case "Eccl","Ecclesiastes","傳","傳道","傳道書","Ec","ec","Книга Екклезиаста","или Проповедника","Truyền-đạo","伝道の書","伝道","伝","伝道書","전도서":
				bible_short_name = "傳"
				switch chap_string {
					case "":
						print_string = "傳道書" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_var_string("傳","傳道書", chap_string, "1")
							default:
								print_string = Bible_print_all_var_string("傳","傳道書", chap_string, sec_string)
						}
				}
			case "Song","Song of Solomon","歌","雅歌","So","so","sng","Sng","Песнь песней Соломона","Nhã-ca","아가":
				bible_short_name = "歌"
				switch chap_string {
					case "":
						print_string = "雅歌" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_var_string("歌","雅歌", chap_string, "1")
							default:
								print_string = Bible_print_all_var_string("歌","雅歌", chap_string, sec_string)
						}
				}
			case "Is","Isaiah","賽","以賽","以賽亞","以賽亞書","Isa","isa","Книга пророка Исаии","Ê-sai","イザヤ書","イザヤ","이사야":
				bible_short_name = "賽"
				switch chap_string {
					case "":
						print_string = "以賽亞書" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_var_string("賽","以賽亞書", chap_string, "1")
							default:
								print_string = Bible_print_all_var_string("賽","以賽亞書", chap_string, sec_string)
						}
				}
			case "Jer","Jeremiah","耶","耶利米","耶利米書","jer","예레미야","エレミヤ","エレミヤ書","Giê-rê-mi","Книга пророка Иеремии":
				bible_short_name = "耶"
				switch chap_string {
					case "":
						print_string = "耶利米書" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_var_string("耶","耶利米書", chap_string, "1")
							default:
								print_string = Bible_print_all_var_string("耶","耶利米書", chap_string, sec_string)
						}
				}
			case "Lam","Lamentations","哀","哀歌","耶利米哀歌","La","lam","예레미야애가","Ca-thương","Плач Иеремии":
				bible_short_name = "哀"
				switch chap_string {
					case "":
						print_string = "耶利米哀歌" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_var_string("哀","耶利米哀歌", chap_string, "1")
							default:
								print_string = Bible_print_all_var_string("哀","耶利米哀歌", chap_string, sec_string)
						}
				}
			case "Ezek","Ezekiel","結","以西結","以西結書","Eze","eze","에스겔","エゼキエル書","エゼキエル","Ê-xê-chi-ên","Книга пророка Иезекииля":
				bible_short_name = "結"
				switch chap_string {
					case "":
						print_string = "以西結書" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_var_string("結","以西結書", chap_string, "1")
							default:
								print_string = Bible_print_all_var_string("結","以西結書", chap_string, sec_string)
						}
				}
			case "Dan","Daniel","但","但以理","但以理書","Da","da","Книга пророка Даниила","Đa-ni-ên","ダニエル書","ダニエル","다니엘":
				bible_short_name = "但"
				switch chap_string {
					case "":
						print_string = "但以理書" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_var_string("但","但以理書", chap_string, "1")
							default:
								print_string = Bible_print_all_var_string("但","但以理書", chap_string, sec_string)
						}
				}
			case "Hos","Hosea","何","何西","何西阿","何西阿書","Ho","ho","Книга пророка Осии","Ô-sê","ホセア書","ホセア","호세아":
				bible_short_name = "何"
				switch chap_string {
					case "":
						print_string = "何西阿書" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_var_string("何","何西阿書", chap_string, "1")
							default:
								print_string = Bible_print_all_var_string("何","何西阿書", chap_string, sec_string)
						}
				}
			case "Joel","珥","約珥","約珥書","Joe","joe","Книга пророка Иоиля","Giô-ên","ヨエル書","ヨエル","요엘":
				bible_short_name = "珥"
				switch chap_string {
					case "":
						print_string = "約珥書" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_var_string("珥","約珥書", chap_string, "1")
							default:
								print_string = Bible_print_all_var_string("珥","約珥書", chap_string, sec_string)
						}
				}
			case "Amos","摩","阿摩司書","Am","am","Книга пророка Амоса","A-mốt","アモス書","アモス","아모스":
				bible_short_name = "摩"
				switch chap_string {
					case "":
						print_string = "阿摩司書" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_var_string("摩","阿摩司書", chap_string, "1")
							default:
								print_string = Bible_print_all_var_string("摩","阿摩司書", chap_string, sec_string)
						}
				}
			case "Obad","Obadiah","俄","俄巴底亞","俄巴底亞書","Ob","ob","오바댜","オバデヤ書","オバデヤ","Áp-đia","Книга пророка Авдия":
				bible_short_name = "俄"
				switch chap_string {
					case "":
						print_string = "俄巴底亞書" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_var_string("俄","俄巴底亞書", chap_string, "1")
							default:
								print_string = Bible_print_all_var_string("俄","俄巴底亞書", chap_string, sec_string)
						}
				}
			case "Jon","Jonah","拿","約拿","約拿書","jon","요나","ヨナ書","ヨナ","Giô-na","Книга пророка Ионы":
				bible_short_name = "拿"
				switch chap_string {
					case "":
						print_string = "約拿書" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_var_string("拿","約拿書", chap_string, "1")
							default:
								print_string = Bible_print_all_var_string("拿","約拿書", chap_string, sec_string)
						}
				}
			case "Micah","彌","彌迦","彌迦書","Mic","mic","Книга пророка Михея","Mi-chê","ミカ書","ミカ","미가":
				bible_short_name = "彌"
				switch chap_string {
					case "":
						print_string = "彌迦書" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_var_string("彌","彌迦書", chap_string, "1")
							default:
								print_string = Bible_print_all_var_string("彌","彌迦書", chap_string, sec_string)
						}
				}
			case "Nah","Nahum","鴻","那鴻","那鴻書","Na","na","Книга пророка Наума","Na-hum","ナホム書","ナホム","나훔":
				bible_short_name = "鴻"
				switch chap_string {
					case "":
						print_string = "那鴻書" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_var_string("鴻","那鴻書", chap_string, "1")
							default:
								print_string = Bible_print_all_var_string("鴻","那鴻書", chap_string, sec_string)
						}
				}
			case "Habakkuk","哈","哈巴","哈巴谷","哈巴谷書","Hab","hab","Книга пророка Аввакума","Ha-ba-cúc","ハバクク書","ハバクク","ハバ","クク","ハバ書","하박국":
				bible_short_name = "哈"
				switch chap_string {
					case "":
						print_string = "哈巴谷書" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_var_string("哈","哈巴谷書", chap_string, "1")
							default:
								print_string = Bible_print_all_var_string("哈","哈巴谷書", chap_string, sec_string)
						}
				}
			case "Zeph","Zephaniah","番","西番雅","西番雅書","Zep","zep","스바냐","ゼパニヤ書","ゼパニヤ","Sô-phô-ni","Книга пророка Софонии":
				bible_short_name = "番"
				switch chap_string {
					case "":
						print_string = "西番雅書" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_var_string("番","西番雅書", chap_string, "1")
							default:
								print_string = Bible_print_all_var_string("番","西番雅書", chap_string, sec_string)
						}
				}
			case "Haggai","該","哈該","哈該書","Hag","hag","학개","ハガイ書","ハガイ","A-ghê","Книга пророка Аггея":
				bible_short_name = "該"
				switch chap_string {
					case "":
						print_string = "哈該書" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_var_string("該","哈該書", chap_string, "1")
							default:
								print_string = Bible_print_all_var_string("該","哈該書", chap_string, sec_string)
						}
				}
			case "Zech","Zechariah","亞","撒迦利亞","撒迦利亞書","Zec","zec","Книга пророка Захарии","Xa-cha-ri","스가랴","ゼカリヤ書","ゼカリヤ":
				bible_short_name = "亞"
				switch chap_string {
					case "":
						print_string = "撒迦利亞書" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_var_string("亞","撒迦利亞書", chap_string, "1")
							default:
								print_string = Bible_print_all_var_string("亞","撒迦利亞書", chap_string, sec_string)
						}
				}
			case "Malachi","瑪","瑪拉","瑪拉基","瑪拉基書","Mal","mal","말라기","マラキ書","マラキ","Ma-la-chi","Книга пророка Малахии":
				bible_short_name = "瑪"
				switch chap_string {
					case "":
						print_string = "瑪拉基書" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_var_string("瑪","瑪拉基書", chap_string, "1")
							default:
								print_string = Bible_print_all_var_string("瑪","瑪拉基書", chap_string, sec_string)
						}
				}
			case "Matt","Matthew","太","馬太","馬太福音","Mt","mt","마태복음","マタイによる福音書","マタイ","マタイによる","Ma-thi-ơ","От Матфея святое благовествование":
				bible_short_name = "太"
				switch chap_string {
					case "":
						print_string = "馬太福音" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_var_string("太","馬太福音", chap_string, "1")
							default:
								print_string = Bible_print_all_var_string("太","馬太福音", chap_string, sec_string)
						}
				}
			case "Mark","可","馬可","馬可福音","Mr","mr","マルコによる福音書","マルコ","マルコによる","마가복음","Mác","От Марка святое благовествование":
				bible_short_name = "可"
				switch chap_string {
					case "":
						print_string = "馬可福音" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_var_string("可","馬可福音", chap_string, "1")
							default:
								print_string = Bible_print_all_var_string("可","馬可福音", chap_string, sec_string)
						}
				}
			case "Luke","路","路加","路加福音","Lu","lu","От Луки святое благовествование","Lu-ca","ルカによる福音書","ルカ","ルカによる","누가복음":
				bible_short_name = "路"
				switch chap_string {
					case "":
						print_string = "路加福音" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_var_string("路","路加福音", chap_string, "1")
							default:
								print_string = Bible_print_all_var_string("路","路加福音", chap_string, sec_string)
						}
				}
			case "John","約","約翰","約翰福音","Joh","joh","От Иоанна святое благовествование","Giăng","ヨハネによる福音書","ヨハネ","ヨハネによる","요한복음":
				bible_short_name = "約"
				switch chap_string {
					case "":
						print_string = "約翰福音" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_var_string("約","約翰福音", chap_string, "1")
							default:
								print_string = Bible_print_all_var_string("約","約翰福音", chap_string, sec_string)
						}
				}
			case "Acts","徒","使徒","使徒行傳","Ac","ac","Деяния святых апостолов","Công-vụ Các Sứ-đồ","使徒行伝","사도행전":
				bible_short_name = "徒"
				switch chap_string {
					case "":
						print_string = "使徒行傳" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_var_string("徒","使徒行傳", chap_string, "1")
							default:
								print_string = Bible_print_all_var_string("徒","使徒行傳", chap_string, sec_string)
						}
				}
			case "Rom","Romans","羅","羅馬","羅馬書","Ro","ro","Послание к Римлянам","Rô-ma","ローマ","ローマ人への手紙","로마서":
				bible_short_name = "羅"
				switch chap_string {
					case "":
						print_string = "羅馬書" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_var_string("羅","羅馬書", chap_string, "1")
							default:
								print_string = Bible_print_all_var_string("羅","羅馬書", chap_string, sec_string)
						}
				}
			case "1 Cor","First Corinthians","林前","哥林多前","哥林多前書","1Co","1co","Первое послание к Коринфянам","1 Cô-rinh-tô","コリント人への第一の手紙","コリント一","コリント人への第一","고린도전서":
				bible_short_name = "林前"
				switch chap_string {
					case "":
						print_string = "哥林多前書" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_var_string("林前","哥林多前書", chap_string, "1")
							default:
								print_string = Bible_print_all_var_string("林前","哥林多前書", chap_string, sec_string)
						}
				}
			case "2 Cor","Second Corinthians","林後","哥林多後","哥林多後書","2Co","2co","Второе послание к Коринфянам","2 Cô-rinh-tô","コリント人への第二の手紙","コリント二","コリント人への第二の","고린도후서":
				bible_short_name = "林後"
				switch chap_string {
					case "":
						print_string = "哥林多後書" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_var_string("林後","哥林多後書", chap_string, "1")
							default:
								print_string = Bible_print_all_var_string("林後","哥林多後書", chap_string, sec_string)
						}
				}
			case "Gal","Galatians","加","加拉太","加拉太書","Ga","ga","Послание к Галатам","Ga-la-ti","ガラテヤ","ガラテヤ人への手紙","갈라디아서":
				bible_short_name = "加"
				switch chap_string {
					case "":
						print_string = "加拉太書" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_var_string("加","加拉太書", chap_string, "1")
							default:
								print_string = Bible_print_all_var_string("加","加拉太書", chap_string, sec_string)
						}
				}
			case "Ephesians","弗","以弗所","以弗所書","Eph","eph","Послание к Ефесянам","Ê-phê-sô","エペソ人への手紙","エペソ","エペソ人","エペソ人の手紙","에베소서":
				bible_short_name = "弗"
				switch chap_string {
					case "":
						print_string = "以弗所書" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_var_string("弗","以弗所書", chap_string, "1")
							default:
								print_string = Bible_print_all_var_string("弗","以弗所書", chap_string, sec_string)
						}
				}
			case "Phil","Philippians","腓","腓立","腓立比","腓立比書","Php","php","빌립보서","ピリピ","ピリピ人.ピリピ人への手紙","Послание к Филиппийцам","Phi-líp":
				bible_short_name = "腓"
				switch chap_string {
					case "":
						print_string = "腓立比書" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_var_string("腓","腓立比書", chap_string, "1")
							default:
								print_string = Bible_print_all_var_string("腓","腓立比書", chap_string, sec_string)
						}
				}
			case "Col","col","Colossians","西","歌羅西","歌羅","歌羅西書","Послание к Колоссянам","Cô-lô-se","コロサイ人への手紙","コロサイ","コロ","골로새서":
				bible_short_name = "西"
				switch chap_string {
					case "":
						print_string = "歌羅西書" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_var_string("西","歌羅西書", chap_string, "1")
							default:
								print_string = Bible_print_all_var_string("西","歌羅西書", chap_string, sec_string)
						}
				}
			case "1 Thess","First Thessalonians","帖前","帖撒羅尼迦前","帖撒羅尼迦前書","1Th","1th","데살로니가전서","テサロニケ人への第一の手紙","テサ一","テサロニケ一","1 Tê-sa-lô-ni-ca","Первое послание к Фессалоникийцам (Солунянам)":
				bible_short_name = "帖前"
				switch chap_string {
					case "":
						print_string = "帖撒羅尼迦前書" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_var_string("帖前","帖撒羅尼迦前書", chap_string, "1")
							default:
								print_string = Bible_print_all_var_string("帖前","帖撒羅尼迦前書", chap_string, sec_string)
						}
				}
			case "2 Thess","Second Thessalonians","帖後","帖撒羅尼迦後","帖撒羅尼迦後書","2Th","2th","데살로니가후서","テサロニケ人への第二の手紙","テサ二","テサロニケ二","2 Tê-sa-lô-ni-ca","Второе послание к Фессалоникийцам (Солунянам)":
				bible_short_name = "帖後"
				switch chap_string {
					case "":
						print_string = "帖撒羅尼迦後書" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_var_string("帖後","帖撒羅尼迦後書", chap_string, "1")
							default:
								print_string = Bible_print_all_var_string("帖後","帖撒羅尼迦後書", chap_string, sec_string)
						}
				}
			case "1 Tim","First Timothy","提前","提摩太前","提摩太前書","1Ti","1ti","Первое послание к Тимофею","1 Ti-mô-thê","テモテヘの第一の手紙","テモテ一","디모데전서":
				bible_short_name = "提前"
				switch chap_string {
					case "":
						print_string = "提摩太前書" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_var_string("提前","提摩太前書", chap_string, "1")
							default:
								print_string = Bible_print_all_var_string("提前","提摩太前書", chap_string, sec_string)
						}
				}
			case "2 Tim","Second Timothy","提後","提摩太後","提摩太後書","2Ti","2ti","Второе послание к Тимофею","2 Ti-mô-thê","テモテヘの第二の手紙","テモテ二","디모데후서":
				bible_short_name = "提後"
				switch chap_string {
					case "":
						print_string = "提摩太後書" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_var_string("提後","提摩太後書", chap_string, "1")
							default:
								print_string = Bible_print_all_var_string("提後","提摩太後書", chap_string, sec_string)
						}
				}
			case "Titus","多","提多","提多書","Tit","tit","Послание к Титу","Tít","テトスヘの手紙","テトス","디도서":
				bible_short_name = "多"
				switch chap_string {
					case "":
						print_string = "提多書" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_var_string("多","提多書", chap_string, "1")
							default:
								print_string = Bible_print_all_var_string("多","提多書", chap_string, sec_string)
						}
				}
			case "Philem","Philemon","門","腓利","腓利門","腓利門書","Phm","phm","Послание к Филимону","Phi-lê-môn","ピレモンヘの手紙","ピレモン","빌레몬서":
				bible_short_name = "門"
				switch chap_string {
					case "":
						print_string = "腓利門書" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_var_string("門","腓利門書", chap_string, "1")
							default:
								print_string = Bible_print_all_var_string("門","腓利門書", chap_string, sec_string)
						}
				}
			case "Heb","Hebrews","來","希伯來","希伯來書","heb","Послание к Евреям","Hê-bơ-rơ","ヘブル人への手紙","ヘブル","히브리서":
				bible_short_name = "來"
				switch chap_string {
					case "":
						print_string = "希伯來書" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_var_string("來","希伯來書", chap_string, "1")
							default:
								print_string = Bible_print_all_var_string("來","希伯來書", chap_string, sec_string)
						}
				}
			case "James","雅","雅各","雅各書","Jas","jas","Послание Иакова","Gia-cơ","ヤコブの手紙","ヤコブ","야고보서":
				bible_short_name = "雅"
				switch chap_string {
					case "":
						print_string = "雅各書" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_var_string("雅","雅各書", chap_string, "1")
							default:
								print_string = Bible_print_all_var_string("雅","雅各書", chap_string, sec_string)
						}
				}
			case "1 Pet","First Peter","彼前","彼得前","彼得前書","1Pe","1pe","Первое послание Петра","1 Phi-e-rơ","ペテロの第一の手紙","ペテロ一","베드로전서":
				bible_short_name = "彼前"
				switch chap_string {
					case "":
						print_string = "彼得前書" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_var_string("彼前","彼得前書", chap_string, "1")
							default:
								print_string = Bible_print_all_var_string("彼前","彼得前書", chap_string, sec_string)
						}
				}
			case "2 Pet","Second Peter","彼後","彼得後","彼得後書","2Pe","2pe","Второе послание Петра","2 Phi-e-rơ","ペテロの第二の手紙","ペテロ","베드로후서":
				bible_short_name = "彼後"
				switch chap_string {
					case "":
						print_string = "彼得後書" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_var_string("彼後","彼得後書", chap_string, "1")
							default:
								print_string = Bible_print_all_var_string("彼後","彼得後書", chap_string, sec_string)
						}
				}
			case "1 John","First John","約一","約翰一書","約翰1","約翰1書","1Jo","1jo","Первое послание Иоанна","1 Giăng","ヨハネの第一の手紙","ヨハネ一","요한일서":
				bible_short_name = "約一"
				switch chap_string {
					case "":
						print_string = "約翰一書" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_var_string("約一","約翰一書", chap_string, "1")
							default:
								print_string = Bible_print_all_var_string("約一","約翰一書", chap_string, sec_string)
						}
				}
			case "2 John","second John","約二","約翰二書","約翰2","約翰2書","2Jo","Второе послание Иоанна","2 Giăng","ヨハネの第二の手紙","ヨハネ二","요한2서":
				bible_short_name = "約二"
				switch chap_string {
					case "":
						print_string = "約翰二書" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_var_string("約二","約翰二書", chap_string, "1")
							default:
								print_string = Bible_print_all_var_string("約二","約翰二書", chap_string, sec_string)
						}
				}
			case "3 John","Third John","約三","約翰三書","約翰3","約翰3書","3Jo","3jo","Третье послание Иоанна","3 Giăng","ヨハネの第三の手紙","ヨハネ三","요한3서":
				bible_short_name = "約三"
				switch chap_string {
					case "":
						print_string = "約翰三書" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_var_string("約三","約翰三書", chap_string, "1")
							default:
								print_string = Bible_print_all_var_string("約三","約翰三書", chap_string, sec_string)
						}
				}
			case "Jude","猶","猶大","猶大書","jude","Послание Иуды","Giu-đe","ユダの手紙","ユダ","유다서":
				bible_short_name = "猶"
				switch chap_string {
					case "":
						print_string = "猶大書" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_var_string("猶","猶大書", chap_string, "1")
							default:
								print_string = Bible_print_all_var_string("猶","猶大書", chap_string, sec_string)
						}
				}
			case "Rev","Revelation","啟","啟示","啟示錄","Re","re","ｒｅ","Ｒｅ","rev","Откровение ап. Иоанна Богослова (Апокалипсис)","Khải-huyền","ヨハネの黙示録","黙示録","요한계시록":
				bible_short_name = "啟"
				switch chap_string {
					case "":
						print_string = "啟示錄" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_var_string("啟","啟示錄", chap_string, "1")
							default:
								print_string = Bible_print_all_var_string("啟","啟示錄", chap_string, sec_string)
						}
				}
			default:
				print_string = "聖經"
				//print_string = "你是要找 " +  reg.ReplaceAllString(text, "$3") + " 對嗎？\n對不起，我還沒學呢...\n"
		}
	case "多國聖經","多語聖經","多語言聖經","多國語聖經","多國語言聖經","allbible","all bible","All bible","All Bible":
		log.Print(reg.ReplaceAllString(text, "$3"))
		switch reg.ReplaceAllString(text, "$3") {
			case "Gen","Genesis","創","創世","創世紀","創世記","Ge","ge","gen","창세기","Sáng-thế Ký","Бытие":
				bible_short_name = "創"
				switch chap_string {
					case "":
						print_string = "創世記" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_string("創","創世記", chap_string, "1")
							default:
								print_string = Bible_print_all_string("創","創世記", chap_string, sec_string)
						}
				}
			case "ex","Ex","Exodus","埃及","出","出埃及","出埃及記","출애굽기","エジプト","出エジプト","出エジプト記","Xuất Ê-díp-tô Ký","Исход":
				bible_short_name = "出"
				switch chap_string {
					case "":
						print_string = "出埃及記" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_string("出","出埃及記", chap_string, "1")
							default:
								print_string = Bible_print_all_string("出","出埃及記", chap_string, sec_string)
						}
				}
			case "Lev","Leviticus","利","利未","利未記","Le","le","Левит","Lê-vi Ký","レビ記","レビ","레위기":
				bible_short_name = "利"
				switch chap_string {
					case "":
						print_string = "利未記" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_string("利","利未記", chap_string, "1")
							default:
								print_string = Bible_print_all_string("利","利未記", chap_string, sec_string)
						}
				}
			case "Num","Numbers","民","民數","民數記","Nu","nu","민수기","民数","民数記","Dân-số Ký","Числа":
				bible_short_name = "民"
				switch chap_string {
					case "":
						print_string = "民數記" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_string("民","民數記", chap_string, "1")
							default:
								print_string = Bible_print_all_string("民","民數記", chap_string, sec_string)
						}
				}
			case "Deut","Deuteronomy","申","申命","申命記","De","de","신명기","Phục-truyền Luật-lệ Ký","Второзаконие":
				bible_short_name = "申"
				switch chap_string {
					case "":
						print_string = "申命記" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_string("申","申命記", chap_string, "1")
							default:
								print_string = Bible_print_all_string("申","申命記", chap_string, sec_string)
						}
				}
			case "Josh","Joshua","約書亞","約書亞記","Jos","jos","여호수아","ヨシュア記","ヨシュア","Giô-suê","Книга Иисуса Навина":
				bible_short_name = "書"
				switch chap_string {
					case "":
						print_string = "約書亞記" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_string("書","約書亞記", chap_string, "1")
							default:
								print_string = Bible_print_all_string("書","約書亞記", chap_string, sec_string)
						}
				}
			case "Judg","Judges","士","士師","士師記","Jud","jud","jdg","Jdg","Книга Судей израилевых","Các Quan Xét","사사기":
				bible_short_name = "士"
				switch chap_string {
					case "":
						print_string = "士師記" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_string("士","士師記", chap_string, "1")
							default:
								print_string = Bible_print_all_string("士","士師記", chap_string, sec_string)
						}
				}
			case "Ruth","路得","路得記","Ru","ru","Rut","rut","룻기","ルツ","ルツ記","Ru-tơ","Книга Руфи":
				bible_short_name = "得"
				switch chap_string {
					case "":
						print_string = "路得記" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_string("得","路得記", chap_string, "1")
							default:
								print_string = Bible_print_all_string("得","路得記", chap_string, sec_string)
						}
				}
			case "1 Sam","First Samuel","撒上","撒母耳記上","1Sa","1sa","サムエル記上","サムエル上","サム上","사무엘상","1 Sa-mu-ên","Первая книга Царств":
				bible_short_name = "撒上"
				switch chap_string {
					case "":
						print_string = "撒母耳記上" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_string("撒上","撒母耳記上", chap_string, "1")
							default:
								print_string = Bible_print_all_string("撒上","撒母耳記上", chap_string, sec_string)
						}
				}
			case "2 Sam","Second Samuel","撒下","撒母耳記下","2Sa","2sa","사무엘하","サムエル記下","サムエル下","サム下","2 Sa-mu-ên","Вторая книга Царств":
				bible_short_name = "撒下"
				switch chap_string {
					case "":
						print_string = "撒母耳記下" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_string("撒下","撒母耳記下", chap_string, "1")
							default:
								print_string = Bible_print_all_string("撒下","撒母耳記下", chap_string, sec_string)
						}
				}
			case "1 Kin","First Kings","王上","列王上","列王紀上","列王記上","1Ki","1ki","열왕기상","Третья книга Царств","1 Các Vua":
				bible_short_name = "王上"
				switch chap_string {
					case "":
						print_string = "列王紀上" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_string("王上","列王紀上", chap_string, "1")
							default:
								print_string = Bible_print_all_string("王上","列王紀上", chap_string, sec_string)
						}
				}
			case "2 Kin","Second Kings","王下","列王下","列王記下","列王紀下","2Ki","2ki","열왕기하","2 Các Vua","Четвертая книга Царств":
				bible_short_name = "王下"
				switch chap_string {
					case "":
						print_string = "列王紀下" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_string("王下","列王紀下", chap_string, "1")
							default:
								print_string = Bible_print_all_string("王下","列王紀下", chap_string, sec_string)
						}
				}
			case "1 Chr","First Chronicles","歷上","代上","歷代志上","歷代上","1Ch","1ch","Первая книга Паралипоменон","1 Sử-ký","歴上","歴代上","歴代志上","역대상":
				bible_short_name = "代上"
				switch chap_string {
					case "":
						print_string = "歷代志上" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_string("代上","歷代志上", chap_string, "1")
							default:
								print_string = Bible_print_all_string("代上","歷代志上", chap_string, sec_string)
						}
				}
			case "2 Chr","Second Chronicles","代下","歷下","歷代下","歷代志下","2Ch","2ch","역대하","歴代志下","歴代下","歴下","2 Sử-ký","Вторая книга Паралипоменон":
				bible_short_name = "代下"
				switch chap_string {
					case "":
						print_string = "歷代志下" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_string("代下","歷代志下", chap_string, "1")
							default:
								print_string = Bible_print_all_string("代下","歷代志下", chap_string, sec_string)
						}
				}
			case "Ezra","拉","以斯拉","以斯拉記","Ezr","ezr","Первая книга Ездры","Ê-xơ-ra","エズラ","エズラ記","에스라":
				bible_short_name = "拉"
				switch chap_string {
					case "":
						print_string = "以斯拉記" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_string("拉","以斯拉記", chap_string, "1")
							default:
								print_string = Bible_print_all_string("拉","以斯拉記", chap_string, sec_string)
						}
				}
			case "Neh","Nehemiah","尼","尼希米","尼希米記","Ne","ne","느헤미야","ネヘミヤ書","ネヘミヤ","Nê-hê-mi","Книга Неемии":
				bible_short_name = "尼"
				switch chap_string {
					case "":
						print_string = "尼希米記" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_string("尼","尼希米記", chap_string, "1")
							default:
								print_string = Bible_print_all_string("尼","尼希米記", chap_string, sec_string)
						}
				}
			case "Esth","Esther","斯","以斯帖","以斯帖記","Es","est","Есфирь","Ê-xơ-tê","エステル","エステル記","에스더":
				bible_short_name = "斯"
				switch chap_string {
					case "":
						print_string = "以斯帖記" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_string("斯","以斯帖記", chap_string, "1")
							default:
								print_string = Bible_print_all_string("斯","以斯帖記", chap_string, sec_string)
						}
				}
			case "Job","job","伯","約伯","約伯記","Книга Иова","Gióp","ヨブ","ヨブ記","욥기":
				bible_short_name = "伯"
				switch chap_string {
					case "":
						print_string = "約伯記" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_string("伯","約伯記", chap_string, "1")
							default:
								print_string = Bible_print_all_string("伯","約伯記", chap_string, sec_string)
						}
				}
			case "Ps","Psalms","詩","詩篇","ps","시편","Thi-thiên","Псалтирь":
				bible_short_name = "詩"
				switch chap_string {
					case "":
						print_string = "詩篇" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_string("詩","詩篇", chap_string, "1")
							default:
								print_string = Bible_print_all_string("詩","詩篇", chap_string, sec_string)
						}
				}
			case "Prov","Proverbs","箴","箴言","Pr","pr","Притчи Соломона","Châm-ngôn","잠언":
				bible_short_name = "箴"
				switch chap_string {
					case "":
						print_string = "箴言" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_string("箴","箴言", chap_string, "1")
							default:
								print_string = Bible_print_all_string("箴","箴言", chap_string, sec_string)
						}
				}
			case "Eccl","Ecclesiastes","傳","傳道","傳道書","Ec","ec","Книга Екклезиаста","или Проповедника","Truyền-đạo","伝道の書","伝道","伝","伝道書","전도서":
				bible_short_name = "傳"
				switch chap_string {
					case "":
						print_string = "傳道書" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_string("傳","傳道書", chap_string, "1")
							default:
								print_string = Bible_print_all_string("傳","傳道書", chap_string, sec_string)
						}
				}
			case "Song","Song of Solomon","歌","雅歌","So","so","sng","Sng","Песнь песней Соломона","Nhã-ca","아가":
				bible_short_name = "歌"
				switch chap_string {
					case "":
						print_string = "雅歌" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_string("歌","雅歌", chap_string, "1")
							default:
								print_string = Bible_print_all_string("歌","雅歌", chap_string, sec_string)
						}
				}
			case "Is","Isaiah","賽","以賽","以賽亞","以賽亞書","Isa","isa","Книга пророка Исаии","Ê-sai","イザヤ書","イザヤ","이사야":
				bible_short_name = "賽"
				switch chap_string {
					case "":
						print_string = "以賽亞書" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_string("賽","以賽亞書", chap_string, "1")
							default:
								print_string = Bible_print_all_string("賽","以賽亞書", chap_string, sec_string)
						}
				}
			case "Jer","Jeremiah","耶","耶利米","耶利米書","jer","예레미야","エレミヤ","エレミヤ書","Giê-rê-mi","Книга пророка Иеремии":
				bible_short_name = "耶"
				switch chap_string {
					case "":
						print_string = "耶利米書" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_string("耶","耶利米書", chap_string, "1")
							default:
								print_string = Bible_print_all_string("耶","耶利米書", chap_string, sec_string)
						}
				}
			case "Lam","Lamentations","哀","哀歌","耶利米哀歌","La","lam","예레미야애가","Ca-thương","Плач Иеремии":
				bible_short_name = "哀"
				switch chap_string {
					case "":
						print_string = "耶利米哀歌" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_string("哀","耶利米哀歌", chap_string, "1")
							default:
								print_string = Bible_print_all_string("哀","耶利米哀歌", chap_string, sec_string)
						}
				}
			case "Ezek","Ezekiel","結","以西結","以西結書","Eze","eze","에스겔","エゼキエル書","エゼキエル","Ê-xê-chi-ên","Книга пророка Иезекииля":
				bible_short_name = "結"
				switch chap_string {
					case "":
						print_string = "以西結書" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_string("結","以西結書", chap_string, "1")
							default:
								print_string = Bible_print_all_string("結","以西結書", chap_string, sec_string)
						}
				}
			case "Dan","Daniel","但","但以理","但以理書","Da","da","Книга пророка Даниила","Đa-ni-ên","ダニエル書","ダニエル","다니엘":
				bible_short_name = "但"
				switch chap_string {
					case "":
						print_string = "但以理書" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_string("但","但以理書", chap_string, "1")
							default:
								print_string = Bible_print_all_string("但","但以理書", chap_string, sec_string)
						}
				}
			case "Hos","Hosea","何","何西","何西阿","何西阿書","Ho","ho","Книга пророка Осии","Ô-sê","ホセア書","ホセア","호세아":
				bible_short_name = "何"
				switch chap_string {
					case "":
						print_string = "何西阿書" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_string("何","何西阿書", chap_string, "1")
							default:
								print_string = Bible_print_all_string("何","何西阿書", chap_string, sec_string)
						}
				}
			case "Joel","珥","約珥","約珥書","Joe","joe","Книга пророка Иоиля","Giô-ên","ヨエル書","ヨエル","요엘":
				bible_short_name = "珥"
				switch chap_string {
					case "":
						print_string = "約珥書" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_string("珥","約珥書", chap_string, "1")
							default:
								print_string = Bible_print_all_string("珥","約珥書", chap_string, sec_string)
						}
				}
			case "Amos","摩","阿摩司書","Am","am","Книга пророка Амоса","A-mốt","アモス書","アモス","아모스":
				bible_short_name = "摩"
				switch chap_string {
					case "":
						print_string = "阿摩司書" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_string("摩","阿摩司書", chap_string, "1")
							default:
								print_string = Bible_print_all_string("摩","阿摩司書", chap_string, sec_string)
						}
				}
			case "Obad","Obadiah","俄","俄巴底亞","俄巴底亞書","Ob","ob","오바댜","オバデヤ書","オバデヤ","Áp-đia","Книга пророка Авдия":
				bible_short_name = "俄"
				switch chap_string {
					case "":
						print_string = "俄巴底亞書" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_string("俄","俄巴底亞書", chap_string, "1")
							default:
								print_string = Bible_print_all_string("俄","俄巴底亞書", chap_string, sec_string)
						}
				}
			case "Jon","Jonah","拿","約拿","約拿書","jon","요나","ヨナ書","ヨナ","Giô-na","Книга пророка Ионы":
				bible_short_name = "拿"
				switch chap_string {
					case "":
						print_string = "約拿書" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_string("拿","約拿書", chap_string, "1")
							default:
								print_string = Bible_print_all_string("拿","約拿書", chap_string, sec_string)
						}
				}
			case "Micah","彌","彌迦","彌迦書","Mic","mic","Книга пророка Михея","Mi-chê","ミカ書","ミカ","미가":
				bible_short_name = "彌"
				switch chap_string {
					case "":
						print_string = "彌迦書" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_string("彌","彌迦書", chap_string, "1")
							default:
								print_string = Bible_print_all_string("彌","彌迦書", chap_string, sec_string)
						}
				}
			case "Nah","Nahum","鴻","那鴻","那鴻書","Na","na","Книга пророка Наума","Na-hum","ナホム書","ナホム","나훔":
				bible_short_name = "鴻"
				switch chap_string {
					case "":
						print_string = "那鴻書" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_string("鴻","那鴻書", chap_string, "1")
							default:
								print_string = Bible_print_all_string("鴻","那鴻書", chap_string, sec_string)
						}
				}
			case "Habakkuk","哈","哈巴","哈巴谷","哈巴谷書","Hab","hab","Книга пророка Аввакума","Ha-ba-cúc","ハバクク書","ハバクク","ハバ","クク","ハバ書","하박국":
				bible_short_name = "哈"
				switch chap_string {
					case "":
						print_string = "哈巴谷書" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_string("哈","哈巴谷書", chap_string, "1")
							default:
								print_string = Bible_print_all_string("哈","哈巴谷書", chap_string, sec_string)
						}
				}
			case "Zeph","Zephaniah","番","西番雅","西番雅書","Zep","zep","스바냐","ゼパニヤ書","ゼパニヤ","Sô-phô-ni","Книга пророка Софонии":
				bible_short_name = "番"
				switch chap_string {
					case "":
						print_string = "西番雅書" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_string("番","西番雅書", chap_string, "1")
							default:
								print_string = Bible_print_all_string("番","西番雅書", chap_string, sec_string)
						}
				}
			case "Haggai","該","哈該","哈該書","Hag","hag","학개","ハガイ書","ハガイ","A-ghê","Книга пророка Аггея":
				bible_short_name = "該"
				switch chap_string {
					case "":
						print_string = "哈該書" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_string("該","哈該書", chap_string, "1")
							default:
								print_string = Bible_print_all_string("該","哈該書", chap_string, sec_string)
						}
				}
			case "Zech","Zechariah","亞","撒迦利亞","撒迦利亞書","Zec","zec","Книга пророка Захарии","Xa-cha-ri","스가랴","ゼカリヤ書","ゼカリヤ":
				bible_short_name = "亞"
				switch chap_string {
					case "":
						print_string = "撒迦利亞書" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_string("亞","撒迦利亞書", chap_string, "1")
							default:
								print_string = Bible_print_all_string("亞","撒迦利亞書", chap_string, sec_string)
						}
				}
			case "Malachi","瑪","瑪拉","瑪拉基","瑪拉基書","Mal","mal","말라기","マラキ書","マラキ","Ma-la-chi","Книга пророка Малахии":
				bible_short_name = "瑪"
				switch chap_string {
					case "":
						print_string = "瑪拉基書" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_string("瑪","瑪拉基書", chap_string, "1")
							default:
								print_string = Bible_print_all_string("瑪","瑪拉基書", chap_string, sec_string)
						}
				}
			case "Matt","Matthew","太","馬太","馬太福音","Mt","mt","마태복음","マタイによる福音書","マタイ","マタイによる","Ma-thi-ơ","От Матфея святое благовествование":
				bible_short_name = "太"
				switch chap_string {
					case "":
						print_string = "馬太福音" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_string("太","馬太福音", chap_string, "1")
							default:
								print_string = Bible_print_all_string("太","馬太福音", chap_string, sec_string)
						}
				}
			case "Mark","可","馬可","馬可福音","Mr","mr","マルコによる福音書","マルコ","マルコによる","마가복음","Mác","От Марка святое благовествование":
				bible_short_name = "可"
				switch chap_string {
					case "":
						print_string = "馬可福音" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_string("可","馬可福音", chap_string, "1")
							default:
								print_string = Bible_print_all_string("可","馬可福音", chap_string, sec_string)
						}
				}
			case "Luke","路","路加","路加福音","Lu","lu","От Луки святое благовествование","Lu-ca","ルカによる福音書","ルカ","ルカによる","누가복음":
				bible_short_name = "路"
				switch chap_string {
					case "":
						print_string = "路加福音" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_string("路","路加福音", chap_string, "1")
							default:
								print_string = Bible_print_all_string("路","路加福音", chap_string, sec_string)
						}
				}
			case "John","約","約翰","約翰福音","Joh","joh","От Иоанна святое благовествование","Giăng","ヨハネによる福音書","ヨハネ","ヨハネによる","요한복음":
				bible_short_name = "約"
				switch chap_string {
					case "":
						print_string = "約翰福音" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_string("約","約翰福音", chap_string, "1")
							default:
								print_string = Bible_print_all_string("約","約翰福音", chap_string, sec_string)
						}
				}
			case "Acts","徒","使徒","使徒行傳","Ac","ac","Деяния святых апостолов","Công-vụ Các Sứ-đồ","使徒行伝","사도행전":
				bible_short_name = "徒"
				switch chap_string {
					case "":
						print_string = "使徒行傳" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_string("徒","使徒行傳", chap_string, "1")
							default:
								print_string = Bible_print_all_string("徒","使徒行傳", chap_string, sec_string)
						}
				}
			case "Rom","Romans","羅","羅馬","羅馬書","Ro","ro","Послание к Римлянам","Rô-ma","ローマ","ローマ人への手紙","로마서":
				bible_short_name = "羅"
				switch chap_string {
					case "":
						print_string = "羅馬書" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_string("羅","羅馬書", chap_string, "1")
							default:
								print_string = Bible_print_all_string("羅","羅馬書", chap_string, sec_string)
						}
				}
			case "1 Cor","First Corinthians","林前","哥林多前","哥林多前書","1Co","1co","Первое послание к Коринфянам","1 Cô-rinh-tô","コリント人への第一の手紙","コリント一","コリント人への第一","고린도전서":
				bible_short_name = "林前"
				switch chap_string {
					case "":
						print_string = "哥林多前書" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_string("林前","哥林多前書", chap_string, "1")
							default:
								print_string = Bible_print_all_string("林前","哥林多前書", chap_string, sec_string)
						}
				}
			case "2 Cor","Second Corinthians","林後","哥林多後","哥林多後書","2Co","2co","Второе послание к Коринфянам","2 Cô-rinh-tô","コリント人への第二の手紙","コリント二","コリント人への第二の","고린도후서":
				bible_short_name = "林後"
				switch chap_string {
					case "":
						print_string = "哥林多後書" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_string("林後","哥林多後書", chap_string, "1")
							default:
								print_string = Bible_print_all_string("林後","哥林多後書", chap_string, sec_string)
						}
				}
			case "Gal","Galatians","加","加拉太","加拉太書","Ga","ga","Послание к Галатам","Ga-la-ti","ガラテヤ","ガラテヤ人への手紙","갈라디아서":
				bible_short_name = "加"
				switch chap_string {
					case "":
						print_string = "加拉太書" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_string("加","加拉太書", chap_string, "1")
							default:
								print_string = Bible_print_all_string("加","加拉太書", chap_string, sec_string)
						}
				}
			case "Ephesians","弗","以弗所","以弗所書","Eph","eph","Послание к Ефесянам","Ê-phê-sô","エペソ人への手紙","エペソ","エペソ人","エペソ人の手紙","에베소서":
				bible_short_name = "弗"
				switch chap_string {
					case "":
						print_string = "以弗所書" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_string("弗","以弗所書", chap_string, "1")
							default:
								print_string = Bible_print_all_string("弗","以弗所書", chap_string, sec_string)
						}
				}
			case "Phil","Philippians","腓","腓立","腓立比","腓立比書","Php","php","빌립보서","ピリピ","ピリピ人.ピリピ人への手紙","Послание к Филиппийцам","Phi-líp":
				bible_short_name = "腓"
				switch chap_string {
					case "":
						print_string = "腓立比書" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_string("腓","腓立比書", chap_string, "1")
							default:
								print_string = Bible_print_all_string("腓","腓立比書", chap_string, sec_string)
						}
				}
			case "Col","col","Colossians","西","歌羅西","歌羅","歌羅西書","Послание к Колоссянам","Cô-lô-se","コロサイ人への手紙","コロサイ","コロ","골로새서":
				bible_short_name = "西"
				switch chap_string {
					case "":
						print_string = "歌羅西書" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_string("西","歌羅西書", chap_string, "1")
							default:
								print_string = Bible_print_all_string("西","歌羅西書", chap_string, sec_string)
						}
				}
			case "1 Thess","First Thessalonians","帖前","帖撒羅尼迦前","帖撒羅尼迦前書","1Th","1th","데살로니가전서","テサロニケ人への第一の手紙","テサ一","テサロニケ一","1 Tê-sa-lô-ni-ca","Первое послание к Фессалоникийцам (Солунянам)":
				bible_short_name = "帖前"
				switch chap_string {
					case "":
						print_string = "帖撒羅尼迦前書" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_string("帖前","帖撒羅尼迦前書", chap_string, "1")
							default:
								print_string = Bible_print_all_string("帖前","帖撒羅尼迦前書", chap_string, sec_string)
						}
				}
			case "2 Thess","Second Thessalonians","帖後","帖撒羅尼迦後","帖撒羅尼迦後書","2Th","2th","데살로니가후서","テサロニケ人への第二の手紙","テサ二","テサロニケ二","2 Tê-sa-lô-ni-ca","Второе послание к Фессалоникийцам (Солунянам)":
				bible_short_name = "帖後"
				switch chap_string {
					case "":
						print_string = "帖撒羅尼迦後書" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_string("帖後","帖撒羅尼迦後書", chap_string, "1")
							default:
								print_string = Bible_print_all_string("帖後","帖撒羅尼迦後書", chap_string, sec_string)
						}
				}
			case "1 Tim","First Timothy","提前","提摩太前","提摩太前書","1Ti","1ti","Первое послание к Тимофею","1 Ti-mô-thê","テモテヘの第一の手紙","テモテ一","디모데전서":
				bible_short_name = "提前"
				switch chap_string {
					case "":
						print_string = "提摩太前書" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_string("提前","提摩太前書", chap_string, "1")
							default:
								print_string = Bible_print_all_string("提前","提摩太前書", chap_string, sec_string)
						}
				}
			case "2 Tim","Second Timothy","提後","提摩太後","提摩太後書","2Ti","2ti","Второе послание к Тимофею","2 Ti-mô-thê","テモテヘの第二の手紙","テモテ二","디모데후서":
				bible_short_name = "提後"
				switch chap_string {
					case "":
						print_string = "提摩太後書" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_string("提後","提摩太後書", chap_string, "1")
							default:
								print_string = Bible_print_all_string("提後","提摩太後書", chap_string, sec_string)
						}
				}
			case "Titus","多","提多","提多書","Tit","tit","Послание к Титу","Tít","テトスヘの手紙","テトス","디도서":
				bible_short_name = "多"
				switch chap_string {
					case "":
						print_string = "提多書" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_string("多","提多書", chap_string, "1")
							default:
								print_string = Bible_print_all_string("多","提多書", chap_string, sec_string)
						}
				}
			case "Philem","Philemon","門","腓利","腓利門","腓利門書","Phm","phm","Послание к Филимону","Phi-lê-môn","ピレモンヘの手紙","ピレモン","빌레몬서":
				bible_short_name = "門"
				switch chap_string {
					case "":
						print_string = "腓利門書" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_string("門","腓利門書", chap_string, "1")
							default:
								print_string = Bible_print_all_string("門","腓利門書", chap_string, sec_string)
						}
				}
			case "Heb","Hebrews","來","希伯來","希伯來書","heb","Послание к Евреям","Hê-bơ-rơ","ヘブル人への手紙","ヘブル","히브리서":
				bible_short_name = "來"
				switch chap_string {
					case "":
						print_string = "希伯來書" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_string("來","希伯來書", chap_string, "1")
							default:
								print_string = Bible_print_all_string("來","希伯來書", chap_string, sec_string)
						}
				}
			case "James","雅","雅各","雅各書","Jas","jas","Послание Иакова","Gia-cơ","ヤコブの手紙","ヤコブ","야고보서":
				bible_short_name = "雅"
				switch chap_string {
					case "":
						print_string = "雅各書" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_string("雅","雅各書", chap_string, "1")
							default:
								print_string = Bible_print_all_string("雅","雅各書", chap_string, sec_string)
						}
				}
			case "1 Pet","First Peter","彼前","彼得前","彼得前書","1Pe","1pe","Первое послание Петра","1 Phi-e-rơ","ペテロの第一の手紙","ペテロ一","베드로전서":
				bible_short_name = "彼前"
				switch chap_string {
					case "":
						print_string = "彼得前書" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_string("彼前","彼得前書", chap_string, "1")
							default:
								print_string = Bible_print_all_string("彼前","彼得前書", chap_string, sec_string)
						}
				}
			case "2 Pet","Second Peter","彼後","彼得後","彼得後書","2Pe","2pe","Второе послание Петра","2 Phi-e-rơ","ペテロの第二の手紙","ペテロ","베드로후서":
				bible_short_name = "彼後"
				switch chap_string {
					case "":
						print_string = "彼得後書" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_string("彼後","彼得後書", chap_string, "1")
							default:
								print_string = Bible_print_all_string("彼後","彼得後書", chap_string, sec_string)
						}
				}
			case "1 John","First John","約一","約翰一書","約翰1","約翰1書","1Jo","1jo","Первое послание Иоанна","1 Giăng","ヨハネの第一の手紙","ヨハネ一","요한일서":
				bible_short_name = "約一"
				switch chap_string {
					case "":
						print_string = "約翰一書" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_string("約一","約翰一書", chap_string, "1")
							default:
								print_string = Bible_print_all_string("約一","約翰一書", chap_string, sec_string)
						}
				}
			case "2 John","second John","約二","約翰二書","約翰2","約翰2書","2Jo","Второе послание Иоанна","2 Giăng","ヨハネの第二の手紙","ヨハネ二","요한2서":
				bible_short_name = "約二"
				switch chap_string {
					case "":
						print_string = "約翰二書" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_string("約二","約翰二書", chap_string, "1")
							default:
								print_string = Bible_print_all_string("約二","約翰二書", chap_string, sec_string)
						}
				}
			case "3 John","Third John","約三","約翰三書","約翰3","約翰3書","3Jo","3jo","Третье послание Иоанна","3 Giăng","ヨハネの第三の手紙","ヨハネ三","요한3서":
				bible_short_name = "約三"
				switch chap_string {
					case "":
						print_string = "約翰三書" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_string("約三","約翰三書", chap_string, "1")
							default:
								print_string = Bible_print_all_string("約三","約翰三書", chap_string, sec_string)
						}
				}
			case "Jude","猶","猶大","猶大書","jude","Послание Иуды","Giu-đe","ユダの手紙","ユダ","유다서":
				bible_short_name = "猶"
				switch chap_string {
					case "":
						print_string = "猶大書" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_string("猶","猶大書", chap_string, "1")
							default:
								print_string = Bible_print_all_string("猶","猶大書", chap_string, sec_string)
						}
				}
			case "Rev","Revelation","啟","啟示","啟示錄","Re","re","ｒｅ","Ｒｅ","rev","Откровение ап. Иоанна Богослова (Апокалипсис)","Khải-huyền","ヨハネの黙示録","黙示録","요한계시록":
				bible_short_name = "啟"
				switch chap_string {
					case "":
						print_string = "啟示錄" //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_all_string("啟","啟示錄", chap_string, "1")
							default:
								print_string = Bible_print_all_string("啟","啟示錄", chap_string, sec_string)
						}
				}
			default:
				//print_string = "你是要找 " +  reg.ReplaceAllString(text, "$3") + " 對嗎？\n對不起，我還沒學呢...\n"
				print_string = "聖經"
		}
	case "馬索拉聖經","bhs","Bhs","BHS","ＢＨＳ","Ｂｈｓ","ｂｈｓ":
		log.Print(reg.ReplaceAllString(text, "$3"))
		switch reg.ReplaceAllString(text, "$3") {
			case "Gen","Genesis","創","創世","創世紀","創世記","Ge","ge","gen","창세기","Sáng-thế Ký","Бытие":
				bible_short_name = "創"
				switch chap_string {
					case "":
						print_string = "創世記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("創","創世記", chap_string, "1","bhs")
							default:
								print_string = Bible_print_string("創","創世記", chap_string, sec_string,"bhs")
						}
				}
			case "ex","Ex","Exodus","埃及","出","出埃及","出埃及記","출애굽기","エジプト","出エジプト","出エジプト記","Xuất Ê-díp-tô Ký","Исход":
				bible_short_name = "出"
				switch chap_string {
					case "":
						print_string = "出埃及記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("出","出埃及記", chap_string, "1","bhs")
							default:
								print_string = Bible_print_string("出","出埃及記", chap_string, sec_string,"bhs")
						}
				}
			case "Lev","Leviticus","利","利未","利未記","Le","le","Левит","Lê-vi Ký","レビ記","レビ","레위기":
				bible_short_name = "利"
				switch chap_string {
					case "":
						print_string = "利未記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("利","利未記", chap_string, "1","bhs")
							default:
								print_string = Bible_print_string("利","利未記", chap_string, sec_string,"bhs")
						}
				}
			case "Num","Numbers","民","民數","民數記","Nu","nu","민수기","民数","民数記","Dân-số Ký","Числа":
				bible_short_name = "民"
				switch chap_string {
					case "":
						print_string = "民數記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("民","民數記", chap_string, "1","bhs")
							default:
								print_string = Bible_print_string("民","民數記", chap_string, sec_string,"bhs")
						}
				}
			case "Deut","Deuteronomy","申","申命","申命記","De","de","신명기","Phục-truyền Luật-lệ Ký","Второзаконие":
				bible_short_name = "申"
				switch chap_string {
					case "":
						print_string = "申命記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("申","申命記", chap_string, "1","bhs")
							default:
								print_string = Bible_print_string("申","申命記", chap_string, sec_string,"bhs")
						}
				}
			case "Josh","Joshua","約書亞","約書亞記","Jos","jos","여호수아","ヨシュア記","ヨシュア","Giô-suê","Книга Иисуса Навина":
				bible_short_name = "書"
				switch chap_string {
					case "":
						print_string = "約書亞記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("書","約書亞記", chap_string, "1","bhs")
							default:
								print_string = Bible_print_string("書","約書亞記", chap_string, sec_string,"bhs")
						}
				}
			case "Judg","Judges","士","士師","士師記","Jud","jud","jdg","Jdg","Книга Судей израилевых","Các Quan Xét","사사기":
				bible_short_name = "士"
				switch chap_string {
					case "":
						print_string = "士師記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("士","士師記", chap_string, "1","bhs")
							default:
								print_string = Bible_print_string("士","士師記", chap_string, sec_string,"bhs")
						}
				}
			case "Ruth","路得","路得記","Ru","ru","Rut","rut","룻기","ルツ","ルツ記","Ru-tơ","Книга Руфи":
				bible_short_name = "得"
				switch chap_string {
					case "":
						print_string = "路得記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("得","路得記", chap_string, "1","bhs")
							default:
								print_string = Bible_print_string("得","路得記", chap_string, sec_string,"bhs")
						}
				}
			case "1 Sam","First Samuel","撒上","撒母耳記上","1Sa","1sa","サムエル記上","サムエル上","サム上","사무엘상","1 Sa-mu-ên","Первая книга Царств":
				bible_short_name = "撒上"
				switch chap_string {
					case "":
						print_string = "撒母耳記上"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("撒上","撒母耳記上", chap_string, "1","bhs")
							default:
								print_string = Bible_print_string("撒上","撒母耳記上", chap_string, sec_string,"bhs")
						}
				}
			case "2 Sam","Second Samuel","撒下","撒母耳記下","2Sa","2sa","사무엘하","サムエル記下","サムエル下","サム下","2 Sa-mu-ên","Вторая книга Царств":
				bible_short_name = "撒下"
				switch chap_string {
					case "":
						print_string = "撒母耳記下"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("撒下","撒母耳記下", chap_string, "1","bhs")
							default:
								print_string = Bible_print_string("撒下","撒母耳記下", chap_string, sec_string,"bhs")
						}
				}
			case "1 Kin","First Kings","王上","列王上","列王紀上","列王記上","1Ki","1ki","열왕기상","Третья книга Царств","1 Các Vua":
				bible_short_name = "王上"
				switch chap_string {
					case "":
						print_string = "列王紀上"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("王上","列王紀上", chap_string, "1","bhs")
							default:
								print_string = Bible_print_string("王上","列王紀上", chap_string, sec_string,"bhs")
						}
				}
			case "2 Kin","Second Kings","王下","列王下","列王記下","列王紀下","2Ki","2ki","열왕기하","2 Các Vua","Четвертая книга Царств":
				bible_short_name = "王下"
				switch chap_string {
					case "":
						print_string = "列王紀下"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("王下","列王紀下", chap_string, "1","bhs")
							default:
								print_string = Bible_print_string("王下","列王紀下", chap_string, sec_string,"bhs")
						}
				}
			case "1 Chr","First Chronicles","歷上","代上","歷代志上","歷代上","1Ch","1ch","Первая книга Паралипоменон","1 Sử-ký","歴上","歴代上","歴代志上","역대상":
				bible_short_name = "代上"
				switch chap_string {
					case "":
						print_string = "歷代志上"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("代上","歷代志上", chap_string, "1","bhs")
							default:
								print_string = Bible_print_string("代上","歷代志上", chap_string, sec_string,"bhs")
						}
				}
			case "2 Chr","Second Chronicles","代下","歷下","歷代下","歷代志下","2Ch","2ch","역대하","歴代志下","歴代下","歴下","2 Sử-ký","Вторая книга Паралипоменон":
				bible_short_name = "代下"
				switch chap_string {
					case "":
						print_string = "歷代志下"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("代下","歷代志下", chap_string, "1","bhs")
							default:
								print_string = Bible_print_string("代下","歷代志下", chap_string, sec_string,"bhs")
						}
				}
			case "Ezra","拉","以斯拉","以斯拉記","Ezr","ezr","Первая книга Ездры","Ê-xơ-ra","エズラ","エズラ記","에스라":
				bible_short_name = "拉"
				switch chap_string {
					case "":
						print_string = "以斯拉記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("拉","以斯拉記", chap_string, "1","bhs")
							default:
								print_string = Bible_print_string("拉","以斯拉記", chap_string, sec_string,"bhs")
						}
				}
			case "Neh","Nehemiah","尼","尼希米","尼希米記","Ne","ne","느헤미야","ネヘミヤ書","ネヘミヤ","Nê-hê-mi","Книга Неемии":
				bible_short_name = "尼"
				switch chap_string {
					case "":
						print_string = "尼希米記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("尼","尼希米記", chap_string, "1","bhs")
							default:
								print_string = Bible_print_string("尼","尼希米記", chap_string, sec_string,"bhs")
						}
				}
			case "Esth","Esther","斯","以斯帖","以斯帖記","Es","est","Есфирь","Ê-xơ-tê","エステル","エステル記","에스더":
				bible_short_name = "斯"
				switch chap_string {
					case "":
						print_string = "以斯帖記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("斯","以斯帖記", chap_string, "1","bhs")
							default:
								print_string = Bible_print_string("斯","以斯帖記", chap_string, sec_string,"bhs")
						}
				}
			case "Job","job","伯","約伯","約伯記","Книга Иова","Gióp","ヨブ","ヨブ記","욥기":
				bible_short_name = "伯"
				switch chap_string {
					case "":
						print_string = "約伯記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("伯","約伯記", chap_string, "1","bhs")
							default:
								print_string = Bible_print_string("伯","約伯記", chap_string, sec_string,"bhs")
						}
				}
			case "Ps","Psalms","詩","詩篇","ps","시편","Thi-thiên","Псалтирь":
				bible_short_name = "詩"
				switch chap_string {
					case "":
						print_string = "詩篇"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("詩","詩篇", chap_string, "1","bhs")
							default:
								print_string = Bible_print_string("詩","詩篇", chap_string, sec_string,"bhs")
						}
				}
			case "Prov","Proverbs","箴","箴言","Pr","pr","Притчи Соломона","Châm-ngôn","잠언":
				bible_short_name = "箴"
				switch chap_string {
					case "":
						print_string = "箴言"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("箴","箴言", chap_string, "1","bhs")
							default:
								print_string = Bible_print_string("箴","箴言", chap_string, sec_string,"bhs")
						}
				}
			case "Eccl","Ecclesiastes","傳","傳道","傳道書","Ec","ec","Книга Екклезиаста","или Проповедника","Truyền-đạo","伝道の書","伝道","伝","伝道書","전도서":
				bible_short_name = "傳"
				switch chap_string {
					case "":
						print_string = "傳道書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("傳","傳道書", chap_string, "1","bhs")
							default:
								print_string = Bible_print_string("傳","傳道書", chap_string, sec_string,"bhs")
						}
				}
			case "Song","Song of Solomon","歌","雅歌","So","so","sng","Sng","Песнь песней Соломона","Nhã-ca","아가":
				bible_short_name = "歌"
				switch chap_string {
					case "":
						print_string = "雅歌"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("歌","雅歌", chap_string, "1","bhs")
							default:
								print_string = Bible_print_string("歌","雅歌", chap_string, sec_string,"bhs")
						}
				}
			case "Is","Isaiah","賽","以賽","以賽亞","以賽亞書","Isa","isa","Книга пророка Исаии","Ê-sai","イザヤ書","イザヤ","이사야":
				bible_short_name = "賽"
				switch chap_string {
					case "":
						print_string = "以賽亞書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("賽","以賽亞書", chap_string, "1","bhs")
							default:
								print_string = Bible_print_string("賽","以賽亞書", chap_string, sec_string,"bhs")
						}
				}
			case "Jer","Jeremiah","耶","耶利米","耶利米書","jer","예레미야","エレミヤ","エレミヤ書","Giê-rê-mi","Книга пророка Иеремии":
				bible_short_name = "耶"
				switch chap_string {
					case "":
						print_string = "耶利米書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("耶","耶利米書", chap_string, "1","bhs")
							default:
								print_string = Bible_print_string("耶","耶利米書", chap_string, sec_string,"bhs")
						}
				}
			case "Lam","Lamentations","哀","哀歌","耶利米哀歌","La","lam","예레미야애가","Ca-thương","Плач Иеремии":
				bible_short_name = "哀"
				switch chap_string {
					case "":
						print_string = "耶利米哀歌"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("哀","耶利米哀歌", chap_string, "1","bhs")
							default:
								print_string = Bible_print_string("哀","耶利米哀歌", chap_string, sec_string,"bhs")
						}
				}
			case "Ezek","Ezekiel","結","以西結","以西結書","Eze","eze","에스겔","エゼキエル書","エゼキエル","Ê-xê-chi-ên","Книга пророка Иезекииля":
				bible_short_name = "結"
				switch chap_string {
					case "":
						print_string = "以西結書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("結","以西結書", chap_string, "1","bhs")
							default:
								print_string = Bible_print_string("結","以西結書", chap_string, sec_string,"bhs")
						}
				}
			case "Dan","Daniel","但","但以理","但以理書","Da","da","Книга пророка Даниила","Đa-ni-ên","ダニエル書","ダニエル","다니엘":
				bible_short_name = "但"
				switch chap_string {
					case "":
						print_string = "但以理書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("但","但以理書", chap_string, "1","bhs")
							default:
								print_string = Bible_print_string("但","但以理書", chap_string, sec_string,"bhs")
						}
				}
			case "Hos","Hosea","何","何西","何西阿","何西阿書","Ho","ho","Книга пророка Осии","Ô-sê","ホセア書","ホセア","호세아":
				bible_short_name = "何"
				switch chap_string {
					case "":
						print_string = "何西阿書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("何","何西阿書", chap_string, "1","bhs")
							default:
								print_string = Bible_print_string("何","何西阿書", chap_string, sec_string,"bhs")
						}
				}
			case "Joel","珥","約珥","約珥書","Joe","joe","Книга пророка Иоиля","Giô-ên","ヨエル書","ヨエル","요엘":
				bible_short_name = "珥"
				switch chap_string {
					case "":
						print_string = "約珥書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("珥","約珥書", chap_string, "1","bhs")
							default:
								print_string = Bible_print_string("珥","約珥書", chap_string, sec_string,"bhs")
						}
				}
			case "Amos","摩","阿摩司書","Am","am","Книга пророка Амоса","A-mốt","アモス書","アモス","아모스":
				bible_short_name = "摩"
				switch chap_string {
					case "":
						print_string = "阿摩司書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("摩","阿摩司書", chap_string, "1","bhs")
							default:
								print_string = Bible_print_string("摩","阿摩司書", chap_string, sec_string,"bhs")
						}
				}
			case "Obad","Obadiah","俄","俄巴底亞","俄巴底亞書","Ob","ob","오바댜","オバデヤ書","オバデヤ","Áp-đia","Книга пророка Авдия":
				bible_short_name = "俄"
				switch chap_string {
					case "":
						print_string = "俄巴底亞書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("俄","俄巴底亞書", chap_string, "1","bhs")
							default:
								print_string = Bible_print_string("俄","俄巴底亞書", chap_string, sec_string,"bhs")
						}
				}
			case "Jon","Jonah","拿","約拿","約拿書","jon","요나","ヨナ書","ヨナ","Giô-na","Книга пророка Ионы":
				bible_short_name = "拿"
				switch chap_string {
					case "":
						print_string = "約拿書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("拿","約拿書", chap_string, "1","bhs")
							default:
								print_string = Bible_print_string("拿","約拿書", chap_string, sec_string,"bhs")
						}
				}
			case "Micah","彌","彌迦","彌迦書","Mic","mic","Книга пророка Михея","Mi-chê","ミカ書","ミカ","미가":
				bible_short_name = "彌"
				switch chap_string {
					case "":
						print_string = "彌迦書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("彌","彌迦書", chap_string, "1","bhs")
							default:
								print_string = Bible_print_string("彌","彌迦書", chap_string, sec_string,"bhs")
						}
				}
			case "Nah","Nahum","鴻","那鴻","那鴻書","Na","na","Книга пророка Наума","Na-hum","ナホム書","ナホム","나훔":
				bible_short_name = "鴻"
				switch chap_string {
					case "":
						print_string = "那鴻書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("鴻","那鴻書", chap_string, "1","bhs")
							default:
								print_string = Bible_print_string("鴻","那鴻書", chap_string, sec_string,"bhs")
						}
				}
			case "Habakkuk","哈","哈巴","哈巴谷","哈巴谷書","Hab","hab","Книга пророка Аввакума","Ha-ba-cúc","ハバクク書","ハバクク","ハバ","クク","ハバ書","하박국":
				bible_short_name = "哈"
				switch chap_string {
					case "":
						print_string = "哈巴谷書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("哈","哈巴谷書", chap_string, "1","bhs")
							default:
								print_string = Bible_print_string("哈","哈巴谷書", chap_string, sec_string,"bhs")
						}
				}
			case "Zeph","Zephaniah","番","西番雅","西番雅書","Zep","zep","스바냐","ゼパニヤ書","ゼパニヤ","Sô-phô-ni","Книга пророка Софонии":
				bible_short_name = "番"
				switch chap_string {
					case "":
						print_string = "西番雅書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("番","西番雅書", chap_string, "1","bhs")
							default:
								print_string = Bible_print_string("番","西番雅書", chap_string, sec_string,"bhs")
						}
				}
			case "Haggai","該","哈該","哈該書","Hag","hag","학개","ハガイ書","ハガイ","A-ghê","Книга пророка Аггея":
				bible_short_name = "該"
				switch chap_string {
					case "":
						print_string = "哈該書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("該","哈該書", chap_string, "1","bhs")
							default:
								print_string = Bible_print_string("該","哈該書", chap_string, sec_string,"bhs")
						}
				}
			case "Zech","Zechariah","亞","撒迦利亞","撒迦利亞書","Zec","zec","Книга пророка Захарии","Xa-cha-ri","스가랴","ゼカリヤ書","ゼカリヤ":
				bible_short_name = "亞"
				switch chap_string {
					case "":
						print_string = "撒迦利亞書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("亞","撒迦利亞書", chap_string, "1","bhs")
							default:
								print_string = Bible_print_string("亞","撒迦利亞書", chap_string, sec_string,"bhs")
						}
				}
			case "Malachi","瑪","","瑪拉","瑪拉基","瑪拉基書","Mal","mal","말라기","マラキ書","マラキ","Ma-la-chi","Книга пророка Малахии":
				bible_short_name = "瑪"
				switch chap_string {
					case "":
						print_string = "瑪拉基書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("瑪","瑪拉基書", chap_string, "1","bhs")
							default:
								print_string = Bible_print_string("瑪","瑪拉基書", chap_string, sec_string,"bhs")
						}
				}
			case "Matt","Matthew","太","馬太","馬太福音","Mt","mt","마태복음","マタイによる福音書","マタイ","マタイによる","Ma-thi-ơ","От Матфея святое благовествование":
				bible_short_name = "太"
				switch chap_string {
					case "":
						print_string = "馬太福音"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("太","馬太福音", chap_string, "1","bhs")
							default:
								print_string = Bible_print_string("太","馬太福音", chap_string, sec_string,"bhs")
						}
				}
			case "Mark","可","馬可","馬可福音","Mr","mr","マルコによる福音書","マルコ","マルコによる","마가복음","Mác","От Марка святое благовествование":
				bible_short_name = "可"
				switch chap_string {
					case "":
						print_string = "馬可福音"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("可","馬可福音", chap_string, "1","bhs")
							default:
								print_string = Bible_print_string("可","馬可福音", chap_string, sec_string,"bhs")
						}
				}
			case "Luke","路","路加","路加福音","Lu","lu","От Луки святое благовествование","Lu-ca","ルカによる福音書","ルカ","ルカによる","누가복음":
				bible_short_name = "路"
				switch chap_string {
					case "":
						print_string = "路加福音"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("路","路加福音", chap_string, "1","bhs")
							default:
								print_string = Bible_print_string("路","路加福音", chap_string, sec_string,"bhs")
						}
				}
			case "John","約","約翰","約翰福音","Joh","joh","От Иоанна святое благовествование","Giăng","ヨハネによる福音書","ヨハネ","ヨハネによる","요한복음":
				bible_short_name = "約"
				switch chap_string {
					case "":
						print_string = "約翰福音"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("約","約翰福音", chap_string, "1","bhs")
							default:
								print_string = Bible_print_string("約","約翰福音", chap_string, sec_string,"bhs")
						}
				}
			case "Acts","徒","使徒","使徒行傳","Ac","ac","Деяния святых апостолов","Công-vụ Các Sứ-đồ","使徒行伝","사도행전":
				bible_short_name = "徒"
				switch chap_string {
					case "":
						print_string = "使徒行傳"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("徒","使徒行傳", chap_string, "1","bhs")
							default:
								print_string = Bible_print_string("徒","使徒行傳", chap_string, sec_string,"bhs")
						}
				}
			case "Rom","Romans","羅","羅馬","羅馬書","Ro","ro","Послание к Римлянам","Rô-ma","ローマ","ローマ人への手紙","로마서":
				bible_short_name = "羅"
				switch chap_string {
					case "":
						print_string = "羅馬書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("羅","羅馬書", chap_string, "1","bhs")
							default:
								print_string = Bible_print_string("羅","羅馬書", chap_string, sec_string,"bhs")
						}
				}
			case "1 Cor","First Corinthians","林前","哥林多前","哥林多前書","1Co","1co","Первое послание к Коринфянам","1 Cô-rinh-tô","コリント人への第一の手紙","コリント一","コリント人への第一","고린도전서":
				bible_short_name = "林前"
				switch chap_string {
					case "":
						print_string = "哥林多前書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("林前","哥林多前書", chap_string, "1","bhs")
							default:
								print_string = Bible_print_string("林前","哥林多前書", chap_string, sec_string,"bhs")
						}
				}
			case "2 Cor","Second Corinthians","林後","哥林多後","哥林多後書","2Co","2co","Второе послание к Коринфянам","2 Cô-rinh-tô","コリント人への第二の手紙","コリント二","コリント人への第二の","고린도후서":
				bible_short_name = "林後"
				switch chap_string {
					case "":
						print_string = "哥林多後書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("林後","哥林多後書", chap_string, "1","bhs")
							default:
								print_string = Bible_print_string("林後","哥林多後書", chap_string, sec_string,"bhs")
						}
				}
			case "Gal","Galatians","加","加拉太","加拉太書","Ga","ga","Послание к Галатам","Ga-la-ti","ガラテヤ","ガラテヤ人への手紙","갈라디아서":
				bible_short_name = "加"
				switch chap_string {
					case "":
						print_string = "加拉太書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("加","加拉太書", chap_string, "1","bhs")
							default:
								print_string = Bible_print_string("加","加拉太書", chap_string, sec_string,"bhs")
						}
				}
			case "Ephesians","弗","以弗所","以弗所書","Eph","eph","Послание к Ефесянам","Ê-phê-sô","エペソ人への手紙","エペソ","エペソ人","エペソ人の手紙","에베소서":
				bible_short_name = "弗"
				switch chap_string {
					case "":
						print_string = "以弗所書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("弗","以弗所書", chap_string, "1","bhs")
							default:
								print_string = Bible_print_string("弗","以弗所書", chap_string, sec_string,"bhs")
						}
				}
			case "Phil","Philippians","腓","腓立","腓立比","腓立比書","Php","php","빌립보서","ピリピ","ピリピ人.ピリピ人への手紙","Послание к Филиппийцам","Phi-líp":
				bible_short_name = "腓"
				switch chap_string {
					case "":
						print_string = "腓立比書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("腓","腓立比書", chap_string, "1","bhs")
							default:
								print_string = Bible_print_string("腓","腓立比書", chap_string, sec_string,"bhs")
						}
				}
			case "Col","col","Colossians","西","歌羅西","歌羅","歌羅西書","Послание к Колоссянам","Cô-lô-se","コロサイ人への手紙","コロサイ","コロ","골로새서":
				bible_short_name = "西"
				switch chap_string {
					case "":
						print_string = "歌羅西書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("西","歌羅西書", chap_string, "1","bhs")
							default:
								print_string = Bible_print_string("西","歌羅西書", chap_string, sec_string,"bhs")
						}
				}
			case "1 Thess","First Thessalonians","帖前","帖撒羅尼迦前","帖撒羅尼迦前書","1Th","1th","데살로니가전서","テサロニケ人への第一の手紙","テサ一","テサロニケ一","1 Tê-sa-lô-ni-ca","Первое послание к Фессалоникийцам (Солунянам)":
				bible_short_name = "帖前"
				switch chap_string {
					case "":
						print_string = "帖撒羅尼迦前書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("帖前","帖撒羅尼迦前書", chap_string, "1","bhs")
							default:
								print_string = Bible_print_string("帖前","帖撒羅尼迦前書", chap_string, sec_string,"bhs")
						}
				}
			case "2 Thess","Second Thessalonians","帖後","帖撒羅尼迦後","帖撒羅尼迦後書","2Th","2th","데살로니가후서","テサロニケ人への第二の手紙","テサ二","テサロニケ二","2 Tê-sa-lô-ni-ca","Второе послание к Фессалоникийцам (Солунянам)":
				bible_short_name = "帖後"
				switch chap_string {
					case "":
						print_string = "帖撒羅尼迦後書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("帖後","帖撒羅尼迦後書", chap_string, "1","bhs")
							default:
								print_string = Bible_print_string("帖後","帖撒羅尼迦後書", chap_string, sec_string,"bhs")
						}
				}
			case "1 Tim","First Timothy","提前","提摩太前","提摩太前書","1Ti","1ti","Первое послание к Тимофею","1 Ti-mô-thê","テモテヘの第一の手紙","テモテ一","디모데전서":
				bible_short_name = "提前"
				switch chap_string {
					case "":
						print_string = "提摩太前書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("提前","提摩太前書", chap_string, "1","bhs")
							default:
								print_string = Bible_print_string("提前","提摩太前書", chap_string, sec_string,"bhs")
						}
				}
			case "2 Tim","Second Timothy","提後","提摩太後","提摩太後書","2Ti","2ti","Второе послание к Тимофею","2 Ti-mô-thê","テモテヘの第二の手紙","テモテ二","디모데후서":
				bible_short_name = "提後"
				switch chap_string {
					case "":
						print_string = "提摩太後書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("提後","提摩太後書", chap_string, "1","bhs")
							default:
								print_string = Bible_print_string("提後","提摩太後書", chap_string, sec_string,"bhs")
						}
				}
			case "Titus","多","提多","提多書","Tit","tit","Послание к Титу","Tít","テトスヘの手紙","テトス","디도서":
				bible_short_name = "多"
				switch chap_string {
					case "":
						print_string = "提多書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("多","提多書", chap_string, "1","bhs")
							default:
								print_string = Bible_print_string("多","提多書", chap_string, sec_string,"bhs")
						}
				}
			case "Philem","Philemon","門","腓利","腓利門","腓利門書","Phm","phm","Послание к Филимону","Phi-lê-môn","ピレモンヘの手紙","ピレモン","빌레몬서":
				bible_short_name = "門"
				switch chap_string {
					case "":
						print_string = "腓利門書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("門","腓利門書", chap_string, "1","bhs")
							default:
								print_string = Bible_print_string("門","腓利門書", chap_string, sec_string,"bhs")
						}
				}
			case "Heb","Hebrews","來","希伯來","希伯來書","heb","Послание к Евреям","Hê-bơ-rơ","ヘブル人への手紙","ヘブル","히브리서":
				bible_short_name = "來"
				switch chap_string {
					case "":
						print_string = "希伯來書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("來","希伯來書", chap_string, "1","bhs")
							default:
								print_string = Bible_print_string("來","希伯來書", chap_string, sec_string,"bhs")
						}
				}
			case "James","雅","雅各","雅各書","Jas","jas","Послание Иакова","Gia-cơ","ヤコブの手紙","ヤコブ","야고보서":
				bible_short_name = "雅"
				switch chap_string {
					case "":
						print_string = "雅各書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("雅","雅各書", chap_string, "1","bhs")
							default:
								print_string = Bible_print_string("雅","雅各書", chap_string, sec_string,"bhs")
						}
				}
			case "1 Pet","First Peter","彼前","彼得前","彼得前書","1Pe","1pe","Первое послание Петра","1 Phi-e-rơ","ペテロの第一の手紙","ペテロ一","베드로전서":
				bible_short_name = "彼前"
				switch chap_string {
					case "":
						print_string = "彼得前書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("彼前","彼得前書", chap_string, "1","bhs")
							default:
								print_string = Bible_print_string("彼前","彼得前書", chap_string, sec_string,"bhs")
						}
				}
			case "2 Pet","Second Peter","彼後","彼得後","彼得後書","2Pe","2pe","Второе послание Петра","2 Phi-e-rơ","ペテロの第二の手紙","ペテロ","베드로후서":
				bible_short_name = "彼後"
				switch chap_string {
					case "":
						print_string = "彼得後書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("彼後","彼得後書", chap_string, "1","bhs")
							default:
								print_string = Bible_print_string("彼後","彼得後書", chap_string, sec_string,"bhs")
						}
				}
			case "1 John","First John","約一","約翰一書","約翰1","約翰1書","1Jo","1jo","Первое послание Иоанна","1 Giăng","ヨハネの第一の手紙","ヨハネ一","요한일서":
				bible_short_name = "約一"
				switch chap_string {
					case "":
						print_string = "約翰一書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("約一","約翰一書", chap_string, "1","bhs")
							default:
								print_string = Bible_print_string("約一","約翰一書", chap_string, sec_string,"bhs")
						}
				}
			case "2 John","second John","約二","約翰二書","約翰2","約翰2書","2Jo","Второе послание Иоанна","2 Giăng","ヨハネの第二の手紙","ヨハネ二","요한2서":
				bible_short_name = "約二"
				switch chap_string {
					case "":
						print_string = "約翰二書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("約二","約翰二書", chap_string, "1","bhs")
							default:
								print_string = Bible_print_string("約二","約翰二書", chap_string, sec_string,"bhs")
						}
				}
			case "3 John","Third John","約三","約翰三書","約翰3","約翰3書","3Jo","3jo","Третье послание Иоанна","3 Giăng","ヨハネの第三の手紙","ヨハネ三","요한3서":
				bible_short_name = "約三"
				switch chap_string {
					case "":
						print_string = "約翰三書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("約三","約翰三書", chap_string, "1","bhs")
							default:
								print_string = Bible_print_string("約三","約翰三書", chap_string, sec_string,"bhs")
						}
				}
			case "Jude","猶","猶大","猶大書","jude","Послание Иуды","Giu-đe","ユダの手紙","ユダ","유다서":
				bible_short_name = "猶"
				switch chap_string {
					case "":
						print_string = "猶大書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("猶","猶大書", chap_string, "1","bhs")
							default:
								print_string = Bible_print_string("猶","猶大書", chap_string, sec_string,"bhs")
						}
				}
			case "Rev","Revelation","啟","啟示","啟示錄","Re","re","ｒｅ","Ｒｅ","rev","Откровение ап. Иоанна Богослова (Апокалипсис)","Khải-huyền","ヨハネの黙示録","黙示録","요한계시록":
				bible_short_name = "啟"
				switch chap_string {
					case "":
						print_string = "啟示錄"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("啟","啟示錄", chap_string, "1","bhs")
							default:
								print_string = Bible_print_string("啟","啟示錄", chap_string, sec_string,"bhs")
						}
				}
			default:
				print_string = "聖經"
				//print_string = "你是要找 " +  reg.ReplaceAllString(text, "$3") + " 對嗎？\n對不起，我還沒學呢...\n"
		}
	case "希臘聖經","lxx","LXX","Lxx","ＬＸＸ","Ｌｘｘ","ｌｘｘ":
		log.Print(reg.ReplaceAllString(text, "$3"))
		switch reg.ReplaceAllString(text, "$3") {
			case "Gen","Genesis","創","創世","創世紀","創世記","Ge","ge","gen","창세기","Sáng-thế Ký","Бытие":
				bible_short_name = "創"
				switch chap_string {
					case "":
						print_string = "創世記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("創","創世記", chap_string, "1","lxx")
							default:
								print_string = Bible_print_string("創","創世記", chap_string, sec_string,"lxx")
						}
				}
			case "ex","Ex","Exodus","埃及","出","出埃及","出埃及記","출애굽기","エジプト","出エジプト","出エジプト記","Xuất Ê-díp-tô Ký","Исход":
				bible_short_name = "出"
				switch chap_string {
					case "":
						print_string = "出埃及記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("出","出埃及記", chap_string, "1","lxx")
							default:
								print_string = Bible_print_string("出","出埃及記", chap_string, sec_string,"lxx")
						}
				}
			case "Lev","Leviticus","利","利未","利未記","Le","le","Левит","Lê-vi Ký","レビ記","レビ","레위기":
				bible_short_name = "利"
				switch chap_string {
					case "":
						print_string = "利未記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("利","利未記", chap_string, "1","lxx")
							default:
								print_string = Bible_print_string("利","利未記", chap_string, sec_string,"lxx")
						}
				}
			case "Num","Numbers","民","民數","民數記","Nu","nu","민수기","民数","民数記","Dân-số Ký","Числа":
				bible_short_name = "民"
				switch chap_string {
					case "":
						print_string = "民數記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("民","民數記", chap_string, "1","lxx")
							default:
								print_string = Bible_print_string("民","民數記", chap_string, sec_string,"lxx")
						}
				}
			case "Deut","Deuteronomy","申","申命","申命記","De","de","신명기","Phục-truyền Luật-lệ Ký","Второзаконие":
				bible_short_name = "申"
				switch chap_string {
					case "":
						print_string = "申命記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("申","申命記", chap_string, "1","lxx")
							default:
								print_string = Bible_print_string("申","申命記", chap_string, sec_string,"lxx")
						}
				}
			case "Josh","Joshua","約書亞","約書亞記","Jos","jos","여호수아","ヨシュア記","ヨシュア","Giô-suê","Книга Иисуса Навина":
				bible_short_name = "書"
				switch chap_string {
					case "":
						print_string = "約書亞記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("書","約書亞記", chap_string, "1","lxx")
							default:
								print_string = Bible_print_string("書","約書亞記", chap_string, sec_string,"lxx")
						}
				}
			case "Judg","Judges","士","士師","士師記","Jud","jud","jdg","Jdg","Книга Судей израилевых","Các Quan Xét","사사기":
				bible_short_name = "士"
				switch chap_string {
					case "":
						print_string = "士師記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("士","士師記", chap_string, "1","lxx")
							default:
								print_string = Bible_print_string("士","士師記", chap_string, sec_string,"lxx")
						}
				}
			case "Ruth","路得","路得記","Ru","ru","Rut","rut","룻기","ルツ","ルツ記","Ru-tơ","Книга Руфи":
				bible_short_name = "得"
				switch chap_string {
					case "":
						print_string = "路得記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("得","路得記", chap_string, "1","lxx")
							default:
								print_string = Bible_print_string("得","路得記", chap_string, sec_string,"lxx")
						}
				}
			case "1 Sam","First Samuel","撒上","撒母耳記上","1Sa","1sa","サムエル記上","サムエル上","サム上","사무엘상","1 Sa-mu-ên","Первая книга Царств":
				bible_short_name = "撒上"
				switch chap_string {
					case "":
						print_string = "撒母耳記上"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("撒上","撒母耳記上", chap_string, "1","lxx")
							default:
								print_string = Bible_print_string("撒上","撒母耳記上", chap_string, sec_string,"lxx")
						}
				}
			case "2 Sam","Second Samuel","撒下","撒母耳記下","2Sa","2sa","사무엘하","サムエル記下","サムエル下","サム下","2 Sa-mu-ên","Вторая книга Царств":
				bible_short_name = "撒下"
				switch chap_string {
					case "":
						print_string = "撒母耳記下"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("撒下","撒母耳記下", chap_string, "1","lxx")
							default:
								print_string = Bible_print_string("撒下","撒母耳記下", chap_string, sec_string,"lxx")
						}
				}
			case "1 Kin","First Kings","王上","列王上","列王紀上","列王記上","1Ki","1ki","열왕기상","Третья книга Царств","1 Các Vua":
				bible_short_name = "王上"
				switch chap_string {
					case "":
						print_string = "列王紀上"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("王上","列王紀上", chap_string, "1","lxx")
							default:
								print_string = Bible_print_string("王上","列王紀上", chap_string, sec_string,"lxx")
						}
				}
			case "2 Kin","Second Kings","王下","列王下","列王記下","列王紀下","2Ki","2ki","열왕기하","2 Các Vua","Четвертая книга Царств":
				bible_short_name = "王下"
				switch chap_string {
					case "":
						print_string = "列王紀下"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("王下","列王紀下", chap_string, "1","lxx")
							default:
								print_string = Bible_print_string("王下","列王紀下", chap_string, sec_string,"lxx")
						}
				}
			case "1 Chr","First Chronicles","歷上","代上","歷代志上","歷代上","1Ch","1ch","Первая книга Паралипоменон","1 Sử-ký","歴上","歴代上","歴代志上","역대상":
				bible_short_name = "代上"
				switch chap_string {
					case "":
						print_string = "歷代志上"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("代上","歷代志上", chap_string, "1","lxx")
							default:
								print_string = Bible_print_string("代上","歷代志上", chap_string, sec_string,"lxx")
						}
				}
			case "2 Chr","Second Chronicles","代下","歷下","歷代下","歷代志下","2Ch","2ch","역대하","歴代志下","歴代下","歴下","2 Sử-ký","Вторая книга Паралипоменон":
				bible_short_name = "代下"
				switch chap_string {
					case "":
						print_string = "歷代志下"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("代下","歷代志下", chap_string, "1","lxx")
							default:
								print_string = Bible_print_string("代下","歷代志下", chap_string, sec_string,"lxx")
						}
				}
			case "Ezra","拉","以斯拉","以斯拉記","Ezr","ezr","Первая книга Ездры","Ê-xơ-ra","エズラ","エズラ記","에스라":
				bible_short_name = "拉"
				switch chap_string {
					case "":
						print_string = "以斯拉記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("拉","以斯拉記", chap_string, "1","lxx")
							default:
								print_string = Bible_print_string("拉","以斯拉記", chap_string, sec_string,"lxx")
						}
				}
			case "Neh","Nehemiah","尼","尼希米","尼希米記","Ne","ne","느헤미야","ネヘミヤ書","ネヘミヤ","Nê-hê-mi","Книга Неемии":
				bible_short_name = "尼"
				switch chap_string {
					case "":
						print_string = "尼希米記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("尼","尼希米記", chap_string, "1","lxx")
							default:
								print_string = Bible_print_string("尼","尼希米記", chap_string, sec_string,"lxx")
						}
				}
			case "Esth","Esther","斯","以斯帖","以斯帖記","Es","est","Есфирь","Ê-xơ-tê","エステル","エステル記","에스더":
				bible_short_name = "斯"
				switch chap_string {
					case "":
						print_string = "以斯帖記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("斯","以斯帖記", chap_string, "1","lxx")
							default:
								print_string = Bible_print_string("斯","以斯帖記", chap_string, sec_string,"lxx")
						}
				}
			case "Job","job","伯","約伯","約伯記","Книга Иова","Gióp","ヨブ","ヨブ記","욥기":
				bible_short_name = "伯"
				switch chap_string {
					case "":
						print_string = "約伯記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("伯","約伯記", chap_string, "1","lxx")
							default:
								print_string = Bible_print_string("伯","約伯記", chap_string, sec_string,"lxx")
						}
				}
			case "Ps","Psalms","詩","詩篇","ps","시편","Thi-thiên","Псалтирь":
				bible_short_name = "詩"
				switch chap_string {
					case "":
						print_string = "詩篇"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("詩","詩篇", chap_string, "1","lxx")
							default:
								print_string = Bible_print_string("詩","詩篇", chap_string, sec_string,"lxx")
						}
				}
			case "Prov","Proverbs","箴","箴言","Pr","pr","Притчи Соломона","Châm-ngôn","잠언":
				bible_short_name = "箴"
				switch chap_string {
					case "":
						print_string = "箴言"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("箴","箴言", chap_string, "1","lxx")
							default:
								print_string = Bible_print_string("箴","箴言", chap_string, sec_string,"lxx")
						}
				}
			case "Eccl","Ecclesiastes","傳","傳道","傳道書","Ec","ec","Книга Екклезиаста","или Проповедника","Truyền-đạo","伝道の書","伝道","伝","伝道書","전도서":
				bible_short_name = "傳"
				switch chap_string {
					case "":
						print_string = "傳道書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("傳","傳道書", chap_string, "1","lxx")
							default:
								print_string = Bible_print_string("傳","傳道書", chap_string, sec_string,"lxx")
						}
				}
			case "Song","Song of Solomon","歌","雅歌","So","so","sng","Sng","Песнь песней Соломона","Nhã-ca","아가":
				bible_short_name = "歌"
				switch chap_string {
					case "":
						print_string = "雅歌"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("歌","雅歌", chap_string, "1","lxx")
							default:
								print_string = Bible_print_string("歌","雅歌", chap_string, sec_string,"lxx")
						}
				}
			case "Is","Isaiah","賽","以賽","以賽亞","以賽亞書","Isa","isa","Книга пророка Исаии","Ê-sai","イザヤ書","イザヤ","이사야":
				bible_short_name = "賽"
				switch chap_string {
					case "":
						print_string = "以賽亞書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("賽","以賽亞書", chap_string, "1","lxx")
							default:
								print_string = Bible_print_string("賽","以賽亞書", chap_string, sec_string,"lxx")
						}
				}
			case "Jer","Jeremiah","耶","耶利米","耶利米書","jer","예레미야","エレミヤ","エレミヤ書","Giê-rê-mi","Книга пророка Иеремии":
				bible_short_name = "耶"
				switch chap_string {
					case "":
						print_string = "耶利米書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("耶","耶利米書", chap_string, "1","lxx")
							default:
								print_string = Bible_print_string("耶","耶利米書", chap_string, sec_string,"lxx")
						}
				}
			case "Lam","Lamentations","哀","哀歌","耶利米哀歌","La","lam","예레미야애가","Ca-thương","Плач Иеремии":
				bible_short_name = "哀"
				switch chap_string {
					case "":
						print_string = "耶利米哀歌"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("哀","耶利米哀歌", chap_string, "1","lxx")
							default:
								print_string = Bible_print_string("哀","耶利米哀歌", chap_string, sec_string,"lxx")
						}
				}
			case "Ezek","Ezekiel","結","以西結","以西結書","Eze","eze","에스겔","エゼキエル書","エゼキエル","Ê-xê-chi-ên","Книга пророка Иезекииля":
				bible_short_name = "結"
				switch chap_string {
					case "":
						print_string = "以西結書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("結","以西結書", chap_string, "1","lxx")
							default:
								print_string = Bible_print_string("結","以西結書", chap_string, sec_string,"lxx")
						}
				}
			case "Dan","Daniel","但","但以理","但以理書","Da","da","Книга пророка Даниила","Đa-ni-ên","ダニエル書","ダニエル","다니엘":
				bible_short_name = "但"
				switch chap_string {
					case "":
						print_string = "但以理書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("但","但以理書", chap_string, "1","lxx")
							default:
								print_string = Bible_print_string("但","但以理書", chap_string, sec_string,"lxx")
						}
				}
			case "Hos","Hosea","何","何西","何西阿","何西阿書","Ho","ho","Книга пророка Осии","Ô-sê","ホセア書","ホセア","호세아":
				bible_short_name = "何"
				switch chap_string {
					case "":
						print_string = "何西阿書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("何","何西阿書", chap_string, "1","lxx")
							default:
								print_string = Bible_print_string("何","何西阿書", chap_string, sec_string,"lxx")
						}
				}
			case "Joel","珥","約珥","約珥書","Joe","joe","Книга пророка Иоиля","Giô-ên","ヨエル書","ヨエル","요엘":
				bible_short_name = "珥"
				switch chap_string {
					case "":
						print_string = "約珥書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("珥","約珥書", chap_string, "1","lxx")
							default:
								print_string = Bible_print_string("珥","約珥書", chap_string, sec_string,"lxx")
						}
				}
			case "Amos","摩","阿摩司書","Am","am","Книга пророка Амоса","A-mốt","アモス書","アモス","아모스":
				bible_short_name = "摩"
				switch chap_string {
					case "":
						print_string = "阿摩司書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("摩","阿摩司書", chap_string, "1","lxx")
							default:
								print_string = Bible_print_string("摩","阿摩司書", chap_string, sec_string,"lxx")
						}
				}
			case "Obad","Obadiah","俄","俄巴底亞","俄巴底亞書","Ob","ob","오바댜","オバデヤ書","オバデヤ","Áp-đia","Книга пророка Авдия":
				bible_short_name = "俄"
				switch chap_string {
					case "":
						print_string = "俄巴底亞書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("俄","俄巴底亞書", chap_string, "1","lxx")
							default:
								print_string = Bible_print_string("俄","俄巴底亞書", chap_string, sec_string,"lxx")
						}
				}
			case "Jon","Jonah","拿","約拿","約拿書","jon","요나","ヨナ書","ヨナ","Giô-na","Книга пророка Ионы":
				bible_short_name = "拿"
				switch chap_string {
					case "":
						print_string = "約拿書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("拿","約拿書", chap_string, "1","lxx")
							default:
								print_string = Bible_print_string("拿","約拿書", chap_string, sec_string,"lxx")
						}
				}
			case "Micah","彌","彌迦","彌迦書","Mic","mic","Книга пророка Михея","Mi-chê","ミカ書","ミカ","미가":
				bible_short_name = "彌"
				switch chap_string {
					case "":
						print_string = "彌迦書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("彌","彌迦書", chap_string, "1","lxx")
							default:
								print_string = Bible_print_string("彌","彌迦書", chap_string, sec_string,"lxx")
						}
				}
			case "Nah","Nahum","鴻","那鴻","那鴻書","Na","na","Книга пророка Наума","Na-hum","ナホム書","ナホム","나훔":
				bible_short_name = "鴻"
				switch chap_string {
					case "":
						print_string = "那鴻書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("鴻","那鴻書", chap_string, "1","lxx")
							default:
								print_string = Bible_print_string("鴻","那鴻書", chap_string, sec_string,"lxx")
						}
				}
			case "Habakkuk","哈","哈巴","哈巴谷","哈巴谷書","Hab","hab","Книга пророка Аввакума","Ha-ba-cúc","ハバクク書","ハバクク","ハバ","クク","ハバ書","하박국":
				bible_short_name = "哈"
				switch chap_string {
					case "":
						print_string = "哈巴谷書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("哈","哈巴谷書", chap_string, "1","lxx")
							default:
								print_string = Bible_print_string("哈","哈巴谷書", chap_string, sec_string,"lxx")
						}
				}
			case "Zeph","Zephaniah","番","西番雅","西番雅書","Zep","zep","스바냐","ゼパニヤ書","ゼパニヤ","Sô-phô-ni","Книга пророка Софонии":
				bible_short_name = "番"
				switch chap_string {
					case "":
						print_string = "西番雅書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("番","西番雅書", chap_string, "1","lxx")
							default:
								print_string = Bible_print_string("番","西番雅書", chap_string, sec_string,"lxx")
						}
				}
			case "Haggai","該","哈該","哈該書","Hag","hag","학개","ハガイ書","ハガイ","A-ghê","Книга пророка Аггея":
				bible_short_name = "該"
				switch chap_string {
					case "":
						print_string = "哈該書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("該","哈該書", chap_string, "1","lxx")
							default:
								print_string = Bible_print_string("該","哈該書", chap_string, sec_string,"lxx")
						}
				}
			case "Zech","Zechariah","亞","撒迦利亞","撒迦利亞書","Zec","zec","Книга пророка Захарии","Xa-cha-ri","스가랴","ゼカリヤ書","ゼカリヤ":
				bible_short_name = "亞"
				switch chap_string {
					case "":
						print_string = "撒迦利亞書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("亞","撒迦利亞書", chap_string, "1","lxx")
							default:
								print_string = Bible_print_string("亞","撒迦利亞書", chap_string, sec_string,"lxx")
						}
				}
			case "Malachi","瑪","","瑪拉","瑪拉基","瑪拉基書","Mal","mal","말라기","マラキ書","マラキ","Ma-la-chi","Книга пророка Малахии":
				bible_short_name = "瑪"
				switch chap_string {
					case "":
						print_string = "瑪拉基書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("瑪","瑪拉基書", chap_string, "1","lxx")
							default:
								print_string = Bible_print_string("瑪","瑪拉基書", chap_string, sec_string,"lxx")
						}
				}
			case "Matt","Matthew","太","馬太","馬太福音","Mt","mt","마태복음","マタイによる福音書","マタイ","マタイによる","Ma-thi-ơ","От Матфея святое благовествование":
				bible_short_name = "太"
				switch chap_string {
					case "":
						print_string = "馬太福音"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("太","馬太福音", chap_string, "1","lxx")
							default:
								print_string = Bible_print_string("太","馬太福音", chap_string, sec_string,"lxx")
						}
				}
			case "Mark","可","馬可","馬可福音","Mr","mr","マルコによる福音書","マルコ","マルコによる","마가복음","Mác","От Марка святое благовествование":
				bible_short_name = "可"
				switch chap_string {
					case "":
						print_string = "馬可福音"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("可","馬可福音", chap_string, "1","lxx")
							default:
								print_string = Bible_print_string("可","馬可福音", chap_string, sec_string,"lxx")
						}
				}
			case "Luke","路","路加","路加福音","Lu","lu","От Луки святое благовествование","Lu-ca","ルカによる福音書","ルカ","ルカによる","누가복음":
				bible_short_name = "路"
				switch chap_string {
					case "":
						print_string = "路加福音"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("路","路加福音", chap_string, "1","lxx")
							default:
								print_string = Bible_print_string("路","路加福音", chap_string, sec_string,"lxx")
						}
				}
			case "John","約","約翰","約翰福音","Joh","joh","От Иоанна святое благовествование","Giăng","ヨハネによる福音書","ヨハネ","ヨハネによる","요한복음":
				bible_short_name = "約"
				switch chap_string {
					case "":
						print_string = "約翰福音"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("約","約翰福音", chap_string, "1","lxx")
							default:
								print_string = Bible_print_string("約","約翰福音", chap_string, sec_string,"lxx")
						}
				}
			case "Acts","徒","使徒","使徒行傳","Ac","ac","Деяния святых апостолов","Công-vụ Các Sứ-đồ","使徒行伝","사도행전":
				bible_short_name = "徒"
				switch chap_string {
					case "":
						print_string = "使徒行傳"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("徒","使徒行傳", chap_string, "1","lxx")
							default:
								print_string = Bible_print_string("徒","使徒行傳", chap_string, sec_string,"lxx")
						}
				}
			case "Rom","Romans","羅","羅馬","羅馬書","Ro","ro","Послание к Римлянам","Rô-ma","ローマ","ローマ人への手紙","로마서":
				bible_short_name = "羅"
				switch chap_string {
					case "":
						print_string = "羅馬書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("羅","羅馬書", chap_string, "1","lxx")
							default:
								print_string = Bible_print_string("羅","羅馬書", chap_string, sec_string,"lxx")
						}
				}
			case "1 Cor","First Corinthians","林前","哥林多前","哥林多前書","1Co","1co","Первое послание к Коринфянам","1 Cô-rinh-tô","コリント人への第一の手紙","コリント一","コリント人への第一","고린도전서":
				bible_short_name = "林前"
				switch chap_string {
					case "":
						print_string = "哥林多前書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("林前","哥林多前書", chap_string, "1","lxx")
							default:
								print_string = Bible_print_string("林前","哥林多前書", chap_string, sec_string,"lxx")
						}
				}
			case "2 Cor","Second Corinthians","林後","哥林多後","哥林多後書","2Co","2co","Второе послание к Коринфянам","2 Cô-rinh-tô","コリント人への第二の手紙","コリント二","コリント人への第二の","고린도후서":
				bible_short_name = "林後"
				switch chap_string {
					case "":
						print_string = "哥林多後書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("林後","哥林多後書", chap_string, "1","lxx")
							default:
								print_string = Bible_print_string("林後","哥林多後書", chap_string, sec_string,"lxx")
						}
				}
			case "Gal","Galatians","加","加拉太","加拉太書","Ga","ga","Послание к Галатам","Ga-la-ti","ガラテヤ","ガラテヤ人への手紙","갈라디아서":
				bible_short_name = "加"
				switch chap_string {
					case "":
						print_string = "加拉太書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("加","加拉太書", chap_string, "1","lxx")
							default:
								print_string = Bible_print_string("加","加拉太書", chap_string, sec_string,"lxx")
						}
				}
			case "Ephesians","弗","以弗所","以弗所書","Eph","eph","Послание к Ефесянам","Ê-phê-sô","エペソ人への手紙","エペソ","エペソ人","エペソ人の手紙","에베소서":
				bible_short_name = "弗"
				switch chap_string {
					case "":
						print_string = "以弗所書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("弗","以弗所書", chap_string, "1","lxx")
							default:
								print_string = Bible_print_string("弗","以弗所書", chap_string, sec_string,"lxx")
						}
				}
			case "Phil","Philippians","腓","腓立","腓立比","腓立比書","Php","php","빌립보서","ピリピ","ピリピ人.ピリピ人への手紙","Послание к Филиппийцам","Phi-líp":
				bible_short_name = "腓"
				switch chap_string {
					case "":
						print_string = "腓立比書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("腓","腓立比書", chap_string, "1","lxx")
							default:
								print_string = Bible_print_string("腓","腓立比書", chap_string, sec_string,"lxx")
						}
				}
			case "Col","col","Colossians","西","歌羅西","歌羅","歌羅西書","Послание к Колоссянам","Cô-lô-se","コロサイ人への手紙","コロサイ","コロ","골로새서":
				bible_short_name = "西"
				switch chap_string {
					case "":
						print_string = "歌羅西書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("西","歌羅西書", chap_string, "1","lxx")
							default:
								print_string = Bible_print_string("西","歌羅西書", chap_string, sec_string,"lxx")
						}
				}
			case "1 Thess","First Thessalonians","帖前","帖撒羅尼迦前","帖撒羅尼迦前書","1Th","1th","데살로니가전서","テサロニケ人への第一の手紙","テサ一","テサロニケ一","1 Tê-sa-lô-ni-ca","Первое послание к Фессалоникийцам (Солунянам)":
				bible_short_name = "帖前"
				switch chap_string {
					case "":
						print_string = "帖撒羅尼迦前書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("帖前","帖撒羅尼迦前書", chap_string, "1","lxx")
							default:
								print_string = Bible_print_string("帖前","帖撒羅尼迦前書", chap_string, sec_string,"lxx")
						}
				}
			case "2 Thess","Second Thessalonians","帖後","帖撒羅尼迦後","帖撒羅尼迦後書","2Th","2th","데살로니가후서","テサロニケ人への第二の手紙","テサ二","テサロニケ二","2 Tê-sa-lô-ni-ca","Второе послание к Фессалоникийцам (Солунянам)":
				bible_short_name = "帖後"
				switch chap_string {
					case "":
						print_string = "帖撒羅尼迦後書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("帖後","帖撒羅尼迦後書", chap_string, "1","lxx")
							default:
								print_string = Bible_print_string("帖後","帖撒羅尼迦後書", chap_string, sec_string,"lxx")
						}
				}
			case "1 Tim","First Timothy","提前","提摩太前","提摩太前書","1Ti","1ti","Первое послание к Тимофею","1 Ti-mô-thê","テモテヘの第一の手紙","テモテ一","디모데전서":
				bible_short_name = "提前"
				switch chap_string {
					case "":
						print_string = "提摩太前書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("提前","提摩太前書", chap_string, "1","lxx")
							default:
								print_string = Bible_print_string("提前","提摩太前書", chap_string, sec_string,"lxx")
						}
				}
			case "2 Tim","Second Timothy","提後","提摩太後","提摩太後書","2Ti","2ti","Второе послание к Тимофею","2 Ti-mô-thê","テモテヘの第二の手紙","テモテ二","디모데후서":
				bible_short_name = "提後"
				switch chap_string {
					case "":
						print_string = "提摩太後書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("提後","提摩太後書", chap_string, "1","lxx")
							default:
								print_string = Bible_print_string("提後","提摩太後書", chap_string, sec_string,"lxx")
						}
				}
			case "Titus","多","提多","提多書","Tit","tit","Послание к Титу","Tít","テトスヘの手紙","テトス","디도서":
				bible_short_name = "多"
				switch chap_string {
					case "":
						print_string = "提多書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("多","提多書", chap_string, "1","lxx")
							default:
								print_string = Bible_print_string("多","提多書", chap_string, sec_string,"lxx")
						}
				}
			case "Philem","Philemon","門","腓利","腓利門","腓利門書","Phm","phm","Послание к Филимону","Phi-lê-môn","ピレモンヘの手紙","ピレモン","빌레몬서":
				bible_short_name = "門"
				switch chap_string {
					case "":
						print_string = "腓利門書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("門","腓利門書", chap_string, "1","lxx")
							default:
								print_string = Bible_print_string("門","腓利門書", chap_string, sec_string,"lxx")
						}
				}
			case "Heb","Hebrews","來","希伯來","希伯來書","heb","Послание к Евреям","Hê-bơ-rơ","ヘブル人への手紙","ヘブル","히브리서":
				bible_short_name = "來"
				switch chap_string {
					case "":
						print_string = "希伯來書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("來","希伯來書", chap_string, "1","lxx")
							default:
								print_string = Bible_print_string("來","希伯來書", chap_string, sec_string,"lxx")
						}
				}
			case "James","雅","雅各","雅各書","Jas","jas","Послание Иакова","Gia-cơ","ヤコブの手紙","ヤコブ","야고보서":
				bible_short_name = "雅"
				switch chap_string {
					case "":
						print_string = "雅各書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("雅","雅各書", chap_string, "1","lxx")
							default:
								print_string = Bible_print_string("雅","雅各書", chap_string, sec_string,"lxx")
						}
				}
			case "1 Pet","First Peter","彼前","彼得前","彼得前書","1Pe","1pe","Первое послание Петра","1 Phi-e-rơ","ペテロの第一の手紙","ペテロ一","베드로전서":
				bible_short_name = "彼前"
				switch chap_string {
					case "":
						print_string = "彼得前書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("彼前","彼得前書", chap_string, "1","lxx")
							default:
								print_string = Bible_print_string("彼前","彼得前書", chap_string, sec_string,"lxx")
						}
				}
			case "2 Pet","Second Peter","彼後","彼得後","彼得後書","2Pe","2pe","Второе послание Петра","2 Phi-e-rơ","ペテロの第二の手紙","ペテロ","베드로후서":
				bible_short_name = "彼後"
				switch chap_string {
					case "":
						print_string = "彼得後書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("彼後","彼得後書", chap_string, "1","lxx")
							default:
								print_string = Bible_print_string("彼後","彼得後書", chap_string, sec_string,"lxx")
						}
				}
			case "1 John","First John","約一","約翰一書","約翰1","約翰1書","1Jo","1jo","Первое послание Иоанна","1 Giăng","ヨハネの第一の手紙","ヨハネ一","요한일서":
				bible_short_name = "約一"
				switch chap_string {
					case "":
						print_string = "約翰一書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("約一","約翰一書", chap_string, "1","lxx")
							default:
								print_string = Bible_print_string("約一","約翰一書", chap_string, sec_string,"lxx")
						}
				}
			case "2 John","second John","約二","約翰二書","約翰2","約翰2書","2Jo","Второе послание Иоанна","2 Giăng","ヨハネの第二の手紙","ヨハネ二","요한2서":
				bible_short_name = "約二"
				switch chap_string {
					case "":
						print_string = "約翰二書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("約二","約翰二書", chap_string, "1","lxx")
							default:
								print_string = Bible_print_string("約二","約翰二書", chap_string, sec_string,"lxx")
						}
				}
			case "3 John","Third John","約三","約翰三書","約翰3","約翰3書","3Jo","3jo","Третье послание Иоанна","3 Giăng","ヨハネの第三の手紙","ヨハネ三","요한3서":
				bible_short_name = "約三"
				switch chap_string {
					case "":
						print_string = "約翰三書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("約三","約翰三書", chap_string, "1","lxx")
							default:
								print_string = Bible_print_string("約三","約翰三書", chap_string, sec_string,"lxx")
						}
				}
			case "Jude","猶","猶大","猶大書","jude","Послание Иуды","Giu-đe","ユダの手紙","ユダ","유다서":
				bible_short_name = "猶"
				switch chap_string {
					case "":
						print_string = "猶大書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("猶","猶大書", chap_string, "1","lxx")
							default:
								print_string = Bible_print_string("猶","猶大書", chap_string, sec_string,"lxx")
						}
				}
			case "Rev","Revelation","啟","啟示","啟示錄","Re","re","ｒｅ","Ｒｅ","rev","Откровение ап. Иоанна Богослова (Апокалипсис)","Khải-huyền","ヨハネの黙示録","黙示録","요한계시록":
				bible_short_name = "啟"
				switch chap_string {
					case "":
						print_string = "啟示錄"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("啟","啟示錄", chap_string, "1","lxx")
							default:
								print_string = Bible_print_string("啟","啟示錄", chap_string, sec_string,"lxx")
						}
				}
			default:
				print_string = "聖經"
				//print_string = "你是要找 " +  reg.ReplaceAllString(text, "$3") + " 對嗎？\n對不起，我還沒學呢...\n"
		}
	case "英文聖經ERV","erv","ERV","Erv","ＥＲＶ","Ｅｒｖ","ｅｒｖ":
		log.Print(reg.ReplaceAllString(text, "$3"))
		switch reg.ReplaceAllString(text, "$3") {
			case "Gen","Genesis","創","創世","創世紀","創世記","Ge","ge","gen","창세기","Sáng-thế Ký","Бытие":
				bible_short_name = "創"
				switch chap_string {
					case "":
						print_string = "創世記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("創","創世記", chap_string, "1","erv")
							default:
								print_string = Bible_print_string("創","創世記", chap_string, sec_string,"erv")
						}
				}
			case "ex","Ex","Exodus","埃及","出","出埃及","出埃及記","출애굽기","エジプト","出エジプト","出エジプト記","Xuất Ê-díp-tô Ký","Исход":
				bible_short_name = "出"
				switch chap_string {
					case "":
						print_string = "出埃及記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("出","出埃及記", chap_string, "1","erv")
							default:
								print_string = Bible_print_string("出","出埃及記", chap_string, sec_string,"erv")
						}
				}
			case "Lev","Leviticus","利","利未","利未記","Le","le","Левит","Lê-vi Ký","レビ記","レビ","레위기":
				bible_short_name = "利"
				switch chap_string {
					case "":
						print_string = "利未記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("利","利未記", chap_string, "1","erv")
							default:
								print_string = Bible_print_string("利","利未記", chap_string, sec_string,"erv")
						}
				}
			case "Num","Numbers","民","民數","民數記","Nu","nu","민수기","民数","民数記","Dân-số Ký","Числа":
				bible_short_name = "民"
				switch chap_string {
					case "":
						print_string = "民數記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("民","民數記", chap_string, "1","erv")
							default:
								print_string = Bible_print_string("民","民數記", chap_string, sec_string,"erv")
						}
				}
			case "Deut","Deuteronomy","申","申命","申命記","De","de","신명기","Phục-truyền Luật-lệ Ký","Второзаконие":
				bible_short_name = "申"
				switch chap_string {
					case "":
						print_string = "申命記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("申","申命記", chap_string, "1","erv")
							default:
								print_string = Bible_print_string("申","申命記", chap_string, sec_string,"erv")
						}
				}
			case "Josh","Joshua","約書亞","約書亞記","Jos","jos","여호수아","ヨシュア記","ヨシュア","Giô-suê","Книга Иисуса Навина":
				bible_short_name = "書"
				switch chap_string {
					case "":
						print_string = "約書亞記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("書","約書亞記", chap_string, "1","erv")
							default:
								print_string = Bible_print_string("書","約書亞記", chap_string, sec_string,"erv")
						}
				}
			case "Judg","Judges","士","士師","士師記","Jud","jud","jdg","Jdg","Книга Судей израилевых","Các Quan Xét","사사기":
				bible_short_name = "士"
				switch chap_string {
					case "":
						print_string = "士師記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("士","士師記", chap_string, "1","erv")
							default:
								print_string = Bible_print_string("士","士師記", chap_string, sec_string,"erv")
						}
				}
			case "Ruth","路得","路得記","Ru","ru","Rut","rut","룻기","ルツ","ルツ記","Ru-tơ","Книга Руфи":
				bible_short_name = "得"
				switch chap_string {
					case "":
						print_string = "路得記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("得","路得記", chap_string, "1","erv")
							default:
								print_string = Bible_print_string("得","路得記", chap_string, sec_string,"erv")
						}
				}
			case "1 Sam","First Samuel","撒上","撒母耳記上","1Sa","1sa","サムエル記上","サムエル上","サム上","사무엘상","1 Sa-mu-ên","Первая книга Царств":
				bible_short_name = "撒上"
				switch chap_string {
					case "":
						print_string = "撒母耳記上"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("撒上","撒母耳記上", chap_string, "1","erv")
							default:
								print_string = Bible_print_string("撒上","撒母耳記上", chap_string, sec_string,"erv")
						}
				}
			case "2 Sam","Second Samuel","撒下","撒母耳記下","2Sa","2sa","사무엘하","サムエル記下","サムエル下","サム下","2 Sa-mu-ên","Вторая книга Царств":
				bible_short_name = "撒下"
				switch chap_string {
					case "":
						print_string = "撒母耳記下"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("撒下","撒母耳記下", chap_string, "1","erv")
							default:
								print_string = Bible_print_string("撒下","撒母耳記下", chap_string, sec_string,"erv")
						}
				}
			case "1 Kin","First Kings","王上","列王上","列王紀上","列王記上","1Ki","1ki","열왕기상","Третья книга Царств","1 Các Vua":
				bible_short_name = "王上"
				switch chap_string {
					case "":
						print_string = "列王紀上"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("王上","列王紀上", chap_string, "1","erv")
							default:
								print_string = Bible_print_string("王上","列王紀上", chap_string, sec_string,"erv")
						}
				}
			case "2 Kin","Second Kings","王下","列王下","列王記下","列王紀下","2Ki","2ki","열왕기하","2 Các Vua","Четвертая книга Царств":
				bible_short_name = "王下"
				switch chap_string {
					case "":
						print_string = "列王紀下"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("王下","列王紀下", chap_string, "1","erv")
							default:
								print_string = Bible_print_string("王下","列王紀下", chap_string, sec_string,"erv")
						}
				}
			case "1 Chr","First Chronicles","歷上","代上","歷代志上","歷代上","1Ch","1ch","Первая книга Паралипоменон","1 Sử-ký","歴上","歴代上","歴代志上","역대상":
				bible_short_name = "代上"
				switch chap_string {
					case "":
						print_string = "歷代志上"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("代上","歷代志上", chap_string, "1","erv")
							default:
								print_string = Bible_print_string("代上","歷代志上", chap_string, sec_string,"erv")
						}
				}
			case "2 Chr","Second Chronicles","代下","歷下","歷代下","歷代志下","2Ch","2ch","역대하","歴代志下","歴代下","歴下","2 Sử-ký","Вторая книга Паралипоменон":
				bible_short_name = "代下"
				switch chap_string {
					case "":
						print_string = "歷代志下"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("代下","歷代志下", chap_string, "1","erv")
							default:
								print_string = Bible_print_string("代下","歷代志下", chap_string, sec_string,"erv")
						}
				}
			case "Ezra","拉","以斯拉","以斯拉記","Ezr","ezr","Первая книга Ездры","Ê-xơ-ra","エズラ","エズラ記","에스라":
				bible_short_name = "拉"
				switch chap_string {
					case "":
						print_string = "以斯拉記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("拉","以斯拉記", chap_string, "1","erv")
							default:
								print_string = Bible_print_string("拉","以斯拉記", chap_string, sec_string,"erv")
						}
				}
			case "Neh","Nehemiah","尼","尼希米","尼希米記","Ne","ne","느헤미야","ネヘミヤ書","ネヘミヤ","Nê-hê-mi","Книга Неемии":
				bible_short_name = "尼"
				switch chap_string {
					case "":
						print_string = "尼希米記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("尼","尼希米記", chap_string, "1","erv")
							default:
								print_string = Bible_print_string("尼","尼希米記", chap_string, sec_string,"erv")
						}
				}
			case "Esth","Esther","斯","以斯帖","以斯帖記","Es","est","Есфирь","Ê-xơ-tê","エステル","エステル記","에스더":
				bible_short_name = "斯"
				switch chap_string {
					case "":
						print_string = "以斯帖記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("斯","以斯帖記", chap_string, "1","erv")
							default:
								print_string = Bible_print_string("斯","以斯帖記", chap_string, sec_string,"erv")
						}
				}
			case "Job","job","伯","約伯","約伯記","Книга Иова","Gióp","ヨブ","ヨブ記","욥기":
				bible_short_name = "伯"
				switch chap_string {
					case "":
						print_string = "約伯記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("伯","約伯記", chap_string, "1","erv")
							default:
								print_string = Bible_print_string("伯","約伯記", chap_string, sec_string,"erv")
						}
				}
			case "Ps","Psalms","詩","詩篇","ps","시편","Thi-thiên","Псалтирь":
				bible_short_name = "詩"
				switch chap_string {
					case "":
						print_string = "詩篇"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("詩","詩篇", chap_string, "1","erv")
							default:
								print_string = Bible_print_string("詩","詩篇", chap_string, sec_string,"erv")
						}
				}
			case "Prov","Proverbs","箴","箴言","Pr","pr","Притчи Соломона","Châm-ngôn","잠언":
				bible_short_name = "箴"
				switch chap_string {
					case "":
						print_string = "箴言"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("箴","箴言", chap_string, "1","erv")
							default:
								print_string = Bible_print_string("箴","箴言", chap_string, sec_string,"erv")
						}
				}
			case "Eccl","Ecclesiastes","傳","傳道","傳道書","Ec","ec","Книга Екклезиаста","или Проповедника","Truyền-đạo","伝道の書","伝道","伝","伝道書","전도서":
				bible_short_name = "傳"
				switch chap_string {
					case "":
						print_string = "傳道書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("傳","傳道書", chap_string, "1","erv")
							default:
								print_string = Bible_print_string("傳","傳道書", chap_string, sec_string,"erv")
						}
				}
			case "Song","Song of Solomon","歌","雅歌","So","so","sng","Sng","Песнь песней Соломона","Nhã-ca","아가":
				bible_short_name = "歌"
				switch chap_string {
					case "":
						print_string = "雅歌"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("歌","雅歌", chap_string, "1","erv")
							default:
								print_string = Bible_print_string("歌","雅歌", chap_string, sec_string,"erv")
						}
				}
			case "Is","Isaiah","賽","以賽","以賽亞","以賽亞書","Isa","isa","Книга пророка Исаии","Ê-sai","イザヤ書","イザヤ","이사야":
				bible_short_name = "賽"
				switch chap_string {
					case "":
						print_string = "以賽亞書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("賽","以賽亞書", chap_string, "1","erv")
							default:
								print_string = Bible_print_string("賽","以賽亞書", chap_string, sec_string,"erv")
						}
				}
			case "Jer","Jeremiah","耶","耶利米","耶利米書","jer","예레미야","エレミヤ","エレミヤ書","Giê-rê-mi","Книга пророка Иеремии":
				bible_short_name = "耶"
				switch chap_string {
					case "":
						print_string = "耶利米書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("耶","耶利米書", chap_string, "1","erv")
							default:
								print_string = Bible_print_string("耶","耶利米書", chap_string, sec_string,"erv")
						}
				}
			case "Lam","Lamentations","哀","哀歌","耶利米哀歌","La","lam","예레미야애가","Ca-thương","Плач Иеремии":
				bible_short_name = "哀"
				switch chap_string {
					case "":
						print_string = "耶利米哀歌"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("哀","耶利米哀歌", chap_string, "1","erv")
							default:
								print_string = Bible_print_string("哀","耶利米哀歌", chap_string, sec_string,"erv")
						}
				}
			case "Ezek","Ezekiel","結","以西結","以西結書","Eze","eze","에스겔","エゼキエル書","エゼキエル","Ê-xê-chi-ên","Книга пророка Иезекииля":
				bible_short_name = "結"
				switch chap_string {
					case "":
						print_string = "以西結書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("結","以西結書", chap_string, "1","erv")
							default:
								print_string = Bible_print_string("結","以西結書", chap_string, sec_string,"erv")
						}
				}
			case "Dan","Daniel","但","但以理","但以理書","Da","da","Книга пророка Даниила","Đa-ni-ên","ダニエル書","ダニエル","다니엘":
				bible_short_name = "但"
				switch chap_string {
					case "":
						print_string = "但以理書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("但","但以理書", chap_string, "1","erv")
							default:
								print_string = Bible_print_string("但","但以理書", chap_string, sec_string,"erv")
						}
				}
			case "Hos","Hosea","何","何西","何西阿","何西阿書","Ho","ho","Книга пророка Осии","Ô-sê","ホセア書","ホセア","호세아":
				bible_short_name = "何"
				switch chap_string {
					case "":
						print_string = "何西阿書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("何","何西阿書", chap_string, "1","erv")
							default:
								print_string = Bible_print_string("何","何西阿書", chap_string, sec_string,"erv")
						}
				}
			case "Joel","珥","約珥","約珥書","Joe","joe","Книга пророка Иоиля","Giô-ên","ヨエル書","ヨエル","요엘":
				bible_short_name = "珥"
				switch chap_string {
					case "":
						print_string = "約珥書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("珥","約珥書", chap_string, "1","erv")
							default:
								print_string = Bible_print_string("珥","約珥書", chap_string, sec_string,"erv")
						}
				}
			case "Amos","摩","阿摩司書","Am","am","Книга пророка Амоса","A-mốt","アモス書","アモス","아모스":
				bible_short_name = "摩"
				switch chap_string {
					case "":
						print_string = "阿摩司書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("摩","阿摩司書", chap_string, "1","erv")
							default:
								print_string = Bible_print_string("摩","阿摩司書", chap_string, sec_string,"erv")
						}
				}
			case "Obad","Obadiah","俄","俄巴底亞","俄巴底亞書","Ob","ob","오바댜","オバデヤ書","オバデヤ","Áp-đia","Книга пророка Авдия":
				bible_short_name = "俄"
				switch chap_string {
					case "":
						print_string = "俄巴底亞書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("俄","俄巴底亞書", chap_string, "1","erv")
							default:
								print_string = Bible_print_string("俄","俄巴底亞書", chap_string, sec_string,"erv")
						}
				}
			case "Jon","Jonah","拿","約拿","約拿書","jon","요나","ヨナ書","ヨナ","Giô-na","Книга пророка Ионы":
				bible_short_name = "拿"
				switch chap_string {
					case "":
						print_string = "約拿書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("拿","約拿書", chap_string, "1","erv")
							default:
								print_string = Bible_print_string("拿","約拿書", chap_string, sec_string,"erv")
						}
				}
			case "Micah","彌","彌迦","彌迦書","Mic","mic","Книга пророка Михея","Mi-chê","ミカ書","ミカ","미가":
				bible_short_name = "彌"
				switch chap_string {
					case "":
						print_string = "彌迦書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("彌","彌迦書", chap_string, "1","erv")
							default:
								print_string = Bible_print_string("彌","彌迦書", chap_string, sec_string,"erv")
						}
				}
			case "Nah","Nahum","鴻","那鴻","那鴻書","Na","na","Книга пророка Наума","Na-hum","ナホム書","ナホム","나훔":
				bible_short_name = "鴻"
				switch chap_string {
					case "":
						print_string = "那鴻書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("鴻","那鴻書", chap_string, "1","erv")
							default:
								print_string = Bible_print_string("鴻","那鴻書", chap_string, sec_string,"erv")
						}
				}
			case "Habakkuk","哈","哈巴","哈巴谷","哈巴谷書","Hab","hab","Книга пророка Аввакума","Ha-ba-cúc","ハバクク書","ハバクク","ハバ","クク","ハバ書","하박국":
				bible_short_name = "哈"
				switch chap_string {
					case "":
						print_string = "哈巴谷書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("哈","哈巴谷書", chap_string, "1","erv")
							default:
								print_string = Bible_print_string("哈","哈巴谷書", chap_string, sec_string,"erv")
						}
				}
			case "Zeph","Zephaniah","番","西番雅","西番雅書","Zep","zep","스바냐","ゼパニヤ書","ゼパニヤ","Sô-phô-ni","Книга пророка Софонии":
				bible_short_name = "番"
				switch chap_string {
					case "":
						print_string = "西番雅書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("番","西番雅書", chap_string, "1","erv")
							default:
								print_string = Bible_print_string("番","西番雅書", chap_string, sec_string,"erv")
						}
				}
			case "Haggai","該","哈該","哈該書","Hag","hag","학개","ハガイ書","ハガイ","A-ghê","Книга пророка Аггея":
				bible_short_name = "該"
				switch chap_string {
					case "":
						print_string = "哈該書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("該","哈該書", chap_string, "1","erv")
							default:
								print_string = Bible_print_string("該","哈該書", chap_string, sec_string,"erv")
						}
				}
			case "Zech","Zechariah","亞","撒迦利亞","撒迦利亞書","Zec","zec","Книга пророка Захарии","Xa-cha-ri","스가랴","ゼカリヤ書","ゼカリヤ":
				bible_short_name = "亞"
				switch chap_string {
					case "":
						print_string = "撒迦利亞書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("亞","撒迦利亞書", chap_string, "1","erv")
							default:
								print_string = Bible_print_string("亞","撒迦利亞書", chap_string, sec_string,"erv")
						}
				}
			case "Malachi","瑪","","瑪拉","瑪拉基","瑪拉基書","Mal","mal","말라기","マラキ書","マラキ","Ma-la-chi","Книга пророка Малахии":
				bible_short_name = "瑪"
				switch chap_string {
					case "":
						print_string = "瑪拉基書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("瑪","瑪拉基書", chap_string, "1","erv")
							default:
								print_string = Bible_print_string("瑪","瑪拉基書", chap_string, sec_string,"erv")
						}
				}
			case "Matt","Matthew","太","馬太","馬太福音","Mt","mt","마태복음","マタイによる福音書","マタイ","マタイによる","Ma-thi-ơ","От Матфея святое благовествование":
				bible_short_name = "太"
				switch chap_string {
					case "":
						print_string = "馬太福音"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("太","馬太福音", chap_string, "1","erv")
							default:
								print_string = Bible_print_string("太","馬太福音", chap_string, sec_string,"erv")
						}
				}
			case "Mark","可","馬可","馬可福音","Mr","mr","マルコによる福音書","マルコ","マルコによる","마가복음","Mác","От Марка святое благовествование":
				bible_short_name = "可"
				switch chap_string {
					case "":
						print_string = "馬可福音"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("可","馬可福音", chap_string, "1","erv")
							default:
								print_string = Bible_print_string("可","馬可福音", chap_string, sec_string,"erv")
						}
				}
			case "Luke","路","路加","路加福音","Lu","lu","От Луки святое благовествование","Lu-ca","ルカによる福音書","ルカ","ルカによる","누가복음":
				bible_short_name = "路"
				switch chap_string {
					case "":
						print_string = "路加福音"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("路","路加福音", chap_string, "1","erv")
							default:
								print_string = Bible_print_string("路","路加福音", chap_string, sec_string,"erv")
						}
				}
			case "John","約","約翰","約翰福音","Joh","joh","От Иоанна святое благовествование","Giăng","ヨハネによる福音書","ヨハネ","ヨハネによる","요한복음":
				bible_short_name = "約"
				switch chap_string {
					case "":
						print_string = "約翰福音"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("約","約翰福音", chap_string, "1","erv")
							default:
								print_string = Bible_print_string("約","約翰福音", chap_string, sec_string,"erv")
						}
				}
			case "Acts","徒","使徒","使徒行傳","Ac","ac","Деяния святых апостолов","Công-vụ Các Sứ-đồ","使徒行伝","사도행전":
				bible_short_name = "徒"
				switch chap_string {
					case "":
						print_string = "使徒行傳"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("徒","使徒行傳", chap_string, "1","erv")
							default:
								print_string = Bible_print_string("徒","使徒行傳", chap_string, sec_string,"erv")
						}
				}
			case "Rom","Romans","羅","羅馬","羅馬書","Ro","ro","Послание к Римлянам","Rô-ma","ローマ","ローマ人への手紙","로마서":
				bible_short_name = "羅"
				switch chap_string {
					case "":
						print_string = "羅馬書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("羅","羅馬書", chap_string, "1","erv")
							default:
								print_string = Bible_print_string("羅","羅馬書", chap_string, sec_string,"erv")
						}
				}
			case "1 Cor","First Corinthians","林前","哥林多前","哥林多前書","1Co","1co","Первое послание к Коринфянам","1 Cô-rinh-tô","コリント人への第一の手紙","コリント一","コリント人への第一","고린도전서":
				bible_short_name = "林前"
				switch chap_string {
					case "":
						print_string = "哥林多前書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("林前","哥林多前書", chap_string, "1","erv")
							default:
								print_string = Bible_print_string("林前","哥林多前書", chap_string, sec_string,"erv")
						}
				}
			case "2 Cor","Second Corinthians","林後","哥林多後","哥林多後書","2Co","2co","Второе послание к Коринфянам","2 Cô-rinh-tô","コリント人への第二の手紙","コリント二","コリント人への第二の","고린도후서":
				bible_short_name = "林後"
				switch chap_string {
					case "":
						print_string = "哥林多後書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("林後","哥林多後書", chap_string, "1","erv")
							default:
								print_string = Bible_print_string("林後","哥林多後書", chap_string, sec_string,"erv")
						}
				}
			case "Gal","Galatians","加","加拉太","加拉太書","Ga","ga","Послание к Галатам","Ga-la-ti","ガラテヤ","ガラテヤ人への手紙","갈라디아서":
				bible_short_name = "加"
				switch chap_string {
					case "":
						print_string = "加拉太書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("加","加拉太書", chap_string, "1","erv")
							default:
								print_string = Bible_print_string("加","加拉太書", chap_string, sec_string,"erv")
						}
				}
			case "Ephesians","弗","以弗所","以弗所書","Eph","eph","Послание к Ефесянам","Ê-phê-sô","エペソ人への手紙","エペソ","エペソ人","エペソ人の手紙","에베소서":
				bible_short_name = "弗"
				switch chap_string {
					case "":
						print_string = "以弗所書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("弗","以弗所書", chap_string, "1","erv")
							default:
								print_string = Bible_print_string("弗","以弗所書", chap_string, sec_string,"erv")
						}
				}
			case "Phil","Philippians","腓","腓立","腓立比","腓立比書","Php","php","빌립보서","ピリピ","ピリピ人.ピリピ人への手紙","Послание к Филиппийцам","Phi-líp":
				bible_short_name = "腓"
				switch chap_string {
					case "":
						print_string = "腓立比書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("腓","腓立比書", chap_string, "1","erv")
							default:
								print_string = Bible_print_string("腓","腓立比書", chap_string, sec_string,"erv")
						}
				}
			case "Col","col","Colossians","西","歌羅西","歌羅","歌羅西書","Послание к Колоссянам","Cô-lô-se","コロサイ人への手紙","コロサイ","コロ","골로새서":
				bible_short_name = "西"
				switch chap_string {
					case "":
						print_string = "歌羅西書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("西","歌羅西書", chap_string, "1","erv")
							default:
								print_string = Bible_print_string("西","歌羅西書", chap_string, sec_string,"erv")
						}
				}
			case "1 Thess","First Thessalonians","帖前","帖撒羅尼迦前","帖撒羅尼迦前書","1Th","1th","데살로니가전서","テサロニケ人への第一の手紙","テサ一","テサロニケ一","1 Tê-sa-lô-ni-ca","Первое послание к Фессалоникийцам (Солунянам)":
				bible_short_name = "帖前"
				switch chap_string {
					case "":
						print_string = "帖撒羅尼迦前書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("帖前","帖撒羅尼迦前書", chap_string, "1","erv")
							default:
								print_string = Bible_print_string("帖前","帖撒羅尼迦前書", chap_string, sec_string,"erv")
						}
				}
			case "2 Thess","Second Thessalonians","帖後","帖撒羅尼迦後","帖撒羅尼迦後書","2Th","2th","데살로니가후서","テサロニケ人への第二の手紙","テサ二","テサロニケ二","2 Tê-sa-lô-ni-ca","Второе послание к Фессалоникийцам (Солунянам)":
				bible_short_name = "帖後"
				switch chap_string {
					case "":
						print_string = "帖撒羅尼迦後書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("帖後","帖撒羅尼迦後書", chap_string, "1","erv")
							default:
								print_string = Bible_print_string("帖後","帖撒羅尼迦後書", chap_string, sec_string,"erv")
						}
				}
			case "1 Tim","First Timothy","提前","提摩太前","提摩太前書","1Ti","1ti","Первое послание к Тимофею","1 Ti-mô-thê","テモテヘの第一の手紙","テモテ一","디모데전서":
				bible_short_name = "提前"
				switch chap_string {
					case "":
						print_string = "提摩太前書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("提前","提摩太前書", chap_string, "1","erv")
							default:
								print_string = Bible_print_string("提前","提摩太前書", chap_string, sec_string,"erv")
						}
				}
			case "2 Tim","Second Timothy","提後","提摩太後","提摩太後書","2Ti","2ti","Второе послание к Тимофею","2 Ti-mô-thê","テモテヘの第二の手紙","テモテ二","디모데후서":
				bible_short_name = "提後"
				switch chap_string {
					case "":
						print_string = "提摩太後書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("提後","提摩太後書", chap_string, "1","erv")
							default:
								print_string = Bible_print_string("提後","提摩太後書", chap_string, sec_string,"erv")
						}
				}
			case "Titus","多","提多","提多書","Tit","tit","Послание к Титу","Tít","テトスヘの手紙","テトス","디도서":
				bible_short_name = "多"
				switch chap_string {
					case "":
						print_string = "提多書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("多","提多書", chap_string, "1","erv")
							default:
								print_string = Bible_print_string("多","提多書", chap_string, sec_string,"erv")
						}
				}
			case "Philem","Philemon","門","腓利","腓利門","腓利門書","Phm","phm","Послание к Филимону","Phi-lê-môn","ピレモンヘの手紙","ピレモン","빌레몬서":
				bible_short_name = "門"
				switch chap_string {
					case "":
						print_string = "腓利門書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("門","腓利門書", chap_string, "1","erv")
							default:
								print_string = Bible_print_string("門","腓利門書", chap_string, sec_string,"erv")
						}
				}
			case "Heb","Hebrews","來","希伯來","希伯來書","heb","Послание к Евреям","Hê-bơ-rơ","ヘブル人への手紙","ヘブル","히브리서":
				bible_short_name = "來"
				switch chap_string {
					case "":
						print_string = "希伯來書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("來","希伯來書", chap_string, "1","erv")
							default:
								print_string = Bible_print_string("來","希伯來書", chap_string, sec_string,"erv")
						}
				}
			case "James","雅","雅各","雅各書","Jas","jas","Послание Иакова","Gia-cơ","ヤコブの手紙","ヤコブ","야고보서":
				bible_short_name = "雅"
				switch chap_string {
					case "":
						print_string = "雅各書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("雅","雅各書", chap_string, "1","erv")
							default:
								print_string = Bible_print_string("雅","雅各書", chap_string, sec_string,"erv")
						}
				}
			case "1 Pet","First Peter","彼前","彼得前","彼得前書","1Pe","1pe","Первое послание Петра","1 Phi-e-rơ","ペテロの第一の手紙","ペテロ一","베드로전서":
				bible_short_name = "彼前"
				switch chap_string {
					case "":
						print_string = "彼得前書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("彼前","彼得前書", chap_string, "1","erv")
							default:
								print_string = Bible_print_string("彼前","彼得前書", chap_string, sec_string,"erv")
						}
				}
			case "2 Pet","Second Peter","彼後","彼得後","彼得後書","2Pe","2pe","Второе послание Петра","2 Phi-e-rơ","ペテロの第二の手紙","ペテロ","베드로후서":
				bible_short_name = "彼後"
				switch chap_string {
					case "":
						print_string = "彼得後書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("彼後","彼得後書", chap_string, "1","erv")
							default:
								print_string = Bible_print_string("彼後","彼得後書", chap_string, sec_string,"erv")
						}
				}
			case "1 John","First John","約一","約翰一書","約翰1","約翰1書","1Jo","1jo","Первое послание Иоанна","1 Giăng","ヨハネの第一の手紙","ヨハネ一","요한일서":
				bible_short_name = "約一"
				switch chap_string {
					case "":
						print_string = "約翰一書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("約一","約翰一書", chap_string, "1","erv")
							default:
								print_string = Bible_print_string("約一","約翰一書", chap_string, sec_string,"erv")
						}
				}
			case "2 John","second John","約二","約翰二書","約翰2","約翰2書","2Jo","Второе послание Иоанна","2 Giăng","ヨハネの第二の手紙","ヨハネ二","요한2서":
				bible_short_name = "約二"
				switch chap_string {
					case "":
						print_string = "約翰二書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("約二","約翰二書", chap_string, "1","erv")
							default:
								print_string = Bible_print_string("約二","約翰二書", chap_string, sec_string,"erv")
						}
				}
			case "3 John","Third John","約三","約翰三書","約翰3","約翰3書","3Jo","3jo","Третье послание Иоанна","3 Giăng","ヨハネの第三の手紙","ヨハネ三","요한3서":
				bible_short_name = "約三"
				switch chap_string {
					case "":
						print_string = "約翰三書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("約三","約翰三書", chap_string, "1","erv")
							default:
								print_string = Bible_print_string("約三","約翰三書", chap_string, sec_string,"erv")
						}
				}
			case "Jude","猶","猶大","猶大書","jude","Послание Иуды","Giu-đe","ユダの手紙","ユダ","유다서":
				bible_short_name = "猶"
				switch chap_string {
					case "":
						print_string = "猶大書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("猶","猶大書", chap_string, "1","erv")
							default:
								print_string = Bible_print_string("猶","猶大書", chap_string, sec_string,"erv")
						}
				}
			case "Rev","Revelation","啟","啟示","啟示錄","Re","re","ｒｅ","Ｒｅ","rev","Откровение ап. Иоанна Богослова (Апокалипсис)","Khải-huyền","ヨハネの黙示録","黙示録","요한계시록":
				bible_short_name = "啟"
				switch chap_string {
					case "":
						print_string = "啟示錄"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("啟","啟示錄", chap_string, "1","erv")
							default:
								print_string = Bible_print_string("啟","啟示錄", chap_string, sec_string,"erv")
						}
				}
			default:
				print_string = "聖經"
				//print_string = "你是要找 " +  reg.ReplaceAllString(text, "$3") + " 對嗎？\n對不起，我還沒學呢...\n"
		}
	case "英文聖經Darby","darby","DARBY","Ｄａｒｂｙ","ＤＡＲＢＹ","ｄａｒｂｙ":
		log.Print(reg.ReplaceAllString(text, "$3"))
		switch reg.ReplaceAllString(text, "$3") {
			case "Gen","Genesis","創","創世","創世紀","創世記","Ge","ge","gen","창세기","Sáng-thế Ký","Бытие":
				bible_short_name = "創"
				switch chap_string {
					case "":
						print_string = "創世記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("創","創世記", chap_string, "1","darby")
							default:
								print_string = Bible_print_string("創","創世記", chap_string, sec_string,"darby")
						}
				}
			case "ex","Ex","Exodus","埃及","出","出埃及","出埃及記","출애굽기","エジプト","出エジプト","出エジプト記","Xuất Ê-díp-tô Ký","Исход":
				bible_short_name = "出"
				switch chap_string {
					case "":
						print_string = "出埃及記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("出","出埃及記", chap_string, "1","darby")
							default:
								print_string = Bible_print_string("出","出埃及記", chap_string, sec_string,"darby")
						}
				}
			case "Lev","Leviticus","利","利未","利未記","Le","le","Левит","Lê-vi Ký","レビ記","レビ","레위기":
				bible_short_name = "利"
				switch chap_string {
					case "":
						print_string = "利未記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("利","利未記", chap_string, "1","darby")
							default:
								print_string = Bible_print_string("利","利未記", chap_string, sec_string,"darby")
						}
				}
			case "Num","Numbers","民","民數","民數記","Nu","nu","민수기","民数","民数記","Dân-số Ký","Числа":
				bible_short_name = "民"
				switch chap_string {
					case "":
						print_string = "民數記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("民","民數記", chap_string, "1","darby")
							default:
								print_string = Bible_print_string("民","民數記", chap_string, sec_string,"darby")
						}
				}
			case "Deut","Deuteronomy","申","申命","申命記","De","de","신명기","Phục-truyền Luật-lệ Ký","Второзаконие":
				bible_short_name = "申"
				switch chap_string {
					case "":
						print_string = "申命記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("申","申命記", chap_string, "1","darby")
							default:
								print_string = Bible_print_string("申","申命記", chap_string, sec_string,"darby")
						}
				}
			case "Josh","Joshua","約書亞","約書亞記","Jos","jos","여호수아","ヨシュア記","ヨシュア","Giô-suê","Книга Иисуса Навина":
				bible_short_name = "書"
				switch chap_string {
					case "":
						print_string = "約書亞記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("書","約書亞記", chap_string, "1","darby")
							default:
								print_string = Bible_print_string("書","約書亞記", chap_string, sec_string,"darby")
						}
				}
			case "Judg","Judges","士","士師","士師記","Jud","jud","jdg","Jdg","Книга Судей израилевых","Các Quan Xét","사사기":
				bible_short_name = "士"
				switch chap_string {
					case "":
						print_string = "士師記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("士","士師記", chap_string, "1","darby")
							default:
								print_string = Bible_print_string("士","士師記", chap_string, sec_string,"darby")
						}
				}
			case "Ruth","路得","路得記","Ru","ru","Rut","rut","룻기","ルツ","ルツ記","Ru-tơ","Книга Руфи":
				bible_short_name = "得"
				switch chap_string {
					case "":
						print_string = "路得記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("得","路得記", chap_string, "1","darby")
							default:
								print_string = Bible_print_string("得","路得記", chap_string, sec_string,"darby")
						}
				}
			case "1 Sam","First Samuel","撒上","撒母耳記上","1Sa","1sa","サムエル記上","サムエル上","サム上","사무엘상","1 Sa-mu-ên","Первая книга Царств":
				bible_short_name = "撒上"
				switch chap_string {
					case "":
						print_string = "撒母耳記上"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("撒上","撒母耳記上", chap_string, "1","darby")
							default:
								print_string = Bible_print_string("撒上","撒母耳記上", chap_string, sec_string,"darby")
						}
				}
			case "2 Sam","Second Samuel","撒下","撒母耳記下","2Sa","2sa","사무엘하","サムエル記下","サムエル下","サム下","2 Sa-mu-ên","Вторая книга Царств":
				bible_short_name = "撒下"
				switch chap_string {
					case "":
						print_string = "撒母耳記下"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("撒下","撒母耳記下", chap_string, "1","darby")
							default:
								print_string = Bible_print_string("撒下","撒母耳記下", chap_string, sec_string,"darby")
						}
				}
			case "1 Kin","First Kings","王上","列王上","列王紀上","列王記上","1Ki","1ki","열왕기상","Третья книга Царств","1 Các Vua":
				bible_short_name = "王上"
				switch chap_string {
					case "":
						print_string = "列王紀上"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("王上","列王紀上", chap_string, "1","darby")
							default:
								print_string = Bible_print_string("王上","列王紀上", chap_string, sec_string,"darby")
						}
				}
			case "2 Kin","Second Kings","王下","列王下","列王記下","列王紀下","2Ki","2ki","열왕기하","2 Các Vua","Четвертая книга Царств":
				bible_short_name = "王下"
				switch chap_string {
					case "":
						print_string = "列王紀下"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("王下","列王紀下", chap_string, "1","darby")
							default:
								print_string = Bible_print_string("王下","列王紀下", chap_string, sec_string,"darby")
						}
				}
			case "1 Chr","First Chronicles","歷上","代上","歷代志上","歷代上","1Ch","1ch","Первая книга Паралипоменон","1 Sử-ký","歴上","歴代上","歴代志上","역대상":
				bible_short_name = "代上"
				switch chap_string {
					case "":
						print_string = "歷代志上"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("代上","歷代志上", chap_string, "1","darby")
							default:
								print_string = Bible_print_string("代上","歷代志上", chap_string, sec_string,"darby")
						}
				}
			case "2 Chr","Second Chronicles","代下","歷下","歷代下","歷代志下","2Ch","2ch","역대하","歴代志下","歴代下","歴下","2 Sử-ký","Вторая книга Паралипоменон":
				bible_short_name = "代下"
				switch chap_string {
					case "":
						print_string = "歷代志下"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("代下","歷代志下", chap_string, "1","darby")
							default:
								print_string = Bible_print_string("代下","歷代志下", chap_string, sec_string,"darby")
						}
				}
			case "Ezra","拉","以斯拉","以斯拉記","Ezr","ezr","Первая книга Ездры","Ê-xơ-ra","エズラ","エズラ記","에스라":
				bible_short_name = "拉"
				switch chap_string {
					case "":
						print_string = "以斯拉記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("拉","以斯拉記", chap_string, "1","darby")
							default:
								print_string = Bible_print_string("拉","以斯拉記", chap_string, sec_string,"darby")
						}
				}
			case "Neh","Nehemiah","尼","尼希米","尼希米記","Ne","ne","느헤미야","ネヘミヤ書","ネヘミヤ","Nê-hê-mi","Книга Неемии":
				bible_short_name = "尼"
				switch chap_string {
					case "":
						print_string = "尼希米記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("尼","尼希米記", chap_string, "1","darby")
							default:
								print_string = Bible_print_string("尼","尼希米記", chap_string, sec_string,"darby")
						}
				}
			case "Esth","Esther","斯","以斯帖","以斯帖記","Es","est","Есфирь","Ê-xơ-tê","エステル","エステル記","에스더":
				bible_short_name = "斯"
				switch chap_string {
					case "":
						print_string = "以斯帖記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("斯","以斯帖記", chap_string, "1","darby")
							default:
								print_string = Bible_print_string("斯","以斯帖記", chap_string, sec_string,"darby")
						}
				}
			case "Job","job","伯","約伯","約伯記","Книга Иова","Gióp","ヨブ","ヨブ記","욥기":
				bible_short_name = "伯"
				switch chap_string {
					case "":
						print_string = "約伯記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("伯","約伯記", chap_string, "1","darby")
							default:
								print_string = Bible_print_string("伯","約伯記", chap_string, sec_string,"darby")
						}
				}
			case "Ps","Psalms","詩","詩篇","ps","시편","Thi-thiên","Псалтирь":
				bible_short_name = "詩"
				switch chap_string {
					case "":
						print_string = "詩篇"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("詩","詩篇", chap_string, "1","darby")
							default:
								print_string = Bible_print_string("詩","詩篇", chap_string, sec_string,"darby")
						}
				}
			case "Prov","Proverbs","箴","箴言","Pr","pr","Притчи Соломона","Châm-ngôn","잠언":
				bible_short_name = "箴"
				switch chap_string {
					case "":
						print_string = "箴言"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("箴","箴言", chap_string, "1","darby")
							default:
								print_string = Bible_print_string("箴","箴言", chap_string, sec_string,"darby")
						}
				}
			case "Eccl","Ecclesiastes","傳","傳道","傳道書","Ec","ec","Книга Екклезиаста","или Проповедника","Truyền-đạo","伝道の書","伝道","伝","伝道書","전도서":
				bible_short_name = "傳"
				switch chap_string {
					case "":
						print_string = "傳道書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("傳","傳道書", chap_string, "1","darby")
							default:
								print_string = Bible_print_string("傳","傳道書", chap_string, sec_string,"darby")
						}
				}
			case "Song","Song of Solomon","歌","雅歌","So","so","sng","Sng","Песнь песней Соломона","Nhã-ca","아가":
				bible_short_name = "歌"
				switch chap_string {
					case "":
						print_string = "雅歌"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("歌","雅歌", chap_string, "1","darby")
							default:
								print_string = Bible_print_string("歌","雅歌", chap_string, sec_string,"darby")
						}
				}
			case "Is","Isaiah","賽","以賽","以賽亞","以賽亞書","Isa","isa","Книга пророка Исаии","Ê-sai","イザヤ書","イザヤ","이사야":
				bible_short_name = "賽"
				switch chap_string {
					case "":
						print_string = "以賽亞書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("賽","以賽亞書", chap_string, "1","darby")
							default:
								print_string = Bible_print_string("賽","以賽亞書", chap_string, sec_string,"darby")
						}
				}
			case "Jer","Jeremiah","耶","耶利米","耶利米書","jer","예레미야","エレミヤ","エレミヤ書","Giê-rê-mi","Книга пророка Иеремии":
				bible_short_name = "耶"
				switch chap_string {
					case "":
						print_string = "耶利米書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("耶","耶利米書", chap_string, "1","darby")
							default:
								print_string = Bible_print_string("耶","耶利米書", chap_string, sec_string,"darby")
						}
				}
			case "Lam","Lamentations","哀","哀歌","耶利米哀歌","La","lam","예레미야애가","Ca-thương","Плач Иеремии":
				bible_short_name = "哀"
				switch chap_string {
					case "":
						print_string = "耶利米哀歌"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("哀","耶利米哀歌", chap_string, "1","darby")
							default:
								print_string = Bible_print_string("哀","耶利米哀歌", chap_string, sec_string,"darby")
						}
				}
			case "Ezek","Ezekiel","結","以西結","以西結書","Eze","eze","에스겔","エゼキエル書","エゼキエル","Ê-xê-chi-ên","Книга пророка Иезекииля":
				bible_short_name = "結"
				switch chap_string {
					case "":
						print_string = "以西結書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("結","以西結書", chap_string, "1","darby")
							default:
								print_string = Bible_print_string("結","以西結書", chap_string, sec_string,"darby")
						}
				}
			case "Dan","Daniel","但","但以理","但以理書","Da","da","Книга пророка Даниила","Đa-ni-ên","ダニエル書","ダニエル","다니엘":
				bible_short_name = "但"
				switch chap_string {
					case "":
						print_string = "但以理書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("但","但以理書", chap_string, "1","darby")
							default:
								print_string = Bible_print_string("但","但以理書", chap_string, sec_string,"darby")
						}
				}
			case "Hos","Hosea","何","何西","何西阿","何西阿書","Ho","ho","Книга пророка Осии","Ô-sê","ホセア書","ホセア","호세아":
				bible_short_name = "何"
				switch chap_string {
					case "":
						print_string = "何西阿書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("何","何西阿書", chap_string, "1","darby")
							default:
								print_string = Bible_print_string("何","何西阿書", chap_string, sec_string,"darby")
						}
				}
			case "Joel","珥","約珥","約珥書","Joe","joe","Книга пророка Иоиля","Giô-ên","ヨエル書","ヨエル","요엘":
				bible_short_name = "珥"
				switch chap_string {
					case "":
						print_string = "約珥書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("珥","約珥書", chap_string, "1","darby")
							default:
								print_string = Bible_print_string("珥","約珥書", chap_string, sec_string,"darby")
						}
				}
			case "Amos","摩","阿摩司書","Am","am","Книга пророка Амоса","A-mốt","アモス書","アモス","아모스":
				bible_short_name = "摩"
				switch chap_string {
					case "":
						print_string = "阿摩司書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("摩","阿摩司書", chap_string, "1","darby")
							default:
								print_string = Bible_print_string("摩","阿摩司書", chap_string, sec_string,"darby")
						}
				}
			case "Obad","Obadiah","俄","俄巴底亞","俄巴底亞書","Ob","ob","오바댜","オバデヤ書","オバデヤ","Áp-đia","Книга пророка Авдия":
				bible_short_name = "俄"
				switch chap_string {
					case "":
						print_string = "俄巴底亞書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("俄","俄巴底亞書", chap_string, "1","darby")
							default:
								print_string = Bible_print_string("俄","俄巴底亞書", chap_string, sec_string,"darby")
						}
				}
			case "Jon","Jonah","拿","約拿","約拿書","jon","요나","ヨナ書","ヨナ","Giô-na","Книга пророка Ионы":
				bible_short_name = "拿"
				switch chap_string {
					case "":
						print_string = "約拿書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("拿","約拿書", chap_string, "1","darby")
							default:
								print_string = Bible_print_string("拿","約拿書", chap_string, sec_string,"darby")
						}
				}
			case "Micah","彌","彌迦","彌迦書","Mic","mic","Книга пророка Михея","Mi-chê","ミカ書","ミカ","미가":
				bible_short_name = "彌"
				switch chap_string {
					case "":
						print_string = "彌迦書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("彌","彌迦書", chap_string, "1","darby")
							default:
								print_string = Bible_print_string("彌","彌迦書", chap_string, sec_string,"darby")
						}
				}
			case "Nah","Nahum","鴻","那鴻","那鴻書","Na","na","Книга пророка Наума","Na-hum","ナホム書","ナホム","나훔":
				bible_short_name = "鴻"
				switch chap_string {
					case "":
						print_string = "那鴻書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("鴻","那鴻書", chap_string, "1","darby")
							default:
								print_string = Bible_print_string("鴻","那鴻書", chap_string, sec_string,"darby")
						}
				}
			case "Habakkuk","哈","哈巴","哈巴谷","哈巴谷書","Hab","hab","Книга пророка Аввакума","Ha-ba-cúc","ハバクク書","ハバクク","ハバ","クク","ハバ書","하박국":
				bible_short_name = "哈"
				switch chap_string {
					case "":
						print_string = "哈巴谷書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("哈","哈巴谷書", chap_string, "1","darby")
							default:
								print_string = Bible_print_string("哈","哈巴谷書", chap_string, sec_string,"darby")
						}
				}
			case "Zeph","Zephaniah","番","西番雅","西番雅書","Zep","zep","스바냐","ゼパニヤ書","ゼパニヤ","Sô-phô-ni","Книга пророка Софонии":
				bible_short_name = "番"
				switch chap_string {
					case "":
						print_string = "西番雅書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("番","西番雅書", chap_string, "1","darby")
							default:
								print_string = Bible_print_string("番","西番雅書", chap_string, sec_string,"darby")
						}
				}
			case "Haggai","該","哈該","哈該書","Hag","hag","학개","ハガイ書","ハガイ","A-ghê","Книга пророка Аггея":
				bible_short_name = "該"
				switch chap_string {
					case "":
						print_string = "哈該書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("該","哈該書", chap_string, "1","darby")
							default:
								print_string = Bible_print_string("該","哈該書", chap_string, sec_string,"darby")
						}
				}
			case "Zech","Zechariah","亞","撒迦利亞","撒迦利亞書","Zec","zec","Книга пророка Захарии","Xa-cha-ri","스가랴","ゼカリヤ書","ゼカリヤ":
				bible_short_name = "亞"
				switch chap_string {
					case "":
						print_string = "撒迦利亞書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("亞","撒迦利亞書", chap_string, "1","darby")
							default:
								print_string = Bible_print_string("亞","撒迦利亞書", chap_string, sec_string,"darby")
						}
				}
			case "Malachi","瑪","","瑪拉","瑪拉基","瑪拉基書","Mal","mal","말라기","マラキ書","マラキ","Ma-la-chi","Книга пророка Малахии":
				bible_short_name = "瑪"
				switch chap_string {
					case "":
						print_string = "瑪拉基書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("瑪","瑪拉基書", chap_string, "1","darby")
							default:
								print_string = Bible_print_string("瑪","瑪拉基書", chap_string, sec_string,"darby")
						}
				}
			case "Matt","Matthew","太","馬太","馬太福音","Mt","mt","마태복음","マタイによる福音書","マタイ","マタイによる","Ma-thi-ơ","От Матфея святое благовествование":
				bible_short_name = "太"
				switch chap_string {
					case "":
						print_string = "馬太福音"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("太","馬太福音", chap_string, "1","darby")
							default:
								print_string = Bible_print_string("太","馬太福音", chap_string, sec_string,"darby")
						}
				}
			case "Mark","可","馬可","馬可福音","Mr","mr","マルコによる福音書","マルコ","マルコによる","마가복음","Mác","От Марка святое благовествование":
				bible_short_name = "可"
				switch chap_string {
					case "":
						print_string = "馬可福音"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("可","馬可福音", chap_string, "1","darby")
							default:
								print_string = Bible_print_string("可","馬可福音", chap_string, sec_string,"darby")
						}
				}
			case "Luke","路","路加","路加福音","Lu","lu","От Луки святое благовествование","Lu-ca","ルカによる福音書","ルカ","ルカによる","누가복음":
				bible_short_name = "路"
				switch chap_string {
					case "":
						print_string = "路加福音"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("路","路加福音", chap_string, "1","darby")
							default:
								print_string = Bible_print_string("路","路加福音", chap_string, sec_string,"darby")
						}
				}
			case "John","約","約翰","約翰福音","Joh","joh","От Иоанна святое благовествование","Giăng","ヨハネによる福音書","ヨハネ","ヨハネによる","요한복음":
				bible_short_name = "約"
				switch chap_string {
					case "":
						print_string = "約翰福音"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("約","約翰福音", chap_string, "1","darby")
							default:
								print_string = Bible_print_string("約","約翰福音", chap_string, sec_string,"darby")
						}
				}
			case "Acts","徒","使徒","使徒行傳","Ac","ac","Деяния святых апостолов","Công-vụ Các Sứ-đồ","使徒行伝","사도행전":
				bible_short_name = "徒"
				switch chap_string {
					case "":
						print_string = "使徒行傳"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("徒","使徒行傳", chap_string, "1","darby")
							default:
								print_string = Bible_print_string("徒","使徒行傳", chap_string, sec_string,"darby")
						}
				}
			case "Rom","Romans","羅","羅馬","羅馬書","Ro","ro","Послание к Римлянам","Rô-ma","ローマ","ローマ人への手紙","로마서":
				bible_short_name = "羅"
				switch chap_string {
					case "":
						print_string = "羅馬書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("羅","羅馬書", chap_string, "1","darby")
							default:
								print_string = Bible_print_string("羅","羅馬書", chap_string, sec_string,"darby")
						}
				}
			case "1 Cor","First Corinthians","林前","哥林多前","哥林多前書","1Co","1co","Первое послание к Коринфянам","1 Cô-rinh-tô","コリント人への第一の手紙","コリント一","コリント人への第一","고린도전서":
				bible_short_name = "林前"
				switch chap_string {
					case "":
						print_string = "哥林多前書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("林前","哥林多前書", chap_string, "1","darby")
							default:
								print_string = Bible_print_string("林前","哥林多前書", chap_string, sec_string,"darby")
						}
				}
			case "2 Cor","Second Corinthians","林後","哥林多後","哥林多後書","2Co","2co","Второе послание к Коринфянам","2 Cô-rinh-tô","コリント人への第二の手紙","コリント二","コリント人への第二の","고린도후서":
				bible_short_name = "林後"
				switch chap_string {
					case "":
						print_string = "哥林多後書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("林後","哥林多後書", chap_string, "1","darby")
							default:
								print_string = Bible_print_string("林後","哥林多後書", chap_string, sec_string,"darby")
						}
				}
			case "Gal","Galatians","加","加拉太","加拉太書","Ga","ga","Послание к Галатам","Ga-la-ti","ガラテヤ","ガラテヤ人への手紙","갈라디아서":
				bible_short_name = "加"
				switch chap_string {
					case "":
						print_string = "加拉太書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("加","加拉太書", chap_string, "1","darby")
							default:
								print_string = Bible_print_string("加","加拉太書", chap_string, sec_string,"darby")
						}
				}
			case "Ephesians","弗","以弗所","以弗所書","Eph","eph","Послание к Ефесянам","Ê-phê-sô","エペソ人への手紙","エペソ","エペソ人","エペソ人の手紙","에베소서":
				bible_short_name = "弗"
				switch chap_string {
					case "":
						print_string = "以弗所書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("弗","以弗所書", chap_string, "1","darby")
							default:
								print_string = Bible_print_string("弗","以弗所書", chap_string, sec_string,"darby")
						}
				}
			case "Phil","Philippians","腓","腓立","腓立比","腓立比書","Php","php","빌립보서","ピリピ","ピリピ人.ピリピ人への手紙","Послание к Филиппийцам","Phi-líp":
				bible_short_name = "腓"
				switch chap_string {
					case "":
						print_string = "腓立比書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("腓","腓立比書", chap_string, "1","darby")
							default:
								print_string = Bible_print_string("腓","腓立比書", chap_string, sec_string,"darby")
						}
				}
			case "Col","col","Colossians","西","歌羅西","歌羅","歌羅西書","Послание к Колоссянам","Cô-lô-se","コロサイ人への手紙","コロサイ","コロ","골로새서":
				bible_short_name = "西"
				switch chap_string {
					case "":
						print_string = "歌羅西書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("西","歌羅西書", chap_string, "1","darby")
							default:
								print_string = Bible_print_string("西","歌羅西書", chap_string, sec_string,"darby")
						}
				}
			case "1 Thess","First Thessalonians","帖前","帖撒羅尼迦前","帖撒羅尼迦前書","1Th","1th","데살로니가전서","テサロニケ人への第一の手紙","テサ一","テサロニケ一","1 Tê-sa-lô-ni-ca","Первое послание к Фессалоникийцам (Солунянам)":
				bible_short_name = "帖前"
				switch chap_string {
					case "":
						print_string = "帖撒羅尼迦前書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("帖前","帖撒羅尼迦前書", chap_string, "1","darby")
							default:
								print_string = Bible_print_string("帖前","帖撒羅尼迦前書", chap_string, sec_string,"darby")
						}
				}
			case "2 Thess","Second Thessalonians","帖後","帖撒羅尼迦後","帖撒羅尼迦後書","2Th","2th","데살로니가후서","テサロニケ人への第二の手紙","テサ二","テサロニケ二","2 Tê-sa-lô-ni-ca","Второе послание к Фессалоникийцам (Солунянам)":
				bible_short_name = "帖後"
				switch chap_string {
					case "":
						print_string = "帖撒羅尼迦後書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("帖後","帖撒羅尼迦後書", chap_string, "1","darby")
							default:
								print_string = Bible_print_string("帖後","帖撒羅尼迦後書", chap_string, sec_string,"darby")
						}
				}
			case "1 Tim","First Timothy","提前","提摩太前","提摩太前書","1Ti","1ti","Первое послание к Тимофею","1 Ti-mô-thê","テモテヘの第一の手紙","テモテ一","디모데전서":
				bible_short_name = "提前"
				switch chap_string {
					case "":
						print_string = "提摩太前書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("提前","提摩太前書", chap_string, "1","darby")
							default:
								print_string = Bible_print_string("提前","提摩太前書", chap_string, sec_string,"darby")
						}
				}
			case "2 Tim","Second Timothy","提後","提摩太後","提摩太後書","2Ti","2ti","Второе послание к Тимофею","2 Ti-mô-thê","テモテヘの第二の手紙","テモテ二","디모데후서":
				bible_short_name = "提後"
				switch chap_string {
					case "":
						print_string = "提摩太後書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("提後","提摩太後書", chap_string, "1","darby")
							default:
								print_string = Bible_print_string("提後","提摩太後書", chap_string, sec_string,"darby")
						}
				}
			case "Titus","多","提多","提多書","Tit","tit","Послание к Титу","Tít","テトスヘの手紙","テトス","디도서":
				bible_short_name = "多"
				switch chap_string {
					case "":
						print_string = "提多書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("多","提多書", chap_string, "1","darby")
							default:
								print_string = Bible_print_string("多","提多書", chap_string, sec_string,"darby")
						}
				}
			case "Philem","Philemon","門","腓利","腓利門","腓利門書","Phm","phm","Послание к Филимону","Phi-lê-môn","ピレモンヘの手紙","ピレモン","빌레몬서":
				bible_short_name = "門"
				switch chap_string {
					case "":
						print_string = "腓利門書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("門","腓利門書", chap_string, "1","darby")
							default:
								print_string = Bible_print_string("門","腓利門書", chap_string, sec_string,"darby")
						}
				}
			case "Heb","Hebrews","來","希伯來","希伯來書","heb","Послание к Евреям","Hê-bơ-rơ","ヘブル人への手紙","ヘブル","히브리서":
				bible_short_name = "來"
				switch chap_string {
					case "":
						print_string = "希伯來書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("來","希伯來書", chap_string, "1","darby")
							default:
								print_string = Bible_print_string("來","希伯來書", chap_string, sec_string,"darby")
						}
				}
			case "James","雅","雅各","雅各書","Jas","jas","Послание Иакова","Gia-cơ","ヤコブの手紙","ヤコブ","야고보서":
				bible_short_name = "雅"
				switch chap_string {
					case "":
						print_string = "雅各書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("雅","雅各書", chap_string, "1","darby")
							default:
								print_string = Bible_print_string("雅","雅各書", chap_string, sec_string,"darby")
						}
				}
			case "1 Pet","First Peter","彼前","彼得前","彼得前書","1Pe","1pe","Первое послание Петра","1 Phi-e-rơ","ペテロの第一の手紙","ペテロ一","베드로전서":
				bible_short_name = "彼前"
				switch chap_string {
					case "":
						print_string = "彼得前書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("彼前","彼得前書", chap_string, "1","darby")
							default:
								print_string = Bible_print_string("彼前","彼得前書", chap_string, sec_string,"darby")
						}
				}
			case "2 Pet","Second Peter","彼後","彼得後","彼得後書","2Pe","2pe","Второе послание Петра","2 Phi-e-rơ","ペテロの第二の手紙","ペテロ","베드로후서":
				bible_short_name = "彼後"
				switch chap_string {
					case "":
						print_string = "彼得後書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("彼後","彼得後書", chap_string, "1","darby")
							default:
								print_string = Bible_print_string("彼後","彼得後書", chap_string, sec_string,"darby")
						}
				}
			case "1 John","First John","約一","約翰一書","約翰1","約翰1書","1Jo","1jo","Первое послание Иоанна","1 Giăng","ヨハネの第一の手紙","ヨハネ一","요한일서":
				bible_short_name = "約一"
				switch chap_string {
					case "":
						print_string = "約翰一書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("約一","約翰一書", chap_string, "1","darby")
							default:
								print_string = Bible_print_string("約一","約翰一書", chap_string, sec_string,"darby")
						}
				}
			case "2 John","second John","約二","約翰二書","約翰2","約翰2書","2Jo","Второе послание Иоанна","2 Giăng","ヨハネの第二の手紙","ヨハネ二","요한2서":
				bible_short_name = "約二"
				switch chap_string {
					case "":
						print_string = "約翰二書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("約二","約翰二書", chap_string, "1","darby")
							default:
								print_string = Bible_print_string("約二","約翰二書", chap_string, sec_string,"darby")
						}
				}
			case "3 John","Third John","約三","約翰三書","約翰3","約翰3書","3Jo","3jo","Третье послание Иоанна","3 Giăng","ヨハネの第三の手紙","ヨハネ三","요한3서":
				bible_short_name = "約三"
				switch chap_string {
					case "":
						print_string = "約翰三書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("約三","約翰三書", chap_string, "1","darby")
							default:
								print_string = Bible_print_string("約三","約翰三書", chap_string, sec_string,"darby")
						}
				}
			case "Jude","猶","猶大","猶大書","jude","Послание Иуды","Giu-đe","ユダの手紙","ユダ","유다서":
				bible_short_name = "猶"
				switch chap_string {
					case "":
						print_string = "猶大書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("猶","猶大書", chap_string, "1","darby")
							default:
								print_string = Bible_print_string("猶","猶大書", chap_string, sec_string,"darby")
						}
				}
			case "Rev","Revelation","啟","啟示","啟示錄","Re","re","ｒｅ","Ｒｅ","rev","Откровение ап. Иоанна Богослова (Апокалипсис)","Khải-huyền","ヨハネの黙示録","黙示録","요한계시록":
				bible_short_name = "啟"
				switch chap_string {
					case "":
						print_string = "啟示錄"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("啟","啟示錄", chap_string, "1","darby")
							default:
								print_string = Bible_print_string("啟","啟示錄", chap_string, sec_string,"darby")
						}
				}
			default:
				print_string = "聖經"
				//print_string = "你是要找 " +  reg.ReplaceAllString(text, "$3") + " 對嗎？\n對不起，我還沒學呢...\n"
		}
	case "英文聖經ASV","ASV","Asv","asv","ＡＳＶ","Ａｓｖ","ａｓｖ":
		log.Print(reg.ReplaceAllString(text, "$3"))
		switch reg.ReplaceAllString(text, "$3") {
			case "Gen","Genesis","創","創世","創世紀","創世記","Ge","ge","gen","창세기","Sáng-thế Ký","Бытие":
				bible_short_name = "創"
				switch chap_string {
					case "":
						print_string = "創世記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("創","創世記", chap_string, "1","asv")
							default:
								print_string = Bible_print_string("創","創世記", chap_string, sec_string,"asv")
						}
				}
			case "ex","Ex","Exodus","埃及","出","出埃及","出埃及記","출애굽기","エジプト","出エジプト","出エジプト記","Xuất Ê-díp-tô Ký","Исход":
				bible_short_name = "出"
				switch chap_string {
					case "":
						print_string = "出埃及記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("出","出埃及記", chap_string, "1","asv")
							default:
								print_string = Bible_print_string("出","出埃及記", chap_string, sec_string,"asv")
						}
				}
			case "Lev","Leviticus","利","利未","利未記","Le","le","Левит","Lê-vi Ký","レビ記","レビ","레위기":
				bible_short_name = "利"
				switch chap_string {
					case "":
						print_string = "利未記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("利","利未記", chap_string, "1","asv")
							default:
								print_string = Bible_print_string("利","利未記", chap_string, sec_string,"asv")
						}
				}
			case "Num","Numbers","民","民數","民數記","Nu","nu","민수기","民数","民数記","Dân-số Ký","Числа":
				bible_short_name = "民"
				switch chap_string {
					case "":
						print_string = "民數記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("民","民數記", chap_string, "1","asv")
							default:
								print_string = Bible_print_string("民","民數記", chap_string, sec_string,"asv")
						}
				}
			case "Deut","Deuteronomy","申","申命","申命記","De","de","신명기","Phục-truyền Luật-lệ Ký","Второзаконие":
				bible_short_name = "申"
				switch chap_string {
					case "":
						print_string = "申命記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("申","申命記", chap_string, "1","asv")
							default:
								print_string = Bible_print_string("申","申命記", chap_string, sec_string,"asv")
						}
				}
			case "Josh","Joshua","約書亞","約書亞記","Jos","jos","여호수아","ヨシュア記","ヨシュア","Giô-suê","Книга Иисуса Навина":
				bible_short_name = "書"
				switch chap_string {
					case "":
						print_string = "約書亞記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("書","約書亞記", chap_string, "1","asv")
							default:
								print_string = Bible_print_string("書","約書亞記", chap_string, sec_string,"asv")
						}
				}
			case "Judg","Judges","士","士師","士師記","Jud","jud","jdg","Jdg","Книга Судей израилевых","Các Quan Xét","사사기":
				bible_short_name = "士"
				switch chap_string {
					case "":
						print_string = "士師記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("士","士師記", chap_string, "1","asv")
							default:
								print_string = Bible_print_string("士","士師記", chap_string, sec_string,"asv")
						}
				}
			case "Ruth","路得","路得記","Ru","ru","Rut","rut","룻기","ルツ","ルツ記","Ru-tơ","Книга Руфи":
				bible_short_name = "得"
				switch chap_string {
					case "":
						print_string = "路得記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("得","路得記", chap_string, "1","asv")
							default:
								print_string = Bible_print_string("得","路得記", chap_string, sec_string,"asv")
						}
				}
			case "1 Sam","First Samuel","撒上","撒母耳記上","1Sa","1sa","サムエル記上","サムエル上","サム上","사무엘상","1 Sa-mu-ên","Первая книга Царств":
				bible_short_name = "撒上"
				switch chap_string {
					case "":
						print_string = "撒母耳記上"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("撒上","撒母耳記上", chap_string, "1","asv")
							default:
								print_string = Bible_print_string("撒上","撒母耳記上", chap_string, sec_string,"asv")
						}
				}
			case "2 Sam","Second Samuel","撒下","撒母耳記下","2Sa","2sa","사무엘하","サムエル記下","サムエル下","サム下","2 Sa-mu-ên","Вторая книга Царств":
				bible_short_name = "撒下"
				switch chap_string {
					case "":
						print_string = "撒母耳記下"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("撒下","撒母耳記下", chap_string, "1","asv")
							default:
								print_string = Bible_print_string("撒下","撒母耳記下", chap_string, sec_string,"asv")
						}
				}
			case "1 Kin","First Kings","王上","列王上","列王紀上","列王記上","1Ki","1ki","열왕기상","Третья книга Царств","1 Các Vua":
				bible_short_name = "王上"
				switch chap_string {
					case "":
						print_string = "列王紀上"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("王上","列王紀上", chap_string, "1","asv")
							default:
								print_string = Bible_print_string("王上","列王紀上", chap_string, sec_string,"asv")
						}
				}
			case "2 Kin","Second Kings","王下","列王下","列王記下","列王紀下","2Ki","2ki","열왕기하","2 Các Vua","Четвертая книга Царств":
				bible_short_name = "王下"
				switch chap_string {
					case "":
						print_string = "列王紀下"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("王下","列王紀下", chap_string, "1","asv")
							default:
								print_string = Bible_print_string("王下","列王紀下", chap_string, sec_string,"asv")
						}
				}
			case "1 Chr","First Chronicles","歷上","代上","歷代志上","歷代上","1Ch","1ch","Первая книга Паралипоменон","1 Sử-ký","歴上","歴代上","歴代志上","역대상":
				bible_short_name = "代上"
				switch chap_string {
					case "":
						print_string = "歷代志上"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("代上","歷代志上", chap_string, "1","asv")
							default:
								print_string = Bible_print_string("代上","歷代志上", chap_string, sec_string,"asv")
						}
				}
			case "2 Chr","Second Chronicles","代下","歷下","歷代下","歷代志下","2Ch","2ch","역대하","歴代志下","歴代下","歴下","2 Sử-ký","Вторая книга Паралипоменон":
				bible_short_name = "代下"
				switch chap_string {
					case "":
						print_string = "歷代志下"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("代下","歷代志下", chap_string, "1","asv")
							default:
								print_string = Bible_print_string("代下","歷代志下", chap_string, sec_string,"asv")
						}
				}
			case "Ezra","拉","以斯拉","以斯拉記","Ezr","ezr","Первая книга Ездры","Ê-xơ-ra","エズラ","エズラ記","에스라":
				bible_short_name = "拉"
				switch chap_string {
					case "":
						print_string = "以斯拉記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("拉","以斯拉記", chap_string, "1","asv")
							default:
								print_string = Bible_print_string("拉","以斯拉記", chap_string, sec_string,"asv")
						}
				}
			case "Neh","Nehemiah","尼","尼希米","尼希米記","Ne","ne","느헤미야","ネヘミヤ書","ネヘミヤ","Nê-hê-mi","Книга Неемии":
				bible_short_name = "尼"
				switch chap_string {
					case "":
						print_string = "尼希米記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("尼","尼希米記", chap_string, "1","asv")
							default:
								print_string = Bible_print_string("尼","尼希米記", chap_string, sec_string,"asv")
						}
				}
			case "Esth","Esther","斯","以斯帖","以斯帖記","Es","est","Есфирь","Ê-xơ-tê","エステル","エステル記","에스더":
				bible_short_name = "斯"
				switch chap_string {
					case "":
						print_string = "以斯帖記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("斯","以斯帖記", chap_string, "1","asv")
							default:
								print_string = Bible_print_string("斯","以斯帖記", chap_string, sec_string,"asv")
						}
				}
			case "Job","job","伯","約伯","約伯記","Книга Иова","Gióp","ヨブ","ヨブ記","욥기":
				bible_short_name = "伯"
				switch chap_string {
					case "":
						print_string = "約伯記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("伯","約伯記", chap_string, "1","asv")
							default:
								print_string = Bible_print_string("伯","約伯記", chap_string, sec_string,"asv")
						}
				}
			case "Ps","Psalms","詩","詩篇","ps","시편","Thi-thiên","Псалтирь":
				bible_short_name = "詩"
				switch chap_string {
					case "":
						print_string = "詩篇"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("詩","詩篇", chap_string, "1","asv")
							default:
								print_string = Bible_print_string("詩","詩篇", chap_string, sec_string,"asv")
						}
				}
			case "Prov","Proverbs","箴","箴言","Pr","pr","Притчи Соломона","Châm-ngôn","잠언":
				bible_short_name = "箴"
				switch chap_string {
					case "":
						print_string = "箴言"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("箴","箴言", chap_string, "1","asv")
							default:
								print_string = Bible_print_string("箴","箴言", chap_string, sec_string,"asv")
						}
				}
			case "Eccl","Ecclesiastes","傳","傳道","傳道書","Ec","ec","Книга Екклезиаста","или Проповедника","Truyền-đạo","伝道の書","伝道","伝","伝道書","전도서":
				bible_short_name = "傳"
				switch chap_string {
					case "":
						print_string = "傳道書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("傳","傳道書", chap_string, "1","asv")
							default:
								print_string = Bible_print_string("傳","傳道書", chap_string, sec_string,"asv")
						}
				}
			case "Song","Song of Solomon","歌","雅歌","So","so","sng","Sng","Песнь песней Соломона","Nhã-ca","아가":
				bible_short_name = "歌"
				switch chap_string {
					case "":
						print_string = "雅歌"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("歌","雅歌", chap_string, "1","asv")
							default:
								print_string = Bible_print_string("歌","雅歌", chap_string, sec_string,"asv")
						}
				}
			case "Is","Isaiah","賽","以賽","以賽亞","以賽亞書","Isa","isa","Книга пророка Исаии","Ê-sai","イザヤ書","イザヤ","이사야":
				bible_short_name = "賽"
				switch chap_string {
					case "":
						print_string = "以賽亞書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("賽","以賽亞書", chap_string, "1","asv")
							default:
								print_string = Bible_print_string("賽","以賽亞書", chap_string, sec_string,"asv")
						}
				}
			case "Jer","Jeremiah","耶","耶利米","耶利米書","jer","예레미야","エレミヤ","エレミヤ書","Giê-rê-mi","Книга пророка Иеремии":
				bible_short_name = "耶"
				switch chap_string {
					case "":
						print_string = "耶利米書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("耶","耶利米書", chap_string, "1","asv")
							default:
								print_string = Bible_print_string("耶","耶利米書", chap_string, sec_string,"asv")
						}
				}
			case "Lam","Lamentations","哀","哀歌","耶利米哀歌","La","lam","예레미야애가","Ca-thương","Плач Иеремии":
				bible_short_name = "哀"
				switch chap_string {
					case "":
						print_string = "耶利米哀歌"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("哀","耶利米哀歌", chap_string, "1","asv")
							default:
								print_string = Bible_print_string("哀","耶利米哀歌", chap_string, sec_string,"asv")
						}
				}
			case "Ezek","Ezekiel","結","以西結","以西結書","Eze","eze","에스겔","エゼキエル書","エゼキエル","Ê-xê-chi-ên","Книга пророка Иезекииля":
				bible_short_name = "結"
				switch chap_string {
					case "":
						print_string = "以西結書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("結","以西結書", chap_string, "1","asv")
							default:
								print_string = Bible_print_string("結","以西結書", chap_string, sec_string,"asv")
						}
				}
			case "Dan","Daniel","但","但以理","但以理書","Da","da","Книга пророка Даниила","Đa-ni-ên","ダニエル書","ダニエル","다니엘":
				bible_short_name = "但"
				switch chap_string {
					case "":
						print_string = "但以理書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("但","但以理書", chap_string, "1","asv")
							default:
								print_string = Bible_print_string("但","但以理書", chap_string, sec_string,"asv")
						}
				}
			case "Hos","Hosea","何","何西","何西阿","何西阿書","Ho","ho","Книга пророка Осии","Ô-sê","ホセア書","ホセア","호세아":
				bible_short_name = "何"
				switch chap_string {
					case "":
						print_string = "何西阿書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("何","何西阿書", chap_string, "1","asv")
							default:
								print_string = Bible_print_string("何","何西阿書", chap_string, sec_string,"asv")
						}
				}
			case "Joel","珥","約珥","約珥書","Joe","joe","Книга пророка Иоиля","Giô-ên","ヨエル書","ヨエル","요엘":
				bible_short_name = "珥"
				switch chap_string {
					case "":
						print_string = "約珥書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("珥","約珥書", chap_string, "1","asv")
							default:
								print_string = Bible_print_string("珥","約珥書", chap_string, sec_string,"asv")
						}
				}
			case "Amos","摩","阿摩司書","Am","am","Книга пророка Амоса","A-mốt","アモス書","アモス","아모스":
				bible_short_name = "摩"
				switch chap_string {
					case "":
						print_string = "阿摩司書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("摩","阿摩司書", chap_string, "1","asv")
							default:
								print_string = Bible_print_string("摩","阿摩司書", chap_string, sec_string,"asv")
						}
				}
			case "Obad","Obadiah","俄","俄巴底亞","俄巴底亞書","Ob","ob","오바댜","オバデヤ書","オバデヤ","Áp-đia","Книга пророка Авдия":
				bible_short_name = "俄"
				switch chap_string {
					case "":
						print_string = "俄巴底亞書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("俄","俄巴底亞書", chap_string, "1","asv")
							default:
								print_string = Bible_print_string("俄","俄巴底亞書", chap_string, sec_string,"asv")
						}
				}
			case "Jon","Jonah","拿","約拿","約拿書","jon","요나","ヨナ書","ヨナ","Giô-na","Книга пророка Ионы":
				bible_short_name = "拿"
				switch chap_string {
					case "":
						print_string = "約拿書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("拿","約拿書", chap_string, "1","asv")
							default:
								print_string = Bible_print_string("拿","約拿書", chap_string, sec_string,"asv")
						}
				}
			case "Micah","彌","彌迦","彌迦書","Mic","mic","Книга пророка Михея","Mi-chê","ミカ書","ミカ","미가":
				bible_short_name = "彌"
				switch chap_string {
					case "":
						print_string = "彌迦書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("彌","彌迦書", chap_string, "1","asv")
							default:
								print_string = Bible_print_string("彌","彌迦書", chap_string, sec_string,"asv")
						}
				}
			case "Nah","Nahum","鴻","那鴻","那鴻書","Na","na","Книга пророка Наума","Na-hum","ナホム書","ナホム","나훔":
				bible_short_name = "鴻"
				switch chap_string {
					case "":
						print_string = "那鴻書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("鴻","那鴻書", chap_string, "1","asv")
							default:
								print_string = Bible_print_string("鴻","那鴻書", chap_string, sec_string,"asv")
						}
				}
			case "Habakkuk","哈","哈巴","哈巴谷","哈巴谷書","Hab","hab","Книга пророка Аввакума","Ha-ba-cúc","ハバクク書","ハバクク","ハバ","クク","ハバ書","하박국":
				bible_short_name = "哈"
				switch chap_string {
					case "":
						print_string = "哈巴谷書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("哈","哈巴谷書", chap_string, "1","asv")
							default:
								print_string = Bible_print_string("哈","哈巴谷書", chap_string, sec_string,"asv")
						}
				}
			case "Zeph","Zephaniah","番","西番雅","西番雅書","Zep","zep","스바냐","ゼパニヤ書","ゼパニヤ","Sô-phô-ni","Книга пророка Софонии":
				bible_short_name = "番"
				switch chap_string {
					case "":
						print_string = "西番雅書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("番","西番雅書", chap_string, "1","asv")
							default:
								print_string = Bible_print_string("番","西番雅書", chap_string, sec_string,"asv")
						}
				}
			case "Haggai","該","哈該","哈該書","Hag","hag","학개","ハガイ書","ハガイ","A-ghê","Книга пророка Аггея":
				bible_short_name = "該"
				switch chap_string {
					case "":
						print_string = "哈該書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("該","哈該書", chap_string, "1","asv")
							default:
								print_string = Bible_print_string("該","哈該書", chap_string, sec_string,"asv")
						}
				}
			case "Zech","Zechariah","亞","撒迦利亞","撒迦利亞書","Zec","zec","Книга пророка Захарии","Xa-cha-ri","스가랴","ゼカリヤ書","ゼカリヤ":
				bible_short_name = "亞"
				switch chap_string {
					case "":
						print_string = "撒迦利亞書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("亞","撒迦利亞書", chap_string, "1","asv")
							default:
								print_string = Bible_print_string("亞","撒迦利亞書", chap_string, sec_string,"asv")
						}
				}
			case "Malachi","瑪","","瑪拉","瑪拉基","瑪拉基書","Mal","mal","말라기","マラキ書","マラキ","Ma-la-chi","Книга пророка Малахии":
				bible_short_name = "瑪"
				switch chap_string {
					case "":
						print_string = "瑪拉基書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("瑪","瑪拉基書", chap_string, "1","asv")
							default:
								print_string = Bible_print_string("瑪","瑪拉基書", chap_string, sec_string,"asv")
						}
				}
			case "Matt","Matthew","太","馬太","馬太福音","Mt","mt","마태복음","マタイによる福音書","マタイ","マタイによる","Ma-thi-ơ","От Матфея святое благовествование":
				bible_short_name = "太"
				switch chap_string {
					case "":
						print_string = "馬太福音"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("太","馬太福音", chap_string, "1","asv")
							default:
								print_string = Bible_print_string("太","馬太福音", chap_string, sec_string,"asv")
						}
				}
			case "Mark","可","馬可","馬可福音","Mr","mr","マルコによる福音書","マルコ","マルコによる","마가복음","Mác","От Марка святое благовествование":
				bible_short_name = "可"
				switch chap_string {
					case "":
						print_string = "馬可福音"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("可","馬可福音", chap_string, "1","asv")
							default:
								print_string = Bible_print_string("可","馬可福音", chap_string, sec_string,"asv")
						}
				}
			case "Luke","路","路加","路加福音","Lu","lu","От Луки святое благовествование","Lu-ca","ルカによる福音書","ルカ","ルカによる","누가복음":
				bible_short_name = "路"
				switch chap_string {
					case "":
						print_string = "路加福音"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("路","路加福音", chap_string, "1","asv")
							default:
								print_string = Bible_print_string("路","路加福音", chap_string, sec_string,"asv")
						}
				}
			case "John","約","約翰","約翰福音","Joh","joh","От Иоанна святое благовествование","Giăng","ヨハネによる福音書","ヨハネ","ヨハネによる","요한복음":
				bible_short_name = "約"
				switch chap_string {
					case "":
						print_string = "約翰福音"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("約","約翰福音", chap_string, "1","asv")
							default:
								print_string = Bible_print_string("約","約翰福音", chap_string, sec_string,"asv")
						}
				}
			case "Acts","徒","使徒","使徒行傳","Ac","ac","Деяния святых апостолов","Công-vụ Các Sứ-đồ","使徒行伝","사도행전":
				bible_short_name = "徒"
				switch chap_string {
					case "":
						print_string = "使徒行傳"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("徒","使徒行傳", chap_string, "1","asv")
							default:
								print_string = Bible_print_string("徒","使徒行傳", chap_string, sec_string,"asv")
						}
				}
			case "Rom","Romans","羅","羅馬","羅馬書","Ro","ro","Послание к Римлянам","Rô-ma","ローマ","ローマ人への手紙","로마서":
				bible_short_name = "羅"
				switch chap_string {
					case "":
						print_string = "羅馬書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("羅","羅馬書", chap_string, "1","asv")
							default:
								print_string = Bible_print_string("羅","羅馬書", chap_string, sec_string,"asv")
						}
				}
			case "1 Cor","First Corinthians","林前","哥林多前","哥林多前書","1Co","1co","Первое послание к Коринфянам","1 Cô-rinh-tô","コリント人への第一の手紙","コリント一","コリント人への第一","고린도전서":
				bible_short_name = "林前"
				switch chap_string {
					case "":
						print_string = "哥林多前書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("林前","哥林多前書", chap_string, "1","asv")
							default:
								print_string = Bible_print_string("林前","哥林多前書", chap_string, sec_string,"asv")
						}
				}
			case "2 Cor","Second Corinthians","林後","哥林多後","哥林多後書","2Co","2co","Второе послание к Коринфянам","2 Cô-rinh-tô","コリント人への第二の手紙","コリント二","コリント人への第二の","고린도후서":
				bible_short_name = "林後"
				switch chap_string {
					case "":
						print_string = "哥林多後書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("林後","哥林多後書", chap_string, "1","asv")
							default:
								print_string = Bible_print_string("林後","哥林多後書", chap_string, sec_string,"asv")
						}
				}
			case "Gal","Galatians","加","加拉太","加拉太書","Ga","ga","Послание к Галатам","Ga-la-ti","ガラテヤ","ガラテヤ人への手紙","갈라디아서":
				bible_short_name = "加"
				switch chap_string {
					case "":
						print_string = "加拉太書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("加","加拉太書", chap_string, "1","asv")
							default:
								print_string = Bible_print_string("加","加拉太書", chap_string, sec_string,"asv")
						}
				}
			case "Ephesians","弗","以弗所","以弗所書","Eph","eph","Послание к Ефесянам","Ê-phê-sô","エペソ人への手紙","エペソ","エペソ人","エペソ人の手紙","에베소서":
				bible_short_name = "弗"
				switch chap_string {
					case "":
						print_string = "以弗所書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("弗","以弗所書", chap_string, "1","asv")
							default:
								print_string = Bible_print_string("弗","以弗所書", chap_string, sec_string,"asv")
						}
				}
			case "Phil","Philippians","腓","腓立","腓立比","腓立比書","Php","php","빌립보서","ピリピ","ピリピ人.ピリピ人への手紙","Послание к Филиппийцам","Phi-líp":
				bible_short_name = "腓"
				switch chap_string {
					case "":
						print_string = "腓立比書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("腓","腓立比書", chap_string, "1","asv")
							default:
								print_string = Bible_print_string("腓","腓立比書", chap_string, sec_string,"asv")
						}
				}
			case "Col","col","Colossians","西","歌羅西","歌羅","歌羅西書","Послание к Колоссянам","Cô-lô-se","コロサイ人への手紙","コロサイ","コロ","골로새서":
				bible_short_name = "西"
				switch chap_string {
					case "":
						print_string = "歌羅西書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("西","歌羅西書", chap_string, "1","asv")
							default:
								print_string = Bible_print_string("西","歌羅西書", chap_string, sec_string,"asv")
						}
				}
			case "1 Thess","First Thessalonians","帖前","帖撒羅尼迦前","帖撒羅尼迦前書","1Th","1th","데살로니가전서","テサロニケ人への第一の手紙","テサ一","テサロニケ一","1 Tê-sa-lô-ni-ca","Первое послание к Фессалоникийцам (Солунянам)":
				bible_short_name = "帖前"
				switch chap_string {
					case "":
						print_string = "帖撒羅尼迦前書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("帖前","帖撒羅尼迦前書", chap_string, "1","asv")
							default:
								print_string = Bible_print_string("帖前","帖撒羅尼迦前書", chap_string, sec_string,"asv")
						}
				}
			case "2 Thess","Second Thessalonians","帖後","帖撒羅尼迦後","帖撒羅尼迦後書","2Th","2th","데살로니가후서","テサロニケ人への第二の手紙","テサ二","テサロニケ二","2 Tê-sa-lô-ni-ca","Второе послание к Фессалоникийцам (Солунянам)":
				bible_short_name = "帖後"
				switch chap_string {
					case "":
						print_string = "帖撒羅尼迦後書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("帖後","帖撒羅尼迦後書", chap_string, "1","asv")
							default:
								print_string = Bible_print_string("帖後","帖撒羅尼迦後書", chap_string, sec_string,"asv")
						}
				}
			case "1 Tim","First Timothy","提前","提摩太前","提摩太前書","1Ti","1ti","Первое послание к Тимофею","1 Ti-mô-thê","テモテヘの第一の手紙","テモテ一","디모데전서":
				bible_short_name = "提前"
				switch chap_string {
					case "":
						print_string = "提摩太前書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("提前","提摩太前書", chap_string, "1","asv")
							default:
								print_string = Bible_print_string("提前","提摩太前書", chap_string, sec_string,"asv")
						}
				}
			case "2 Tim","Second Timothy","提後","提摩太後","提摩太後書","2Ti","2ti","Второе послание к Тимофею","2 Ti-mô-thê","テモテヘの第二の手紙","テモテ二","디모데후서":
				bible_short_name = "提後"
				switch chap_string {
					case "":
						print_string = "提摩太後書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("提後","提摩太後書", chap_string, "1","asv")
							default:
								print_string = Bible_print_string("提後","提摩太後書", chap_string, sec_string,"asv")
						}
				}
			case "Titus","多","提多","提多書","Tit","tit","Послание к Титу","Tít","テトスヘの手紙","テトス","디도서":
				bible_short_name = "多"
				switch chap_string {
					case "":
						print_string = "提多書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("多","提多書", chap_string, "1","asv")
							default:
								print_string = Bible_print_string("多","提多書", chap_string, sec_string,"asv")
						}
				}
			case "Philem","Philemon","門","腓利","腓利門","腓利門書","Phm","phm","Послание к Филимону","Phi-lê-môn","ピレモンヘの手紙","ピレモン","빌레몬서":
				bible_short_name = "門"
				switch chap_string {
					case "":
						print_string = "腓利門書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("門","腓利門書", chap_string, "1","asv")
							default:
								print_string = Bible_print_string("門","腓利門書", chap_string, sec_string,"asv")
						}
				}
			case "Heb","Hebrews","來","希伯來","希伯來書","heb","Послание к Евреям","Hê-bơ-rơ","ヘブル人への手紙","ヘブル","히브리서":
				bible_short_name = "來"
				switch chap_string {
					case "":
						print_string = "希伯來書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("來","希伯來書", chap_string, "1","asv")
							default:
								print_string = Bible_print_string("來","希伯來書", chap_string, sec_string,"asv")
						}
				}
			case "James","雅","雅各","雅各書","Jas","jas","Послание Иакова","Gia-cơ","ヤコブの手紙","ヤコブ","야고보서":
				bible_short_name = "雅"
				switch chap_string {
					case "":
						print_string = "雅各書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("雅","雅各書", chap_string, "1","asv")
							default:
								print_string = Bible_print_string("雅","雅各書", chap_string, sec_string,"asv")
						}
				}
			case "1 Pet","First Peter","彼前","彼得前","彼得前書","1Pe","1pe","Первое послание Петра","1 Phi-e-rơ","ペテロの第一の手紙","ペテロ一","베드로전서":
				bible_short_name = "彼前"
				switch chap_string {
					case "":
						print_string = "彼得前書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("彼前","彼得前書", chap_string, "1","asv")
							default:
								print_string = Bible_print_string("彼前","彼得前書", chap_string, sec_string,"asv")
						}
				}
			case "2 Pet","Second Peter","彼後","彼得後","彼得後書","2Pe","2pe","Второе послание Петра","2 Phi-e-rơ","ペテロの第二の手紙","ペテロ","베드로후서":
				bible_short_name = "彼後"
				switch chap_string {
					case "":
						print_string = "彼得後書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("彼後","彼得後書", chap_string, "1","asv")
							default:
								print_string = Bible_print_string("彼後","彼得後書", chap_string, sec_string,"asv")
						}
				}
			case "1 John","First John","約一","約翰一書","約翰1","約翰1書","1Jo","1jo","Первое послание Иоанна","1 Giăng","ヨハネの第一の手紙","ヨハネ一","요한일서":
				bible_short_name = "約一"
				switch chap_string {
					case "":
						print_string = "約翰一書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("約一","約翰一書", chap_string, "1","asv")
							default:
								print_string = Bible_print_string("約一","約翰一書", chap_string, sec_string,"asv")
						}
				}
			case "2 John","second John","約二","約翰二書","約翰2","約翰2書","2Jo","Второе послание Иоанна","2 Giăng","ヨハネの第二の手紙","ヨハネ二","요한2서":
				bible_short_name = "約二"
				switch chap_string {
					case "":
						print_string = "約翰二書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("約二","約翰二書", chap_string, "1","asv")
							default:
								print_string = Bible_print_string("約二","約翰二書", chap_string, sec_string,"asv")
						}
				}
			case "3 John","Third John","約三","約翰三書","約翰3","約翰3書","3Jo","3jo","Третье послание Иоанна","3 Giăng","ヨハネの第三の手紙","ヨハネ三","요한3서":
				bible_short_name = "約三"
				switch chap_string {
					case "":
						print_string = "約翰三書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("約三","約翰三書", chap_string, "1","asv")
							default:
								print_string = Bible_print_string("約三","約翰三書", chap_string, sec_string,"asv")
						}
				}
			case "Jude","猶","猶大","猶大書","jude","Послание Иуды","Giu-đe","ユダの手紙","ユダ","유다서":
				bible_short_name = "猶"
				switch chap_string {
					case "":
						print_string = "猶大書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("猶","猶大書", chap_string, "1","asv")
							default:
								print_string = Bible_print_string("猶","猶大書", chap_string, sec_string,"asv")
						}
				}
			case "Rev","Revelation","啟","啟示","啟示錄","Re","re","ｒｅ","Ｒｅ","rev","Откровение ап. Иоанна Богослова (Апокалипсис)","Khải-huyền","ヨハネの黙示録","黙示録","요한계시록":
				bible_short_name = "啟"
				switch chap_string {
					case "":
						print_string = "啟示錄"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("啟","啟示錄", chap_string, "1","asv")
							default:
								print_string = Bible_print_string("啟","啟示錄", chap_string, sec_string,"asv")
						}
				}
			default:
				print_string = "聖經"
				//print_string = "你是要找 " +  reg.ReplaceAllString(text, "$3") + " 對嗎？\n對不起，我還沒學呢...\n"
		}
	case "英文聖經WEB","WEB","Web","web","ＷＥＢ","Ｗｅｂ","ｗｅｂ":
		log.Print(reg.ReplaceAllString(text, "$3"))
		switch reg.ReplaceAllString(text, "$3") {
			case "Gen","Genesis","創","創世","創世紀","創世記","Ge","ge","gen","창세기","Sáng-thế Ký","Бытие":
				bible_short_name = "創"
				switch chap_string {
					case "":
						print_string = "創世記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("創","創世記", chap_string, "1","web")
							default:
								print_string = Bible_print_string("創","創世記", chap_string, sec_string,"web")
						}
				}
			case "ex","Ex","Exodus","埃及","出","出埃及","出埃及記","출애굽기","エジプト","出エジプト","出エジプト記","Xuất Ê-díp-tô Ký","Исход":
				bible_short_name = "出"
				switch chap_string {
					case "":
						print_string = "出埃及記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("出","出埃及記", chap_string, "1","web")
							default:
								print_string = Bible_print_string("出","出埃及記", chap_string, sec_string,"web")
						}
				}
			case "Lev","Leviticus","利","利未","利未記","Le","le","Левит","Lê-vi Ký","レビ記","レビ","레위기":
				bible_short_name = "利"
				switch chap_string {
					case "":
						print_string = "利未記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("利","利未記", chap_string, "1","web")
							default:
								print_string = Bible_print_string("利","利未記", chap_string, sec_string,"web")
						}
				}
			case "Num","Numbers","民","民數","民數記","Nu","nu","민수기","民数","民数記","Dân-số Ký","Числа":
				bible_short_name = "民"
				switch chap_string {
					case "":
						print_string = "民數記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("民","民數記", chap_string, "1","web")
							default:
								print_string = Bible_print_string("民","民數記", chap_string, sec_string,"web")
						}
				}
			case "Deut","Deuteronomy","申","申命","申命記","De","de","신명기","Phục-truyền Luật-lệ Ký","Второзаконие":
				bible_short_name = "申"
				switch chap_string {
					case "":
						print_string = "申命記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("申","申命記", chap_string, "1","web")
							default:
								print_string = Bible_print_string("申","申命記", chap_string, sec_string,"web")
						}
				}
			case "Josh","Joshua","約書亞","約書亞記","Jos","jos","여호수아","ヨシュア記","ヨシュア","Giô-suê","Книга Иисуса Навина":
				bible_short_name = "書"
				switch chap_string {
					case "":
						print_string = "約書亞記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("書","約書亞記", chap_string, "1","web")
							default:
								print_string = Bible_print_string("書","約書亞記", chap_string, sec_string,"web")
						}
				}
			case "Judg","Judges","士","士師","士師記","Jud","jud","jdg","Jdg","Книга Судей израилевых","Các Quan Xét","사사기":
				bible_short_name = "士"
				switch chap_string {
					case "":
						print_string = "士師記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("士","士師記", chap_string, "1","web")
							default:
								print_string = Bible_print_string("士","士師記", chap_string, sec_string,"web")
						}
				}
			case "Ruth","路得","路得記","Ru","ru","Rut","rut","룻기","ルツ","ルツ記","Ru-tơ","Книга Руфи":
				bible_short_name = "得"
				switch chap_string {
					case "":
						print_string = "路得記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("得","路得記", chap_string, "1","web")
							default:
								print_string = Bible_print_string("得","路得記", chap_string, sec_string,"web")
						}
				}
			case "1 Sam","First Samuel","撒上","撒母耳記上","1Sa","1sa","サムエル記上","サムエル上","サム上","사무엘상","1 Sa-mu-ên","Первая книга Царств":
				bible_short_name = "撒上"
				switch chap_string {
					case "":
						print_string = "撒母耳記上"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("撒上","撒母耳記上", chap_string, "1","web")
							default:
								print_string = Bible_print_string("撒上","撒母耳記上", chap_string, sec_string,"web")
						}
				}
			case "2 Sam","Second Samuel","撒下","撒母耳記下","2Sa","2sa","사무엘하","サムエル記下","サムエル下","サム下","2 Sa-mu-ên","Вторая книга Царств":
				bible_short_name = "撒下"
				switch chap_string {
					case "":
						print_string = "撒母耳記下"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("撒下","撒母耳記下", chap_string, "1","web")
							default:
								print_string = Bible_print_string("撒下","撒母耳記下", chap_string, sec_string,"web")
						}
				}
			case "1 Kin","First Kings","王上","列王上","列王紀上","列王記上","1Ki","1ki","열왕기상","Третья книга Царств","1 Các Vua":
				bible_short_name = "王上"
				switch chap_string {
					case "":
						print_string = "列王紀上"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("王上","列王紀上", chap_string, "1","web")
							default:
								print_string = Bible_print_string("王上","列王紀上", chap_string, sec_string,"web")
						}
				}
			case "2 Kin","Second Kings","王下","列王下","列王記下","列王紀下","2Ki","2ki","열왕기하","2 Các Vua","Четвертая книга Царств":
				bible_short_name = "王下"
				switch chap_string {
					case "":
						print_string = "列王紀下"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("王下","列王紀下", chap_string, "1","web")
							default:
								print_string = Bible_print_string("王下","列王紀下", chap_string, sec_string,"web")
						}
				}
			case "1 Chr","First Chronicles","歷上","代上","歷代志上","歷代上","1Ch","1ch","Первая книга Паралипоменон","1 Sử-ký","歴上","歴代上","歴代志上","역대상":
				bible_short_name = "代上"
				switch chap_string {
					case "":
						print_string = "歷代志上"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("代上","歷代志上", chap_string, "1","web")
							default:
								print_string = Bible_print_string("代上","歷代志上", chap_string, sec_string,"web")
						}
				}
			case "2 Chr","Second Chronicles","代下","歷下","歷代下","歷代志下","2Ch","2ch","역대하","歴代志下","歴代下","歴下","2 Sử-ký","Вторая книга Паралипоменон":
				bible_short_name = "代下"
				switch chap_string {
					case "":
						print_string = "歷代志下"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("代下","歷代志下", chap_string, "1","web")
							default:
								print_string = Bible_print_string("代下","歷代志下", chap_string, sec_string,"web")
						}
				}
			case "Ezra","拉","以斯拉","以斯拉記","Ezr","ezr","Первая книга Ездры","Ê-xơ-ra","エズラ","エズラ記","에스라":
				bible_short_name = "拉"
				switch chap_string {
					case "":
						print_string = "以斯拉記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("拉","以斯拉記", chap_string, "1","web")
							default:
								print_string = Bible_print_string("拉","以斯拉記", chap_string, sec_string,"web")
						}
				}
			case "Neh","Nehemiah","尼","尼希米","尼希米記","Ne","ne","느헤미야","ネヘミヤ書","ネヘミヤ","Nê-hê-mi","Книга Неемии":
				bible_short_name = "尼"
				switch chap_string {
					case "":
						print_string = "尼希米記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("尼","尼希米記", chap_string, "1","web")
							default:
								print_string = Bible_print_string("尼","尼希米記", chap_string, sec_string,"web")
						}
				}
			case "Esth","Esther","斯","以斯帖","以斯帖記","Es","est","Есфирь","Ê-xơ-tê","エステル","エステル記","에스더":
				bible_short_name = "斯"
				switch chap_string {
					case "":
						print_string = "以斯帖記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("斯","以斯帖記", chap_string, "1","web")
							default:
								print_string = Bible_print_string("斯","以斯帖記", chap_string, sec_string,"web")
						}
				}
			case "Job","job","伯","約伯","約伯記","Книга Иова","Gióp","ヨブ","ヨブ記","욥기":
				bible_short_name = "伯"
				switch chap_string {
					case "":
						print_string = "約伯記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("伯","約伯記", chap_string, "1","web")
							default:
								print_string = Bible_print_string("伯","約伯記", chap_string, sec_string,"web")
						}
				}
			case "Ps","Psalms","詩","詩篇","ps","시편","Thi-thiên","Псалтирь":
				bible_short_name = "詩"
				switch chap_string {
					case "":
						print_string = "詩篇"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("詩","詩篇", chap_string, "1","web")
							default:
								print_string = Bible_print_string("詩","詩篇", chap_string, sec_string,"web")
						}
				}
			case "Prov","Proverbs","箴","箴言","Pr","pr","Притчи Соломона","Châm-ngôn","잠언":
				bible_short_name = "箴"
				switch chap_string {
					case "":
						print_string = "箴言"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("箴","箴言", chap_string, "1","web")
							default:
								print_string = Bible_print_string("箴","箴言", chap_string, sec_string,"web")
						}
				}
			case "Eccl","Ecclesiastes","傳","傳道","傳道書","Ec","ec","Книга Екклезиаста","или Проповедника","Truyền-đạo","伝道の書","伝道","伝","伝道書","전도서":
				bible_short_name = "傳"
				switch chap_string {
					case "":
						print_string = "傳道書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("傳","傳道書", chap_string, "1","web")
							default:
								print_string = Bible_print_string("傳","傳道書", chap_string, sec_string,"web")
						}
				}
			case "Song","Song of Solomon","歌","雅歌","So","so","sng","Sng","Песнь песней Соломона","Nhã-ca","아가":
				bible_short_name = "歌"
				switch chap_string {
					case "":
						print_string = "雅歌"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("歌","雅歌", chap_string, "1","web")
							default:
								print_string = Bible_print_string("歌","雅歌", chap_string, sec_string,"web")
						}
				}
			case "Is","Isaiah","賽","以賽","以賽亞","以賽亞書","Isa","isa","Книга пророка Исаии","Ê-sai","イザヤ書","イザヤ","이사야":
				bible_short_name = "賽"
				switch chap_string {
					case "":
						print_string = "以賽亞書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("賽","以賽亞書", chap_string, "1","web")
							default:
								print_string = Bible_print_string("賽","以賽亞書", chap_string, sec_string,"web")
						}
				}
			case "Jer","Jeremiah","耶","耶利米","耶利米書","jer","예레미야","エレミヤ","エレミヤ書","Giê-rê-mi","Книга пророка Иеремии":
				bible_short_name = "耶"
				switch chap_string {
					case "":
						print_string = "耶利米書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("耶","耶利米書", chap_string, "1","web")
							default:
								print_string = Bible_print_string("耶","耶利米書", chap_string, sec_string,"web")
						}
				}
			case "Lam","Lamentations","哀","哀歌","耶利米哀歌","La","lam","예레미야애가","Ca-thương","Плач Иеремии":
				bible_short_name = "哀"
				switch chap_string {
					case "":
						print_string = "耶利米哀歌"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("哀","耶利米哀歌", chap_string, "1","web")
							default:
								print_string = Bible_print_string("哀","耶利米哀歌", chap_string, sec_string,"web")
						}
				}
			case "Ezek","Ezekiel","結","以西結","以西結書","Eze","eze","에스겔","エゼキエル書","エゼキエル","Ê-xê-chi-ên","Книга пророка Иезекииля":
				bible_short_name = "結"
				switch chap_string {
					case "":
						print_string = "以西結書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("結","以西結書", chap_string, "1","web")
							default:
								print_string = Bible_print_string("結","以西結書", chap_string, sec_string,"web")
						}
				}
			case "Dan","Daniel","但","但以理","但以理書","Da","da","Книга пророка Даниила","Đa-ni-ên","ダニエル書","ダニエル","다니엘":
				bible_short_name = "但"
				switch chap_string {
					case "":
						print_string = "但以理書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("但","但以理書", chap_string, "1","web")
							default:
								print_string = Bible_print_string("但","但以理書", chap_string, sec_string,"web")
						}
				}
			case "Hos","Hosea","何","何西","何西阿","何西阿書","Ho","ho","Книга пророка Осии","Ô-sê","ホセア書","ホセア","호세아":
				bible_short_name = "何"
				switch chap_string {
					case "":
						print_string = "何西阿書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("何","何西阿書", chap_string, "1","web")
							default:
								print_string = Bible_print_string("何","何西阿書", chap_string, sec_string,"web")
						}
				}
			case "Joel","珥","約珥","約珥書","Joe","joe","Книга пророка Иоиля","Giô-ên","ヨエル書","ヨエル","요엘":
				bible_short_name = "珥"
				switch chap_string {
					case "":
						print_string = "約珥書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("珥","約珥書", chap_string, "1","web")
							default:
								print_string = Bible_print_string("珥","約珥書", chap_string, sec_string,"web")
						}
				}
			case "Amos","摩","阿摩司書","Am","am","Книга пророка Амоса","A-mốt","アモス書","アモス","아모스":
				bible_short_name = "摩"
				switch chap_string {
					case "":
						print_string = "阿摩司書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("摩","阿摩司書", chap_string, "1","web")
							default:
								print_string = Bible_print_string("摩","阿摩司書", chap_string, sec_string,"web")
						}
				}
			case "Obad","Obadiah","俄","俄巴底亞","俄巴底亞書","Ob","ob","오바댜","オバデヤ書","オバデヤ","Áp-đia","Книга пророка Авдия":
				bible_short_name = "俄"
				switch chap_string {
					case "":
						print_string = "俄巴底亞書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("俄","俄巴底亞書", chap_string, "1","web")
							default:
								print_string = Bible_print_string("俄","俄巴底亞書", chap_string, sec_string,"web")
						}
				}
			case "Jon","Jonah","拿","約拿","約拿書","jon","요나","ヨナ書","ヨナ","Giô-na","Книга пророка Ионы":
				bible_short_name = "拿"
				switch chap_string {
					case "":
						print_string = "約拿書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("拿","約拿書", chap_string, "1","web")
							default:
								print_string = Bible_print_string("拿","約拿書", chap_string, sec_string,"web")
						}
				}
			case "Micah","彌","彌迦","彌迦書","Mic","mic","Книга пророка Михея","Mi-chê","ミカ書","ミカ","미가":
				bible_short_name = "彌"
				switch chap_string {
					case "":
						print_string = "彌迦書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("彌","彌迦書", chap_string, "1","web")
							default:
								print_string = Bible_print_string("彌","彌迦書", chap_string, sec_string,"web")
						}
				}
			case "Nah","Nahum","鴻","那鴻","那鴻書","Na","na","Книга пророка Наума","Na-hum","ナホム書","ナホム","나훔":
				bible_short_name = "鴻"
				switch chap_string {
					case "":
						print_string = "那鴻書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("鴻","那鴻書", chap_string, "1","web")
							default:
								print_string = Bible_print_string("鴻","那鴻書", chap_string, sec_string,"web")
						}
				}
			case "Habakkuk","哈","哈巴","哈巴谷","哈巴谷書","Hab","hab","Книга пророка Аввакума","Ha-ba-cúc","ハバクク書","ハバクク","ハバ","クク","ハバ書","하박국":
				bible_short_name = "哈"
				switch chap_string {
					case "":
						print_string = "哈巴谷書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("哈","哈巴谷書", chap_string, "1","web")
							default:
								print_string = Bible_print_string("哈","哈巴谷書", chap_string, sec_string,"web")
						}
				}
			case "Zeph","Zephaniah","番","西番雅","西番雅書","Zep","zep","스바냐","ゼパニヤ書","ゼパニヤ","Sô-phô-ni","Книга пророка Софонии":
				bible_short_name = "番"
				switch chap_string {
					case "":
						print_string = "西番雅書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("番","西番雅書", chap_string, "1","web")
							default:
								print_string = Bible_print_string("番","西番雅書", chap_string, sec_string,"web")
						}
				}
			case "Haggai","該","哈該","哈該書","Hag","hag","학개","ハガイ書","ハガイ","A-ghê","Книга пророка Аггея":
				bible_short_name = "該"
				switch chap_string {
					case "":
						print_string = "哈該書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("該","哈該書", chap_string, "1","web")
							default:
								print_string = Bible_print_string("該","哈該書", chap_string, sec_string,"web")
						}
				}
			case "Zech","Zechariah","亞","撒迦利亞","撒迦利亞書","Zec","zec","Книга пророка Захарии","Xa-cha-ri","스가랴","ゼカリヤ書","ゼカリヤ":
				bible_short_name = "亞"
				switch chap_string {
					case "":
						print_string = "撒迦利亞書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("亞","撒迦利亞書", chap_string, "1","web")
							default:
								print_string = Bible_print_string("亞","撒迦利亞書", chap_string, sec_string,"web")
						}
				}
			case "Malachi","瑪","","瑪拉","瑪拉基","瑪拉基書","Mal","mal","말라기","マラキ書","マラキ","Ma-la-chi","Книга пророка Малахии":
				bible_short_name = "瑪"
				switch chap_string {
					case "":
						print_string = "瑪拉基書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("瑪","瑪拉基書", chap_string, "1","web")
							default:
								print_string = Bible_print_string("瑪","瑪拉基書", chap_string, sec_string,"web")
						}
				}
			case "Matt","Matthew","太","馬太","馬太福音","Mt","mt","마태복음","マタイによる福音書","マタイ","マタイによる","Ma-thi-ơ","От Матфея святое благовествование":
				bible_short_name = "太"
				switch chap_string {
					case "":
						print_string = "馬太福音"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("太","馬太福音", chap_string, "1","web")
							default:
								print_string = Bible_print_string("太","馬太福音", chap_string, sec_string,"web")
						}
				}
			case "Mark","可","馬可","馬可福音","Mr","mr","マルコによる福音書","マルコ","マルコによる","마가복음","Mác","От Марка святое благовествование":
				bible_short_name = "可"
				switch chap_string {
					case "":
						print_string = "馬可福音"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("可","馬可福音", chap_string, "1","web")
							default:
								print_string = Bible_print_string("可","馬可福音", chap_string, sec_string,"web")
						}
				}
			case "Luke","路","路加","路加福音","Lu","lu","От Луки святое благовествование","Lu-ca","ルカによる福音書","ルカ","ルカによる","누가복음":
				bible_short_name = "路"
				switch chap_string {
					case "":
						print_string = "路加福音"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("路","路加福音", chap_string, "1","web")
							default:
								print_string = Bible_print_string("路","路加福音", chap_string, sec_string,"web")
						}
				}
			case "John","約","約翰","約翰福音","Joh","joh","От Иоанна святое благовествование","Giăng","ヨハネによる福音書","ヨハネ","ヨハネによる","요한복음":
				bible_short_name = "約"
				switch chap_string {
					case "":
						print_string = "約翰福音"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("約","約翰福音", chap_string, "1","web")
							default:
								print_string = Bible_print_string("約","約翰福音", chap_string, sec_string,"web")
						}
				}
			case "Acts","徒","使徒","使徒行傳","Ac","ac","Деяния святых апостолов","Công-vụ Các Sứ-đồ","使徒行伝","사도행전":
				bible_short_name = "徒"
				switch chap_string {
					case "":
						print_string = "使徒行傳"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("徒","使徒行傳", chap_string, "1","web")
							default:
								print_string = Bible_print_string("徒","使徒行傳", chap_string, sec_string,"web")
						}
				}
			case "Rom","Romans","羅","羅馬","羅馬書","Ro","ro","Послание к Римлянам","Rô-ma","ローマ","ローマ人への手紙","로마서":
				bible_short_name = "羅"
				switch chap_string {
					case "":
						print_string = "羅馬書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("羅","羅馬書", chap_string, "1","web")
							default:
								print_string = Bible_print_string("羅","羅馬書", chap_string, sec_string,"web")
						}
				}
			case "1 Cor","First Corinthians","林前","哥林多前","哥林多前書","1Co","1co","Первое послание к Коринфянам","1 Cô-rinh-tô","コリント人への第一の手紙","コリント一","コリント人への第一","고린도전서":
				bible_short_name = "林前"
				switch chap_string {
					case "":
						print_string = "哥林多前書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("林前","哥林多前書", chap_string, "1","web")
							default:
								print_string = Bible_print_string("林前","哥林多前書", chap_string, sec_string,"web")
						}
				}
			case "2 Cor","Second Corinthians","林後","哥林多後","哥林多後書","2Co","2co","Второе послание к Коринфянам","2 Cô-rinh-tô","コリント人への第二の手紙","コリント二","コリント人への第二の","고린도후서":
				bible_short_name = "林後"
				switch chap_string {
					case "":
						print_string = "哥林多後書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("林後","哥林多後書", chap_string, "1","web")
							default:
								print_string = Bible_print_string("林後","哥林多後書", chap_string, sec_string,"web")
						}
				}
			case "Gal","Galatians","加","加拉太","加拉太書","Ga","ga","Послание к Галатам","Ga-la-ti","ガラテヤ","ガラテヤ人への手紙","갈라디아서":
				bible_short_name = "加"
				switch chap_string {
					case "":
						print_string = "加拉太書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("加","加拉太書", chap_string, "1","web")
							default:
								print_string = Bible_print_string("加","加拉太書", chap_string, sec_string,"web")
						}
				}
			case "Ephesians","弗","以弗所","以弗所書","Eph","eph","Послание к Ефесянам","Ê-phê-sô","エペソ人への手紙","エペソ","エペソ人","エペソ人の手紙","에베소서":
				bible_short_name = "弗"
				switch chap_string {
					case "":
						print_string = "以弗所書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("弗","以弗所書", chap_string, "1","web")
							default:
								print_string = Bible_print_string("弗","以弗所書", chap_string, sec_string,"web")
						}
				}
			case "Phil","Philippians","腓","腓立","腓立比","腓立比書","Php","php","빌립보서","ピリピ","ピリピ人.ピリピ人への手紙","Послание к Филиппийцам","Phi-líp":
				bible_short_name = "腓"
				switch chap_string {
					case "":
						print_string = "腓立比書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("腓","腓立比書", chap_string, "1","web")
							default:
								print_string = Bible_print_string("腓","腓立比書", chap_string, sec_string,"web")
						}
				}
			case "Col","col","Colossians","西","歌羅西","歌羅","歌羅西書","Послание к Колоссянам","Cô-lô-se","コロサイ人への手紙","コロサイ","コロ","골로새서":
				bible_short_name = "西"
				switch chap_string {
					case "":
						print_string = "歌羅西書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("西","歌羅西書", chap_string, "1","web")
							default:
								print_string = Bible_print_string("西","歌羅西書", chap_string, sec_string,"web")
						}
				}
			case "1 Thess","First Thessalonians","帖前","帖撒羅尼迦前","帖撒羅尼迦前書","1Th","1th","데살로니가전서","テサロニケ人への第一の手紙","テサ一","テサロニケ一","1 Tê-sa-lô-ni-ca","Первое послание к Фессалоникийцам (Солунянам)":
				bible_short_name = "帖前"
				switch chap_string {
					case "":
						print_string = "帖撒羅尼迦前書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("帖前","帖撒羅尼迦前書", chap_string, "1","web")
							default:
								print_string = Bible_print_string("帖前","帖撒羅尼迦前書", chap_string, sec_string,"web")
						}
				}
			case "2 Thess","Second Thessalonians","帖後","帖撒羅尼迦後","帖撒羅尼迦後書","2Th","2th","데살로니가후서","テサロニケ人への第二の手紙","テサ二","テサロニケ二","2 Tê-sa-lô-ni-ca","Второе послание к Фессалоникийцам (Солунянам)":
				bible_short_name = "帖後"
				switch chap_string {
					case "":
						print_string = "帖撒羅尼迦後書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("帖後","帖撒羅尼迦後書", chap_string, "1","web")
							default:
								print_string = Bible_print_string("帖後","帖撒羅尼迦後書", chap_string, sec_string,"web")
						}
				}
			case "1 Tim","First Timothy","提前","提摩太前","提摩太前書","1Ti","1ti","Первое послание к Тимофею","1 Ti-mô-thê","テモテヘの第一の手紙","テモテ一","디모데전서":
				bible_short_name = "提前"
				switch chap_string {
					case "":
						print_string = "提摩太前書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("提前","提摩太前書", chap_string, "1","web")
							default:
								print_string = Bible_print_string("提前","提摩太前書", chap_string, sec_string,"web")
						}
				}
			case "2 Tim","Second Timothy","提後","提摩太後","提摩太後書","2Ti","2ti","Второе послание к Тимофею","2 Ti-mô-thê","テモテヘの第二の手紙","テモテ二","디모데후서":
				bible_short_name = "提後"
				switch chap_string {
					case "":
						print_string = "提摩太後書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("提後","提摩太後書", chap_string, "1","web")
							default:
								print_string = Bible_print_string("提後","提摩太後書", chap_string, sec_string,"web")
						}
				}
			case "Titus","多","提多","提多書","Tit","tit","Послание к Титу","Tít","テトスヘの手紙","テトス","디도서":
				bible_short_name = "多"
				switch chap_string {
					case "":
						print_string = "提多書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("多","提多書", chap_string, "1","web")
							default:
								print_string = Bible_print_string("多","提多書", chap_string, sec_string,"web")
						}
				}
			case "Philem","Philemon","門","腓利","腓利門","腓利門書","Phm","phm","Послание к Филимону","Phi-lê-môn","ピレモンヘの手紙","ピレモン","빌레몬서":
				bible_short_name = "門"
				switch chap_string {
					case "":
						print_string = "腓利門書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("門","腓利門書", chap_string, "1","web")
							default:
								print_string = Bible_print_string("門","腓利門書", chap_string, sec_string,"web")
						}
				}
			case "Heb","Hebrews","來","希伯來","希伯來書","heb","Послание к Евреям","Hê-bơ-rơ","ヘブル人への手紙","ヘブル","히브리서":
				bible_short_name = "來"
				switch chap_string {
					case "":
						print_string = "希伯來書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("來","希伯來書", chap_string, "1","web")
							default:
								print_string = Bible_print_string("來","希伯來書", chap_string, sec_string,"web")
						}
				}
			case "James","雅","雅各","雅各書","Jas","jas","Послание Иакова","Gia-cơ","ヤコブの手紙","ヤコブ","야고보서":
				bible_short_name = "雅"
				switch chap_string {
					case "":
						print_string = "雅各書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("雅","雅各書", chap_string, "1","web")
							default:
								print_string = Bible_print_string("雅","雅各書", chap_string, sec_string,"web")
						}
				}
			case "1 Pet","First Peter","彼前","彼得前","彼得前書","1Pe","1pe","Первое послание Петра","1 Phi-e-rơ","ペテロの第一の手紙","ペテロ一","베드로전서":
				bible_short_name = "彼前"
				switch chap_string {
					case "":
						print_string = "彼得前書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("彼前","彼得前書", chap_string, "1","web")
							default:
								print_string = Bible_print_string("彼前","彼得前書", chap_string, sec_string,"web")
						}
				}
			case "2 Pet","Second Peter","彼後","彼得後","彼得後書","2Pe","2pe","Второе послание Петра","2 Phi-e-rơ","ペテロの第二の手紙","ペテロ","베드로후서":
				bible_short_name = "彼後"
				switch chap_string {
					case "":
						print_string = "彼得後書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("彼後","彼得後書", chap_string, "1","web")
							default:
								print_string = Bible_print_string("彼後","彼得後書", chap_string, sec_string,"web")
						}
				}
			case "1 John","First John","約一","約翰一書","約翰1","約翰1書","1Jo","1jo","Первое послание Иоанна","1 Giăng","ヨハネの第一の手紙","ヨハネ一","요한일서":
				bible_short_name = "約一"
				switch chap_string {
					case "":
						print_string = "約翰一書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("約一","約翰一書", chap_string, "1","web")
							default:
								print_string = Bible_print_string("約一","約翰一書", chap_string, sec_string,"web")
						}
				}
			case "2 John","second John","約二","約翰二書","約翰2","約翰2書","2Jo","Второе послание Иоанна","2 Giăng","ヨハネの第二の手紙","ヨハネ二","요한2서":
				bible_short_name = "約二"
				switch chap_string {
					case "":
						print_string = "約翰二書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("約二","約翰二書", chap_string, "1","web")
							default:
								print_string = Bible_print_string("約二","約翰二書", chap_string, sec_string,"web")
						}
				}
			case "3 John","Third John","約三","約翰三書","約翰3","約翰3書","3Jo","3jo","Третье послание Иоанна","3 Giăng","ヨハネの第三の手紙","ヨハネ三","요한3서":
				bible_short_name = "約三"
				switch chap_string {
					case "":
						print_string = "約翰三書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("約三","約翰三書", chap_string, "1","web")
							default:
								print_string = Bible_print_string("約三","約翰三書", chap_string, sec_string,"web")
						}
				}
			case "Jude","猶","猶大","猶大書","jude","Послание Иуды","Giu-đe","ユダの手紙","ユダ","유다서":
				bible_short_name = "猶"
				switch chap_string {
					case "":
						print_string = "猶大書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("猶","猶大書", chap_string, "1","web")
							default:
								print_string = Bible_print_string("猶","猶大書", chap_string, sec_string,"web")
						}
				}
			case "Rev","Revelation","啟","啟示","啟示錄","Re","re","ｒｅ","Ｒｅ","rev","Откровение ап. Иоанна Богослова (Апокалипсис)","Khải-huyền","ヨハネの黙示録","黙示録","요한계시록":
				bible_short_name = "啟"
				switch chap_string {
					case "":
						print_string = "啟示錄"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("啟","啟示錄", chap_string, "1","web")
							default:
								print_string = Bible_print_string("啟","啟示錄", chap_string, sec_string,"web")
						}
				}
			default:
				print_string = "聖經"
				//print_string = "你是要找 " +  reg.ReplaceAllString(text, "$3") + " 對嗎？\n對不起，我還沒學呢...\n"
		}
	case "英文聖經BBE","BBE","Bbe","bbe","ＢＢＥ","Ｂｂｅ","ｂｂｅ":
		log.Print(reg.ReplaceAllString(text, "$3"))
		switch reg.ReplaceAllString(text, "$3") {
			case "Gen","Genesis","創","創世","創世紀","創世記","Ge","ge","gen","창세기","Sáng-thế Ký","Бытие":
				bible_short_name = "創"
				switch chap_string {
					case "":
						print_string = "創世記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("創","創世記", chap_string, "1","bbe")
							default:
								print_string = Bible_print_string("創","創世記", chap_string, sec_string,"bbe")
						}
				}
			case "ex","Ex","Exodus","埃及","出","出埃及","出埃及記","출애굽기","エジプト","出エジプト","出エジプト記","Xuất Ê-díp-tô Ký","Исход":
				bible_short_name = "出"
				switch chap_string {
					case "":
						print_string = "出埃及記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("出","出埃及記", chap_string, "1","bbe")
							default:
								print_string = Bible_print_string("出","出埃及記", chap_string, sec_string,"bbe")
						}
				}
			case "Lev","Leviticus","利","利未","利未記","Le","le","Левит","Lê-vi Ký","レビ記","レビ","레위기":
				bible_short_name = "利"
				switch chap_string {
					case "":
						print_string = "利未記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("利","利未記", chap_string, "1","bbe")
							default:
								print_string = Bible_print_string("利","利未記", chap_string, sec_string,"bbe")
						}
				}
			case "Num","Numbers","民","民數","民數記","Nu","nu","민수기","民数","民数記","Dân-số Ký","Числа":
				bible_short_name = "民"
				switch chap_string {
					case "":
						print_string = "民數記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("民","民數記", chap_string, "1","bbe")
							default:
								print_string = Bible_print_string("民","民數記", chap_string, sec_string,"bbe")
						}
				}
			case "Deut","Deuteronomy","申","申命","申命記","De","de","신명기","Phục-truyền Luật-lệ Ký","Второзаконие":
				bible_short_name = "申"
				switch chap_string {
					case "":
						print_string = "申命記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("申","申命記", chap_string, "1","bbe")
							default:
								print_string = Bible_print_string("申","申命記", chap_string, sec_string,"bbe")
						}
				}
			case "Josh","Joshua","約書亞","約書亞記","Jos","jos","여호수아","ヨシュア記","ヨシュア","Giô-suê","Книга Иисуса Навина":
				bible_short_name = "書"
				switch chap_string {
					case "":
						print_string = "約書亞記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("書","約書亞記", chap_string, "1","bbe")
							default:
								print_string = Bible_print_string("書","約書亞記", chap_string, sec_string,"bbe")
						}
				}
			case "Judg","Judges","士","士師","士師記","Jud","jud","jdg","Jdg","Книга Судей израилевых","Các Quan Xét","사사기":
				bible_short_name = "士"
				switch chap_string {
					case "":
						print_string = "士師記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("士","士師記", chap_string, "1","bbe")
							default:
								print_string = Bible_print_string("士","士師記", chap_string, sec_string,"bbe")
						}
				}
			case "Ruth","路得","路得記","Ru","ru","Rut","rut","룻기","ルツ","ルツ記","Ru-tơ","Книга Руфи":
				bible_short_name = "得"
				switch chap_string {
					case "":
						print_string = "路得記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("得","路得記", chap_string, "1","bbe")
							default:
								print_string = Bible_print_string("得","路得記", chap_string, sec_string,"bbe")
						}
				}
			case "1 Sam","First Samuel","撒上","撒母耳記上","1Sa","1sa","サムエル記上","サムエル上","サム上","사무엘상","1 Sa-mu-ên","Первая книга Царств":
				bible_short_name = "撒上"
				switch chap_string {
					case "":
						print_string = "撒母耳記上"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("撒上","撒母耳記上", chap_string, "1","bbe")
							default:
								print_string = Bible_print_string("撒上","撒母耳記上", chap_string, sec_string,"bbe")
						}
				}
			case "2 Sam","Second Samuel","撒下","撒母耳記下","2Sa","2sa","사무엘하","サムエル記下","サムエル下","サム下","2 Sa-mu-ên","Вторая книга Царств":
				bible_short_name = "撒下"
				switch chap_string {
					case "":
						print_string = "撒母耳記下"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("撒下","撒母耳記下", chap_string, "1","bbe")
							default:
								print_string = Bible_print_string("撒下","撒母耳記下", chap_string, sec_string,"bbe")
						}
				}
			case "1 Kin","First Kings","王上","列王上","列王紀上","列王記上","1Ki","1ki","열왕기상","Третья книга Царств","1 Các Vua":
				bible_short_name = "王上"
				switch chap_string {
					case "":
						print_string = "列王紀上"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("王上","列王紀上", chap_string, "1","bbe")
							default:
								print_string = Bible_print_string("王上","列王紀上", chap_string, sec_string,"bbe")
						}
				}
			case "2 Kin","Second Kings","王下","列王下","列王記下","列王紀下","2Ki","2ki","열왕기하","2 Các Vua","Четвертая книга Царств":
				bible_short_name = "王下"
				switch chap_string {
					case "":
						print_string = "列王紀下"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("王下","列王紀下", chap_string, "1","bbe")
							default:
								print_string = Bible_print_string("王下","列王紀下", chap_string, sec_string,"bbe")
						}
				}
			case "1 Chr","First Chronicles","歷上","代上","歷代志上","歷代上","1Ch","1ch","Первая книга Паралипоменон","1 Sử-ký","歴上","歴代上","歴代志上","역대상":
				bible_short_name = "代上"
				switch chap_string {
					case "":
						print_string = "歷代志上"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("代上","歷代志上", chap_string, "1","bbe")
							default:
								print_string = Bible_print_string("代上","歷代志上", chap_string, sec_string,"bbe")
						}
				}
			case "2 Chr","Second Chronicles","代下","歷下","歷代下","歷代志下","2Ch","2ch","역대하","歴代志下","歴代下","歴下","2 Sử-ký","Вторая книга Паралипоменон":
				bible_short_name = "代下"
				switch chap_string {
					case "":
						print_string = "歷代志下"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("代下","歷代志下", chap_string, "1","bbe")
							default:
								print_string = Bible_print_string("代下","歷代志下", chap_string, sec_string,"bbe")
						}
				}
			case "Ezra","拉","以斯拉","以斯拉記","Ezr","ezr","Первая книга Ездры","Ê-xơ-ra","エズラ","エズラ記","에스라":
				bible_short_name = "拉"
				switch chap_string {
					case "":
						print_string = "以斯拉記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("拉","以斯拉記", chap_string, "1","bbe")
							default:
								print_string = Bible_print_string("拉","以斯拉記", chap_string, sec_string,"bbe")
						}
				}
			case "Neh","Nehemiah","尼","尼希米","尼希米記","Ne","ne","느헤미야","ネヘミヤ書","ネヘミヤ","Nê-hê-mi","Книга Неемии":
				bible_short_name = "尼"
				switch chap_string {
					case "":
						print_string = "尼希米記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("尼","尼希米記", chap_string, "1","bbe")
							default:
								print_string = Bible_print_string("尼","尼希米記", chap_string, sec_string,"bbe")
						}
				}
			case "Esth","Esther","斯","以斯帖","以斯帖記","Es","est","Есфирь","Ê-xơ-tê","エステル","エステル記","에스더":
				bible_short_name = "斯"
				switch chap_string {
					case "":
						print_string = "以斯帖記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("斯","以斯帖記", chap_string, "1","bbe")
							default:
								print_string = Bible_print_string("斯","以斯帖記", chap_string, sec_string,"bbe")
						}
				}
			case "Job","job","伯","約伯","約伯記","Книга Иова","Gióp","ヨブ","ヨブ記","욥기":
				bible_short_name = "伯"
				switch chap_string {
					case "":
						print_string = "約伯記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("伯","約伯記", chap_string, "1","bbe")
							default:
								print_string = Bible_print_string("伯","約伯記", chap_string, sec_string,"bbe")
						}
				}
			case "Ps","Psalms","詩","詩篇","ps","시편","Thi-thiên","Псалтирь":
				bible_short_name = "詩"
				switch chap_string {
					case "":
						print_string = "詩篇"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("詩","詩篇", chap_string, "1","bbe")
							default:
								print_string = Bible_print_string("詩","詩篇", chap_string, sec_string,"bbe")
						}
				}
			case "Prov","Proverbs","箴","箴言","Pr","pr","Притчи Соломона","Châm-ngôn","잠언":
				bible_short_name = "箴"
				switch chap_string {
					case "":
						print_string = "箴言"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("箴","箴言", chap_string, "1","bbe")
							default:
								print_string = Bible_print_string("箴","箴言", chap_string, sec_string,"bbe")
						}
				}
			case "Eccl","Ecclesiastes","傳","傳道","傳道書","Ec","ec","Книга Екклезиаста","или Проповедника","Truyền-đạo","伝道の書","伝道","伝","伝道書","전도서":
				bible_short_name = "傳"
				switch chap_string {
					case "":
						print_string = "傳道書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("傳","傳道書", chap_string, "1","bbe")
							default:
								print_string = Bible_print_string("傳","傳道書", chap_string, sec_string,"bbe")
						}
				}
			case "Song","Song of Solomon","歌","雅歌","So","so","sng","Sng","Песнь песней Соломона","Nhã-ca","아가":
				bible_short_name = "歌"
				switch chap_string {
					case "":
						print_string = "雅歌"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("歌","雅歌", chap_string, "1","bbe")
							default:
								print_string = Bible_print_string("歌","雅歌", chap_string, sec_string,"bbe")
						}
				}
			case "Is","Isaiah","賽","以賽","以賽亞","以賽亞書","Isa","isa","Книга пророка Исаии","Ê-sai","イザヤ書","イザヤ","이사야":
				bible_short_name = "賽"
				switch chap_string {
					case "":
						print_string = "以賽亞書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("賽","以賽亞書", chap_string, "1","bbe")
							default:
								print_string = Bible_print_string("賽","以賽亞書", chap_string, sec_string,"bbe")
						}
				}
			case "Jer","Jeremiah","耶","耶利米","耶利米書","jer","예레미야","エレミヤ","エレミヤ書","Giê-rê-mi","Книга пророка Иеремии":
				bible_short_name = "耶"
				switch chap_string {
					case "":
						print_string = "耶利米書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("耶","耶利米書", chap_string, "1","bbe")
							default:
								print_string = Bible_print_string("耶","耶利米書", chap_string, sec_string,"bbe")
						}
				}
			case "Lam","Lamentations","哀","哀歌","耶利米哀歌","La","lam","예레미야애가","Ca-thương","Плач Иеремии":
				bible_short_name = "哀"
				switch chap_string {
					case "":
						print_string = "耶利米哀歌"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("哀","耶利米哀歌", chap_string, "1","bbe")
							default:
								print_string = Bible_print_string("哀","耶利米哀歌", chap_string, sec_string,"bbe")
						}
				}
			case "Ezek","Ezekiel","結","以西結","以西結書","Eze","eze","에스겔","エゼキエル書","エゼキエル","Ê-xê-chi-ên","Книга пророка Иезекииля":
				bible_short_name = "結"
				switch chap_string {
					case "":
						print_string = "以西結書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("結","以西結書", chap_string, "1","bbe")
							default:
								print_string = Bible_print_string("結","以西結書", chap_string, sec_string,"bbe")
						}
				}
			case "Dan","Daniel","但","但以理","但以理書","Da","da","Книга пророка Даниила","Đa-ni-ên","ダニエル書","ダニエル","다니엘":
				bible_short_name = "但"
				switch chap_string {
					case "":
						print_string = "但以理書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("但","但以理書", chap_string, "1","bbe")
							default:
								print_string = Bible_print_string("但","但以理書", chap_string, sec_string,"bbe")
						}
				}
			case "Hos","Hosea","何","何西","何西阿","何西阿書","Ho","ho","Книга пророка Осии","Ô-sê","ホセア書","ホセア","호세아":
				bible_short_name = "何"
				switch chap_string {
					case "":
						print_string = "何西阿書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("何","何西阿書", chap_string, "1","bbe")
							default:
								print_string = Bible_print_string("何","何西阿書", chap_string, sec_string,"bbe")
						}
				}
			case "Joel","珥","約珥","約珥書","Joe","joe","Книга пророка Иоиля","Giô-ên","ヨエル書","ヨエル","요엘":
				bible_short_name = "珥"
				switch chap_string {
					case "":
						print_string = "約珥書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("珥","約珥書", chap_string, "1","bbe")
							default:
								print_string = Bible_print_string("珥","約珥書", chap_string, sec_string,"bbe")
						}
				}
			case "Amos","摩","阿摩司書","Am","am","Книга пророка Амоса","A-mốt","アモス書","アモス","아모스":
				bible_short_name = "摩"
				switch chap_string {
					case "":
						print_string = "阿摩司書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("摩","阿摩司書", chap_string, "1","bbe")
							default:
								print_string = Bible_print_string("摩","阿摩司書", chap_string, sec_string,"bbe")
						}
				}
			case "Obad","Obadiah","俄","俄巴底亞","俄巴底亞書","Ob","ob","오바댜","オバデヤ書","オバデヤ","Áp-đia","Книга пророка Авдия":
				bible_short_name = "俄"
				switch chap_string {
					case "":
						print_string = "俄巴底亞書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("俄","俄巴底亞書", chap_string, "1","bbe")
							default:
								print_string = Bible_print_string("俄","俄巴底亞書", chap_string, sec_string,"bbe")
						}
				}
			case "Jon","Jonah","拿","約拿","約拿書","jon","요나","ヨナ書","ヨナ","Giô-na","Книга пророка Ионы":
				bible_short_name = "拿"
				switch chap_string {
					case "":
						print_string = "約拿書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("拿","約拿書", chap_string, "1","bbe")
							default:
								print_string = Bible_print_string("拿","約拿書", chap_string, sec_string,"bbe")
						}
				}
			case "Micah","彌","彌迦","彌迦書","Mic","mic","Книга пророка Михея","Mi-chê","ミカ書","ミカ","미가":
				bible_short_name = "彌"
				switch chap_string {
					case "":
						print_string = "彌迦書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("彌","彌迦書", chap_string, "1","bbe")
							default:
								print_string = Bible_print_string("彌","彌迦書", chap_string, sec_string,"bbe")
						}
				}
			case "Nah","Nahum","鴻","那鴻","那鴻書","Na","na","Книга пророка Наума","Na-hum","ナホム書","ナホム","나훔":
				bible_short_name = "鴻"
				switch chap_string {
					case "":
						print_string = "那鴻書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("鴻","那鴻書", chap_string, "1","bbe")
							default:
								print_string = Bible_print_string("鴻","那鴻書", chap_string, sec_string,"bbe")
						}
				}
			case "Habakkuk","哈","哈巴","哈巴谷","哈巴谷書","Hab","hab","Книга пророка Аввакума","Ha-ba-cúc","ハバクク書","ハバクク","ハバ","クク","ハバ書","하박국":
				bible_short_name = "哈"
				switch chap_string {
					case "":
						print_string = "哈巴谷書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("哈","哈巴谷書", chap_string, "1","bbe")
							default:
								print_string = Bible_print_string("哈","哈巴谷書", chap_string, sec_string,"bbe")
						}
				}
			case "Zeph","Zephaniah","番","西番雅","西番雅書","Zep","zep","스바냐","ゼパニヤ書","ゼパニヤ","Sô-phô-ni","Книга пророка Софонии":
				bible_short_name = "番"
				switch chap_string {
					case "":
						print_string = "西番雅書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("番","西番雅書", chap_string, "1","bbe")
							default:
								print_string = Bible_print_string("番","西番雅書", chap_string, sec_string,"bbe")
						}
				}
			case "Haggai","該","哈該","哈該書","Hag","hag","학개","ハガイ書","ハガイ","A-ghê","Книга пророка Аггея":
				bible_short_name = "該"
				switch chap_string {
					case "":
						print_string = "哈該書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("該","哈該書", chap_string, "1","bbe")
							default:
								print_string = Bible_print_string("該","哈該書", chap_string, sec_string,"bbe")
						}
				}
			case "Zech","Zechariah","亞","撒迦利亞","撒迦利亞書","Zec","zec","Книга пророка Захарии","Xa-cha-ri","스가랴","ゼカリヤ書","ゼカリヤ":
				bible_short_name = "亞"
				switch chap_string {
					case "":
						print_string = "撒迦利亞書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("亞","撒迦利亞書", chap_string, "1","bbe")
							default:
								print_string = Bible_print_string("亞","撒迦利亞書", chap_string, sec_string,"bbe")
						}
				}
			case "Malachi","瑪","","瑪拉","瑪拉基","瑪拉基書","Mal","mal","말라기","マラキ書","マラキ","Ma-la-chi","Книга пророка Малахии":
				bible_short_name = "瑪"
				switch chap_string {
					case "":
						print_string = "瑪拉基書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("瑪","瑪拉基書", chap_string, "1","bbe")
							default:
								print_string = Bible_print_string("瑪","瑪拉基書", chap_string, sec_string,"bbe")
						}
				}
			case "Matt","Matthew","太","馬太","馬太福音","Mt","mt","마태복음","マタイによる福音書","マタイ","マタイによる","Ma-thi-ơ","От Матфея святое благовествование":
				bible_short_name = "太"
				switch chap_string {
					case "":
						print_string = "馬太福音"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("太","馬太福音", chap_string, "1","bbe")
							default:
								print_string = Bible_print_string("太","馬太福音", chap_string, sec_string,"bbe")
						}
				}
			case "Mark","可","馬可","馬可福音","Mr","mr","マルコによる福音書","マルコ","マルコによる","마가복음","Mác","От Марка святое благовествование":
				bible_short_name = "可"
				switch chap_string {
					case "":
						print_string = "馬可福音"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("可","馬可福音", chap_string, "1","bbe")
							default:
								print_string = Bible_print_string("可","馬可福音", chap_string, sec_string,"bbe")
						}
				}
			case "Luke","路","路加","路加福音","Lu","lu","От Луки святое благовествование","Lu-ca","ルカによる福音書","ルカ","ルカによる","누가복음":
				bible_short_name = "路"
				switch chap_string {
					case "":
						print_string = "路加福音"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("路","路加福音", chap_string, "1","bbe")
							default:
								print_string = Bible_print_string("路","路加福音", chap_string, sec_string,"bbe")
						}
				}
			case "John","約","約翰","約翰福音","Joh","joh","От Иоанна святое благовествование","Giăng","ヨハネによる福音書","ヨハネ","ヨハネによる","요한복음":
				bible_short_name = "約"
				switch chap_string {
					case "":
						print_string = "約翰福音"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("約","約翰福音", chap_string, "1","bbe")
							default:
								print_string = Bible_print_string("約","約翰福音", chap_string, sec_string,"bbe")
						}
				}
			case "Acts","徒","使徒","使徒行傳","Ac","ac","Деяния святых апостолов","Công-vụ Các Sứ-đồ","使徒行伝","사도행전":
				bible_short_name = "徒"
				switch chap_string {
					case "":
						print_string = "使徒行傳"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("徒","使徒行傳", chap_string, "1","bbe")
							default:
								print_string = Bible_print_string("徒","使徒行傳", chap_string, sec_string,"bbe")
						}
				}
			case "Rom","Romans","羅","羅馬","羅馬書","Ro","ro","Послание к Римлянам","Rô-ma","ローマ","ローマ人への手紙","로마서":
				bible_short_name = "羅"
				switch chap_string {
					case "":
						print_string = "羅馬書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("羅","羅馬書", chap_string, "1","bbe")
							default:
								print_string = Bible_print_string("羅","羅馬書", chap_string, sec_string,"bbe")
						}
				}
			case "1 Cor","First Corinthians","林前","哥林多前","哥林多前書","1Co","1co","Первое послание к Коринфянам","1 Cô-rinh-tô","コリント人への第一の手紙","コリント一","コリント人への第一","고린도전서":
				bible_short_name = "林前"
				switch chap_string {
					case "":
						print_string = "哥林多前書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("林前","哥林多前書", chap_string, "1","bbe")
							default:
								print_string = Bible_print_string("林前","哥林多前書", chap_string, sec_string,"bbe")
						}
				}
			case "2 Cor","Second Corinthians","林後","哥林多後","哥林多後書","2Co","2co","Второе послание к Коринфянам","2 Cô-rinh-tô","コリント人への第二の手紙","コリント二","コリント人への第二の","고린도후서":
				bible_short_name = "林後"
				switch chap_string {
					case "":
						print_string = "哥林多後書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("林後","哥林多後書", chap_string, "1","bbe")
							default:
								print_string = Bible_print_string("林後","哥林多後書", chap_string, sec_string,"bbe")
						}
				}
			case "Gal","Galatians","加","加拉太","加拉太書","Ga","ga","Послание к Галатам","Ga-la-ti","ガラテヤ","ガラテヤ人への手紙","갈라디아서":
				bible_short_name = "加"
				switch chap_string {
					case "":
						print_string = "加拉太書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("加","加拉太書", chap_string, "1","bbe")
							default:
								print_string = Bible_print_string("加","加拉太書", chap_string, sec_string,"bbe")
						}
				}
			case "Ephesians","弗","以弗所","以弗所書","Eph","eph","Послание к Ефесянам","Ê-phê-sô","エペソ人への手紙","エペソ","エペソ人","エペソ人の手紙","에베소서":
				bible_short_name = "弗"
				switch chap_string {
					case "":
						print_string = "以弗所書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("弗","以弗所書", chap_string, "1","bbe")
							default:
								print_string = Bible_print_string("弗","以弗所書", chap_string, sec_string,"bbe")
						}
				}
			case "Phil","Philippians","腓","腓立","腓立比","腓立比書","Php","php","빌립보서","ピリピ","ピリピ人.ピリピ人への手紙","Послание к Филиппийцам","Phi-líp":
				bible_short_name = "腓"
				switch chap_string {
					case "":
						print_string = "腓立比書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("腓","腓立比書", chap_string, "1","bbe")
							default:
								print_string = Bible_print_string("腓","腓立比書", chap_string, sec_string,"bbe")
						}
				}
			case "Col","col","Colossians","西","歌羅西","歌羅","歌羅西書","Послание к Колоссянам","Cô-lô-se","コロサイ人への手紙","コロサイ","コロ","골로새서":
				bible_short_name = "西"
				switch chap_string {
					case "":
						print_string = "歌羅西書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("西","歌羅西書", chap_string, "1","bbe")
							default:
								print_string = Bible_print_string("西","歌羅西書", chap_string, sec_string,"bbe")
						}
				}
			case "1 Thess","First Thessalonians","帖前","帖撒羅尼迦前","帖撒羅尼迦前書","1Th","1th","데살로니가전서","テサロニケ人への第一の手紙","テサ一","テサロニケ一","1 Tê-sa-lô-ni-ca","Первое послание к Фессалоникийцам (Солунянам)":
				bible_short_name = "帖前"
				switch chap_string {
					case "":
						print_string = "帖撒羅尼迦前書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("帖前","帖撒羅尼迦前書", chap_string, "1","bbe")
							default:
								print_string = Bible_print_string("帖前","帖撒羅尼迦前書", chap_string, sec_string,"bbe")
						}
				}
			case "2 Thess","Second Thessalonians","帖後","帖撒羅尼迦後","帖撒羅尼迦後書","2Th","2th","데살로니가후서","テサロニケ人への第二の手紙","テサ二","テサロニケ二","2 Tê-sa-lô-ni-ca","Второе послание к Фессалоникийцам (Солунянам)":
				bible_short_name = "帖後"
				switch chap_string {
					case "":
						print_string = "帖撒羅尼迦後書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("帖後","帖撒羅尼迦後書", chap_string, "1","bbe")
							default:
								print_string = Bible_print_string("帖後","帖撒羅尼迦後書", chap_string, sec_string,"bbe")
						}
				}
			case "1 Tim","First Timothy","提前","提摩太前","提摩太前書","1Ti","1ti","Первое послание к Тимофею","1 Ti-mô-thê","テモテヘの第一の手紙","テモテ一","디모데전서":
				bible_short_name = "提前"
				switch chap_string {
					case "":
						print_string = "提摩太前書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("提前","提摩太前書", chap_string, "1","bbe")
							default:
								print_string = Bible_print_string("提前","提摩太前書", chap_string, sec_string,"bbe")
						}
				}
			case "2 Tim","Second Timothy","提後","提摩太後","提摩太後書","2Ti","2ti","Второе послание к Тимофею","2 Ti-mô-thê","テモテヘの第二の手紙","テモテ二","디모데후서":
				bible_short_name = "提後"
				switch chap_string {
					case "":
						print_string = "提摩太後書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("提後","提摩太後書", chap_string, "1","bbe")
							default:
								print_string = Bible_print_string("提後","提摩太後書", chap_string, sec_string,"bbe")
						}
				}
			case "Titus","多","提多","提多書","Tit","tit","Послание к Титу","Tít","テトスヘの手紙","テトス","디도서":
				bible_short_name = "多"
				switch chap_string {
					case "":
						print_string = "提多書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("多","提多書", chap_string, "1","bbe")
							default:
								print_string = Bible_print_string("多","提多書", chap_string, sec_string,"bbe")
						}
				}
			case "Philem","Philemon","門","腓利","腓利門","腓利門書","Phm","phm","Послание к Филимону","Phi-lê-môn","ピレモンヘの手紙","ピレモン","빌레몬서":
				bible_short_name = "門"
				switch chap_string {
					case "":
						print_string = "腓利門書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("門","腓利門書", chap_string, "1","bbe")
							default:
								print_string = Bible_print_string("門","腓利門書", chap_string, sec_string,"bbe")
						}
				}
			case "Heb","Hebrews","來","希伯來","希伯來書","heb","Послание к Евреям","Hê-bơ-rơ","ヘブル人への手紙","ヘブル","히브리서":
				bible_short_name = "來"
				switch chap_string {
					case "":
						print_string = "希伯來書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("來","希伯來書", chap_string, "1","bbe")
							default:
								print_string = Bible_print_string("來","希伯來書", chap_string, sec_string,"bbe")
						}
				}
			case "James","雅","雅各","雅各書","Jas","jas","Послание Иакова","Gia-cơ","ヤコブの手紙","ヤコブ","야고보서":
				bible_short_name = "雅"
				switch chap_string {
					case "":
						print_string = "雅各書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("雅","雅各書", chap_string, "1","bbe")
							default:
								print_string = Bible_print_string("雅","雅各書", chap_string, sec_string,"bbe")
						}
				}
			case "1 Pet","First Peter","彼前","彼得前","彼得前書","1Pe","1pe","Первое послание Петра","1 Phi-e-rơ","ペテロの第一の手紙","ペテロ一","베드로전서":
				bible_short_name = "彼前"
				switch chap_string {
					case "":
						print_string = "彼得前書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("彼前","彼得前書", chap_string, "1","bbe")
							default:
								print_string = Bible_print_string("彼前","彼得前書", chap_string, sec_string,"bbe")
						}
				}
			case "2 Pet","Second Peter","彼後","彼得後","彼得後書","2Pe","2pe","Второе послание Петра","2 Phi-e-rơ","ペテロの第二の手紙","ペテロ","베드로후서":
				bible_short_name = "彼後"
				switch chap_string {
					case "":
						print_string = "彼得後書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("彼後","彼得後書", chap_string, "1","bbe")
							default:
								print_string = Bible_print_string("彼後","彼得後書", chap_string, sec_string,"bbe")
						}
				}
			case "1 John","First John","約一","約翰一書","約翰1","約翰1書","1Jo","1jo","Первое послание Иоанна","1 Giăng","ヨハネの第一の手紙","ヨハネ一","요한일서":
				bible_short_name = "約一"
				switch chap_string {
					case "":
						print_string = "約翰一書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("約一","約翰一書", chap_string, "1","bbe")
							default:
								print_string = Bible_print_string("約一","約翰一書", chap_string, sec_string,"bbe")
						}
				}
			case "2 John","second John","約二","約翰二書","約翰2","約翰2書","2Jo","Второе послание Иоанна","2 Giăng","ヨハネの第二の手紙","ヨハネ二","요한2서":
				bible_short_name = "約二"
				switch chap_string {
					case "":
						print_string = "約翰二書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("約二","約翰二書", chap_string, "1","bbe")
							default:
								print_string = Bible_print_string("約二","約翰二書", chap_string, sec_string,"bbe")
						}
				}
			case "3 John","Third John","約三","約翰三書","約翰3","約翰3書","3Jo","3jo","Третье послание Иоанна","3 Giăng","ヨハネの第三の手紙","ヨハネ三","요한3서":
				bible_short_name = "約三"
				switch chap_string {
					case "":
						print_string = "約翰三書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("約三","約翰三書", chap_string, "1","bbe")
							default:
								print_string = Bible_print_string("約三","約翰三書", chap_string, sec_string,"bbe")
						}
				}
			case "Jude","猶","猶大","猶大書","jude","Послание Иуды","Giu-đe","ユダの手紙","ユダ","유다서":
				bible_short_name = "猶"
				switch chap_string {
					case "":
						print_string = "猶大書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("猶","猶大書", chap_string, "1","bbe")
							default:
								print_string = Bible_print_string("猶","猶大書", chap_string, sec_string,"bbe")
						}
				}
			case "Rev","Revelation","啟","啟示","啟示錄","Re","re","ｒｅ","Ｒｅ","rev","Откровение ап. Иоанна Богослова (Апокалипсис)","Khải-huyền","ヨハネの黙示録","黙示録","요한계시록":
				bible_short_name = "啟"
				switch chap_string {
					case "":
						print_string = "啟示錄"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("啟","啟示錄", chap_string, "1","bbe")
							default:
								print_string = Bible_print_string("啟","啟示錄", chap_string, sec_string,"bbe")
						}
				}
			default:
				print_string = "聖經"
				//print_string = "你是要找 " +  reg.ReplaceAllString(text, "$3") + " 對嗎？\n對不起，我還沒學呢...\n"
		}
	case "客家聖經":
		log.Print(reg.ReplaceAllString(text, "$3"))
		switch reg.ReplaceAllString(text, "$3") {

			case "Gen","Genesis","創","創世","創世紀","創世記","Ge","ge","gen","창세기","Sáng-thế Ký","Бытие":
				bible_short_name = "創"
				switch chap_string {
					case "":
						print_string = "創世記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("創","創世記", chap_string, "1","hakka")
							default:
								print_string = Bible_print_string("創","創世記", chap_string, sec_string,"hakka")
						}
				}
			case "ex","Ex","Exodus","埃及","出","出埃及","出埃及記","출애굽기","エジプト","出エジプト","出エジプト記","Xuất Ê-díp-tô Ký","Исход":
				bible_short_name = "出"
				switch chap_string {
					case "":
						print_string = "出埃及記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("出","出埃及記", chap_string, "1","hakka")
							default:
								print_string = Bible_print_string("出","出埃及記", chap_string, sec_string,"hakka")
						}
				}
			case "Lev","Leviticus","利","利未","利未記","Le","le","Левит","Lê-vi Ký","レビ記","レビ","레위기":
				bible_short_name = "利"
				switch chap_string {
					case "":
						print_string = "利未記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("利","利未記", chap_string, "1","hakka")
							default:
								print_string = Bible_print_string("利","利未記", chap_string, sec_string,"hakka")
						}
				}
			case "Num","Numbers","民","民數","民數記","Nu","nu","민수기","民数","民数記","Dân-số Ký","Числа":
				bible_short_name = "民"
				switch chap_string {
					case "":
						print_string = "民數記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("民","民數記", chap_string, "1","hakka")
							default:
								print_string = Bible_print_string("民","民數記", chap_string, sec_string,"hakka")
						}
				}
			case "Deut","Deuteronomy","申","申命","申命記","De","de","신명기","Phục-truyền Luật-lệ Ký","Второзаконие":
				bible_short_name = "申"
				switch chap_string {
					case "":
						print_string = "申命記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("申","申命記", chap_string, "1","hakka")
							default:
								print_string = Bible_print_string("申","申命記", chap_string, sec_string,"hakka")
						}
				}
			case "Josh","Joshua","約書亞","約書亞記","Jos","jos","여호수아","ヨシュア記","ヨシュア","Giô-suê","Книга Иисуса Навина":
				bible_short_name = "書"
				switch chap_string {
					case "":
						print_string = "約書亞記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("書","約書亞記", chap_string, "1","hakka")
							default:
								print_string = Bible_print_string("書","約書亞記", chap_string, sec_string,"hakka")
						}
				}
			case "Judg","Judges","士","士師","士師記","Jud","jud","jdg","Jdg","Книга Судей израилевых","Các Quan Xét","사사기":
				bible_short_name = "士"
				switch chap_string {
					case "":
						print_string = "士師記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("士","士師記", chap_string, "1","hakka")
							default:
								print_string = Bible_print_string("士","士師記", chap_string, sec_string,"hakka")
						}
				}
			case "Ruth","路得","路得記","Ru","ru","Rut","rut","룻기","ルツ","ルツ記","Ru-tơ","Книга Руфи":
				bible_short_name = "得"
				switch chap_string {
					case "":
						print_string = "路得記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("得","路得記", chap_string, "1","hakka")
							default:
								print_string = Bible_print_string("得","路得記", chap_string, sec_string,"hakka")
						}
				}
			case "1 Sam","First Samuel","撒上","撒母耳記上","1Sa","1sa","サムエル記上","サムエル上","サム上","사무엘상","1 Sa-mu-ên","Первая книга Царств":
				bible_short_name = "撒上"
				switch chap_string {
					case "":
						print_string = "撒母耳記上"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("撒上","撒母耳記上", chap_string, "1","hakka")
							default:
								print_string = Bible_print_string("撒上","撒母耳記上", chap_string, sec_string,"hakka")
						}
				}
			case "2 Sam","Second Samuel","撒下","撒母耳記下","2Sa","2sa","사무엘하","サムエル記下","サムエル下","サム下","2 Sa-mu-ên","Вторая книга Царств":
				bible_short_name = "撒下"
				switch chap_string {
					case "":
						print_string = "撒母耳記下"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("撒下","撒母耳記下", chap_string, "1","hakka")
							default:
								print_string = Bible_print_string("撒下","撒母耳記下", chap_string, sec_string,"hakka")
						}
				}
			case "1 Kin","First Kings","王上","列王上","列王紀上","列王記上","1Ki","1ki","열왕기상","Третья книга Царств","1 Các Vua":
				bible_short_name = "王上"
				switch chap_string {
					case "":
						print_string = "列王紀上"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("王上","列王紀上", chap_string, "1","hakka")
							default:
								print_string = Bible_print_string("王上","列王紀上", chap_string, sec_string,"hakka")
						}
				}
			case "2 Kin","Second Kings","王下","列王下","列王記下","列王紀下","2Ki","2ki","열왕기하","2 Các Vua","Четвертая книга Царств":
				bible_short_name = "王下"
				switch chap_string {
					case "":
						print_string = "列王紀下"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("王下","列王紀下", chap_string, "1","hakka")
							default:
								print_string = Bible_print_string("王下","列王紀下", chap_string, sec_string,"hakka")
						}
				}
			case "1 Chr","First Chronicles","歷上","代上","歷代志上","歷代上","1Ch","1ch","Первая книга Паралипоменон","1 Sử-ký","歴上","歴代上","歴代志上","역대상":
				bible_short_name = "代上"
				switch chap_string {
					case "":
						print_string = "歷代志上"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("代上","歷代志上", chap_string, "1","hakka")
							default:
								print_string = Bible_print_string("代上","歷代志上", chap_string, sec_string,"hakka")
						}
				}
			case "2 Chr","Second Chronicles","代下","歷下","歷代下","歷代志下","2Ch","2ch","역대하","歴代志下","歴代下","歴下","2 Sử-ký","Вторая книга Паралипоменон":
				bible_short_name = "代下"
				switch chap_string {
					case "":
						print_string = "歷代志下"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("代下","歷代志下", chap_string, "1","hakka")
							default:
								print_string = Bible_print_string("代下","歷代志下", chap_string, sec_string,"hakka")
						}
				}
			case "Ezra","拉","以斯拉","以斯拉記","Ezr","ezr","Первая книга Ездры","Ê-xơ-ra","エズラ","エズラ記","에스라":
				bible_short_name = "拉"
				switch chap_string {
					case "":
						print_string = "以斯拉記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("拉","以斯拉記", chap_string, "1","hakka")
							default:
								print_string = Bible_print_string("拉","以斯拉記", chap_string, sec_string,"hakka")
						}
				}
			case "Neh","Nehemiah","尼","尼希米","尼希米記","Ne","ne","느헤미야","ネヘミヤ書","ネヘミヤ","Nê-hê-mi","Книга Неемии":
				bible_short_name = "尼"
				switch chap_string {
					case "":
						print_string = "尼希米記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("尼","尼希米記", chap_string, "1","hakka")
							default:
								print_string = Bible_print_string("尼","尼希米記", chap_string, sec_string,"hakka")
						}
				}
			case "Esth","Esther","斯","以斯帖","以斯帖記","Es","est","Есфирь","Ê-xơ-tê","エステル","エステル記","에스더":
				bible_short_name = "斯"
				switch chap_string {
					case "":
						print_string = "以斯帖記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("斯","以斯帖記", chap_string, "1","hakka")
							default:
								print_string = Bible_print_string("斯","以斯帖記", chap_string, sec_string,"hakka")
						}
				}
			case "Job","job","伯","約伯","約伯記","Книга Иова","Gióp","ヨブ","ヨブ記","욥기":
				bible_short_name = "伯"
				switch chap_string {
					case "":
						print_string = "約伯記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("伯","約伯記", chap_string, "1","hakka")
							default:
								print_string = Bible_print_string("伯","約伯記", chap_string, sec_string,"hakka")
						}
				}
			case "Ps","Psalms","詩","詩篇","ps","시편","Thi-thiên","Псалтирь":
				bible_short_name = "詩"
				switch chap_string {
					case "":
						print_string = "詩篇"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("詩","詩篇", chap_string, "1","hakka")
							default:
								print_string = Bible_print_string("詩","詩篇", chap_string, sec_string,"hakka")
						}
				}
			case "Prov","Proverbs","箴","箴言","Pr","pr","Притчи Соломона","Châm-ngôn","잠언":
				bible_short_name = "箴"
				switch chap_string {
					case "":
						print_string = "箴言"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("箴","箴言", chap_string, "1","hakka")
							default:
								print_string = Bible_print_string("箴","箴言", chap_string, sec_string,"hakka")
						}
				}
			case "Eccl","Ecclesiastes","傳","傳道","傳道書","Ec","ec","Книга Екклезиаста","или Проповедника","Truyền-đạo","伝道の書","伝道","伝","伝道書","전도서":
				bible_short_name = "傳"
				switch chap_string {
					case "":
						print_string = "傳道書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("傳","傳道書", chap_string, "1","hakka")
							default:
								print_string = Bible_print_string("傳","傳道書", chap_string, sec_string,"hakka")
						}
				}
			case "Song","Song of Solomon","歌","雅歌","So","so","sng","Sng","Песнь песней Соломона","Nhã-ca","아가":
				bible_short_name = "歌"
				switch chap_string {
					case "":
						print_string = "雅歌"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("歌","雅歌", chap_string, "1","hakka")
							default:
								print_string = Bible_print_string("歌","雅歌", chap_string, sec_string,"hakka")
						}
				}
			case "Is","Isaiah","賽","以賽","以賽亞","以賽亞書","Isa","isa","Книга пророка Исаии","Ê-sai","イザヤ書","イザヤ","이사야":
				bible_short_name = "賽"
				switch chap_string {
					case "":
						print_string = "以賽亞書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("賽","以賽亞書", chap_string, "1","hakka")
							default:
								print_string = Bible_print_string("賽","以賽亞書", chap_string, sec_string,"hakka")
						}
				}
			case "Jer","Jeremiah","耶","耶利米","耶利米書","jer","예레미야","エレミヤ","エレミヤ書","Giê-rê-mi","Книга пророка Иеремии":
				bible_short_name = "耶"
				switch chap_string {
					case "":
						print_string = "耶利米書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("耶","耶利米書", chap_string, "1","hakka")
							default:
								print_string = Bible_print_string("耶","耶利米書", chap_string, sec_string,"hakka")
						}
				}
			case "Lam","Lamentations","哀","哀歌","耶利米哀歌","La","lam","예레미야애가","Ca-thương","Плач Иеремии":
				bible_short_name = "哀"
				switch chap_string {
					case "":
						print_string = "耶利米哀歌"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("哀","耶利米哀歌", chap_string, "1","hakka")
							default:
								print_string = Bible_print_string("哀","耶利米哀歌", chap_string, sec_string,"hakka")
						}
				}
			case "Ezek","Ezekiel","結","以西結","以西結書","Eze","eze","에스겔","エゼキエル書","エゼキエル","Ê-xê-chi-ên","Книга пророка Иезекииля":
				bible_short_name = "結"
				switch chap_string {
					case "":
						print_string = "以西結書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("結","以西結書", chap_string, "1","hakka")
							default:
								print_string = Bible_print_string("結","以西結書", chap_string, sec_string,"hakka")
						}
				}
			case "Dan","Daniel","但","但以理","但以理書","Da","da","Книга пророка Даниила","Đa-ni-ên","ダニエル書","ダニエル","다니엘":
				bible_short_name = "但"
				switch chap_string {
					case "":
						print_string = "但以理書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("但","但以理書", chap_string, "1","hakka")
							default:
								print_string = Bible_print_string("但","但以理書", chap_string, sec_string,"hakka")
						}
				}
			case "Hos","Hosea","何","何西","何西阿","何西阿書","Ho","ho","Книга пророка Осии","Ô-sê","ホセア書","ホセア","호세아":
				bible_short_name = "何"
				switch chap_string {
					case "":
						print_string = "何西阿書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("何","何西阿書", chap_string, "1","hakka")
							default:
								print_string = Bible_print_string("何","何西阿書", chap_string, sec_string,"hakka")
						}
				}
			case "Joel","珥","約珥","約珥書","Joe","joe","Книга пророка Иоиля","Giô-ên","ヨエル書","ヨエル","요엘":
				bible_short_name = "珥"
				switch chap_string {
					case "":
						print_string = "約珥書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("珥","約珥書", chap_string, "1","hakka")
							default:
								print_string = Bible_print_string("珥","約珥書", chap_string, sec_string,"hakka")
						}
				}
			case "Amos","摩","阿摩司書","Am","am","Книга пророка Амоса","A-mốt","アモス書","アモス","아모스":
				bible_short_name = "摩"
				switch chap_string {
					case "":
						print_string = "阿摩司書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("摩","阿摩司書", chap_string, "1","hakka")
							default:
								print_string = Bible_print_string("摩","阿摩司書", chap_string, sec_string,"hakka")
						}
				}
			case "Obad","Obadiah","俄","俄巴底亞","俄巴底亞書","Ob","ob","오바댜","オバデヤ書","オバデヤ","Áp-đia","Книга пророка Авдия":
				bible_short_name = "俄"
				switch chap_string {
					case "":
						print_string = "俄巴底亞書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("俄","俄巴底亞書", chap_string, "1","hakka")
							default:
								print_string = Bible_print_string("俄","俄巴底亞書", chap_string, sec_string,"hakka")
						}
				}
			case "Jon","Jonah","拿","約拿","約拿書","jon","요나","ヨナ書","ヨナ","Giô-na","Книга пророка Ионы":
				bible_short_name = "拿"
				switch chap_string {
					case "":
						print_string = "約拿書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("拿","約拿書", chap_string, "1","hakka")
							default:
								print_string = Bible_print_string("拿","約拿書", chap_string, sec_string,"hakka")
						}
				}
			case "Micah","彌","彌迦","彌迦書","Mic","mic","Книга пророка Михея","Mi-chê","ミカ書","ミカ","미가":
				bible_short_name = "彌"
				switch chap_string {
					case "":
						print_string = "彌迦書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("彌","彌迦書", chap_string, "1","hakka")
							default:
								print_string = Bible_print_string("彌","彌迦書", chap_string, sec_string,"hakka")
						}
				}
			case "Nah","Nahum","鴻","那鴻","那鴻書","Na","na","Книга пророка Наума","Na-hum","ナホム書","ナホム","나훔":
				bible_short_name = "鴻"
				switch chap_string {
					case "":
						print_string = "那鴻書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("鴻","那鴻書", chap_string, "1","hakka")
							default:
								print_string = Bible_print_string("鴻","那鴻書", chap_string, sec_string,"hakka")
						}
				}
			case "Habakkuk","哈","哈巴","哈巴谷","哈巴谷書","Hab","hab","Книга пророка Аввакума","Ha-ba-cúc","ハバクク書","ハバクク","ハバ","クク","ハバ書","하박국":
				bible_short_name = "哈"
				switch chap_string {
					case "":
						print_string = "哈巴谷書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("哈","哈巴谷書", chap_string, "1","hakka")
							default:
								print_string = Bible_print_string("哈","哈巴谷書", chap_string, sec_string,"hakka")
						}
				}
			case "Zeph","Zephaniah","番","西番雅","西番雅書","Zep","zep","스바냐","ゼパニヤ書","ゼパニヤ","Sô-phô-ni","Книга пророка Софонии":
				bible_short_name = "番"
				switch chap_string {
					case "":
						print_string = "西番雅書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("番","西番雅書", chap_string, "1","hakka")
							default:
								print_string = Bible_print_string("番","西番雅書", chap_string, sec_string,"hakka")
						}
				}
			case "Haggai","該","哈該","哈該書","Hag","hag","학개","ハガイ書","ハガイ","A-ghê","Книга пророка Аггея":
				bible_short_name = "該"
				switch chap_string {
					case "":
						print_string = "哈該書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("該","哈該書", chap_string, "1","hakka")
							default:
								print_string = Bible_print_string("該","哈該書", chap_string, sec_string,"hakka")
						}
				}
			case "Zech","Zechariah","亞","撒迦利亞","撒迦利亞書","Zec","zec","Книга пророка Захарии","Xa-cha-ri","스가랴","ゼカリヤ書","ゼカリヤ":
				bible_short_name = "亞"
				switch chap_string {
					case "":
						print_string = "撒迦利亞書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("亞","撒迦利亞書", chap_string, "1","hakka")
							default:
								print_string = Bible_print_string("亞","撒迦利亞書", chap_string, sec_string,"hakka")
						}
				}
			case "Malachi","瑪","","瑪拉","瑪拉基","瑪拉基書","Mal","mal","말라기","マラキ書","マラキ","Ma-la-chi","Книга пророка Малахии":
				bible_short_name = "瑪"
				switch chap_string {
					case "":
						print_string = "瑪拉基書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("瑪","瑪拉基書", chap_string, "1","hakka")
							default:
								print_string = Bible_print_string("瑪","瑪拉基書", chap_string, sec_string,"hakka")
						}
				}
			case "Matt","Matthew","太","馬太","馬太福音","Mt","mt","마태복음","マタイによる福音書","マタイ","マタイによる","Ma-thi-ơ","От Матфея святое благовествование":
				bible_short_name = "太"
				switch chap_string {
					case "":
						print_string = "馬太福音"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("太","馬太福音", chap_string, "1","hakka")
							default:
								print_string = Bible_print_string("太","馬太福音", chap_string, sec_string,"hakka")
						}
				}
			case "Mark","可","馬可","馬可福音","Mr","mr","マルコによる福音書","マルコ","マルコによる","마가복음","Mác","От Марка святое благовествование":
				bible_short_name = "可"
				switch chap_string {
					case "":
						print_string = "馬可福音"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("可","馬可福音", chap_string, "1","hakka")
							default:
								print_string = Bible_print_string("可","馬可福音", chap_string, sec_string,"hakka")
						}
				}
			case "Luke","路","路加","路加福音","Lu","lu","От Луки святое благовествование","Lu-ca","ルカによる福音書","ルカ","ルカによる","누가복음":
				bible_short_name = "路"
				switch chap_string {
					case "":
						print_string = "路加福音"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("路","路加福音", chap_string, "1","hakka")
							default:
								print_string = Bible_print_string("路","路加福音", chap_string, sec_string,"hakka")
						}
				}
			case "John","約","約翰","約翰福音","Joh","joh","От Иоанна святое благовествование","Giăng","ヨハネによる福音書","ヨハネ","ヨハネによる","요한복음":
				bible_short_name = "約"
				switch chap_string {
					case "":
						print_string = "約翰福音"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("約","約翰福音", chap_string, "1","hakka")
							default:
								print_string = Bible_print_string("約","約翰福音", chap_string, sec_string,"hakka")
						}
				}
			case "Acts","徒","使徒","使徒行傳","Ac","ac","Деяния святых апостолов","Công-vụ Các Sứ-đồ","使徒行伝","사도행전":
				bible_short_name = "徒"
				switch chap_string {
					case "":
						print_string = "使徒行傳"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("徒","使徒行傳", chap_string, "1","hakka")
							default:
								print_string = Bible_print_string("徒","使徒行傳", chap_string, sec_string,"hakka")
						}
				}
			case "Rom","Romans","羅","羅馬","羅馬書","Ro","ro","Послание к Римлянам","Rô-ma","ローマ","ローマ人への手紙","로마서":
				bible_short_name = "羅"
				switch chap_string {
					case "":
						print_string = "羅馬書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("羅","羅馬書", chap_string, "1","hakka")
							default:
								print_string = Bible_print_string("羅","羅馬書", chap_string, sec_string,"hakka")
						}
				}
			case "1 Cor","First Corinthians","林前","哥林多前","哥林多前書","1Co","1co","Первое послание к Коринфянам","1 Cô-rinh-tô","コリント人への第一の手紙","コリント一","コリント人への第一","고린도전서":
				bible_short_name = "林前"
				switch chap_string {
					case "":
						print_string = "哥林多前書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("林前","哥林多前書", chap_string, "1","hakka")
							default:
								print_string = Bible_print_string("林前","哥林多前書", chap_string, sec_string,"hakka")
						}
				}
			case "2 Cor","Second Corinthians","林後","哥林多後","哥林多後書","2Co","2co","Второе послание к Коринфянам","2 Cô-rinh-tô","コリント人への第二の手紙","コリント二","コリント人への第二の","고린도후서":
				bible_short_name = "林後"
				switch chap_string {
					case "":
						print_string = "哥林多後書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("林後","哥林多後書", chap_string, "1","hakka")
							default:
								print_string = Bible_print_string("林後","哥林多後書", chap_string, sec_string,"hakka")
						}
				}
			case "Gal","Galatians","加","加拉太","加拉太書","Ga","ga","Послание к Галатам","Ga-la-ti","ガラテヤ","ガラテヤ人への手紙","갈라디아서":
				bible_short_name = "加"
				switch chap_string {
					case "":
						print_string = "加拉太書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("加","加拉太書", chap_string, "1","hakka")
							default:
								print_string = Bible_print_string("加","加拉太書", chap_string, sec_string,"hakka")
						}
				}
			case "Ephesians","弗","以弗所","以弗所書","Eph","eph","Послание к Ефесянам","Ê-phê-sô","エペソ人への手紙","エペソ","エペソ人","エペソ人の手紙","에베소서":
				bible_short_name = "弗"
				switch chap_string {
					case "":
						print_string = "以弗所書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("弗","以弗所書", chap_string, "1","hakka")
							default:
								print_string = Bible_print_string("弗","以弗所書", chap_string, sec_string,"hakka")
						}
				}
			case "Phil","Philippians","腓","腓立","腓立比","腓立比書","Php","php","빌립보서","ピリピ","ピリピ人.ピリピ人への手紙","Послание к Филиппийцам","Phi-líp":
				bible_short_name = "腓"
				switch chap_string {
					case "":
						print_string = "腓立比書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("腓","腓立比書", chap_string, "1","hakka")
							default:
								print_string = Bible_print_string("腓","腓立比書", chap_string, sec_string,"hakka")
						}
				}
			case "Col","col","Colossians","西","歌羅西","歌羅","歌羅西書","Послание к Колоссянам","Cô-lô-se","コロサイ人への手紙","コロサイ","コロ","골로새서":
				bible_short_name = "西"
				switch chap_string {
					case "":
						print_string = "歌羅西書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("西","歌羅西書", chap_string, "1","hakka")
							default:
								print_string = Bible_print_string("西","歌羅西書", chap_string, sec_string,"hakka")
						}
				}
			case "1 Thess","First Thessalonians","帖前","帖撒羅尼迦前","帖撒羅尼迦前書","1Th","1th","데살로니가전서","テサロニケ人への第一の手紙","テサ一","テサロニケ一","1 Tê-sa-lô-ni-ca","Первое послание к Фессалоникийцам (Солунянам)":
				bible_short_name = "帖前"
				switch chap_string {
					case "":
						print_string = "帖撒羅尼迦前書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("帖前","帖撒羅尼迦前書", chap_string, "1","hakka")
							default:
								print_string = Bible_print_string("帖前","帖撒羅尼迦前書", chap_string, sec_string,"hakka")
						}
				}
			case "2 Thess","Second Thessalonians","帖後","帖撒羅尼迦後","帖撒羅尼迦後書","2Th","2th","데살로니가후서","テサロニケ人への第二の手紙","テサ二","テサロニケ二","2 Tê-sa-lô-ni-ca","Второе послание к Фессалоникийцам (Солунянам)":
				bible_short_name = "帖後"
				switch chap_string {
					case "":
						print_string = "帖撒羅尼迦後書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("帖後","帖撒羅尼迦後書", chap_string, "1","hakka")
							default:
								print_string = Bible_print_string("帖後","帖撒羅尼迦後書", chap_string, sec_string,"hakka")
						}
				}
			case "1 Tim","First Timothy","提前","提摩太前","提摩太前書","1Ti","1ti","Первое послание к Тимофею","1 Ti-mô-thê","テモテヘの第一の手紙","テモテ一","디모데전서":
				bible_short_name = "提前"
				switch chap_string {
					case "":
						print_string = "提摩太前書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("提前","提摩太前書", chap_string, "1","hakka")
							default:
								print_string = Bible_print_string("提前","提摩太前書", chap_string, sec_string,"hakka")
						}
				}
			case "2 Tim","Second Timothy","提後","提摩太後","提摩太後書","2Ti","2ti","Второе послание к Тимофею","2 Ti-mô-thê","テモテヘの第二の手紙","テモテ二","디모데후서":
				bible_short_name = "提後"
				switch chap_string {
					case "":
						print_string = "提摩太後書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("提後","提摩太後書", chap_string, "1","hakka")
							default:
								print_string = Bible_print_string("提後","提摩太後書", chap_string, sec_string,"hakka")
						}
				}
			case "Titus","多","提多","提多書","Tit","tit","Послание к Титу","Tít","テトスヘの手紙","テトス","디도서":
				bible_short_name = "多"
				switch chap_string {
					case "":
						print_string = "提多書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("多","提多書", chap_string, "1","hakka")
							default:
								print_string = Bible_print_string("多","提多書", chap_string, sec_string,"hakka")
						}
				}
			case "Philem","Philemon","門","腓利","腓利門","腓利門書","Phm","phm","Послание к Филимону","Phi-lê-môn","ピレモンヘの手紙","ピレモン","빌레몬서":
				bible_short_name = "門"
				switch chap_string {
					case "":
						print_string = "腓利門書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("門","腓利門書", chap_string, "1","hakka")
							default:
								print_string = Bible_print_string("門","腓利門書", chap_string, sec_string,"hakka")
						}
				}
			case "Heb","Hebrews","來","希伯來","希伯來書","heb","Послание к Евреям","Hê-bơ-rơ","ヘブル人への手紙","ヘブル","히브리서":
				bible_short_name = "來"
				switch chap_string {
					case "":
						print_string = "希伯來書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("來","希伯來書", chap_string, "1","hakka")
							default:
								print_string = Bible_print_string("來","希伯來書", chap_string, sec_string,"hakka")
						}
				}
			case "James","雅","雅各","雅各書","Jas","jas","Послание Иакова","Gia-cơ","ヤコブの手紙","ヤコブ","야고보서":
				bible_short_name = "雅"
				switch chap_string {
					case "":
						print_string = "雅各書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("雅","雅各書", chap_string, "1","hakka")
							default:
								print_string = Bible_print_string("雅","雅各書", chap_string, sec_string,"hakka")
						}
				}
			case "1 Pet","First Peter","彼前","彼得前","彼得前書","1Pe","1pe","Первое послание Петра","1 Phi-e-rơ","ペテロの第一の手紙","ペテロ一","베드로전서":
				bible_short_name = "彼前"
				switch chap_string {
					case "":
						print_string = "彼得前書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("彼前","彼得前書", chap_string, "1","hakka")
							default:
								print_string = Bible_print_string("彼前","彼得前書", chap_string, sec_string,"hakka")
						}
				}
			case "2 Pet","Second Peter","彼後","彼得後","彼得後書","2Pe","2pe","Второе послание Петра","2 Phi-e-rơ","ペテロの第二の手紙","ペテロ","베드로후서":
				bible_short_name = "彼後"
				switch chap_string {
					case "":
						print_string = "彼得後書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("彼後","彼得後書", chap_string, "1","hakka")
							default:
								print_string = Bible_print_string("彼後","彼得後書", chap_string, sec_string,"hakka")
						}
				}
			case "1 John","First John","約一","約翰一書","約翰1","約翰1書","1Jo","1jo","Первое послание Иоанна","1 Giăng","ヨハネの第一の手紙","ヨハネ一","요한일서":
				bible_short_name = "約一"
				switch chap_string {
					case "":
						print_string = "約翰一書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("約一","約翰一書", chap_string, "1","hakka")
							default:
								print_string = Bible_print_string("約一","約翰一書", chap_string, sec_string,"hakka")
						}
				}
			case "2 John","second John","約二","約翰二書","約翰2","約翰2書","2Jo","Второе послание Иоанна","2 Giăng","ヨハネの第二の手紙","ヨハネ二","요한2서":
				bible_short_name = "約二"
				switch chap_string {
					case "":
						print_string = "約翰二書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("約二","約翰二書", chap_string, "1","hakka")
							default:
								print_string = Bible_print_string("約二","約翰二書", chap_string, sec_string,"hakka")
						}
				}
			case "3 John","Third John","約三","約翰三書","約翰3","約翰3書","3Jo","3jo","Третье послание Иоанна","3 Giăng","ヨハネの第三の手紙","ヨハネ三","요한3서":
				bible_short_name = "約三"
				switch chap_string {
					case "":
						print_string = "約翰三書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("約三","約翰三書", chap_string, "1","hakka")
							default:
								print_string = Bible_print_string("約三","約翰三書", chap_string, sec_string,"hakka")
						}
				}
			case "Jude","猶","猶大","猶大書","jude","Послание Иуды","Giu-đe","ユダの手紙","ユダ","유다서":
				bible_short_name = "猶"
				switch chap_string {
					case "":
						print_string = "猶大書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("猶","猶大書", chap_string, "1","hakka")
							default:
								print_string = Bible_print_string("猶","猶大書", chap_string, sec_string,"hakka")
						}
				}
			case "Rev","Revelation","啟","啟示","啟示錄","Re","re","ｒｅ","Ｒｅ","rev","Откровение ап. Иоанна Богослова (Апокалипсис)","Khải-huyền","ヨハネの黙示録","黙示録","요한계시록":
				bible_short_name = "啟"
				switch chap_string {
					case "":
						print_string = "啟示錄"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("啟","啟示錄", chap_string, "1","hakka")
							default:
								print_string = Bible_print_string("啟","啟示錄", chap_string, sec_string,"hakka")
						}
				}
			default:
				print_string = "聖經"
				//print_string = "你是要找 " +  reg.ReplaceAllString(text, "$3") + " 對嗎？\n對不起，我還沒學呢...\n"
		}
	case "台語聖經馬雅各全羅":
		log.Print(reg.ReplaceAllString(text, "$3"))
		switch reg.ReplaceAllString(text, "$3") {

			case "Gen","Genesis","創","創世","創世紀","創世記","Ge","ge","gen","창세기","Sáng-thế Ký","Бытие":
				bible_short_name = "創"
				switch chap_string {
					case "":
						print_string = "創世記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("創","創世記", chap_string, "1","prebklcl")
							default:
								print_string = Bible_print_string("創","創世記", chap_string, sec_string,"prebklcl")
						}
				}
			case "ex","Ex","Exodus","埃及","出","出埃及","出埃及記","출애굽기","エジプト","出エジプト","出エジプト記","Xuất Ê-díp-tô Ký","Исход":
				bible_short_name = "出"
				switch chap_string {
					case "":
						print_string = "出埃及記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("出","出埃及記", chap_string, "1","prebklcl")
							default:
								print_string = Bible_print_string("出","出埃及記", chap_string, sec_string,"prebklcl")
						}
				}
			case "Lev","Leviticus","利","利未","利未記","Le","le","Левит","Lê-vi Ký","レビ記","レビ","레위기":
				bible_short_name = "利"
				switch chap_string {
					case "":
						print_string = "利未記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("利","利未記", chap_string, "1","prebklcl")
							default:
								print_string = Bible_print_string("利","利未記", chap_string, sec_string,"prebklcl")
						}
				}
			case "Num","Numbers","民","民數","民數記","Nu","nu","민수기","民数","民数記","Dân-số Ký","Числа":
				bible_short_name = "民"
				switch chap_string {
					case "":
						print_string = "民數記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("民","民數記", chap_string, "1","prebklcl")
							default:
								print_string = Bible_print_string("民","民數記", chap_string, sec_string,"prebklcl")
						}
				}
			case "Deut","Deuteronomy","申","申命","申命記","De","de","신명기","Phục-truyền Luật-lệ Ký","Второзаконие":
				bible_short_name = "申"
				switch chap_string {
					case "":
						print_string = "申命記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("申","申命記", chap_string, "1","prebklcl")
							default:
								print_string = Bible_print_string("申","申命記", chap_string, sec_string,"prebklcl")
						}
				}
			case "Josh","Joshua","約書亞","約書亞記","Jos","jos","여호수아","ヨシュア記","ヨシュア","Giô-suê","Книга Иисуса Навина":
				bible_short_name = "書"
				switch chap_string {
					case "":
						print_string = "約書亞記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("書","約書亞記", chap_string, "1","prebklcl")
							default:
								print_string = Bible_print_string("書","約書亞記", chap_string, sec_string,"prebklcl")
						}
				}
			case "Judg","Judges","士","士師","士師記","Jud","jud","jdg","Jdg","Книга Судей израилевых","Các Quan Xét","사사기":
				bible_short_name = "士"
				switch chap_string {
					case "":
						print_string = "士師記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("士","士師記", chap_string, "1","prebklcl")
							default:
								print_string = Bible_print_string("士","士師記", chap_string, sec_string,"prebklcl")
						}
				}
			case "Ruth","路得","路得記","Ru","ru","Rut","rut","룻기","ルツ","ルツ記","Ru-tơ","Книга Руфи":
				bible_short_name = "得"
				switch chap_string {
					case "":
						print_string = "路得記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("得","路得記", chap_string, "1","prebklcl")
							default:
								print_string = Bible_print_string("得","路得記", chap_string, sec_string,"prebklcl")
						}
				}
			case "1 Sam","First Samuel","撒上","撒母耳記上","1Sa","1sa","サムエル記上","サムエル上","サム上","사무엘상","1 Sa-mu-ên","Первая книга Царств":
				bible_short_name = "撒上"
				switch chap_string {
					case "":
						print_string = "撒母耳記上"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("撒上","撒母耳記上", chap_string, "1","prebklcl")
							default:
								print_string = Bible_print_string("撒上","撒母耳記上", chap_string, sec_string,"prebklcl")
						}
				}
			case "2 Sam","Second Samuel","撒下","撒母耳記下","2Sa","2sa","사무엘하","サムエル記下","サムエル下","サム下","2 Sa-mu-ên","Вторая книга Царств":
				bible_short_name = "撒下"
				switch chap_string {
					case "":
						print_string = "撒母耳記下"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("撒下","撒母耳記下", chap_string, "1","prebklcl")
							default:
								print_string = Bible_print_string("撒下","撒母耳記下", chap_string, sec_string,"prebklcl")
						}
				}
			case "1 Kin","First Kings","王上","列王上","列王紀上","列王記上","1Ki","1ki","열왕기상","Третья книга Царств","1 Các Vua":
				bible_short_name = "王上"
				switch chap_string {
					case "":
						print_string = "列王紀上"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("王上","列王紀上", chap_string, "1","prebklcl")
							default:
								print_string = Bible_print_string("王上","列王紀上", chap_string, sec_string,"prebklcl")
						}
				}
			case "2 Kin","Second Kings","王下","列王下","列王記下","列王紀下","2Ki","2ki","열왕기하","2 Các Vua","Четвертая книга Царств":
				bible_short_name = "王下"
				switch chap_string {
					case "":
						print_string = "列王紀下"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("王下","列王紀下", chap_string, "1","prebklcl")
							default:
								print_string = Bible_print_string("王下","列王紀下", chap_string, sec_string,"prebklcl")
						}
				}
			case "1 Chr","First Chronicles","歷上","代上","歷代志上","歷代上","1Ch","1ch","Первая книга Паралипоменон","1 Sử-ký","歴上","歴代上","歴代志上","역대상":
				bible_short_name = "代上"
				switch chap_string {
					case "":
						print_string = "歷代志上"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("代上","歷代志上", chap_string, "1","prebklcl")
							default:
								print_string = Bible_print_string("代上","歷代志上", chap_string, sec_string,"prebklcl")
						}
				}
			case "2 Chr","Second Chronicles","代下","歷下","歷代下","歷代志下","2Ch","2ch","역대하","歴代志下","歴代下","歴下","2 Sử-ký","Вторая книга Паралипоменон":
				bible_short_name = "代下"
				switch chap_string {
					case "":
						print_string = "歷代志下"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("代下","歷代志下", chap_string, "1","prebklcl")
							default:
								print_string = Bible_print_string("代下","歷代志下", chap_string, sec_string,"prebklcl")
						}
				}
			case "Ezra","拉","以斯拉","以斯拉記","Ezr","ezr","Первая книга Ездры","Ê-xơ-ra","エズラ","エズラ記","에스라":
				bible_short_name = "拉"
				switch chap_string {
					case "":
						print_string = "以斯拉記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("拉","以斯拉記", chap_string, "1","prebklcl")
							default:
								print_string = Bible_print_string("拉","以斯拉記", chap_string, sec_string,"prebklcl")
						}
				}
			case "Neh","Nehemiah","尼","尼希米","尼希米記","Ne","ne","느헤미야","ネヘミヤ書","ネヘミヤ","Nê-hê-mi","Книга Неемии":
				bible_short_name = "尼"
				switch chap_string {
					case "":
						print_string = "尼希米記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("尼","尼希米記", chap_string, "1","prebklcl")
							default:
								print_string = Bible_print_string("尼","尼希米記", chap_string, sec_string,"prebklcl")
						}
				}
			case "Esth","Esther","斯","以斯帖","以斯帖記","Es","est","Есфирь","Ê-xơ-tê","エステル","エステル記","에스더":
				bible_short_name = "斯"
				switch chap_string {
					case "":
						print_string = "以斯帖記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("斯","以斯帖記", chap_string, "1","prebklcl")
							default:
								print_string = Bible_print_string("斯","以斯帖記", chap_string, sec_string,"prebklcl")
						}
				}
			case "Job","job","伯","約伯","約伯記","Книга Иова","Gióp","ヨブ","ヨブ記","욥기":
				bible_short_name = "伯"
				switch chap_string {
					case "":
						print_string = "約伯記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("伯","約伯記", chap_string, "1","prebklcl")
							default:
								print_string = Bible_print_string("伯","約伯記", chap_string, sec_string,"prebklcl")
						}
				}
			case "Ps","Psalms","詩","詩篇","ps","시편","Thi-thiên","Псалтирь":
				bible_short_name = "詩"
				switch chap_string {
					case "":
						print_string = "詩篇"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("詩","詩篇", chap_string, "1","prebklcl")
							default:
								print_string = Bible_print_string("詩","詩篇", chap_string, sec_string,"prebklcl")
						}
				}
			case "Prov","Proverbs","箴","箴言","Pr","pr","Притчи Соломона","Châm-ngôn","잠언":
				bible_short_name = "箴"
				switch chap_string {
					case "":
						print_string = "箴言"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("箴","箴言", chap_string, "1","prebklcl")
							default:
								print_string = Bible_print_string("箴","箴言", chap_string, sec_string,"prebklcl")
						}
				}
			case "Eccl","Ecclesiastes","傳","傳道","傳道書","Ec","ec","Книга Екклезиаста","или Проповедника","Truyền-đạo","伝道の書","伝道","伝","伝道書","전도서":
				bible_short_name = "傳"
				switch chap_string {
					case "":
						print_string = "傳道書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("傳","傳道書", chap_string, "1","prebklcl")
							default:
								print_string = Bible_print_string("傳","傳道書", chap_string, sec_string,"prebklcl")
						}
				}
			case "Song","Song of Solomon","歌","雅歌","So","so","sng","Sng","Песнь песней Соломона","Nhã-ca","아가":
				bible_short_name = "歌"
				switch chap_string {
					case "":
						print_string = "雅歌"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("歌","雅歌", chap_string, "1","prebklcl")
							default:
								print_string = Bible_print_string("歌","雅歌", chap_string, sec_string,"prebklcl")
						}
				}
			case "Is","Isaiah","賽","以賽","以賽亞","以賽亞書","Isa","isa","Книга пророка Исаии","Ê-sai","イザヤ書","イザヤ","이사야":
				bible_short_name = "賽"
				switch chap_string {
					case "":
						print_string = "以賽亞書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("賽","以賽亞書", chap_string, "1","prebklcl")
							default:
								print_string = Bible_print_string("賽","以賽亞書", chap_string, sec_string,"prebklcl")
						}
				}
			case "Jer","Jeremiah","耶","耶利米","耶利米書","jer","예레미야","エレミヤ","エレミヤ書","Giê-rê-mi","Книга пророка Иеремии":
				bible_short_name = "耶"
				switch chap_string {
					case "":
						print_string = "耶利米書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("耶","耶利米書", chap_string, "1","prebklcl")
							default:
								print_string = Bible_print_string("耶","耶利米書", chap_string, sec_string,"prebklcl")
						}
				}
			case "Lam","Lamentations","哀","哀歌","耶利米哀歌","La","lam","예레미야애가","Ca-thương","Плач Иеремии":
				bible_short_name = "哀"
				switch chap_string {
					case "":
						print_string = "耶利米哀歌"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("哀","耶利米哀歌", chap_string, "1","prebklcl")
							default:
								print_string = Bible_print_string("哀","耶利米哀歌", chap_string, sec_string,"prebklcl")
						}
				}
			case "Ezek","Ezekiel","結","以西結","以西結書","Eze","eze","에스겔","エゼキエル書","エゼキエル","Ê-xê-chi-ên","Книга пророка Иезекииля":
				bible_short_name = "結"
				switch chap_string {
					case "":
						print_string = "以西結書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("結","以西結書", chap_string, "1","prebklcl")
							default:
								print_string = Bible_print_string("結","以西結書", chap_string, sec_string,"prebklcl")
						}
				}
			case "Dan","Daniel","但","但以理","但以理書","Da","da","Книга пророка Даниила","Đa-ni-ên","ダニエル書","ダニエル","다니엘":
				bible_short_name = "但"
				switch chap_string {
					case "":
						print_string = "但以理書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("但","但以理書", chap_string, "1","prebklcl")
							default:
								print_string = Bible_print_string("但","但以理書", chap_string, sec_string,"prebklcl")
						}
				}
			case "Hos","Hosea","何","何西","何西阿","何西阿書","Ho","ho","Книга пророка Осии","Ô-sê","ホセア書","ホセア","호세아":
				bible_short_name = "何"
				switch chap_string {
					case "":
						print_string = "何西阿書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("何","何西阿書", chap_string, "1","prebklcl")
							default:
								print_string = Bible_print_string("何","何西阿書", chap_string, sec_string,"prebklcl")
						}
				}
			case "Joel","珥","約珥","約珥書","Joe","joe","Книга пророка Иоиля","Giô-ên","ヨエル書","ヨエル","요엘":
				bible_short_name = "珥"
				switch chap_string {
					case "":
						print_string = "約珥書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("珥","約珥書", chap_string, "1","prebklcl")
							default:
								print_string = Bible_print_string("珥","約珥書", chap_string, sec_string,"prebklcl")
						}
				}
			case "Amos","摩","阿摩司書","Am","am","Книга пророка Амоса","A-mốt","アモス書","アモス","아모스":
				bible_short_name = "摩"
				switch chap_string {
					case "":
						print_string = "阿摩司書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("摩","阿摩司書", chap_string, "1","prebklcl")
							default:
								print_string = Bible_print_string("摩","阿摩司書", chap_string, sec_string,"prebklcl")
						}
				}
			case "Obad","Obadiah","俄","俄巴底亞","俄巴底亞書","Ob","ob","오바댜","オバデヤ書","オバデヤ","Áp-đia","Книга пророка Авдия":
				bible_short_name = "俄"
				switch chap_string {
					case "":
						print_string = "俄巴底亞書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("俄","俄巴底亞書", chap_string, "1","prebklcl")
							default:
								print_string = Bible_print_string("俄","俄巴底亞書", chap_string, sec_string,"prebklcl")
						}
				}
			case "Jon","Jonah","拿","約拿","約拿書","jon","요나","ヨナ書","ヨナ","Giô-na","Книга пророка Ионы":
				bible_short_name = "拿"
				switch chap_string {
					case "":
						print_string = "約拿書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("拿","約拿書", chap_string, "1","prebklcl")
							default:
								print_string = Bible_print_string("拿","約拿書", chap_string, sec_string,"prebklcl")
						}
				}
			case "Micah","彌","彌迦","彌迦書","Mic","mic","Книга пророка Михея","Mi-chê","ミカ書","ミカ","미가":
				bible_short_name = "彌"
				switch chap_string {
					case "":
						print_string = "彌迦書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("彌","彌迦書", chap_string, "1","prebklcl")
							default:
								print_string = Bible_print_string("彌","彌迦書", chap_string, sec_string,"prebklcl")
						}
				}
			case "Nah","Nahum","鴻","那鴻","那鴻書","Na","na","Книга пророка Наума","Na-hum","ナホム書","ナホム","나훔":
				bible_short_name = "鴻"
				switch chap_string {
					case "":
						print_string = "那鴻書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("鴻","那鴻書", chap_string, "1","prebklcl")
							default:
								print_string = Bible_print_string("鴻","那鴻書", chap_string, sec_string,"prebklcl")
						}
				}
			case "Habakkuk","哈","哈巴","哈巴谷","哈巴谷書","Hab","hab","Книга пророка Аввакума","Ha-ba-cúc","ハバクク書","ハバクク","ハバ","クク","ハバ書","하박국":
				bible_short_name = "哈"
				switch chap_string {
					case "":
						print_string = "哈巴谷書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("哈","哈巴谷書", chap_string, "1","prebklcl")
							default:
								print_string = Bible_print_string("哈","哈巴谷書", chap_string, sec_string,"prebklcl")
						}
				}
			case "Zeph","Zephaniah","番","西番雅","西番雅書","Zep","zep","스바냐","ゼパニヤ書","ゼパニヤ","Sô-phô-ni","Книга пророка Софонии":
				bible_short_name = "番"
				switch chap_string {
					case "":
						print_string = "西番雅書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("番","西番雅書", chap_string, "1","prebklcl")
							default:
								print_string = Bible_print_string("番","西番雅書", chap_string, sec_string,"prebklcl")
						}
				}
			case "Haggai","該","哈該","哈該書","Hag","hag","학개","ハガイ書","ハガイ","A-ghê","Книга пророка Аггея":
				bible_short_name = "該"
				switch chap_string {
					case "":
						print_string = "哈該書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("該","哈該書", chap_string, "1","prebklcl")
							default:
								print_string = Bible_print_string("該","哈該書", chap_string, sec_string,"prebklcl")
						}
				}
			case "Zech","Zechariah","亞","撒迦利亞","撒迦利亞書","Zec","zec","Книга пророка Захарии","Xa-cha-ri","스가랴","ゼカリヤ書","ゼカリヤ":
				bible_short_name = "亞"
				switch chap_string {
					case "":
						print_string = "撒迦利亞書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("亞","撒迦利亞書", chap_string, "1","prebklcl")
							default:
								print_string = Bible_print_string("亞","撒迦利亞書", chap_string, sec_string,"prebklcl")
						}
				}
			case "Malachi","瑪","","瑪拉","瑪拉基","瑪拉基書","Mal","mal","말라기","マラキ書","マラキ","Ma-la-chi","Книга пророка Малахии":
				bible_short_name = "瑪"
				switch chap_string {
					case "":
						print_string = "瑪拉基書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("瑪","瑪拉基書", chap_string, "1","prebklcl")
							default:
								print_string = Bible_print_string("瑪","瑪拉基書", chap_string, sec_string,"prebklcl")
						}
				}
			case "Matt","Matthew","太","馬太","馬太福音","Mt","mt","마태복음","マタイによる福音書","マタイ","マタイによる","Ma-thi-ơ","От Матфея святое благовествование":
				bible_short_name = "太"
				switch chap_string {
					case "":
						print_string = "馬太福音"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("太","馬太福音", chap_string, "1","prebklcl")
							default:
								print_string = Bible_print_string("太","馬太福音", chap_string, sec_string,"prebklcl")
						}
				}
			case "Mark","可","馬可","馬可福音","Mr","mr","マルコによる福音書","マルコ","マルコによる","마가복음","Mác","От Марка святое благовествование":
				bible_short_name = "可"
				switch chap_string {
					case "":
						print_string = "馬可福音"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("可","馬可福音", chap_string, "1","prebklcl")
							default:
								print_string = Bible_print_string("可","馬可福音", chap_string, sec_string,"prebklcl")
						}
				}
			case "Luke","路","路加","路加福音","Lu","lu","От Луки святое благовествование","Lu-ca","ルカによる福音書","ルカ","ルカによる","누가복음":
				bible_short_name = "路"
				switch chap_string {
					case "":
						print_string = "路加福音"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("路","路加福音", chap_string, "1","prebklcl")
							default:
								print_string = Bible_print_string("路","路加福音", chap_string, sec_string,"prebklcl")
						}
				}
			case "John","約","約翰","約翰福音","Joh","joh","От Иоанна святое благовествование","Giăng","ヨハネによる福音書","ヨハネ","ヨハネによる","요한복음":
				bible_short_name = "約"
				switch chap_string {
					case "":
						print_string = "約翰福音"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("約","約翰福音", chap_string, "1","prebklcl")
							default:
								print_string = Bible_print_string("約","約翰福音", chap_string, sec_string,"prebklcl")
						}
				}
			case "Acts","徒","使徒","使徒行傳","Ac","ac","Деяния святых апостолов","Công-vụ Các Sứ-đồ","使徒行伝","사도행전":
				bible_short_name = "徒"
				switch chap_string {
					case "":
						print_string = "使徒行傳"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("徒","使徒行傳", chap_string, "1","prebklcl")
							default:
								print_string = Bible_print_string("徒","使徒行傳", chap_string, sec_string,"prebklcl")
						}
				}
			case "Rom","Romans","羅","羅馬","羅馬書","Ro","ro","Послание к Римлянам","Rô-ma","ローマ","ローマ人への手紙","로마서":
				bible_short_name = "羅"
				switch chap_string {
					case "":
						print_string = "羅馬書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("羅","羅馬書", chap_string, "1","prebklcl")
							default:
								print_string = Bible_print_string("羅","羅馬書", chap_string, sec_string,"prebklcl")
						}
				}
			case "1 Cor","First Corinthians","林前","哥林多前","哥林多前書","1Co","1co","Первое послание к Коринфянам","1 Cô-rinh-tô","コリント人への第一の手紙","コリント一","コリント人への第一","고린도전서":
				bible_short_name = "林前"
				switch chap_string {
					case "":
						print_string = "哥林多前書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("林前","哥林多前書", chap_string, "1","prebklcl")
							default:
								print_string = Bible_print_string("林前","哥林多前書", chap_string, sec_string,"prebklcl")
						}
				}
			case "2 Cor","Second Corinthians","林後","哥林多後","哥林多後書","2Co","2co","Второе послание к Коринфянам","2 Cô-rinh-tô","コリント人への第二の手紙","コリント二","コリント人への第二の","고린도후서":
				bible_short_name = "林後"
				switch chap_string {
					case "":
						print_string = "哥林多後書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("林後","哥林多後書", chap_string, "1","prebklcl")
							default:
								print_string = Bible_print_string("林後","哥林多後書", chap_string, sec_string,"prebklcl")
						}
				}
			case "Gal","Galatians","加","加拉太","加拉太書","Ga","ga","Послание к Галатам","Ga-la-ti","ガラテヤ","ガラテヤ人への手紙","갈라디아서":
				bible_short_name = "加"
				switch chap_string {
					case "":
						print_string = "加拉太書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("加","加拉太書", chap_string, "1","prebklcl")
							default:
								print_string = Bible_print_string("加","加拉太書", chap_string, sec_string,"prebklcl")
						}
				}
			case "Ephesians","弗","以弗所","以弗所書","Eph","eph","Послание к Ефесянам","Ê-phê-sô","エペソ人への手紙","エペソ","エペソ人","エペソ人の手紙","에베소서":
				bible_short_name = "弗"
				switch chap_string {
					case "":
						print_string = "以弗所書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("弗","以弗所書", chap_string, "1","prebklcl")
							default:
								print_string = Bible_print_string("弗","以弗所書", chap_string, sec_string,"prebklcl")
						}
				}
			case "Phil","Philippians","腓","腓立","腓立比","腓立比書","Php","php","빌립보서","ピリピ","ピリピ人.ピリピ人への手紙","Послание к Филиппийцам","Phi-líp":
				bible_short_name = "腓"
				switch chap_string {
					case "":
						print_string = "腓立比書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("腓","腓立比書", chap_string, "1","prebklcl")
							default:
								print_string = Bible_print_string("腓","腓立比書", chap_string, sec_string,"prebklcl")
						}
				}
			case "Col","col","Colossians","西","歌羅西","歌羅","歌羅西書","Послание к Колоссянам","Cô-lô-se","コロサイ人への手紙","コロサイ","コロ","골로새서":
				bible_short_name = "西"
				switch chap_string {
					case "":
						print_string = "歌羅西書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("西","歌羅西書", chap_string, "1","prebklcl")
							default:
								print_string = Bible_print_string("西","歌羅西書", chap_string, sec_string,"prebklcl")
						}
				}
			case "1 Thess","First Thessalonians","帖前","帖撒羅尼迦前","帖撒羅尼迦前書","1Th","1th","데살로니가전서","テサロニケ人への第一の手紙","テサ一","テサロニケ一","1 Tê-sa-lô-ni-ca","Первое послание к Фессалоникийцам (Солунянам)":
				bible_short_name = "帖前"
				switch chap_string {
					case "":
						print_string = "帖撒羅尼迦前書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("帖前","帖撒羅尼迦前書", chap_string, "1","prebklcl")
							default:
								print_string = Bible_print_string("帖前","帖撒羅尼迦前書", chap_string, sec_string,"prebklcl")
						}
				}
			case "2 Thess","Second Thessalonians","帖後","帖撒羅尼迦後","帖撒羅尼迦後書","2Th","2th","데살로니가후서","テサロニケ人への第二の手紙","テサ二","テサロニケ二","2 Tê-sa-lô-ni-ca","Второе послание к Фессалоникийцам (Солунянам)":
				bible_short_name = "帖後"
				switch chap_string {
					case "":
						print_string = "帖撒羅尼迦後書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("帖後","帖撒羅尼迦後書", chap_string, "1","prebklcl")
							default:
								print_string = Bible_print_string("帖後","帖撒羅尼迦後書", chap_string, sec_string,"prebklcl")
						}
				}
			case "1 Tim","First Timothy","提前","提摩太前","提摩太前書","1Ti","1ti","Первое послание к Тимофею","1 Ti-mô-thê","テモテヘの第一の手紙","テモテ一","디모데전서":
				bible_short_name = "提前"
				switch chap_string {
					case "":
						print_string = "提摩太前書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("提前","提摩太前書", chap_string, "1","prebklcl")
							default:
								print_string = Bible_print_string("提前","提摩太前書", chap_string, sec_string,"prebklcl")
						}
				}
			case "2 Tim","Second Timothy","提後","提摩太後","提摩太後書","2Ti","2ti","Второе послание к Тимофею","2 Ti-mô-thê","テモテヘの第二の手紙","テモテ二","디모데후서":
				bible_short_name = "提後"
				switch chap_string {
					case "":
						print_string = "提摩太後書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("提後","提摩太後書", chap_string, "1","prebklcl")
							default:
								print_string = Bible_print_string("提後","提摩太後書", chap_string, sec_string,"prebklcl")
						}
				}
			case "Titus","多","提多","提多書","Tit","tit","Послание к Титу","Tít","テトスヘの手紙","テトス","디도서":
				bible_short_name = "多"
				switch chap_string {
					case "":
						print_string = "提多書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("多","提多書", chap_string, "1","prebklcl")
							default:
								print_string = Bible_print_string("多","提多書", chap_string, sec_string,"prebklcl")
						}
				}
			case "Philem","Philemon","門","腓利","腓利門","腓利門書","Phm","phm","Послание к Филимону","Phi-lê-môn","ピレモンヘの手紙","ピレモン","빌레몬서":
				bible_short_name = "門"
				switch chap_string {
					case "":
						print_string = "腓利門書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("門","腓利門書", chap_string, "1","prebklcl")
							default:
								print_string = Bible_print_string("門","腓利門書", chap_string, sec_string,"prebklcl")
						}
				}
			case "Heb","Hebrews","來","希伯來","希伯來書","heb","Послание к Евреям","Hê-bơ-rơ","ヘブル人への手紙","ヘブル","히브리서":
				bible_short_name = "來"
				switch chap_string {
					case "":
						print_string = "希伯來書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("來","希伯來書", chap_string, "1","prebklcl")
							default:
								print_string = Bible_print_string("來","希伯來書", chap_string, sec_string,"prebklcl")
						}
				}
			case "James","雅","雅各","雅各書","Jas","jas","Послание Иакова","Gia-cơ","ヤコブの手紙","ヤコブ","야고보서":
				bible_short_name = "雅"
				switch chap_string {
					case "":
						print_string = "雅各書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("雅","雅各書", chap_string, "1","prebklcl")
							default:
								print_string = Bible_print_string("雅","雅各書", chap_string, sec_string,"prebklcl")
						}
				}
			case "1 Pet","First Peter","彼前","彼得前","彼得前書","1Pe","1pe","Первое послание Петра","1 Phi-e-rơ","ペテロの第一の手紙","ペテロ一","베드로전서":
				bible_short_name = "彼前"
				switch chap_string {
					case "":
						print_string = "彼得前書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("彼前","彼得前書", chap_string, "1","prebklcl")
							default:
								print_string = Bible_print_string("彼前","彼得前書", chap_string, sec_string,"prebklcl")
						}
				}
			case "2 Pet","Second Peter","彼後","彼得後","彼得後書","2Pe","2pe","Второе послание Петра","2 Phi-e-rơ","ペテロの第二の手紙","ペテロ","베드로후서":
				bible_short_name = "彼後"
				switch chap_string {
					case "":
						print_string = "彼得後書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("彼後","彼得後書", chap_string, "1","prebklcl")
							default:
								print_string = Bible_print_string("彼後","彼得後書", chap_string, sec_string,"prebklcl")
						}
				}
			case "1 John","First John","約一","約翰一書","約翰1","約翰1書","1Jo","1jo","Первое послание Иоанна","1 Giăng","ヨハネの第一の手紙","ヨハネ一","요한일서":
				bible_short_name = "約一"
				switch chap_string {
					case "":
						print_string = "約翰一書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("約一","約翰一書", chap_string, "1","prebklcl")
							default:
								print_string = Bible_print_string("約一","約翰一書", chap_string, sec_string,"prebklcl")
						}
				}
			case "2 John","second John","約二","約翰二書","約翰2","約翰2書","2Jo","Второе послание Иоанна","2 Giăng","ヨハネの第二の手紙","ヨハネ二","요한2서":
				bible_short_name = "約二"
				switch chap_string {
					case "":
						print_string = "約翰二書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("約二","約翰二書", chap_string, "1","prebklcl")
							default:
								print_string = Bible_print_string("約二","約翰二書", chap_string, sec_string,"prebklcl")
						}
				}
			case "3 John","Third John","約三","約翰三書","約翰3","約翰3書","3Jo","3jo","Третье послание Иоанна","3 Giăng","ヨハネの第三の手紙","ヨハネ三","요한3서":
				bible_short_name = "約三"
				switch chap_string {
					case "":
						print_string = "約翰三書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("約三","約翰三書", chap_string, "1","prebklcl")
							default:
								print_string = Bible_print_string("約三","約翰三書", chap_string, sec_string,"prebklcl")
						}
				}
			case "Jude","猶","猶大","猶大書","jude","Послание Иуды","Giu-đe","ユダの手紙","ユダ","유다서":
				bible_short_name = "猶"
				switch chap_string {
					case "":
						print_string = "猶大書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("猶","猶大書", chap_string, "1","prebklcl")
							default:
								print_string = Bible_print_string("猶","猶大書", chap_string, sec_string,"prebklcl")
						}
				}
			case "Rev","Revelation","啟","啟示","啟示錄","Re","re","ｒｅ","Ｒｅ","rev","Откровение ап. Иоанна Богослова (Апокалипсис)","Khải-huyền","ヨハネの黙示録","黙示録","요한계시록":
				bible_short_name = "啟"
				switch chap_string {
					case "":
						print_string = "啟示錄"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("啟","啟示錄", chap_string, "1","prebklcl")
							default:
								print_string = Bible_print_string("啟","啟示錄", chap_string, sec_string,"prebklcl")
						}
				}
			default:
				print_string = "聖經"
				//print_string = "你是要找 " +  reg.ReplaceAllString(text, "$3") + " 對嗎？\n對不起，我還沒學呢...\n"
		}
	case "台語聖經馬雅各漢羅":
		log.Print(reg.ReplaceAllString(text, "$3"))
		switch reg.ReplaceAllString(text, "$3") {

			case "Gen","Genesis","創","創世","創世紀","創世記","Ge","ge","gen","창세기","Sáng-thế Ký","Бытие":
				bible_short_name = "創"
				switch chap_string {
					case "":
						print_string = "創世記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("創","創世記", chap_string, "1","prebklhl")
							default:
								print_string = Bible_print_string("創","創世記", chap_string, sec_string,"prebklhl")
						}
				}
			case "ex","Ex","Exodus","埃及","出","出埃及","出埃及記","출애굽기","エジプト","出エジプト","出エジプト記","Xuất Ê-díp-tô Ký","Исход":
				bible_short_name = "出"
				switch chap_string {
					case "":
						print_string = "出埃及記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("出","出埃及記", chap_string, "1","prebklhl")
							default:
								print_string = Bible_print_string("出","出埃及記", chap_string, sec_string,"prebklhl")
						}
				}
			case "Lev","Leviticus","利","利未","利未記","Le","le","Левит","Lê-vi Ký","レビ記","レビ","레위기":
				bible_short_name = "利"
				switch chap_string {
					case "":
						print_string = "利未記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("利","利未記", chap_string, "1","prebklhl")
							default:
								print_string = Bible_print_string("利","利未記", chap_string, sec_string,"prebklhl")
						}
				}
			case "Num","Numbers","民","民數","民數記","Nu","nu","민수기","民数","民数記","Dân-số Ký","Числа":
				bible_short_name = "民"
				switch chap_string {
					case "":
						print_string = "民數記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("民","民數記", chap_string, "1","prebklhl")
							default:
								print_string = Bible_print_string("民","民數記", chap_string, sec_string,"prebklhl")
						}
				}
			case "Deut","Deuteronomy","申","申命","申命記","De","de","신명기","Phục-truyền Luật-lệ Ký","Второзаконие":
				bible_short_name = "申"
				switch chap_string {
					case "":
						print_string = "申命記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("申","申命記", chap_string, "1","prebklhl")
							default:
								print_string = Bible_print_string("申","申命記", chap_string, sec_string,"prebklhl")
						}
				}
			case "Josh","Joshua","約書亞","約書亞記","Jos","jos","여호수아","ヨシュア記","ヨシュア","Giô-suê","Книга Иисуса Навина":
				bible_short_name = "書"
				switch chap_string {
					case "":
						print_string = "約書亞記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("書","約書亞記", chap_string, "1","prebklhl")
							default:
								print_string = Bible_print_string("書","約書亞記", chap_string, sec_string,"prebklhl")
						}
				}
			case "Judg","Judges","士","士師","士師記","Jud","jud","jdg","Jdg","Книга Судей израилевых","Các Quan Xét","사사기":
				bible_short_name = "士"
				switch chap_string {
					case "":
						print_string = "士師記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("士","士師記", chap_string, "1","prebklhl")
							default:
								print_string = Bible_print_string("士","士師記", chap_string, sec_string,"prebklhl")
						}
				}
			case "Ruth","路得","路得記","Ru","ru","Rut","rut","룻기","ルツ","ルツ記","Ru-tơ","Книга Руфи":
				bible_short_name = "得"
				switch chap_string {
					case "":
						print_string = "路得記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("得","路得記", chap_string, "1","prebklhl")
							default:
								print_string = Bible_print_string("得","路得記", chap_string, sec_string,"prebklhl")
						}
				}
			case "1 Sam","First Samuel","撒上","撒母耳記上","1Sa","1sa","サムエル記上","サムエル上","サム上","사무엘상","1 Sa-mu-ên","Первая книга Царств":
				bible_short_name = "撒上"
				switch chap_string {
					case "":
						print_string = "撒母耳記上"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("撒上","撒母耳記上", chap_string, "1","prebklhl")
							default:
								print_string = Bible_print_string("撒上","撒母耳記上", chap_string, sec_string,"prebklhl")
						}
				}
			case "2 Sam","Second Samuel","撒下","撒母耳記下","2Sa","2sa","사무엘하","サムエル記下","サムエル下","サム下","2 Sa-mu-ên","Вторая книга Царств":
				bible_short_name = "撒下"
				switch chap_string {
					case "":
						print_string = "撒母耳記下"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("撒下","撒母耳記下", chap_string, "1","prebklhl")
							default:
								print_string = Bible_print_string("撒下","撒母耳記下", chap_string, sec_string,"prebklhl")
						}
				}
			case "1 Kin","First Kings","王上","列王上","列王紀上","列王記上","1Ki","1ki","열왕기상","Третья книга Царств","1 Các Vua":
				bible_short_name = "王上"
				switch chap_string {
					case "":
						print_string = "列王紀上"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("王上","列王紀上", chap_string, "1","prebklhl")
							default:
								print_string = Bible_print_string("王上","列王紀上", chap_string, sec_string,"prebklhl")
						}
				}
			case "2 Kin","Second Kings","王下","列王下","列王記下","列王紀下","2Ki","2ki","열왕기하","2 Các Vua","Четвертая книга Царств":
				bible_short_name = "王下"
				switch chap_string {
					case "":
						print_string = "列王紀下"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("王下","列王紀下", chap_string, "1","prebklhl")
							default:
								print_string = Bible_print_string("王下","列王紀下", chap_string, sec_string,"prebklhl")
						}
				}
			case "1 Chr","First Chronicles","歷上","代上","歷代志上","歷代上","1Ch","1ch","Первая книга Паралипоменон","1 Sử-ký","歴上","歴代上","歴代志上","역대상":
				bible_short_name = "代上"
				switch chap_string {
					case "":
						print_string = "歷代志上"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("代上","歷代志上", chap_string, "1","prebklhl")
							default:
								print_string = Bible_print_string("代上","歷代志上", chap_string, sec_string,"prebklhl")
						}
				}
			case "2 Chr","Second Chronicles","代下","歷下","歷代下","歷代志下","2Ch","2ch","역대하","歴代志下","歴代下","歴下","2 Sử-ký","Вторая книга Паралипоменон":
				bible_short_name = "代下"
				switch chap_string {
					case "":
						print_string = "歷代志下"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("代下","歷代志下", chap_string, "1","prebklhl")
							default:
								print_string = Bible_print_string("代下","歷代志下", chap_string, sec_string,"prebklhl")
						}
				}
			case "Ezra","拉","以斯拉","以斯拉記","Ezr","ezr","Первая книга Ездры","Ê-xơ-ra","エズラ","エズラ記","에스라":
				bible_short_name = "拉"
				switch chap_string {
					case "":
						print_string = "以斯拉記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("拉","以斯拉記", chap_string, "1","prebklhl")
							default:
								print_string = Bible_print_string("拉","以斯拉記", chap_string, sec_string,"prebklhl")
						}
				}
			case "Neh","Nehemiah","尼","尼希米","尼希米記","Ne","ne","느헤미야","ネヘミヤ書","ネヘミヤ","Nê-hê-mi","Книга Неемии":
				bible_short_name = "尼"
				switch chap_string {
					case "":
						print_string = "尼希米記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("尼","尼希米記", chap_string, "1","prebklhl")
							default:
								print_string = Bible_print_string("尼","尼希米記", chap_string, sec_string,"prebklhl")
						}
				}
			case "Esth","Esther","斯","以斯帖","以斯帖記","Es","est","Есфирь","Ê-xơ-tê","エステル","エステル記","에스더":
				bible_short_name = "斯"
				switch chap_string {
					case "":
						print_string = "以斯帖記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("斯","以斯帖記", chap_string, "1","prebklhl")
							default:
								print_string = Bible_print_string("斯","以斯帖記", chap_string, sec_string,"prebklhl")
						}
				}
			case "Job","job","伯","約伯","約伯記","Книга Иова","Gióp","ヨブ","ヨブ記","욥기":
				bible_short_name = "伯"
				switch chap_string {
					case "":
						print_string = "約伯記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("伯","約伯記", chap_string, "1","prebklhl")
							default:
								print_string = Bible_print_string("伯","約伯記", chap_string, sec_string,"prebklhl")
						}
				}
			case "Ps","Psalms","詩","詩篇","ps","시편","Thi-thiên","Псалтирь":
				bible_short_name = "詩"
				switch chap_string {
					case "":
						print_string = "詩篇"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("詩","詩篇", chap_string, "1","prebklhl")
							default:
								print_string = Bible_print_string("詩","詩篇", chap_string, sec_string,"prebklhl")
						}
				}
			case "Prov","Proverbs","箴","箴言","Pr","pr","Притчи Соломона","Châm-ngôn","잠언":
				bible_short_name = "箴"
				switch chap_string {
					case "":
						print_string = "箴言"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("箴","箴言", chap_string, "1","prebklhl")
							default:
								print_string = Bible_print_string("箴","箴言", chap_string, sec_string,"prebklhl")
						}
				}
			case "Eccl","Ecclesiastes","傳","傳道","傳道書","Ec","ec","Книга Екклезиаста","или Проповедника","Truyền-đạo","伝道の書","伝道","伝","伝道書","전도서":
				bible_short_name = "傳"
				switch chap_string {
					case "":
						print_string = "傳道書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("傳","傳道書", chap_string, "1","prebklhl")
							default:
								print_string = Bible_print_string("傳","傳道書", chap_string, sec_string,"prebklhl")
						}
				}
			case "Song","Song of Solomon","歌","雅歌","So","so","sng","Sng","Песнь песней Соломона","Nhã-ca","아가":
				bible_short_name = "歌"
				switch chap_string {
					case "":
						print_string = "雅歌"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("歌","雅歌", chap_string, "1","prebklhl")
							default:
								print_string = Bible_print_string("歌","雅歌", chap_string, sec_string,"prebklhl")
						}
				}
			case "Is","Isaiah","賽","以賽","以賽亞","以賽亞書","Isa","isa","Книга пророка Исаии","Ê-sai","イザヤ書","イザヤ","이사야":
				bible_short_name = "賽"
				switch chap_string {
					case "":
						print_string = "以賽亞書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("賽","以賽亞書", chap_string, "1","prebklhl")
							default:
								print_string = Bible_print_string("賽","以賽亞書", chap_string, sec_string,"prebklhl")
						}
				}
			case "Jer","Jeremiah","耶","耶利米","耶利米書","jer","예레미야","エレミヤ","エレミヤ書","Giê-rê-mi","Книга пророка Иеремии":
				bible_short_name = "耶"
				switch chap_string {
					case "":
						print_string = "耶利米書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("耶","耶利米書", chap_string, "1","prebklhl")
							default:
								print_string = Bible_print_string("耶","耶利米書", chap_string, sec_string,"prebklhl")
						}
				}
			case "Lam","Lamentations","哀","哀歌","耶利米哀歌","La","lam","예레미야애가","Ca-thương","Плач Иеремии":
				bible_short_name = "哀"
				switch chap_string {
					case "":
						print_string = "耶利米哀歌"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("哀","耶利米哀歌", chap_string, "1","prebklhl")
							default:
								print_string = Bible_print_string("哀","耶利米哀歌", chap_string, sec_string,"prebklhl")
						}
				}
			case "Ezek","Ezekiel","結","以西結","以西結書","Eze","eze","에스겔","エゼキエル書","エゼキエル","Ê-xê-chi-ên","Книга пророка Иезекииля":
				bible_short_name = "結"
				switch chap_string {
					case "":
						print_string = "以西結書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("結","以西結書", chap_string, "1","prebklhl")
							default:
								print_string = Bible_print_string("結","以西結書", chap_string, sec_string,"prebklhl")
						}
				}
			case "Dan","Daniel","但","但以理","但以理書","Da","da","Книга пророка Даниила","Đa-ni-ên","ダニエル書","ダニエル","다니엘":
				bible_short_name = "但"
				switch chap_string {
					case "":
						print_string = "但以理書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("但","但以理書", chap_string, "1","prebklhl")
							default:
								print_string = Bible_print_string("但","但以理書", chap_string, sec_string,"prebklhl")
						}
				}
			case "Hos","Hosea","何","何西","何西阿","何西阿書","Ho","ho","Книга пророка Осии","Ô-sê","ホセア書","ホセア","호세아":
				bible_short_name = "何"
				switch chap_string {
					case "":
						print_string = "何西阿書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("何","何西阿書", chap_string, "1","prebklhl")
							default:
								print_string = Bible_print_string("何","何西阿書", chap_string, sec_string,"prebklhl")
						}
				}
			case "Joel","珥","約珥","約珥書","Joe","joe","Книга пророка Иоиля","Giô-ên","ヨエル書","ヨエル","요엘":
				bible_short_name = "珥"
				switch chap_string {
					case "":
						print_string = "約珥書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("珥","約珥書", chap_string, "1","prebklhl")
							default:
								print_string = Bible_print_string("珥","約珥書", chap_string, sec_string,"prebklhl")
						}
				}
			case "Amos","摩","阿摩司書","Am","am","Книга пророка Амоса","A-mốt","アモス書","アモス","아모스":
				bible_short_name = "摩"
				switch chap_string {
					case "":
						print_string = "阿摩司書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("摩","阿摩司書", chap_string, "1","prebklhl")
							default:
								print_string = Bible_print_string("摩","阿摩司書", chap_string, sec_string,"prebklhl")
						}
				}
			case "Obad","Obadiah","俄","俄巴底亞","俄巴底亞書","Ob","ob","오바댜","オバデヤ書","オバデヤ","Áp-đia","Книга пророка Авдия":
				bible_short_name = "俄"
				switch chap_string {
					case "":
						print_string = "俄巴底亞書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("俄","俄巴底亞書", chap_string, "1","prebklhl")
							default:
								print_string = Bible_print_string("俄","俄巴底亞書", chap_string, sec_string,"prebklhl")
						}
				}
			case "Jon","Jonah","拿","約拿","約拿書","jon","요나","ヨナ書","ヨナ","Giô-na","Книга пророка Ионы":
				bible_short_name = "拿"
				switch chap_string {
					case "":
						print_string = "約拿書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("拿","約拿書", chap_string, "1","prebklhl")
							default:
								print_string = Bible_print_string("拿","約拿書", chap_string, sec_string,"prebklhl")
						}
				}
			case "Micah","彌","彌迦","彌迦書","Mic","mic","Книга пророка Михея","Mi-chê","ミカ書","ミカ","미가":
				bible_short_name = "彌"
				switch chap_string {
					case "":
						print_string = "彌迦書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("彌","彌迦書", chap_string, "1","prebklhl")
							default:
								print_string = Bible_print_string("彌","彌迦書", chap_string, sec_string,"prebklhl")
						}
				}
			case "Nah","Nahum","鴻","那鴻","那鴻書","Na","na","Книга пророка Наума","Na-hum","ナホム書","ナホム","나훔":
				bible_short_name = "鴻"
				switch chap_string {
					case "":
						print_string = "那鴻書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("鴻","那鴻書", chap_string, "1","prebklhl")
							default:
								print_string = Bible_print_string("鴻","那鴻書", chap_string, sec_string,"prebklhl")
						}
				}
			case "Habakkuk","哈","哈巴","哈巴谷","哈巴谷書","Hab","hab","Книга пророка Аввакума","Ha-ba-cúc","ハバクク書","ハバクク","ハバ","クク","ハバ書","하박국":
				bible_short_name = "哈"
				switch chap_string {
					case "":
						print_string = "哈巴谷書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("哈","哈巴谷書", chap_string, "1","prebklhl")
							default:
								print_string = Bible_print_string("哈","哈巴谷書", chap_string, sec_string,"prebklhl")
						}
				}
			case "Zeph","Zephaniah","番","西番雅","西番雅書","Zep","zep","스바냐","ゼパニヤ書","ゼパニヤ","Sô-phô-ni","Книга пророка Софонии":
				bible_short_name = "番"
				switch chap_string {
					case "":
						print_string = "西番雅書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("番","西番雅書", chap_string, "1","prebklhl")
							default:
								print_string = Bible_print_string("番","西番雅書", chap_string, sec_string,"prebklhl")
						}
				}
			case "Haggai","該","哈該","哈該書","Hag","hag","학개","ハガイ書","ハガイ","A-ghê","Книга пророка Аггея":
				bible_short_name = "該"
				switch chap_string {
					case "":
						print_string = "哈該書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("該","哈該書", chap_string, "1","prebklhl")
							default:
								print_string = Bible_print_string("該","哈該書", chap_string, sec_string,"prebklhl")
						}
				}
			case "Zech","Zechariah","亞","撒迦利亞","撒迦利亞書","Zec","zec","Книга пророка Захарии","Xa-cha-ri","스가랴","ゼカリヤ書","ゼカリヤ":
				bible_short_name = "亞"
				switch chap_string {
					case "":
						print_string = "撒迦利亞書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("亞","撒迦利亞書", chap_string, "1","prebklhl")
							default:
								print_string = Bible_print_string("亞","撒迦利亞書", chap_string, sec_string,"prebklhl")
						}
				}
			case "Malachi","瑪","","瑪拉","瑪拉基","瑪拉基書","Mal","mal","말라기","マラキ書","マラキ","Ma-la-chi","Книга пророка Малахии":
				bible_short_name = "瑪"
				switch chap_string {
					case "":
						print_string = "瑪拉基書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("瑪","瑪拉基書", chap_string, "1","prebklhl")
							default:
								print_string = Bible_print_string("瑪","瑪拉基書", chap_string, sec_string,"prebklhl")
						}
				}
			case "Matt","Matthew","太","馬太","馬太福音","Mt","mt","마태복음","マタイによる福音書","マタイ","マタイによる","Ma-thi-ơ","От Матфея святое благовествование":
				bible_short_name = "太"
				switch chap_string {
					case "":
						print_string = "馬太福音"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("太","馬太福音", chap_string, "1","prebklhl")
							default:
								print_string = Bible_print_string("太","馬太福音", chap_string, sec_string,"prebklhl")
						}
				}
			case "Mark","可","馬可","馬可福音","Mr","mr","マルコによる福音書","マルコ","マルコによる","마가복음","Mác","От Марка святое благовествование":
				bible_short_name = "可"
				switch chap_string {
					case "":
						print_string = "馬可福音"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("可","馬可福音", chap_string, "1","prebklhl")
							default:
								print_string = Bible_print_string("可","馬可福音", chap_string, sec_string,"prebklhl")
						}
				}
			case "Luke","路","路加","路加福音","Lu","lu","От Луки святое благовествование","Lu-ca","ルカによる福音書","ルカ","ルカによる","누가복음":
				bible_short_name = "路"
				switch chap_string {
					case "":
						print_string = "路加福音"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("路","路加福音", chap_string, "1","prebklhl")
							default:
								print_string = Bible_print_string("路","路加福音", chap_string, sec_string,"prebklhl")
						}
				}
			case "John","約","約翰","約翰福音","Joh","joh","От Иоанна святое благовествование","Giăng","ヨハネによる福音書","ヨハネ","ヨハネによる","요한복음":
				bible_short_name = "約"
				switch chap_string {
					case "":
						print_string = "約翰福音"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("約","約翰福音", chap_string, "1","prebklhl")
							default:
								print_string = Bible_print_string("約","約翰福音", chap_string, sec_string,"prebklhl")
						}
				}
			case "Acts","徒","使徒","使徒行傳","Ac","ac","Деяния святых апостолов","Công-vụ Các Sứ-đồ","使徒行伝","사도행전":
				bible_short_name = "徒"
				switch chap_string {
					case "":
						print_string = "使徒行傳"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("徒","使徒行傳", chap_string, "1","prebklhl")
							default:
								print_string = Bible_print_string("徒","使徒行傳", chap_string, sec_string,"prebklhl")
						}
				}
			case "Rom","Romans","羅","羅馬","羅馬書","Ro","ro","Послание к Римлянам","Rô-ma","ローマ","ローマ人への手紙","로마서":
				bible_short_name = "羅"
				switch chap_string {
					case "":
						print_string = "羅馬書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("羅","羅馬書", chap_string, "1","prebklhl")
							default:
								print_string = Bible_print_string("羅","羅馬書", chap_string, sec_string,"prebklhl")
						}
				}
			case "1 Cor","First Corinthians","林前","哥林多前","哥林多前書","1Co","1co","Первое послание к Коринфянам","1 Cô-rinh-tô","コリント人への第一の手紙","コリント一","コリント人への第一","고린도전서":
				bible_short_name = "林前"
				switch chap_string {
					case "":
						print_string = "哥林多前書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("林前","哥林多前書", chap_string, "1","prebklhl")
							default:
								print_string = Bible_print_string("林前","哥林多前書", chap_string, sec_string,"prebklhl")
						}
				}
			case "2 Cor","Second Corinthians","林後","哥林多後","哥林多後書","2Co","2co","Второе послание к Коринфянам","2 Cô-rinh-tô","コリント人への第二の手紙","コリント二","コリント人への第二の","고린도후서":
				bible_short_name = "林後"
				switch chap_string {
					case "":
						print_string = "哥林多後書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("林後","哥林多後書", chap_string, "1","prebklhl")
							default:
								print_string = Bible_print_string("林後","哥林多後書", chap_string, sec_string,"prebklhl")
						}
				}
			case "Gal","Galatians","加","加拉太","加拉太書","Ga","ga","Послание к Галатам","Ga-la-ti","ガラテヤ","ガラテヤ人への手紙","갈라디아서":
				bible_short_name = "加"
				switch chap_string {
					case "":
						print_string = "加拉太書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("加","加拉太書", chap_string, "1","prebklhl")
							default:
								print_string = Bible_print_string("加","加拉太書", chap_string, sec_string,"prebklhl")
						}
				}
			case "Ephesians","弗","以弗所","以弗所書","Eph","eph","Послание к Ефесянам","Ê-phê-sô","エペソ人への手紙","エペソ","エペソ人","エペソ人の手紙","에베소서":
				bible_short_name = "弗"
				switch chap_string {
					case "":
						print_string = "以弗所書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("弗","以弗所書", chap_string, "1","prebklhl")
							default:
								print_string = Bible_print_string("弗","以弗所書", chap_string, sec_string,"prebklhl")
						}
				}
			case "Phil","Philippians","腓","腓立","腓立比","腓立比書","Php","php","빌립보서","ピリピ","ピリピ人.ピリピ人への手紙","Послание к Филиппийцам","Phi-líp":
				bible_short_name = "腓"
				switch chap_string {
					case "":
						print_string = "腓立比書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("腓","腓立比書", chap_string, "1","prebklhl")
							default:
								print_string = Bible_print_string("腓","腓立比書", chap_string, sec_string,"prebklhl")
						}
				}
			case "Col","col","Colossians","西","歌羅西","歌羅","歌羅西書","Послание к Колоссянам","Cô-lô-se","コロサイ人への手紙","コロサイ","コロ","골로새서":
				bible_short_name = "西"
				switch chap_string {
					case "":
						print_string = "歌羅西書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("西","歌羅西書", chap_string, "1","prebklhl")
							default:
								print_string = Bible_print_string("西","歌羅西書", chap_string, sec_string,"prebklhl")
						}
				}
			case "1 Thess","First Thessalonians","帖前","帖撒羅尼迦前","帖撒羅尼迦前書","1Th","1th","데살로니가전서","テサロニケ人への第一の手紙","テサ一","テサロニケ一","1 Tê-sa-lô-ni-ca","Первое послание к Фессалоникийцам (Солунянам)":
				bible_short_name = "帖前"
				switch chap_string {
					case "":
						print_string = "帖撒羅尼迦前書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("帖前","帖撒羅尼迦前書", chap_string, "1","prebklhl")
							default:
								print_string = Bible_print_string("帖前","帖撒羅尼迦前書", chap_string, sec_string,"prebklhl")
						}
				}
			case "2 Thess","Second Thessalonians","帖後","帖撒羅尼迦後","帖撒羅尼迦後書","2Th","2th","데살로니가후서","テサロニケ人への第二の手紙","テサ二","テサロニケ二","2 Tê-sa-lô-ni-ca","Второе послание к Фессалоникийцам (Солунянам)":
				bible_short_name = "帖後"
				switch chap_string {
					case "":
						print_string = "帖撒羅尼迦後書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("帖後","帖撒羅尼迦後書", chap_string, "1","prebklhl")
							default:
								print_string = Bible_print_string("帖後","帖撒羅尼迦後書", chap_string, sec_string,"prebklhl")
						}
				}
			case "1 Tim","First Timothy","提前","提摩太前","提摩太前書","1Ti","1ti","Первое послание к Тимофею","1 Ti-mô-thê","テモテヘの第一の手紙","テモテ一","디모데전서":
				bible_short_name = "提前"
				switch chap_string {
					case "":
						print_string = "提摩太前書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("提前","提摩太前書", chap_string, "1","prebklhl")
							default:
								print_string = Bible_print_string("提前","提摩太前書", chap_string, sec_string,"prebklhl")
						}
				}
			case "2 Tim","Second Timothy","提後","提摩太後","提摩太後書","2Ti","2ti","Второе послание к Тимофею","2 Ti-mô-thê","テモテヘの第二の手紙","テモテ二","디모데후서":
				bible_short_name = "提後"
				switch chap_string {
					case "":
						print_string = "提摩太後書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("提後","提摩太後書", chap_string, "1","prebklhl")
							default:
								print_string = Bible_print_string("提後","提摩太後書", chap_string, sec_string,"prebklhl")
						}
				}
			case "Titus","多","提多","提多書","Tit","tit","Послание к Титу","Tít","テトスヘの手紙","テトス","디도서":
				bible_short_name = "多"
				switch chap_string {
					case "":
						print_string = "提多書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("多","提多書", chap_string, "1","prebklhl")
							default:
								print_string = Bible_print_string("多","提多書", chap_string, sec_string,"prebklhl")
						}
				}
			case "Philem","Philemon","門","腓利","腓利門","腓利門書","Phm","phm","Послание к Филимону","Phi-lê-môn","ピレモンヘの手紙","ピレモン","빌레몬서":
				bible_short_name = "門"
				switch chap_string {
					case "":
						print_string = "腓利門書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("門","腓利門書", chap_string, "1","prebklhl")
							default:
								print_string = Bible_print_string("門","腓利門書", chap_string, sec_string,"prebklhl")
						}
				}
			case "Heb","Hebrews","來","希伯來","希伯來書","heb","Послание к Евреям","Hê-bơ-rơ","ヘブル人への手紙","ヘブル","히브리서":
				bible_short_name = "來"
				switch chap_string {
					case "":
						print_string = "希伯來書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("來","希伯來書", chap_string, "1","prebklhl")
							default:
								print_string = Bible_print_string("來","希伯來書", chap_string, sec_string,"prebklhl")
						}
				}
			case "James","雅","雅各","雅各書","Jas","jas","Послание Иакова","Gia-cơ","ヤコブの手紙","ヤコブ","야고보서":
				bible_short_name = "雅"
				switch chap_string {
					case "":
						print_string = "雅各書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("雅","雅各書", chap_string, "1","prebklhl")
							default:
								print_string = Bible_print_string("雅","雅各書", chap_string, sec_string,"prebklhl")
						}
				}
			case "1 Pet","First Peter","彼前","彼得前","彼得前書","1Pe","1pe","Первое послание Петра","1 Phi-e-rơ","ペテロの第一の手紙","ペテロ一","베드로전서":
				bible_short_name = "彼前"
				switch chap_string {
					case "":
						print_string = "彼得前書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("彼前","彼得前書", chap_string, "1","prebklhl")
							default:
								print_string = Bible_print_string("彼前","彼得前書", chap_string, sec_string,"prebklhl")
						}
				}
			case "2 Pet","Second Peter","彼後","彼得後","彼得後書","2Pe","2pe","Второе послание Петра","2 Phi-e-rơ","ペテロの第二の手紙","ペテロ","베드로후서":
				bible_short_name = "彼後"
				switch chap_string {
					case "":
						print_string = "彼得後書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("彼後","彼得後書", chap_string, "1","prebklhl")
							default:
								print_string = Bible_print_string("彼後","彼得後書", chap_string, sec_string,"prebklhl")
						}
				}
			case "1 John","First John","約一","約翰一書","約翰1","約翰1書","1Jo","1jo","Первое послание Иоанна","1 Giăng","ヨハネの第一の手紙","ヨハネ一","요한일서":
				bible_short_name = "約一"
				switch chap_string {
					case "":
						print_string = "約翰一書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("約一","約翰一書", chap_string, "1","prebklhl")
							default:
								print_string = Bible_print_string("約一","約翰一書", chap_string, sec_string,"prebklhl")
						}
				}
			case "2 John","second John","約二","約翰二書","約翰2","約翰2書","2Jo","Второе послание Иоанна","2 Giăng","ヨハネの第二の手紙","ヨハネ二","요한2서":
				bible_short_name = "約二"
				switch chap_string {
					case "":
						print_string = "約翰二書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("約二","約翰二書", chap_string, "1","prebklhl")
							default:
								print_string = Bible_print_string("約二","約翰二書", chap_string, sec_string,"prebklhl")
						}
				}
			case "3 John","Third John","約三","約翰三書","約翰3","約翰3書","3Jo","3jo","Третье послание Иоанна","3 Giăng","ヨハネの第三の手紙","ヨハネ三","요한3서":
				bible_short_name = "約三"
				switch chap_string {
					case "":
						print_string = "約翰三書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("約三","約翰三書", chap_string, "1","prebklhl")
							default:
								print_string = Bible_print_string("約三","約翰三書", chap_string, sec_string,"prebklhl")
						}
				}
			case "Jude","猶","猶大","猶大書","jude","Послание Иуды","Giu-đe","ユダの手紙","ユダ","유다서":
				bible_short_name = "猶"
				switch chap_string {
					case "":
						print_string = "猶大書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("猶","猶大書", chap_string, "1","prebklhl")
							default:
								print_string = Bible_print_string("猶","猶大書", chap_string, sec_string,"prebklhl")
						}
				}
			case "Rev","Revelation","啟","啟示","啟示錄","Re","re","ｒｅ","Ｒｅ","rev","Откровение ап. Иоанна Богослова (Апокалипсис)","Khải-huyền","ヨハネの黙示録","黙示録","요한계시록":
				bible_short_name = "啟"
				switch chap_string {
					case "":
						print_string = "啟示錄"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("啟","啟示錄", chap_string, "1","prebklhl")
							default:
								print_string = Bible_print_string("啟","啟示錄", chap_string, sec_string,"prebklhl")
						}
				}
			default:
				print_string = "聖經"
				//print_string = "你是要找 " +  reg.ReplaceAllString(text, "$3") + " 對嗎？\n對不起，我還沒學呢...\n"
		}
	case "台語聖經巴克禮全羅":
		log.Print(reg.ReplaceAllString(text, "$3"))
		switch reg.ReplaceAllString(text, "$3") {
			case "Gen","Genesis","創","創世","創世紀","創世記","Ge","ge","gen","창세기","Sáng-thế Ký","Бытие":
				bible_short_name = "創"
				switch chap_string {
					case "":
						print_string = "創世記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("創","創世記", chap_string, "1","bklcl")
							default:
								print_string = Bible_print_string("創","創世記", chap_string, sec_string,"bklcl")
						}
				}
			case "ex","Ex","Exodus","埃及","出","出埃及","出埃及記","출애굽기","エジプト","出エジプト","出エジプト記","Xuất Ê-díp-tô Ký","Исход":
				bible_short_name = "出"
				switch chap_string {
					case "":
						print_string = "出埃及記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("出","出埃及記", chap_string, "1","bklcl")
							default:
								print_string = Bible_print_string("出","出埃及記", chap_string, sec_string,"bklcl")
						}
				}
			case "Lev","Leviticus","利","利未","利未記","Le","le","Левит","Lê-vi Ký","レビ記","レビ","레위기":
				bible_short_name = "利"
				switch chap_string {
					case "":
						print_string = "利未記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("利","利未記", chap_string, "1","bklcl")
							default:
								print_string = Bible_print_string("利","利未記", chap_string, sec_string,"bklcl")
						}
				}
			case "Num","Numbers","民","民數","民數記","Nu","nu","민수기","民数","民数記","Dân-số Ký","Числа":
				bible_short_name = "民"
				switch chap_string {
					case "":
						print_string = "民數記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("民","民數記", chap_string, "1","bklcl")
							default:
								print_string = Bible_print_string("民","民數記", chap_string, sec_string,"bklcl")
						}
				}
			case "Deut","Deuteronomy","申","申命","申命記","De","de","신명기","Phục-truyền Luật-lệ Ký","Второзаконие":
				bible_short_name = "申"
				switch chap_string {
					case "":
						print_string = "申命記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("申","申命記", chap_string, "1","bklcl")
							default:
								print_string = Bible_print_string("申","申命記", chap_string, sec_string,"bklcl")
						}
				}
			case "Josh","Joshua","約書亞","約書亞記","Jos","jos","여호수아","ヨシュア記","ヨシュア","Giô-suê","Книга Иисуса Навина":
				bible_short_name = "書"
				switch chap_string {
					case "":
						print_string = "約書亞記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("書","約書亞記", chap_string, "1","bklcl")
							default:
								print_string = Bible_print_string("書","約書亞記", chap_string, sec_string,"bklcl")
						}
				}
			case "Judg","Judges","士","士師","士師記","Jud","jud","jdg","Jdg","Книга Судей израилевых","Các Quan Xét","사사기":
				bible_short_name = "士"
				switch chap_string {
					case "":
						print_string = "士師記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("士","士師記", chap_string, "1","bklcl")
							default:
								print_string = Bible_print_string("士","士師記", chap_string, sec_string,"bklcl")
						}
				}
			case "Ruth","路得","路得記","Ru","ru","Rut","rut","룻기","ルツ","ルツ記","Ru-tơ","Книга Руфи":
				bible_short_name = "得"
				switch chap_string {
					case "":
						print_string = "路得記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("得","路得記", chap_string, "1","bklcl")
							default:
								print_string = Bible_print_string("得","路得記", chap_string, sec_string,"bklcl")
						}
				}
			case "1 Sam","First Samuel","撒上","撒母耳記上","1Sa","1sa","サムエル記上","サムエル上","サム上","사무엘상","1 Sa-mu-ên","Первая книга Царств":
				bible_short_name = "撒上"
				switch chap_string {
					case "":
						print_string = "撒母耳記上"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("撒上","撒母耳記上", chap_string, "1","bklcl")
							default:
								print_string = Bible_print_string("撒上","撒母耳記上", chap_string, sec_string,"bklcl")
						}
				}
			case "2 Sam","Second Samuel","撒下","撒母耳記下","2Sa","2sa","사무엘하","サムエル記下","サムエル下","サム下","2 Sa-mu-ên","Вторая книга Царств":
				bible_short_name = "撒下"
				switch chap_string {
					case "":
						print_string = "撒母耳記下"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("撒下","撒母耳記下", chap_string, "1","bklcl")
							default:
								print_string = Bible_print_string("撒下","撒母耳記下", chap_string, sec_string,"bklcl")
						}
				}
			case "1 Kin","First Kings","王上","列王上","列王紀上","列王記上","1Ki","1ki","열왕기상","Третья книга Царств","1 Các Vua":
				bible_short_name = "王上"
				switch chap_string {
					case "":
						print_string = "列王紀上"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("王上","列王紀上", chap_string, "1","bklcl")
							default:
								print_string = Bible_print_string("王上","列王紀上", chap_string, sec_string,"bklcl")
						}
				}
			case "2 Kin","Second Kings","王下","列王下","列王記下","列王紀下","2Ki","2ki","열왕기하","2 Các Vua","Четвертая книга Царств":
				bible_short_name = "王下"
				switch chap_string {
					case "":
						print_string = "列王紀下"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("王下","列王紀下", chap_string, "1","bklcl")
							default:
								print_string = Bible_print_string("王下","列王紀下", chap_string, sec_string,"bklcl")
						}
				}
			case "1 Chr","First Chronicles","歷上","代上","歷代志上","歷代上","1Ch","1ch","Первая книга Паралипоменон","1 Sử-ký","歴上","歴代上","歴代志上","역대상":
				bible_short_name = "代上"
				switch chap_string {
					case "":
						print_string = "歷代志上"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("代上","歷代志上", chap_string, "1","bklcl")
							default:
								print_string = Bible_print_string("代上","歷代志上", chap_string, sec_string,"bklcl")
						}
				}
			case "2 Chr","Second Chronicles","代下","歷下","歷代下","歷代志下","2Ch","2ch","역대하","歴代志下","歴代下","歴下","2 Sử-ký","Вторая книга Паралипоменон":
				bible_short_name = "代下"
				switch chap_string {
					case "":
						print_string = "歷代志下"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("代下","歷代志下", chap_string, "1","bklcl")
							default:
								print_string = Bible_print_string("代下","歷代志下", chap_string, sec_string,"bklcl")
						}
				}
			case "Ezra","拉","以斯拉","以斯拉記","Ezr","ezr","Первая книга Ездры","Ê-xơ-ra","エズラ","エズラ記","에스라":
				bible_short_name = "拉"
				switch chap_string {
					case "":
						print_string = "以斯拉記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("拉","以斯拉記", chap_string, "1","bklcl")
							default:
								print_string = Bible_print_string("拉","以斯拉記", chap_string, sec_string,"bklcl")
						}
				}
			case "Neh","Nehemiah","尼","尼希米","尼希米記","Ne","ne","느헤미야","ネヘミヤ書","ネヘミヤ","Nê-hê-mi","Книга Неемии":
				bible_short_name = "尼"
				switch chap_string {
					case "":
						print_string = "尼希米記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("尼","尼希米記", chap_string, "1","bklcl")
							default:
								print_string = Bible_print_string("尼","尼希米記", chap_string, sec_string,"bklcl")
						}
				}
			case "Esth","Esther","斯","以斯帖","以斯帖記","Es","est","Есфирь","Ê-xơ-tê","エステル","エステル記","에스더":
				bible_short_name = "斯"
				switch chap_string {
					case "":
						print_string = "以斯帖記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("斯","以斯帖記", chap_string, "1","bklcl")
							default:
								print_string = Bible_print_string("斯","以斯帖記", chap_string, sec_string,"bklcl")
						}
				}
			case "Job","job","伯","約伯","約伯記","Книга Иова","Gióp","ヨブ","ヨブ記","욥기":
				bible_short_name = "伯"
				switch chap_string {
					case "":
						print_string = "約伯記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("伯","約伯記", chap_string, "1","bklcl")
							default:
								print_string = Bible_print_string("伯","約伯記", chap_string, sec_string,"bklcl")
						}
				}
			case "Ps","Psalms","詩","詩篇","ps","시편","Thi-thiên","Псалтирь":
				bible_short_name = "詩"
				switch chap_string {
					case "":
						print_string = "詩篇"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("詩","詩篇", chap_string, "1","bklcl")
							default:
								print_string = Bible_print_string("詩","詩篇", chap_string, sec_string,"bklcl")
						}
				}
			case "Prov","Proverbs","箴","箴言","Pr","pr","Притчи Соломона","Châm-ngôn","잠언":
				bible_short_name = "箴"
				switch chap_string {
					case "":
						print_string = "箴言"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("箴","箴言", chap_string, "1","bklcl")
							default:
								print_string = Bible_print_string("箴","箴言", chap_string, sec_string,"bklcl")
						}
				}
			case "Eccl","Ecclesiastes","傳","傳道","傳道書","Ec","ec","Книга Екклезиаста","или Проповедника","Truyền-đạo","伝道の書","伝道","伝","伝道書","전도서":
				bible_short_name = "傳"
				switch chap_string {
					case "":
						print_string = "傳道書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("傳","傳道書", chap_string, "1","bklcl")
							default:
								print_string = Bible_print_string("傳","傳道書", chap_string, sec_string,"bklcl")
						}
				}
			case "Song","Song of Solomon","歌","雅歌","So","so","sng","Sng","Песнь песней Соломона","Nhã-ca","아가":
				bible_short_name = "歌"
				switch chap_string {
					case "":
						print_string = "雅歌"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("歌","雅歌", chap_string, "1","bklcl")
							default:
								print_string = Bible_print_string("歌","雅歌", chap_string, sec_string,"bklcl")
						}
				}
			case "Is","Isaiah","賽","以賽","以賽亞","以賽亞書","Isa","isa","Книга пророка Исаии","Ê-sai","イザヤ書","イザヤ","이사야":
				bible_short_name = "賽"
				switch chap_string {
					case "":
						print_string = "以賽亞書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("賽","以賽亞書", chap_string, "1","bklcl")
							default:
								print_string = Bible_print_string("賽","以賽亞書", chap_string, sec_string,"bklcl")
						}
				}
			case "Jer","Jeremiah","耶","耶利米","耶利米書","jer","예레미야","エレミヤ","エレミヤ書","Giê-rê-mi","Книга пророка Иеремии":
				bible_short_name = "耶"
				switch chap_string {
					case "":
						print_string = "耶利米書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("耶","耶利米書", chap_string, "1","bklcl")
							default:
								print_string = Bible_print_string("耶","耶利米書", chap_string, sec_string,"bklcl")
						}
				}
			case "Lam","Lamentations","哀","哀歌","耶利米哀歌","La","lam","예레미야애가","Ca-thương","Плач Иеремии":
				bible_short_name = "哀"
				switch chap_string {
					case "":
						print_string = "耶利米哀歌"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("哀","耶利米哀歌", chap_string, "1","bklcl")
							default:
								print_string = Bible_print_string("哀","耶利米哀歌", chap_string, sec_string,"bklcl")
						}
				}
			case "Ezek","Ezekiel","結","以西結","以西結書","Eze","eze","에스겔","エゼキエル書","エゼキエル","Ê-xê-chi-ên","Книга пророка Иезекииля":
				bible_short_name = "結"
				switch chap_string {
					case "":
						print_string = "以西結書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("結","以西結書", chap_string, "1","bklcl")
							default:
								print_string = Bible_print_string("結","以西結書", chap_string, sec_string,"bklcl")
						}
				}
			case "Dan","Daniel","但","但以理","但以理書","Da","da","Книга пророка Даниила","Đa-ni-ên","ダニエル書","ダニエル","다니엘":
				bible_short_name = "但"
				switch chap_string {
					case "":
						print_string = "但以理書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("但","但以理書", chap_string, "1","bklcl")
							default:
								print_string = Bible_print_string("但","但以理書", chap_string, sec_string,"bklcl")
						}
				}
			case "Hos","Hosea","何","何西","何西阿","何西阿書","Ho","ho","Книга пророка Осии","Ô-sê","ホセア書","ホセア","호세아":
				bible_short_name = "何"
				switch chap_string {
					case "":
						print_string = "何西阿書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("何","何西阿書", chap_string, "1","bklcl")
							default:
								print_string = Bible_print_string("何","何西阿書", chap_string, sec_string,"bklcl")
						}
				}
			case "Joel","珥","約珥","約珥書","Joe","joe","Книга пророка Иоиля","Giô-ên","ヨエル書","ヨエル","요엘":
				bible_short_name = "珥"
				switch chap_string {
					case "":
						print_string = "約珥書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("珥","約珥書", chap_string, "1","bklcl")
							default:
								print_string = Bible_print_string("珥","約珥書", chap_string, sec_string,"bklcl")
						}
				}
			case "Amos","摩","阿摩司書","Am","am","Книга пророка Амоса","A-mốt","アモス書","アモス","아모스":
				bible_short_name = "摩"
				switch chap_string {
					case "":
						print_string = "阿摩司書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("摩","阿摩司書", chap_string, "1","bklcl")
							default:
								print_string = Bible_print_string("摩","阿摩司書", chap_string, sec_string,"bklcl")
						}
				}
			case "Obad","Obadiah","俄","俄巴底亞","俄巴底亞書","Ob","ob","오바댜","オバデヤ書","オバデヤ","Áp-đia","Книга пророка Авдия":
				bible_short_name = "俄"
				switch chap_string {
					case "":
						print_string = "俄巴底亞書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("俄","俄巴底亞書", chap_string, "1","bklcl")
							default:
								print_string = Bible_print_string("俄","俄巴底亞書", chap_string, sec_string,"bklcl")
						}
				}
			case "Jon","Jonah","拿","約拿","約拿書","jon","요나","ヨナ書","ヨナ","Giô-na","Книга пророка Ионы":
				bible_short_name = "拿"
				switch chap_string {
					case "":
						print_string = "約拿書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("拿","約拿書", chap_string, "1","bklcl")
							default:
								print_string = Bible_print_string("拿","約拿書", chap_string, sec_string,"bklcl")
						}
				}
			case "Micah","彌","彌迦","彌迦書","Mic","mic","Книга пророка Михея","Mi-chê","ミカ書","ミカ","미가":
				bible_short_name = "彌"
				switch chap_string {
					case "":
						print_string = "彌迦書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("彌","彌迦書", chap_string, "1","bklcl")
							default:
								print_string = Bible_print_string("彌","彌迦書", chap_string, sec_string,"bklcl")
						}
				}
			case "Nah","Nahum","鴻","那鴻","那鴻書","Na","na","Книга пророка Наума","Na-hum","ナホム書","ナホム","나훔":
				bible_short_name = "鴻"
				switch chap_string {
					case "":
						print_string = "那鴻書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("鴻","那鴻書", chap_string, "1","bklcl")
							default:
								print_string = Bible_print_string("鴻","那鴻書", chap_string, sec_string,"bklcl")
						}
				}
			case "Habakkuk","哈","哈巴","哈巴谷","哈巴谷書","Hab","hab","Книга пророка Аввакума","Ha-ba-cúc","ハバクク書","ハバクク","ハバ","クク","ハバ書","하박국":
				bible_short_name = "哈"
				switch chap_string {
					case "":
						print_string = "哈巴谷書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("哈","哈巴谷書", chap_string, "1","bklcl")
							default:
								print_string = Bible_print_string("哈","哈巴谷書", chap_string, sec_string,"bklcl")
						}
				}
			case "Zeph","Zephaniah","番","西番雅","西番雅書","Zep","zep","스바냐","ゼパニヤ書","ゼパニヤ","Sô-phô-ni","Книга пророка Софонии":
				bible_short_name = "番"
				switch chap_string {
					case "":
						print_string = "西番雅書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("番","西番雅書", chap_string, "1","bklcl")
							default:
								print_string = Bible_print_string("番","西番雅書", chap_string, sec_string,"bklcl")
						}
				}
			case "Haggai","該","哈該","哈該書","Hag","hag","학개","ハガイ書","ハガイ","A-ghê","Книга пророка Аггея":
				bible_short_name = "該"
				switch chap_string {
					case "":
						print_string = "哈該書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("該","哈該書", chap_string, "1","bklcl")
							default:
								print_string = Bible_print_string("該","哈該書", chap_string, sec_string,"bklcl")
						}
				}
			case "Zech","Zechariah","亞","撒迦利亞","撒迦利亞書","Zec","zec","Книга пророка Захарии","Xa-cha-ri","스가랴","ゼカリヤ書","ゼカリヤ":
				bible_short_name = "亞"
				switch chap_string {
					case "":
						print_string = "撒迦利亞書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("亞","撒迦利亞書", chap_string, "1","bklcl")
							default:
								print_string = Bible_print_string("亞","撒迦利亞書", chap_string, sec_string,"bklcl")
						}
				}
			case "Malachi","瑪","","瑪拉","瑪拉基","瑪拉基書","Mal","mal","말라기","マラキ書","マラキ","Ma-la-chi","Книга пророка Малахии":
				bible_short_name = "瑪"
				switch chap_string {
					case "":
						print_string = "瑪拉基書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("瑪","瑪拉基書", chap_string, "1","bklcl")
							default:
								print_string = Bible_print_string("瑪","瑪拉基書", chap_string, sec_string,"bklcl")
						}
				}
			case "Matt","Matthew","太","馬太","馬太福音","Mt","mt","마태복음","マタイによる福音書","マタイ","マタイによる","Ma-thi-ơ","От Матфея святое благовествование":
				bible_short_name = "太"
				switch chap_string {
					case "":
						print_string = "馬太福音"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("太","馬太福音", chap_string, "1","bklcl")
							default:
								print_string = Bible_print_string("太","馬太福音", chap_string, sec_string,"bklcl")
						}
				}
			case "Mark","可","馬可","馬可福音","Mr","mr","マルコによる福音書","マルコ","マルコによる","마가복음","Mác","От Марка святое благовествование":
				bible_short_name = "可"
				switch chap_string {
					case "":
						print_string = "馬可福音"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("可","馬可福音", chap_string, "1","bklcl")
							default:
								print_string = Bible_print_string("可","馬可福音", chap_string, sec_string,"bklcl")
						}
				}
			case "Luke","路","路加","路加福音","Lu","lu","От Луки святое благовествование","Lu-ca","ルカによる福音書","ルカ","ルカによる","누가복음":
				bible_short_name = "路"
				switch chap_string {
					case "":
						print_string = "路加福音"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("路","路加福音", chap_string, "1","bklcl")
							default:
								print_string = Bible_print_string("路","路加福音", chap_string, sec_string,"bklcl")
						}
				}
			case "John","約","約翰","約翰福音","Joh","joh","От Иоанна святое благовествование","Giăng","ヨハネによる福音書","ヨハネ","ヨハネによる","요한복음":
				bible_short_name = "約"
				switch chap_string {
					case "":
						print_string = "約翰福音"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("約","約翰福音", chap_string, "1","bklcl")
							default:
								print_string = Bible_print_string("約","約翰福音", chap_string, sec_string,"bklcl")
						}
				}
			case "Acts","徒","使徒","使徒行傳","Ac","ac","Деяния святых апостолов","Công-vụ Các Sứ-đồ","使徒行伝","사도행전":
				bible_short_name = "徒"
				switch chap_string {
					case "":
						print_string = "使徒行傳"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("徒","使徒行傳", chap_string, "1","bklcl")
							default:
								print_string = Bible_print_string("徒","使徒行傳", chap_string, sec_string,"bklcl")
						}
				}
			case "Rom","Romans","羅","羅馬","羅馬書","Ro","ro","Послание к Римлянам","Rô-ma","ローマ","ローマ人への手紙","로마서":
				bible_short_name = "羅"
				switch chap_string {
					case "":
						print_string = "羅馬書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("羅","羅馬書", chap_string, "1","bklcl")
							default:
								print_string = Bible_print_string("羅","羅馬書", chap_string, sec_string,"bklcl")
						}
				}
			case "1 Cor","First Corinthians","林前","哥林多前","哥林多前書","1Co","1co","Первое послание к Коринфянам","1 Cô-rinh-tô","コリント人への第一の手紙","コリント一","コリント人への第一","고린도전서":
				bible_short_name = "林前"
				switch chap_string {
					case "":
						print_string = "哥林多前書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("林前","哥林多前書", chap_string, "1","bklcl")
							default:
								print_string = Bible_print_string("林前","哥林多前書", chap_string, sec_string,"bklcl")
						}
				}
			case "2 Cor","Second Corinthians","林後","哥林多後","哥林多後書","2Co","2co","Второе послание к Коринфянам","2 Cô-rinh-tô","コリント人への第二の手紙","コリント二","コリント人への第二の","고린도후서":
				bible_short_name = "林後"
				switch chap_string {
					case "":
						print_string = "哥林多後書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("林後","哥林多後書", chap_string, "1","bklcl")
							default:
								print_string = Bible_print_string("林後","哥林多後書", chap_string, sec_string,"bklcl")
						}
				}
			case "Gal","Galatians","加","加拉太","加拉太書","Ga","ga","Послание к Галатам","Ga-la-ti","ガラテヤ","ガラテヤ人への手紙","갈라디아서":
				bible_short_name = "加"
				switch chap_string {
					case "":
						print_string = "加拉太書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("加","加拉太書", chap_string, "1","bklcl")
							default:
								print_string = Bible_print_string("加","加拉太書", chap_string, sec_string,"bklcl")
						}
				}
			case "Ephesians","弗","以弗所","以弗所書","Eph","eph","Послание к Ефесянам","Ê-phê-sô","エペソ人への手紙","エペソ","エペソ人","エペソ人の手紙","에베소서":
				bible_short_name = "弗"
				switch chap_string {
					case "":
						print_string = "以弗所書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("弗","以弗所書", chap_string, "1","bklcl")
							default:
								print_string = Bible_print_string("弗","以弗所書", chap_string, sec_string,"bklcl")
						}
				}
			case "Phil","Philippians","腓","腓立","腓立比","腓立比書","Php","php","빌립보서","ピリピ","ピリピ人.ピリピ人への手紙","Послание к Филиппийцам","Phi-líp":
				bible_short_name = "腓"
				switch chap_string {
					case "":
						print_string = "腓立比書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("腓","腓立比書", chap_string, "1","bklcl")
							default:
								print_string = Bible_print_string("腓","腓立比書", chap_string, sec_string,"bklcl")
						}
				}
			case "Col","col","Colossians","西","歌羅西","歌羅","歌羅西書","Послание к Колоссянам","Cô-lô-se","コロサイ人への手紙","コロサイ","コロ","골로새서":
				bible_short_name = "西"
				switch chap_string {
					case "":
						print_string = "歌羅西書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("西","歌羅西書", chap_string, "1","bklcl")
							default:
								print_string = Bible_print_string("西","歌羅西書", chap_string, sec_string,"bklcl")
						}
				}
			case "1 Thess","First Thessalonians","帖前","帖撒羅尼迦前","帖撒羅尼迦前書","1Th","1th","데살로니가전서","テサロニケ人への第一の手紙","テサ一","テサロニケ一","1 Tê-sa-lô-ni-ca","Первое послание к Фессалоникийцам (Солунянам)":
				bible_short_name = "帖前"
				switch chap_string {
					case "":
						print_string = "帖撒羅尼迦前書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("帖前","帖撒羅尼迦前書", chap_string, "1","bklcl")
							default:
								print_string = Bible_print_string("帖前","帖撒羅尼迦前書", chap_string, sec_string,"bklcl")
						}
				}
			case "2 Thess","Second Thessalonians","帖後","帖撒羅尼迦後","帖撒羅尼迦後書","2Th","2th","데살로니가후서","テサロニケ人への第二の手紙","テサ二","テサロニケ二","2 Tê-sa-lô-ni-ca","Второе послание к Фессалоникийцам (Солунянам)":
				bible_short_name = "帖後"
				switch chap_string {
					case "":
						print_string = "帖撒羅尼迦後書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("帖後","帖撒羅尼迦後書", chap_string, "1","bklcl")
							default:
								print_string = Bible_print_string("帖後","帖撒羅尼迦後書", chap_string, sec_string,"bklcl")
						}
				}
			case "1 Tim","First Timothy","提前","提摩太前","提摩太前書","1Ti","1ti","Первое послание к Тимофею","1 Ti-mô-thê","テモテヘの第一の手紙","テモテ一","디모데전서":
				bible_short_name = "提前"
				switch chap_string {
					case "":
						print_string = "提摩太前書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("提前","提摩太前書", chap_string, "1","bklcl")
							default:
								print_string = Bible_print_string("提前","提摩太前書", chap_string, sec_string,"bklcl")
						}
				}
			case "2 Tim","Second Timothy","提後","提摩太後","提摩太後書","2Ti","2ti","Второе послание к Тимофею","2 Ti-mô-thê","テモテヘの第二の手紙","テモテ二","디모데후서":
				bible_short_name = "提後"
				switch chap_string {
					case "":
						print_string = "提摩太後書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("提後","提摩太後書", chap_string, "1","bklcl")
							default:
								print_string = Bible_print_string("提後","提摩太後書", chap_string, sec_string,"bklcl")
						}
				}
			case "Titus","多","提多","提多書","Tit","tit","Послание к Титу","Tít","テトスヘの手紙","テトス","디도서":
				bible_short_name = "多"
				switch chap_string {
					case "":
						print_string = "提多書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("多","提多書", chap_string, "1","bklcl")
							default:
								print_string = Bible_print_string("多","提多書", chap_string, sec_string,"bklcl")
						}
				}
			case "Philem","Philemon","門","腓利","腓利門","腓利門書","Phm","phm","Послание к Филимону","Phi-lê-môn","ピレモンヘの手紙","ピレモン","빌레몬서":
				bible_short_name = "門"
				switch chap_string {
					case "":
						print_string = "腓利門書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("門","腓利門書", chap_string, "1","bklcl")
							default:
								print_string = Bible_print_string("門","腓利門書", chap_string, sec_string,"bklcl")
						}
				}
			case "Heb","Hebrews","來","希伯來","希伯來書","heb","Послание к Евреям","Hê-bơ-rơ","ヘブル人への手紙","ヘブル","히브리서":
				bible_short_name = "來"
				switch chap_string {
					case "":
						print_string = "希伯來書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("來","希伯來書", chap_string, "1","bklcl")
							default:
								print_string = Bible_print_string("來","希伯來書", chap_string, sec_string,"bklcl")
						}
				}
			case "James","雅","雅各","雅各書","Jas","jas","Послание Иакова","Gia-cơ","ヤコブの手紙","ヤコブ","야고보서":
				bible_short_name = "雅"
				switch chap_string {
					case "":
						print_string = "雅各書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("雅","雅各書", chap_string, "1","bklcl")
							default:
								print_string = Bible_print_string("雅","雅各書", chap_string, sec_string,"bklcl")
						}
				}
			case "1 Pet","First Peter","彼前","彼得前","彼得前書","1Pe","1pe","Первое послание Петра","1 Phi-e-rơ","ペテロの第一の手紙","ペテロ一","베드로전서":
				bible_short_name = "彼前"
				switch chap_string {
					case "":
						print_string = "彼得前書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("彼前","彼得前書", chap_string, "1","bklcl")
							default:
								print_string = Bible_print_string("彼前","彼得前書", chap_string, sec_string,"bklcl")
						}
				}
			case "2 Pet","Second Peter","彼後","彼得後","彼得後書","2Pe","2pe","Второе послание Петра","2 Phi-e-rơ","ペテロの第二の手紙","ペテロ","베드로후서":
				bible_short_name = "彼後"
				switch chap_string {
					case "":
						print_string = "彼得後書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("彼後","彼得後書", chap_string, "1","bklcl")
							default:
								print_string = Bible_print_string("彼後","彼得後書", chap_string, sec_string,"bklcl")
						}
				}
			case "1 John","First John","約一","約翰一書","約翰1","約翰1書","1Jo","1jo","Первое послание Иоанна","1 Giăng","ヨハネの第一の手紙","ヨハネ一","요한일서":
				bible_short_name = "約一"
				switch chap_string {
					case "":
						print_string = "約翰一書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("約一","約翰一書", chap_string, "1","bklcl")
							default:
								print_string = Bible_print_string("約一","約翰一書", chap_string, sec_string,"bklcl")
						}
				}
			case "2 John","second John","約二","約翰二書","約翰2","約翰2書","2Jo","Второе послание Иоанна","2 Giăng","ヨハネの第二の手紙","ヨハネ二","요한2서":
				bible_short_name = "約二"
				switch chap_string {
					case "":
						print_string = "約翰二書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("約二","約翰二書", chap_string, "1","bklcl")
							default:
								print_string = Bible_print_string("約二","約翰二書", chap_string, sec_string,"bklcl")
						}
				}
			case "3 John","Third John","約三","約翰三書","約翰3","約翰3書","3Jo","3jo","Третье послание Иоанна","3 Giăng","ヨハネの第三の手紙","ヨハネ三","요한3서":
				bible_short_name = "約三"
				switch chap_string {
					case "":
						print_string = "約翰三書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("約三","約翰三書", chap_string, "1","bklcl")
							default:
								print_string = Bible_print_string("約三","約翰三書", chap_string, sec_string,"bklcl")
						}
				}
			case "Jude","猶","猶大","猶大書","jude","Послание Иуды","Giu-đe","ユダの手紙","ユダ","유다서":
				bible_short_name = "猶"
				switch chap_string {
					case "":
						print_string = "猶大書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("猶","猶大書", chap_string, "1","bklcl")
							default:
								print_string = Bible_print_string("猶","猶大書", chap_string, sec_string,"bklcl")
						}
				}
			case "Rev","Revelation","啟","啟示","啟示錄","Re","re","ｒｅ","Ｒｅ","rev","Откровение ап. Иоанна Богослова (Апокалипсис)","Khải-huyền","ヨハネの黙示録","黙示録","요한계시록":
				bible_short_name = "啟"
				switch chap_string {
					case "":
						print_string = "啟示錄"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("啟","啟示錄", chap_string, "1","bklcl")
							default:
								print_string = Bible_print_string("啟","啟示錄", chap_string, sec_string,"bklcl")
						}
				}
			default:
				print_string = "聖經"
				//print_string = "你是要找 " +  reg.ReplaceAllString(text, "$3") + " 對嗎？\n對不起，我還沒學呢...\n"
		}
	case "台語聖經巴克禮漢羅":
		log.Print(reg.ReplaceAllString(text, "$3"))
		switch reg.ReplaceAllString(text, "$3") {
			case "Gen","Genesis","創","創世","創世紀","創世記","Ge","ge","gen","창세기","Sáng-thế Ký","Бытие":
				bible_short_name = "創"
				switch chap_string {
					case "":
						print_string = "創世記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("創","創世記", chap_string, "1","bklhl")
							default:
								print_string = Bible_print_string("創","創世記", chap_string, sec_string,"bklhl")
						}
				}
			case "ex","Ex","Exodus","埃及","出","出埃及","出埃及記","출애굽기","エジプト","出エジプト","出エジプト記","Xuất Ê-díp-tô Ký","Исход":
				bible_short_name = "出"
				switch chap_string {
					case "":
						print_string = "出埃及記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("出","出埃及記", chap_string, "1","bklhl")
							default:
								print_string = Bible_print_string("出","出埃及記", chap_string, sec_string,"bklhl")
						}
				}
			case "Lev","Leviticus","利","利未","利未記","Le","le","Левит","Lê-vi Ký","レビ記","レビ","레위기":
				bible_short_name = "利"
				switch chap_string {
					case "":
						print_string = "利未記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("利","利未記", chap_string, "1","bklhl")
							default:
								print_string = Bible_print_string("利","利未記", chap_string, sec_string,"bklhl")
						}
				}
			case "Num","Numbers","民","民數","民數記","Nu","nu","민수기","民数","民数記","Dân-số Ký","Числа":
				bible_short_name = "民"
				switch chap_string {
					case "":
						print_string = "民數記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("民","民數記", chap_string, "1","bklhl")
							default:
								print_string = Bible_print_string("民","民數記", chap_string, sec_string,"bklhl")
						}
				}
			case "Deut","Deuteronomy","申","申命","申命記","De","de","신명기","Phục-truyền Luật-lệ Ký","Второзаконие":
				bible_short_name = "申"
				switch chap_string {
					case "":
						print_string = "申命記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("申","申命記", chap_string, "1","bklhl")
							default:
								print_string = Bible_print_string("申","申命記", chap_string, sec_string,"bklhl")
						}
				}
			case "Josh","Joshua","約書亞","約書亞記","Jos","jos","여호수아","ヨシュア記","ヨシュア","Giô-suê","Книга Иисуса Навина":
				bible_short_name = "書"
				switch chap_string {
					case "":
						print_string = "約書亞記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("書","約書亞記", chap_string, "1","bklhl")
							default:
								print_string = Bible_print_string("書","約書亞記", chap_string, sec_string,"bklhl")
						}
				}
			case "Judg","Judges","士","士師","士師記","Jud","jud","jdg","Jdg","Книга Судей израилевых","Các Quan Xét","사사기":
				bible_short_name = "士"
				switch chap_string {
					case "":
						print_string = "士師記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("士","士師記", chap_string, "1","bklhl")
							default:
								print_string = Bible_print_string("士","士師記", chap_string, sec_string,"bklhl")
						}
				}
			case "Ruth","路得","路得記","Ru","ru","Rut","rut","룻기","ルツ","ルツ記","Ru-tơ","Книга Руфи":
				bible_short_name = "得"
				switch chap_string {
					case "":
						print_string = "路得記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("得","路得記", chap_string, "1","bklhl")
							default:
								print_string = Bible_print_string("得","路得記", chap_string, sec_string,"bklhl")
						}
				}
			case "1 Sam","First Samuel","撒上","撒母耳記上","1Sa","1sa","サムエル記上","サムエル上","サム上","사무엘상","1 Sa-mu-ên","Первая книга Царств":
				bible_short_name = "撒上"
				switch chap_string {
					case "":
						print_string = "撒母耳記上"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("撒上","撒母耳記上", chap_string, "1","bklhl")
							default:
								print_string = Bible_print_string("撒上","撒母耳記上", chap_string, sec_string,"bklhl")
						}
				}
			case "2 Sam","Second Samuel","撒下","撒母耳記下","2Sa","2sa","사무엘하","サムエル記下","サムエル下","サム下","2 Sa-mu-ên","Вторая книга Царств":
				bible_short_name = "撒下"
				switch chap_string {
					case "":
						print_string = "撒母耳記下"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("撒下","撒母耳記下", chap_string, "1","bklhl")
							default:
								print_string = Bible_print_string("撒下","撒母耳記下", chap_string, sec_string,"bklhl")
						}
				}
			case "1 Kin","First Kings","王上","列王上","列王紀上","列王記上","1Ki","1ki","열왕기상","Третья книга Царств","1 Các Vua":
				bible_short_name = "王上"
				switch chap_string {
					case "":
						print_string = "列王紀上"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("王上","列王紀上", chap_string, "1","bklhl")
							default:
								print_string = Bible_print_string("王上","列王紀上", chap_string, sec_string,"bklhl")
						}
				}
			case "2 Kin","Second Kings","王下","列王下","列王記下","列王紀下","2Ki","2ki","열왕기하","2 Các Vua","Четвертая книга Царств":
				bible_short_name = "王下"
				switch chap_string {
					case "":
						print_string = "列王紀下"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("王下","列王紀下", chap_string, "1","bklhl")
							default:
								print_string = Bible_print_string("王下","列王紀下", chap_string, sec_string,"bklhl")
						}
				}
			case "1 Chr","First Chronicles","歷上","代上","歷代志上","歷代上","1Ch","1ch","Первая книга Паралипоменон","1 Sử-ký","歴上","歴代上","歴代志上","역대상":
				bible_short_name = "代上"
				switch chap_string {
					case "":
						print_string = "歷代志上"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("代上","歷代志上", chap_string, "1","bklhl")
							default:
								print_string = Bible_print_string("代上","歷代志上", chap_string, sec_string,"bklhl")
						}
				}
			case "2 Chr","Second Chronicles","代下","歷下","歷代下","歷代志下","2Ch","2ch","역대하","歴代志下","歴代下","歴下","2 Sử-ký","Вторая книга Паралипоменон":
				bible_short_name = "代下"
				switch chap_string {
					case "":
						print_string = "歷代志下"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("代下","歷代志下", chap_string, "1","bklhl")
							default:
								print_string = Bible_print_string("代下","歷代志下", chap_string, sec_string,"bklhl")
						}
				}
			case "Ezra","拉","以斯拉","以斯拉記","Ezr","ezr","Первая книга Ездры","Ê-xơ-ra","エズラ","エズラ記","에스라":
				bible_short_name = "拉"
				switch chap_string {
					case "":
						print_string = "以斯拉記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("拉","以斯拉記", chap_string, "1","bklhl")
							default:
								print_string = Bible_print_string("拉","以斯拉記", chap_string, sec_string,"bklhl")
						}
				}
			case "Neh","Nehemiah","尼","尼希米","尼希米記","Ne","ne","느헤미야","ネヘミヤ書","ネヘミヤ","Nê-hê-mi","Книга Неемии":
				bible_short_name = "尼"
				switch chap_string {
					case "":
						print_string = "尼希米記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("尼","尼希米記", chap_string, "1","bklhl")
							default:
								print_string = Bible_print_string("尼","尼希米記", chap_string, sec_string,"bklhl")
						}
				}
			case "Esth","Esther","斯","以斯帖","以斯帖記","Es","est","Есфирь","Ê-xơ-tê","エステル","エステル記","에스더":
				bible_short_name = "斯"
				switch chap_string {
					case "":
						print_string = "以斯帖記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("斯","以斯帖記", chap_string, "1","bklhl")
							default:
								print_string = Bible_print_string("斯","以斯帖記", chap_string, sec_string,"bklhl")
						}
				}
			case "Job","job","伯","約伯","約伯記","Книга Иова","Gióp","ヨブ","ヨブ記","욥기":
				bible_short_name = "伯"
				switch chap_string {
					case "":
						print_string = "約伯記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("伯","約伯記", chap_string, "1","bklhl")
							default:
								print_string = Bible_print_string("伯","約伯記", chap_string, sec_string,"bklhl")
						}
				}
			case "Ps","Psalms","詩","詩篇","ps","시편","Thi-thiên","Псалтирь":
				bible_short_name = "詩"
				switch chap_string {
					case "":
						print_string = "詩篇"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("詩","詩篇", chap_string, "1","bklhl")
							default:
								print_string = Bible_print_string("詩","詩篇", chap_string, sec_string,"bklhl")
						}
				}
			case "Prov","Proverbs","箴","箴言","Pr","pr","Притчи Соломона","Châm-ngôn","잠언":
				bible_short_name = "箴"
				switch chap_string {
					case "":
						print_string = "箴言"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("箴","箴言", chap_string, "1","bklhl")
							default:
								print_string = Bible_print_string("箴","箴言", chap_string, sec_string,"bklhl")
						}
				}
			case "Eccl","Ecclesiastes","傳","傳道","傳道書","Ec","ec","Книга Екклезиаста","или Проповедника","Truyền-đạo","伝道の書","伝道","伝","伝道書","전도서":
				bible_short_name = "傳"
				switch chap_string {
					case "":
						print_string = "傳道書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("傳","傳道書", chap_string, "1","bklhl")
							default:
								print_string = Bible_print_string("傳","傳道書", chap_string, sec_string,"bklhl")
						}
				}
			case "Song","Song of Solomon","歌","雅歌","So","so","sng","Sng","Песнь песней Соломона","Nhã-ca","아가":
				bible_short_name = "歌"
				switch chap_string {
					case "":
						print_string = "雅歌"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("歌","雅歌", chap_string, "1","bklhl")
							default:
								print_string = Bible_print_string("歌","雅歌", chap_string, sec_string,"bklhl")
						}
				}
			case "Is","Isaiah","賽","以賽","以賽亞","以賽亞書","Isa","isa","Книга пророка Исаии","Ê-sai","イザヤ書","イザヤ","이사야":
				bible_short_name = "賽"
				switch chap_string {
					case "":
						print_string = "以賽亞書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("賽","以賽亞書", chap_string, "1","bklhl")
							default:
								print_string = Bible_print_string("賽","以賽亞書", chap_string, sec_string,"bklhl")
						}
				}
			case "Jer","Jeremiah","耶","耶利米","耶利米書","jer","예레미야","エレミヤ","エレミヤ書","Giê-rê-mi","Книга пророка Иеремии":
				bible_short_name = "耶"
				switch chap_string {
					case "":
						print_string = "耶利米書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("耶","耶利米書", chap_string, "1","bklhl")
							default:
								print_string = Bible_print_string("耶","耶利米書", chap_string, sec_string,"bklhl")
						}
				}
			case "Lam","Lamentations","哀","哀歌","耶利米哀歌","La","lam","예레미야애가","Ca-thương","Плач Иеремии":
				bible_short_name = "哀"
				switch chap_string {
					case "":
						print_string = "耶利米哀歌"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("哀","耶利米哀歌", chap_string, "1","bklhl")
							default:
								print_string = Bible_print_string("哀","耶利米哀歌", chap_string, sec_string,"bklhl")
						}
				}
			case "Ezek","Ezekiel","結","以西結","以西結書","Eze","eze","에스겔","エゼキエル書","エゼキエル","Ê-xê-chi-ên","Книга пророка Иезекииля":
				bible_short_name = "結"
				switch chap_string {
					case "":
						print_string = "以西結書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("結","以西結書", chap_string, "1","bklhl")
							default:
								print_string = Bible_print_string("結","以西結書", chap_string, sec_string,"bklhl")
						}
				}
			case "Dan","Daniel","但","但以理","但以理書","Da","da","Книга пророка Даниила","Đa-ni-ên","ダニエル書","ダニエル","다니엘":
				bible_short_name = "但"
				switch chap_string {
					case "":
						print_string = "但以理書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("但","但以理書", chap_string, "1","bklhl")
							default:
								print_string = Bible_print_string("但","但以理書", chap_string, sec_string,"bklhl")
						}
				}
			case "Hos","Hosea","何","何西","何西阿","何西阿書","Ho","ho","Книга пророка Осии","Ô-sê","ホセア書","ホセア","호세아":
				bible_short_name = "何"
				switch chap_string {
					case "":
						print_string = "何西阿書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("何","何西阿書", chap_string, "1","bklhl")
							default:
								print_string = Bible_print_string("何","何西阿書", chap_string, sec_string,"bklhl")
						}
				}
			case "Joel","珥","約珥","約珥書","Joe","joe","Книга пророка Иоиля","Giô-ên","ヨエル書","ヨエル","요엘":
				bible_short_name = "珥"
				switch chap_string {
					case "":
						print_string = "約珥書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("珥","約珥書", chap_string, "1","bklhl")
							default:
								print_string = Bible_print_string("珥","約珥書", chap_string, sec_string,"bklhl")
						}
				}
			case "Amos","摩","阿摩司書","Am","am","Книга пророка Амоса","A-mốt","アモス書","アモス","아모스":
				bible_short_name = "摩"
				switch chap_string {
					case "":
						print_string = "阿摩司書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("摩","阿摩司書", chap_string, "1","bklhl")
							default:
								print_string = Bible_print_string("摩","阿摩司書", chap_string, sec_string,"bklhl")
						}
				}
			case "Obad","Obadiah","俄","俄巴底亞","俄巴底亞書","Ob","ob","오바댜","オバデヤ書","オバデヤ","Áp-đia","Книга пророка Авдия":
				bible_short_name = "俄"
				switch chap_string {
					case "":
						print_string = "俄巴底亞書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("俄","俄巴底亞書", chap_string, "1","bklhl")
							default:
								print_string = Bible_print_string("俄","俄巴底亞書", chap_string, sec_string,"bklhl")
						}
				}
			case "Jon","Jonah","拿","約拿","約拿書","jon","요나","ヨナ書","ヨナ","Giô-na","Книга пророка Ионы":
				bible_short_name = "拿"
				switch chap_string {
					case "":
						print_string = "約拿書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("拿","約拿書", chap_string, "1","bklhl")
							default:
								print_string = Bible_print_string("拿","約拿書", chap_string, sec_string,"bklhl")
						}
				}
			case "Micah","彌","彌迦","彌迦書","Mic","mic","Книга пророка Михея","Mi-chê","ミカ書","ミカ","미가":
				bible_short_name = "彌"
				switch chap_string {
					case "":
						print_string = "彌迦書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("彌","彌迦書", chap_string, "1","bklhl")
							default:
								print_string = Bible_print_string("彌","彌迦書", chap_string, sec_string,"bklhl")
						}
				}
			case "Nah","Nahum","鴻","那鴻","那鴻書","Na","na","Книга пророка Наума","Na-hum","ナホム書","ナホム","나훔":
				bible_short_name = "鴻"
				switch chap_string {
					case "":
						print_string = "那鴻書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("鴻","那鴻書", chap_string, "1","bklhl")
							default:
								print_string = Bible_print_string("鴻","那鴻書", chap_string, sec_string,"bklhl")
						}
				}
			case "Habakkuk","哈","哈巴","哈巴谷","哈巴谷書","Hab","hab","Книга пророка Аввакума","Ha-ba-cúc","ハバクク書","ハバクク","ハバ","クク","ハバ書","하박국":
				bible_short_name = "哈"
				switch chap_string {
					case "":
						print_string = "哈巴谷書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("哈","哈巴谷書", chap_string, "1","bklhl")
							default:
								print_string = Bible_print_string("哈","哈巴谷書", chap_string, sec_string,"bklhl")
						}
				}
			case "Zeph","Zephaniah","番","西番雅","西番雅書","Zep","zep","스바냐","ゼパニヤ書","ゼパニヤ","Sô-phô-ni","Книга пророка Софонии":
				bible_short_name = "番"
				switch chap_string {
					case "":
						print_string = "西番雅書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("番","西番雅書", chap_string, "1","bklhl")
							default:
								print_string = Bible_print_string("番","西番雅書", chap_string, sec_string,"bklhl")
						}
				}
			case "Haggai","該","哈該","哈該書","Hag","hag","학개","ハガイ書","ハガイ","A-ghê","Книга пророка Аггея":
				bible_short_name = "該"
				switch chap_string {
					case "":
						print_string = "哈該書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("該","哈該書", chap_string, "1","bklhl")
							default:
								print_string = Bible_print_string("該","哈該書", chap_string, sec_string,"bklhl")
						}
				}
			case "Zech","Zechariah","亞","撒迦利亞","撒迦利亞書","Zec","zec","Книга пророка Захарии","Xa-cha-ri","스가랴","ゼカリヤ書","ゼカリヤ":
				bible_short_name = "亞"
				switch chap_string {
					case "":
						print_string = "撒迦利亞書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("亞","撒迦利亞書", chap_string, "1","bklhl")
							default:
								print_string = Bible_print_string("亞","撒迦利亞書", chap_string, sec_string,"bklhl")
						}
				}
			case "Malachi","瑪","瑪拉","瑪拉基","瑪拉基書","Mal","mal","말라기","マラキ書","マラキ","Ma-la-chi","Книга пророка Малахии":
				bible_short_name = "瑪"
				switch chap_string {
					case "":
						print_string = "瑪拉基書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("瑪","瑪拉基書", chap_string, "1","bklhl")
							default:
								print_string = Bible_print_string("瑪","瑪拉基書", chap_string, sec_string,"bklhl")
						}
				}
			case "Matt","Matthew","太","馬太","馬太福音","Mt","mt","마태복음","マタイによる福音書","マタイ","マタイによる","Ma-thi-ơ","От Матфея святое благовествование":
				bible_short_name = "太"
				switch chap_string {
					case "":
						print_string = "馬太福音"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("太","馬太福音", chap_string, "1","bklhl")
							default:
								print_string = Bible_print_string("太","馬太福音", chap_string, sec_string,"bklhl")
						}
				}
			case "Mark","可","馬可","馬可福音","Mr","mr","マルコによる福音書","マルコ","マルコによる","마가복음","Mác","От Марка святое благовествование":
				bible_short_name = "可"
				switch chap_string {
					case "":
						print_string = "馬可福音"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("可","馬可福音", chap_string, "1","bklhl")
							default:
								print_string = Bible_print_string("可","馬可福音", chap_string, sec_string,"bklhl")
						}
				}
			case "Luke","路","路加","路加福音","Lu","lu","От Луки святое благовествование","Lu-ca","ルカによる福音書","ルカ","ルカによる","누가복음":
				bible_short_name = "路"
				switch chap_string {
					case "":
						print_string = "路加福音"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("路","路加福音", chap_string, "1","bklhl")
							default:
								print_string = Bible_print_string("路","路加福音", chap_string, sec_string,"bklhl")
						}
				}
			case "John","約","約翰","約翰福音","Joh","joh","От Иоанна святое благовествование","Giăng","ヨハネによる福音書","ヨハネ","ヨハネによる","요한복음":
				bible_short_name = "約"
				switch chap_string {
					case "":
						print_string = "約翰福音"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("約","約翰福音", chap_string, "1","bklhl")
							default:
								print_string = Bible_print_string("約","約翰福音", chap_string, sec_string,"bklhl")
						}
				}
			case "Acts","徒","使徒","使徒行傳","Ac","ac","Деяния святых апостолов","Công-vụ Các Sứ-đồ","使徒行伝","사도행전":
				bible_short_name = "徒"
				switch chap_string {
					case "":
						print_string = "使徒行傳"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("徒","使徒行傳", chap_string, "1","bklhl")
							default:
								print_string = Bible_print_string("徒","使徒行傳", chap_string, sec_string,"bklhl")
						}
				}
			case "Rom","Romans","羅","羅馬","羅馬書","Ro","ro","Послание к Римлянам","Rô-ma","ローマ","ローマ人への手紙","로마서":
				bible_short_name = "羅"
				switch chap_string {
					case "":
						print_string = "羅馬書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("羅","羅馬書", chap_string, "1","bklhl")
							default:
								print_string = Bible_print_string("羅","羅馬書", chap_string, sec_string,"bklhl")
						}
				}
			case "1 Cor","First Corinthians","林前","哥林多前","哥林多前書","1Co","1co","Первое послание к Коринфянам","1 Cô-rinh-tô","コリント人への第一の手紙","コリント一","コリント人への第一","고린도전서":
				bible_short_name = "林前"
				switch chap_string {
					case "":
						print_string = "哥林多前書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("林前","哥林多前書", chap_string, "1","bklhl")
							default:
								print_string = Bible_print_string("林前","哥林多前書", chap_string, sec_string,"bklhl")
						}
				}
			case "2 Cor","Second Corinthians","林後","哥林多後","哥林多後書","2Co","2co","Второе послание к Коринфянам","2 Cô-rinh-tô","コリント人への第二の手紙","コリント二","コリント人への第二の","고린도후서":
				bible_short_name = "林後"
				switch chap_string {
					case "":
						print_string = "哥林多後書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("林後","哥林多後書", chap_string, "1","bklhl")
							default:
								print_string = Bible_print_string("林後","哥林多後書", chap_string, sec_string,"bklhl")
						}
				}
			case "Gal","Galatians","加","加拉太","加拉太書","Ga","ga","Послание к Галатам","Ga-la-ti","ガラテヤ","ガラテヤ人への手紙","갈라디아서":
				bible_short_name = "加"
				switch chap_string {
					case "":
						print_string = "加拉太書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("加","加拉太書", chap_string, "1","bklhl")
							default:
								print_string = Bible_print_string("加","加拉太書", chap_string, sec_string,"bklhl")
						}
				}
			case "Ephesians","弗","以弗所","以弗所書","Eph","eph","Послание к Ефесянам","Ê-phê-sô","エペソ人への手紙","エペソ","エペソ人","エペソ人の手紙","에베소서":
				bible_short_name = "弗"
				switch chap_string {
					case "":
						print_string = "以弗所書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("弗","以弗所書", chap_string, "1","bklhl")
							default:
								print_string = Bible_print_string("弗","以弗所書", chap_string, sec_string,"bklhl")
						}
				}
			case "Phil","Philippians","腓","腓立","腓立比","腓立比書","Php","php","빌립보서","ピリピ","ピリピ人.ピリピ人への手紙","Послание к Филиппийцам","Phi-líp":
				bible_short_name = "腓"
				switch chap_string {
					case "":
						print_string = "腓立比書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("腓","腓立比書", chap_string, "1","bklhl")
							default:
								print_string = Bible_print_string("腓","腓立比書", chap_string, sec_string,"bklhl")
						}
				}
			case "Col","col","Colossians","西","歌羅西","歌羅","歌羅西書","Послание к Колоссянам","Cô-lô-se","コロサイ人への手紙","コロサイ","コロ","골로새서":
				bible_short_name = "西"
				switch chap_string {
					case "":
						print_string = "歌羅西書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("西","歌羅西書", chap_string, "1","bklhl")
							default:
								print_string = Bible_print_string("西","歌羅西書", chap_string, sec_string,"bklhl")
						}
				}
			case "1 Thess","First Thessalonians","帖前","帖撒羅尼迦前","帖撒羅尼迦前書","1Th","1th","데살로니가전서","テサロニケ人への第一の手紙","テサ一","テサロニケ一","1 Tê-sa-lô-ni-ca","Первое послание к Фессалоникийцам (Солунянам)":
				bible_short_name = "帖前"
				switch chap_string {
					case "":
						print_string = "帖撒羅尼迦前書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("帖前","帖撒羅尼迦前書", chap_string, "1","bklhl")
							default:
								print_string = Bible_print_string("帖前","帖撒羅尼迦前書", chap_string, sec_string,"bklhl")
						}
				}
			case "2 Thess","Second Thessalonians","帖後","帖撒羅尼迦後","帖撒羅尼迦後書","2Th","2th","데살로니가후서","テサロニケ人への第二の手紙","テサ二","テサロニケ二","2 Tê-sa-lô-ni-ca","Второе послание к Фессалоникийцам (Солунянам)":
				bible_short_name = "帖後"
				switch chap_string {
					case "":
						print_string = "帖撒羅尼迦後書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("帖後","帖撒羅尼迦後書", chap_string, "1","bklhl")
							default:
								print_string = Bible_print_string("帖後","帖撒羅尼迦後書", chap_string, sec_string,"bklhl")
						}
				}
			case "1 Tim","First Timothy","提前","提摩太前","提摩太前書","1Ti","1ti","Первое послание к Тимофею","1 Ti-mô-thê","テモテヘの第一の手紙","テモテ一","디모데전서":
				bible_short_name = "提前"
				switch chap_string {
					case "":
						print_string = "提摩太前書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("提前","提摩太前書", chap_string, "1","bklhl")
							default:
								print_string = Bible_print_string("提前","提摩太前書", chap_string, sec_string,"bklhl")
						}
				}
			case "2 Tim","Second Timothy","提後","提摩太後","提摩太後書","2Ti","2ti","Второе послание к Тимофею","2 Ti-mô-thê","テモテヘの第二の手紙","テモテ二","디모데후서":
				bible_short_name = "提後"
				switch chap_string {
					case "":
						print_string = "提摩太後書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("提後","提摩太後書", chap_string, "1","bklhl")
							default:
								print_string = Bible_print_string("提後","提摩太後書", chap_string, sec_string,"bklhl")
						}
				}
			case "Titus","多","提多","提多書","Tit","tit","Послание к Титу","Tít","テトスヘの手紙","テトス","디도서":
				bible_short_name = "多"
				switch chap_string {
					case "":
						print_string = "提多書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("多","提多書", chap_string, "1","bklhl")
							default:
								print_string = Bible_print_string("多","提多書", chap_string, sec_string,"bklhl")
						}
				}
			case "Philem","Philemon","門","腓利","腓利門","腓利門書","Phm","phm","Послание к Филимону","Phi-lê-môn","ピレモンヘの手紙","ピレモン","빌레몬서":
				bible_short_name = "門"
				switch chap_string {
					case "":
						print_string = "腓利門書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("門","腓利門書", chap_string, "1","bklhl")
							default:
								print_string = Bible_print_string("門","腓利門書", chap_string, sec_string,"bklhl")
						}
				}
			case "Heb","Hebrews","來","希伯來","希伯來書","heb","Послание к Евреям","Hê-bơ-rơ","ヘブル人への手紙","ヘブル","히브리서":
				bible_short_name = "來"
				switch chap_string {
					case "":
						print_string = "希伯來書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("來","希伯來書", chap_string, "1","bklhl")
							default:
								print_string = Bible_print_string("來","希伯來書", chap_string, sec_string,"bklhl")
						}
				}
			case "James","雅","雅各","雅各書","Jas","jas","Послание Иакова","Gia-cơ","ヤコブの手紙","ヤコブ","야고보서":
				bible_short_name = "雅"
				switch chap_string {
					case "":
						print_string = "雅各書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("雅","雅各書", chap_string, "1","bklhl")
							default:
								print_string = Bible_print_string("雅","雅各書", chap_string, sec_string,"bklhl")
						}
				}
			case "1 Pet","First Peter","彼前","彼得前","彼得前書","1Pe","1pe","Первое послание Петра","1 Phi-e-rơ","ペテロの第一の手紙","ペテロ一","베드로전서":
				bible_short_name = "彼前"
				switch chap_string {
					case "":
						print_string = "彼得前書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("彼前","彼得前書", chap_string, "1","bklhl")
							default:
								print_string = Bible_print_string("彼前","彼得前書", chap_string, sec_string,"bklhl")
						}
				}
			case "2 Pet","Second Peter","彼後","彼得後","彼得後書","2Pe","2pe","Второе послание Петра","2 Phi-e-rơ","ペテロの第二の手紙","ペテロ","베드로후서":
				bible_short_name = "彼後"
				switch chap_string {
					case "":
						print_string = "彼得後書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("彼後","彼得後書", chap_string, "1","bklhl")
							default:
								print_string = Bible_print_string("彼後","彼得後書", chap_string, sec_string,"bklhl")
						}
				}
			case "1 John","First John","約一","約翰一書","約翰1","約翰1書","1Jo","1jo","Первое послание Иоанна","1 Giăng","ヨハネの第一の手紙","ヨハネ一","요한일서":
				bible_short_name = "約一"
				switch chap_string {
					case "":
						print_string = "約翰一書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("約一","約翰一書", chap_string, "1","bklhl")
							default:
								print_string = Bible_print_string("約一","約翰一書", chap_string, sec_string,"bklhl")
						}
				}
			case "2 John","second John","約二","約翰二書","約翰2","約翰2書","2Jo","Второе послание Иоанна","2 Giăng","ヨハネの第二の手紙","ヨハネ二","요한2서":
				bible_short_name = "約二"
				switch chap_string {
					case "":
						print_string = "約翰二書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("約二","約翰二書", chap_string, "1","bklhl")
							default:
								print_string = Bible_print_string("約二","約翰二書", chap_string, sec_string,"bklhl")
						}
				}
			case "3 John","Third John","約三","約翰三書","約翰3","約翰3書","3Jo","3jo","Третье послание Иоанна","3 Giăng","ヨハネの第三の手紙","ヨハネ三","요한3서":
				bible_short_name = "約三"
				switch chap_string {
					case "":
						print_string = "約翰三書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("約三","約翰三書", chap_string, "1","bklhl")
							default:
								print_string = Bible_print_string("約三","約翰三書", chap_string, sec_string,"bklhl")
						}
				}
			case "Jude","猶","猶大","猶大書","jude","Послание Иуды","Giu-đe","ユダの手紙","ユダ","유다서":
				bible_short_name = "猶"
				switch chap_string {
					case "":
						print_string = "猶大書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("猶","猶大書", chap_string, "1","bklhl")
							default:
								print_string = Bible_print_string("猶","猶大書", chap_string, sec_string,"bklhl")
						}
				}
			case "Rev","Revelation","啟","啟示","啟示錄","Re","re","ｒｅ","Ｒｅ","rev","Откровение ап. Иоанна Богослова (Апокалипсис)","Khải-huyền","ヨハネの黙示録","黙示録","요한계시록":
				bible_short_name = "啟"
				switch chap_string {
					case "":
						print_string = "啟示錄"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("啟","啟示錄", chap_string, "1","bklhl")
							default:
								print_string = Bible_print_string("啟","啟示錄", chap_string, sec_string,"bklhl")
						}
				}
			default:
				print_string = "聖經"
				//print_string = "你是要找 " +  reg.ReplaceAllString(text, "$3") + " 對嗎？\n對不起，我還沒學呢...\n"
		}
	case "台語聖經全羅","全民台語聖經全羅":
		log.Print(reg.ReplaceAllString(text, "$3"))
		switch reg.ReplaceAllString(text, "$3") {
			case "Gen","Genesis","創","創世","創世紀","創世記","Ge","ge","gen","창세기","Sáng-thế Ký","Бытие":
				bible_short_name = "創"
				switch chap_string {
					case "":
						print_string = "創世記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("創","創世記", chap_string, "1","sgebklcl")
							default:
								print_string = Bible_print_string("創","創世記", chap_string, sec_string,"sgebklcl")
						}
				}
			case "ex","Ex","Exodus","埃及","出","出埃及","出埃及記","출애굽기","エジプト","出エジプト","出エジプト記","Xuất Ê-díp-tô Ký","Исход":
				bible_short_name = "出"
				switch chap_string {
					case "":
						print_string = "出埃及記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("出","出埃及記", chap_string, "1","sgebklcl")
							default:
								print_string = Bible_print_string("出","出埃及記", chap_string, sec_string,"sgebklcl")
						}
				}
			case "Lev","Leviticus","利","利未","利未記","Le","le","Левит","Lê-vi Ký","レビ記","レビ","레위기":
				bible_short_name = "利"
				switch chap_string {
					case "":
						print_string = "利未記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("利","利未記", chap_string, "1","sgebklcl")
							default:
								print_string = Bible_print_string("利","利未記", chap_string, sec_string,"sgebklcl")
						}
				}
			case "Num","Numbers","民","民數","民數記","Nu","nu","민수기","民数","民数記","Dân-số Ký","Числа":
				bible_short_name = "民"
				switch chap_string {
					case "":
						print_string = "民數記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("民","民數記", chap_string, "1","sgebklcl")
							default:
								print_string = Bible_print_string("民","民數記", chap_string, sec_string,"sgebklcl")
						}
				}
			case "Deut","Deuteronomy","申","申命","申命記","De","de","신명기","Phục-truyền Luật-lệ Ký","Второзаконие":
				bible_short_name = "申"
				switch chap_string {
					case "":
						print_string = "申命記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("申","申命記", chap_string, "1","sgebklcl")
							default:
								print_string = Bible_print_string("申","申命記", chap_string, sec_string,"sgebklcl")
						}
				}
			case "Josh","Joshua","約書亞","約書亞記","Jos","jos","여호수아","ヨシュア記","ヨシュア","Giô-suê","Книга Иисуса Навина":
				bible_short_name = "書"
				switch chap_string {
					case "":
						print_string = "約書亞記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("書","約書亞記", chap_string, "1","sgebklcl")
							default:
								print_string = Bible_print_string("書","約書亞記", chap_string, sec_string,"sgebklcl")
						}
				}
			case "Judg","Judges","士","士師","士師記","Jud","jud","jdg","Jdg","Книга Судей израилевых","Các Quan Xét","사사기":
				bible_short_name = "士"
				switch chap_string {
					case "":
						print_string = "士師記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("士","士師記", chap_string, "1","sgebklcl")
							default:
								print_string = Bible_print_string("士","士師記", chap_string, sec_string,"sgebklcl")
						}
				}
			case "Ruth","路得","路得記","Ru","ru","Rut","rut","룻기","ルツ","ルツ記","Ru-tơ","Книга Руфи":
				bible_short_name = "得"
				switch chap_string {
					case "":
						print_string = "路得記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("得","路得記", chap_string, "1","sgebklcl")
							default:
								print_string = Bible_print_string("得","路得記", chap_string, sec_string,"sgebklcl")
						}
				}
			case "1 Sam","First Samuel","撒上","撒母耳記上","1Sa","1sa","サムエル記上","サムエル上","サム上","사무엘상","1 Sa-mu-ên","Первая книга Царств":
				bible_short_name = "撒上"
				switch chap_string {
					case "":
						print_string = "撒母耳記上"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("撒上","撒母耳記上", chap_string, "1","sgebklcl")
							default:
								print_string = Bible_print_string("撒上","撒母耳記上", chap_string, sec_string,"sgebklcl")
						}
				}
			case "2 Sam","Second Samuel","撒下","撒母耳記下","2Sa","2sa","사무엘하","サムエル記下","サムエル下","サム下","2 Sa-mu-ên","Вторая книга Царств":
				bible_short_name = "撒下"
				switch chap_string {
					case "":
						print_string = "撒母耳記下"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("撒下","撒母耳記下", chap_string, "1","sgebklcl")
							default:
								print_string = Bible_print_string("撒下","撒母耳記下", chap_string, sec_string,"sgebklcl")
						}
				}
			case "1 Kin","First Kings","王上","列王上","列王紀上","列王記上","1Ki","1ki","열왕기상","Третья книга Царств","1 Các Vua":
				bible_short_name = "王上"
				switch chap_string {
					case "":
						print_string = "列王紀上"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("王上","列王紀上", chap_string, "1","sgebklcl")
							default:
								print_string = Bible_print_string("王上","列王紀上", chap_string, sec_string,"sgebklcl")
						}
				}
			case "2 Kin","Second Kings","王下","列王下","列王記下","列王紀下","2Ki","2ki","열왕기하","2 Các Vua","Четвертая книга Царств":
				bible_short_name = "王下"
				switch chap_string {
					case "":
						print_string = "列王紀下"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("王下","列王紀下", chap_string, "1","sgebklcl")
							default:
								print_string = Bible_print_string("王下","列王紀下", chap_string, sec_string,"sgebklcl")
						}
				}
			case "1 Chr","First Chronicles","歷上","代上","歷代志上","歷代上","1Ch","1ch","Первая книга Паралипоменон","1 Sử-ký","歴上","歴代上","歴代志上","역대상":
				bible_short_name = "代上"
				switch chap_string {
					case "":
						print_string = "歷代志上"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("代上","歷代志上", chap_string, "1","sgebklcl")
							default:
								print_string = Bible_print_string("代上","歷代志上", chap_string, sec_string,"sgebklcl")
						}
				}
			case "2 Chr","Second Chronicles","代下","歷下","歷代下","歷代志下","2Ch","2ch","역대하","歴代志下","歴代下","歴下","2 Sử-ký","Вторая книга Паралипоменон":
				bible_short_name = "代下"
				switch chap_string {
					case "":
						print_string = "歷代志下"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("代下","歷代志下", chap_string, "1","sgebklcl")
							default:
								print_string = Bible_print_string("代下","歷代志下", chap_string, sec_string,"sgebklcl")
						}
				}
			case "Ezra","拉","以斯拉","以斯拉記","Ezr","ezr","Первая книга Ездры","Ê-xơ-ra","エズラ","エズラ記","에스라":
				bible_short_name = "拉"
				switch chap_string {
					case "":
						print_string = "以斯拉記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("拉","以斯拉記", chap_string, "1","sgebklcl")
							default:
								print_string = Bible_print_string("拉","以斯拉記", chap_string, sec_string,"sgebklcl")
						}
				}
			case "Neh","Nehemiah","尼","尼希米","尼希米記","Ne","ne","느헤미야","ネヘミヤ書","ネヘミヤ","Nê-hê-mi","Книга Неемии":
				bible_short_name = "尼"
				switch chap_string {
					case "":
						print_string = "尼希米記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("尼","尼希米記", chap_string, "1","sgebklcl")
							default:
								print_string = Bible_print_string("尼","尼希米記", chap_string, sec_string,"sgebklcl")
						}
				}
			case "Esth","Esther","斯","以斯帖","以斯帖記","Es","est","Есфирь","Ê-xơ-tê","エステル","エステル記","에스더":
				bible_short_name = "斯"
				switch chap_string {
					case "":
						print_string = "以斯帖記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("斯","以斯帖記", chap_string, "1","sgebklcl")
							default:
								print_string = Bible_print_string("斯","以斯帖記", chap_string, sec_string,"sgebklcl")
						}
				}
			case "Job","job","伯","約伯","約伯記","Книга Иова","Gióp","ヨブ","ヨブ記","욥기":
				bible_short_name = "伯"
				switch chap_string {
					case "":
						print_string = "約伯記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("伯","約伯記", chap_string, "1","sgebklcl")
							default:
								print_string = Bible_print_string("伯","約伯記", chap_string, sec_string,"sgebklcl")
						}
				}
			case "Ps","Psalms","詩","詩篇","ps","시편","Thi-thiên","Псалтирь":
				bible_short_name = "詩"
				switch chap_string {
					case "":
						print_string = "詩篇"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("詩","詩篇", chap_string, "1","sgebklcl")
							default:
								print_string = Bible_print_string("詩","詩篇", chap_string, sec_string,"sgebklcl")
						}
				}
			case "Prov","Proverbs","箴","箴言","Pr","pr","Притчи Соломона","Châm-ngôn","잠언":
				bible_short_name = "箴"
				switch chap_string {
					case "":
						print_string = "箴言"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("箴","箴言", chap_string, "1","sgebklcl")
							default:
								print_string = Bible_print_string("箴","箴言", chap_string, sec_string,"sgebklcl")
						}
				}
			case "Eccl","Ecclesiastes","傳","傳道","傳道書","Ec","ec","Книга Екклезиаста","или Проповедника","Truyền-đạo","伝道の書","伝道","伝","伝道書","전도서":
				bible_short_name = "傳"
				switch chap_string {
					case "":
						print_string = "傳道書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("傳","傳道書", chap_string, "1","sgebklcl")
							default:
								print_string = Bible_print_string("傳","傳道書", chap_string, sec_string,"sgebklcl")
						}
				}
			case "Song","Song of Solomon","歌","雅歌","So","so","sng","Sng","Песнь песней Соломона","Nhã-ca","아가":
				bible_short_name = "歌"
				switch chap_string {
					case "":
						print_string = "雅歌"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("歌","雅歌", chap_string, "1","sgebklcl")
							default:
								print_string = Bible_print_string("歌","雅歌", chap_string, sec_string,"sgebklcl")
						}
				}
			case "Is","Isaiah","賽","以賽","以賽亞","以賽亞書","Isa","isa","Книга пророка Исаии","Ê-sai","イザヤ書","イザヤ","이사야":
				bible_short_name = "賽"
				switch chap_string {
					case "":
						print_string = "以賽亞書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("賽","以賽亞書", chap_string, "1","sgebklcl")
							default:
								print_string = Bible_print_string("賽","以賽亞書", chap_string, sec_string,"sgebklcl")
						}
				}
			case "Jer","Jeremiah","耶","耶利米","耶利米書","jer","예레미야","エレミヤ","エレミヤ書","Giê-rê-mi","Книга пророка Иеремии":
				bible_short_name = "耶"
				switch chap_string {
					case "":
						print_string = "耶利米書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("耶","耶利米書", chap_string, "1","sgebklcl")
							default:
								print_string = Bible_print_string("耶","耶利米書", chap_string, sec_string,"sgebklcl")
						}
				}
			case "Lam","Lamentations","哀","哀歌","耶利米哀歌","La","lam","예레미야애가","Ca-thương","Плач Иеремии":
				bible_short_name = "哀"
				switch chap_string {
					case "":
						print_string = "耶利米哀歌"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("哀","耶利米哀歌", chap_string, "1","sgebklcl")
							default:
								print_string = Bible_print_string("哀","耶利米哀歌", chap_string, sec_string,"sgebklcl")
						}
				}
			case "Ezek","Ezekiel","結","以西結","以西結書","Eze","eze","에스겔","エゼキエル書","エゼキエル","Ê-xê-chi-ên","Книга пророка Иезекииля":
				bible_short_name = "結"
				switch chap_string {
					case "":
						print_string = "以西結書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("結","以西結書", chap_string, "1","sgebklcl")
							default:
								print_string = Bible_print_string("結","以西結書", chap_string, sec_string,"sgebklcl")
						}
				}
			case "Dan","Daniel","但","但以理","但以理書","Da","da","Книга пророка Даниила","Đa-ni-ên","ダニエル書","ダニエル","다니엘":
				bible_short_name = "但"
				switch chap_string {
					case "":
						print_string = "但以理書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("但","但以理書", chap_string, "1","sgebklcl")
							default:
								print_string = Bible_print_string("但","但以理書", chap_string, sec_string,"sgebklcl")
						}
				}
			case "Hos","Hosea","何","何西","何西阿","何西阿書","Ho","ho","Книга пророка Осии","Ô-sê","ホセア書","ホセア","호세아":
				bible_short_name = "何"
				switch chap_string {
					case "":
						print_string = "何西阿書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("何","何西阿書", chap_string, "1","sgebklcl")
							default:
								print_string = Bible_print_string("何","何西阿書", chap_string, sec_string,"sgebklcl")
						}
				}
			case "Joel","珥","約珥","約珥書","Joe","joe","Книга пророка Иоиля","Giô-ên","ヨエル書","ヨエル","요엘":
				bible_short_name = "珥"
				switch chap_string {
					case "":
						print_string = "約珥書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("珥","約珥書", chap_string, "1","sgebklcl")
							default:
								print_string = Bible_print_string("珥","約珥書", chap_string, sec_string,"sgebklcl")
						}
				}
			case "Amos","摩","阿摩司書","Am","am","Книга пророка Амоса","A-mốt","アモス書","アモス","아모스":
				bible_short_name = "摩"
				switch chap_string {
					case "":
						print_string = "阿摩司書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("摩","阿摩司書", chap_string, "1","sgebklcl")
							default:
								print_string = Bible_print_string("摩","阿摩司書", chap_string, sec_string,"sgebklcl")
						}
				}
			case "Obad","Obadiah","俄","俄巴底亞","俄巴底亞書","Ob","ob","오바댜","オバデヤ書","オバデヤ","Áp-đia","Книга пророка Авдия":
				bible_short_name = "俄"
				switch chap_string {
					case "":
						print_string = "俄巴底亞書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("俄","俄巴底亞書", chap_string, "1","sgebklcl")
							default:
								print_string = Bible_print_string("俄","俄巴底亞書", chap_string, sec_string,"sgebklcl")
						}
				}
			case "Jon","Jonah","拿","約拿","約拿書","jon","요나","ヨナ書","ヨナ","Giô-na","Книга пророка Ионы":
				bible_short_name = "拿"
				switch chap_string {
					case "":
						print_string = "約拿書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("拿","約拿書", chap_string, "1","sgebklcl")
							default:
								print_string = Bible_print_string("拿","約拿書", chap_string, sec_string,"sgebklcl")
						}
				}
			case "Micah","彌","彌迦","彌迦書","Mic","mic","Книга пророка Михея","Mi-chê","ミカ書","ミカ","미가":
				bible_short_name = "彌"
				switch chap_string {
					case "":
						print_string = "彌迦書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("彌","彌迦書", chap_string, "1","sgebklcl")
							default:
								print_string = Bible_print_string("彌","彌迦書", chap_string, sec_string,"sgebklcl")
						}
				}
			case "Nah","Nahum","鴻","那鴻","那鴻書","Na","na","Книга пророка Наума","Na-hum","ナホム書","ナホム","나훔":
				bible_short_name = "鴻"
				switch chap_string {
					case "":
						print_string = "那鴻書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("鴻","那鴻書", chap_string, "1","sgebklcl")
							default:
								print_string = Bible_print_string("鴻","那鴻書", chap_string, sec_string,"sgebklcl")
						}
				}
			case "Habakkuk","哈","哈巴","哈巴谷","哈巴谷書","Hab","hab","Книга пророка Аввакума","Ha-ba-cúc","ハバクク書","ハバクク","ハバ","クク","ハバ書","하박국":
				bible_short_name = "哈"
				switch chap_string {
					case "":
						print_string = "哈巴谷書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("哈","哈巴谷書", chap_string, "1","sgebklcl")
							default:
								print_string = Bible_print_string("哈","哈巴谷書", chap_string, sec_string,"sgebklcl")
						}
				}
			case "Zeph","Zephaniah","番","西番雅","西番雅書","Zep","zep","스바냐","ゼパニヤ書","ゼパニヤ","Sô-phô-ni","Книга пророка Софонии":
				bible_short_name = "番"
				switch chap_string {
					case "":
						print_string = "西番雅書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("番","西番雅書", chap_string, "1","sgebklcl")
							default:
								print_string = Bible_print_string("番","西番雅書", chap_string, sec_string,"sgebklcl")
						}
				}
			case "Haggai","該","哈該","哈該書","Hag","hag","학개","ハガイ書","ハガイ","A-ghê","Книга пророка Аггея":
				bible_short_name = "該"
				switch chap_string {
					case "":
						print_string = "哈該書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("該","哈該書", chap_string, "1","sgebklcl")
							default:
								print_string = Bible_print_string("該","哈該書", chap_string, sec_string,"sgebklcl")
						}
				}
			case "Zech","Zechariah","亞","撒迦利亞","撒迦利亞書","Zec","zec","Книга пророка Захарии","Xa-cha-ri","스가랴","ゼカリヤ書","ゼカリヤ":
				bible_short_name = "亞"
				switch chap_string {
					case "":
						print_string = "撒迦利亞書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("亞","撒迦利亞書", chap_string, "1","sgebklcl")
							default:
								print_string = Bible_print_string("亞","撒迦利亞書", chap_string, sec_string,"sgebklcl")
						}
				}
			case "Malachi","瑪","瑪拉","瑪拉基","瑪拉基書","Mal","mal","말라기","マラキ書","マラキ","Ma-la-chi","Книга пророка Малахии":
				bible_short_name = "瑪"
				switch chap_string {
					case "":
						print_string = "瑪拉基書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("瑪","瑪拉基書", chap_string, "1","sgebklcl")
							default:
								print_string = Bible_print_string("瑪","瑪拉基書", chap_string, sec_string,"sgebklcl")
						}
				}
			case "Matt","Matthew","太","馬太","馬太福音","Mt","mt","마태복음","マタイによる福音書","マタイ","マタイによる","Ma-thi-ơ","От Матфея святое благовествование":
				bible_short_name = "太"
				switch chap_string {
					case "":
						print_string = "馬太福音"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("太","馬太福音", chap_string, "1","sgebklcl")
							default:
								print_string = Bible_print_string("太","馬太福音", chap_string, sec_string,"sgebklcl")
						}
				}
			case "Mark","可","馬可","馬可福音","Mr","mr","マルコによる福音書","マルコ","マルコによる","마가복음","Mác","От Марка святое благовествование":
				bible_short_name = "可"
				switch chap_string {
					case "":
						print_string = "馬可福音"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("可","馬可福音", chap_string, "1","sgebklcl")
							default:
								print_string = Bible_print_string("可","馬可福音", chap_string, sec_string,"sgebklcl")
						}
				}
			case "Luke","路","路加","路加福音","Lu","lu","От Луки святое благовествование","Lu-ca","ルカによる福音書","ルカ","ルカによる","누가복음":
				bible_short_name = "路"
				switch chap_string {
					case "":
						print_string = "路加福音"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("路","路加福音", chap_string, "1","sgebklcl")
							default:
								print_string = Bible_print_string("路","路加福音", chap_string, sec_string,"sgebklcl")
						}
				}
			case "John","約","約翰","約翰福音","Joh","joh","От Иоанна святое благовествование","Giăng","ヨハネによる福音書","ヨハネ","ヨハネによる","요한복음":
				bible_short_name = "約"
				switch chap_string {
					case "":
						print_string = "約翰福音"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("約","約翰福音", chap_string, "1","sgebklcl")
							default:
								print_string = Bible_print_string("約","約翰福音", chap_string, sec_string,"sgebklcl")
						}
				}
			case "Acts","徒","使徒","使徒行傳","Ac","ac","Деяния святых апостолов","Công-vụ Các Sứ-đồ","使徒行伝","사도행전":
				bible_short_name = "徒"
				switch chap_string {
					case "":
						print_string = "使徒行傳"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("徒","使徒行傳", chap_string, "1","sgebklcl")
							default:
								print_string = Bible_print_string("徒","使徒行傳", chap_string, sec_string,"sgebklcl")
						}
				}
			case "Rom","Romans","羅","羅馬","羅馬書","Ro","ro","Послание к Римлянам","Rô-ma","ローマ","ローマ人への手紙","로마서":
				bible_short_name = "羅"
				switch chap_string {
					case "":
						print_string = "羅馬書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("羅","羅馬書", chap_string, "1","sgebklcl")
							default:
								print_string = Bible_print_string("羅","羅馬書", chap_string, sec_string,"sgebklcl")
						}
				}
			case "1 Cor","First Corinthians","林前","哥林多前","哥林多前書","1Co","1co","Первое послание к Коринфянам","1 Cô-rinh-tô","コリント人への第一の手紙","コリント一","コリント人への第一","고린도전서":
				bible_short_name = "林前"
				switch chap_string {
					case "":
						print_string = "哥林多前書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("林前","哥林多前書", chap_string, "1","sgebklcl")
							default:
								print_string = Bible_print_string("林前","哥林多前書", chap_string, sec_string,"sgebklcl")
						}
				}
			case "2 Cor","Second Corinthians","林後","哥林多後","哥林多後書","2Co","2co","Второе послание к Коринфянам","2 Cô-rinh-tô","コリント人への第二の手紙","コリント二","コリント人への第二の","고린도후서":
				bible_short_name = "林後"
				switch chap_string {
					case "":
						print_string = "哥林多後書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("林後","哥林多後書", chap_string, "1","sgebklcl")
							default:
								print_string = Bible_print_string("林後","哥林多後書", chap_string, sec_string,"sgebklcl")
						}
				}
			case "Gal","Galatians","加","加拉太","加拉太書","Ga","ga","Послание к Галатам","Ga-la-ti","ガラテヤ","ガラテヤ人への手紙","갈라디아서":
				bible_short_name = "加"
				switch chap_string {
					case "":
						print_string = "加拉太書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("加","加拉太書", chap_string, "1","sgebklcl")
							default:
								print_string = Bible_print_string("加","加拉太書", chap_string, sec_string,"sgebklcl")
						}
				}
			case "Ephesians","弗","以弗所","以弗所書","Eph","eph","Послание к Ефесянам","Ê-phê-sô","エペソ人への手紙","エペソ","エペソ人","エペソ人の手紙","에베소서":
				bible_short_name = "弗"
				switch chap_string {
					case "":
						print_string = "以弗所書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("弗","以弗所書", chap_string, "1","sgebklcl")
							default:
								print_string = Bible_print_string("弗","以弗所書", chap_string, sec_string,"sgebklcl")
						}
				}
			case "Phil","Philippians","腓","腓立","腓立比","腓立比書","Php","php","빌립보서","ピリピ","ピリピ人.ピリピ人への手紙","Послание к Филиппийцам","Phi-líp":
				bible_short_name = "腓"
				switch chap_string {
					case "":
						print_string = "腓立比書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("腓","腓立比書", chap_string, "1","sgebklcl")
							default:
								print_string = Bible_print_string("腓","腓立比書", chap_string, sec_string,"sgebklcl")
						}
				}
			case "Col","col","Colossians","西","歌羅西","歌羅","歌羅西書","Послание к Колоссянам","Cô-lô-se","コロサイ人への手紙","コロサイ","コロ","골로새서":
				bible_short_name = "西"
				switch chap_string {
					case "":
						print_string = "歌羅西書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("西","歌羅西書", chap_string, "1","sgebklcl")
							default:
								print_string = Bible_print_string("西","歌羅西書", chap_string, sec_string,"sgebklcl")
						}
				}
			case "1 Thess","First Thessalonians","帖前","帖撒羅尼迦前","帖撒羅尼迦前書","1Th","1th","데살로니가전서","テサロニケ人への第一の手紙","テサ一","テサロニケ一","1 Tê-sa-lô-ni-ca","Первое послание к Фессалоникийцам (Солунянам)":
				bible_short_name = "帖前"
				switch chap_string {
					case "":
						print_string = "帖撒羅尼迦前書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("帖前","帖撒羅尼迦前書", chap_string, "1","sgebklcl")
							default:
								print_string = Bible_print_string("帖前","帖撒羅尼迦前書", chap_string, sec_string,"sgebklcl")
						}
				}
			case "2 Thess","Second Thessalonians","帖後","帖撒羅尼迦後","帖撒羅尼迦後書","2Th","2th","데살로니가후서","テサロニケ人への第二の手紙","テサ二","テサロニケ二","2 Tê-sa-lô-ni-ca","Второе послание к Фессалоникийцам (Солунянам)":
				bible_short_name = "帖後"
				switch chap_string {
					case "":
						print_string = "帖撒羅尼迦後書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("帖後","帖撒羅尼迦後書", chap_string, "1","sgebklcl")
							default:
								print_string = Bible_print_string("帖後","帖撒羅尼迦後書", chap_string, sec_string,"sgebklcl")
						}
				}
			case "1 Tim","First Timothy","提前","提摩太前","提摩太前書","1Ti","1ti","Первое послание к Тимофею","1 Ti-mô-thê","テモテヘの第一の手紙","テモテ一","디모데전서":
				bible_short_name = "提前"
				switch chap_string {
					case "":
						print_string = "提摩太前書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("提前","提摩太前書", chap_string, "1","sgebklcl")
							default:
								print_string = Bible_print_string("提前","提摩太前書", chap_string, sec_string,"sgebklcl")
						}
				}
			case "2 Tim","Second Timothy","提後","提摩太後","提摩太後書","2Ti","2ti","Второе послание к Тимофею","2 Ti-mô-thê","テモテヘの第二の手紙","テモテ二","디모데후서":
				bible_short_name = "提後"
				switch chap_string {
					case "":
						print_string = "提摩太後書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("提後","提摩太後書", chap_string, "1","sgebklcl")
							default:
								print_string = Bible_print_string("提後","提摩太後書", chap_string, sec_string,"sgebklcl")
						}
				}
			case "Titus","多","提多","提多書","Tit","tit","Послание к Титу","Tít","テトスヘの手紙","テトス","디도서":
				bible_short_name = "多"
				switch chap_string {
					case "":
						print_string = "提多書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("多","提多書", chap_string, "1","sgebklcl")
							default:
								print_string = Bible_print_string("多","提多書", chap_string, sec_string,"sgebklcl")
						}
				}
			case "Philem","Philemon","門","腓利","腓利門","腓利門書","Phm","phm","Послание к Филимону","Phi-lê-môn","ピレモンヘの手紙","ピレモン","빌레몬서":
				bible_short_name = "門"
				switch chap_string {
					case "":
						print_string = "腓利門書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("門","腓利門書", chap_string, "1","sgebklcl")
							default:
								print_string = Bible_print_string("門","腓利門書", chap_string, sec_string,"sgebklcl")
						}
				}
			case "Heb","Hebrews","來","希伯來","希伯來書","heb","Послание к Евреям","Hê-bơ-rơ","ヘブル人への手紙","ヘブル","히브리서":
				bible_short_name = "來"
				switch chap_string {
					case "":
						print_string = "希伯來書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("來","希伯來書", chap_string, "1","sgebklcl")
							default:
								print_string = Bible_print_string("來","希伯來書", chap_string, sec_string,"sgebklcl")
						}
				}
			case "James","雅","雅各","雅各書","Jas","jas","Послание Иакова","Gia-cơ","ヤコブの手紙","ヤコブ","야고보서":
				bible_short_name = "雅"
				switch chap_string {
					case "":
						print_string = "雅各書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("雅","雅各書", chap_string, "1","sgebklcl")
							default:
								print_string = Bible_print_string("雅","雅各書", chap_string, sec_string,"sgebklcl")
						}
				}
			case "1 Pet","First Peter","彼前","彼得前","彼得前書","1Pe","1pe","Первое послание Петра","1 Phi-e-rơ","ペテロの第一の手紙","ペテロ一","베드로전서":
				bible_short_name = "彼前"
				switch chap_string {
					case "":
						print_string = "彼得前書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("彼前","彼得前書", chap_string, "1","sgebklcl")
							default:
								print_string = Bible_print_string("彼前","彼得前書", chap_string, sec_string,"sgebklcl")
						}
				}
			case "2 Pet","Second Peter","彼後","彼得後","彼得後書","2Pe","2pe","Второе послание Петра","2 Phi-e-rơ","ペテロの第二の手紙","ペテロ","베드로후서":
				bible_short_name = "彼後"
				switch chap_string {
					case "":
						print_string = "彼得後書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("彼後","彼得後書", chap_string, "1","sgebklcl")
							default:
								print_string = Bible_print_string("彼後","彼得後書", chap_string, sec_string,"sgebklcl")
						}
				}
			case "1 John","First John","約一","約翰一書","約翰1","約翰1書","1Jo","1jo","Первое послание Иоанна","1 Giăng","ヨハネの第一の手紙","ヨハネ一","요한일서":
				bible_short_name = "約一"
				switch chap_string {
					case "":
						print_string = "約翰一書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("約一","約翰一書", chap_string, "1","sgebklcl")
							default:
								print_string = Bible_print_string("約一","約翰一書", chap_string, sec_string,"sgebklcl")
						}
				}
			case "2 John","second John","約二","約翰二書","約翰2","約翰2書","2Jo","Второе послание Иоанна","2 Giăng","ヨハネの第二の手紙","ヨハネ二","요한2서":
				bible_short_name = "約二"
				switch chap_string {
					case "":
						print_string = "約翰二書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("約二","約翰二書", chap_string, "1","sgebklcl")
							default:
								print_string = Bible_print_string("約二","約翰二書", chap_string, sec_string,"sgebklcl")
						}
				}
			case "3 John","Third John","約三","約翰三書","約翰3","約翰3書","3Jo","3jo","Третье послание Иоанна","3 Giăng","ヨハネの第三の手紙","ヨハネ三","요한3서":
				bible_short_name = "約三"
				switch chap_string {
					case "":
						print_string = "約翰三書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("約三","約翰三書", chap_string, "1","sgebklcl")
							default:
								print_string = Bible_print_string("約三","約翰三書", chap_string, sec_string,"sgebklcl")
						}
				}
			case "Jude","猶","猶大","猶大書","jude","Послание Иуды","Giu-đe","ユダの手紙","ユダ","유다서":
				bible_short_name = "猶"
				switch chap_string {
					case "":
						print_string = "猶大書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("猶","猶大書", chap_string, "1","sgebklcl")
							default:
								print_string = Bible_print_string("猶","猶大書", chap_string, sec_string,"sgebklcl")
						}
				}
			case "Rev","Revelation","啟","啟示","啟示錄","Re","re","ｒｅ","Ｒｅ","rev","Откровение ап. Иоанна Богослова (Апокалипсис)","Khải-huyền","ヨハネの黙示録","黙示録","요한계시록":
				bible_short_name = "啟"
				switch chap_string {
					case "":
						print_string = "啟示錄"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("啟","啟示錄", chap_string, "1","sgebklcl")
							default:
								print_string = Bible_print_string("啟","啟示錄", chap_string, sec_string,"sgebklcl")
						}
				}
			default:
				print_string = "聖經"
				//print_string = "你是要找 " +  reg.ReplaceAllString(text, "$3") + " 對嗎？\n對不起，我還沒學呢...\n"
		}
	case "台語聖經","閩南語聖經","台語聖經漢羅","全民台語聖經漢羅": //台語（全民台語聖經漢羅）
		log.Print(reg.ReplaceAllString(text, "$3"))
		switch reg.ReplaceAllString(text, "$3") {
			case "Gen","Genesis","創","創世","創世紀","創世記","Ge","ge","gen","창세기","Sáng-thế Ký","Бытие":
				bible_short_name = "創"
				switch chap_string {
					case "":
						print_string = "創世記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("創","創世記", chap_string, "1","sgebklhl")
							default:
								print_string = Bible_print_string("創","創世記", chap_string, sec_string,"sgebklhl")
						}
				}
			case "ex","Ex","Exodus","埃及","出","出埃及","出埃及記","출애굽기","エジプト","出エジプト","出エジプト記","Xuất Ê-díp-tô Ký","Исход":
				bible_short_name = "出"
				switch chap_string {
					case "":
						print_string = "出埃及記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("出","出埃及記", chap_string, "1","sgebklhl")
							default:
								print_string = Bible_print_string("出","出埃及記", chap_string, sec_string,"sgebklhl")
						}
				}
			case "Lev","Leviticus","利","利未","利未記","Le","le","Левит","Lê-vi Ký","レビ記","レビ","레위기":
				bible_short_name = "利"
				switch chap_string {
					case "":
						print_string = "利未記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("利","利未記", chap_string, "1","sgebklhl")
							default:
								print_string = Bible_print_string("利","利未記", chap_string, sec_string,"sgebklhl")
						}
				}
			case "Num","Numbers","民","民數","民數記","Nu","nu","민수기","民数","民数記","Dân-số Ký","Числа":
				bible_short_name = "民"
				switch chap_string {
					case "":
						print_string = "民數記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("民","民數記", chap_string, "1","sgebklhl")
							default:
								print_string = Bible_print_string("民","民數記", chap_string, sec_string,"sgebklhl")
						}
				}
			case "Deut","Deuteronomy","申","申命","申命記","De","de","신명기","Phục-truyền Luật-lệ Ký","Второзаконие":
				bible_short_name = "申"
				switch chap_string {
					case "":
						print_string = "申命記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("申","申命記", chap_string, "1","sgebklhl")
							default:
								print_string = Bible_print_string("申","申命記", chap_string, sec_string,"sgebklhl")
						}
				}
			case "Josh","Joshua","約書亞","約書亞記","Jos","jos","여호수아","ヨシュア記","ヨシュア","Giô-suê","Книга Иисуса Навина":
				bible_short_name = "書"
				switch chap_string {
					case "":
						print_string = "約書亞記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("書","約書亞記", chap_string, "1","sgebklhl")
							default:
								print_string = Bible_print_string("書","約書亞記", chap_string, sec_string,"sgebklhl")
						}
				}
			case "Judg","Judges","士","士師","士師記","Jud","jud","jdg","Jdg","Книга Судей израилевых","Các Quan Xét","사사기":
				bible_short_name = "士"
				switch chap_string {
					case "":
						print_string = "士師記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("士","士師記", chap_string, "1","sgebklhl")
							default:
								print_string = Bible_print_string("士","士師記", chap_string, sec_string,"sgebklhl")
						}
				}
			case "Ruth","路得","路得記","Ru","ru","Rut","rut","룻기","ルツ","ルツ記","Ru-tơ","Книга Руфи":
				bible_short_name = "得"
				switch chap_string {
					case "":
						print_string = "路得記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("得","路得記", chap_string, "1","sgebklhl")
							default:
								print_string = Bible_print_string("得","路得記", chap_string, sec_string,"sgebklhl")
						}
				}
			case "1 Sam","First Samuel","撒上","撒母耳記上","1Sa","1sa","サムエル記上","サムエル上","サム上","사무엘상","1 Sa-mu-ên","Первая книга Царств":
				bible_short_name = "撒上"
				switch chap_string {
					case "":
						print_string = "撒母耳記上"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("撒上","撒母耳記上", chap_string, "1","sgebklhl")
							default:
								print_string = Bible_print_string("撒上","撒母耳記上", chap_string, sec_string,"sgebklhl")
						}
				}
			case "2 Sam","Second Samuel","撒下","撒母耳記下","2Sa","2sa","사무엘하","サムエル記下","サムエル下","サム下","2 Sa-mu-ên","Вторая книга Царств":
				bible_short_name = "撒下"
				switch chap_string {
					case "":
						print_string = "撒母耳記下"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("撒下","撒母耳記下", chap_string, "1","sgebklhl")
							default:
								print_string = Bible_print_string("撒下","撒母耳記下", chap_string, sec_string,"sgebklhl")
						}
				}
			case "1 Kin","First Kings","王上","列王上","列王紀上","列王記上","1Ki","1ki","열왕기상","Третья книга Царств","1 Các Vua":
				bible_short_name = "王上"
				switch chap_string {
					case "":
						print_string = "列王紀上"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("王上","列王紀上", chap_string, "1","sgebklhl")
							default:
								print_string = Bible_print_string("王上","列王紀上", chap_string, sec_string,"sgebklhl")
						}
				}
			case "2 Kin","Second Kings","王下","列王下","列王記下","列王紀下","2Ki","2ki","열왕기하","2 Các Vua","Четвертая книга Царств":
				bible_short_name = "王下"
				switch chap_string {
					case "":
						print_string = "列王紀下"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("王下","列王紀下", chap_string, "1","sgebklhl")
							default:
								print_string = Bible_print_string("王下","列王紀下", chap_string, sec_string,"sgebklhl")
						}
				}
			case "1 Chr","First Chronicles","歷上","代上","歷代志上","歷代上","1Ch","1ch","Первая книга Паралипоменон","1 Sử-ký","歴上","歴代上","歴代志上","역대상":
				bible_short_name = "代上"
				switch chap_string {
					case "":
						print_string = "歷代志上"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("代上","歷代志上", chap_string, "1","sgebklhl")
							default:
								print_string = Bible_print_string("代上","歷代志上", chap_string, sec_string,"sgebklhl")
						}
				}
			case "2 Chr","Second Chronicles","代下","歷下","歷代下","歷代志下","2Ch","2ch","역대하","歴代志下","歴代下","歴下","2 Sử-ký","Вторая книга Паралипоменон":
				bible_short_name = "代下"
				switch chap_string {
					case "":
						print_string = "歷代志下"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("代下","歷代志下", chap_string, "1","sgebklhl")
							default:
								print_string = Bible_print_string("代下","歷代志下", chap_string, sec_string,"sgebklhl")
						}
				}
			case "Ezra","拉","以斯拉","以斯拉記","Ezr","ezr","Первая книга Ездры","Ê-xơ-ra","エズラ","エズラ記","에스라":
				bible_short_name = "拉"
				switch chap_string {
					case "":
						print_string = "以斯拉記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("拉","以斯拉記", chap_string, "1","sgebklhl")
							default:
								print_string = Bible_print_string("拉","以斯拉記", chap_string, sec_string,"sgebklhl")
						}
				}
			case "Neh","Nehemiah","尼","尼希米","尼希米記","Ne","ne","느헤미야","ネヘミヤ書","ネヘミヤ","Nê-hê-mi","Книга Неемии":
				bible_short_name = "尼"
				switch chap_string {
					case "":
						print_string = "尼希米記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("尼","尼希米記", chap_string, "1","sgebklhl")
							default:
								print_string = Bible_print_string("尼","尼希米記", chap_string, sec_string,"sgebklhl")
						}
				}
			case "Esth","Esther","斯","以斯帖","以斯帖記","Es","est","Есфирь","Ê-xơ-tê","エステル","エステル記","에스더":
				bible_short_name = "斯"
				switch chap_string {
					case "":
						print_string = "以斯帖記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("斯","以斯帖記", chap_string, "1","sgebklhl")
							default:
								print_string = Bible_print_string("斯","以斯帖記", chap_string, sec_string,"sgebklhl")
						}
				}
			case "Job","job","伯","約伯","約伯記","Книга Иова","Gióp","ヨブ","ヨブ記","욥기":
				bible_short_name = "伯"
				switch chap_string {
					case "":
						print_string = "約伯記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("伯","約伯記", chap_string, "1","sgebklhl")
							default:
								print_string = Bible_print_string("伯","約伯記", chap_string, sec_string,"sgebklhl")
						}
				}
			case "Ps","Psalms","詩","詩篇","ps","시편","Thi-thiên","Псалтирь":
				bible_short_name = "詩"
				switch chap_string {
					case "":
						print_string = "詩篇"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("詩","詩篇", chap_string, "1","sgebklhl")
							default:
								print_string = Bible_print_string("詩","詩篇", chap_string, sec_string,"sgebklhl")
						}
				}
			case "Prov","Proverbs","箴","箴言","Pr","pr","Притчи Соломона","Châm-ngôn","잠언":
				bible_short_name = "箴"
				switch chap_string {
					case "":
						print_string = "箴言"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("箴","箴言", chap_string, "1","sgebklhl")
							default:
								print_string = Bible_print_string("箴","箴言", chap_string, sec_string,"sgebklhl")
						}
				}
			case "Eccl","Ecclesiastes","傳","傳道","傳道書","Ec","ec","Книга Екклезиаста","или Проповедника","Truyền-đạo","伝道の書","伝道","伝","伝道書","전도서":
				bible_short_name = "傳"
				switch chap_string {
					case "":
						print_string = "傳道書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("傳","傳道書", chap_string, "1","sgebklhl")
							default:
								print_string = Bible_print_string("傳","傳道書", chap_string, sec_string,"sgebklhl")
						}
				}
			case "Song","Song of Solomon","歌","雅歌","So","so","sng","Sng","Песнь песней Соломона","Nhã-ca","아가":
				bible_short_name = "歌"
				switch chap_string {
					case "":
						print_string = "雅歌"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("歌","雅歌", chap_string, "1","sgebklhl")
							default:
								print_string = Bible_print_string("歌","雅歌", chap_string, sec_string,"sgebklhl")
						}
				}
			case "Is","Isaiah","賽","以賽","以賽亞","以賽亞書","Isa","isa","Книга пророка Исаии","Ê-sai","イザヤ書","イザヤ","이사야":
				bible_short_name = "賽"
				switch chap_string {
					case "":
						print_string = "以賽亞書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("賽","以賽亞書", chap_string, "1","sgebklhl")
							default:
								print_string = Bible_print_string("賽","以賽亞書", chap_string, sec_string,"sgebklhl")
						}
				}
			case "Jer","Jeremiah","耶","耶利米","耶利米書","jer","예레미야","エレミヤ","エレミヤ書","Giê-rê-mi","Книга пророка Иеремии":
				bible_short_name = "耶"
				switch chap_string {
					case "":
						print_string = "耶利米書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("耶","耶利米書", chap_string, "1","sgebklhl")
							default:
								print_string = Bible_print_string("耶","耶利米書", chap_string, sec_string,"sgebklhl")
						}
				}
			case "Lam","Lamentations","哀","哀歌","耶利米哀歌","La","lam","예레미야애가","Ca-thương","Плач Иеремии":
				bible_short_name = "哀"
				switch chap_string {
					case "":
						print_string = "耶利米哀歌"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("哀","耶利米哀歌", chap_string, "1","sgebklhl")
							default:
								print_string = Bible_print_string("哀","耶利米哀歌", chap_string, sec_string,"sgebklhl")
						}
				}
			case "Ezek","Ezekiel","結","以西結","以西結書","Eze","eze","에스겔","エゼキエル書","エゼキエル","Ê-xê-chi-ên","Книга пророка Иезекииля":
				bible_short_name = "結"
				switch chap_string {
					case "":
						print_string = "以西結書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("結","以西結書", chap_string, "1","sgebklhl")
							default:
								print_string = Bible_print_string("結","以西結書", chap_string, sec_string,"sgebklhl")
						}
				}
			case "Dan","Daniel","但","但以理","但以理書","Da","da","Книга пророка Даниила","Đa-ni-ên","ダニエル書","ダニエル","다니엘":
				bible_short_name = "但"
				switch chap_string {
					case "":
						print_string = "但以理書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("但","但以理書", chap_string, "1","sgebklhl")
							default:
								print_string = Bible_print_string("但","但以理書", chap_string, sec_string,"sgebklhl")
						}
				}
			case "Hos","Hosea","何","何西","何西阿","何西阿書","Ho","ho","Книга пророка Осии","Ô-sê","ホセア書","ホセア","호세아":
				bible_short_name = "何"
				switch chap_string {
					case "":
						print_string = "何西阿書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("何","何西阿書", chap_string, "1","sgebklhl")
							default:
								print_string = Bible_print_string("何","何西阿書", chap_string, sec_string,"sgebklhl")
						}
				}
			case "Joel","珥","約珥","約珥書","Joe","joe","Книга пророка Иоиля","Giô-ên","ヨエル書","ヨエル","요엘":
				bible_short_name = "珥"
				switch chap_string {
					case "":
						print_string = "約珥書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("珥","約珥書", chap_string, "1","sgebklhl")
							default:
								print_string = Bible_print_string("珥","約珥書", chap_string, sec_string,"sgebklhl")
						}
				}
			case "Amos","摩","阿摩司書","Am","am","Книга пророка Амоса","A-mốt","アモス書","アモス","아모스":
				bible_short_name = "摩"
				switch chap_string {
					case "":
						print_string = "阿摩司書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("摩","阿摩司書", chap_string, "1","sgebklhl")
							default:
								print_string = Bible_print_string("摩","阿摩司書", chap_string, sec_string,"sgebklhl")
						}
				}
			case "Obad","Obadiah","俄","俄巴底亞","俄巴底亞書","Ob","ob","오바댜","オバデヤ書","オバデヤ","Áp-đia","Книга пророка Авдия":
				bible_short_name = "俄"
				switch chap_string {
					case "":
						print_string = "俄巴底亞書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("俄","俄巴底亞書", chap_string, "1","sgebklhl")
							default:
								print_string = Bible_print_string("俄","俄巴底亞書", chap_string, sec_string,"sgebklhl")
						}
				}
			case "Jon","Jonah","拿","約拿","約拿書","jon","요나","ヨナ書","ヨナ","Giô-na","Книга пророка Ионы":
				bible_short_name = "拿"
				switch chap_string {
					case "":
						print_string = "約拿書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("拿","約拿書", chap_string, "1","sgebklhl")
							default:
								print_string = Bible_print_string("拿","約拿書", chap_string, sec_string,"sgebklhl")
						}
				}
			case "Micah","彌","彌迦","彌迦書","Mic","mic","Книга пророка Михея","Mi-chê","ミカ書","ミカ","미가":
				bible_short_name = "彌"
				switch chap_string {
					case "":
						print_string = "彌迦書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("彌","彌迦書", chap_string, "1","sgebklhl")
							default:
								print_string = Bible_print_string("彌","彌迦書", chap_string, sec_string,"sgebklhl")
						}
				}
			case "Nah","Nahum","鴻","那鴻","那鴻書","Na","na","Книга пророка Наума","Na-hum","ナホム書","ナホム","나훔":
				bible_short_name = "鴻"
				switch chap_string {
					case "":
						print_string = "那鴻書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("鴻","那鴻書", chap_string, "1","sgebklhl")
							default:
								print_string = Bible_print_string("鴻","那鴻書", chap_string, sec_string,"sgebklhl")
						}
				}
			case "Habakkuk","哈","哈巴","哈巴谷","哈巴谷書","Hab","hab","Книга пророка Аввакума","Ha-ba-cúc","ハバクク書","ハバクク","ハバ","クク","ハバ書","하박국":
				bible_short_name = "哈"
				switch chap_string {
					case "":
						print_string = "哈巴谷書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("哈","哈巴谷書", chap_string, "1","sgebklhl")
							default:
								print_string = Bible_print_string("哈","哈巴谷書", chap_string, sec_string,"sgebklhl")
						}
				}
			case "Zeph","Zephaniah","番","西番雅","西番雅書","Zep","zep","스바냐","ゼパニヤ書","ゼパニヤ","Sô-phô-ni","Книга пророка Софонии":
				bible_short_name = "番"
				switch chap_string {
					case "":
						print_string = "西番雅書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("番","西番雅書", chap_string, "1","sgebklhl")
							default:
								print_string = Bible_print_string("番","西番雅書", chap_string, sec_string,"sgebklhl")
						}
				}
			case "Haggai","該","哈該","哈該書","Hag","hag","학개","ハガイ書","ハガイ","A-ghê","Книга пророка Аггея":
				bible_short_name = "該"
				switch chap_string {
					case "":
						print_string = "哈該書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("該","哈該書", chap_string, "1","sgebklhl")
							default:
								print_string = Bible_print_string("該","哈該書", chap_string, sec_string,"sgebklhl")
						}
				}
			case "Zech","Zechariah","亞","撒迦利亞","撒迦利亞書","Zec","zec","Книга пророка Захарии","Xa-cha-ri","스가랴","ゼカリヤ書","ゼカリヤ":
				bible_short_name = "亞"
				switch chap_string {
					case "":
						print_string = "撒迦利亞書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("亞","撒迦利亞書", chap_string, "1","sgebklhl")
							default:
								print_string = Bible_print_string("亞","撒迦利亞書", chap_string, sec_string,"sgebklhl")
						}
				}
			case "Malachi","瑪","瑪拉","瑪拉基","瑪拉基書","Mal","mal","말라기","マラキ書","マラキ","Ma-la-chi","Книга пророка Малахии":
				bible_short_name = "瑪"
				switch chap_string {
					case "":
						print_string = "瑪拉基書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("瑪","瑪拉基書", chap_string, "1","sgebklhl")
							default:
								print_string = Bible_print_string("瑪","瑪拉基書", chap_string, sec_string,"sgebklhl")
						}
				}
			case "Matt","Matthew","太","馬太","馬太福音","Mt","mt","마태복음","マタイによる福音書","マタイ","マタイによる","Ma-thi-ơ","От Матфея святое благовествование":
				bible_short_name = "太"
				switch chap_string {
					case "":
						print_string = "馬太福音"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("太","馬太福音", chap_string, "1","sgebklhl")
							default:
								print_string = Bible_print_string("太","馬太福音", chap_string, sec_string,"sgebklhl")
						}
				}
			case "Mark","可","馬可","馬可福音","Mr","mr","マルコによる福音書","マルコ","マルコによる","마가복음","Mác","От Марка святое благовествование":
				bible_short_name = "可"
				switch chap_string {
					case "":
						print_string = "馬可福音"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("可","馬可福音", chap_string, "1","sgebklhl")
							default:
								print_string = Bible_print_string("可","馬可福音", chap_string, sec_string,"sgebklhl")
						}
				}
			case "Luke","路","路加","路加福音","Lu","lu","От Луки святое благовествование","Lu-ca","ルカによる福音書","ルカ","ルカによる","누가복음":
				bible_short_name = "路"
				switch chap_string {
					case "":
						print_string = "路加福音"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("路","路加福音", chap_string, "1","sgebklhl")
							default:
								print_string = Bible_print_string("路","路加福音", chap_string, sec_string,"sgebklhl")
						}
				}
			case "John","約","約翰","約翰福音","Joh","joh","От Иоанна святое благовествование","Giăng","ヨハネによる福音書","ヨハネ","ヨハネによる","요한복음":
				bible_short_name = "約"
				switch chap_string {
					case "":
						print_string = "約翰福音"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("約","約翰福音", chap_string, "1","sgebklhl")
							default:
								print_string = Bible_print_string("約","約翰福音", chap_string, sec_string,"sgebklhl")
						}
				}
			case "Acts","徒","使徒","使徒行傳","Ac","ac","Деяния святых апостолов","Công-vụ Các Sứ-đồ","使徒行伝","사도행전":
				bible_short_name = "徒"
				switch chap_string {
					case "":
						print_string = "使徒行傳"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("徒","使徒行傳", chap_string, "1","sgebklhl")
							default:
								print_string = Bible_print_string("徒","使徒行傳", chap_string, sec_string,"sgebklhl")
						}
				}
			case "Rom","Romans","羅","羅馬","羅馬書","Ro","ro","Послание к Римлянам","Rô-ma","ローマ","ローマ人への手紙","로마서":
				bible_short_name = "羅"
				switch chap_string {
					case "":
						print_string = "羅馬書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("羅","羅馬書", chap_string, "1","sgebklhl")
							default:
								print_string = Bible_print_string("羅","羅馬書", chap_string, sec_string,"sgebklhl")
						}
				}
			case "1 Cor","First Corinthians","林前","哥林多前","哥林多前書","1Co","1co","Первое послание к Коринфянам","1 Cô-rinh-tô","コリント人への第一の手紙","コリント一","コリント人への第一","고린도전서":
				bible_short_name = "林前"
				switch chap_string {
					case "":
						print_string = "哥林多前書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("林前","哥林多前書", chap_string, "1","sgebklhl")
							default:
								print_string = Bible_print_string("林前","哥林多前書", chap_string, sec_string,"sgebklhl")
						}
				}
			case "2 Cor","Second Corinthians","林後","哥林多後","哥林多後書","2Co","2co","Второе послание к Коринфянам","2 Cô-rinh-tô","コリント人への第二の手紙","コリント二","コリント人への第二の","고린도후서":
				bible_short_name = "林後"
				switch chap_string {
					case "":
						print_string = "哥林多後書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("林後","哥林多後書", chap_string, "1","sgebklhl")
							default:
								print_string = Bible_print_string("林後","哥林多後書", chap_string, sec_string,"sgebklhl")
						}
				}
			case "Gal","Galatians","加","加拉太","加拉太書","Ga","ga","Послание к Галатам","Ga-la-ti","ガラテヤ","ガラテヤ人への手紙","갈라디아서":
				bible_short_name = "加"
				switch chap_string {
					case "":
						print_string = "加拉太書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("加","加拉太書", chap_string, "1","sgebklhl")
							default:
								print_string = Bible_print_string("加","加拉太書", chap_string, sec_string,"sgebklhl")
						}
				}
			case "Ephesians","弗","以弗所","以弗所書","Eph","eph","Послание к Ефесянам","Ê-phê-sô","エペソ人への手紙","エペソ","エペソ人","エペソ人の手紙","에베소서":
				bible_short_name = "弗"
				switch chap_string {
					case "":
						print_string = "以弗所書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("弗","以弗所書", chap_string, "1","sgebklhl")
							default:
								print_string = Bible_print_string("弗","以弗所書", chap_string, sec_string,"sgebklhl")
						}
				}
			case "Phil","Philippians","腓","腓立","腓立比","腓立比書","Php","php","빌립보서","ピリピ","ピリピ人.ピリピ人への手紙","Послание к Филиппийцам","Phi-líp":
				bible_short_name = "腓"
				switch chap_string {
					case "":
						print_string = "腓立比書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("腓","腓立比書", chap_string, "1","sgebklhl")
							default:
								print_string = Bible_print_string("腓","腓立比書", chap_string, sec_string,"sgebklhl")
						}
				}
			case "Col","col","Colossians","西","歌羅西","歌羅","歌羅西書","Послание к Колоссянам","Cô-lô-se","コロサイ人への手紙","コロサイ","コロ","골로새서":
				bible_short_name = "西"
				switch chap_string {
					case "":
						print_string = "歌羅西書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("西","歌羅西書", chap_string, "1","sgebklhl")
							default:
								print_string = Bible_print_string("西","歌羅西書", chap_string, sec_string,"sgebklhl")
						}
				}
			case "1 Thess","First Thessalonians","帖前","帖撒羅尼迦前","帖撒羅尼迦前書","1Th","1th","데살로니가전서","テサロニケ人への第一の手紙","テサ一","テサロニケ一","1 Tê-sa-lô-ni-ca","Первое послание к Фессалоникийцам (Солунянам)":
				bible_short_name = "帖前"
				switch chap_string {
					case "":
						print_string = "帖撒羅尼迦前書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("帖前","帖撒羅尼迦前書", chap_string, "1","sgebklhl")
							default:
								print_string = Bible_print_string("帖前","帖撒羅尼迦前書", chap_string, sec_string,"sgebklhl")
						}
				}
			case "2 Thess","Second Thessalonians","帖後","帖撒羅尼迦後","帖撒羅尼迦後書","2Th","2th","데살로니가후서","テサロニケ人への第二の手紙","テサ二","テサロニケ二","2 Tê-sa-lô-ni-ca","Второе послание к Фессалоникийцам (Солунянам)":
				bible_short_name = "帖後"
				switch chap_string {
					case "":
						print_string = "帖撒羅尼迦後書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("帖後","帖撒羅尼迦後書", chap_string, "1","sgebklhl")
							default:
								print_string = Bible_print_string("帖後","帖撒羅尼迦後書", chap_string, sec_string,"sgebklhl")
						}
				}
			case "1 Tim","First Timothy","提前","提摩太前","提摩太前書","1Ti","1ti","Первое послание к Тимофею","1 Ti-mô-thê","テモテヘの第一の手紙","テモテ一","디모데전서":
				bible_short_name = "提前"
				switch chap_string {
					case "":
						print_string = "提摩太前書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("提前","提摩太前書", chap_string, "1","sgebklhl")
							default:
								print_string = Bible_print_string("提前","提摩太前書", chap_string, sec_string,"sgebklhl")
						}
				}
			case "2 Tim","Second Timothy","提後","提摩太後","提摩太後書","2Ti","2ti","Второе послание к Тимофею","2 Ti-mô-thê","テモテヘの第二の手紙","テモテ二","디모데후서":
				bible_short_name = "提後"
				switch chap_string {
					case "":
						print_string = "提摩太後書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("提後","提摩太後書", chap_string, "1","sgebklhl")
							default:
								print_string = Bible_print_string("提後","提摩太後書", chap_string, sec_string,"sgebklhl")
						}
				}
			case "Titus","多","提多","提多書","Tit","tit","Послание к Титу","Tít","テトスヘの手紙","テトス","디도서":
				bible_short_name = "多"
				switch chap_string {
					case "":
						print_string = "提多書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("多","提多書", chap_string, "1","sgebklhl")
							default:
								print_string = Bible_print_string("多","提多書", chap_string, sec_string,"sgebklhl")
						}
				}
			case "Philem","Philemon","門","腓利","腓利門","腓利門書","Phm","phm","Послание к Филимону","Phi-lê-môn","ピレモンヘの手紙","ピレモン","빌레몬서":
				bible_short_name = "門"
				switch chap_string {
					case "":
						print_string = "腓利門書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("門","腓利門書", chap_string, "1","sgebklhl")
							default:
								print_string = Bible_print_string("門","腓利門書", chap_string, sec_string,"sgebklhl")
						}
				}
			case "Heb","Hebrews","來","希伯來","希伯來書","heb","Послание к Евреям","Hê-bơ-rơ","ヘブル人への手紙","ヘブル","히브리서":
				bible_short_name = "來"
				switch chap_string {
					case "":
						print_string = "希伯來書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("來","希伯來書", chap_string, "1","sgebklhl")
							default:
								print_string = Bible_print_string("來","希伯來書", chap_string, sec_string,"sgebklhl")
						}
				}
			case "James","雅","雅各","雅各書","Jas","jas","Послание Иакова","Gia-cơ","ヤコブの手紙","ヤコブ","야고보서":
				bible_short_name = "雅"
				switch chap_string {
					case "":
						print_string = "雅各書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("雅","雅各書", chap_string, "1","sgebklhl")
							default:
								print_string = Bible_print_string("雅","雅各書", chap_string, sec_string,"sgebklhl")
						}
				}
			case "1 Pet","First Peter","彼前","彼得前","彼得前書","1Pe","1pe","Первое послание Петра","1 Phi-e-rơ","ペテロの第一の手紙","ペテロ一","베드로전서":
				bible_short_name = "彼前"
				switch chap_string {
					case "":
						print_string = "彼得前書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("彼前","彼得前書", chap_string, "1","sgebklhl")
							default:
								print_string = Bible_print_string("彼前","彼得前書", chap_string, sec_string,"sgebklhl")
						}
				}
			case "2 Pet","Second Peter","彼後","彼得後","彼得後書","2Pe","2pe","Второе послание Петра","2 Phi-e-rơ","ペテロの第二の手紙","ペテロ","베드로후서":
				bible_short_name = "彼後"
				switch chap_string {
					case "":
						print_string = "彼得後書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("彼後","彼得後書", chap_string, "1","sgebklhl")
							default:
								print_string = Bible_print_string("彼後","彼得後書", chap_string, sec_string,"sgebklhl")
						}
				}
			case "1 John","First John","約一","約翰一書","約翰1","約翰1書","1Jo","1jo","Первое послание Иоанна","1 Giăng","ヨハネの第一の手紙","ヨハネ一","요한일서":
				bible_short_name = "約一"
				switch chap_string {
					case "":
						print_string = "約翰一書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("約一","約翰一書", chap_string, "1","sgebklhl")
							default:
								print_string = Bible_print_string("約一","約翰一書", chap_string, sec_string,"sgebklhl")
						}
				}
			case "2 John","second John","約二","約翰二書","約翰2","約翰2書","2Jo","Второе послание Иоанна","2 Giăng","ヨハネの第二の手紙","ヨハネ二","요한2서":
				bible_short_name = "約二"
				switch chap_string {
					case "":
						print_string = "約翰二書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("約二","約翰二書", chap_string, "1","sgebklhl")
							default:
								print_string = Bible_print_string("約二","約翰二書", chap_string, sec_string,"sgebklhl")
						}
				}
			case "3 John","Third John","約三","約翰三書","約翰3","約翰3書","3Jo","3jo","Третье послание Иоанна","3 Giăng","ヨハネの第三の手紙","ヨハネ三","요한3서":
				bible_short_name = "約三"
				switch chap_string {
					case "":
						print_string = "約翰三書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("約三","約翰三書", chap_string, "1","sgebklhl")
							default:
								print_string = Bible_print_string("約三","約翰三書", chap_string, sec_string,"sgebklhl")
						}
				}
			case "Jude","猶","猶大","猶大書","jude","Послание Иуды","Giu-đe","ユダの手紙","ユダ","유다서":
				bible_short_name = "猶"
				switch chap_string {
					case "":
						print_string = "猶大書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("猶","猶大書", chap_string, "1","sgebklhl")
							default:
								print_string = Bible_print_string("猶","猶大書", chap_string, sec_string,"sgebklhl")
						}
				}
			case "Rev","Revelation","啟","啟示","啟示錄","Re","re","ｒｅ","Ｒｅ","rev","Откровение ап. Иоанна Богослова (Апокалипсис)","Khải-huyền","ヨハネの黙示録","黙示録","요한계시록":
				bible_short_name = "啟"
				switch chap_string {
					case "":
						print_string = "啟示錄"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("啟","啟示錄", chap_string, "1","sgebklhl")
							default:
								print_string = Bible_print_string("啟","啟示錄", chap_string, sec_string,"sgebklhl")
						}
				}
			default:
				print_string = "聖經"
				//print_string = "你是要找 " +  reg.ReplaceAllString(text, "$3") + " 對嗎？\n對不起，我還沒學呢...\n"
		}
	case "文言文聖經","深文理和合本":
		log.Print(reg.ReplaceAllString(text, "$3"))
		switch reg.ReplaceAllString(text, "$3") {
			case "Gen","Genesis","創","創世","創世紀","創世記","Ge","ge","gen","창세기","Sáng-thế Ký","Бытие":
				bible_short_name = "創"
				switch chap_string {
					case "":
						print_string = "創世記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("創","創世記", chap_string, "1","wlunv")
							default:
								print_string = Bible_print_string("創","創世記", chap_string, sec_string,"wlunv")
						}
				}
			case "ex","Ex","Exodus","埃及","出","出埃及","出埃及記","출애굽기","エジプト","出エジプト","出エジプト記","Xuất Ê-díp-tô Ký","Исход":
				bible_short_name = "出"
				switch chap_string {
					case "":
						print_string = "出埃及記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("出","出埃及記", chap_string, "1","wlunv")
							default:
								print_string = Bible_print_string("出","出埃及記", chap_string, sec_string,"wlunv")
						}
				}
			case "Lev","Leviticus","利","利未","利未記","Le","le","Левит","Lê-vi Ký","レビ記","レビ","레위기":
				bible_short_name = "利"
				switch chap_string {
					case "":
						print_string = "利未記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("利","利未記", chap_string, "1","wlunv")
							default:
								print_string = Bible_print_string("利","利未記", chap_string, sec_string,"wlunv")
						}
				}
			case "Num","Numbers","民","民數","民數記","Nu","nu","민수기","民数","民数記","Dân-số Ký","Числа":
				bible_short_name = "民"
				switch chap_string {
					case "":
						print_string = "民數記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("民","民數記", chap_string, "1","wlunv")
							default:
								print_string = Bible_print_string("民","民數記", chap_string, sec_string,"wlunv")
						}
				}
			case "Deut","Deuteronomy","申","申命","申命記","De","de","신명기","Phục-truyền Luật-lệ Ký","Второзаконие":
				bible_short_name = "申"
				switch chap_string {
					case "":
						print_string = "申命記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("申","申命記", chap_string, "1","wlunv")
							default:
								print_string = Bible_print_string("申","申命記", chap_string, sec_string,"wlunv")
						}
				}
			case "Josh","Joshua","約書亞","約書亞記","Jos","jos","여호수아","ヨシュア記","ヨシュア","Giô-suê","Книга Иисуса Навина":
				bible_short_name = "書"
				switch chap_string {
					case "":
						print_string = "約書亞記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("書","約書亞記", chap_string, "1","wlunv")
							default:
								print_string = Bible_print_string("書","約書亞記", chap_string, sec_string,"wlunv")
						}
				}
			case "Judg","Judges","士","士師","士師記","Jud","jud","jdg","Jdg","Книга Судей израилевых","Các Quan Xét","사사기":
				bible_short_name = "士"
				switch chap_string {
					case "":
						print_string = "士師記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("士","士師記", chap_string, "1","wlunv")
							default:
								print_string = Bible_print_string("士","士師記", chap_string, sec_string,"wlunv")
						}
				}
			case "Ruth","路得","路得記","Ru","ru","Rut","rut","룻기","ルツ","ルツ記","Ru-tơ","Книга Руфи":
				bible_short_name = "得"
				switch chap_string {
					case "":
						print_string = "路得記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("得","路得記", chap_string, "1","wlunv")
							default:
								print_string = Bible_print_string("得","路得記", chap_string, sec_string,"wlunv")
						}
				}
			case "1 Sam","First Samuel","撒上","撒母耳記上","1Sa","1sa","サムエル記上","サムエル上","サム上","사무엘상","1 Sa-mu-ên","Первая книга Царств":
				bible_short_name = "撒上"
				switch chap_string {
					case "":
						print_string = "撒母耳記上"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("撒上","撒母耳記上", chap_string, "1","wlunv")
							default:
								print_string = Bible_print_string("撒上","撒母耳記上", chap_string, sec_string,"wlunv")
						}
				}
			case "2 Sam","Second Samuel","撒下","撒母耳記下","2Sa","2sa","사무엘하","サムエル記下","サムエル下","サム下","2 Sa-mu-ên","Вторая книга Царств":
				bible_short_name = "撒下"
				switch chap_string {
					case "":
						print_string = "撒母耳記下"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("撒下","撒母耳記下", chap_string, "1","wlunv")
							default:
								print_string = Bible_print_string("撒下","撒母耳記下", chap_string, sec_string,"wlunv")
						}
				}
			case "1 Kin","First Kings","王上","列王上","列王紀上","列王記上","1Ki","1ki","열왕기상","Третья книга Царств","1 Các Vua":
				bible_short_name = "王上"
				switch chap_string {
					case "":
						print_string = "列王紀上"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("王上","列王紀上", chap_string, "1","wlunv")
							default:
								print_string = Bible_print_string("王上","列王紀上", chap_string, sec_string,"wlunv")
						}
				}
			case "2 Kin","Second Kings","王下","列王下","列王記下","列王紀下","2Ki","2ki","열왕기하","2 Các Vua","Четвертая книга Царств":
				bible_short_name = "王下"
				switch chap_string {
					case "":
						print_string = "列王紀下"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("王下","列王紀下", chap_string, "1","wlunv")
							default:
								print_string = Bible_print_string("王下","列王紀下", chap_string, sec_string,"wlunv")
						}
				}
			case "1 Chr","First Chronicles","歷上","代上","歷代志上","歷代上","1Ch","1ch","Первая книга Паралипоменон","1 Sử-ký","歴上","歴代上","歴代志上","역대상":
				bible_short_name = "代上"
				switch chap_string {
					case "":
						print_string = "歷代志上"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("代上","歷代志上", chap_string, "1","wlunv")
							default:
								print_string = Bible_print_string("代上","歷代志上", chap_string, sec_string,"wlunv")
						}
				}
			case "2 Chr","Second Chronicles","代下","歷下","歷代下","歷代志下","2Ch","2ch","역대하","歴代志下","歴代下","歴下","2 Sử-ký","Вторая книга Паралипоменон":
				bible_short_name = "代下"
				switch chap_string {
					case "":
						print_string = "歷代志下"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("代下","歷代志下", chap_string, "1","wlunv")
							default:
								print_string = Bible_print_string("代下","歷代志下", chap_string, sec_string,"wlunv")
						}
				}
			case "Ezra","拉","以斯拉","以斯拉記","Ezr","ezr","Первая книга Ездры","Ê-xơ-ra","エズラ","エズラ記","에스라":
				bible_short_name = "拉"
				switch chap_string {
					case "":
						print_string = "以斯拉記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("拉","以斯拉記", chap_string, "1","wlunv")
							default:
								print_string = Bible_print_string("拉","以斯拉記", chap_string, sec_string,"wlunv")
						}
				}
			case "Neh","Nehemiah","尼","尼希米","尼希米記","Ne","ne","느헤미야","ネヘミヤ書","ネヘミヤ","Nê-hê-mi","Книга Неемии":
				bible_short_name = "尼"
				switch chap_string {
					case "":
						print_string = "尼希米記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("尼","尼希米記", chap_string, "1","wlunv")
							default:
								print_string = Bible_print_string("尼","尼希米記", chap_string, sec_string,"wlunv")
						}
				}
			case "Esth","Esther","斯","以斯帖","以斯帖記","Es","est","Есфирь","Ê-xơ-tê","エステル","エステル記","에스더":
				bible_short_name = "斯"
				switch chap_string {
					case "":
						print_string = "以斯帖記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("斯","以斯帖記", chap_string, "1","wlunv")
							default:
								print_string = Bible_print_string("斯","以斯帖記", chap_string, sec_string,"wlunv")
						}
				}
			case "Job","job","伯","約伯","約伯記","Книга Иова","Gióp","ヨブ","ヨブ記","욥기":
				bible_short_name = "伯"
				switch chap_string {
					case "":
						print_string = "約伯記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("伯","約伯記", chap_string, "1","wlunv")
							default:
								print_string = Bible_print_string("伯","約伯記", chap_string, sec_string,"wlunv")
						}
				}
			case "Ps","Psalms","詩","詩篇","ps","시편","Thi-thiên","Псалтирь":
				bible_short_name = "詩"
				switch chap_string {
					case "":
						print_string = "詩篇"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("詩","詩篇", chap_string, "1","wlunv")
							default:
								print_string = Bible_print_string("詩","詩篇", chap_string, sec_string,"wlunv")
						}
				}
			case "Prov","Proverbs","箴","箴言","Pr","pr","Притчи Соломона","Châm-ngôn","잠언":
				bible_short_name = "箴"
				switch chap_string {
					case "":
						print_string = "箴言"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("箴","箴言", chap_string, "1","wlunv")
							default:
								print_string = Bible_print_string("箴","箴言", chap_string, sec_string,"wlunv")
						}
				}
			case "Eccl","Ecclesiastes","傳","傳道","傳道書","Ec","ec","Книга Екклезиаста","или Проповедника","Truyền-đạo","伝道の書","伝道","伝","伝道書","전도서":
				bible_short_name = "傳"
				switch chap_string {
					case "":
						print_string = "傳道書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("傳","傳道書", chap_string, "1","wlunv")
							default:
								print_string = Bible_print_string("傳","傳道書", chap_string, sec_string,"wlunv")
						}
				}
			case "Song","Song of Solomon","歌","雅歌","So","so","sng","Sng","Песнь песней Соломона","Nhã-ca","아가":
				bible_short_name = "歌"
				switch chap_string {
					case "":
						print_string = "雅歌"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("歌","雅歌", chap_string, "1","wlunv")
							default:
								print_string = Bible_print_string("歌","雅歌", chap_string, sec_string,"wlunv")
						}
				}
			case "Is","Isaiah","賽","以賽","以賽亞","以賽亞書","Isa","isa","Книга пророка Исаии","Ê-sai","イザヤ書","イザヤ","이사야":
				bible_short_name = "賽"
				switch chap_string {
					case "":
						print_string = "以賽亞書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("賽","以賽亞書", chap_string, "1","wlunv")
							default:
								print_string = Bible_print_string("賽","以賽亞書", chap_string, sec_string,"wlunv")
						}
				}
			case "Jer","Jeremiah","耶","耶利米","耶利米書","jer","예레미야","エレミヤ","エレミヤ書","Giê-rê-mi","Книга пророка Иеремии":
				bible_short_name = "耶"
				switch chap_string {
					case "":
						print_string = "耶利米書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("耶","耶利米書", chap_string, "1","wlunv")
							default:
								print_string = Bible_print_string("耶","耶利米書", chap_string, sec_string,"wlunv")
						}
				}
			case "Lam","Lamentations","哀","哀歌","耶利米哀歌","La","lam","예레미야애가","Ca-thương","Плач Иеремии":
				bible_short_name = "哀"
				switch chap_string {
					case "":
						print_string = "耶利米哀歌"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("哀","耶利米哀歌", chap_string, "1","wlunv")
							default:
								print_string = Bible_print_string("哀","耶利米哀歌", chap_string, sec_string,"wlunv")
						}
				}
			case "Ezek","Ezekiel","結","以西結","以西結書","Eze","eze","에스겔","エゼキエル書","エゼキエル","Ê-xê-chi-ên","Книга пророка Иезекииля":
				bible_short_name = "結"
				switch chap_string {
					case "":
						print_string = "以西結書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("結","以西結書", chap_string, "1","wlunv")
							default:
								print_string = Bible_print_string("結","以西結書", chap_string, sec_string,"wlunv")
						}
				}
			case "Dan","Daniel","但","但以理","但以理書","Da","da","Книга пророка Даниила","Đa-ni-ên","ダニエル書","ダニエル","다니엘":
				bible_short_name = "但"
				switch chap_string {
					case "":
						print_string = "但以理書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("但","但以理書", chap_string, "1","wlunv")
							default:
								print_string = Bible_print_string("但","但以理書", chap_string, sec_string,"wlunv")
						}
				}
			case "Hos","Hosea","何","何西","何西阿","何西阿書","Ho","ho","Книга пророка Осии","Ô-sê","ホセア書","ホセア","호세아":
				bible_short_name = "何"
				switch chap_string {
					case "":
						print_string = "何西阿書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("何","何西阿書", chap_string, "1","wlunv")
							default:
								print_string = Bible_print_string("何","何西阿書", chap_string, sec_string,"wlunv")
						}
				}
			case "Joel","珥","約珥","約珥書","Joe","joe","Книга пророка Иоиля","Giô-ên","ヨエル書","ヨエル","요엘":
				bible_short_name = "珥"
				switch chap_string {
					case "":
						print_string = "約珥書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("珥","約珥書", chap_string, "1","wlunv")
							default:
								print_string = Bible_print_string("珥","約珥書", chap_string, sec_string,"wlunv")
						}
				}
			case "Amos","摩","阿摩司書","Am","am","Книга пророка Амоса","A-mốt","アモス書","アモス","아모스":
				bible_short_name = "摩"
				switch chap_string {
					case "":
						print_string = "阿摩司書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("摩","阿摩司書", chap_string, "1","wlunv")
							default:
								print_string = Bible_print_string("摩","阿摩司書", chap_string, sec_string,"wlunv")
						}
				}
			case "Obad","Obadiah","俄","俄巴底亞","俄巴底亞書","Ob","ob","오바댜","オバデヤ書","オバデヤ","Áp-đia","Книга пророка Авдия":
				bible_short_name = "俄"
				switch chap_string {
					case "":
						print_string = "俄巴底亞書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("俄","俄巴底亞書", chap_string, "1","wlunv")
							default:
								print_string = Bible_print_string("俄","俄巴底亞書", chap_string, sec_string,"wlunv")
						}
				}
			case "Jon","Jonah","拿","約拿","約拿書","jon","요나","ヨナ書","ヨナ","Giô-na","Книга пророка Ионы":
				bible_short_name = "拿"
				switch chap_string {
					case "":
						print_string = "約拿書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("拿","約拿書", chap_string, "1","wlunv")
							default:
								print_string = Bible_print_string("拿","約拿書", chap_string, sec_string,"wlunv")
						}
				}
			case "Micah","彌","彌迦","彌迦書","Mic","mic","Книга пророка Михея","Mi-chê","ミカ書","ミカ","미가":
				bible_short_name = "彌"
				switch chap_string {
					case "":
						print_string = "彌迦書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("彌","彌迦書", chap_string, "1","wlunv")
							default:
								print_string = Bible_print_string("彌","彌迦書", chap_string, sec_string,"wlunv")
						}
				}
			case "Nah","Nahum","鴻","那鴻","那鴻書","Na","na","Книга пророка Наума","Na-hum","ナホム書","ナホム","나훔":
				bible_short_name = "鴻"
				switch chap_string {
					case "":
						print_string = "那鴻書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("鴻","那鴻書", chap_string, "1","wlunv")
							default:
								print_string = Bible_print_string("鴻","那鴻書", chap_string, sec_string,"wlunv")
						}
				}
			case "Habakkuk","哈","哈巴","哈巴谷","哈巴谷書","Hab","hab","Книга пророка Аввакума","Ha-ba-cúc","ハバクク書","ハバクク","ハバ","クク","ハバ書","하박국":
				bible_short_name = "哈"
				switch chap_string {
					case "":
						print_string = "哈巴谷書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("哈","哈巴谷書", chap_string, "1","wlunv")
							default:
								print_string = Bible_print_string("哈","哈巴谷書", chap_string, sec_string,"wlunv")
						}
				}
			case "Zeph","Zephaniah","番","西番雅","西番雅書","Zep","zep","스바냐","ゼパニヤ書","ゼパニヤ","Sô-phô-ni","Книга пророка Софонии":
				bible_short_name = "番"
				switch chap_string {
					case "":
						print_string = "西番雅書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("番","西番雅書", chap_string, "1","wlunv")
							default:
								print_string = Bible_print_string("番","西番雅書", chap_string, sec_string,"wlunv")
						}
				}
			case "Haggai","該","哈該","哈該書","Hag","hag","학개","ハガイ書","ハガイ","A-ghê","Книга пророка Аггея":
				bible_short_name = "該"
				switch chap_string {
					case "":
						print_string = "哈該書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("該","哈該書", chap_string, "1","wlunv")
							default:
								print_string = Bible_print_string("該","哈該書", chap_string, sec_string,"wlunv")
						}
				}
			case "Zech","Zechariah","亞","撒迦利亞","撒迦利亞書","Zec","zec","Книга пророка Захарии","Xa-cha-ri","스가랴","ゼカリヤ書","ゼカリヤ":
				bible_short_name = "亞"
				switch chap_string {
					case "":
						print_string = "撒迦利亞書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("亞","撒迦利亞書", chap_string, "1","wlunv")
							default:
								print_string = Bible_print_string("亞","撒迦利亞書", chap_string, sec_string,"wlunv")
						}
				}
			case "Malachi","瑪","瑪拉","瑪拉基","瑪拉基書","Mal","mal","말라기","マラキ書","マラキ","Ma-la-chi","Книга пророка Малахии":
				bible_short_name = "瑪"
				switch chap_string {
					case "":
						print_string = "瑪拉基書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("瑪","瑪拉基書", chap_string, "1","wlunv")
							default:
								print_string = Bible_print_string("瑪","瑪拉基書", chap_string, sec_string,"wlunv")
						}
				}
			case "Matt","Matthew","太","馬太","馬太福音","Mt","mt","마태복음","マタイによる福音書","マタイ","マタイによる","Ma-thi-ơ","От Матфея святое благовествование":
				bible_short_name = "太"
				switch chap_string {
					case "":
						print_string = "馬太福音"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("太","馬太福音", chap_string, "1","wlunv")
							default:
								print_string = Bible_print_string("太","馬太福音", chap_string, sec_string,"wlunv")
						}
				}
			case "Mark","可","馬可","馬可福音","Mr","mr","マルコによる福音書","マルコ","マルコによる","마가복음","Mác","От Марка святое благовествование":
				bible_short_name = "可"
				switch chap_string {
					case "":
						print_string = "馬可福音"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("可","馬可福音", chap_string, "1","wlunv")
							default:
								print_string = Bible_print_string("可","馬可福音", chap_string, sec_string,"wlunv")
						}
				}
			case "Luke","路","路加","路加福音","Lu","lu","От Луки святое благовествование","Lu-ca","ルカによる福音書","ルカ","ルカによる","누가복음":
				bible_short_name = "路"
				switch chap_string {
					case "":
						print_string = "路加福音"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("路","路加福音", chap_string, "1","wlunv")
							default:
								print_string = Bible_print_string("路","路加福音", chap_string, sec_string,"wlunv")
						}
				}
			case "John","約","約翰","約翰福音","Joh","joh","От Иоанна святое благовествование","Giăng","ヨハネによる福音書","ヨハネ","ヨハネによる","요한복음":
				bible_short_name = "約"
				switch chap_string {
					case "":
						print_string = "約翰福音"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("約","約翰福音", chap_string, "1","wlunv")
							default:
								print_string = Bible_print_string("約","約翰福音", chap_string, sec_string,"wlunv")
						}
				}
			case "Acts","徒","使徒","使徒行傳","Ac","ac","Деяния святых апостолов","Công-vụ Các Sứ-đồ","使徒行伝","사도행전":
				bible_short_name = "徒"
				switch chap_string {
					case "":
						print_string = "使徒行傳"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("徒","使徒行傳", chap_string, "1","wlunv")
							default:
								print_string = Bible_print_string("徒","使徒行傳", chap_string, sec_string,"wlunv")
						}
				}
			case "Rom","Romans","羅","羅馬","羅馬書","Ro","ro","Послание к Римлянам","Rô-ma","ローマ","ローマ人への手紙","로마서":
				bible_short_name = "羅"
				switch chap_string {
					case "":
						print_string = "羅馬書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("羅","羅馬書", chap_string, "1","wlunv")
							default:
								print_string = Bible_print_string("羅","羅馬書", chap_string, sec_string,"wlunv")
						}
				}
			case "1 Cor","First Corinthians","林前","哥林多前","哥林多前書","1Co","1co","Первое послание к Коринфянам","1 Cô-rinh-tô","コリント人への第一の手紙","コリント一","コリント人への第一","고린도전서":
				bible_short_name = "林前"
				switch chap_string {
					case "":
						print_string = "哥林多前書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("林前","哥林多前書", chap_string, "1","wlunv")
							default:
								print_string = Bible_print_string("林前","哥林多前書", chap_string, sec_string,"wlunv")
						}
				}
			case "2 Cor","Second Corinthians","林後","哥林多後","哥林多後書","2Co","2co","Второе послание к Коринфянам","2 Cô-rinh-tô","コリント人への第二の手紙","コリント二","コリント人への第二の","고린도후서":
				bible_short_name = "林後"
				switch chap_string {
					case "":
						print_string = "哥林多後書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("林後","哥林多後書", chap_string, "1","wlunv")
							default:
								print_string = Bible_print_string("林後","哥林多後書", chap_string, sec_string,"wlunv")
						}
				}
			case "Gal","Galatians","加","加拉太","加拉太書","Ga","ga","Послание к Галатам","Ga-la-ti","ガラテヤ","ガラテヤ人への手紙","갈라디아서":
				bible_short_name = "加"
				switch chap_string {
					case "":
						print_string = "加拉太書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("加","加拉太書", chap_string, "1","wlunv")
							default:
								print_string = Bible_print_string("加","加拉太書", chap_string, sec_string,"wlunv")
						}
				}
			case "Ephesians","弗","以弗所","以弗所書","Eph","eph","Послание к Ефесянам","Ê-phê-sô","エペソ人への手紙","エペソ","エペソ人","エペソ人の手紙","에베소서":
				bible_short_name = "弗"
				switch chap_string {
					case "":
						print_string = "以弗所書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("弗","以弗所書", chap_string, "1","wlunv")
							default:
								print_string = Bible_print_string("弗","以弗所書", chap_string, sec_string,"wlunv")
						}
				}
			case "Phil","Philippians","腓","腓立","腓立比","腓立比書","Php","php","빌립보서","ピリピ","ピリピ人.ピリピ人への手紙","Послание к Филиппийцам","Phi-líp":
				bible_short_name = "腓"
				switch chap_string {
					case "":
						print_string = "腓立比書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("腓","腓立比書", chap_string, "1","wlunv")
							default:
								print_string = Bible_print_string("腓","腓立比書", chap_string, sec_string,"wlunv")
						}
				}
			case "Col","col","Colossians","西","歌羅西","歌羅","歌羅西書","Послание к Колоссянам","Cô-lô-se","コロサイ人への手紙","コロサイ","コロ","골로새서":
				bible_short_name = "西"
				switch chap_string {
					case "":
						print_string = "歌羅西書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("西","歌羅西書", chap_string, "1","wlunv")
							default:
								print_string = Bible_print_string("西","歌羅西書", chap_string, sec_string,"wlunv")
						}
				}
			case "1 Thess","First Thessalonians","帖前","帖撒羅尼迦前","帖撒羅尼迦前書","1Th","1th","데살로니가전서","テサロニケ人への第一の手紙","テサ一","テサロニケ一","1 Tê-sa-lô-ni-ca","Первое послание к Фессалоникийцам (Солунянам)":
				bible_short_name = "帖前"
				switch chap_string {
					case "":
						print_string = "帖撒羅尼迦前書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("帖前","帖撒羅尼迦前書", chap_string, "1","wlunv")
							default:
								print_string = Bible_print_string("帖前","帖撒羅尼迦前書", chap_string, sec_string,"wlunv")
						}
				}
			case "2 Thess","Second Thessalonians","帖後","帖撒羅尼迦後","帖撒羅尼迦後書","2Th","2th","데살로니가후서","テサロニケ人への第二の手紙","テサ二","テサロニケ二","2 Tê-sa-lô-ni-ca","Второе послание к Фессалоникийцам (Солунянам)":
				bible_short_name = "帖後"
				switch chap_string {
					case "":
						print_string = "帖撒羅尼迦後書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("帖後","帖撒羅尼迦後書", chap_string, "1","wlunv")
							default:
								print_string = Bible_print_string("帖後","帖撒羅尼迦後書", chap_string, sec_string,"wlunv")
						}
				}
			case "1 Tim","First Timothy","提前","提摩太前","提摩太前書","1Ti","1ti","Первое послание к Тимофею","1 Ti-mô-thê","テモテヘの第一の手紙","テモテ一","디모데전서":
				bible_short_name = "提前"
				switch chap_string {
					case "":
						print_string = "提摩太前書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("提前","提摩太前書", chap_string, "1","wlunv")
							default:
								print_string = Bible_print_string("提前","提摩太前書", chap_string, sec_string,"wlunv")
						}
				}
			case "2 Tim","Second Timothy","提後","提摩太後","提摩太後書","2Ti","2ti","Второе послание к Тимофею","2 Ti-mô-thê","テモテヘの第二の手紙","テモテ二","디모데후서":
				bible_short_name = "提後"
				switch chap_string {
					case "":
						print_string = "提摩太後書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("提後","提摩太後書", chap_string, "1","wlunv")
							default:
								print_string = Bible_print_string("提後","提摩太後書", chap_string, sec_string,"wlunv")
						}
				}
			case "Titus","多","提多","提多書","Tit","tit","Послание к Титу","Tít","テトスヘの手紙","テトス","디도서":
				bible_short_name = "多"
				switch chap_string {
					case "":
						print_string = "提多書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("多","提多書", chap_string, "1","wlunv")
							default:
								print_string = Bible_print_string("多","提多書", chap_string, sec_string,"wlunv")
						}
				}
			case "Philem","Philemon","門","腓利","腓利門","腓利門書","Phm","phm","Послание к Филимону","Phi-lê-môn","ピレモンヘの手紙","ピレモン","빌레몬서":
				bible_short_name = "門"
				switch chap_string {
					case "":
						print_string = "腓利門書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("門","腓利門書", chap_string, "1","wlunv")
							default:
								print_string = Bible_print_string("門","腓利門書", chap_string, sec_string,"wlunv")
						}
				}
			case "Heb","Hebrews","來","希伯來","希伯來書","heb","Послание к Евреям","Hê-bơ-rơ","ヘブル人への手紙","ヘブル","히브리서":
				bible_short_name = "來"
				switch chap_string {
					case "":
						print_string = "希伯來書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("來","希伯來書", chap_string, "1","wlunv")
							default:
								print_string = Bible_print_string("來","希伯來書", chap_string, sec_string,"wlunv")
						}
				}
			case "James","雅","雅各","雅各書","Jas","jas","Послание Иакова","Gia-cơ","ヤコブの手紙","ヤコブ","야고보서":
				bible_short_name = "雅"
				switch chap_string {
					case "":
						print_string = "雅各書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("雅","雅各書", chap_string, "1","wlunv")
							default:
								print_string = Bible_print_string("雅","雅各書", chap_string, sec_string,"wlunv")
						}
				}
			case "1 Pet","First Peter","彼前","彼得前","彼得前書","1Pe","1pe","Первое послание Петра","1 Phi-e-rơ","ペテロの第一の手紙","ペテロ一","베드로전서":
				bible_short_name = "彼前"
				switch chap_string {
					case "":
						print_string = "彼得前書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("彼前","彼得前書", chap_string, "1","wlunv")
							default:
								print_string = Bible_print_string("彼前","彼得前書", chap_string, sec_string,"wlunv")
						}
				}
			case "2 Pet","Second Peter","彼後","彼得後","彼得後書","2Pe","2pe","Второе послание Петра","2 Phi-e-rơ","ペテロの第二の手紙","ペテロ","베드로후서":
				bible_short_name = "彼後"
				switch chap_string {
					case "":
						print_string = "彼得後書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("彼後","彼得後書", chap_string, "1","wlunv")
							default:
								print_string = Bible_print_string("彼後","彼得後書", chap_string, sec_string,"wlunv")
						}
				}
			case "1 John","First John","約一","約翰一書","約翰1","約翰1書","1Jo","1jo","Первое послание Иоанна","1 Giăng","ヨハネの第一の手紙","ヨハネ一","요한일서":
				bible_short_name = "約一"
				switch chap_string {
					case "":
						print_string = "約翰一書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("約一","約翰一書", chap_string, "1","wlunv")
							default:
								print_string = Bible_print_string("約一","約翰一書", chap_string, sec_string,"wlunv")
						}
				}
			case "2 John","second John","約二","約翰二書","約翰2","約翰2書","2Jo","Второе послание Иоанна","2 Giăng","ヨハネの第二の手紙","ヨハネ二","요한2서":
				bible_short_name = "約二"
				switch chap_string {
					case "":
						print_string = "約翰二書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("約二","約翰二書", chap_string, "1","wlunv")
							default:
								print_string = Bible_print_string("約二","約翰二書", chap_string, sec_string,"wlunv")
						}
				}
			case "3 John","Third John","約三","約翰三書","約翰3","約翰3書","3Jo","3jo","Третье послание Иоанна","3 Giăng","ヨハネの第三の手紙","ヨハネ三","요한3서":
				bible_short_name = "約三"
				switch chap_string {
					case "":
						print_string = "約翰三書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("約三","約翰三書", chap_string, "1","wlunv")
							default:
								print_string = Bible_print_string("約三","約翰三書", chap_string, sec_string,"wlunv")
						}
				}
			case "Jude","猶","猶大","猶大書","jude","Послание Иуды","Giu-đe","ユダの手紙","ユダ","유다서":
				bible_short_name = "猶"
				switch chap_string {
					case "":
						print_string = "猶大書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("猶","猶大書", chap_string, "1","wlunv")
							default:
								print_string = Bible_print_string("猶","猶大書", chap_string, sec_string,"wlunv")
						}
				}
			case "Rev","Revelation","啟","啟示","啟示錄","Re","re","ｒｅ","Ｒｅ","rev","Откровение ап. Иоанна Богослова (Апокалипсис)","Khải-huyền","ヨハネの黙示録","黙示録","요한계시록":
				bible_short_name = "啟"
				switch chap_string {
					case "":
						print_string = "啟示錄"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("啟","啟示錄", chap_string, "1","wlunv")
							default:
								print_string = Bible_print_string("啟","啟示錄", chap_string, sec_string,"wlunv")
						}
				}
			default:
				print_string = "聖經"
				//print_string = "你是要找 " +  reg.ReplaceAllString(text, "$3") + " 對嗎？\n對不起，我還沒學呢...\n"
		}
	case "中文聖經譯本修訂版","tcv","TCV","Ｔｃｖ","ＴＣＶ":
		log.Print(reg.ReplaceAllString(text, "$3"))
		switch reg.ReplaceAllString(text, "$3") {
			case "Gen","Genesis","創","創世","創世紀","創世記","Ge","ge","gen","창세기","Sáng-thế Ký","Бытие":
				bible_short_name = "創"
				switch chap_string {
					case "":
						print_string = "創世記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("創","創世記", chap_string, "1","tcv")
							default:
								print_string = Bible_print_string("創","創世記", chap_string, sec_string,"tcv")
						}
				}
			case "ex","Ex","Exodus","埃及","出","出埃及","出埃及記","출애굽기","エジプト","出エジプト","出エジプト記","Xuất Ê-díp-tô Ký","Исход":
				bible_short_name = "出"
				switch chap_string {
					case "":
						print_string = "出埃及記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("出","出埃及記", chap_string, "1","tcv")
							default:
								print_string = Bible_print_string("出","出埃及記", chap_string, sec_string,"tcv")
						}
				}
			case "Lev","Leviticus","利","利未","利未記","Le","le","Левит","Lê-vi Ký","レビ記","レビ","레위기":
				bible_short_name = "利"
				switch chap_string {
					case "":
						print_string = "利未記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("利","利未記", chap_string, "1","tcv")
							default:
								print_string = Bible_print_string("利","利未記", chap_string, sec_string,"tcv")
						}
				}
			case "Num","Numbers","民","民數","民數記","Nu","nu","민수기","民数","民数記","Dân-số Ký","Числа":
				bible_short_name = "民"
				switch chap_string {
					case "":
						print_string = "民數記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("民","民數記", chap_string, "1","tcv")
							default:
								print_string = Bible_print_string("民","民數記", chap_string, sec_string,"tcv")
						}
				}
			case "Deut","Deuteronomy","申","申命","申命記","De","de","신명기","Phục-truyền Luật-lệ Ký","Второзаконие":
				bible_short_name = "申"
				switch chap_string {
					case "":
						print_string = "申命記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("申","申命記", chap_string, "1","tcv")
							default:
								print_string = Bible_print_string("申","申命記", chap_string, sec_string,"tcv")
						}
				}
			case "Josh","Joshua","約書亞","約書亞記","Jos","jos","여호수아","ヨシュア記","ヨシュア","Giô-suê","Книга Иисуса Навина":
				bible_short_name = "書"
				switch chap_string {
					case "":
						print_string = "約書亞記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("書","約書亞記", chap_string, "1","tcv")
							default:
								print_string = Bible_print_string("書","約書亞記", chap_string, sec_string,"tcv")
						}
				}
			case "Judg","Judges","士","士師","士師記","Jud","jud","jdg","Jdg","Книга Судей израилевых","Các Quan Xét","사사기":
				bible_short_name = "士"
				switch chap_string {
					case "":
						print_string = "士師記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("士","士師記", chap_string, "1","tcv")
							default:
								print_string = Bible_print_string("士","士師記", chap_string, sec_string,"tcv")
						}
				}
			case "Ruth","路得","路得記","Ru","ru","Rut","rut","룻기","ルツ","ルツ記","Ru-tơ","Книга Руфи":
				bible_short_name = "得"
				switch chap_string {
					case "":
						print_string = "路得記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("得","路得記", chap_string, "1","tcv")
							default:
								print_string = Bible_print_string("得","路得記", chap_string, sec_string,"tcv")
						}
				}
			case "1 Sam","First Samuel","撒上","撒母耳記上","1Sa","1sa","サムエル記上","サムエル上","サム上","사무엘상","1 Sa-mu-ên","Первая книга Царств":
				bible_short_name = "撒上"
				switch chap_string {
					case "":
						print_string = "撒母耳記上"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("撒上","撒母耳記上", chap_string, "1","tcv")
							default:
								print_string = Bible_print_string("撒上","撒母耳記上", chap_string, sec_string,"tcv")
						}
				}
			case "2 Sam","Second Samuel","撒下","撒母耳記下","2Sa","2sa","사무엘하","サムエル記下","サムエル下","サム下","2 Sa-mu-ên","Вторая книга Царств":
				bible_short_name = "撒下"
				switch chap_string {
					case "":
						print_string = "撒母耳記下"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("撒下","撒母耳記下", chap_string, "1","tcv")
							default:
								print_string = Bible_print_string("撒下","撒母耳記下", chap_string, sec_string,"tcv")
						}
				}
			case "1 Kin","First Kings","王上","列王上","列王紀上","列王記上","1Ki","1ki","열왕기상","Третья книга Царств","1 Các Vua":
				bible_short_name = "王上"
				switch chap_string {
					case "":
						print_string = "列王紀上"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("王上","列王紀上", chap_string, "1","tcv")
							default:
								print_string = Bible_print_string("王上","列王紀上", chap_string, sec_string,"tcv")
						}
				}
			case "2 Kin","Second Kings","王下","列王下","列王記下","列王紀下","2Ki","2ki","열왕기하","2 Các Vua","Четвертая книга Царств":
				bible_short_name = "王下"
				switch chap_string {
					case "":
						print_string = "列王紀下"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("王下","列王紀下", chap_string, "1","tcv")
							default:
								print_string = Bible_print_string("王下","列王紀下", chap_string, sec_string,"tcv")
						}
				}
			case "1 Chr","First Chronicles","歷上","代上","歷代志上","歷代上","1Ch","1ch","Первая книга Паралипоменон","1 Sử-ký","歴上","歴代上","歴代志上","역대상":
				bible_short_name = "代上"
				switch chap_string {
					case "":
						print_string = "歷代志上"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("代上","歷代志上", chap_string, "1","tcv")
							default:
								print_string = Bible_print_string("代上","歷代志上", chap_string, sec_string,"tcv")
						}
				}
			case "2 Chr","Second Chronicles","代下","歷下","歷代下","歷代志下","2Ch","2ch","역대하","歴代志下","歴代下","歴下","2 Sử-ký","Вторая книга Паралипоменон":
				bible_short_name = "代下"
				switch chap_string {
					case "":
						print_string = "歷代志下"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("代下","歷代志下", chap_string, "1","tcv")
							default:
								print_string = Bible_print_string("代下","歷代志下", chap_string, sec_string,"tcv")
						}
				}
			case "Ezra","拉","以斯拉","以斯拉記","Ezr","ezr","Первая книга Ездры","Ê-xơ-ra","エズラ","エズラ記","에스라":
				bible_short_name = "拉"
				switch chap_string {
					case "":
						print_string = "以斯拉記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("拉","以斯拉記", chap_string, "1","tcv")
							default:
								print_string = Bible_print_string("拉","以斯拉記", chap_string, sec_string,"tcv")
						}
				}
			case "Neh","Nehemiah","尼","尼希米","尼希米記","Ne","ne","느헤미야","ネヘミヤ書","ネヘミヤ","Nê-hê-mi","Книга Неемии":
				bible_short_name = "尼"
				switch chap_string {
					case "":
						print_string = "尼希米記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("尼","尼希米記", chap_string, "1","tcv")
							default:
								print_string = Bible_print_string("尼","尼希米記", chap_string, sec_string,"tcv")
						}
				}
			case "Esth","Esther","斯","以斯帖","以斯帖記","Es","est","Есфирь","Ê-xơ-tê","エステル","エステル記","에스더":
				bible_short_name = "斯"
				switch chap_string {
					case "":
						print_string = "以斯帖記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("斯","以斯帖記", chap_string, "1","tcv")
							default:
								print_string = Bible_print_string("斯","以斯帖記", chap_string, sec_string,"tcv")
						}
				}
			case "Job","job","伯","約伯","約伯記","Книга Иова","Gióp","ヨブ","ヨブ記","욥기":
				bible_short_name = "伯"
				switch chap_string {
					case "":
						print_string = "約伯記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("伯","約伯記", chap_string, "1","tcv")
							default:
								print_string = Bible_print_string("伯","約伯記", chap_string, sec_string,"tcv")
						}
				}
			case "Ps","Psalms","詩","詩篇","ps","시편","Thi-thiên","Псалтирь":
				bible_short_name = "詩"
				switch chap_string {
					case "":
						print_string = "詩篇"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("詩","詩篇", chap_string, "1","tcv")
							default:
								print_string = Bible_print_string("詩","詩篇", chap_string, sec_string,"tcv")
						}
				}
			case "Prov","Proverbs","箴","箴言","Pr","pr","Притчи Соломона","Châm-ngôn","잠언":
				bible_short_name = "箴"
				switch chap_string {
					case "":
						print_string = "箴言"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("箴","箴言", chap_string, "1","tcv")
							default:
								print_string = Bible_print_string("箴","箴言", chap_string, sec_string,"tcv")
						}
				}
			case "Eccl","Ecclesiastes","傳","傳道","傳道書","Ec","ec","Книга Екклезиаста","или Проповедника","Truyền-đạo","伝道の書","伝道","伝","伝道書","전도서":
				bible_short_name = "傳"
				switch chap_string {
					case "":
						print_string = "傳道書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("傳","傳道書", chap_string, "1","tcv")
							default:
								print_string = Bible_print_string("傳","傳道書", chap_string, sec_string,"tcv")
						}
				}
			case "Song","Song of Solomon","歌","雅歌","So","so","sng","Sng","Песнь песней Соломона","Nhã-ca","아가":
				bible_short_name = "歌"
				switch chap_string {
					case "":
						print_string = "雅歌"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("歌","雅歌", chap_string, "1","tcv")
							default:
								print_string = Bible_print_string("歌","雅歌", chap_string, sec_string,"tcv")
						}
				}
			case "Is","Isaiah","賽","以賽","以賽亞","以賽亞書","Isa","isa","Книга пророка Исаии","Ê-sai","イザヤ書","イザヤ","이사야":
				bible_short_name = "賽"
				switch chap_string {
					case "":
						print_string = "以賽亞書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("賽","以賽亞書", chap_string, "1","tcv")
							default:
								print_string = Bible_print_string("賽","以賽亞書", chap_string, sec_string,"tcv")
						}
				}
			case "Jer","Jeremiah","耶","耶利米","耶利米書","jer","예레미야","エレミヤ","エレミヤ書","Giê-rê-mi","Книга пророка Иеремии":
				bible_short_name = "耶"
				switch chap_string {
					case "":
						print_string = "耶利米書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("耶","耶利米書", chap_string, "1","tcv")
							default:
								print_string = Bible_print_string("耶","耶利米書", chap_string, sec_string,"tcv")
						}
				}
			case "Lam","Lamentations","哀","哀歌","耶利米哀歌","La","lam","예레미야애가","Ca-thương","Плач Иеремии":
				bible_short_name = "哀"
				switch chap_string {
					case "":
						print_string = "耶利米哀歌"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("哀","耶利米哀歌", chap_string, "1","tcv")
							default:
								print_string = Bible_print_string("哀","耶利米哀歌", chap_string, sec_string,"tcv")
						}
				}
			case "Ezek","Ezekiel","結","以西結","以西結書","Eze","eze","에스겔","エゼキエル書","エゼキエル","Ê-xê-chi-ên","Книга пророка Иезекииля":
				bible_short_name = "結"
				switch chap_string {
					case "":
						print_string = "以西結書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("結","以西結書", chap_string, "1","tcv")
							default:
								print_string = Bible_print_string("結","以西結書", chap_string, sec_string,"tcv")
						}
				}
			case "Dan","Daniel","但","但以理","但以理書","Da","da","Книга пророка Даниила","Đa-ni-ên","ダニエル書","ダニエル","다니엘":
				bible_short_name = "但"
				switch chap_string {
					case "":
						print_string = "但以理書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("但","但以理書", chap_string, "1","tcv")
							default:
								print_string = Bible_print_string("但","但以理書", chap_string, sec_string,"tcv")
						}
				}
			case "Hos","Hosea","何","何西","何西阿","何西阿書","Ho","ho","Книга пророка Осии","Ô-sê","ホセア書","ホセア","호세아":
				bible_short_name = "何"
				switch chap_string {
					case "":
						print_string = "何西阿書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("何","何西阿書", chap_string, "1","tcv")
							default:
								print_string = Bible_print_string("何","何西阿書", chap_string, sec_string,"tcv")
						}
				}
			case "Joel","珥","約珥","約珥書","Joe","joe","Книга пророка Иоиля","Giô-ên","ヨエル書","ヨエル","요엘":
				bible_short_name = "珥"
				switch chap_string {
					case "":
						print_string = "約珥書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("珥","約珥書", chap_string, "1","tcv")
							default:
								print_string = Bible_print_string("珥","約珥書", chap_string, sec_string,"tcv")
						}
				}
			case "Amos","摩","阿摩司書","Am","am","Книга пророка Амоса","A-mốt","アモス書","アモス","아모스":
				bible_short_name = "摩"
				switch chap_string {
					case "":
						print_string = "阿摩司書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("摩","阿摩司書", chap_string, "1","tcv")
							default:
								print_string = Bible_print_string("摩","阿摩司書", chap_string, sec_string,"tcv")
						}
				}
			case "Obad","Obadiah","俄","俄巴底亞","俄巴底亞書","Ob","ob","오바댜","オバデヤ書","オバデヤ","Áp-đia","Книга пророка Авдия":
				bible_short_name = "俄"
				switch chap_string {
					case "":
						print_string = "俄巴底亞書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("俄","俄巴底亞書", chap_string, "1","tcv")
							default:
								print_string = Bible_print_string("俄","俄巴底亞書", chap_string, sec_string,"tcv")
						}
				}
			case "Jon","Jonah","拿","約拿","約拿書","jon","요나","ヨナ書","ヨナ","Giô-na","Книга пророка Ионы":
				bible_short_name = "拿"
				switch chap_string {
					case "":
						print_string = "約拿書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("拿","約拿書", chap_string, "1","tcv")
							default:
								print_string = Bible_print_string("拿","約拿書", chap_string, sec_string,"tcv")
						}
				}
			case "Micah","彌","彌迦","彌迦書","Mic","mic","Книга пророка Михея","Mi-chê","ミカ書","ミカ","미가":
				bible_short_name = "彌"
				switch chap_string {
					case "":
						print_string = "彌迦書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("彌","彌迦書", chap_string, "1","tcv")
							default:
								print_string = Bible_print_string("彌","彌迦書", chap_string, sec_string,"tcv")
						}
				}
			case "Nah","Nahum","鴻","那鴻","那鴻書","Na","na","Книга пророка Наума","Na-hum","ナホム書","ナホム","나훔":
				bible_short_name = "鴻"
				switch chap_string {
					case "":
						print_string = "那鴻書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("鴻","那鴻書", chap_string, "1","tcv")
							default:
								print_string = Bible_print_string("鴻","那鴻書", chap_string, sec_string,"tcv")
						}
				}
			case "Habakkuk","哈","哈巴","哈巴谷","哈巴谷書","Hab","hab","Книга пророка Аввакума","Ha-ba-cúc","ハバクク書","ハバクク","ハバ","クク","ハバ書","하박국":
				bible_short_name = "哈"
				switch chap_string {
					case "":
						print_string = "哈巴谷書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("哈","哈巴谷書", chap_string, "1","tcv")
							default:
								print_string = Bible_print_string("哈","哈巴谷書", chap_string, sec_string,"tcv")
						}
				}
			case "Zeph","Zephaniah","番","西番雅","西番雅書","Zep","zep","스바냐","ゼパニヤ書","ゼパニヤ","Sô-phô-ni","Книга пророка Софонии":
				bible_short_name = "番"
				switch chap_string {
					case "":
						print_string = "西番雅書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("番","西番雅書", chap_string, "1","tcv")
							default:
								print_string = Bible_print_string("番","西番雅書", chap_string, sec_string,"tcv")
						}
				}
			case "Haggai","該","哈該","哈該書","Hag","hag","학개","ハガイ書","ハガイ","A-ghê","Книга пророка Аггея":
				bible_short_name = "該"
				switch chap_string {
					case "":
						print_string = "哈該書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("該","哈該書", chap_string, "1","tcv")
							default:
								print_string = Bible_print_string("該","哈該書", chap_string, sec_string,"tcv")
						}
				}
			case "Zech","Zechariah","亞","撒迦利亞","撒迦利亞書","Zec","zec","Книга пророка Захарии","Xa-cha-ri","스가랴","ゼカリヤ書","ゼカリヤ":
				bible_short_name = "亞"
				switch chap_string {
					case "":
						print_string = "撒迦利亞書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("亞","撒迦利亞書", chap_string, "1","tcv")
							default:
								print_string = Bible_print_string("亞","撒迦利亞書", chap_string, sec_string,"tcv")
						}
				}
			case "Malachi","瑪","瑪拉","瑪拉基","瑪拉基書","Mal","mal","말라기","マラキ書","マラキ","Ma-la-chi","Книга пророка Малахии":
				bible_short_name = "瑪"
				switch chap_string {
					case "":
						print_string = "瑪拉基書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("瑪","瑪拉基書", chap_string, "1","tcv")
							default:
								print_string = Bible_print_string("瑪","瑪拉基書", chap_string, sec_string,"tcv")
						}
				}
			case "Matt","Matthew","太","馬太","馬太福音","Mt","mt","마태복음","マタイによる福音書","マタイ","マタイによる","Ma-thi-ơ","От Матфея святое благовествование":
				bible_short_name = "太"
				switch chap_string {
					case "":
						print_string = "馬太福音"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("太","馬太福音", chap_string, "1","tcv")
							default:
								print_string = Bible_print_string("太","馬太福音", chap_string, sec_string,"tcv")
						}
				}
			case "Mark","可","馬可","馬可福音","Mr","mr","マルコによる福音書","マルコ","マルコによる","마가복음","Mác","От Марка святое благовествование":
				bible_short_name = "可"
				switch chap_string {
					case "":
						print_string = "馬可福音"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("可","馬可福音", chap_string, "1","tcv")
							default:
								print_string = Bible_print_string("可","馬可福音", chap_string, sec_string,"tcv")
						}
				}
			case "Luke","路","路加","路加福音","Lu","lu","От Луки святое благовествование","Lu-ca","ルカによる福音書","ルカ","ルカによる","누가복음":
				bible_short_name = "路"
				switch chap_string {
					case "":
						print_string = "路加福音"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("路","路加福音", chap_string, "1","tcv")
							default:
								print_string = Bible_print_string("路","路加福音", chap_string, sec_string,"tcv")
						}
				}
			case "John","約","約翰","約翰福音","Joh","joh","От Иоанна святое благовествование","Giăng","ヨハネによる福音書","ヨハネ","ヨハネによる","요한복음":
				bible_short_name = "約"
				switch chap_string {
					case "":
						print_string = "約翰福音"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("約","約翰福音", chap_string, "1","tcv")
							default:
								print_string = Bible_print_string("約","約翰福音", chap_string, sec_string,"tcv")
						}
				}
			case "Acts","徒","使徒","使徒行傳","Ac","ac","Деяния святых апостолов","Công-vụ Các Sứ-đồ","使徒行伝","사도행전":
				bible_short_name = "徒"
				switch chap_string {
					case "":
						print_string = "使徒行傳"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("徒","使徒行傳", chap_string, "1","tcv")
							default:
								print_string = Bible_print_string("徒","使徒行傳", chap_string, sec_string,"tcv")
						}
				}
			case "Rom","Romans","羅","羅馬","羅馬書","Ro","ro","Послание к Римлянам","Rô-ma","ローマ","ローマ人への手紙","로마서":
				bible_short_name = "羅"
				switch chap_string {
					case "":
						print_string = "羅馬書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("羅","羅馬書", chap_string, "1","tcv")
							default:
								print_string = Bible_print_string("羅","羅馬書", chap_string, sec_string,"tcv")
						}
				}
			case "1 Cor","First Corinthians","林前","哥林多前","哥林多前書","1Co","1co","Первое послание к Коринфянам","1 Cô-rinh-tô","コリント人への第一の手紙","コリント一","コリント人への第一","고린도전서":
				bible_short_name = "林前"
				switch chap_string {
					case "":
						print_string = "哥林多前書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("林前","哥林多前書", chap_string, "1","tcv")
							default:
								print_string = Bible_print_string("林前","哥林多前書", chap_string, sec_string,"tcv")
						}
				}
			case "2 Cor","Second Corinthians","林後","哥林多後","哥林多後書","2Co","2co","Второе послание к Коринфянам","2 Cô-rinh-tô","コリント人への第二の手紙","コリント二","コリント人への第二の","고린도후서":
				bible_short_name = "林後"
				switch chap_string {
					case "":
						print_string = "哥林多後書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("林後","哥林多後書", chap_string, "1","tcv")
							default:
								print_string = Bible_print_string("林後","哥林多後書", chap_string, sec_string,"tcv")
						}
				}
			case "Gal","Galatians","加","加拉太","加拉太書","Ga","ga","Послание к Галатам","Ga-la-ti","ガラテヤ","ガラテヤ人への手紙","갈라디아서":
				bible_short_name = "加"
				switch chap_string {
					case "":
						print_string = "加拉太書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("加","加拉太書", chap_string, "1","tcv")
							default:
								print_string = Bible_print_string("加","加拉太書", chap_string, sec_string,"tcv")
						}
				}
			case "Ephesians","弗","以弗所","以弗所書","Eph","eph","Послание к Ефесянам","Ê-phê-sô","エペソ人への手紙","エペソ","エペソ人","エペソ人の手紙","에베소서":
				bible_short_name = "弗"
				switch chap_string {
					case "":
						print_string = "以弗所書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("弗","以弗所書", chap_string, "1","tcv")
							default:
								print_string = Bible_print_string("弗","以弗所書", chap_string, sec_string,"tcv")
						}
				}
			case "Phil","Philippians","腓","腓立","腓立比","腓立比書","Php","php","빌립보서","ピリピ","ピリピ人.ピリピ人への手紙","Послание к Филиппийцам","Phi-líp":
				bible_short_name = "腓"
				switch chap_string {
					case "":
						print_string = "腓立比書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("腓","腓立比書", chap_string, "1","tcv")
							default:
								print_string = Bible_print_string("腓","腓立比書", chap_string, sec_string,"tcv")
						}
				}
			case "Col","col","Colossians","西","歌羅西","歌羅","歌羅西書","Послание к Колоссянам","Cô-lô-se","コロサイ人への手紙","コロサイ","コロ","골로새서":
				bible_short_name = "西"
				switch chap_string {
					case "":
						print_string = "歌羅西書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("西","歌羅西書", chap_string, "1","tcv")
							default:
								print_string = Bible_print_string("西","歌羅西書", chap_string, sec_string,"tcv")
						}
				}
			case "1 Thess","First Thessalonians","帖前","帖撒羅尼迦前","帖撒羅尼迦前書","1Th","1th","데살로니가전서","テサロニケ人への第一の手紙","テサ一","テサロニケ一","1 Tê-sa-lô-ni-ca","Первое послание к Фессалоникийцам (Солунянам)":
				bible_short_name = "帖前"
				switch chap_string {
					case "":
						print_string = "帖撒羅尼迦前書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("帖前","帖撒羅尼迦前書", chap_string, "1","tcv")
							default:
								print_string = Bible_print_string("帖前","帖撒羅尼迦前書", chap_string, sec_string,"tcv")
						}
				}
			case "2 Thess","Second Thessalonians","帖後","帖撒羅尼迦後","帖撒羅尼迦後書","2Th","2th","데살로니가후서","テサロニケ人への第二の手紙","テサ二","テサロニケ二","2 Tê-sa-lô-ni-ca","Второе послание к Фессалоникийцам (Солунянам)":
				bible_short_name = "帖後"
				switch chap_string {
					case "":
						print_string = "帖撒羅尼迦後書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("帖後","帖撒羅尼迦後書", chap_string, "1","tcv")
							default:
								print_string = Bible_print_string("帖後","帖撒羅尼迦後書", chap_string, sec_string,"tcv")
						}
				}
			case "1 Tim","First Timothy","提前","提摩太前","提摩太前書","1Ti","1ti","Первое послание к Тимофею","1 Ti-mô-thê","テモテヘの第一の手紙","テモテ一","디모데전서":
				bible_short_name = "提前"
				switch chap_string {
					case "":
						print_string = "提摩太前書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("提前","提摩太前書", chap_string, "1","tcv")
							default:
								print_string = Bible_print_string("提前","提摩太前書", chap_string, sec_string,"tcv")
						}
				}
			case "2 Tim","Second Timothy","提後","提摩太後","提摩太後書","2Ti","2ti","Второе послание к Тимофею","2 Ti-mô-thê","テモテヘの第二の手紙","テモテ二","디모데후서":
				bible_short_name = "提後"
				switch chap_string {
					case "":
						print_string = "提摩太後書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("提後","提摩太後書", chap_string, "1","tcv")
							default:
								print_string = Bible_print_string("提後","提摩太後書", chap_string, sec_string,"tcv")
						}
				}
			case "Titus","多","提多","提多書","Tit","tit","Послание к Титу","Tít","テトスヘの手紙","テトス","디도서":
				bible_short_name = "多"
				switch chap_string {
					case "":
						print_string = "提多書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("多","提多書", chap_string, "1","tcv")
							default:
								print_string = Bible_print_string("多","提多書", chap_string, sec_string,"tcv")
						}
				}
			case "Philem","Philemon","門","腓利","腓利門","腓利門書","Phm","phm","Послание к Филимону","Phi-lê-môn","ピレモンヘの手紙","ピレモン","빌레몬서":
				bible_short_name = "門"
				switch chap_string {
					case "":
						print_string = "腓利門書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("門","腓利門書", chap_string, "1","tcv")
							default:
								print_string = Bible_print_string("門","腓利門書", chap_string, sec_string,"tcv")
						}
				}
			case "Heb","Hebrews","來","希伯來","希伯來書","heb","Послание к Евреям","Hê-bơ-rơ","ヘブル人への手紙","ヘブル","히브리서":
				bible_short_name = "來"
				switch chap_string {
					case "":
						print_string = "希伯來書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("來","希伯來書", chap_string, "1","tcv")
							default:
								print_string = Bible_print_string("來","希伯來書", chap_string, sec_string,"tcv")
						}
				}
			case "James","雅","雅各","雅各書","Jas","jas","Послание Иакова","Gia-cơ","ヤコブの手紙","ヤコブ","야고보서":
				bible_short_name = "雅"
				switch chap_string {
					case "":
						print_string = "雅各書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("雅","雅各書", chap_string, "1","tcv")
							default:
								print_string = Bible_print_string("雅","雅各書", chap_string, sec_string,"tcv")
						}
				}
			case "1 Pet","First Peter","彼前","彼得前","彼得前書","1Pe","1pe","Первое послание Петра","1 Phi-e-rơ","ペテロの第一の手紙","ペテロ一","베드로전서":
				bible_short_name = "彼前"
				switch chap_string {
					case "":
						print_string = "彼得前書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("彼前","彼得前書", chap_string, "1","tcv")
							default:
								print_string = Bible_print_string("彼前","彼得前書", chap_string, sec_string,"tcv")
						}
				}
			case "2 Pet","Second Peter","彼後","彼得後","彼得後書","2Pe","2pe","Второе послание Петра","2 Phi-e-rơ","ペテロの第二の手紙","ペテロ","베드로후서":
				bible_short_name = "彼後"
				switch chap_string {
					case "":
						print_string = "彼得後書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("彼後","彼得後書", chap_string, "1","tcv")
							default:
								print_string = Bible_print_string("彼後","彼得後書", chap_string, sec_string,"tcv")
						}
				}
			case "1 John","First John","約一","約翰一書","約翰1","約翰1書","1Jo","1jo","Первое послание Иоанна","1 Giăng","ヨハネの第一の手紙","ヨハネ一","요한일서":
				bible_short_name = "約一"
				switch chap_string {
					case "":
						print_string = "約翰一書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("約一","約翰一書", chap_string, "1","tcv")
							default:
								print_string = Bible_print_string("約一","約翰一書", chap_string, sec_string,"tcv")
						}
				}
			case "2 John","second John","約二","約翰二書","約翰2","約翰2書","2Jo","Второе послание Иоанна","2 Giăng","ヨハネの第二の手紙","ヨハネ二","요한2서":
				bible_short_name = "約二"
				switch chap_string {
					case "":
						print_string = "約翰二書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("約二","約翰二書", chap_string, "1","tcv")
							default:
								print_string = Bible_print_string("約二","約翰二書", chap_string, sec_string,"tcv")
						}
				}
			case "3 John","Third John","約三","約翰三書","約翰3","約翰3書","3Jo","3jo","Третье послание Иоанна","3 Giăng","ヨハネの第三の手紙","ヨハネ三","요한3서":
				bible_short_name = "約三"
				switch chap_string {
					case "":
						print_string = "約翰三書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("約三","約翰三書", chap_string, "1","tcv")
							default:
								print_string = Bible_print_string("約三","約翰三書", chap_string, sec_string,"tcv")
						}
				}
			case "Jude","猶","猶大","猶大書","jude","Послание Иуды","Giu-đe","ユダの手紙","ユダ","유다서":
				bible_short_name = "猶"
				switch chap_string {
					case "":
						print_string = "猶大書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("猶","猶大書", chap_string, "1","tcv")
							default:
								print_string = Bible_print_string("猶","猶大書", chap_string, sec_string,"tcv")
						}
				}
			case "Rev","Revelation","啟","啟示","啟示錄","Re","re","ｒｅ","Ｒｅ","rev","Откровение ап. Иоанна Богослова (Апокалипсис)","Khải-huyền","ヨハネの黙示録","黙示録","요한계시록":
				bible_short_name = "啟"
				switch chap_string {
					case "":
						print_string = "啟示錄"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("啟","啟示錄", chap_string, "1","tcv")
							default:
								print_string = Bible_print_string("啟","啟示錄", chap_string, sec_string,"tcv")
						}
				}
			default:
				print_string = "聖經"
				//print_string = "你是要找 " +  reg.ReplaceAllString(text, "$3") + " 對嗎？\n對不起，我還沒學呢...\n"
		}
	case "中文聖經新譯本","ncv","Ncv","Ｎｃｖ","ｎｃｖ":
		log.Print(reg.ReplaceAllString(text, "$3"))
		switch reg.ReplaceAllString(text, "$3") {
			case "Gen","Genesis","創","創世","創世紀","創世記","Ge","ge","gen","창세기","Sáng-thế Ký","Бытие":
				bible_short_name = "創"
				switch chap_string {
					case "":
						print_string = "創世記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("創","創世記", chap_string, "1","ncv")
							default:
								print_string = Bible_print_string("創","創世記", chap_string, sec_string,"ncv")
						}
				}
			case "ex","Ex","Exodus","埃及","出","出埃及","出埃及記","출애굽기","エジプト","出エジプト","出エジプト記","Xuất Ê-díp-tô Ký","Исход":
				bible_short_name = "出"
				switch chap_string {
					case "":
						print_string = "出埃及記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("出","出埃及記", chap_string, "1","ncv")
							default:
								print_string = Bible_print_string("出","出埃及記", chap_string, sec_string,"ncv")
						}
				}
			case "Lev","Leviticus","利","利未","利未記","Le","le","Левит","Lê-vi Ký","レビ記","レビ","레위기":
				bible_short_name = "利"
				switch chap_string {
					case "":
						print_string = "利未記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("利","利未記", chap_string, "1","ncv")
							default:
								print_string = Bible_print_string("利","利未記", chap_string, sec_string,"ncv")
						}
				}
			case "Num","Numbers","民","民數","民數記","Nu","nu","민수기","民数","民数記","Dân-số Ký","Числа":
				bible_short_name = "民"
				switch chap_string {
					case "":
						print_string = "民數記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("民","民數記", chap_string, "1","ncv")
							default:
								print_string = Bible_print_string("民","民數記", chap_string, sec_string,"ncv")
						}
				}
			case "Deut","Deuteronomy","申","申命","申命記","De","de","신명기","Phục-truyền Luật-lệ Ký","Второзаконие":
				bible_short_name = "申"
				switch chap_string {
					case "":
						print_string = "申命記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("申","申命記", chap_string, "1","ncv")
							default:
								print_string = Bible_print_string("申","申命記", chap_string, sec_string,"ncv")
						}
				}
			case "Josh","Joshua","約書亞","約書亞記","Jos","jos","여호수아","ヨシュア記","ヨシュア","Giô-suê","Книга Иисуса Навина":
				bible_short_name = "書"
				switch chap_string {
					case "":
						print_string = "約書亞記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("書","約書亞記", chap_string, "1","ncv")
							default:
								print_string = Bible_print_string("書","約書亞記", chap_string, sec_string,"ncv")
						}
				}
			case "Judg","Judges","士","士師","士師記","Jud","jud","jdg","Jdg","Книга Судей израилевых","Các Quan Xét","사사기":
				bible_short_name = "士"
				switch chap_string {
					case "":
						print_string = "士師記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("士","士師記", chap_string, "1","ncv")
							default:
								print_string = Bible_print_string("士","士師記", chap_string, sec_string,"ncv")
						}
				}
			case "Ruth","路得","路得記","Ru","ru","Rut","rut","룻기","ルツ","ルツ記","Ru-tơ","Книга Руфи":
				bible_short_name = "得"
				switch chap_string {
					case "":
						print_string = "路得記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("得","路得記", chap_string, "1","ncv")
							default:
								print_string = Bible_print_string("得","路得記", chap_string, sec_string,"ncv")
						}
				}
			case "1 Sam","First Samuel","撒上","撒母耳記上","1Sa","1sa","サムエル記上","サムエル上","サム上","사무엘상","1 Sa-mu-ên","Первая книга Царств":
				bible_short_name = "撒上"
				switch chap_string {
					case "":
						print_string = "撒母耳記上"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("撒上","撒母耳記上", chap_string, "1","ncv")
							default:
								print_string = Bible_print_string("撒上","撒母耳記上", chap_string, sec_string,"ncv")
						}
				}
			case "2 Sam","Second Samuel","撒下","撒母耳記下","2Sa","2sa","사무엘하","サムエル記下","サムエル下","サム下","2 Sa-mu-ên","Вторая книга Царств":
				bible_short_name = "撒下"
				switch chap_string {
					case "":
						print_string = "撒母耳記下"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("撒下","撒母耳記下", chap_string, "1","ncv")
							default:
								print_string = Bible_print_string("撒下","撒母耳記下", chap_string, sec_string,"ncv")
						}
				}
			case "1 Kin","First Kings","王上","列王上","列王紀上","列王記上","1Ki","1ki","열왕기상","Третья книга Царств","1 Các Vua":
				bible_short_name = "王上"
				switch chap_string {
					case "":
						print_string = "列王紀上"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("王上","列王紀上", chap_string, "1","ncv")
							default:
								print_string = Bible_print_string("王上","列王紀上", chap_string, sec_string,"ncv")
						}
				}
			case "2 Kin","Second Kings","王下","列王下","列王記下","列王紀下","2Ki","2ki","열왕기하","2 Các Vua","Четвертая книга Царств":
				bible_short_name = "王下"
				switch chap_string {
					case "":
						print_string = "列王紀下"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("王下","列王紀下", chap_string, "1","ncv")
							default:
								print_string = Bible_print_string("王下","列王紀下", chap_string, sec_string,"ncv")
						}
				}
			case "1 Chr","First Chronicles","歷上","代上","歷代志上","歷代上","1Ch","1ch","Первая книга Паралипоменон","1 Sử-ký","歴上","歴代上","歴代志上","역대상":
				bible_short_name = "代上"
				switch chap_string {
					case "":
						print_string = "歷代志上"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("代上","歷代志上", chap_string, "1","ncv")
							default:
								print_string = Bible_print_string("代上","歷代志上", chap_string, sec_string,"ncv")
						}
				}
			case "2 Chr","Second Chronicles","代下","歷下","歷代下","歷代志下","2Ch","2ch","역대하","歴代志下","歴代下","歴下","2 Sử-ký","Вторая книга Паралипоменон":
				bible_short_name = "代下"
				switch chap_string {
					case "":
						print_string = "歷代志下"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("代下","歷代志下", chap_string, "1","ncv")
							default:
								print_string = Bible_print_string("代下","歷代志下", chap_string, sec_string,"ncv")
						}
				}
			case "Ezra","拉","以斯拉","以斯拉記","Ezr","ezr","Первая книга Ездры","Ê-xơ-ra","エズラ","エズラ記","에스라":
				bible_short_name = "拉"
				switch chap_string {
					case "":
						print_string = "以斯拉記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("拉","以斯拉記", chap_string, "1","ncv")
							default:
								print_string = Bible_print_string("拉","以斯拉記", chap_string, sec_string,"ncv")
						}
				}
			case "Neh","Nehemiah","尼","尼希米","尼希米記","Ne","ne","느헤미야","ネヘミヤ書","ネヘミヤ","Nê-hê-mi","Книга Неемии":
				bible_short_name = "尼"
				switch chap_string {
					case "":
						print_string = "尼希米記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("尼","尼希米記", chap_string, "1","ncv")
							default:
								print_string = Bible_print_string("尼","尼希米記", chap_string, sec_string,"ncv")
						}
				}
			case "Esth","Esther","斯","以斯帖","以斯帖記","Es","est","Есфирь","Ê-xơ-tê","エステル","エステル記","에스더":
				bible_short_name = "斯"
				switch chap_string {
					case "":
						print_string = "以斯帖記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("斯","以斯帖記", chap_string, "1","ncv")
							default:
								print_string = Bible_print_string("斯","以斯帖記", chap_string, sec_string,"ncv")
						}
				}
			case "Job","job","伯","約伯","約伯記","Книга Иова","Gióp","ヨブ","ヨブ記","욥기":
				bible_short_name = "伯"
				switch chap_string {
					case "":
						print_string = "約伯記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("伯","約伯記", chap_string, "1","ncv")
							default:
								print_string = Bible_print_string("伯","約伯記", chap_string, sec_string,"ncv")
						}
				}
			case "Ps","Psalms","詩","詩篇","ps","시편","Thi-thiên","Псалтирь":
				bible_short_name = "詩"
				switch chap_string {
					case "":
						print_string = "詩篇"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("詩","詩篇", chap_string, "1","ncv")
							default:
								print_string = Bible_print_string("詩","詩篇", chap_string, sec_string,"ncv")
						}
				}
			case "Prov","Proverbs","箴","箴言","Pr","pr","Притчи Соломона","Châm-ngôn","잠언":
				bible_short_name = "箴"
				switch chap_string {
					case "":
						print_string = "箴言"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("箴","箴言", chap_string, "1","ncv")
							default:
								print_string = Bible_print_string("箴","箴言", chap_string, sec_string,"ncv")
						}
				}
			case "Eccl","Ecclesiastes","傳","傳道","傳道書","Ec","ec","Книга Екклезиаста","или Проповедника","Truyền-đạo","伝道の書","伝道","伝","伝道書","전도서":
				bible_short_name = "傳"
				switch chap_string {
					case "":
						print_string = "傳道書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("傳","傳道書", chap_string, "1","ncv")
							default:
								print_string = Bible_print_string("傳","傳道書", chap_string, sec_string,"ncv")
						}
				}
			case "Song","Song of Solomon","歌","雅歌","So","so","sng","Sng","Песнь песней Соломона","Nhã-ca","아가":
				bible_short_name = "歌"
				switch chap_string {
					case "":
						print_string = "雅歌"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("歌","雅歌", chap_string, "1","ncv")
							default:
								print_string = Bible_print_string("歌","雅歌", chap_string, sec_string,"ncv")
						}
				}
			case "Is","Isaiah","賽","以賽","以賽亞","以賽亞書","Isa","isa","Книга пророка Исаии","Ê-sai","イザヤ書","イザヤ","이사야":
				bible_short_name = "賽"
				switch chap_string {
					case "":
						print_string = "以賽亞書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("賽","以賽亞書", chap_string, "1","ncv")
							default:
								print_string = Bible_print_string("賽","以賽亞書", chap_string, sec_string,"ncv")
						}
				}
			case "Jer","Jeremiah","耶","耶利米","耶利米書","jer","예레미야","エレミヤ","エレミヤ書","Giê-rê-mi","Книга пророка Иеремии":
				bible_short_name = "耶"
				switch chap_string {
					case "":
						print_string = "耶利米書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("耶","耶利米書", chap_string, "1","ncv")
							default:
								print_string = Bible_print_string("耶","耶利米書", chap_string, sec_string,"ncv")
						}
				}
			case "Lam","Lamentations","哀","哀歌","耶利米哀歌","La","lam","예레미야애가","Ca-thương","Плач Иеремии":
				bible_short_name = "哀"
				switch chap_string {
					case "":
						print_string = "耶利米哀歌"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("哀","耶利米哀歌", chap_string, "1","ncv")
							default:
								print_string = Bible_print_string("哀","耶利米哀歌", chap_string, sec_string,"ncv")
						}
				}
			case "Ezek","Ezekiel","結","以西結","以西結書","Eze","eze","에스겔","エゼキエル書","エゼキエル","Ê-xê-chi-ên","Книга пророка Иезекииля":
				bible_short_name = "結"
				switch chap_string {
					case "":
						print_string = "以西結書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("結","以西結書", chap_string, "1","ncv")
							default:
								print_string = Bible_print_string("結","以西結書", chap_string, sec_string,"ncv")
						}
				}
			case "Dan","Daniel","但","但以理","但以理書","Da","da","Книга пророка Даниила","Đa-ni-ên","ダニエル書","ダニエル","다니엘":
				bible_short_name = "但"
				switch chap_string {
					case "":
						print_string = "但以理書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("但","但以理書", chap_string, "1","ncv")
							default:
								print_string = Bible_print_string("但","但以理書", chap_string, sec_string,"ncv")
						}
				}
			case "Hos","Hosea","何","何西","何西阿","何西阿書","Ho","ho","Книга пророка Осии","Ô-sê","ホセア書","ホセア","호세아":
				bible_short_name = "何"
				switch chap_string {
					case "":
						print_string = "何西阿書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("何","何西阿書", chap_string, "1","ncv")
							default:
								print_string = Bible_print_string("何","何西阿書", chap_string, sec_string,"ncv")
						}
				}
			case "Joel","珥","約珥","約珥書","Joe","joe","Книга пророка Иоиля","Giô-ên","ヨエル書","ヨエル","요엘":
				bible_short_name = "珥"
				switch chap_string {
					case "":
						print_string = "約珥書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("珥","約珥書", chap_string, "1","ncv")
							default:
								print_string = Bible_print_string("珥","約珥書", chap_string, sec_string,"ncv")
						}
				}
			case "Amos","摩","阿摩司書","Am","am","Книга пророка Амоса","A-mốt","アモス書","アモス","아모스":
				bible_short_name = "摩"
				switch chap_string {
					case "":
						print_string = "阿摩司書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("摩","阿摩司書", chap_string, "1","ncv")
							default:
								print_string = Bible_print_string("摩","阿摩司書", chap_string, sec_string,"ncv")
						}
				}
			case "Obad","Obadiah","俄","俄巴底亞","俄巴底亞書","Ob","ob","오바댜","オバデヤ書","オバデヤ","Áp-đia","Книга пророка Авдия":
				bible_short_name = "俄"
				switch chap_string {
					case "":
						print_string = "俄巴底亞書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("俄","俄巴底亞書", chap_string, "1","ncv")
							default:
								print_string = Bible_print_string("俄","俄巴底亞書", chap_string, sec_string,"ncv")
						}
				}
			case "Jon","Jonah","拿","約拿","約拿書","jon","요나","ヨナ書","ヨナ","Giô-na","Книга пророка Ионы":
				bible_short_name = "拿"
				switch chap_string {
					case "":
						print_string = "約拿書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("拿","約拿書", chap_string, "1","ncv")
							default:
								print_string = Bible_print_string("拿","約拿書", chap_string, sec_string,"ncv")
						}
				}
			case "Micah","彌","彌迦","彌迦書","Mic","mic","Книга пророка Михея","Mi-chê","ミカ書","ミカ","미가":
				bible_short_name = "彌"
				switch chap_string {
					case "":
						print_string = "彌迦書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("彌","彌迦書", chap_string, "1","ncv")
							default:
								print_string = Bible_print_string("彌","彌迦書", chap_string, sec_string,"ncv")
						}
				}
			case "Nah","Nahum","鴻","那鴻","那鴻書","Na","na","Книга пророка Наума","Na-hum","ナホム書","ナホム","나훔":
				bible_short_name = "鴻"
				switch chap_string {
					case "":
						print_string = "那鴻書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("鴻","那鴻書", chap_string, "1","ncv")
							default:
								print_string = Bible_print_string("鴻","那鴻書", chap_string, sec_string,"ncv")
						}
				}
			case "Habakkuk","哈","哈巴","哈巴谷","哈巴谷書","Hab","hab","Книга пророка Аввакума","Ha-ba-cúc","ハバクク書","ハバクク","ハバ","クク","ハバ書","하박국":
				bible_short_name = "哈"
				switch chap_string {
					case "":
						print_string = "哈巴谷書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("哈","哈巴谷書", chap_string, "1","ncv")
							default:
								print_string = Bible_print_string("哈","哈巴谷書", chap_string, sec_string,"ncv")
						}
				}
			case "Zeph","Zephaniah","番","西番雅","西番雅書","Zep","zep","스바냐","ゼパニヤ書","ゼパニヤ","Sô-phô-ni","Книга пророка Софонии":
				bible_short_name = "番"
				switch chap_string {
					case "":
						print_string = "西番雅書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("番","西番雅書", chap_string, "1","ncv")
							default:
								print_string = Bible_print_string("番","西番雅書", chap_string, sec_string,"ncv")
						}
				}
			case "Haggai","該","哈該","哈該書","Hag","hag","학개","ハガイ書","ハガイ","A-ghê","Книга пророка Аггея":
				bible_short_name = "該"
				switch chap_string {
					case "":
						print_string = "哈該書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("該","哈該書", chap_string, "1","ncv")
							default:
								print_string = Bible_print_string("該","哈該書", chap_string, sec_string,"ncv")
						}
				}
			case "Zech","Zechariah","亞","撒迦利亞","撒迦利亞書","Zec","zec","Книга пророка Захарии","Xa-cha-ri","스가랴","ゼカリヤ書","ゼカリヤ":
				bible_short_name = "亞"
				switch chap_string {
					case "":
						print_string = "撒迦利亞書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("亞","撒迦利亞書", chap_string, "1","ncv")
							default:
								print_string = Bible_print_string("亞","撒迦利亞書", chap_string, sec_string,"ncv")
						}
				}
			case "Malachi","瑪","瑪拉","瑪拉基","瑪拉基書","Mal","mal","말라기","マラキ書","マラキ","Ma-la-chi","Книга пророка Малахии":
				bible_short_name = "瑪"
				switch chap_string {
					case "":
						print_string = "瑪拉基書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("瑪","瑪拉基書", chap_string, "1","ncv")
							default:
								print_string = Bible_print_string("瑪","瑪拉基書", chap_string, sec_string,"ncv")
						}
				}
			case "Matt","Matthew","太","馬太","馬太福音","Mt","mt","마태복음","マタイによる福音書","マタイ","マタイによる","Ma-thi-ơ","От Матфея святое благовествование":
				bible_short_name = "太"
				switch chap_string {
					case "":
						print_string = "馬太福音"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("太","馬太福音", chap_string, "1","ncv")
							default:
								print_string = Bible_print_string("太","馬太福音", chap_string, sec_string,"ncv")
						}
				}
			case "Mark","可","馬可","馬可福音","Mr","mr","マルコによる福音書","マルコ","マルコによる","마가복음","Mác","От Марка святое благовествование":
				bible_short_name = "可"
				switch chap_string {
					case "":
						print_string = "馬可福音"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("可","馬可福音", chap_string, "1","ncv")
							default:
								print_string = Bible_print_string("可","馬可福音", chap_string, sec_string,"ncv")
						}
				}
			case "Luke","路","路加","路加福音","Lu","lu","От Луки святое благовествование","Lu-ca","ルカによる福音書","ルカ","ルカによる","누가복음":
				bible_short_name = "路"
				switch chap_string {
					case "":
						print_string = "路加福音"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("路","路加福音", chap_string, "1","ncv")
							default:
								print_string = Bible_print_string("路","路加福音", chap_string, sec_string,"ncv")
						}
				}
			case "John","約","約翰","約翰福音","Joh","joh","От Иоанна святое благовествование","Giăng","ヨハネによる福音書","ヨハネ","ヨハネによる","요한복음":
				bible_short_name = "約"
				switch chap_string {
					case "":
						print_string = "約翰福音"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("約","約翰福音", chap_string, "1","ncv")
							default:
								print_string = Bible_print_string("約","約翰福音", chap_string, sec_string,"ncv")
						}
				}
			case "Acts","徒","使徒","使徒行傳","Ac","ac","Деяния святых апостолов","Công-vụ Các Sứ-đồ","使徒行伝","사도행전":
				bible_short_name = "徒"
				switch chap_string {
					case "":
						print_string = "使徒行傳"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("徒","使徒行傳", chap_string, "1","ncv")
							default:
								print_string = Bible_print_string("徒","使徒行傳", chap_string, sec_string,"ncv")
						}
				}
			case "Rom","Romans","羅","羅馬","羅馬書","Ro","ro","Послание к Римлянам","Rô-ma","ローマ","ローマ人への手紙","로마서":
				bible_short_name = "羅"
				switch chap_string {
					case "":
						print_string = "羅馬書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("羅","羅馬書", chap_string, "1","ncv")
							default:
								print_string = Bible_print_string("羅","羅馬書", chap_string, sec_string,"ncv")
						}
				}
			case "1 Cor","First Corinthians","林前","哥林多前","哥林多前書","1Co","1co","Первое послание к Коринфянам","1 Cô-rinh-tô","コリント人への第一の手紙","コリント一","コリント人への第一","고린도전서":
				bible_short_name = "林前"
				switch chap_string {
					case "":
						print_string = "哥林多前書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("林前","哥林多前書", chap_string, "1","ncv")
							default:
								print_string = Bible_print_string("林前","哥林多前書", chap_string, sec_string,"ncv")
						}
				}
			case "2 Cor","Second Corinthians","林後","哥林多後","哥林多後書","2Co","2co","Второе послание к Коринфянам","2 Cô-rinh-tô","コリント人への第二の手紙","コリント二","コリント人への第二の","고린도후서":
				bible_short_name = "林後"
				switch chap_string {
					case "":
						print_string = "哥林多後書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("林後","哥林多後書", chap_string, "1","ncv")
							default:
								print_string = Bible_print_string("林後","哥林多後書", chap_string, sec_string,"ncv")
						}
				}
			case "Gal","Galatians","加","加拉太","加拉太書","Ga","ga","Послание к Галатам","Ga-la-ti","ガラテヤ","ガラテヤ人への手紙","갈라디아서":
				bible_short_name = "加"
				switch chap_string {
					case "":
						print_string = "加拉太書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("加","加拉太書", chap_string, "1","ncv")
							default:
								print_string = Bible_print_string("加","加拉太書", chap_string, sec_string,"ncv")
						}
				}
			case "Ephesians","弗","以弗所","以弗所書","Eph","eph","Послание к Ефесянам","Ê-phê-sô","エペソ人への手紙","エペソ","エペソ人","エペソ人の手紙","에베소서":
				bible_short_name = "弗"
				switch chap_string {
					case "":
						print_string = "以弗所書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("弗","以弗所書", chap_string, "1","ncv")
							default:
								print_string = Bible_print_string("弗","以弗所書", chap_string, sec_string,"ncv")
						}
				}
			case "Phil","Philippians","腓","腓立","腓立比","腓立比書","Php","php","빌립보서","ピリピ","ピリピ人.ピリピ人への手紙","Послание к Филиппийцам","Phi-líp":
				bible_short_name = "腓"
				switch chap_string {
					case "":
						print_string = "腓立比書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("腓","腓立比書", chap_string, "1","ncv")
							default:
								print_string = Bible_print_string("腓","腓立比書", chap_string, sec_string,"ncv")
						}
				}
			case "Col","col","Colossians","西","歌羅西","歌羅","歌羅西書","Послание к Колоссянам","Cô-lô-se","コロサイ人への手紙","コロサイ","コロ","골로새서":
				bible_short_name = "西"
				switch chap_string {
					case "":
						print_string = "歌羅西書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("西","歌羅西書", chap_string, "1","ncv")
							default:
								print_string = Bible_print_string("西","歌羅西書", chap_string, sec_string,"ncv")
						}
				}
			case "1 Thess","First Thessalonians","帖前","帖撒羅尼迦前","帖撒羅尼迦前書","1Th","1th","데살로니가전서","テサロニケ人への第一の手紙","テサ一","テサロニケ一","1 Tê-sa-lô-ni-ca","Первое послание к Фессалоникийцам (Солунянам)":
				bible_short_name = "帖前"
				switch chap_string {
					case "":
						print_string = "帖撒羅尼迦前書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("帖前","帖撒羅尼迦前書", chap_string, "1","ncv")
							default:
								print_string = Bible_print_string("帖前","帖撒羅尼迦前書", chap_string, sec_string,"ncv")
						}
				}
			case "2 Thess","Second Thessalonians","帖後","帖撒羅尼迦後","帖撒羅尼迦後書","2Th","2th","데살로니가후서","テサロニケ人への第二の手紙","テサ二","テサロニケ二","2 Tê-sa-lô-ni-ca","Второе послание к Фессалоникийцам (Солунянам)":
				bible_short_name = "帖後"
				switch chap_string {
					case "":
						print_string = "帖撒羅尼迦後書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("帖後","帖撒羅尼迦後書", chap_string, "1","ncv")
							default:
								print_string = Bible_print_string("帖後","帖撒羅尼迦後書", chap_string, sec_string,"ncv")
						}
				}
			case "1 Tim","First Timothy","提前","提摩太前","提摩太前書","1Ti","1ti","Первое послание к Тимофею","1 Ti-mô-thê","テモテヘの第一の手紙","テモテ一","디모데전서":
				bible_short_name = "提前"
				switch chap_string {
					case "":
						print_string = "提摩太前書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("提前","提摩太前書", chap_string, "1","ncv")
							default:
								print_string = Bible_print_string("提前","提摩太前書", chap_string, sec_string,"ncv")
						}
				}
			case "2 Tim","Second Timothy","提後","提摩太後","提摩太後書","2Ti","2ti","Второе послание к Тимофею","2 Ti-mô-thê","テモテヘの第二の手紙","テモテ二","디모데후서":
				bible_short_name = "提後"
				switch chap_string {
					case "":
						print_string = "提摩太後書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("提後","提摩太後書", chap_string, "1","ncv")
							default:
								print_string = Bible_print_string("提後","提摩太後書", chap_string, sec_string,"ncv")
						}
				}
			case "Titus","多","提多","提多書","Tit","tit","Послание к Титу","Tít","テトスヘの手紙","テトス","디도서":
				bible_short_name = "多"
				switch chap_string {
					case "":
						print_string = "提多書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("多","提多書", chap_string, "1","ncv")
							default:
								print_string = Bible_print_string("多","提多書", chap_string, sec_string,"ncv")
						}
				}
			case "Philem","Philemon","門","腓利","腓利門","腓利門書","Phm","phm","Послание к Филимону","Phi-lê-môn","ピレモンヘの手紙","ピレモン","빌레몬서":
				bible_short_name = "門"
				switch chap_string {
					case "":
						print_string = "腓利門書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("門","腓利門書", chap_string, "1","ncv")
							default:
								print_string = Bible_print_string("門","腓利門書", chap_string, sec_string,"ncv")
						}
				}
			case "Heb","Hebrews","來","希伯來","希伯來書","heb","Послание к Евреям","Hê-bơ-rơ","ヘブル人への手紙","ヘブル","히브리서":
				bible_short_name = "來"
				switch chap_string {
					case "":
						print_string = "希伯來書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("來","希伯來書", chap_string, "1","ncv")
							default:
								print_string = Bible_print_string("來","希伯來書", chap_string, sec_string,"ncv")
						}
				}
			case "James","雅","雅各","雅各書","Jas","jas","Послание Иакова","Gia-cơ","ヤコブの手紙","ヤコブ","야고보서":
				bible_short_name = "雅"
				switch chap_string {
					case "":
						print_string = "雅各書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("雅","雅各書", chap_string, "1","ncv")
							default:
								print_string = Bible_print_string("雅","雅各書", chap_string, sec_string,"ncv")
						}
				}
			case "1 Pet","First Peter","彼前","彼得前","彼得前書","1Pe","1pe","Первое послание Петра","1 Phi-e-rơ","ペテロの第一の手紙","ペテロ一","베드로전서":
				bible_short_name = "彼前"
				switch chap_string {
					case "":
						print_string = "彼得前書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("彼前","彼得前書", chap_string, "1","ncv")
							default:
								print_string = Bible_print_string("彼前","彼得前書", chap_string, sec_string,"ncv")
						}
				}
			case "2 Pet","Second Peter","彼後","彼得後","彼得後書","2Pe","2pe","Второе послание Петра","2 Phi-e-rơ","ペテロの第二の手紙","ペテロ","베드로후서":
				bible_short_name = "彼後"
				switch chap_string {
					case "":
						print_string = "彼得後書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("彼後","彼得後書", chap_string, "1","ncv")
							default:
								print_string = Bible_print_string("彼後","彼得後書", chap_string, sec_string,"ncv")
						}
				}
			case "1 John","First John","約一","約翰一書","約翰1","約翰1書","1Jo","1jo","Первое послание Иоанна","1 Giăng","ヨハネの第一の手紙","ヨハネ一","요한일서":
				bible_short_name = "約一"
				switch chap_string {
					case "":
						print_string = "約翰一書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("約一","約翰一書", chap_string, "1","ncv")
							default:
								print_string = Bible_print_string("約一","約翰一書", chap_string, sec_string,"ncv")
						}
				}
			case "2 John","second John","約二","約翰二書","約翰2","約翰2書","2Jo","Второе послание Иоанна","2 Giăng","ヨハネの第二の手紙","ヨハネ二","요한2서":
				bible_short_name = "約二"
				switch chap_string {
					case "":
						print_string = "約翰二書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("約二","約翰二書", chap_string, "1","ncv")
							default:
								print_string = Bible_print_string("約二","約翰二書", chap_string, sec_string,"ncv")
						}
				}
			case "3 John","Third John","約三","約翰三書","約翰3","約翰3書","3Jo","3jo","Третье послание Иоанна","3 Giăng","ヨハネの第三の手紙","ヨハネ三","요한3서":
				bible_short_name = "約三"
				switch chap_string {
					case "":
						print_string = "約翰三書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("約三","約翰三書", chap_string, "1","ncv")
							default:
								print_string = Bible_print_string("約三","約翰三書", chap_string, sec_string,"ncv")
						}
				}
			case "Jude","猶","猶大","猶大書","jude","Послание Иуды","Giu-đe","ユダの手紙","ユダ","유다서":
				bible_short_name = "猶"
				switch chap_string {
					case "":
						print_string = "猶大書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("猶","猶大書", chap_string, "1","ncv")
							default:
								print_string = Bible_print_string("猶","猶大書", chap_string, sec_string,"ncv")
						}
				}
			case "Rev","Revelation","啟","啟示","啟示錄","Re","re","ｒｅ","Ｒｅ","rev","Откровение ап. Иоанна Богослова (Апокалипсис)","Khải-huyền","ヨハネの黙示録","黙示録","요한계시록":
				bible_short_name = "啟"
				switch chap_string {
					case "":
						print_string = "啟示錄"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("啟","啟示錄", chap_string, "1","ncv")
							default:
								print_string = Bible_print_string("啟","啟示錄", chap_string, sec_string,"ncv")
						}
				}
			default:
				print_string = "聖經"
				//print_string = "你是要找 " +  reg.ReplaceAllString(text, "$3") + " 對嗎？\n對不起，我還沒學呢...\n"
		}
	case "中文聖經","中文聖經和合本修訂版","Rcuv","rcuv","ｒｃｕｖ","Ｒｃｕｖ":
		log.Print(reg.ReplaceAllString(text, "$3"))
		switch reg.ReplaceAllString(text, "$3") {
			case "Gen","Genesis","創","創世","創世紀","創世記","Ge","ge","gen","창세기","Sáng-thế Ký","Бытие":
				bible_short_name = "創"
				switch chap_string {
					case "":
						print_string = "創世記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("創","創世記", chap_string, "1","rcuv")
							default:
								print_string = Bible_print_string("創","創世記", chap_string, sec_string,"rcuv")
						}
				}
			case "ex","Ex","Exodus","埃及","出","出埃及","出埃及記","출애굽기","エジプト","出エジプト","出エジプト記","Xuất Ê-díp-tô Ký","Исход":
				bible_short_name = "出"
				switch chap_string {
					case "":
						print_string = "出埃及記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("出","出埃及記", chap_string, "1","rcuv")
							default:
								print_string = Bible_print_string("出","出埃及記", chap_string, sec_string,"rcuv")
						}
				}
			case "Lev","Leviticus","利","利未","利未記","Le","le","Левит","Lê-vi Ký","レビ記","レビ","레위기":
				bible_short_name = "利"
				switch chap_string {
					case "":
						print_string = "利未記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("利","利未記", chap_string, "1","rcuv")
							default:
								print_string = Bible_print_string("利","利未記", chap_string, sec_string,"rcuv")
						}
				}
			case "Num","Numbers","民","民數","民數記","Nu","nu","민수기","民数","民数記","Dân-số Ký","Числа":
				bible_short_name = "民"
				switch chap_string {
					case "":
						print_string = "民數記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("民","民數記", chap_string, "1","rcuv")
							default:
								print_string = Bible_print_string("民","民數記", chap_string, sec_string,"rcuv")
						}
				}
			case "Deut","Deuteronomy","申","申命","申命記","De","de","신명기","Phục-truyền Luật-lệ Ký","Второзаконие":
				bible_short_name = "申"
				switch chap_string {
					case "":
						print_string = "申命記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("申","申命記", chap_string, "1","rcuv")
							default:
								print_string = Bible_print_string("申","申命記", chap_string, sec_string,"rcuv")
						}
				}
			case "Josh","Joshua","約書亞","約書亞記","Jos","jos","여호수아","ヨシュア記","ヨシュア","Giô-suê","Книга Иисуса Навина":
				bible_short_name = "書"
				switch chap_string {
					case "":
						print_string = "約書亞記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("書","約書亞記", chap_string, "1","rcuv")
							default:
								print_string = Bible_print_string("書","約書亞記", chap_string, sec_string,"rcuv")
						}
				}
			case "Judg","Judges","士","士師","士師記","Jud","jud","jdg","Jdg","Книга Судей израилевых","Các Quan Xét","사사기":
				bible_short_name = "士"
				switch chap_string {
					case "":
						print_string = "士師記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("士","士師記", chap_string, "1","rcuv")
							default:
								print_string = Bible_print_string("士","士師記", chap_string, sec_string,"rcuv")
						}
				}
			case "Ruth","路得","路得記","Ru","ru","Rut","rut","룻기","ルツ","ルツ記","Ru-tơ","Книга Руфи":
				bible_short_name = "得"
				switch chap_string {
					case "":
						print_string = "路得記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("得","路得記", chap_string, "1","rcuv")
							default:
								print_string = Bible_print_string("得","路得記", chap_string, sec_string,"rcuv")
						}
				}
			case "1 Sam","First Samuel","撒上","撒母耳記上","1Sa","1sa","サムエル記上","サムエル上","サム上","사무엘상","1 Sa-mu-ên","Первая книга Царств":
				bible_short_name = "撒上"
				switch chap_string {
					case "":
						print_string = "撒母耳記上"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("撒上","撒母耳記上", chap_string, "1","rcuv")
							default:
								print_string = Bible_print_string("撒上","撒母耳記上", chap_string, sec_string,"rcuv")
						}
				}
			case "2 Sam","Second Samuel","撒下","撒母耳記下","2Sa","2sa","사무엘하","サムエル記下","サムエル下","サム下","2 Sa-mu-ên","Вторая книга Царств":
				bible_short_name = "撒下"
				switch chap_string {
					case "":
						print_string = "撒母耳記下"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("撒下","撒母耳記下", chap_string, "1","rcuv")
							default:
								print_string = Bible_print_string("撒下","撒母耳記下", chap_string, sec_string,"rcuv")
						}
				}
			case "1 Kin","First Kings","王上","列王上","列王紀上","列王記上","1Ki","1ki","열왕기상","Третья книга Царств","1 Các Vua":
				bible_short_name = "王上"
				switch chap_string {
					case "":
						print_string = "列王紀上"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("王上","列王紀上", chap_string, "1","rcuv")
							default:
								print_string = Bible_print_string("王上","列王紀上", chap_string, sec_string,"rcuv")
						}
				}
			case "2 Kin","Second Kings","王下","列王下","列王記下","列王紀下","2Ki","2ki","열왕기하","2 Các Vua","Четвертая книга Царств":
				bible_short_name = "王下"
				switch chap_string {
					case "":
						print_string = "列王紀下"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("王下","列王紀下", chap_string, "1","rcuv")
							default:
								print_string = Bible_print_string("王下","列王紀下", chap_string, sec_string,"rcuv")
						}
				}
			case "1 Chr","First Chronicles","歷上","代上","歷代志上","歷代上","1Ch","1ch","Первая книга Паралипоменон","1 Sử-ký","歴上","歴代上","歴代志上","역대상":
				bible_short_name = "代上"
				switch chap_string {
					case "":
						print_string = "歷代志上"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("代上","歷代志上", chap_string, "1","rcuv")
							default:
								print_string = Bible_print_string("代上","歷代志上", chap_string, sec_string,"rcuv")
						}
				}
			case "2 Chr","Second Chronicles","代下","歷下","歷代下","歷代志下","2Ch","2ch","역대하","歴代志下","歴代下","歴下","2 Sử-ký","Вторая книга Паралипоменон":
				bible_short_name = "代下"
				switch chap_string {
					case "":
						print_string = "歷代志下"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("代下","歷代志下", chap_string, "1","rcuv")
							default:
								print_string = Bible_print_string("代下","歷代志下", chap_string, sec_string,"rcuv")
						}
				}
			case "Ezra","拉","以斯拉","以斯拉記","Ezr","ezr","Первая книга Ездры","Ê-xơ-ra","エズラ","エズラ記","에스라":
				bible_short_name = "拉"
				switch chap_string {
					case "":
						print_string = "以斯拉記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("拉","以斯拉記", chap_string, "1","rcuv")
							default:
								print_string = Bible_print_string("拉","以斯拉記", chap_string, sec_string,"rcuv")
						}
				}
			case "Neh","Nehemiah","尼","尼希米","尼希米記","Ne","ne","느헤미야","ネヘミヤ書","ネヘミヤ","Nê-hê-mi","Книга Неемии":
				bible_short_name = "尼"
				switch chap_string {
					case "":
						print_string = "尼希米記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("尼","尼希米記", chap_string, "1","rcuv")
							default:
								print_string = Bible_print_string("尼","尼希米記", chap_string, sec_string,"rcuv")
						}
				}
			case "Esth","Esther","斯","以斯帖","以斯帖記","Es","est","Есфирь","Ê-xơ-tê","エステル","エステル記","에스더":
				bible_short_name = "斯"
				switch chap_string {
					case "":
						print_string = "以斯帖記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("斯","以斯帖記", chap_string, "1","rcuv")
							default:
								print_string = Bible_print_string("斯","以斯帖記", chap_string, sec_string,"rcuv")
						}
				}
			case "Job","job","伯","約伯","約伯記","Книга Иова","Gióp","ヨブ","ヨブ記","욥기":
				bible_short_name = "伯"
				switch chap_string {
					case "":
						print_string = "約伯記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("伯","約伯記", chap_string, "1","rcuv")
							default:
								print_string = Bible_print_string("伯","約伯記", chap_string, sec_string,"rcuv")
						}
				}
			case "Ps","Psalms","詩","詩篇","ps","시편","Thi-thiên","Псалтирь":
				bible_short_name = "詩"
				switch chap_string {
					case "":
						print_string = "詩篇"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("詩","詩篇", chap_string, "1","rcuv")
							default:
								print_string = Bible_print_string("詩","詩篇", chap_string, sec_string,"rcuv")
						}
				}
			case "Prov","Proverbs","箴","箴言","Pr","pr","Притчи Соломона","Châm-ngôn","잠언":
				bible_short_name = "箴"
				switch chap_string {
					case "":
						print_string = "箴言"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("箴","箴言", chap_string, "1","rcuv")
							default:
								print_string = Bible_print_string("箴","箴言", chap_string, sec_string,"rcuv")
						}
				}
			case "Eccl","Ecclesiastes","傳","傳道","傳道書","Ec","ec","Книга Екклезиаста","или Проповедника","Truyền-đạo","伝道の書","伝道","伝","伝道書","전도서":
				bible_short_name = "傳"
				switch chap_string {
					case "":
						print_string = "傳道書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("傳","傳道書", chap_string, "1","rcuv")
							default:
								print_string = Bible_print_string("傳","傳道書", chap_string, sec_string,"rcuv")
						}
				}
			case "Song","Song of Solomon","歌","雅歌","So","so","sng","Sng","Песнь песней Соломона","Nhã-ca","아가":
				bible_short_name = "歌"
				switch chap_string {
					case "":
						print_string = "雅歌"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("歌","雅歌", chap_string, "1","rcuv")
							default:
								print_string = Bible_print_string("歌","雅歌", chap_string, sec_string,"rcuv")
						}
				}
			case "Is","Isaiah","賽","以賽","以賽亞","以賽亞書","Isa","isa","Книга пророка Исаии","Ê-sai","イザヤ書","イザヤ","이사야":
				bible_short_name = "賽"
				switch chap_string {
					case "":
						print_string = "以賽亞書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("賽","以賽亞書", chap_string, "1","rcuv")
							default:
								print_string = Bible_print_string("賽","以賽亞書", chap_string, sec_string,"rcuv")
						}
				}
			case "Jer","Jeremiah","耶","耶利米","耶利米書","jer","예레미야","エレミヤ","エレミヤ書","Giê-rê-mi","Книга пророка Иеремии":
				bible_short_name = "耶"
				switch chap_string {
					case "":
						print_string = "耶利米書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("耶","耶利米書", chap_string, "1","rcuv")
							default:
								print_string = Bible_print_string("耶","耶利米書", chap_string, sec_string,"rcuv")
						}
				}
			case "Lam","Lamentations","哀","哀歌","耶利米哀歌","La","lam","예레미야애가","Ca-thương","Плач Иеремии":
				bible_short_name = "哀"
				switch chap_string {
					case "":
						print_string = "耶利米哀歌"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("哀","耶利米哀歌", chap_string, "1","rcuv")
							default:
								print_string = Bible_print_string("哀","耶利米哀歌", chap_string, sec_string,"rcuv")
						}
				}
			case "Ezek","Ezekiel","結","以西結","以西結書","Eze","eze","에스겔","エゼキエル書","エゼキエル","Ê-xê-chi-ên","Книга пророка Иезекииля":
				bible_short_name = "結"
				switch chap_string {
					case "":
						print_string = "以西結書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("結","以西結書", chap_string, "1","rcuv")
							default:
								print_string = Bible_print_string("結","以西結書", chap_string, sec_string,"rcuv")
						}
				}
			case "Dan","Daniel","但","但以理","但以理書","Da","da","Книга пророка Даниила","Đa-ni-ên","ダニエル書","ダニエル","다니엘":
				bible_short_name = "但"
				switch chap_string {
					case "":
						print_string = "但以理書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("但","但以理書", chap_string, "1","rcuv")
							default:
								print_string = Bible_print_string("但","但以理書", chap_string, sec_string,"rcuv")
						}
				}
			case "Hos","Hosea","何","何西","何西阿","何西阿書","Ho","ho","Книга пророка Осии","Ô-sê","ホセア書","ホセア","호세아":
				bible_short_name = "何"
				switch chap_string {
					case "":
						print_string = "何西阿書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("何","何西阿書", chap_string, "1","rcuv")
							default:
								print_string = Bible_print_string("何","何西阿書", chap_string, sec_string,"rcuv")
						}
				}
			case "Joel","珥","約珥","約珥書","Joe","joe","Книга пророка Иоиля","Giô-ên","ヨエル書","ヨエル","요엘":
				bible_short_name = "珥"
				switch chap_string {
					case "":
						print_string = "約珥書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("珥","約珥書", chap_string, "1","rcuv")
							default:
								print_string = Bible_print_string("珥","約珥書", chap_string, sec_string,"rcuv")
						}
				}
			case "Amos","摩","阿摩司書","Am","am","Книга пророка Амоса","A-mốt","アモス書","アモス","아모스":
				bible_short_name = "摩"
				switch chap_string {
					case "":
						print_string = "阿摩司書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("摩","阿摩司書", chap_string, "1","rcuv")
							default:
								print_string = Bible_print_string("摩","阿摩司書", chap_string, sec_string,"rcuv")
						}
				}
			case "Obad","Obadiah","俄","俄巴底亞","俄巴底亞書","Ob","ob","오바댜","オバデヤ書","オバデヤ","Áp-đia","Книга пророка Авдия":
				bible_short_name = "俄"
				switch chap_string {
					case "":
						print_string = "俄巴底亞書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("俄","俄巴底亞書", chap_string, "1","rcuv")
							default:
								print_string = Bible_print_string("俄","俄巴底亞書", chap_string, sec_string,"rcuv")
						}
				}
			case "Jon","Jonah","拿","約拿","約拿書","jon","요나","ヨナ書","ヨナ","Giô-na","Книга пророка Ионы":
				bible_short_name = "拿"
				switch chap_string {
					case "":
						print_string = "約拿書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("拿","約拿書", chap_string, "1","rcuv")
							default:
								print_string = Bible_print_string("拿","約拿書", chap_string, sec_string,"rcuv")
						}
				}
			case "Micah","彌","彌迦","彌迦書","Mic","mic","Книга пророка Михея","Mi-chê","ミカ書","ミカ","미가":
				bible_short_name = "彌"
				switch chap_string {
					case "":
						print_string = "彌迦書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("彌","彌迦書", chap_string, "1","rcuv")
							default:
								print_string = Bible_print_string("彌","彌迦書", chap_string, sec_string,"rcuv")
						}
				}
			case "Nah","Nahum","鴻","那鴻","那鴻書","Na","na","Книга пророка Наума","Na-hum","ナホム書","ナホム","나훔":
				bible_short_name = "鴻"
				switch chap_string {
					case "":
						print_string = "那鴻書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("鴻","那鴻書", chap_string, "1","rcuv")
							default:
								print_string = Bible_print_string("鴻","那鴻書", chap_string, sec_string,"rcuv")
						}
				}
			case "Habakkuk","哈","哈巴","哈巴谷","哈巴谷書","Hab","hab","Книга пророка Аввакума","Ha-ba-cúc","ハバクク書","ハバクク","ハバ","クク","ハバ書","하박국":
				bible_short_name = "哈"
				switch chap_string {
					case "":
						print_string = "哈巴谷書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("哈","哈巴谷書", chap_string, "1","rcuv")
							default:
								print_string = Bible_print_string("哈","哈巴谷書", chap_string, sec_string,"rcuv")
						}
				}
			case "Zeph","Zephaniah","番","西番雅","西番雅書","Zep","zep","스바냐","ゼパニヤ書","ゼパニヤ","Sô-phô-ni","Книга пророка Софонии":
				bible_short_name = "番"
				switch chap_string {
					case "":
						print_string = "西番雅書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("番","西番雅書", chap_string, "1","rcuv")
							default:
								print_string = Bible_print_string("番","西番雅書", chap_string, sec_string,"rcuv")
						}
				}
			case "Haggai","該","哈該","哈該書","Hag","hag","학개","ハガイ書","ハガイ","A-ghê","Книга пророка Аггея":
				bible_short_name = "該"
				switch chap_string {
					case "":
						print_string = "哈該書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("該","哈該書", chap_string, "1","rcuv")
							default:
								print_string = Bible_print_string("該","哈該書", chap_string, sec_string,"rcuv")
						}
				}
			case "Zech","Zechariah","亞","撒迦利亞","撒迦利亞書","Zec","zec","Книга пророка Захарии","Xa-cha-ri","스가랴","ゼカリヤ書","ゼカリヤ":
				bible_short_name = "亞"
				switch chap_string {
					case "":
						print_string = "撒迦利亞書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("亞","撒迦利亞書", chap_string, "1","rcuv")
							default:
								print_string = Bible_print_string("亞","撒迦利亞書", chap_string, sec_string,"rcuv")
						}
				}
			case "Malachi","瑪","瑪拉","瑪拉基","瑪拉基書","Mal","mal","말라기","マラキ書","マラキ","Ma-la-chi","Книга пророка Малахии":
				bible_short_name = "瑪"
				switch chap_string {
					case "":
						print_string = "瑪拉基書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("瑪","瑪拉基書", chap_string, "1","rcuv")
							default:
								print_string = Bible_print_string("瑪","瑪拉基書", chap_string, sec_string,"rcuv")
						}
				}
			case "Matt","Matthew","太","馬太","馬太福音","Mt","mt","마태복음","マタイによる福音書","マタイ","マタイによる","Ma-thi-ơ","От Матфея святое благовествование":
				bible_short_name = "太"
				switch chap_string {
					case "":
						print_string = "馬太福音"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("太","馬太福音", chap_string, "1","rcuv")
							default:
								print_string = Bible_print_string("太","馬太福音", chap_string, sec_string,"rcuv")
						}
				}
			case "Mark","可","馬可","馬可福音","Mr","mr","マルコによる福音書","マルコ","マルコによる","마가복음","Mác","От Марка святое благовествование":
				bible_short_name = "可"
				switch chap_string {
					case "":
						print_string = "馬可福音"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("可","馬可福音", chap_string, "1","rcuv")
							default:
								print_string = Bible_print_string("可","馬可福音", chap_string, sec_string,"rcuv")
						}
				}
			case "Luke","路","路加","路加福音","Lu","lu","От Луки святое благовествование","Lu-ca","ルカによる福音書","ルカ","ルカによる","누가복음":
				bible_short_name = "路"
				switch chap_string {
					case "":
						print_string = "路加福音"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("路","路加福音", chap_string, "1","rcuv")
							default:
								print_string = Bible_print_string("路","路加福音", chap_string, sec_string,"rcuv")
						}
				}
			case "John","約","約翰","約翰福音","Joh","joh","От Иоанна святое благовествование","Giăng","ヨハネによる福音書","ヨハネ","ヨハネによる","요한복음":
				bible_short_name = "約"
				switch chap_string {
					case "":
						print_string = "約翰福音"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("約","約翰福音", chap_string, "1","rcuv")
							default:
								print_string = Bible_print_string("約","約翰福音", chap_string, sec_string,"rcuv")
						}
				}
			case "Acts","徒","使徒","使徒行傳","Ac","ac","Деяния святых апостолов","Công-vụ Các Sứ-đồ","使徒行伝","사도행전":
				bible_short_name = "徒"
				switch chap_string {
					case "":
						print_string = "使徒行傳"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("徒","使徒行傳", chap_string, "1","rcuv")
							default:
								print_string = Bible_print_string("徒","使徒行傳", chap_string, sec_string,"rcuv")
						}
				}
			case "Rom","Romans","羅","羅馬","羅馬書","Ro","ro","Послание к Римлянам","Rô-ma","ローマ","ローマ人への手紙","로마서":
				bible_short_name = "羅"
				switch chap_string {
					case "":
						print_string = "羅馬書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("羅","羅馬書", chap_string, "1","rcuv")
							default:
								print_string = Bible_print_string("羅","羅馬書", chap_string, sec_string,"rcuv")
						}
				}
			case "1 Cor","First Corinthians","林前","哥林多前","哥林多前書","1Co","1co","Первое послание к Коринфянам","1 Cô-rinh-tô","コリント人への第一の手紙","コリント一","コリント人への第一","고린도전서":
				bible_short_name = "林前"
				switch chap_string {
					case "":
						print_string = "哥林多前書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("林前","哥林多前書", chap_string, "1","rcuv")
							default:
								print_string = Bible_print_string("林前","哥林多前書", chap_string, sec_string,"rcuv")
						}
				}
			case "2 Cor","Second Corinthians","林後","哥林多後","哥林多後書","2Co","2co","Второе послание к Коринфянам","2 Cô-rinh-tô","コリント人への第二の手紙","コリント二","コリント人への第二の","고린도후서":
				bible_short_name = "林後"
				switch chap_string {
					case "":
						print_string = "哥林多後書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("林後","哥林多後書", chap_string, "1","rcuv")
							default:
								print_string = Bible_print_string("林後","哥林多後書", chap_string, sec_string,"rcuv")
						}
				}
			case "Gal","Galatians","加","加拉太","加拉太書","Ga","ga","Послание к Галатам","Ga-la-ti","ガラテヤ","ガラテヤ人への手紙","갈라디아서":
				bible_short_name = "加"
				switch chap_string {
					case "":
						print_string = "加拉太書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("加","加拉太書", chap_string, "1","rcuv")
							default:
								print_string = Bible_print_string("加","加拉太書", chap_string, sec_string,"rcuv")
						}
				}
			case "Ephesians","弗","以弗所","以弗所書","Eph","eph","Послание к Ефесянам","Ê-phê-sô","エペソ人への手紙","エペソ","エペソ人","エペソ人の手紙","에베소서":
				bible_short_name = "弗"
				switch chap_string {
					case "":
						print_string = "以弗所書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("弗","以弗所書", chap_string, "1","rcuv")
							default:
								print_string = Bible_print_string("弗","以弗所書", chap_string, sec_string,"rcuv")
						}
				}
			case "Phil","Philippians","腓","腓立","腓立比","腓立比書","Php","php","빌립보서","ピリピ","ピリピ人.ピリピ人への手紙","Послание к Филиппийцам","Phi-líp":
				bible_short_name = "腓"
				switch chap_string {
					case "":
						print_string = "腓立比書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("腓","腓立比書", chap_string, "1","rcuv")
							default:
								print_string = Bible_print_string("腓","腓立比書", chap_string, sec_string,"rcuv")
						}
				}
			case "Col","col","Colossians","西","歌羅西","歌羅","歌羅西書","Послание к Колоссянам","Cô-lô-se","コロサイ人への手紙","コロサイ","コロ","골로새서":
				bible_short_name = "西"
				switch chap_string {
					case "":
						print_string = "歌羅西書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("西","歌羅西書", chap_string, "1","rcuv")
							default:
								print_string = Bible_print_string("西","歌羅西書", chap_string, sec_string,"rcuv")
						}
				}
			case "1 Thess","First Thessalonians","帖前","帖撒羅尼迦前","帖撒羅尼迦前書","1Th","1th","데살로니가전서","テサロニケ人への第一の手紙","テサ一","テサロニケ一","1 Tê-sa-lô-ni-ca","Первое послание к Фессалоникийцам (Солунянам)":
				bible_short_name = "帖前"
				switch chap_string {
					case "":
						print_string = "帖撒羅尼迦前書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("帖前","帖撒羅尼迦前書", chap_string, "1","rcuv")
							default:
								print_string = Bible_print_string("帖前","帖撒羅尼迦前書", chap_string, sec_string,"rcuv")
						}
				}
			case "2 Thess","Second Thessalonians","帖後","帖撒羅尼迦後","帖撒羅尼迦後書","2Th","2th","데살로니가후서","テサロニケ人への第二の手紙","テサ二","テサロニケ二","2 Tê-sa-lô-ni-ca","Второе послание к Фессалоникийцам (Солунянам)":
				bible_short_name = "帖後"
				switch chap_string {
					case "":
						print_string = "帖撒羅尼迦後書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("帖後","帖撒羅尼迦後書", chap_string, "1","rcuv")
							default:
								print_string = Bible_print_string("帖後","帖撒羅尼迦後書", chap_string, sec_string,"rcuv")
						}
				}
			case "1 Tim","First Timothy","提前","提摩太前","提摩太前書","1Ti","1ti","Первое послание к Тимофею","1 Ti-mô-thê","テモテヘの第一の手紙","テモテ一","디모데전서":
				bible_short_name = "提前"
				switch chap_string {
					case "":
						print_string = "提摩太前書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("提前","提摩太前書", chap_string, "1","rcuv")
							default:
								print_string = Bible_print_string("提前","提摩太前書", chap_string, sec_string,"rcuv")
						}
				}
			case "2 Tim","Second Timothy","提後","提摩太後","提摩太後書","2Ti","2ti","Второе послание к Тимофею","2 Ti-mô-thê","テモテヘの第二の手紙","テモテ二","디모데후서":
				bible_short_name = "提後"
				switch chap_string {
					case "":
						print_string = "提摩太後書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("提後","提摩太後書", chap_string, "1","rcuv")
							default:
								print_string = Bible_print_string("提後","提摩太後書", chap_string, sec_string,"rcuv")
						}
				}
			case "Titus","多","提多","提多書","Tit","tit","Послание к Титу","Tít","テトスヘの手紙","テトス","디도서":
				bible_short_name = "多"
				switch chap_string {
					case "":
						print_string = "提多書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("多","提多書", chap_string, "1","rcuv")
							default:
								print_string = Bible_print_string("多","提多書", chap_string, sec_string,"rcuv")
						}
				}
			case "Philem","Philemon","門","腓利","腓利門","腓利門書","Phm","phm","Послание к Филимону","Phi-lê-môn","ピレモンヘの手紙","ピレモン","빌레몬서":
				bible_short_name = "門"
				switch chap_string {
					case "":
						print_string = "腓利門書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("門","腓利門書", chap_string, "1","rcuv")
							default:
								print_string = Bible_print_string("門","腓利門書", chap_string, sec_string,"rcuv")
						}
				}
			case "Heb","Hebrews","來","希伯來","希伯來書","heb","Послание к Евреям","Hê-bơ-rơ","ヘブル人への手紙","ヘブル","히브리서":
				bible_short_name = "來"
				switch chap_string {
					case "":
						print_string = "希伯來書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("來","希伯來書", chap_string, "1","rcuv")
							default:
								print_string = Bible_print_string("來","希伯來書", chap_string, sec_string,"rcuv")
						}
				}
			case "James","雅","雅各","雅各書","Jas","jas","Послание Иакова","Gia-cơ","ヤコブの手紙","ヤコブ","야고보서":
				bible_short_name = "雅"
				switch chap_string {
					case "":
						print_string = "雅各書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("雅","雅各書", chap_string, "1","rcuv")
							default:
								print_string = Bible_print_string("雅","雅各書", chap_string, sec_string,"rcuv")
						}
				}
			case "1 Pet","First Peter","彼前","彼得前","彼得前書","1Pe","1pe","Первое послание Петра","1 Phi-e-rơ","ペテロの第一の手紙","ペテロ一","베드로전서":
				bible_short_name = "彼前"
				switch chap_string {
					case "":
						print_string = "彼得前書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("彼前","彼得前書", chap_string, "1","rcuv")
							default:
								print_string = Bible_print_string("彼前","彼得前書", chap_string, sec_string,"rcuv")
						}
				}
			case "2 Pet","Second Peter","彼後","彼得後","彼得後書","2Pe","2pe","Второе послание Петра","2 Phi-e-rơ","ペテロの第二の手紙","ペテロ","베드로후서":
				bible_short_name = "彼後"
				switch chap_string {
					case "":
						print_string = "彼得後書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("彼後","彼得後書", chap_string, "1","rcuv")
							default:
								print_string = Bible_print_string("彼後","彼得後書", chap_string, sec_string,"rcuv")
						}
				}
			case "1 John","First John","約一","約翰一書","約翰1","約翰1書","1Jo","1jo","Первое послание Иоанна","1 Giăng","ヨハネの第一の手紙","ヨハネ一","요한일서":
				bible_short_name = "約一"
				switch chap_string {
					case "":
						print_string = "約翰一書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("約一","約翰一書", chap_string, "1","rcuv")
							default:
								print_string = Bible_print_string("約一","約翰一書", chap_string, sec_string,"rcuv")
						}
				}
			case "2 John","second John","約二","約翰二書","約翰2","約翰2書","2Jo","Второе послание Иоанна","2 Giăng","ヨハネの第二の手紙","ヨハネ二","요한2서":
				bible_short_name = "約二"
				switch chap_string {
					case "":
						print_string = "約翰二書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("約二","約翰二書", chap_string, "1","rcuv")
							default:
								print_string = Bible_print_string("約二","約翰二書", chap_string, sec_string,"rcuv")
						}
				}
			case "3 John","Third John","約三","約翰三書","約翰3","約翰3書","3Jo","3jo","Третье послание Иоанна","3 Giăng","ヨハネの第三の手紙","ヨハネ三","요한3서":
				bible_short_name = "約三"
				switch chap_string {
					case "":
						print_string = "約翰三書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("約三","約翰三書", chap_string, "1","rcuv")
							default:
								print_string = Bible_print_string("約三","約翰三書", chap_string, sec_string,"rcuv")
						}
				}
			case "Jude","猶","猶大","猶大書","jude","Послание Иуды","Giu-đe","ユダの手紙","ユダ","유다서":
				bible_short_name = "猶"
				switch chap_string {
					case "":
						print_string = "猶大書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("猶","猶大書", chap_string, "1","rcuv")
							default:
								print_string = Bible_print_string("猶","猶大書", chap_string, sec_string,"rcuv")
						}
				}
			case "Rev","Revelation","啟","啟示","啟示錄","Re","re","ｒｅ","Ｒｅ","rev","Откровение ап. Иоанна Богослова (Апокалипсис)","Khải-huyền","ヨハネの黙示録","黙示録","요한계시록":
				bible_short_name = "啟"
				switch chap_string {
					case "":
						print_string = "啟示錄"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("啟","啟示錄", chap_string, "1","rcuv")
							default:
								print_string = Bible_print_string("啟","啟示錄", chap_string, sec_string,"rcuv")
						}
				}
			default:
				print_string = "聖經"
				//print_string = "你是要找 " +  reg.ReplaceAllString(text, "$3") + " 對嗎？\n對不起，我還沒學呢...\n"
		}
	case "俄文聖經":
		log.Print(reg.ReplaceAllString(text, "$3"))
		switch reg.ReplaceAllString(text, "$3") {
			case "Gen","Genesis","創","創世","創世紀","創世記","Ge","ge","gen","창세기","Sáng-thế Ký","Бытие":
				bible_short_name = "創"
				switch chap_string {
					case "":
						print_string = "創世記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("創","創世記", chap_string, "1","russian")
							default:
								print_string = Bible_print_string("創","創世記", chap_string, sec_string,"russian")
						}
				}
			case "ex","Ex","Exodus","埃及","出","出埃及","出埃及記","출애굽기","エジプト","出エジプト","出エジプト記","Xuất Ê-díp-tô Ký","Исход":
				bible_short_name = "出"
				switch chap_string {
					case "":
						print_string = "出埃及記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("出","出埃及記", chap_string, "1","russian")
							default:
								print_string = Bible_print_string("出","出埃及記", chap_string, sec_string,"russian")
						}
				}
			case "Lev","Leviticus","利","利未","利未記","Le","le","Левит","Lê-vi Ký","レビ記","レビ","레위기":
				bible_short_name = "利"
				switch chap_string {
					case "":
						print_string = "利未記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("利","利未記", chap_string, "1","russian")
							default:
								print_string = Bible_print_string("利","利未記", chap_string, sec_string,"russian")
						}
				}
			case "Num","Numbers","民","民數","民數記","Nu","nu","민수기","民数","民数記","Dân-số Ký","Числа":
				bible_short_name = "民"
				switch chap_string {
					case "":
						print_string = "民數記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("民","民數記", chap_string, "1","russian")
							default:
								print_string = Bible_print_string("民","民數記", chap_string, sec_string,"russian")
						}
				}
			case "Deut","Deuteronomy","申","申命","申命記","De","de","신명기","Phục-truyền Luật-lệ Ký","Второзаконие":
				bible_short_name = "申"
				switch chap_string {
					case "":
						print_string = "申命記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("申","申命記", chap_string, "1","russian")
							default:
								print_string = Bible_print_string("申","申命記", chap_string, sec_string,"russian")
						}
				}
			case "Josh","Joshua","約書亞","約書亞記","Jos","jos","여호수아","ヨシュア記","ヨシュア","Giô-suê","Книга Иисуса Навина":
				bible_short_name = "書"
				switch chap_string {
					case "":
						print_string = "約書亞記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("書","約書亞記", chap_string, "1","russian")
							default:
								print_string = Bible_print_string("書","約書亞記", chap_string, sec_string,"russian")
						}
				}
			case "Judg","Judges","士","士師","士師記","Jud","jud","jdg","Jdg","Книга Судей израилевых","Các Quan Xét","사사기":
				bible_short_name = "士"
				switch chap_string {
					case "":
						print_string = "士師記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("士","士師記", chap_string, "1","russian")
							default:
								print_string = Bible_print_string("士","士師記", chap_string, sec_string,"russian")
						}
				}
			case "Ruth","路得","路得記","Ru","ru","Rut","rut","룻기","ルツ","ルツ記","Ru-tơ","Книга Руфи":
				bible_short_name = "得"
				switch chap_string {
					case "":
						print_string = "路得記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("得","路得記", chap_string, "1","russian")
							default:
								print_string = Bible_print_string("得","路得記", chap_string, sec_string,"russian")
						}
				}
			case "1 Sam","First Samuel","撒上","撒母耳記上","1Sa","1sa","サムエル記上","サムエル上","サム上","사무엘상","1 Sa-mu-ên","Первая книга Царств":
				bible_short_name = "撒上"
				switch chap_string {
					case "":
						print_string = "撒母耳記上"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("撒上","撒母耳記上", chap_string, "1","russian")
							default:
								print_string = Bible_print_string("撒上","撒母耳記上", chap_string, sec_string,"russian")
						}
				}
			case "2 Sam","Second Samuel","撒下","撒母耳記下","2Sa","2sa","사무엘하","サムエル記下","サムエル下","サム下","2 Sa-mu-ên","Вторая книга Царств":
				bible_short_name = "撒下"
				switch chap_string {
					case "":
						print_string = "撒母耳記下"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("撒下","撒母耳記下", chap_string, "1","russian")
							default:
								print_string = Bible_print_string("撒下","撒母耳記下", chap_string, sec_string,"russian")
						}
				}
			case "1 Kin","First Kings","王上","列王上","列王紀上","列王記上","1Ki","1ki","열왕기상","Третья книга Царств","1 Các Vua":
				bible_short_name = "王上"
				switch chap_string {
					case "":
						print_string = "列王紀上"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("王上","列王紀上", chap_string, "1","russian")
							default:
								print_string = Bible_print_string("王上","列王紀上", chap_string, sec_string,"russian")
						}
				}
			case "2 Kin","Second Kings","王下","列王下","列王記下","列王紀下","2Ki","2ki","열왕기하","2 Các Vua","Четвертая книга Царств":
				bible_short_name = "王下"
				switch chap_string {
					case "":
						print_string = "列王紀下"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("王下","列王紀下", chap_string, "1","russian")
							default:
								print_string = Bible_print_string("王下","列王紀下", chap_string, sec_string,"russian")
						}
				}
			case "1 Chr","First Chronicles","歷上","代上","歷代志上","歷代上","1Ch","1ch","Первая книга Паралипоменон","1 Sử-ký","歴上","歴代上","歴代志上","역대상":
				bible_short_name = "代上"
				switch chap_string {
					case "":
						print_string = "歷代志上"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("代上","歷代志上", chap_string, "1","russian")
							default:
								print_string = Bible_print_string("代上","歷代志上", chap_string, sec_string,"russian")
						}
				}
			case "2 Chr","Second Chronicles","代下","歷下","歷代下","歷代志下","2Ch","2ch","역대하","歴代志下","歴代下","歴下","2 Sử-ký","Вторая книга Паралипоменон":
				bible_short_name = "代下"
				switch chap_string {
					case "":
						print_string = "歷代志下"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("代下","歷代志下", chap_string, "1","russian")
							default:
								print_string = Bible_print_string("代下","歷代志下", chap_string, sec_string,"russian")
						}
				}
			case "Ezra","拉","以斯拉","以斯拉記","Ezr","ezr","Первая книга Ездры","Ê-xơ-ra","エズラ","エズラ記","에스라":
				bible_short_name = "拉"
				switch chap_string {
					case "":
						print_string = "以斯拉記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("拉","以斯拉記", chap_string, "1","russian")
							default:
								print_string = Bible_print_string("拉","以斯拉記", chap_string, sec_string,"russian")
						}
				}
			case "Neh","Nehemiah","尼","尼希米","尼希米記","Ne","ne","느헤미야","ネヘミヤ書","ネヘミヤ","Nê-hê-mi","Книга Неемии":
				bible_short_name = "尼"
				switch chap_string {
					case "":
						print_string = "尼希米記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("尼","尼希米記", chap_string, "1","russian")
							default:
								print_string = Bible_print_string("尼","尼希米記", chap_string, sec_string,"russian")
						}
				}
			case "Esth","Esther","斯","以斯帖","以斯帖記","Es","est","Есфирь","Ê-xơ-tê","エステル","エステル記","에스더":
				bible_short_name = "斯"
				switch chap_string {
					case "":
						print_string = "以斯帖記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("斯","以斯帖記", chap_string, "1","russian")
							default:
								print_string = Bible_print_string("斯","以斯帖記", chap_string, sec_string,"russian")
						}
				}
			case "Job","job","伯","約伯","約伯記","Книга Иова","Gióp","ヨブ","ヨブ記","욥기":
				bible_short_name = "伯"
				switch chap_string {
					case "":
						print_string = "約伯記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("伯","約伯記", chap_string, "1","russian")
							default:
								print_string = Bible_print_string("伯","約伯記", chap_string, sec_string,"russian")
						}
				}
			case "Ps","Psalms","詩","詩篇","ps","시편","Thi-thiên","Псалтирь":
				bible_short_name = "詩"
				switch chap_string {
					case "":
						print_string = "詩篇"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("詩","詩篇", chap_string, "1","russian")
							default:
								print_string = Bible_print_string("詩","詩篇", chap_string, sec_string,"russian")
						}
				}
			case "Prov","Proverbs","箴","箴言","Pr","pr","Притчи Соломона","Châm-ngôn","잠언":
				bible_short_name = "箴"
				switch chap_string {
					case "":
						print_string = "箴言"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("箴","箴言", chap_string, "1","russian")
							default:
								print_string = Bible_print_string("箴","箴言", chap_string, sec_string,"russian")
						}
				}
			case "Eccl","Ecclesiastes","傳","傳道","傳道書","Ec","ec","Книга Екклезиаста","или Проповедника","Truyền-đạo","伝道の書","伝道","伝","伝道書","전도서":
				bible_short_name = "傳"
				switch chap_string {
					case "":
						print_string = "傳道書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("傳","傳道書", chap_string, "1","russian")
							default:
								print_string = Bible_print_string("傳","傳道書", chap_string, sec_string,"russian")
						}
				}
			case "Song","Song of Solomon","歌","雅歌","So","so","sng","Sng","Песнь песней Соломона","Nhã-ca","아가":
				bible_short_name = "歌"
				switch chap_string {
					case "":
						print_string = "雅歌"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("歌","雅歌", chap_string, "1","russian")
							default:
								print_string = Bible_print_string("歌","雅歌", chap_string, sec_string,"russian")
						}
				}
			case "Is","Isaiah","賽","以賽","以賽亞","以賽亞書","Isa","isa","Книга пророка Исаии","Ê-sai","イザヤ書","イザヤ","이사야":
				bible_short_name = "賽"
				switch chap_string {
					case "":
						print_string = "以賽亞書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("賽","以賽亞書", chap_string, "1","russian")
							default:
								print_string = Bible_print_string("賽","以賽亞書", chap_string, sec_string,"russian")
						}
				}
			case "Jer","Jeremiah","耶","耶利米","耶利米書","jer","예레미야","エレミヤ","エレミヤ書","Giê-rê-mi","Книга пророка Иеремии":
				bible_short_name = "耶"
				switch chap_string {
					case "":
						print_string = "耶利米書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("耶","耶利米書", chap_string, "1","russian")
							default:
								print_string = Bible_print_string("耶","耶利米書", chap_string, sec_string,"russian")
						}
				}
			case "Lam","Lamentations","哀","哀歌","耶利米哀歌","La","lam","예레미야애가","Ca-thương","Плач Иеремии":
				bible_short_name = "哀"
				switch chap_string {
					case "":
						print_string = "耶利米哀歌"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("哀","耶利米哀歌", chap_string, "1","russian")
							default:
								print_string = Bible_print_string("哀","耶利米哀歌", chap_string, sec_string,"russian")
						}
				}
			case "Ezek","Ezekiel","結","以西結","以西結書","Eze","eze","에스겔","エゼキエル書","エゼキエル","Ê-xê-chi-ên","Книга пророка Иезекииля":
				bible_short_name = "結"
				switch chap_string {
					case "":
						print_string = "以西結書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("結","以西結書", chap_string, "1","russian")
							default:
								print_string = Bible_print_string("結","以西結書", chap_string, sec_string,"russian")
						}
				}
			case "Dan","Daniel","但","但以理","但以理書","Da","da","Книга пророка Даниила","Đa-ni-ên","ダニエル書","ダニエル","다니엘":
				bible_short_name = "但"
				switch chap_string {
					case "":
						print_string = "但以理書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("但","但以理書", chap_string, "1","russian")
							default:
								print_string = Bible_print_string("但","但以理書", chap_string, sec_string,"russian")
						}
				}
			case "Hos","Hosea","何","何西","何西阿","何西阿書","Ho","ho","Книга пророка Осии","Ô-sê","ホセア書","ホセア","호세아":
				bible_short_name = "何"
				switch chap_string {
					case "":
						print_string = "何西阿書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("何","何西阿書", chap_string, "1","russian")
							default:
								print_string = Bible_print_string("何","何西阿書", chap_string, sec_string,"russian")
						}
				}
			case "Joel","珥","約珥","約珥書","Joe","joe","Книга пророка Иоиля","Giô-ên","ヨエル書","ヨエル","요엘":
				bible_short_name = "珥"
				switch chap_string {
					case "":
						print_string = "約珥書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("珥","約珥書", chap_string, "1","russian")
							default:
								print_string = Bible_print_string("珥","約珥書", chap_string, sec_string,"russian")
						}
				}
			case "Amos","摩","阿摩司書","Am","am","Книга пророка Амоса","A-mốt","アモス書","アモス","아모스":
				bible_short_name = "摩"
				switch chap_string {
					case "":
						print_string = "阿摩司書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("摩","阿摩司書", chap_string, "1","russian")
							default:
								print_string = Bible_print_string("摩","阿摩司書", chap_string, sec_string,"russian")
						}
				}
			case "Obad","Obadiah","俄","俄巴底亞","俄巴底亞書","Ob","ob","오바댜","オバデヤ書","オバデヤ","Áp-đia","Книга пророка Авдия":
				bible_short_name = "俄"
				switch chap_string {
					case "":
						print_string = "俄巴底亞書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("俄","俄巴底亞書", chap_string, "1","russian")
							default:
								print_string = Bible_print_string("俄","俄巴底亞書", chap_string, sec_string,"russian")
						}
				}
			case "Jon","Jonah","拿","約拿","約拿書","jon","요나","ヨナ書","ヨナ","Giô-na","Книга пророка Ионы":
				bible_short_name = "拿"
				switch chap_string {
					case "":
						print_string = "約拿書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("拿","約拿書", chap_string, "1","russian")
							default:
								print_string = Bible_print_string("拿","約拿書", chap_string, sec_string,"russian")
						}
				}
			case "Micah","彌","彌迦","彌迦書","Mic","mic","Книга пророка Михея","Mi-chê","ミカ書","ミカ","미가":
				bible_short_name = "彌"
				switch chap_string {
					case "":
						print_string = "彌迦書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("彌","彌迦書", chap_string, "1","russian")
							default:
								print_string = Bible_print_string("彌","彌迦書", chap_string, sec_string,"russian")
						}
				}
			case "Nah","Nahum","鴻","那鴻","那鴻書","Na","na","Книга пророка Наума","Na-hum","ナホム書","ナホム","나훔":
				bible_short_name = "鴻"
				switch chap_string {
					case "":
						print_string = "那鴻書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("鴻","那鴻書", chap_string, "1","russian")
							default:
								print_string = Bible_print_string("鴻","那鴻書", chap_string, sec_string,"russian")
						}
				}
			case "Habakkuk","哈","哈巴","哈巴谷","哈巴谷書","Hab","hab","Книга пророка Аввакума","Ha-ba-cúc","ハバクク書","ハバクク","ハバ","クク","ハバ書","하박국":
				bible_short_name = "哈"
				switch chap_string {
					case "":
						print_string = "哈巴谷書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("哈","哈巴谷書", chap_string, "1","russian")
							default:
								print_string = Bible_print_string("哈","哈巴谷書", chap_string, sec_string,"russian")
						}
				}
			case "Zeph","Zephaniah","番","西番雅","西番雅書","Zep","zep","스바냐","ゼパニヤ書","ゼパニヤ","Sô-phô-ni","Книга пророка Софонии":
				bible_short_name = "番"
				switch chap_string {
					case "":
						print_string = "西番雅書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("番","西番雅書", chap_string, "1","russian")
							default:
								print_string = Bible_print_string("番","西番雅書", chap_string, sec_string,"russian")
						}
				}
			case "Haggai","該","哈該","哈該書","Hag","hag","학개","ハガイ書","ハガイ","A-ghê","Книга пророка Аггея":
				bible_short_name = "該"
				switch chap_string {
					case "":
						print_string = "哈該書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("該","哈該書", chap_string, "1","russian")
							default:
								print_string = Bible_print_string("該","哈該書", chap_string, sec_string,"russian")
						}
				}
			case "Zech","Zechariah","亞","撒迦利亞","撒迦利亞書","Zec","zec","Книга пророка Захарии","Xa-cha-ri","스가랴","ゼカリヤ書","ゼカリヤ":
				bible_short_name = "亞"
				switch chap_string {
					case "":
						print_string = "撒迦利亞書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("亞","撒迦利亞書", chap_string, "1","russian")
							default:
								print_string = Bible_print_string("亞","撒迦利亞書", chap_string, sec_string,"russian")
						}
				}
			case "Malachi","瑪","瑪拉","瑪拉基","瑪拉基書","Mal","mal","말라기","マラキ書","マラキ","Ma-la-chi","Книга пророка Малахии":
				bible_short_name = "瑪"
				switch chap_string {
					case "":
						print_string = "瑪拉基書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("瑪","瑪拉基書", chap_string, "1","russian")
							default:
								print_string = Bible_print_string("瑪","瑪拉基書", chap_string, sec_string,"russian")
						}
				}
			case "Matt","Matthew","太","馬太","馬太福音","Mt","mt","마태복음","マタイによる福音書","マタイ","マタイによる","Ma-thi-ơ","От Матфея святое благовествование":
				bible_short_name = "太"
				switch chap_string {
					case "":
						print_string = "馬太福音"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("太","馬太福音", chap_string, "1","russian")
							default:
								print_string = Bible_print_string("太","馬太福音", chap_string, sec_string,"russian")
						}
				}
			case "Mark","可","馬可","馬可福音","Mr","mr","マルコによる福音書","マルコ","マルコによる","마가복음","Mác","От Марка святое благовествование":
				bible_short_name = "可"
				switch chap_string {
					case "":
						print_string = "馬可福音"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("可","馬可福音", chap_string, "1","russian")
							default:
								print_string = Bible_print_string("可","馬可福音", chap_string, sec_string,"russian")
						}
				}
			case "Luke","路","路加","路加福音","Lu","lu","От Луки святое благовествование","Lu-ca","ルカによる福音書","ルカ","ルカによる","누가복음":
				bible_short_name = "路"
				switch chap_string {
					case "":
						print_string = "路加福音"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("路","路加福音", chap_string, "1","russian")
							default:
								print_string = Bible_print_string("路","路加福音", chap_string, sec_string,"russian")
						}
				}
			case "John","約","約翰","約翰福音","Joh","joh","От Иоанна святое благовествование","Giăng","ヨハネによる福音書","ヨハネ","ヨハネによる","요한복음":
				bible_short_name = "約"
				switch chap_string {
					case "":
						print_string = "約翰福音"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("約","約翰福音", chap_string, "1","russian")
							default:
								print_string = Bible_print_string("約","約翰福音", chap_string, sec_string,"russian")
						}
				}
			case "Acts","徒","使徒","使徒行傳","Ac","ac","Деяния святых апостолов","Công-vụ Các Sứ-đồ","使徒行伝","사도행전":
				bible_short_name = "徒"
				switch chap_string {
					case "":
						print_string = "使徒行傳"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("徒","使徒行傳", chap_string, "1","russian")
							default:
								print_string = Bible_print_string("徒","使徒行傳", chap_string, sec_string,"russian")
						}
				}
			case "Rom","Romans","羅","羅馬","羅馬書","Ro","ro","Послание к Римлянам","Rô-ma","ローマ","ローマ人への手紙","로마서":
				bible_short_name = "羅"
				switch chap_string {
					case "":
						print_string = "羅馬書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("羅","羅馬書", chap_string, "1","russian")
							default:
								print_string = Bible_print_string("羅","羅馬書", chap_string, sec_string,"russian")
						}
				}
			case "1 Cor","First Corinthians","林前","哥林多前","哥林多前書","1Co","1co","Первое послание к Коринфянам","1 Cô-rinh-tô","コリント人への第一の手紙","コリント一","コリント人への第一","고린도전서":
				bible_short_name = "林前"
				switch chap_string {
					case "":
						print_string = "哥林多前書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("林前","哥林多前書", chap_string, "1","russian")
							default:
								print_string = Bible_print_string("林前","哥林多前書", chap_string, sec_string,"russian")
						}
				}
			case "2 Cor","Second Corinthians","林後","哥林多後","哥林多後書","2Co","2co","Второе послание к Коринфянам","2 Cô-rinh-tô","コリント人への第二の手紙","コリント二","コリント人への第二の","고린도후서":
				bible_short_name = "林後"
				switch chap_string {
					case "":
						print_string = "哥林多後書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("林後","哥林多後書", chap_string, "1","russian")
							default:
								print_string = Bible_print_string("林後","哥林多後書", chap_string, sec_string,"russian")
						}
				}
			case "Gal","Galatians","加","加拉太","加拉太書","Ga","ga","Послание к Галатам","Ga-la-ti","ガラテヤ","ガラテヤ人への手紙","갈라디아서":
				bible_short_name = "加"
				switch chap_string {
					case "":
						print_string = "加拉太書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("加","加拉太書", chap_string, "1","russian")
							default:
								print_string = Bible_print_string("加","加拉太書", chap_string, sec_string,"russian")
						}
				}
			case "Ephesians","弗","以弗所","以弗所書","Eph","eph","Послание к Ефесянам","Ê-phê-sô","エペソ人への手紙","エペソ","エペソ人","エペソ人の手紙","에베소서":
				bible_short_name = "弗"
				switch chap_string {
					case "":
						print_string = "以弗所書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("弗","以弗所書", chap_string, "1","russian")
							default:
								print_string = Bible_print_string("弗","以弗所書", chap_string, sec_string,"russian")
						}
				}
			case "Phil","Philippians","腓","腓立","腓立比","腓立比書","Php","php","빌립보서","ピリピ","ピリピ人.ピリピ人への手紙","Послание к Филиппийцам","Phi-líp":
				bible_short_name = "腓"
				switch chap_string {
					case "":
						print_string = "腓立比書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("腓","腓立比書", chap_string, "1","russian")
							default:
								print_string = Bible_print_string("腓","腓立比書", chap_string, sec_string,"russian")
						}
				}
			case "Col","col","Colossians","西","歌羅西","歌羅","歌羅西書","Послание к Колоссянам","Cô-lô-se","コロサイ人への手紙","コロサイ","コロ","골로새서":
				bible_short_name = "西"
				switch chap_string {
					case "":
						print_string = "歌羅西書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("西","歌羅西書", chap_string, "1","russian")
							default:
								print_string = Bible_print_string("西","歌羅西書", chap_string, sec_string,"russian")
						}
				}
			case "1 Thess","First Thessalonians","帖前","帖撒羅尼迦前","帖撒羅尼迦前書","1Th","1th","데살로니가전서","テサロニケ人への第一の手紙","テサ一","テサロニケ一","1 Tê-sa-lô-ni-ca","Первое послание к Фессалоникийцам (Солунянам)":
				bible_short_name = "帖前"
				switch chap_string {
					case "":
						print_string = "帖撒羅尼迦前書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("帖前","帖撒羅尼迦前書", chap_string, "1","russian")
							default:
								print_string = Bible_print_string("帖前","帖撒羅尼迦前書", chap_string, sec_string,"russian")
						}
				}
			case "2 Thess","Second Thessalonians","帖後","帖撒羅尼迦後","帖撒羅尼迦後書","2Th","2th","데살로니가후서","テサロニケ人への第二の手紙","テサ二","テサロニケ二","2 Tê-sa-lô-ni-ca","Второе послание к Фессалоникийцам (Солунянам)":
				bible_short_name = "帖後"
				switch chap_string {
					case "":
						print_string = "帖撒羅尼迦後書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("帖後","帖撒羅尼迦後書", chap_string, "1","russian")
							default:
								print_string = Bible_print_string("帖後","帖撒羅尼迦後書", chap_string, sec_string,"russian")
						}
				}
			case "1 Tim","First Timothy","提前","提摩太前","提摩太前書","1Ti","1ti","Первое послание к Тимофею","1 Ti-mô-thê","テモテヘの第一の手紙","テモテ一","디모데전서":
				bible_short_name = "提前"
				switch chap_string {
					case "":
						print_string = "提摩太前書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("提前","提摩太前書", chap_string, "1","russian")
							default:
								print_string = Bible_print_string("提前","提摩太前書", chap_string, sec_string,"russian")
						}
				}
			case "2 Tim","Second Timothy","提後","提摩太後","提摩太後書","2Ti","2ti","Второе послание к Тимофею","2 Ti-mô-thê","テモテヘの第二の手紙","テモテ二","디모데후서":
				bible_short_name = "提後"
				switch chap_string {
					case "":
						print_string = "提摩太後書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("提後","提摩太後書", chap_string, "1","russian")
							default:
								print_string = Bible_print_string("提後","提摩太後書", chap_string, sec_string,"russian")
						}
				}
			case "Titus","多","提多","提多書","Tit","tit","Послание к Титу","Tít","テトスヘの手紙","テトス","디도서":
				bible_short_name = "多"
				switch chap_string {
					case "":
						print_string = "提多書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("多","提多書", chap_string, "1","russian")
							default:
								print_string = Bible_print_string("多","提多書", chap_string, sec_string,"russian")
						}
				}
			case "Philem","Philemon","門","腓利","腓利門","腓利門書","Phm","phm","Послание к Филимону","Phi-lê-môn","ピレモンヘの手紙","ピレモン","빌레몬서":
				bible_short_name = "門"
				switch chap_string {
					case "":
						print_string = "腓利門書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("門","腓利門書", chap_string, "1","russian")
							default:
								print_string = Bible_print_string("門","腓利門書", chap_string, sec_string,"russian")
						}
				}
			case "Heb","Hebrews","來","希伯來","希伯來書","heb","Послание к Евреям","Hê-bơ-rơ","ヘブル人への手紙","ヘブル","히브리서":
				bible_short_name = "來"
				switch chap_string {
					case "":
						print_string = "希伯來書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("來","希伯來書", chap_string, "1","russian")
							default:
								print_string = Bible_print_string("來","希伯來書", chap_string, sec_string,"russian")
						}
				}
			case "James","雅","雅各","雅各書","Jas","jas","Послание Иакова","Gia-cơ","ヤコブの手紙","ヤコブ","야고보서":
				bible_short_name = "雅"
				switch chap_string {
					case "":
						print_string = "雅各書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("雅","雅各書", chap_string, "1","russian")
							default:
								print_string = Bible_print_string("雅","雅各書", chap_string, sec_string,"russian")
						}
				}
			case "1 Pet","First Peter","彼前","彼得前","彼得前書","1Pe","1pe","Первое послание Петра","1 Phi-e-rơ","ペテロの第一の手紙","ペテロ一","베드로전서":
				bible_short_name = "彼前"
				switch chap_string {
					case "":
						print_string = "彼得前書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("彼前","彼得前書", chap_string, "1","russian")
							default:
								print_string = Bible_print_string("彼前","彼得前書", chap_string, sec_string,"russian")
						}
				}
			case "2 Pet","Second Peter","彼後","彼得後","彼得後書","2Pe","2pe","Второе послание Петра","2 Phi-e-rơ","ペテロの第二の手紙","ペテロ","베드로후서":
				bible_short_name = "彼後"
				switch chap_string {
					case "":
						print_string = "彼得後書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("彼後","彼得後書", chap_string, "1","russian")
							default:
								print_string = Bible_print_string("彼後","彼得後書", chap_string, sec_string,"russian")
						}
				}
			case "1 John","First John","約一","約翰一書","約翰1","約翰1書","1Jo","1jo","Первое послание Иоанна","1 Giăng","ヨハネの第一の手紙","ヨハネ一","요한일서":
				bible_short_name = "約一"
				switch chap_string {
					case "":
						print_string = "約翰一書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("約一","約翰一書", chap_string, "1","russian")
							default:
								print_string = Bible_print_string("約一","約翰一書", chap_string, sec_string,"russian")
						}
				}
			case "2 John","second John","約二","約翰二書","約翰2","約翰2書","2Jo","Второе послание Иоанна","2 Giăng","ヨハネの第二の手紙","ヨハネ二","요한2서":
				bible_short_name = "約二"
				switch chap_string {
					case "":
						print_string = "約翰二書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("約二","約翰二書", chap_string, "1","russian")
							default:
								print_string = Bible_print_string("約二","約翰二書", chap_string, sec_string,"russian")
						}
				}
			case "3 John","Third John","約三","約翰三書","約翰3","約翰3書","3Jo","3jo","Третье послание Иоанна","3 Giăng","ヨハネの第三の手紙","ヨハネ三","요한3서":
				bible_short_name = "約三"
				switch chap_string {
					case "":
						print_string = "約翰三書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("約三","約翰三書", chap_string, "1","russian")
							default:
								print_string = Bible_print_string("約三","約翰三書", chap_string, sec_string,"russian")
						}
				}
			case "Jude","猶","猶大","猶大書","jude","Послание Иуды","Giu-đe","ユダの手紙","ユダ","유다서":
				bible_short_name = "猶"
				switch chap_string {
					case "":
						print_string = "猶大書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("猶","猶大書", chap_string, "1","russian")
							default:
								print_string = Bible_print_string("猶","猶大書", chap_string, sec_string,"russian")
						}
				}
			case "Rev","Revelation","啟","啟示","啟示錄","Re","re","ｒｅ","Ｒｅ","rev","Откровение ап. Иоанна Богослова (Апокалипсис)","Khải-huyền","ヨハネの黙示録","黙示録","요한계시록":
				bible_short_name = "啟"
				switch chap_string {
					case "":
						print_string = "啟示錄"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("啟","啟示錄", chap_string, "1","russian")
							default:
								print_string = Bible_print_string("啟","啟示錄", chap_string, sec_string,"russian")
						}
				}
			default:
				//print_string = "你是要找 " +  reg.ReplaceAllString(text, "$3") + " 對嗎？\n對不起，我還沒學呢...\n"
				print_string = "聖經"
		}
	case "越南聖經":
		log.Print(reg.ReplaceAllString(text, "$3"))
		switch reg.ReplaceAllString(text, "$3") {
			case "Gen","Genesis","創","創世","創世紀","創世記","Ge","ge","gen","창세기","Sáng-thế Ký","Бытие":
				bible_short_name = "創"
				switch chap_string {
					case "":
						print_string = "創世記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("創","創世記", chap_string, "1","vietnamese")
							default:
								print_string = Bible_print_string("創","創世記", chap_string, sec_string,"vietnamese")
						}
				}
			case "ex","Ex","Exodus","埃及","出","出埃及","出埃及記","출애굽기","エジプト","出エジプト","出エジプト記","Xuất Ê-díp-tô Ký","Исход":
				bible_short_name = "出"
				switch chap_string {
					case "":
						print_string = "出埃及記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("出","出埃及記", chap_string, "1","vietnamese")
							default:
								print_string = Bible_print_string("出","出埃及記", chap_string, sec_string,"vietnamese")
						}
				}
			case "Lev","Leviticus","利","利未","利未記","Le","le","Левит","Lê-vi Ký","レビ記","レビ","레위기":
				bible_short_name = "利"
				switch chap_string {
					case "":
						print_string = "利未記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("利","利未記", chap_string, "1","vietnamese")
							default:
								print_string = Bible_print_string("利","利未記", chap_string, sec_string,"vietnamese")
						}
				}
			case "Num","Numbers","民","民數","民數記","Nu","nu","민수기","民数","民数記","Dân-số Ký","Числа":
				bible_short_name = "民"
				switch chap_string {
					case "":
						print_string = "民數記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("民","民數記", chap_string, "1","vietnamese")
							default:
								print_string = Bible_print_string("民","民數記", chap_string, sec_string,"vietnamese")
						}
				}
			case "Deut","Deuteronomy","申","申命","申命記","De","de","신명기","Phục-truyền Luật-lệ Ký","Второзаконие":
				bible_short_name = "申"
				switch chap_string {
					case "":
						print_string = "申命記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("申","申命記", chap_string, "1","vietnamese")
							default:
								print_string = Bible_print_string("申","申命記", chap_string, sec_string,"vietnamese")
						}
				}
			case "Josh","Joshua","約書亞","約書亞記","Jos","jos","여호수아","ヨシュア記","ヨシュア","Giô-suê","Книга Иисуса Навина":
				bible_short_name = "書"
				switch chap_string {
					case "":
						print_string = "約書亞記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("書","約書亞記", chap_string, "1","vietnamese")
							default:
								print_string = Bible_print_string("書","約書亞記", chap_string, sec_string,"vietnamese")
						}
				}
			case "Judg","Judges","士","士師","士師記","Jud","jud","jdg","Jdg","Книга Судей израилевых","Các Quan Xét","사사기":
				bible_short_name = "士"
				switch chap_string {
					case "":
						print_string = "士師記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("士","士師記", chap_string, "1","vietnamese")
							default:
								print_string = Bible_print_string("士","士師記", chap_string, sec_string,"vietnamese")
						}
				}
			case "Ruth","路得","路得記","Ru","ru","Rut","rut","룻기","ルツ","ルツ記","Ru-tơ","Книга Руфи":
				bible_short_name = "得"
				switch chap_string {
					case "":
						print_string = "路得記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("得","路得記", chap_string, "1","vietnamese")
							default:
								print_string = Bible_print_string("得","路得記", chap_string, sec_string,"vietnamese")
						}
				}
			case "1 Sam","First Samuel","撒上","撒母耳記上","1Sa","1sa","サムエル記上","サムエル上","サム上","사무엘상","1 Sa-mu-ên","Первая книга Царств":
				bible_short_name = "撒上"
				switch chap_string {
					case "":
						print_string = "撒母耳記上"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("撒上","撒母耳記上", chap_string, "1","vietnamese")
							default:
								print_string = Bible_print_string("撒上","撒母耳記上", chap_string, sec_string,"vietnamese")
						}
				}
			case "2 Sam","Second Samuel","撒下","撒母耳記下","2Sa","2sa","사무엘하","サムエル記下","サムエル下","サム下","2 Sa-mu-ên","Вторая книга Царств":
				bible_short_name = "撒下"
				switch chap_string {
					case "":
						print_string = "撒母耳記下"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("撒下","撒母耳記下", chap_string, "1","vietnamese")
							default:
								print_string = Bible_print_string("撒下","撒母耳記下", chap_string, sec_string,"vietnamese")
						}
				}
			case "1 Kin","First Kings","王上","列王上","列王紀上","列王記上","1Ki","1ki","열왕기상","Третья книга Царств","1 Các Vua":
				bible_short_name = "王上"
				switch chap_string {
					case "":
						print_string = "列王紀上"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("王上","列王紀上", chap_string, "1","vietnamese")
							default:
								print_string = Bible_print_string("王上","列王紀上", chap_string, sec_string,"vietnamese")
						}
				}
			case "2 Kin","Second Kings","王下","列王下","列王記下","列王紀下","2Ki","2ki","열왕기하","2 Các Vua","Четвертая книга Царств":
				bible_short_name = "王下"
				switch chap_string {
					case "":
						print_string = "列王紀下"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("王下","列王紀下", chap_string, "1","vietnamese")
							default:
								print_string = Bible_print_string("王下","列王紀下", chap_string, sec_string,"vietnamese")
						}
				}
			case "1 Chr","First Chronicles","歷上","代上","歷代志上","歷代上","1Ch","1ch","Первая книга Паралипоменон","1 Sử-ký","歴上","歴代上","歴代志上","역대상":
				bible_short_name = "代上"
				switch chap_string {
					case "":
						print_string = "歷代志上"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("代上","歷代志上", chap_string, "1","vietnamese")
							default:
								print_string = Bible_print_string("代上","歷代志上", chap_string, sec_string,"vietnamese")
						}
				}
			case "2 Chr","Second Chronicles","代下","歷下","歷代下","歷代志下","2Ch","2ch","역대하","歴代志下","歴代下","歴下","2 Sử-ký","Вторая книга Паралипоменон":
				bible_short_name = "代下"
				switch chap_string {
					case "":
						print_string = "歷代志下"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("代下","歷代志下", chap_string, "1","vietnamese")
							default:
								print_string = Bible_print_string("代下","歷代志下", chap_string, sec_string,"vietnamese")
						}
				}
			case "Ezra","拉","以斯拉","以斯拉記","Ezr","ezr","Первая книга Ездры","Ê-xơ-ra","エズラ","エズラ記","에스라":
				bible_short_name = "拉"
				switch chap_string {
					case "":
						print_string = "以斯拉記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("拉","以斯拉記", chap_string, "1","vietnamese")
							default:
								print_string = Bible_print_string("拉","以斯拉記", chap_string, sec_string,"vietnamese")
						}
				}
			case "Neh","Nehemiah","尼","尼希米","尼希米記","Ne","ne","느헤미야","ネヘミヤ書","ネヘミヤ","Nê-hê-mi","Книга Неемии":
				bible_short_name = "尼"
				switch chap_string {
					case "":
						print_string = "尼希米記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("尼","尼希米記", chap_string, "1","vietnamese")
							default:
								print_string = Bible_print_string("尼","尼希米記", chap_string, sec_string,"vietnamese")
						}
				}
			case "Esth","Esther","斯","以斯帖","以斯帖記","Es","est","Есфирь","Ê-xơ-tê","エステル","エステル記","에스더":
				bible_short_name = "斯"
				switch chap_string {
					case "":
						print_string = "以斯帖記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("斯","以斯帖記", chap_string, "1","vietnamese")
							default:
								print_string = Bible_print_string("斯","以斯帖記", chap_string, sec_string,"vietnamese")
						}
				}
			case "Job","job","伯","約伯","約伯記","Книга Иова","Gióp","ヨブ","ヨブ記","욥기":
				bible_short_name = "伯"
				switch chap_string {
					case "":
						print_string = "約伯記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("伯","約伯記", chap_string, "1","vietnamese")
							default:
								print_string = Bible_print_string("伯","約伯記", chap_string, sec_string,"vietnamese")
						}
				}
			case "Ps","Psalms","詩","詩篇","ps","시편","Thi-thiên","Псалтирь":
				bible_short_name = "詩"
				switch chap_string {
					case "":
						print_string = "詩篇"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("詩","詩篇", chap_string, "1","vietnamese")
							default:
								print_string = Bible_print_string("詩","詩篇", chap_string, sec_string,"vietnamese")
						}
				}
			case "Prov","Proverbs","箴","箴言","Pr","pr","Притчи Соломона","Châm-ngôn","잠언":
				bible_short_name = "箴"
				switch chap_string {
					case "":
						print_string = "箴言"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("箴","箴言", chap_string, "1","vietnamese")
							default:
								print_string = Bible_print_string("箴","箴言", chap_string, sec_string,"vietnamese")
						}
				}
			case "Eccl","Ecclesiastes","傳","傳道","傳道書","Ec","ec","Книга Екклезиаста","или Проповедника","Truyền-đạo","伝道の書","伝道","伝","伝道書","전도서":
				bible_short_name = "傳"
				switch chap_string {
					case "":
						print_string = "傳道書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("傳","傳道書", chap_string, "1","vietnamese")
							default:
								print_string = Bible_print_string("傳","傳道書", chap_string, sec_string,"vietnamese")
						}
				}
			case "Song","Song of Solomon","歌","雅歌","So","so","sng","Sng","Песнь песней Соломона","Nhã-ca","아가":
				bible_short_name = "歌"
				switch chap_string {
					case "":
						print_string = "雅歌"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("歌","雅歌", chap_string, "1","vietnamese")
							default:
								print_string = Bible_print_string("歌","雅歌", chap_string, sec_string,"vietnamese")
						}
				}
			case "Is","Isaiah","賽","以賽","以賽亞","以賽亞書","Isa","isa","Книга пророка Исаии","Ê-sai","イザヤ書","イザヤ","이사야":
				bible_short_name = "賽"
				switch chap_string {
					case "":
						print_string = "以賽亞書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("賽","以賽亞書", chap_string, "1","vietnamese")
							default:
								print_string = Bible_print_string("賽","以賽亞書", chap_string, sec_string,"vietnamese")
						}
				}
			case "Jer","Jeremiah","耶","耶利米","耶利米書","jer","예레미야","エレミヤ","エレミヤ書","Giê-rê-mi","Книга пророка Иеремии":
				bible_short_name = "耶"
				switch chap_string {
					case "":
						print_string = "耶利米書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("耶","耶利米書", chap_string, "1","vietnamese")
							default:
								print_string = Bible_print_string("耶","耶利米書", chap_string, sec_string,"vietnamese")
						}
				}
			case "Lam","Lamentations","哀","哀歌","耶利米哀歌","La","lam","예레미야애가","Ca-thương","Плач Иеремии":
				bible_short_name = "哀"
				switch chap_string {
					case "":
						print_string = "耶利米哀歌"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("哀","耶利米哀歌", chap_string, "1","vietnamese")
							default:
								print_string = Bible_print_string("哀","耶利米哀歌", chap_string, sec_string,"vietnamese")
						}
				}
			case "Ezek","Ezekiel","結","以西結","以西結書","Eze","eze","에스겔","エゼキエル書","エゼキエル","Ê-xê-chi-ên","Книга пророка Иезекииля":
				bible_short_name = "結"
				switch chap_string {
					case "":
						print_string = "以西結書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("結","以西結書", chap_string, "1","vietnamese")
							default:
								print_string = Bible_print_string("結","以西結書", chap_string, sec_string,"vietnamese")
						}
				}
			case "Dan","Daniel","但","但以理","但以理書","Da","da","Книга пророка Даниила","Đa-ni-ên","ダニエル書","ダニエル","다니엘":
				bible_short_name = "但"
				switch chap_string {
					case "":
						print_string = "但以理書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("但","但以理書", chap_string, "1","vietnamese")
							default:
								print_string = Bible_print_string("但","但以理書", chap_string, sec_string,"vietnamese")
						}
				}
			case "Hos","Hosea","何","何西","何西阿","何西阿書","Ho","ho","Книга пророка Осии","Ô-sê","ホセア書","ホセア","호세아":
				bible_short_name = "何"
				switch chap_string {
					case "":
						print_string = "何西阿書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("何","何西阿書", chap_string, "1","vietnamese")
							default:
								print_string = Bible_print_string("何","何西阿書", chap_string, sec_string,"vietnamese")
						}
				}
			case "Joel","珥","約珥","約珥書","Joe","joe","Книга пророка Иоиля","Giô-ên","ヨエル書","ヨエル","요엘":
				bible_short_name = "珥"
				switch chap_string {
					case "":
						print_string = "約珥書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("珥","約珥書", chap_string, "1","vietnamese")
							default:
								print_string = Bible_print_string("珥","約珥書", chap_string, sec_string,"vietnamese")
						}
				}
			case "Amos","摩","阿摩司書","Am","am","Книга пророка Амоса","A-mốt","アモス書","アモス","아모스":
				bible_short_name = "摩"
				switch chap_string {
					case "":
						print_string = "阿摩司書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("摩","阿摩司書", chap_string, "1","vietnamese")
							default:
								print_string = Bible_print_string("摩","阿摩司書", chap_string, sec_string,"vietnamese")
						}
				}
			case "Obad","Obadiah","俄","俄巴底亞","俄巴底亞書","Ob","ob","오바댜","オバデヤ書","オバデヤ","Áp-đia","Книга пророка Авдия":
				bible_short_name = "俄"
				switch chap_string {
					case "":
						print_string = "俄巴底亞書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("俄","俄巴底亞書", chap_string, "1","vietnamese")
							default:
								print_string = Bible_print_string("俄","俄巴底亞書", chap_string, sec_string,"vietnamese")
						}
				}
			case "Jon","Jonah","拿","約拿","約拿書","jon","요나","ヨナ書","ヨナ","Giô-na","Книга пророка Ионы":
				bible_short_name = "拿"
				switch chap_string {
					case "":
						print_string = "約拿書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("拿","約拿書", chap_string, "1","vietnamese")
							default:
								print_string = Bible_print_string("拿","約拿書", chap_string, sec_string,"vietnamese")
						}
				}
			case "Micah","彌","彌迦","彌迦書","Mic","mic","Книга пророка Михея","Mi-chê","ミカ書","ミカ","미가":
				bible_short_name = "彌"
				switch chap_string {
					case "":
						print_string = "彌迦書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("彌","彌迦書", chap_string, "1","vietnamese")
							default:
								print_string = Bible_print_string("彌","彌迦書", chap_string, sec_string,"vietnamese")
						}
				}
			case "Nah","Nahum","鴻","那鴻","那鴻書","Na","na","Книга пророка Наума","Na-hum","ナホム書","ナホム","나훔":
				bible_short_name = "鴻"
				switch chap_string {
					case "":
						print_string = "那鴻書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("鴻","那鴻書", chap_string, "1","vietnamese")
							default:
								print_string = Bible_print_string("鴻","那鴻書", chap_string, sec_string,"vietnamese")
						}
				}
			case "Habakkuk","哈","哈巴","哈巴谷","哈巴谷書","Hab","hab","Книга пророка Аввакума","Ha-ba-cúc","ハバクク書","ハバクク","ハバ","クク","ハバ書","하박국":
				bible_short_name = "哈"
				switch chap_string {
					case "":
						print_string = "哈巴谷書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("哈","哈巴谷書", chap_string, "1","vietnamese")
							default:
								print_string = Bible_print_string("哈","哈巴谷書", chap_string, sec_string,"vietnamese")
						}
				}
			case "Zeph","Zephaniah","番","西番雅","西番雅書","Zep","zep","스바냐","ゼパニヤ書","ゼパニヤ","Sô-phô-ni","Книга пророка Софонии":
				bible_short_name = "番"
				switch chap_string {
					case "":
						print_string = "西番雅書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("番","西番雅書", chap_string, "1","vietnamese")
							default:
								print_string = Bible_print_string("番","西番雅書", chap_string, sec_string,"vietnamese")
						}
				}
			case "Haggai","該","哈該","哈該書","Hag","hag","학개","ハガイ書","ハガイ","A-ghê","Книга пророка Аггея":
				bible_short_name = "該"
				switch chap_string {
					case "":
						print_string = "哈該書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("該","哈該書", chap_string, "1","vietnamese")
							default:
								print_string = Bible_print_string("該","哈該書", chap_string, sec_string,"vietnamese")
						}
				}
			case "Zech","Zechariah","亞","撒迦利亞","撒迦利亞書","Zec","zec","Книга пророка Захарии","Xa-cha-ri","스가랴","ゼカリヤ書","ゼカリヤ":
				bible_short_name = "亞"
				switch chap_string {
					case "":
						print_string = "撒迦利亞書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("亞","撒迦利亞書", chap_string, "1","vietnamese")
							default:
								print_string = Bible_print_string("亞","撒迦利亞書", chap_string, sec_string,"vietnamese")
						}
				}
			case "Malachi","瑪","瑪拉","瑪拉基","瑪拉基書","Mal","mal","말라기","マラキ書","マラキ","Ma-la-chi","Книга пророка Малахии":
				bible_short_name = "瑪"
				switch chap_string {
					case "":
						print_string = "瑪拉基書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("瑪","瑪拉基書", chap_string, "1","vietnamese")
							default:
								print_string = Bible_print_string("瑪","瑪拉基書", chap_string, sec_string,"vietnamese")
						}
				}
			case "Matt","Matthew","太","馬太","馬太福音","Mt","mt","마태복음","マタイによる福音書","マタイ","マタイによる","Ma-thi-ơ","От Матфея святое благовествование":
				bible_short_name = "太"
				switch chap_string {
					case "":
						print_string = "馬太福音"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("太","馬太福音", chap_string, "1","vietnamese")
							default:
								print_string = Bible_print_string("太","馬太福音", chap_string, sec_string,"vietnamese")
						}
				}
			case "Mark","可","馬可","馬可福音","Mr","mr","マルコによる福音書","マルコ","マルコによる","마가복음","Mác","От Марка святое благовествование":
				bible_short_name = "可"
				switch chap_string {
					case "":
						print_string = "馬可福音"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("可","馬可福音", chap_string, "1","vietnamese")
							default:
								print_string = Bible_print_string("可","馬可福音", chap_string, sec_string,"vietnamese")
						}
				}
			case "Luke","路","路加","路加福音","Lu","lu","От Луки святое благовествование","Lu-ca","ルカによる福音書","ルカ","ルカによる","누가복음":
				bible_short_name = "路"
				switch chap_string {
					case "":
						print_string = "路加福音"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("路","路加福音", chap_string, "1","vietnamese")
							default:
								print_string = Bible_print_string("路","路加福音", chap_string, sec_string,"vietnamese")
						}
				}
			case "John","約","約翰","約翰福音","Joh","joh","От Иоанна святое благовествование","Giăng","ヨハネによる福音書","ヨハネ","ヨハネによる","요한복음":
				bible_short_name = "約"
				switch chap_string {
					case "":
						print_string = "約翰福音"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("約","約翰福音", chap_string, "1","vietnamese")
							default:
								print_string = Bible_print_string("約","約翰福音", chap_string, sec_string,"vietnamese")
						}
				}
			case "Acts","徒","使徒","使徒行傳","Ac","ac","Деяния святых апостолов","Công-vụ Các Sứ-đồ","使徒行伝","사도행전":
				bible_short_name = "徒"
				switch chap_string {
					case "":
						print_string = "使徒行傳"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("徒","使徒行傳", chap_string, "1","vietnamese")
							default:
								print_string = Bible_print_string("徒","使徒行傳", chap_string, sec_string,"vietnamese")
						}
				}
			case "Rom","Romans","羅","羅馬","羅馬書","Ro","ro","Послание к Римлянам","Rô-ma","ローマ","ローマ人への手紙","로마서":
				bible_short_name = "羅"
				switch chap_string {
					case "":
						print_string = "羅馬書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("羅","羅馬書", chap_string, "1","vietnamese")
							default:
								print_string = Bible_print_string("羅","羅馬書", chap_string, sec_string,"vietnamese")
						}
				}
			case "1 Cor","First Corinthians","林前","哥林多前","哥林多前書","1Co","1co","Первое послание к Коринфянам","1 Cô-rinh-tô","コリント人への第一の手紙","コリント一","コリント人への第一","고린도전서":
				bible_short_name = "林前"
				switch chap_string {
					case "":
						print_string = "哥林多前書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("林前","哥林多前書", chap_string, "1","vietnamese")
							default:
								print_string = Bible_print_string("林前","哥林多前書", chap_string, sec_string,"vietnamese")
						}
				}
			case "2 Cor","Second Corinthians","林後","哥林多後","哥林多後書","2Co","2co","Второе послание к Коринфянам","2 Cô-rinh-tô","コリント人への第二の手紙","コリント二","コリント人への第二の","고린도후서":
				bible_short_name = "林後"
				switch chap_string {
					case "":
						print_string = "哥林多後書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("林後","哥林多後書", chap_string, "1","vietnamese")
							default:
								print_string = Bible_print_string("林後","哥林多後書", chap_string, sec_string,"vietnamese")
						}
				}
			case "Gal","Galatians","加","加拉太","加拉太書","Ga","ga","Послание к Галатам","Ga-la-ti","ガラテヤ","ガラテヤ人への手紙","갈라디아서":
				bible_short_name = "加"
				switch chap_string {
					case "":
						print_string = "加拉太書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("加","加拉太書", chap_string, "1","vietnamese")
							default:
								print_string = Bible_print_string("加","加拉太書", chap_string, sec_string,"vietnamese")
						}
				}
			case "Ephesians","弗","以弗所","以弗所書","Eph","eph","Послание к Ефесянам","Ê-phê-sô","エペソ人への手紙","エペソ","エペソ人","エペソ人の手紙","에베소서":
				bible_short_name = "弗"
				switch chap_string {
					case "":
						print_string = "以弗所書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("弗","以弗所書", chap_string, "1","vietnamese")
							default:
								print_string = Bible_print_string("弗","以弗所書", chap_string, sec_string,"vietnamese")
						}
				}
			case "Phil","Philippians","腓","腓立","腓立比","腓立比書","Php","php","빌립보서","ピリピ","ピリピ人.ピリピ人への手紙","Послание к Филиппийцам","Phi-líp":
				bible_short_name = "腓"
				switch chap_string {
					case "":
						print_string = "腓立比書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("腓","腓立比書", chap_string, "1","vietnamese")
							default:
								print_string = Bible_print_string("腓","腓立比書", chap_string, sec_string,"vietnamese")
						}
				}
			case "Col","col","Colossians","西","歌羅西","歌羅","歌羅西書","Послание к Колоссянам","Cô-lô-se","コロサイ人への手紙","コロサイ","コロ","골로새서":
				bible_short_name = "西"
				switch chap_string {
					case "":
						print_string = "歌羅西書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("西","歌羅西書", chap_string, "1","vietnamese")
							default:
								print_string = Bible_print_string("西","歌羅西書", chap_string, sec_string,"vietnamese")
						}
				}
			case "1 Thess","First Thessalonians","帖前","帖撒羅尼迦前","帖撒羅尼迦前書","1Th","1th","데살로니가전서","テサロニケ人への第一の手紙","テサ一","テサロニケ一","1 Tê-sa-lô-ni-ca","Первое послание к Фессалоникийцам (Солунянам)":
				bible_short_name = "帖前"
				switch chap_string {
					case "":
						print_string = "帖撒羅尼迦前書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("帖前","帖撒羅尼迦前書", chap_string, "1","vietnamese")
							default:
								print_string = Bible_print_string("帖前","帖撒羅尼迦前書", chap_string, sec_string,"vietnamese")
						}
				}
			case "2 Thess","Second Thessalonians","帖後","帖撒羅尼迦後","帖撒羅尼迦後書","2Th","2th","데살로니가후서","テサロニケ人への第二の手紙","テサ二","テサロニケ二","2 Tê-sa-lô-ni-ca","Второе послание к Фессалоникийцам (Солунянам)":
				bible_short_name = "帖後"
				switch chap_string {
					case "":
						print_string = "帖撒羅尼迦後書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("帖後","帖撒羅尼迦後書", chap_string, "1","vietnamese")
							default:
								print_string = Bible_print_string("帖後","帖撒羅尼迦後書", chap_string, sec_string,"vietnamese")
						}
				}
			case "1 Tim","First Timothy","提前","提摩太前","提摩太前書","1Ti","1ti","Первое послание к Тимофею","1 Ti-mô-thê","テモテヘの第一の手紙","テモテ一","디모데전서":
				bible_short_name = "提前"
				switch chap_string {
					case "":
						print_string = "提摩太前書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("提前","提摩太前書", chap_string, "1","vietnamese")
							default:
								print_string = Bible_print_string("提前","提摩太前書", chap_string, sec_string,"vietnamese")
						}
				}
			case "2 Tim","Second Timothy","提後","提摩太後","提摩太後書","2Ti","2ti","Второе послание к Тимофею","2 Ti-mô-thê","テモテヘの第二の手紙","テモテ二","디모데후서":
				bible_short_name = "提後"
				switch chap_string {
					case "":
						print_string = "提摩太後書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("提後","提摩太後書", chap_string, "1","vietnamese")
							default:
								print_string = Bible_print_string("提後","提摩太後書", chap_string, sec_string,"vietnamese")
						}
				}
			case "Titus","多","提多","提多書","Tit","tit","Послание к Титу","Tít","テトスヘの手紙","テトス","디도서":
				bible_short_name = "多"
				switch chap_string {
					case "":
						print_string = "提多書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("多","提多書", chap_string, "1","vietnamese")
							default:
								print_string = Bible_print_string("多","提多書", chap_string, sec_string,"vietnamese")
						}
				}
			case "Philem","Philemon","門","腓利","腓利門","腓利門書","Phm","phm","Послание к Филимону","Phi-lê-môn","ピレモンヘの手紙","ピレモン","빌레몬서":
				bible_short_name = "門"
				switch chap_string {
					case "":
						print_string = "腓利門書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("門","腓利門書", chap_string, "1","vietnamese")
							default:
								print_string = Bible_print_string("門","腓利門書", chap_string, sec_string,"vietnamese")
						}
				}
			case "Heb","Hebrews","來","希伯來","希伯來書","heb","Послание к Евреям","Hê-bơ-rơ","ヘブル人への手紙","ヘブル","히브리서":
				bible_short_name = "來"
				switch chap_string {
					case "":
						print_string = "希伯來書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("來","希伯來書", chap_string, "1","vietnamese")
							default:
								print_string = Bible_print_string("來","希伯來書", chap_string, sec_string,"vietnamese")
						}
				}
			case "James","雅","雅各","雅各書","Jas","jas","Послание Иакова","Gia-cơ","ヤコブの手紙","ヤコブ","야고보서":
				bible_short_name = "雅"
				switch chap_string {
					case "":
						print_string = "雅各書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("雅","雅各書", chap_string, "1","vietnamese")
							default:
								print_string = Bible_print_string("雅","雅各書", chap_string, sec_string,"vietnamese")
						}
				}
			case "1 Pet","First Peter","彼前","彼得前","彼得前書","1Pe","1pe","Первое послание Петра","1 Phi-e-rơ","ペテロの第一の手紙","ペテロ一","베드로전서":
				bible_short_name = "彼前"
				switch chap_string {
					case "":
						print_string = "彼得前書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("彼前","彼得前書", chap_string, "1","vietnamese")
							default:
								print_string = Bible_print_string("彼前","彼得前書", chap_string, sec_string,"vietnamese")
						}
				}
			case "2 Pet","Second Peter","彼後","彼得後","彼得後書","2Pe","2pe","Второе послание Петра","2 Phi-e-rơ","ペテロの第二の手紙","ペテロ","베드로후서":
				bible_short_name = "彼後"
				switch chap_string {
					case "":
						print_string = "彼得後書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("彼後","彼得後書", chap_string, "1","vietnamese")
							default:
								print_string = Bible_print_string("彼後","彼得後書", chap_string, sec_string,"vietnamese")
						}
				}
			case "1 John","First John","約一","約翰一書","約翰1","約翰1書","1Jo","1jo","Первое послание Иоанна","1 Giăng","ヨハネの第一の手紙","ヨハネ一","요한일서":
				bible_short_name = "約一"
				switch chap_string {
					case "":
						print_string = "約翰一書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("約一","約翰一書", chap_string, "1","vietnamese")
							default:
								print_string = Bible_print_string("約一","約翰一書", chap_string, sec_string,"vietnamese")
						}
				}
			case "2 John","second John","約二","約翰二書","約翰2","約翰2書","2Jo","Второе послание Иоанна","2 Giăng","ヨハネの第二の手紙","ヨハネ二","요한2서":
				bible_short_name = "約二"
				switch chap_string {
					case "":
						print_string = "約翰二書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("約二","約翰二書", chap_string, "1","vietnamese")
							default:
								print_string = Bible_print_string("約二","約翰二書", chap_string, sec_string,"vietnamese")
						}
				}
			case "3 John","Third John","約三","約翰三書","約翰3","約翰3書","3Jo","3jo","Третье послание Иоанна","3 Giăng","ヨハネの第三の手紙","ヨハネ三","요한3서":
				bible_short_name = "約三"
				switch chap_string {
					case "":
						print_string = "約翰三書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("約三","約翰三書", chap_string, "1","vietnamese")
							default:
								print_string = Bible_print_string("約三","約翰三書", chap_string, sec_string,"vietnamese")
						}
				}
			case "Jude","猶","猶大","猶大書","jude","Послание Иуды","Giu-đe","ユダの手紙","ユダ","유다서":
				bible_short_name = "猶"
				switch chap_string {
					case "":
						print_string = "猶大書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("猶","猶大書", chap_string, "1","vietnamese")
							default:
								print_string = Bible_print_string("猶","猶大書", chap_string, sec_string,"vietnamese")
						}
				}
			case "Rev","Revelation","啟","啟示","啟示錄","Re","re","ｒｅ","Ｒｅ","rev","Откровение ап. Иоанна Богослова (Апокалипсис)","Khải-huyền","ヨハネの黙示録","黙示録","요한계시록":
				bible_short_name = "啟"
				switch chap_string {
					case "":
						print_string = "啟示錄"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("啟","啟示錄", chap_string, "1","vietnamese")
							default:
								print_string = Bible_print_string("啟","啟示錄", chap_string, sec_string,"vietnamese")
						}
				}
			default:
				//print_string = "你是要找 " +  reg.ReplaceAllString(text, "$3") + " 對嗎？\n對不起，我還沒學呢...\n"
				print_string = "聖經"
		}
	case "韓文聖經","KR bible","Korean","korean","Kr Bible","Kr bible":
		log.Print(reg.ReplaceAllString(text, "$3"))
		switch reg.ReplaceAllString(text, "$3") {
			case "Gen","Genesis","創","創世","創世紀","創世記","Ge","ge","gen","창세기","Sáng-thế Ký","Бытие":
				bible_short_name = "創"
				switch chap_string {
					case "":
						print_string = "創世記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("創","創世記", chap_string, "1","korean")
							default:
								print_string = Bible_print_string("創","創世記", chap_string, sec_string,"korean")
						}
				}
			case "ex","Ex","Exodus","埃及","出","出埃及","出埃及記","출애굽기","エジプト","出エジプト","出エジプト記","Xuất Ê-díp-tô Ký","Исход":
				bible_short_name = "出"
				switch chap_string {
					case "":
						print_string = "出埃及記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("出","出埃及記", chap_string, "1","korean")
							default:
								print_string = Bible_print_string("出","出埃及記", chap_string, sec_string,"korean")
						}
				}
			case "Lev","Leviticus","利","利未","利未記","Le","le","Левит","Lê-vi Ký","レビ記","レビ","레위기":
				bible_short_name = "利"
				switch chap_string {
					case "":
						print_string = "利未記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("利","利未記", chap_string, "1","korean")
							default:
								print_string = Bible_print_string("利","利未記", chap_string, sec_string,"korean")
						}
				}
			case "Num","Numbers","民","民數","民數記","Nu","nu","민수기","民数","民数記","Dân-số Ký","Числа":
				bible_short_name = "民"
				switch chap_string {
					case "":
						print_string = "民數記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("民","民數記", chap_string, "1","korean")
							default:
								print_string = Bible_print_string("民","民數記", chap_string, sec_string,"korean")
						}
				}
			case "Deut","Deuteronomy","申","申命","申命記","De","de","신명기","Phục-truyền Luật-lệ Ký","Второзаконие":
				bible_short_name = "申"
				switch chap_string {
					case "":
						print_string = "申命記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("申","申命記", chap_string, "1","korean")
							default:
								print_string = Bible_print_string("申","申命記", chap_string, sec_string,"korean")
						}
				}
			case "Josh","Joshua","約書亞","約書亞記","Jos","jos","여호수아","ヨシュア記","ヨシュア","Giô-suê","Книга Иисуса Навина":
				bible_short_name = "書"
				switch chap_string {
					case "":
						print_string = "約書亞記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("書","約書亞記", chap_string, "1","korean")
							default:
								print_string = Bible_print_string("書","約書亞記", chap_string, sec_string,"korean")
						}
				}
			case "Judg","Judges","士","士師","士師記","Jud","jud","jdg","Jdg","Книга Судей израилевых","Các Quan Xét","사사기":
				bible_short_name = "士"
				switch chap_string {
					case "":
						print_string = "士師記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("士","士師記", chap_string, "1","korean")
							default:
								print_string = Bible_print_string("士","士師記", chap_string, sec_string,"korean")
						}
				}
			case "Ruth","路得","路得記","Ru","ru","Rut","rut","룻기","ルツ","ルツ記","Ru-tơ","Книга Руфи":
				bible_short_name = "得"
				switch chap_string {
					case "":
						print_string = "路得記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("得","路得記", chap_string, "1","korean")
							default:
								print_string = Bible_print_string("得","路得記", chap_string, sec_string,"korean")
						}
				}
			case "1 Sam","First Samuel","撒上","撒母耳記上","1Sa","1sa","サムエル記上","サムエル上","サム上","사무엘상","1 Sa-mu-ên","Первая книга Царств":
				bible_short_name = "撒上"
				switch chap_string {
					case "":
						print_string = "撒母耳記上"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("撒上","撒母耳記上", chap_string, "1","korean")
							default:
								print_string = Bible_print_string("撒上","撒母耳記上", chap_string, sec_string,"korean")
						}
				}
			case "2 Sam","Second Samuel","撒下","撒母耳記下","2Sa","2sa","사무엘하","サムエル記下","サムエル下","サム下","2 Sa-mu-ên","Вторая книга Царств":
				bible_short_name = "撒下"
				switch chap_string {
					case "":
						print_string = "撒母耳記下"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("撒下","撒母耳記下", chap_string, "1","korean")
							default:
								print_string = Bible_print_string("撒下","撒母耳記下", chap_string, sec_string,"korean")
						}
				}
			case "1 Kin","First Kings","王上","列王上","列王紀上","列王記上","1Ki","1ki","열왕기상","Третья книга Царств","1 Các Vua":
				bible_short_name = "王上"
				switch chap_string {
					case "":
						print_string = "列王紀上"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("王上","列王紀上", chap_string, "1","korean")
							default:
								print_string = Bible_print_string("王上","列王紀上", chap_string, sec_string,"korean")
						}
				}
			case "2 Kin","Second Kings","王下","列王下","列王記下","列王紀下","2Ki","2ki","열왕기하","2 Các Vua","Четвертая книга Царств":
				bible_short_name = "王下"
				switch chap_string {
					case "":
						print_string = "列王紀下"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("王下","列王紀下", chap_string, "1","korean")
							default:
								print_string = Bible_print_string("王下","列王紀下", chap_string, sec_string,"korean")
						}
				}
			case "1 Chr","First Chronicles","歷上","代上","歷代志上","歷代上","1Ch","1ch","Первая книга Паралипоменон","1 Sử-ký","歴上","歴代上","歴代志上","역대상":
				bible_short_name = "代上"
				switch chap_string {
					case "":
						print_string = "歷代志上"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("代上","歷代志上", chap_string, "1","korean")
							default:
								print_string = Bible_print_string("代上","歷代志上", chap_string, sec_string,"korean")
						}
				}
			case "2 Chr","Second Chronicles","代下","歷下","歷代下","歷代志下","2Ch","2ch","역대하","歴代志下","歴代下","歴下","2 Sử-ký","Вторая книга Паралипоменон":
				bible_short_name = "代下"
				switch chap_string {
					case "":
						print_string = "歷代志下"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("代下","歷代志下", chap_string, "1","korean")
							default:
								print_string = Bible_print_string("代下","歷代志下", chap_string, sec_string,"korean")
						}
				}
			case "Ezra","拉","以斯拉","以斯拉記","Ezr","ezr","Первая книга Ездры","Ê-xơ-ra","エズラ","エズラ記","에스라":
				bible_short_name = "拉"
				switch chap_string {
					case "":
						print_string = "以斯拉記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("拉","以斯拉記", chap_string, "1","korean")
							default:
								print_string = Bible_print_string("拉","以斯拉記", chap_string, sec_string,"korean")
						}
				}
			case "Neh","Nehemiah","尼","尼希米","尼希米記","Ne","ne","느헤미야","ネヘミヤ書","ネヘミヤ","Nê-hê-mi","Книга Неемии":
				bible_short_name = "尼"
				switch chap_string {
					case "":
						print_string = "尼希米記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("尼","尼希米記", chap_string, "1","korean")
							default:
								print_string = Bible_print_string("尼","尼希米記", chap_string, sec_string,"korean")
						}
				}
			case "Esth","Esther","斯","以斯帖","以斯帖記","Es","est","Есфирь","Ê-xơ-tê","エステル","エステル記","에스더":
				bible_short_name = "斯"
				switch chap_string {
					case "":
						print_string = "以斯帖記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("斯","以斯帖記", chap_string, "1","korean")
							default:
								print_string = Bible_print_string("斯","以斯帖記", chap_string, sec_string,"korean")
						}
				}
			case "Job","job","伯","約伯","約伯記","Книга Иова","Gióp","ヨブ","ヨブ記","욥기":
				bible_short_name = "伯"
				switch chap_string {
					case "":
						print_string = "約伯記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("伯","約伯記", chap_string, "1","korean")
							default:
								print_string = Bible_print_string("伯","約伯記", chap_string, sec_string,"korean")
						}
				}
			case "Ps","Psalms","詩","詩篇","ps","시편","Thi-thiên","Псалтирь":
				bible_short_name = "詩"
				switch chap_string {
					case "":
						print_string = "詩篇"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("詩","詩篇", chap_string, "1","korean")
							default:
								print_string = Bible_print_string("詩","詩篇", chap_string, sec_string,"korean")
						}
				}
			case "Prov","Proverbs","箴","箴言","Pr","pr","Притчи Соломона","Châm-ngôn","잠언":
				bible_short_name = "箴"
				switch chap_string {
					case "":
						print_string = "箴言"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("箴","箴言", chap_string, "1","korean")
							default:
								print_string = Bible_print_string("箴","箴言", chap_string, sec_string,"korean")
						}
				}
			case "Eccl","Ecclesiastes","傳","傳道","傳道書","Ec","ec","Книга Екклезиаста","или Проповедника","Truyền-đạo","伝道の書","伝道","伝","伝道書","전도서":
				bible_short_name = "傳"
				switch chap_string {
					case "":
						print_string = "傳道書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("傳","傳道書", chap_string, "1","korean")
							default:
								print_string = Bible_print_string("傳","傳道書", chap_string, sec_string,"korean")
						}
				}
			case "Song","Song of Solomon","歌","雅歌","So","so","sng","Sng","Песнь песней Соломона","Nhã-ca","아가":
				bible_short_name = "歌"
				switch chap_string {
					case "":
						print_string = "雅歌"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("歌","雅歌", chap_string, "1","korean")
							default:
								print_string = Bible_print_string("歌","雅歌", chap_string, sec_string,"korean")
						}
				}
			case "Is","Isaiah","賽","以賽","以賽亞","以賽亞書","Isa","isa","Книга пророка Исаии","Ê-sai","イザヤ書","イザヤ","이사야":
				bible_short_name = "賽"
				switch chap_string {
					case "":
						print_string = "以賽亞書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("賽","以賽亞書", chap_string, "1","korean")
							default:
								print_string = Bible_print_string("賽","以賽亞書", chap_string, sec_string,"korean")
						}
				}
			case "Jer","Jeremiah","耶","耶利米","耶利米書","jer","예레미야","エレミヤ","エレミヤ書","Giê-rê-mi","Книга пророка Иеремии":
				bible_short_name = "耶"
				switch chap_string {
					case "":
						print_string = "耶利米書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("耶","耶利米書", chap_string, "1","korean")
							default:
								print_string = Bible_print_string("耶","耶利米書", chap_string, sec_string,"korean")
						}
				}
			case "Lam","Lamentations","哀","哀歌","耶利米哀歌","La","lam","예레미야애가","Ca-thương","Плач Иеремии":
				bible_short_name = "哀"
				switch chap_string {
					case "":
						print_string = "耶利米哀歌"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("哀","耶利米哀歌", chap_string, "1","korean")
							default:
								print_string = Bible_print_string("哀","耶利米哀歌", chap_string, sec_string,"korean")
						}
				}
			case "Ezek","Ezekiel","結","以西結","以西結書","Eze","eze","에스겔","エゼキエル書","エゼキエル","Ê-xê-chi-ên","Книга пророка Иезекииля":
				bible_short_name = "結"
				switch chap_string {
					case "":
						print_string = "以西結書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("結","以西結書", chap_string, "1","korean")
							default:
								print_string = Bible_print_string("結","以西結書", chap_string, sec_string,"korean")
						}
				}
			case "Dan","Daniel","但","但以理","但以理書","Da","da","Книга пророка Даниила","Đa-ni-ên","ダニエル書","ダニエル","다니엘":
				bible_short_name = "但"
				switch chap_string {
					case "":
						print_string = "但以理書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("但","但以理書", chap_string, "1","korean")
							default:
								print_string = Bible_print_string("但","但以理書", chap_string, sec_string,"korean")
						}
				}
			case "Hos","Hosea","何","何西","何西阿","何西阿書","Ho","ho","Книга пророка Осии","Ô-sê","ホセア書","ホセア","호세아":
				bible_short_name = "何"
				switch chap_string {
					case "":
						print_string = "何西阿書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("何","何西阿書", chap_string, "1","korean")
							default:
								print_string = Bible_print_string("何","何西阿書", chap_string, sec_string,"korean")
						}
				}
			case "Joel","珥","約珥","約珥書","Joe","joe","Книга пророка Иоиля","Giô-ên","ヨエル書","ヨエル","요엘":
				bible_short_name = "珥"
				switch chap_string {
					case "":
						print_string = "約珥書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("珥","約珥書", chap_string, "1","korean")
							default:
								print_string = Bible_print_string("珥","約珥書", chap_string, sec_string,"korean")
						}
				}
			case "Amos","摩","阿摩司書","Am","am","Книга пророка Амоса","A-mốt","アモス書","アモス","아모스":
				bible_short_name = "摩"
				switch chap_string {
					case "":
						print_string = "阿摩司書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("摩","阿摩司書", chap_string, "1","korean")
							default:
								print_string = Bible_print_string("摩","阿摩司書", chap_string, sec_string,"korean")
						}
				}
			case "Obad","Obadiah","俄","俄巴底亞","俄巴底亞書","Ob","ob","오바댜","オバデヤ書","オバデヤ","Áp-đia","Книга пророка Авдия":
				bible_short_name = "俄"
				switch chap_string {
					case "":
						print_string = "俄巴底亞書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("俄","俄巴底亞書", chap_string, "1","korean")
							default:
								print_string = Bible_print_string("俄","俄巴底亞書", chap_string, sec_string,"korean")
						}
				}
			case "Jon","Jonah","拿","約拿","約拿書","jon","요나","ヨナ書","ヨナ","Giô-na","Книга пророка Ионы":
				bible_short_name = "拿"
				switch chap_string {
					case "":
						print_string = "約拿書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("拿","約拿書", chap_string, "1","korean")
							default:
								print_string = Bible_print_string("拿","約拿書", chap_string, sec_string,"korean")
						}
				}
			case "Micah","彌","彌迦","彌迦書","Mic","mic","Книга пророка Михея","Mi-chê","ミカ書","ミカ","미가":
				bible_short_name = "彌"
				switch chap_string {
					case "":
						print_string = "彌迦書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("彌","彌迦書", chap_string, "1","korean")
							default:
								print_string = Bible_print_string("彌","彌迦書", chap_string, sec_string,"korean")
						}
				}
			case "Nah","Nahum","鴻","那鴻","那鴻書","Na","na","Книга пророка Наума","Na-hum","ナホム書","ナホム","나훔":
				bible_short_name = "鴻"
				switch chap_string {
					case "":
						print_string = "那鴻書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("鴻","那鴻書", chap_string, "1","korean")
							default:
								print_string = Bible_print_string("鴻","那鴻書", chap_string, sec_string,"korean")
						}
				}
			case "Habakkuk","哈","哈巴","哈巴谷","哈巴谷書","Hab","hab","Книга пророка Аввакума","Ha-ba-cúc","ハバクク書","ハバクク","ハバ","クク","ハバ書","하박국":
				bible_short_name = "哈"
				switch chap_string {
					case "":
						print_string = "哈巴谷書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("哈","哈巴谷書", chap_string, "1","korean")
							default:
								print_string = Bible_print_string("哈","哈巴谷書", chap_string, sec_string,"korean")
						}
				}
			case "Zeph","Zephaniah","番","西番雅","西番雅書","Zep","zep","스바냐","ゼパニヤ書","ゼパニヤ","Sô-phô-ni","Книга пророка Софонии":
				bible_short_name = "番"
				switch chap_string {
					case "":
						print_string = "西番雅書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("番","西番雅書", chap_string, "1","korean")
							default:
								print_string = Bible_print_string("番","西番雅書", chap_string, sec_string,"korean")
						}
				}
			case "Haggai","該","哈該","哈該書","Hag","hag","학개","ハガイ書","ハガイ","A-ghê","Книга пророка Аггея":
				bible_short_name = "該"
				switch chap_string {
					case "":
						print_string = "哈該書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("該","哈該書", chap_string, "1","korean")
							default:
								print_string = Bible_print_string("該","哈該書", chap_string, sec_string,"korean")
						}
				}
			case "Zech","Zechariah","亞","撒迦利亞","撒迦利亞書","Zec","zec","Книга пророка Захарии","Xa-cha-ri","스가랴","ゼカリヤ書","ゼカリヤ":
				bible_short_name = "亞"
				switch chap_string {
					case "":
						print_string = "撒迦利亞書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("亞","撒迦利亞書", chap_string, "1","korean")
							default:
								print_string = Bible_print_string("亞","撒迦利亞書", chap_string, sec_string,"korean")
						}
				}
			case "Malachi","瑪","瑪拉","瑪拉基","瑪拉基書","Mal","mal","말라기","マラキ書","マラキ","Ma-la-chi","Книга пророка Малахии":
				bible_short_name = "瑪"
				switch chap_string {
					case "":
						print_string = "瑪拉基書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("瑪","瑪拉基書", chap_string, "1","korean")
							default:
								print_string = Bible_print_string("瑪","瑪拉基書", chap_string, sec_string,"korean")
						}
				}
			case "Matt","Matthew","太","馬太","馬太福音","Mt","mt","마태복음","マタイによる福音書","マタイ","マタイによる","Ma-thi-ơ","От Матфея святое благовествование":
				bible_short_name = "太"
				switch chap_string {
					case "":
						print_string = "馬太福音"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("太","馬太福音", chap_string, "1","korean")
							default:
								print_string = Bible_print_string("太","馬太福音", chap_string, sec_string,"korean")
						}
				}
			case "Mark","可","馬可","馬可福音","Mr","mr","マルコによる福音書","マルコ","マルコによる","마가복음","Mác","От Марка святое благовествование":
				bible_short_name = "可"
				switch chap_string {
					case "":
						print_string = "馬可福音"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("可","馬可福音", chap_string, "1","korean")
							default:
								print_string = Bible_print_string("可","馬可福音", chap_string, sec_string,"korean")
						}
				}
			case "Luke","路","路加","路加福音","Lu","lu","От Луки святое благовествование","Lu-ca","ルカによる福音書","ルカ","ルカによる","누가복음":
				bible_short_name = "路"
				switch chap_string {
					case "":
						print_string = "路加福音"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("路","路加福音", chap_string, "1","korean")
							default:
								print_string = Bible_print_string("路","路加福音", chap_string, sec_string,"korean")
						}
				}
			case "John","約","約翰","約翰福音","Joh","joh","От Иоанна святое благовествование","Giăng","ヨハネによる福音書","ヨハネ","ヨハネによる","요한복음":
				bible_short_name = "約"
				switch chap_string {
					case "":
						print_string = "約翰福音"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("約","約翰福音", chap_string, "1","korean")
							default:
								print_string = Bible_print_string("約","約翰福音", chap_string, sec_string,"korean")
						}
				}
			case "Acts","徒","使徒","使徒行傳","Ac","ac","Деяния святых апостолов","Công-vụ Các Sứ-đồ","使徒行伝","사도행전":
				bible_short_name = "徒"
				switch chap_string {
					case "":
						print_string = "使徒行傳"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("徒","使徒行傳", chap_string, "1","korean")
							default:
								print_string = Bible_print_string("徒","使徒行傳", chap_string, sec_string,"korean")
						}
				}
			case "Rom","Romans","羅","羅馬","羅馬書","Ro","ro","Послание к Римлянам","Rô-ma","ローマ","ローマ人への手紙","로마서":
				bible_short_name = "羅"
				switch chap_string {
					case "":
						print_string = "羅馬書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("羅","羅馬書", chap_string, "1","korean")
							default:
								print_string = Bible_print_string("羅","羅馬書", chap_string, sec_string,"korean")
						}
				}
			case "1 Cor","First Corinthians","林前","哥林多前","哥林多前書","1Co","1co","Первое послание к Коринфянам","1 Cô-rinh-tô","コリント人への第一の手紙","コリント一","コリント人への第一","고린도전서":
				bible_short_name = "林前"
				switch chap_string {
					case "":
						print_string = "哥林多前書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("林前","哥林多前書", chap_string, "1","korean")
							default:
								print_string = Bible_print_string("林前","哥林多前書", chap_string, sec_string,"korean")
						}
				}
			case "2 Cor","Second Corinthians","林後","哥林多後","哥林多後書","2Co","2co","Второе послание к Коринфянам","2 Cô-rinh-tô","コリント人への第二の手紙","コリント二","コリント人への第二の","고린도후서":
				bible_short_name = "林後"
				switch chap_string {
					case "":
						print_string = "哥林多後書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("林後","哥林多後書", chap_string, "1","korean")
							default:
								print_string = Bible_print_string("林後","哥林多後書", chap_string, sec_string,"korean")
						}
				}
			case "Gal","Galatians","加","加拉太","加拉太書","Ga","ga","Послание к Галатам","Ga-la-ti","ガラテヤ","ガラテヤ人への手紙","갈라디아서":
				bible_short_name = "加"
				switch chap_string {
					case "":
						print_string = "加拉太書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("加","加拉太書", chap_string, "1","korean")
							default:
								print_string = Bible_print_string("加","加拉太書", chap_string, sec_string,"korean")
						}
				}
			case "Ephesians","弗","以弗所","以弗所書","Eph","eph","Послание к Ефесянам","Ê-phê-sô","エペソ人への手紙","エペソ","エペソ人","エペソ人の手紙","에베소서":
				bible_short_name = "弗"
				switch chap_string {
					case "":
						print_string = "以弗所書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("弗","以弗所書", chap_string, "1","korean")
							default:
								print_string = Bible_print_string("弗","以弗所書", chap_string, sec_string,"korean")
						}
				}
			case "Phil","Philippians","腓","腓立","腓立比","腓立比書","Php","php","빌립보서","ピリピ","ピリピ人.ピリピ人への手紙","Послание к Филиппийцам","Phi-líp":
				bible_short_name = "腓"
				switch chap_string {
					case "":
						print_string = "腓立比書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("腓","腓立比書", chap_string, "1","korean")
							default:
								print_string = Bible_print_string("腓","腓立比書", chap_string, sec_string,"korean")
						}
				}
			case "Col","col","Colossians","西","歌羅西","歌羅","歌羅西書","Послание к Колоссянам","Cô-lô-se","コロサイ人への手紙","コロサイ","コロ","골로새서":
				bible_short_name = "西"
				switch chap_string {
					case "":
						print_string = "歌羅西書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("西","歌羅西書", chap_string, "1","korean")
							default:
								print_string = Bible_print_string("西","歌羅西書", chap_string, sec_string,"korean")
						}
				}
			case "1 Thess","First Thessalonians","帖前","帖撒羅尼迦前","帖撒羅尼迦前書","1Th","1th","데살로니가전서","テサロニケ人への第一の手紙","テサ一","テサロニケ一","1 Tê-sa-lô-ni-ca","Первое послание к Фессалоникийцам (Солунянам)":
				bible_short_name = "帖前"
				switch chap_string {
					case "":
						print_string = "帖撒羅尼迦前書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("帖前","帖撒羅尼迦前書", chap_string, "1","korean")
							default:
								print_string = Bible_print_string("帖前","帖撒羅尼迦前書", chap_string, sec_string,"korean")
						}
				}
			case "2 Thess","Second Thessalonians","帖後","帖撒羅尼迦後","帖撒羅尼迦後書","2Th","2th","데살로니가후서","テサロニケ人への第二の手紙","テサ二","テサロニケ二","2 Tê-sa-lô-ni-ca","Второе послание к Фессалоникийцам (Солунянам)":
				bible_short_name = "帖後"
				switch chap_string {
					case "":
						print_string = "帖撒羅尼迦後書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("帖後","帖撒羅尼迦後書", chap_string, "1","korean")
							default:
								print_string = Bible_print_string("帖後","帖撒羅尼迦後書", chap_string, sec_string,"korean")
						}
				}
			case "1 Tim","First Timothy","提前","提摩太前","提摩太前書","1Ti","1ti","Первое послание к Тимофею","1 Ti-mô-thê","テモテヘの第一の手紙","テモテ一","디모데전서":
				bible_short_name = "提前"
				switch chap_string {
					case "":
						print_string = "提摩太前書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("提前","提摩太前書", chap_string, "1","korean")
							default:
								print_string = Bible_print_string("提前","提摩太前書", chap_string, sec_string,"korean")
						}
				}
			case "2 Tim","Second Timothy","提後","提摩太後","提摩太後書","2Ti","2ti","Второе послание к Тимофею","2 Ti-mô-thê","テモテヘの第二の手紙","テモテ二","디모데후서":
				bible_short_name = "提後"
				switch chap_string {
					case "":
						print_string = "提摩太後書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("提後","提摩太後書", chap_string, "1","korean")
							default:
								print_string = Bible_print_string("提後","提摩太後書", chap_string, sec_string,"korean")
						}
				}
			case "Titus","多","提多","提多書","Tit","tit","Послание к Титу","Tít","テトスヘの手紙","テトス","디도서":
				bible_short_name = "多"
				switch chap_string {
					case "":
						print_string = "提多書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("多","提多書", chap_string, "1","korean")
							default:
								print_string = Bible_print_string("多","提多書", chap_string, sec_string,"korean")
						}
				}
			case "Philem","Philemon","門","腓利","腓利門","腓利門書","Phm","phm","Послание к Филимону","Phi-lê-môn","ピレモンヘの手紙","ピレモン","빌레몬서":
				bible_short_name = "門"
				switch chap_string {
					case "":
						print_string = "腓利門書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("門","腓利門書", chap_string, "1","korean")
							default:
								print_string = Bible_print_string("門","腓利門書", chap_string, sec_string,"korean")
						}
				}
			case "Heb","Hebrews","來","希伯來","希伯來書","heb","Послание к Евреям","Hê-bơ-rơ","ヘブル人への手紙","ヘブル","히브리서":
				bible_short_name = "來"
				switch chap_string {
					case "":
						print_string = "希伯來書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("來","希伯來書", chap_string, "1","korean")
							default:
								print_string = Bible_print_string("來","希伯來書", chap_string, sec_string,"korean")
						}
				}
			case "James","雅","雅各","雅各書","Jas","jas","Послание Иакова","Gia-cơ","ヤコブの手紙","ヤコブ","야고보서":
				bible_short_name = "雅"
				switch chap_string {
					case "":
						print_string = "雅各書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("雅","雅各書", chap_string, "1","korean")
							default:
								print_string = Bible_print_string("雅","雅各書", chap_string, sec_string,"korean")
						}
				}
			case "1 Pet","First Peter","彼前","彼得前","彼得前書","1Pe","1pe","Первое послание Петра","1 Phi-e-rơ","ペテロの第一の手紙","ペテロ一","베드로전서":
				bible_short_name = "彼前"
				switch chap_string {
					case "":
						print_string = "彼得前書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("彼前","彼得前書", chap_string, "1","korean")
							default:
								print_string = Bible_print_string("彼前","彼得前書", chap_string, sec_string,"korean")
						}
				}
			case "2 Pet","Second Peter","彼後","彼得後","彼得後書","2Pe","2pe","Второе послание Петра","2 Phi-e-rơ","ペテロの第二の手紙","ペテロ","베드로후서":
				bible_short_name = "彼後"
				switch chap_string {
					case "":
						print_string = "彼得後書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("彼後","彼得後書", chap_string, "1","korean")
							default:
								print_string = Bible_print_string("彼後","彼得後書", chap_string, sec_string,"korean")
						}
				}
			case "1 John","First John","約一","約翰一書","約翰1","約翰1書","1Jo","1jo","Первое послание Иоанна","1 Giăng","ヨハネの第一の手紙","ヨハネ一","요한일서":
				bible_short_name = "約一"
				switch chap_string {
					case "":
						print_string = "約翰一書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("約一","約翰一書", chap_string, "1","korean")
							default:
								print_string = Bible_print_string("約一","約翰一書", chap_string, sec_string,"korean")
						}
				}
			case "2 John","second John","約二","約翰二書","約翰2","約翰2書","2Jo","Второе послание Иоанна","2 Giăng","ヨハネの第二の手紙","ヨハネ二","요한2서":
				bible_short_name = "約二"
				switch chap_string {
					case "":
						print_string = "約翰二書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("約二","約翰二書", chap_string, "1","korean")
							default:
								print_string = Bible_print_string("約二","約翰二書", chap_string, sec_string,"korean")
						}
				}
			case "3 John","Third John","約三","約翰三書","約翰3","約翰3書","3Jo","3jo","Третье послание Иоанна","3 Giăng","ヨハネの第三の手紙","ヨハネ三","요한3서":
				bible_short_name = "約三"
				switch chap_string {
					case "":
						print_string = "約翰三書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("約三","約翰三書", chap_string, "1","korean")
							default:
								print_string = Bible_print_string("約三","約翰三書", chap_string, sec_string,"korean")
						}
				}
			case "Jude","猶","猶大","猶大書","jude","Послание Иуды","Giu-đe","ユダの手紙","ユダ","유다서":
				bible_short_name = "猶"
				switch chap_string {
					case "":
						print_string = "猶大書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("猶","猶大書", chap_string, "1","korean")
							default:
								print_string = Bible_print_string("猶","猶大書", chap_string, sec_string,"korean")
						}
				}
			case "Rev","Revelation","啟","啟示","啟示錄","Re","re","ｒｅ","Ｒｅ","rev","Откровение ап. Иоанна Богослова (Апокалипсис)","Khải-huyền","ヨハネの黙示録","黙示録","요한계시록":
				bible_short_name = "啟"
				switch chap_string {
					case "":
						print_string = "啟示錄"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("啟","啟示錄", chap_string, "1","korean")
							default:
								print_string = Bible_print_string("啟","啟示錄", chap_string, sec_string,"korean")
						}
				}
			case "韓文聖經","KR bible","Kr Bible","Kr bible":
				print_string = "聖經"
			default:
				//print_string = "你是要找 " +  reg.ReplaceAllString(text, "$3") + " 對嗎？\n對不起，我還沒學呢...\n"
				print_string = "聖經"
		}
	case "日文聖經","聖書","日本語聖書","JP bible","JP Bible","Jp bible":
		log.Print(reg.ReplaceAllString(text, "$3"))
		switch reg.ReplaceAllString(text, "$3") {
			case "Gen","Genesis","創","創世","創世紀","創世記","Ge","ge","gen","창세기","Sáng-thế Ký","Бытие":
				bible_short_name = "創"
				switch chap_string {
					case "":
						print_string = "創世記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("創","創世記", chap_string, "1","jp")
							default:
								print_string = Bible_print_string("創","創世記", chap_string, sec_string,"jp")
						}
				}
			case "ex","Ex","Exodus","埃及","出","出埃及","出埃及記","출애굽기","エジプト","出エジプト","出エジプト記","Xuất Ê-díp-tô Ký","Исход":
				bible_short_name = "出"
				switch chap_string {
					case "":
						print_string = "出埃及記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("出","出埃及記", chap_string, "1","jp")
							default:
								print_string = Bible_print_string("出","出埃及記", chap_string, sec_string,"jp")
						}
				}
			case "Lev","Leviticus","利","利未","利未記","Le","le","Левит","Lê-vi Ký","レビ記","レビ","레위기":
				bible_short_name = "利"
				switch chap_string {
					case "":
						print_string = "利未記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("利","利未記", chap_string, "1","jp")
							default:
								print_string = Bible_print_string("利","利未記", chap_string, sec_string,"jp")
						}
				}
			case "Num","Numbers","民","民數","民數記","Nu","nu","민수기","民数","民数記","Dân-số Ký","Числа":
				bible_short_name = "民"
				switch chap_string {
					case "":
						print_string = "民數記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("民","民數記", chap_string, "1","jp")
							default:
								print_string = Bible_print_string("民","民數記", chap_string, sec_string,"jp")
						}
				}
			case "Deut","Deuteronomy","申","申命","申命記","De","de","신명기","Phục-truyền Luật-lệ Ký","Второзаконие":
				bible_short_name = "申"
				switch chap_string {
					case "":
						print_string = "申命記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("申","申命記", chap_string, "1","jp")
							default:
								print_string = Bible_print_string("申","申命記", chap_string, sec_string,"jp")
						}
				}
			case "Josh","Joshua","約書亞","約書亞記","Jos","jos","여호수아","ヨシュア記","ヨシュア","Giô-suê","Книга Иисуса Навина":
				bible_short_name = "書"
				switch chap_string {
					case "":
						print_string = "約書亞記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("書","約書亞記", chap_string, "1","jp")
							default:
								print_string = Bible_print_string("書","約書亞記", chap_string, sec_string,"jp")
						}
				}
			case "Judg","Judges","士","士師","士師記","Jud","jud","jdg","Jdg","Книга Судей израилевых","Các Quan Xét","사사기":
				bible_short_name = "士"
				switch chap_string {
					case "":
						print_string = "士師記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("士","士師記", chap_string, "1","jp")
							default:
								print_string = Bible_print_string("士","士師記", chap_string, sec_string,"jp")
						}
				}
			case "Ruth","路得","路得記","Ru","ru","Rut","rut","룻기","ルツ","ルツ記","Ru-tơ","Книга Руфи":
				bible_short_name = "得"
				switch chap_string {
					case "":
						print_string = "路得記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("得","路得記", chap_string, "1","jp")
							default:
								print_string = Bible_print_string("得","路得記", chap_string, sec_string,"jp")
						}
				}
			case "1 Sam","First Samuel","撒上","撒母耳記上","1Sa","1sa","サムエル記上","サムエル上","サム上","사무엘상","1 Sa-mu-ên","Первая книга Царств":
				bible_short_name = "撒上"
				switch chap_string {
					case "":
						print_string = "撒母耳記上"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("撒上","撒母耳記上", chap_string, "1","jp")
							default:
								print_string = Bible_print_string("撒上","撒母耳記上", chap_string, sec_string,"jp")
						}
				}
			case "2 Sam","Second Samuel","撒下","撒母耳記下","2Sa","2sa","사무엘하","サムエル記下","サムエル下","サム下","2 Sa-mu-ên","Вторая книга Царств":
				bible_short_name = "撒下"
				switch chap_string {
					case "":
						print_string = "撒母耳記下"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("撒下","撒母耳記下", chap_string, "1","jp")
							default:
								print_string = Bible_print_string("撒下","撒母耳記下", chap_string, sec_string,"jp")
						}
				}
			case "1 Kin","First Kings","王上","列王上","列王紀上","列王記上","1Ki","1ki","열왕기상","Третья книга Царств","1 Các Vua":
				bible_short_name = "王上"
				switch chap_string {
					case "":
						print_string = "列王紀上"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("王上","列王紀上", chap_string, "1","jp")
							default:
								print_string = Bible_print_string("王上","列王紀上", chap_string, sec_string,"jp")
						}
				}
			case "2 Kin","Second Kings","王下","列王下","列王記下","列王紀下","2Ki","2ki","열왕기하","2 Các Vua","Четвертая книга Царств":
				bible_short_name = "王下"
				switch chap_string {
					case "":
						print_string = "列王紀下"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("王下","列王紀下", chap_string, "1","jp")
							default:
								print_string = Bible_print_string("王下","列王紀下", chap_string, sec_string,"jp")
						}
				}
			case "1 Chr","First Chronicles","歷上","代上","歷代志上","歷代上","1Ch","1ch","Первая книга Паралипоменон","1 Sử-ký","歴上","歴代上","歴代志上","역대상":
				bible_short_name = "代上"
				switch chap_string {
					case "":
						print_string = "歷代志上"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("代上","歷代志上", chap_string, "1","jp")
							default:
								print_string = Bible_print_string("代上","歷代志上", chap_string, sec_string,"jp")
						}
				}
			case "2 Chr","Second Chronicles","代下","歷下","歷代下","歷代志下","2Ch","2ch","역대하","歴代志下","歴代下","歴下","2 Sử-ký","Вторая книга Паралипоменон":
				bible_short_name = "代下"
				switch chap_string {
					case "":
						print_string = "歷代志下"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("代下","歷代志下", chap_string, "1","jp")
							default:
								print_string = Bible_print_string("代下","歷代志下", chap_string, sec_string,"jp")
						}
				}
			case "Ezra","拉","以斯拉","以斯拉記","Ezr","ezr","Первая книга Ездры","Ê-xơ-ra","エズラ","エズラ記","에스라":
				bible_short_name = "拉"
				switch chap_string {
					case "":
						print_string = "以斯拉記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("拉","以斯拉記", chap_string, "1","jp")
							default:
								print_string = Bible_print_string("拉","以斯拉記", chap_string, sec_string,"jp")
						}
				}
			case "Neh","Nehemiah","尼","尼希米","尼希米記","Ne","ne","느헤미야","ネヘミヤ書","ネヘミヤ","Nê-hê-mi","Книга Неемии":
				bible_short_name = "尼"
				switch chap_string {
					case "":
						print_string = "尼希米記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("尼","尼希米記", chap_string, "1","jp")
							default:
								print_string = Bible_print_string("尼","尼希米記", chap_string, sec_string,"jp")
						}
				}
			case "Esth","Esther","斯","以斯帖","以斯帖記","Es","est","Есфирь","Ê-xơ-tê","エステル","エステル記","에스더":
				bible_short_name = "斯"
				switch chap_string {
					case "":
						print_string = "以斯帖記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("斯","以斯帖記", chap_string, "1","jp")
							default:
								print_string = Bible_print_string("斯","以斯帖記", chap_string, sec_string,"jp")
						}
				}
			case "Job","job","伯","約伯","約伯記","Книга Иова","Gióp","ヨブ","ヨブ記","욥기":
				bible_short_name = "伯"
				switch chap_string {
					case "":
						print_string = "約伯記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("伯","約伯記", chap_string, "1","jp")
							default:
								print_string = Bible_print_string("伯","約伯記", chap_string, sec_string,"jp")
						}
				}
			case "Ps","Psalms","詩","詩篇","ps","시편","Thi-thiên","Псалтирь":
				bible_short_name = "詩"
				switch chap_string {
					case "":
						print_string = "詩篇"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("詩","詩篇", chap_string, "1","jp")
							default:
								print_string = Bible_print_string("詩","詩篇", chap_string, sec_string,"jp")
						}
				}
			case "Prov","Proverbs","箴","箴言","Pr","pr","Притчи Соломона","Châm-ngôn","잠언":
				bible_short_name = "箴"
				switch chap_string {
					case "":
						print_string = "箴言"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("箴","箴言", chap_string, "1","jp")
							default:
								print_string = Bible_print_string("箴","箴言", chap_string, sec_string,"jp")
						}
				}
			case "Eccl","Ecclesiastes","傳","傳道","傳道書","Ec","ec","Книга Екклезиаста","или Проповедника","Truyền-đạo","伝道の書","伝道","伝","伝道書","전도서":
				bible_short_name = "傳"
				switch chap_string {
					case "":
						print_string = "傳道書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("傳","傳道書", chap_string, "1","jp")
							default:
								print_string = Bible_print_string("傳","傳道書", chap_string, sec_string,"jp")
						}
				}
			case "Song","Song of Solomon","歌","雅歌","So","so","sng","Sng","Песнь песней Соломона","Nhã-ca","아가":
				bible_short_name = "歌"
				switch chap_string {
					case "":
						print_string = "雅歌"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("歌","雅歌", chap_string, "1","jp")
							default:
								print_string = Bible_print_string("歌","雅歌", chap_string, sec_string,"jp")
						}
				}
			case "Is","Isaiah","賽","以賽","以賽亞","以賽亞書","Isa","isa","Книга пророка Исаии","Ê-sai","イザヤ書","イザヤ","이사야":
				bible_short_name = "賽"
				switch chap_string {
					case "":
						print_string = "以賽亞書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("賽","以賽亞書", chap_string, "1","jp")
							default:
								print_string = Bible_print_string("賽","以賽亞書", chap_string, sec_string,"jp")
						}
				}
			case "Jer","Jeremiah","耶","耶利米","耶利米書","jer","예레미야","エレミヤ","エレミヤ書","Giê-rê-mi","Книга пророка Иеремии":
				bible_short_name = "耶"
				switch chap_string {
					case "":
						print_string = "耶利米書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("耶","耶利米書", chap_string, "1","jp")
							default:
								print_string = Bible_print_string("耶","耶利米書", chap_string, sec_string,"jp")
						}
				}
			case "Lam","Lamentations","哀","哀歌","耶利米哀歌","La","lam","예레미야애가","Ca-thương","Плач Иеремии":
				bible_short_name = "哀"
				switch chap_string {
					case "":
						print_string = "耶利米哀歌"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("哀","耶利米哀歌", chap_string, "1","jp")
							default:
								print_string = Bible_print_string("哀","耶利米哀歌", chap_string, sec_string,"jp")
						}
				}
			case "Ezek","Ezekiel","結","以西結","以西結書","Eze","eze","에스겔","エゼキエル書","エゼキエル","Ê-xê-chi-ên","Книга пророка Иезекииля":
				bible_short_name = "結"
				switch chap_string {
					case "":
						print_string = "以西結書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("結","以西結書", chap_string, "1","jp")
							default:
								print_string = Bible_print_string("結","以西結書", chap_string, sec_string,"jp")
						}
				}
			case "Dan","Daniel","但","但以理","但以理書","Da","da","Книга пророка Даниила","Đa-ni-ên","ダニエル書","ダニエル","다니엘":
				bible_short_name = "但"
				switch chap_string {
					case "":
						print_string = "但以理書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("但","但以理書", chap_string, "1","jp")
							default:
								print_string = Bible_print_string("但","但以理書", chap_string, sec_string,"jp")
						}
				}
			case "Hos","Hosea","何","何西","何西阿","何西阿書","Ho","ho","Книга пророка Осии","Ô-sê","ホセア書","ホセア","호세아":
				bible_short_name = "何"
				switch chap_string {
					case "":
						print_string = "何西阿書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("何","何西阿書", chap_string, "1","jp")
							default:
								print_string = Bible_print_string("何","何西阿書", chap_string, sec_string,"jp")
						}
				}
			case "Joel","珥","約珥","約珥書","Joe","joe","Книга пророка Иоиля","Giô-ên","ヨエル書","ヨエル","요엘":
				bible_short_name = "珥"
				switch chap_string {
					case "":
						print_string = "約珥書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("珥","約珥書", chap_string, "1","jp")
							default:
								print_string = Bible_print_string("珥","約珥書", chap_string, sec_string,"jp")
						}
				}
			case "Amos","摩","阿摩司書","Am","am","Книга пророка Амоса","A-mốt","アモス書","アモス","아모스":
				bible_short_name = "摩"
				switch chap_string {
					case "":
						print_string = "阿摩司書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("摩","阿摩司書", chap_string, "1","jp")
							default:
								print_string = Bible_print_string("摩","阿摩司書", chap_string, sec_string,"jp")
						}
				}
			case "Obad","Obadiah","俄","俄巴底亞","俄巴底亞書","Ob","ob","오바댜","オバデヤ書","オバデヤ","Áp-đia","Книга пророка Авдия":
				bible_short_name = "俄"
				switch chap_string {
					case "":
						print_string = "俄巴底亞書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("俄","俄巴底亞書", chap_string, "1","jp")
							default:
								print_string = Bible_print_string("俄","俄巴底亞書", chap_string, sec_string,"jp")
						}
				}
			case "Jon","Jonah","拿","約拿","約拿書","jon","요나","ヨナ書","ヨナ","Giô-na","Книга пророка Ионы":
				bible_short_name = "拿"
				switch chap_string {
					case "":
						print_string = "約拿書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("拿","約拿書", chap_string, "1","jp")
							default:
								print_string = Bible_print_string("拿","約拿書", chap_string, sec_string,"jp")
						}
				}
			case "Micah","彌","彌迦","彌迦書","Mic","mic","Книга пророка Михея","Mi-chê","ミカ書","ミカ","미가":
				bible_short_name = "彌"
				switch chap_string {
					case "":
						print_string = "彌迦書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("彌","彌迦書", chap_string, "1","jp")
							default:
								print_string = Bible_print_string("彌","彌迦書", chap_string, sec_string,"jp")
						}
				}
			case "Nah","Nahum","鴻","那鴻","那鴻書","Na","na","Книга пророка Наума","Na-hum","ナホム書","ナホム","나훔":
				bible_short_name = "鴻"
				switch chap_string {
					case "":
						print_string = "那鴻書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("鴻","那鴻書", chap_string, "1","jp")
							default:
								print_string = Bible_print_string("鴻","那鴻書", chap_string, sec_string,"jp")
						}
				}
			case "Habakkuk","哈","哈巴","哈巴谷","哈巴谷書","Hab","hab","Книга пророка Аввакума","Ha-ba-cúc","ハバクク書","ハバクク","ハバ","クク","ハバ書","하박국":
				bible_short_name = "哈"
				switch chap_string {
					case "":
						print_string = "哈巴谷書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("哈","哈巴谷書", chap_string, "1","jp")
							default:
								print_string = Bible_print_string("哈","哈巴谷書", chap_string, sec_string,"jp")
						}
				}
			case "Zeph","Zephaniah","番","西番雅","西番雅書","Zep","zep","스바냐","ゼパニヤ書","ゼパニヤ","Sô-phô-ni","Книга пророка Софонии":
				bible_short_name = "番"
				switch chap_string {
					case "":
						print_string = "西番雅書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("番","西番雅書", chap_string, "1","jp")
							default:
								print_string = Bible_print_string("番","西番雅書", chap_string, sec_string,"jp")
						}
				}
			case "Haggai","該","哈該","哈該書","Hag","hag","학개","ハガイ書","ハガイ","A-ghê","Книга пророка Аггея":
				bible_short_name = "該"
				switch chap_string {
					case "":
						print_string = "哈該書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("該","哈該書", chap_string, "1","jp")
							default:
								print_string = Bible_print_string("該","哈該書", chap_string, sec_string,"jp")
						}
				}
			case "Zech","Zechariah","亞","撒迦利亞","撒迦利亞書","Zec","zec","Книга пророка Захарии","Xa-cha-ri","스가랴","ゼカリヤ書","ゼカリヤ":
				bible_short_name = "亞"
				switch chap_string {
					case "":
						print_string = "撒迦利亞書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("亞","撒迦利亞書", chap_string, "1","jp")
							default:
								print_string = Bible_print_string("亞","撒迦利亞書", chap_string, sec_string,"jp")
						}
				}
			case "Malachi","瑪","瑪拉","瑪拉基","瑪拉基書","Mal","mal","말라기","マラキ書","マラキ","Ma-la-chi","Книга пророка Малахии":
				bible_short_name = "瑪"
				switch chap_string {
					case "":
						print_string = "瑪拉基書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("瑪","瑪拉基書", chap_string, "1","jp")
							default:
								print_string = Bible_print_string("瑪","瑪拉基書", chap_string, sec_string,"jp")
						}
				}
			case "Matt","Matthew","太","馬太","馬太福音","Mt","mt","마태복음","マタイによる福音書","マタイ","マタイによる","Ma-thi-ơ","От Матфея святое благовествование":
				bible_short_name = "太"
				switch chap_string {
					case "":
						print_string = "馬太福音"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("太","馬太福音", chap_string, "1","jp")
							default:
								print_string = Bible_print_string("太","馬太福音", chap_string, sec_string,"jp")
						}
				}
			case "Mark","可","馬可","馬可福音","Mr","mr","マルコによる福音書","マルコ","マルコによる","마가복음","Mác","От Марка святое благовествование":
				bible_short_name = "可"
				switch chap_string {
					case "":
						print_string = "馬可福音"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("可","馬可福音", chap_string, "1","jp")
							default:
								print_string = Bible_print_string("可","馬可福音", chap_string, sec_string,"jp")
						}
				}
			case "Luke","路","路加","路加福音","Lu","lu","От Луки святое благовествование","Lu-ca","ルカによる福音書","ルカ","ルカによる","누가복음":
				bible_short_name = "路"
				switch chap_string {
					case "":
						print_string = "路加福音"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("路","路加福音", chap_string, "1","jp")
							default:
								print_string = Bible_print_string("路","路加福音", chap_string, sec_string,"jp")
						}
				}
			case "John","約","約翰","約翰福音","Joh","joh","От Иоанна святое благовествование","Giăng","ヨハネによる福音書","ヨハネ","ヨハネによる","요한복음":
				bible_short_name = "約"
				switch chap_string {
					case "":
						print_string = "約翰福音"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("約","約翰福音", chap_string, "1","jp")
							default:
								print_string = Bible_print_string("約","約翰福音", chap_string, sec_string,"jp")
						}
				}
			case "Acts","徒","使徒","使徒行傳","Ac","ac","Деяния святых апостолов","Công-vụ Các Sứ-đồ","使徒行伝","사도행전":
				bible_short_name = "徒"
				switch chap_string {
					case "":
						print_string = "使徒行傳"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("徒","使徒行傳", chap_string, "1","jp")
							default:
								print_string = Bible_print_string("徒","使徒行傳", chap_string, sec_string,"jp")
						}
				}
			case "Rom","Romans","羅","羅馬","羅馬書","Ro","ro","Послание к Римлянам","Rô-ma","ローマ","ローマ人への手紙","로마서":
				bible_short_name = "羅"
				switch chap_string {
					case "":
						print_string = "羅馬書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("羅","羅馬書", chap_string, "1","jp")
							default:
								print_string = Bible_print_string("羅","羅馬書", chap_string, sec_string,"jp")
						}
				}
			case "1 Cor","First Corinthians","林前","哥林多前","哥林多前書","1Co","1co","Первое послание к Коринфянам","1 Cô-rinh-tô","コリント人への第一の手紙","コリント一","コリント人への第一","고린도전서":
				bible_short_name = "林前"
				switch chap_string {
					case "":
						print_string = "哥林多前書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("林前","哥林多前書", chap_string, "1","jp")
							default:
								print_string = Bible_print_string("林前","哥林多前書", chap_string, sec_string,"jp")
						}
				}
			case "2 Cor","Second Corinthians","林後","哥林多後","哥林多後書","2Co","2co","Второе послание к Коринфянам","2 Cô-rinh-tô","コリント人への第二の手紙","コリント二","コリント人への第二の","고린도후서":
				bible_short_name = "林後"
				switch chap_string {
					case "":
						print_string = "哥林多後書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("林後","哥林多後書", chap_string, "1","jp")
							default:
								print_string = Bible_print_string("林後","哥林多後書", chap_string, sec_string,"jp")
						}
				}
			case "Gal","Galatians","加","加拉太","加拉太書","Ga","ga","Послание к Галатам","Ga-la-ti","ガラテヤ","ガラテヤ人への手紙","갈라디아서":
				bible_short_name = "加"
				switch chap_string {
					case "":
						print_string = "加拉太書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("加","加拉太書", chap_string, "1","jp")
							default:
								print_string = Bible_print_string("加","加拉太書", chap_string, sec_string,"jp")
						}
				}
			case "Ephesians","弗","以弗所","以弗所書","Eph","eph","Послание к Ефесянам","Ê-phê-sô","エペソ人への手紙","エペソ","エペソ人","エペソ人の手紙","에베소서":
				bible_short_name = "弗"
				switch chap_string {
					case "":
						print_string = "以弗所書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("弗","以弗所書", chap_string, "1","jp")
							default:
								print_string = Bible_print_string("弗","以弗所書", chap_string, sec_string,"jp")
						}
				}
			case "Phil","Philippians","腓","腓立","腓立比","腓立比書","Php","php","빌립보서","ピリピ","ピリピ人.ピリピ人への手紙","Послание к Филиппийцам","Phi-líp":
				bible_short_name = "腓"
				switch chap_string {
					case "":
						print_string = "腓立比書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("腓","腓立比書", chap_string, "1","jp")
							default:
								print_string = Bible_print_string("腓","腓立比書", chap_string, sec_string,"jp")
						}
				}
			case "Col","col","Colossians","西","歌羅西","歌羅","歌羅西書","Послание к Колоссянам","Cô-lô-se","コロサイ人への手紙","コロサイ","コロ","골로새서":
				bible_short_name = "西"
				switch chap_string {
					case "":
						print_string = "歌羅西書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("西","歌羅西書", chap_string, "1","jp")
							default:
								print_string = Bible_print_string("西","歌羅西書", chap_string, sec_string,"jp")
						}
				}
			case "1 Thess","First Thessalonians","帖前","帖撒羅尼迦前","帖撒羅尼迦前書","1Th","1th","데살로니가전서","テサロニケ人への第一の手紙","テサ一","テサロニケ一","1 Tê-sa-lô-ni-ca","Первое послание к Фессалоникийцам (Солунянам)":
				bible_short_name = "帖前"
				switch chap_string {
					case "":
						print_string = "帖撒羅尼迦前書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("帖前","帖撒羅尼迦前書", chap_string, "1","jp")
							default:
								print_string = Bible_print_string("帖前","帖撒羅尼迦前書", chap_string, sec_string,"jp")
						}
				}
			case "2 Thess","Second Thessalonians","帖後","帖撒羅尼迦後","帖撒羅尼迦後書","2Th","2th","데살로니가후서","テサロニケ人への第二の手紙","テサ二","テサロニケ二","2 Tê-sa-lô-ni-ca","Второе послание к Фессалоникийцам (Солунянам)":
				bible_short_name = "帖後"
				switch chap_string {
					case "":
						print_string = "帖撒羅尼迦後書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("帖後","帖撒羅尼迦後書", chap_string, "1","jp")
							default:
								print_string = Bible_print_string("帖後","帖撒羅尼迦後書", chap_string, sec_string,"jp")
						}
				}
			case "1 Tim","First Timothy","提前","提摩太前","提摩太前書","1Ti","1ti","Первое послание к Тимофею","1 Ti-mô-thê","テモテヘの第一の手紙","テモテ一","디모데전서":
				bible_short_name = "提前"
				switch chap_string {
					case "":
						print_string = "提摩太前書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("提前","提摩太前書", chap_string, "1","jp")
							default:
								print_string = Bible_print_string("提前","提摩太前書", chap_string, sec_string,"jp")
						}
				}
			case "2 Tim","Second Timothy","提後","提摩太後","提摩太後書","2Ti","2ti","Второе послание к Тимофею","2 Ti-mô-thê","テモテヘの第二の手紙","テモテ二","디모데후서":
				bible_short_name = "提後"
				switch chap_string {
					case "":
						print_string = "提摩太後書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("提後","提摩太後書", chap_string, "1","jp")
							default:
								print_string = Bible_print_string("提後","提摩太後書", chap_string, sec_string,"jp")
						}
				}
			case "Titus","多","提多","提多書","Tit","tit","Послание к Титу","Tít","テトスヘの手紙","テトス","디도서":
				bible_short_name = "多"
				switch chap_string {
					case "":
						print_string = "提多書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("多","提多書", chap_string, "1","jp")
							default:
								print_string = Bible_print_string("多","提多書", chap_string, sec_string,"jp")
						}
				}
			case "Philem","Philemon","門","腓利","腓利門","腓利門書","Phm","phm","Послание к Филимону","Phi-lê-môn","ピレモンヘの手紙","ピレモン","빌레몬서":
				bible_short_name = "門"
				switch chap_string {
					case "":
						print_string = "腓利門書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("門","腓利門書", chap_string, "1","jp")
							default:
								print_string = Bible_print_string("門","腓利門書", chap_string, sec_string,"jp")
						}
				}
			case "Heb","Hebrews","來","希伯來","希伯來書","heb","Послание к Евреям","Hê-bơ-rơ","ヘブル人への手紙","ヘブル","히브리서":
				bible_short_name = "來"
				switch chap_string {
					case "":
						print_string = "希伯來書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("來","希伯來書", chap_string, "1","jp")
							default:
								print_string = Bible_print_string("來","希伯來書", chap_string, sec_string,"jp")
						}
				}
			case "James","雅","雅各","雅各書","Jas","jas","Послание Иакова","Gia-cơ","ヤコブの手紙","ヤコブ","야고보서":
				bible_short_name = "雅"
				switch chap_string {
					case "":
						print_string = "雅各書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("雅","雅各書", chap_string, "1","jp")
							default:
								print_string = Bible_print_string("雅","雅各書", chap_string, sec_string,"jp")
						}
				}
			case "1 Pet","First Peter","彼前","彼得前","彼得前書","1Pe","1pe","Первое послание Петра","1 Phi-e-rơ","ペテロの第一の手紙","ペテロ一","베드로전서":
				bible_short_name = "彼前"
				switch chap_string {
					case "":
						print_string = "彼得前書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("彼前","彼得前書", chap_string, "1","jp")
							default:
								print_string = Bible_print_string("彼前","彼得前書", chap_string, sec_string,"jp")
						}
				}
			case "2 Pet","Second Peter","彼後","彼得後","彼得後書","2Pe","2pe","Второе послание Петра","2 Phi-e-rơ","ペテロの第二の手紙","ペテロ","베드로후서":
				bible_short_name = "彼後"
				switch chap_string {
					case "":
						print_string = "彼得後書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("彼後","彼得後書", chap_string, "1","jp")
							default:
								print_string = Bible_print_string("彼後","彼得後書", chap_string, sec_string,"jp")
						}
				}
			case "1 John","First John","約一","約翰一書","約翰1","約翰1書","1Jo","1jo","Первое послание Иоанна","1 Giăng","ヨハネの第一の手紙","ヨハネ一","요한일서":
				bible_short_name = "約一"
				switch chap_string {
					case "":
						print_string = "約翰一書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("約一","約翰一書", chap_string, "1","jp")
							default:
								print_string = Bible_print_string("約一","約翰一書", chap_string, sec_string,"jp")
						}
				}
			case "2 John","second John","約二","約翰二書","約翰2","約翰2書","2Jo","Второе послание Иоанна","2 Giăng","ヨハネの第二の手紙","ヨハネ二","요한2서":
				bible_short_name = "約二"
				switch chap_string {
					case "":
						print_string = "約翰二書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("約二","約翰二書", chap_string, "1","jp")
							default:
								print_string = Bible_print_string("約二","約翰二書", chap_string, sec_string,"jp")
						}
				}
			case "3 John","Third John","約三","約翰三書","約翰3","約翰3書","3Jo","3jo","Третье послание Иоанна","3 Giăng","ヨハネの第三の手紙","ヨハネ三","요한3서":
				bible_short_name = "約三"
				switch chap_string {
					case "":
						print_string = "約翰三書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("約三","約翰三書", chap_string, "1","jp")
							default:
								print_string = Bible_print_string("約三","約翰三書", chap_string, sec_string,"jp")
						}
				}
			case "Jude","猶","猶大","猶大書","jude","Послание Иуды","Giu-đe","ユダの手紙","ユダ","유다서":
				bible_short_name = "猶"
				switch chap_string {
					case "":
						print_string = "猶大書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("猶","猶大書", chap_string, "1","jp")
							default:
								print_string = Bible_print_string("猶","猶大書", chap_string, sec_string,"jp")
						}
				}
			case "Rev","Revelation","啟","啟示","啟示錄","Re","re","ｒｅ","Ｒｅ","rev","Откровение ап. Иоанна Богослова (Апокалипсис)","Khải-huyền","ヨハネの黙示録","黙示録","요한계시록":
				bible_short_name = "啟"
				switch chap_string {
					case "":
						print_string = "啟示錄"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("啟","啟示錄", chap_string, "1","jp")
							default:
								print_string = Bible_print_string("啟","啟示錄", chap_string, sec_string,"jp")
						}
				}
			case "日文聖經","聖書","日本語聖書","JP bible","JP Bible","Jp bible":
				print_string = "聖經"
			default:
				//print_string = "你是要找 " +  reg.ReplaceAllString(text, "$3") + " 對嗎？\n對不起，我還沒學呢...\n"
				print_string = "聖經"
		}
	case "英文聖經","英語聖書","Kjv","kjv","Ｋｊｖ","ｋｊｖ","Eng bible","ENG Bible","English bible":
		log.Print(reg.ReplaceAllString(text, "$3"))
		switch reg.ReplaceAllString(text, "$3") {
			case "Gen","Genesis","創","創世","創世紀","創世記","Ge","ge","gen","창세기","Sáng-thế Ký","Бытие":
				bible_short_name = "創"
				switch chap_string {
					case "":
						print_string = "創世記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("創","創世記", chap_string, "1","kjv")
							default:
								print_string = Bible_print_string("創","創世記", chap_string, sec_string,"kjv")
						}
				}
			case "ex","Ex","Exodus","埃及","出","出埃及","出埃及記","출애굽기","エジプト","出エジプト","出エジプト記","Xuất Ê-díp-tô Ký","Исход":
				bible_short_name = "出"
				switch chap_string {
					case "":
						print_string = "出埃及記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("出","出埃及記", chap_string, "1","kjv")
							default:
								print_string = Bible_print_string("出","出埃及記", chap_string, sec_string,"kjv")
						}
				}
			case "Lev","Leviticus","利","利未","利未記","Le","le","Левит","Lê-vi Ký","レビ記","レビ","레위기":
				bible_short_name = "利"
				switch chap_string {
					case "":
						print_string = "利未記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("利","利未記", chap_string, "1","kjv")
							default:
								print_string = Bible_print_string("利","利未記", chap_string, sec_string,"kjv")
						}
				}
			case "Num","Numbers","民","民數","民數記","Nu","nu","민수기","民数","民数記","Dân-số Ký","Числа":
				bible_short_name = "民"
				switch chap_string {
					case "":
						print_string = "民數記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("民","民數記", chap_string, "1","kjv")
							default:
								print_string = Bible_print_string("民","民數記", chap_string, sec_string,"kjv")
						}
				}
			case "Deut","Deuteronomy","申","申命","申命記","De","de","신명기","Phục-truyền Luật-lệ Ký","Второзаконие":
				bible_short_name = "申"
				switch chap_string {
					case "":
						print_string = "申命記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("申","申命記", chap_string, "1","kjv")
							default:
								print_string = Bible_print_string("申","申命記", chap_string, sec_string,"kjv")
						}
				}
			case "Josh","Joshua","約書亞","約書亞記","Jos","jos","여호수아","ヨシュア記","ヨシュア","Giô-suê","Книга Иисуса Навина":
				bible_short_name = "書"
				switch chap_string {
					case "":
						print_string = "約書亞記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("書","約書亞記", chap_string, "1","kjv")
							default:
								print_string = Bible_print_string("書","約書亞記", chap_string, sec_string,"kjv")
						}
				}
			case "Judg","Judges","士","士師","士師記","Jud","jud","jdg","Jdg","Книга Судей израилевых","Các Quan Xét","사사기":
				bible_short_name = "士"
				switch chap_string {
					case "":
						print_string = "士師記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("士","士師記", chap_string, "1","kjv")
							default:
								print_string = Bible_print_string("士","士師記", chap_string, sec_string,"kjv")
						}
				}
			case "Ruth","路得","路得記","Ru","ru","Rut","rut","룻기","ルツ","ルツ記","Ru-tơ","Книга Руфи":
				bible_short_name = "得"
				switch chap_string {
					case "":
						print_string = "路得記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("得","路得記", chap_string, "1","kjv")
							default:
								print_string = Bible_print_string("得","路得記", chap_string, sec_string,"kjv")
						}
				}
			case "1 Sam","First Samuel","撒上","撒母耳記上","1Sa","1sa","サムエル記上","サムエル上","サム上","사무엘상","1 Sa-mu-ên","Первая книга Царств":
				bible_short_name = "撒上"
				switch chap_string {
					case "":
						print_string = "撒母耳記上"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("撒上","撒母耳記上", chap_string, "1","kjv")
							default:
								print_string = Bible_print_string("撒上","撒母耳記上", chap_string, sec_string,"kjv")
						}
				}
			case "2 Sam","Second Samuel","撒下","撒母耳記下","2Sa","2sa","사무엘하","サムエル記下","サムエル下","サム下","2 Sa-mu-ên","Вторая книга Царств":
				bible_short_name = "撒下"
				switch chap_string {
					case "":
						print_string = "撒母耳記下"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("撒下","撒母耳記下", chap_string, "1","kjv")
							default:
								print_string = Bible_print_string("撒下","撒母耳記下", chap_string, sec_string,"kjv")
						}
				}
			case "1 Kin","First Kings","王上","列王上","列王紀上","列王記上","1Ki","1ki","열왕기상","Третья книга Царств","1 Các Vua":
				bible_short_name = "王上"
				switch chap_string {
					case "":
						print_string = "列王紀上"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("王上","列王紀上", chap_string, "1","kjv")
							default:
								print_string = Bible_print_string("王上","列王紀上", chap_string, sec_string,"kjv")
						}
				}
			case "2 Kin","Second Kings","王下","列王下","列王記下","列王紀下","2Ki","2ki","열왕기하","2 Các Vua","Четвертая книга Царств":
				bible_short_name = "王下"
				switch chap_string {
					case "":
						print_string = "列王紀下"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("王下","列王紀下", chap_string, "1","kjv")
							default:
								print_string = Bible_print_string("王下","列王紀下", chap_string, sec_string,"kjv")
						}
				}
			case "1 Chr","First Chronicles","歷上","代上","歷代志上","歷代上","1Ch","1ch","Первая книга Паралипоменон","1 Sử-ký","歴上","歴代上","歴代志上","역대상":
				bible_short_name = "代上"
				switch chap_string {
					case "":
						print_string = "歷代志上"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("代上","歷代志上", chap_string, "1","kjv")
							default:
								print_string = Bible_print_string("代上","歷代志上", chap_string, sec_string,"kjv")
						}
				}
			case "2 Chr","Second Chronicles","代下","歷下","歷代下","歷代志下","2Ch","2ch","역대하","歴代志下","歴代下","歴下","2 Sử-ký","Вторая книга Паралипоменон":
				bible_short_name = "代下"
				switch chap_string {
					case "":
						print_string = "歷代志下"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("代下","歷代志下", chap_string, "1","kjv")
							default:
								print_string = Bible_print_string("代下","歷代志下", chap_string, sec_string,"kjv")
						}
				}
			case "Ezra","拉","以斯拉","以斯拉記","Ezr","ezr","Первая книга Ездры","Ê-xơ-ra","エズラ","エズラ記","에스라":
				bible_short_name = "拉"
				switch chap_string {
					case "":
						print_string = "以斯拉記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("拉","以斯拉記", chap_string, "1","kjv")
							default:
								print_string = Bible_print_string("拉","以斯拉記", chap_string, sec_string,"kjv")
						}
				}
			case "Neh","Nehemiah","尼","尼希米","尼希米記","Ne","ne","느헤미야","ネヘミヤ書","ネヘミヤ","Nê-hê-mi","Книга Неемии":
				bible_short_name = "尼"
				switch chap_string {
					case "":
						print_string = "尼希米記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("尼","尼希米記", chap_string, "1","kjv")
							default:
								print_string = Bible_print_string("尼","尼希米記", chap_string, sec_string,"kjv")
						}
				}
			case "Esth","Esther","斯","以斯帖","以斯帖記","Es","est","Есфирь","Ê-xơ-tê","エステル","エステル記","에스더":
				bible_short_name = "斯"
				switch chap_string {
					case "":
						print_string = "以斯帖記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("斯","以斯帖記", chap_string, "1","kjv")
							default:
								print_string = Bible_print_string("斯","以斯帖記", chap_string, sec_string,"kjv")
						}
				}
			case "Job","job","伯","約伯","約伯記","Книга Иова","Gióp","ヨブ","ヨブ記","욥기":
				bible_short_name = "伯"
				switch chap_string {
					case "":
						print_string = "約伯記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("伯","約伯記", chap_string, "1","kjv")
							default:
								print_string = Bible_print_string("伯","約伯記", chap_string, sec_string,"kjv")
						}
				}
			case "Ps","Psalms","詩","詩篇","ps","시편","Thi-thiên","Псалтирь":
				bible_short_name = "詩"
				switch chap_string {
					case "":
						print_string = "詩篇"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("詩","詩篇", chap_string, "1","kjv")
							default:
								print_string = Bible_print_string("詩","詩篇", chap_string, sec_string,"kjv")
						}
				}
			case "Prov","Proverbs","箴","箴言","Pr","pr","Притчи Соломона","Châm-ngôn","잠언":
				bible_short_name = "箴"
				switch chap_string {
					case "":
						print_string = "箴言"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("箴","箴言", chap_string, "1","kjv")
							default:
								print_string = Bible_print_string("箴","箴言", chap_string, sec_string,"kjv")
						}
				}
			case "Eccl","Ecclesiastes","傳","傳道","傳道書","Ec","ec","Книга Екклезиаста","или Проповедника","Truyền-đạo","伝道の書","伝道","伝","伝道書","전도서":
				bible_short_name = "傳"
				switch chap_string {
					case "":
						print_string = "傳道書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("傳","傳道書", chap_string, "1","kjv")
							default:
								print_string = Bible_print_string("傳","傳道書", chap_string, sec_string,"kjv")
						}
				}
			case "Song","Song of Solomon","歌","雅歌","So","so","sng","Sng","Песнь песней Соломона","Nhã-ca","아가":
				bible_short_name = "歌"
				switch chap_string {
					case "":
						print_string = "雅歌"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("歌","雅歌", chap_string, "1","kjv")
							default:
								print_string = Bible_print_string("歌","雅歌", chap_string, sec_string,"kjv")
						}
				}
			case "Is","Isaiah","賽","以賽","以賽亞","以賽亞書","Isa","isa","Книга пророка Исаии","Ê-sai","イザヤ書","イザヤ","이사야":
				bible_short_name = "賽"
				switch chap_string {
					case "":
						print_string = "以賽亞書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("賽","以賽亞書", chap_string, "1","kjv")
							default:
								print_string = Bible_print_string("賽","以賽亞書", chap_string, sec_string,"kjv")
						}
				}
			case "Jer","Jeremiah","耶","耶利米","耶利米書","jer","예레미야","エレミヤ","エレミヤ書","Giê-rê-mi","Книга пророка Иеремии":
				bible_short_name = "耶"
				switch chap_string {
					case "":
						print_string = "耶利米書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("耶","耶利米書", chap_string, "1","kjv")
							default:
								print_string = Bible_print_string("耶","耶利米書", chap_string, sec_string,"kjv")
						}
				}
			case "Lam","Lamentations","哀","哀歌","耶利米哀歌","La","lam","예레미야애가","Ca-thương","Плач Иеремии":
				bible_short_name = "哀"
				switch chap_string {
					case "":
						print_string = "耶利米哀歌"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("哀","耶利米哀歌", chap_string, "1","kjv")
							default:
								print_string = Bible_print_string("哀","耶利米哀歌", chap_string, sec_string,"kjv")
						}
				}
			case "Ezek","Ezekiel","結","以西結","以西結書","Eze","eze","에스겔","エゼキエル書","エゼキエル","Ê-xê-chi-ên","Книга пророка Иезекииля":
				bible_short_name = "結"
				switch chap_string {
					case "":
						print_string = "以西結書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("結","以西結書", chap_string, "1","kjv")
							default:
								print_string = Bible_print_string("結","以西結書", chap_string, sec_string,"kjv")
						}
				}
			case "Dan","Daniel","但","但以理","但以理書","Da","da","Книга пророка Даниила","Đa-ni-ên","ダニエル書","ダニエル","다니엘":
				bible_short_name = "但"
				switch chap_string {
					case "":
						print_string = "但以理書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("但","但以理書", chap_string, "1","kjv")
							default:
								print_string = Bible_print_string("但","但以理書", chap_string, sec_string,"kjv")
						}
				}
			case "Hos","Hosea","何","何西","何西阿","何西阿書","Ho","ho","Книга пророка Осии","Ô-sê","ホセア書","ホセア","호세아":
				bible_short_name = "何"
				switch chap_string {
					case "":
						print_string = "何西阿書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("何","何西阿書", chap_string, "1","kjv")
							default:
								print_string = Bible_print_string("何","何西阿書", chap_string, sec_string,"kjv")
						}
				}
			case "Joel","珥","約珥","約珥書","Joe","joe","Книга пророка Иоиля","Giô-ên","ヨエル書","ヨエル","요엘":
				bible_short_name = "珥"
				switch chap_string {
					case "":
						print_string = "約珥書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("珥","約珥書", chap_string, "1","kjv")
							default:
								print_string = Bible_print_string("珥","約珥書", chap_string, sec_string,"kjv")
						}
				}
			case "Amos","摩","阿摩司書","Am","am","Книга пророка Амоса","A-mốt","アモス書","アモス","아모스":
				bible_short_name = "摩"
				switch chap_string {
					case "":
						print_string = "阿摩司書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("摩","阿摩司書", chap_string, "1","kjv")
							default:
								print_string = Bible_print_string("摩","阿摩司書", chap_string, sec_string,"kjv")
						}
				}
			case "Obad","Obadiah","俄","俄巴底亞","俄巴底亞書","Ob","ob","오바댜","オバデヤ書","オバデヤ","Áp-đia","Книга пророка Авдия":
				bible_short_name = "俄"
				switch chap_string {
					case "":
						print_string = "俄巴底亞書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("俄","俄巴底亞書", chap_string, "1","kjv")
							default:
								print_string = Bible_print_string("俄","俄巴底亞書", chap_string, sec_string,"kjv")
						}
				}
			case "Jon","Jonah","拿","約拿","約拿書","jon","요나","ヨナ書","ヨナ","Giô-na","Книга пророка Ионы":
				bible_short_name = "拿"
				switch chap_string {
					case "":
						print_string = "約拿書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("拿","約拿書", chap_string, "1","kjv")
							default:
								print_string = Bible_print_string("拿","約拿書", chap_string, sec_string,"kjv")
						}
				}
			case "Micah","彌","彌迦","彌迦書","Mic","mic","Книга пророка Михея","Mi-chê","ミカ書","ミカ","미가":
				bible_short_name = "彌"
				switch chap_string {
					case "":
						print_string = "彌迦書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("彌","彌迦書", chap_string, "1","kjv")
							default:
								print_string = Bible_print_string("彌","彌迦書", chap_string, sec_string,"kjv")
						}
				}
			case "Nah","Nahum","鴻","那鴻","那鴻書","Na","na","Книга пророка Наума","Na-hum","ナホム書","ナホム","나훔":
				bible_short_name = "鴻"
				switch chap_string {
					case "":
						print_string = "那鴻書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("鴻","那鴻書", chap_string, "1","kjv")
							default:
								print_string = Bible_print_string("鴻","那鴻書", chap_string, sec_string,"kjv")
						}
				}
			case "Habakkuk","哈","哈巴","哈巴谷","哈巴谷書","Hab","hab","Книга пророка Аввакума","Ha-ba-cúc","ハバクク書","ハバクク","ハバ","クク","ハバ書","하박국":
				bible_short_name = "哈"
				switch chap_string {
					case "":
						print_string = "哈巴谷書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("哈","哈巴谷書", chap_string, "1","kjv")
							default:
								print_string = Bible_print_string("哈","哈巴谷書", chap_string, sec_string,"kjv")
						}
				}
			case "Zeph","Zephaniah","番","西番雅","西番雅書","Zep","zep","스바냐","ゼパニヤ書","ゼパニヤ","Sô-phô-ni","Книга пророка Софонии":
				bible_short_name = "番"
				switch chap_string {
					case "":
						print_string = "西番雅書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("番","西番雅書", chap_string, "1","kjv")
							default:
								print_string = Bible_print_string("番","西番雅書", chap_string, sec_string,"kjv")
						}
				}
			case "Haggai","該","哈該","哈該書","Hag","hag","학개","ハガイ書","ハガイ","A-ghê","Книга пророка Аггея":
				bible_short_name = "該"
				switch chap_string {
					case "":
						print_string = "哈該書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("該","哈該書", chap_string, "1","kjv")
							default:
								print_string = Bible_print_string("該","哈該書", chap_string, sec_string,"kjv")
						}
				}
			case "Zech","Zechariah","亞","撒迦利亞","撒迦利亞書","Zec","zec","Книга пророка Захарии","Xa-cha-ri","스가랴","ゼカリヤ書","ゼカリヤ":
				bible_short_name = "亞"
				switch chap_string {
					case "":
						print_string = "撒迦利亞書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("亞","撒迦利亞書", chap_string, "1","kjv")
							default:
								print_string = Bible_print_string("亞","撒迦利亞書", chap_string, sec_string,"kjv")
						}
				}
			case "Malachi","瑪","瑪拉","瑪拉基","瑪拉基書","Mal","mal","말라기","マラキ書","マラキ","Ma-la-chi","Книга пророка Малахии":
				bible_short_name = "瑪"
				switch chap_string {
					case "":
						print_string = "瑪拉基書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("瑪","瑪拉基書", chap_string, "1","kjv")
							default:
								print_string = Bible_print_string("瑪","瑪拉基書", chap_string, sec_string,"kjv")
						}
				}
			case "Matt","Matthew","太","馬太","馬太福音","Mt","mt","마태복음","マタイによる福音書","マタイ","マタイによる","Ma-thi-ơ","От Матфея святое благовествование":
				bible_short_name = "太"
				switch chap_string {
					case "":
						print_string = "馬太福音"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("太","馬太福音", chap_string, "1","kjv")
							default:
								print_string = Bible_print_string("太","馬太福音", chap_string, sec_string,"kjv")
						}
				}
			case "Mark","可","馬可","馬可福音","Mr","mr","マルコによる福音書","マルコ","マルコによる","마가복음","Mác","От Марка святое благовествование":
				bible_short_name = "可"
				switch chap_string {
					case "":
						print_string = "馬可福音"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("可","馬可福音", chap_string, "1","kjv")
							default:
								print_string = Bible_print_string("可","馬可福音", chap_string, sec_string,"kjv")
						}
				}
			case "Luke","路","路加","路加福音","Lu","lu","От Луки святое благовествование","Lu-ca","ルカによる福音書","ルカ","ルカによる","누가복음":
				bible_short_name = "路"
				switch chap_string {
					case "":
						print_string = "路加福音"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("路","路加福音", chap_string, "1","kjv")
							default:
								print_string = Bible_print_string("路","路加福音", chap_string, sec_string,"kjv")
						}
				}
			case "John","約","約翰","約翰福音","Joh","joh","От Иоанна святое благовествование","Giăng","ヨハネによる福音書","ヨハネ","ヨハネによる","요한복음":
				bible_short_name = "約"
				switch chap_string {
					case "":
						print_string = "約翰福音"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("約","約翰福音", chap_string, "1","kjv")
							default:
								print_string = Bible_print_string("約","約翰福音", chap_string, sec_string,"kjv")
						}
				}
			case "Acts","徒","使徒","使徒行傳","Ac","ac","Деяния святых апостолов","Công-vụ Các Sứ-đồ","使徒行伝","사도행전":
				bible_short_name = "徒"
				switch chap_string {
					case "":
						print_string = "使徒行傳"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("徒","使徒行傳", chap_string, "1","kjv")
							default:
								print_string = Bible_print_string("徒","使徒行傳", chap_string, sec_string,"kjv")
						}
				}
			case "Rom","Romans","羅","羅馬","羅馬書","Ro","ro","Послание к Римлянам","Rô-ma","ローマ","ローマ人への手紙","로마서":
				bible_short_name = "羅"
				switch chap_string {
					case "":
						print_string = "羅馬書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("羅","羅馬書", chap_string, "1","kjv")
							default:
								print_string = Bible_print_string("羅","羅馬書", chap_string, sec_string,"kjv")
						}
				}
			case "1 Cor","First Corinthians","林前","哥林多前","哥林多前書","1Co","1co","Первое послание к Коринфянам","1 Cô-rinh-tô","コリント人への第一の手紙","コリント一","コリント人への第一","고린도전서":
				bible_short_name = "林前"
				switch chap_string {
					case "":
						print_string = "哥林多前書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("林前","哥林多前書", chap_string, "1","kjv")
							default:
								print_string = Bible_print_string("林前","哥林多前書", chap_string, sec_string,"kjv")
						}
				}
			case "2 Cor","Second Corinthians","林後","哥林多後","哥林多後書","2Co","2co","Второе послание к Коринфянам","2 Cô-rinh-tô","コリント人への第二の手紙","コリント二","コリント人への第二の","고린도후서":
				bible_short_name = "林後"
				switch chap_string {
					case "":
						print_string = "哥林多後書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("林後","哥林多後書", chap_string, "1","kjv")
							default:
								print_string = Bible_print_string("林後","哥林多後書", chap_string, sec_string,"kjv")
						}
				}
			case "Gal","Galatians","加","加拉太","加拉太書","Ga","ga","Послание к Галатам","Ga-la-ti","ガラテヤ","ガラテヤ人への手紙","갈라디아서":
				bible_short_name = "加"
				switch chap_string {
					case "":
						print_string = "加拉太書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("加","加拉太書", chap_string, "1","kjv")
							default:
								print_string = Bible_print_string("加","加拉太書", chap_string, sec_string,"kjv")
						}
				}
			case "Ephesians","弗","以弗所","以弗所書","Eph","eph","Послание к Ефесянам","Ê-phê-sô","エペソ人への手紙","エペソ","エペソ人","エペソ人の手紙","에베소서":
				bible_short_name = "弗"
				switch chap_string {
					case "":
						print_string = "以弗所書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("弗","以弗所書", chap_string, "1","kjv")
							default:
								print_string = Bible_print_string("弗","以弗所書", chap_string, sec_string,"kjv")
						}
				}
			case "Phil","Philippians","腓","腓立","腓立比","腓立比書","Php","php","빌립보서","ピリピ","ピリピ人.ピリピ人への手紙","Послание к Филиппийцам","Phi-líp":
				bible_short_name = "腓"
				switch chap_string {
					case "":
						print_string = "腓立比書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("腓","腓立比書", chap_string, "1","kjv")
							default:
								print_string = Bible_print_string("腓","腓立比書", chap_string, sec_string,"kjv")
						}
				}
			case "Col","col","Colossians","西","歌羅西","歌羅","歌羅西書","Послание к Колоссянам","Cô-lô-se","コロサイ人への手紙","コロサイ","コロ","골로새서":
				bible_short_name = "西"
				switch chap_string {
					case "":
						print_string = "歌羅西書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("西","歌羅西書", chap_string, "1","kjv")
							default:
								print_string = Bible_print_string("西","歌羅西書", chap_string, sec_string,"kjv")
						}
				}
			case "1 Thess","First Thessalonians","帖前","帖撒羅尼迦前","帖撒羅尼迦前書","1Th","1th","데살로니가전서","テサロニケ人への第一の手紙","テサ一","テサロニケ一","1 Tê-sa-lô-ni-ca","Первое послание к Фессалоникийцам (Солунянам)":
				bible_short_name = "帖前"
				switch chap_string {
					case "":
						print_string = "帖撒羅尼迦前書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("帖前","帖撒羅尼迦前書", chap_string, "1","kjv")
							default:
								print_string = Bible_print_string("帖前","帖撒羅尼迦前書", chap_string, sec_string,"kjv")
						}
				}
			case "2 Thess","Second Thessalonians","帖後","帖撒羅尼迦後","帖撒羅尼迦後書","2Th","2th","데살로니가후서","テサロニケ人への第二の手紙","テサ二","テサロニケ二","2 Tê-sa-lô-ni-ca","Второе послание к Фессалоникийцам (Солунянам)":
				bible_short_name = "帖後"
				switch chap_string {
					case "":
						print_string = "帖撒羅尼迦後書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("帖後","帖撒羅尼迦後書", chap_string, "1","kjv")
							default:
								print_string = Bible_print_string("帖後","帖撒羅尼迦後書", chap_string, sec_string,"kjv")
						}
				}
			case "1 Tim","First Timothy","提前","提摩太前","提摩太前書","1Ti","1ti","Первое послание к Тимофею","1 Ti-mô-thê","テモテヘの第一の手紙","テモテ一","디모데전서":
				bible_short_name = "提前"
				switch chap_string {
					case "":
						print_string = "提摩太前書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("提前","提摩太前書", chap_string, "1","kjv")
							default:
								print_string = Bible_print_string("提前","提摩太前書", chap_string, sec_string,"kjv")
						}
				}
			case "2 Tim","Second Timothy","提後","提摩太後","提摩太後書","2Ti","2ti","Второе послание к Тимофею","2 Ti-mô-thê","テモテヘの第二の手紙","テモテ二","디모데후서":
				bible_short_name = "提後"
				switch chap_string {
					case "":
						print_string = "提摩太後書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("提後","提摩太後書", chap_string, "1","kjv")
							default:
								print_string = Bible_print_string("提後","提摩太後書", chap_string, sec_string,"kjv")
						}
				}
			case "Titus","多","提多","提多書","Tit","tit","Послание к Титу","Tít","テトスヘの手紙","テトス","디도서":
				bible_short_name = "多"
				switch chap_string {
					case "":
						print_string = "提多書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("多","提多書", chap_string, "1","kjv")
							default:
								print_string = Bible_print_string("多","提多書", chap_string, sec_string,"kjv")
						}
				}
			case "Philem","Philemon","門","腓利","腓利門","腓利門書","Phm","phm","Послание к Филимону","Phi-lê-môn","ピレモンヘの手紙","ピレモン","빌레몬서":
				bible_short_name = "門"
				switch chap_string {
					case "":
						print_string = "腓利門書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("門","腓利門書", chap_string, "1","kjv")
							default:
								print_string = Bible_print_string("門","腓利門書", chap_string, sec_string,"kjv")
						}
				}
			case "Heb","Hebrews","來","希伯來","希伯來書","heb","Послание к Евреям","Hê-bơ-rơ","ヘブル人への手紙","ヘブル","히브리서":
				bible_short_name = "來"
				switch chap_string {
					case "":
						print_string = "希伯來書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("來","希伯來書", chap_string, "1","kjv")
							default:
								print_string = Bible_print_string("來","希伯來書", chap_string, sec_string,"kjv")
						}
				}
			case "James","雅","雅各","雅各書","Jas","jas","Послание Иакова","Gia-cơ","ヤコブの手紙","ヤコブ","야고보서":
				bible_short_name = "雅"
				switch chap_string {
					case "":
						print_string = "雅各書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("雅","雅各書", chap_string, "1","kjv")
							default:
								print_string = Bible_print_string("雅","雅各書", chap_string, sec_string,"kjv")
						}
				}
			case "1 Pet","First Peter","彼前","彼得前","彼得前書","1Pe","1pe","Первое послание Петра","1 Phi-e-rơ","ペテロの第一の手紙","ペテロ一","베드로전서":
				bible_short_name = "彼前"
				switch chap_string {
					case "":
						print_string = "彼得前書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("彼前","彼得前書", chap_string, "1","kjv")
							default:
								print_string = Bible_print_string("彼前","彼得前書", chap_string, sec_string,"kjv")
						}
				}
			case "2 Pet","Second Peter","彼後","彼得後","彼得後書","2Pe","2pe","Второе послание Петра","2 Phi-e-rơ","ペテロの第二の手紙","ペテロ","베드로후서":
				bible_short_name = "彼後"
				switch chap_string {
					case "":
						print_string = "彼得後書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("彼後","彼得後書", chap_string, "1","kjv")
							default:
								print_string = Bible_print_string("彼後","彼得後書", chap_string, sec_string,"kjv")
						}
				}
			case "1 John","First John","約一","約翰一書","約翰1","約翰1書","1Jo","1jo","Первое послание Иоанна","1 Giăng","ヨハネの第一の手紙","ヨハネ一","요한일서":
				bible_short_name = "約一"
				switch chap_string {
					case "":
						print_string = "約翰一書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("約一","約翰一書", chap_string, "1","kjv")
							default:
								print_string = Bible_print_string("約一","約翰一書", chap_string, sec_string,"kjv")
						}
				}
			case "2 John","second John","約二","約翰二書","約翰2","約翰2書","2Jo","Второе послание Иоанна","2 Giăng","ヨハネの第二の手紙","ヨハネ二","요한2서":
				bible_short_name = "約二"
				switch chap_string {
					case "":
						print_string = "約翰二書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("約二","約翰二書", chap_string, "1","kjv")
							default:
								print_string = Bible_print_string("約二","約翰二書", chap_string, sec_string,"kjv")
						}
				}
			case "3 John","Third John","約三","約翰三書","約翰3","約翰3書","3Jo","3jo","Третье послание Иоанна","3 Giăng","ヨハネの第三の手紙","ヨハネ三","요한3서":
				bible_short_name = "約三"
				switch chap_string {
					case "":
						print_string = "約翰三書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("約三","約翰三書", chap_string, "1","kjv")
							default:
								print_string = Bible_print_string("約三","約翰三書", chap_string, sec_string,"kjv")
						}
				}
			case "Jude","猶","猶大","猶大書","jude","Послание Иуды","Giu-đe","ユダの手紙","ユダ","유다서":
				bible_short_name = "猶"
				switch chap_string {
					case "":
						print_string = "猶大書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("猶","猶大書", chap_string, "1","kjv")
							default:
								print_string = Bible_print_string("猶","猶大書", chap_string, sec_string,"kjv")
						}
				}
			case "Rev","Revelation","啟","啟示","啟示錄","Re","re","ｒｅ","Ｒｅ","rev","Откровение ап. Иоанна Богослова (Апокалипсис)","Khải-huyền","ヨハネの黙示録","黙示録","요한계시록":
				bible_short_name = "啟"
				switch chap_string {
					case "":
						print_string = "啟示錄"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("啟","啟示錄", chap_string, "1","kjv")
							default:
								print_string = Bible_print_string("啟","啟示錄", chap_string, sec_string,"kjv")
						}
				}
			default:
				//print_string = "你是要找 " +  reg.ReplaceAllString(text, "$3") + " 對嗎？\n對不起，我還沒學呢...\n"
				print_string = "聖經"
		}
	case "聖經","bible","Bible","ｂｉｂｌｅ","Ｂｉｂｌｅ":
		print_string = text + "？\n抱歉目前找不到"
		//bible_say := "有喔！有喔！你在找這個對吧！？\n"
		//view-source:http://bible.fhl.net/json/listall.html
		//----JavaScript 偷吃步法（拿 JavaScript 當預處理XD）
							// function getGOGOGO(s_name,fullname){
							//     var lang = 'nstrunv'; //jp,kjv
							//     var str = `\n			case "` + fullname + `","` + s_name + `":
							// 	switch chap_string {
							// 		case "":
							// 			print_string = "` + fullname + `" //不知章節的時候
							// 		default:
							// 			switch sec_string {
							// 				case "":	//不知節的時候
							// 				default:
							// 					print_string = Bible_print_string("` + s_name + `","` + fullname + `", chap_string, sec_string,"` + lang + `")
							// 			}
							// 	}`;
							//     return str;
							// }
							// console.info(getGOGOGO('利','利未記') + getGOGOGO('民','民數記'));


			// function getGOGOGO(s_name,fullname,all_name){
			// 	var lang = 'nstrunv'; //nstrunv,jp,kjv(英文),korean,russian(俄文),vietnamese(越南)
			// 	// rcuv(和合本修訂版 2010)
			// 	// ncv(中文新譯本 2010)
			// 	// tcv(現代中文譯本修訂版 1997)
			// 	// wlunv(文言文（深文理和合本）)
			// 	// sgebklhl(台語（全民台語聖經漢羅）)
			// 	// sgebklcl(台語（全民台語聖經全羅）)
			// 	// bklhl(台語（巴克禮漢羅）)
			// 	// bklcl(台語（巴克禮全羅）)
			// 	// prebklhl(台語（馬雅各漢羅）)
			// 	// prebklcl(台語（馬雅各全羅）)
			// 	// hakka(客家話（新約）)
			
			// 	// bbe(英文 BBE（簡易英文譯本）)
			// 	// web(英文 WEB（環球英文聖經）)
			// 	// asv(英文 ASV（美國標準譯本）)
			// 	// darby(英文 Darby 1890)
			// 	// erv(英文 ERV（English Revised Version 英國修訂譯本）)
			// 	// lxx(舊約 古譯本 七十士譯本)
			// 	// bhs(舊約 馬索拉原文)
			// 	var str = `\n			case "` + all_name.replace(/,/g,`","`) + `":
			// 	bible_short_name = "` + s_name +  `"
			// 	switch chap_string {
			// 		case "":
			// 			print_string = "` + fullname + `"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
			// 		default:
			// 			switch sec_string {
			// 				case "":	//不知節的時候，知章
			// 					print_string = Bible_print_string("` + s_name + `","` + fullname + `", chap_string, "1","` + lang + `")
			// 				default:
			// 					print_string = Bible_print_string("` + s_name + `","` + fullname + `", chap_string, sec_string,"` + lang + `")
			// 			}
			// 	}`;
			// 	return str;
			// }

			// // 多國語言並列的版本	//ncv = 《聖經新譯本》©1976, 1992, 1999, 2001, 2005, 2010版權屬於環球聖經公會
			// function getGOGOGO(s_name,fullname,all_name){
			// 	//var lang = 'nstrunv'; //nstrunv,jp,kjv(英文),korean,russian(俄文),vietnamese(越南) ,tcv=現代中文譯本修訂版(©1997版權屬於聯合聖經公會，由台灣聖經公會授權信望愛站使用。)
			// 	var str = `\n			case "` + all_name.replace(/,/g,`","`) + `":
			// 	bible_short_name = "` + s_name +  `"
			// 	switch chap_string {
			// 		case "":
			// 			print_string = "` + fullname + `" //不知章節的時候 //用來等觸發 UI 及特別說明文字
			// 		default:
			// 			switch sec_string {
			// 				case "":	//不知節的時候，知章
			// 					print_string = Bible_print_all_string("` + s_name + `","` + fullname + `", chap_string, "1")
			// 				default:
			// 					print_string = Bible_print_all_string("` + s_name + `","` + fullname + `", chap_string, sec_string)
			// 			}
			// 	}`;
			// 	return str;
			// }

			// // 全版本並列並列的版本	//ncv = 《聖經新譯本》©1976, 1992, 1999, 2001, 2005, 2010 版權屬於環球聖經公會
			// function getGOGOGO(s_name,fullname,all_name){
			// 	//var lang = 'nstrunv'; //nstrunv,jp,kjv(英文),korean,russian(俄文),vietnamese(越南) ,tcv=現代中文譯本修訂版(©1997版權屬於聯合聖經公會，由台灣聖經公會授權信望愛站使用。)
			// 	var str = `\n			case "` + all_name.replace(/,/g,`","`) + `":
			// 	bible_short_name = "` + s_name +  `"
			// 	switch chap_string {
			// 		case "":
			// 			print_string = "` + fullname + `" //不知章節的時候 //用來等觸發 UI 及特別說明文字
			// 		default:
			// 			switch sec_string {
			// 				case "":	//不知節的時候，知章
			// 					print_string = Bible_print_all_var_string("` + s_name + `","` + fullname + `", chap_string, "1")
			// 				default:
			// 					print_string = Bible_print_all_var_string("` + s_name + `","` + fullname + `", chap_string, sec_string)
			// 			}
			// 	}`;
			// 	return str;
			// }

			//執行量產

			// clear()
			// var end = getGOGOGO('創','創世記','Gen,Genesis,創,創世,創世紀,創世記,Ge,ge,gen,창세기,Sáng-thế Ký,Бытие') + getGOGOGO('出','出埃及記','ex,Ex,Exodus,埃及,出,出埃及,出埃及記,출애굽기,エジプト,出エジプト,出エジプト記,Xuất Ê-díp-tô Ký,Исход') + 
			// 		  getGOGOGO('利','利未記','Lev,Leviticus,利,利未,利未記,Le,le,Левит,Lê-vi Ký,レビ記,レビ,레위기') + getGOGOGO('民','民數記','Num,Numbers,民,民數,民數記,Nu,nu,민수기,民数,民数記,Dân-số Ký,Числа') +
			// 		  getGOGOGO('申','申命記','Deut,Deuteronomy,申,申命,申命記,De,de,신명기,Phục-truyền Luật-lệ Ký,Второзаконие') +  getGOGOGO('書','約書亞記','Josh,Joshua,約書亞,約書亞記,Jos,jos,여호수아,ヨシュア記,ヨシュア,Giô-suê,Книга Иисуса Навина');
			// end += getGOGOGO('士','士師記','Judg,Judges,士,士師,士師記,Jud,jud,jdg,Jdg,Книга Судей израилевых,Các Quan Xét,사사기') + getGOGOGO('得','路得記','Ruth,路得,路得記,Ru,ru,Rut,rut,룻기,ルツ,ルツ記,Ru-tơ,Книга Руфи');
			// end += getGOGOGO('撒上','撒母耳記上','1 Sam,First Samuel,撒上,撒母耳記上,1Sa,1sa,サムエル記上,サムエル上,サム上,사무엘상,1 Sa-mu-ên,Первая книга Царств') + getGOGOGO('撒下','撒母耳記下','2 Sam,Second Samuel,撒下,撒母耳記下,2Sa,2sa,사무엘하,サムエル記下,サムエル下,サム下,2 Sa-mu-ên,Вторая книга Царств');	//10 all=66

			// end += getGOGOGO('王上','列王紀上','1 Kin,First Kings,王上,列王上,列王紀上,列王記上,1Ki,1ki,열왕기상,Третья книга Царств,1 Các Vua') + getGOGOGO('王下','列王紀下','2 Kin,Second Kings,王下,列王下,列王記下,列王紀下,2Ki,2ki,열왕기하,2 Các Vua,Четвертая книга Царств');	//12
			// end += getGOGOGO('代上','歷代志上','1 Chr,First Chronicles,歷上,代上,歷代志上,歷代上,1Ch,1ch,Первая книга Паралипоменон,1 Sử-ký,歴上,歴代上,歴代志上,역대상') + getGOGOGO('代下','歷代志下','2 Chr,Second Chronicles,代下,歷下,歷代下,歷代志下,2Ch,2ch,역대하,歴代志下,歴代下,歴下,2 Sử-ký,Вторая книга Паралипоменон');	//14
			// end += getGOGOGO('拉','以斯拉記','Ezra,拉,以斯拉,以斯拉記,Ezr,ezr,Первая книга Ездры,Ê-xơ-ra,エズラ,エズラ記,에스라') + getGOGOGO('尼','尼希米記','Neh,Nehemiah,尼,尼希米,尼希米記,Ne,ne,느헤미야,ネヘミヤ書,ネヘミヤ,Nê-hê-mi,Книга Неемии');
			// end += getGOGOGO('斯','以斯帖記','Esth,Esther,斯,以斯帖,以斯帖記,Es,est,Есфирь,Ê-xơ-tê,エステル,エステル記,에스더') + getGOGOGO('伯','約伯記','Job,job,伯,約伯,約伯記,Книга Иова,Gióp,ヨブ,ヨブ記,욥기');
			// end += getGOGOGO('詩','詩篇','Ps,Psalms,詩,詩篇,ps,시편,Thi-thiên,Псалтирь') + getGOGOGO('箴','箴言','Prov,Proverbs,箴,箴言,Pr,pr,Притчи Соломона,Châm-ngôn,잠언');	//20

			// end += getGOGOGO('傳','傳道書','Eccl,Ecclesiastes,傳,傳道,傳道書,Ec,ec,Книга Екклезиаста,или Проповедника,Truyền-đạo,伝道の書,伝道,伝,伝道書,전도서') + getGOGOGO('歌','雅歌','Song,Song of Solomon,歌,雅歌,So,so,sng,Sng,Песнь песней Соломона,Nhã-ca,아가');
			// end += getGOGOGO('賽','以賽亞書','Is,Isaiah,賽,以賽,以賽亞,以賽亞書,Isa,isa,Книга пророка Исаии,Ê-sai,イザヤ書,イザヤ,이사야') + getGOGOGO('耶','耶利米書','Jer,Jeremiah,耶,耶利米,耶利米書,jer,예레미야,エレミヤ,エレミヤ書,Giê-rê-mi,Книга пророка Иеремии');
			// end += getGOGOGO('哀','耶利米哀歌','Lam,Lamentations,哀,哀歌,耶利米哀歌,La,lam,예레미야애가,Ca-thương,Плач Иеремии') + getGOGOGO('結','以西結書','Ezek,Ezekiel,結,以西結,以西結書,Eze,eze,에스겔,エゼキエル書,エゼキエル,Ê-xê-chi-ên,Книга пророка Иезекииля');
			// end += getGOGOGO('但','但以理書','Dan,Daniel,但,但以理,但以理書,Da,da,Книга пророка Даниила,Đa-ni-ên,ダニエル書,ダニエル,다니엘') + getGOGOGO('何','何西阿書','Hos,Hosea,何,何西,何西阿,何西阿書,Ho,ho,Книга пророка Осии,Ô-sê,ホセア書,ホセア,호세아');
			// end += getGOGOGO('珥','約珥書','Joel,珥,約珥,約珥書,Joe,joe,Книга пророка Иоиля,Giô-ên,ヨエル書,ヨエル,요엘') + getGOGOGO('摩','阿摩司書','Amos,摩,阿摩司書,Am,am,Книга пророка Амоса,A-mốt,アモス書,アモス,아모스');	//30

			// end += getGOGOGO('俄','俄巴底亞書','Obad,Obadiah,俄,俄巴底亞,俄巴底亞書,Ob,ob,오바댜,オバデヤ書,オバデヤ,Áp-đia,Книга пророка Авдия') + getGOGOGO('拿','約拿書','Jon,Jonah,拿,約拿,約拿書,jon,요나,ヨナ書,ヨナ,Giô-na,Книга пророка Ионы');
			// end += getGOGOGO('彌','彌迦書','Micah,彌,彌迦,彌迦書,Mic,mic,Книга пророка Михея,Mi-chê,ミカ書,ミカ,미가') + getGOGOGO('鴻','那鴻書','Nah,Nahum,鴻,那鴻,那鴻書,Na,na,Книга пророка Наума,Na-hum,ナホム書,ナホム,나훔');
			// end += getGOGOGO('哈','哈巴谷書','Habakkuk,哈,哈巴,哈巴谷,哈巴谷書,Hab,hab,Книга пророка Аввакума,Ha-ba-cúc,ハバクク書,ハバクク,ハバ,クク,ハバ書,하박국') + getGOGOGO('番','西番雅書','Zeph,Zephaniah,番,西番雅,西番雅書,Zep,zep,스바냐,ゼパニヤ書,ゼパニヤ,Sô-phô-ni,Книга пророка Софонии');
			// end += getGOGOGO('該','哈該書','Haggai,該,哈該,哈該書,Hag,hag,학개,ハガイ書,ハガイ,A-ghê,Книга пророка Аггея') + getGOGOGO('亞','撒迦利亞書','Zech,Zechariah,亞,撒迦利亞,撒迦利亞書,Zec,zec,Книга пророка Захарии,Xa-cha-ri,스가랴,ゼカリヤ書,ゼカリヤ');
			// end += getGOGOGO('瑪','瑪拉基書','Malachi,瑪,瑪拉,瑪拉基,瑪拉基書,Mal,mal,말라기,マラキ書,マラキ,Ma-la-chi,Книга пророка Малахии') + getGOGOGO('太','馬太福音','Matt,Matthew,太,馬太,馬太福音,Mt,mt,마태복음,マタイによる福音書,マタイ,マタイによる,Ma-thi-ơ,От Матфея святое благовествование');	//40

			// end += getGOGOGO('可','馬可福音','Mark,可,馬可,馬可福音,Mr,mr,マルコによる福音書,マルコ,マルコによる,마가복음,Mác,От Марка святое благовествование') + getGOGOGO('路','路加福音','Luke,路,路加,路加福音,Lu,lu,От Луки святое благовествование,Lu-ca,ルカによる福音書,ルカ,ルカによる,누가복음');
			// end += getGOGOGO('約','約翰福音','John,約,約翰,約翰福音,Joh,joh,От Иоанна святое благовествование,Giăng,ヨハネによる福音書,ヨハネ,ヨハネによる,요한복음') + getGOGOGO('徒','使徒行傳','Acts,徒,使徒,使徒行傳,Ac,ac,Деяния святых апостолов,Công-vụ Các Sứ-đồ,使徒行伝,사도행전');
			// end += getGOGOGO('羅','羅馬書','Rom,Romans,羅,羅馬,羅馬書,Ro,ro,Послание к Римлянам,Rô-ma,ローマ,ローマ人への手紙,로마서') + getGOGOGO('林前','哥林多前書','1 Cor,First Corinthians,林前,哥林多前,哥林多前書,1Co,1co,Первое послание к Коринфянам,1 Cô-rinh-tô,コリント人への第一の手紙,コリント一,コリント人への第一,고린도전서');
			// end += getGOGOGO('林後','哥林多後書','2 Cor,Second Corinthians,林後,哥林多後,哥林多後書,2Co,2co,Второе послание к Коринфянам,2 Cô-rinh-tô,コリント人への第二の手紙,コリント二,コリント人への第二の,고린도후서') + getGOGOGO('加','加拉太書','Gal,Galatians,加,加拉太,加拉太書,Ga,ga,Послание к Галатам,Ga-la-ti,ガラテヤ,ガラテヤ人への手紙,갈라디아서');
			// end += getGOGOGO('弗','以弗所書','Ephesians,弗,以弗所,以弗所書,Eph,eph,Послание к Ефесянам,Ê-phê-sô,エペソ人への手紙,エペソ,エペソ人,エペソ人の手紙,에베소서') + getGOGOGO('腓','腓立比書','Phil,Philippians,腓,腓立,腓立比,腓立比書,Php,php,빌립보서,ピリピ,ピリピ人.ピリピ人への手紙,Послание к Филиппийцам,Phi-líp');	//50

			// end += getGOGOGO('西','歌羅西書','Col,col,Colossians,西,歌羅西,歌羅,歌羅西書,Послание к Колоссянам,Cô-lô-se,コロサイ人への手紙,コロサイ,コロ,골로새서') + getGOGOGO('帖前','帖撒羅尼迦前書','1 Thess,First Thessalonians,帖前,帖撒羅尼迦前,帖撒羅尼迦前書,1Th,1th,데살로니가전서,テサロニケ人への第一の手紙,テサ一,テサロニケ一,1 Tê-sa-lô-ni-ca,Первое послание к Фессалоникийцам (Солунянам)');
			// end += getGOGOGO('帖後','帖撒羅尼迦後書','2 Thess,Second Thessalonians,帖後,帖撒羅尼迦後,帖撒羅尼迦後書,2Th,2th,데살로니가후서,テサロニケ人への第二の手紙,テサ二,テサロニケ二,2 Tê-sa-lô-ni-ca,Второе послание к Фессалоникийцам (Солунянам)') + getGOGOGO('提前','提摩太前書','1 Tim,First Timothy,提前,提摩太前,提摩太前書,1Ti,1ti,Первое послание к Тимофею,1 Ti-mô-thê,テモテヘの第一の手紙,テモテ一,디모데전서');
			// end += getGOGOGO('提後','提摩太後書','2 Tim,Second Timothy,提後,提摩太後,提摩太後書,2Ti,2ti,Второе послание к Тимофею,2 Ti-mô-thê,テモテヘの第二の手紙,テモテ二,디모데후서') + getGOGOGO('多','提多書','Titus,多,提多,提多書,Tit,tit,Послание к Титу,Tít,テトスヘの手紙,テトス,디도서');
			// end += getGOGOGO('門','腓利門書','Philem,Philemon,門,腓利,腓利門,腓利門書,Phm,phm,Послание к Филимону,Phi-lê-môn,ピレモンヘの手紙,ピレモン,빌레몬서') + getGOGOGO('來','希伯來書','Heb,Hebrews,來,希伯來,希伯來書,heb,Послание к Евреям,Hê-bơ-rơ,ヘブル人への手紙,ヘブル,히브리서');
			// end += getGOGOGO('雅','雅各書','James,雅,雅各,雅各書,Jas,jas,Послание Иакова,Gia-cơ,ヤコブの手紙,ヤコブ,야고보서') + getGOGOGO('彼前','彼得前書','1 Pet,First Peter,彼前,彼得前,彼得前書,1Pe,1pe,Первое послание Петра,1 Phi-e-rơ,ペテロの第一の手紙,ペテロ一,베드로전서');	//60

			// end += getGOGOGO('彼後','彼得後書','2 Pet,Second Peter,彼後,彼得後,彼得後書,2Pe,2pe,Второе послание Петра,2 Phi-e-rơ,ペテロの第二の手紙,ペテロ,베드로후서') + getGOGOGO('約一','約翰一書','1 John,First John,約一,約翰一書,約翰1,約翰1書,1Jo,1jo,Первое послание Иоанна,1 Giăng,ヨハネの第一の手紙,ヨハネ一,요한일서');
			// end += getGOGOGO('約二','約翰二書','2 John,second John,約二,約翰二書,約翰2,約翰2書,2Jo,Второе послание Иоанна,2 Giăng,ヨハネの第二の手紙,ヨハネ二,요한2서') + getGOGOGO('約三','約翰三書','3 John,Third John,約三,約翰三書,約翰3,約翰3書,3Jo,3jo,Третье послание Иоанна,3 Giăng,ヨハネの第三の手紙,ヨハネ三,요한3서');
			// end += getGOGOGO('猶','猶大書','Jude,猶,猶大,猶大書,jude,Послание Иуды,Giu-đe,ユダの手紙,ユダ,유다서') + getGOGOGO('啟','啟示錄','Rev,Revelation,啟,啟示,啟示錄,Re,re,ｒｅ,Ｒｅ,rev,Откровение ап. Иоанна Богослова (Апокалипсис),Khải-huyền,ヨハネの黙示録,黙示録,요한계시록');	//66
			// console.info(end);
		//----JavaScript 偷吃步法
		log.Print(reg.ReplaceAllString(text, "$3"))
		switch reg.ReplaceAllString(text, "$3") {
			case "Gen","Genesis","創","創世","創世紀","創世記","Ge","ge","gen","창세기","Sáng-thế Ký","Бытие":
				bible_short_name = "創"
				switch chap_string {
					// case "1":
					// 	switch sec_string {
					// 		case "1":
					// 			print_string = "起初　神創造天地。"
					// 		default:
					// 			print_string = reg.ReplaceAllString(text, "$3") + " " + reg.ReplaceAllString(text, "$4") + " : " + reg.ReplaceAllString(text, "$6")
					// 	}
					case "":
						print_string = "創世紀"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						// bible_json_string = HttpGET_("http://bible.fhl.net/json/qb.php?chineses=創&chap=" + chap_string + "&sec=" + sec_string)
						// if bible_json_string!="" {
						// 	bible_text_string = GetJson_bible(bible_json_string)
						// }
						// log.Print("GET = ")
						// log.Print("bible_json_string = " + bible_json_string)
						// log.Print("GetJson_bible = " + bible_text_string)
						// if bible_text_string != ""{
						// 	print_string = "[創世紀 " + chap_string + " : " +  sec_string + "]\n" + bible_text_string
						// }
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("創","創世紀", chap_string, "1","nstrunv")
							default:
								//print_string = reg.ReplaceAllString(text, "$3") + " " + reg.ReplaceAllString(text, "$4") + " : " + reg.ReplaceAllString(text, "$6")
								print_string = Bible_print_string("創","創世紀", chap_string, sec_string,"nstrunv")
						}
				}
				// bible_json_string = HttpGET_("http://bible.fhl.net/json/qb.php?chineses=創&chap=" + chap_string + "&sec=" + sec_string)
				// if bible_json_string!="" {
				// 	bible_text_string = GetJson_bible(bible_json_string)
				// }
				// log.Print("GET = ")
				// log.Print("bible_json_string = " + bible_json_string)
				// log.Print("GetJson_bible = " + bible_text_string)
				// if bible_text_string != ""{
				// 	print_string = "[創世紀 " + chap_string + " : " +  sec_string + "]\n" + bible_text_string
				// }
			case "ex","Ex","Exodus","埃及","出","出埃及","出埃及記","출애굽기","エジプト","出エジプト","出エジプト記","Xuất Ê-díp-tô Ký","Исход":
				bible_short_name = "出"
				switch chap_string {
					case "":
						print_string = "出埃及記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("出","出埃及記", chap_string, "1","nstrunv")
							default:
								print_string = Bible_print_string("出","出埃及記", chap_string, sec_string,"nstrunv")
						}
				}
			case "Lev","Leviticus","利","利未","利未記","Le","le","Левит","Lê-vi Ký","レビ記","レビ","레위기":
				bible_short_name = "利"
				switch chap_string {
					case "":
						print_string = "利未記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("利","利未記", chap_string, "1","nstrunv")
							default:
								print_string = Bible_print_string("利","利未記", chap_string, sec_string,"nstrunv")
						}
				}
			case "Num","Numbers","民","民數","民數記","Nu","nu","민수기","民数","民数記","Dân-số Ký","Числа":
				bible_short_name = "民"
				switch chap_string {
					case "":
						print_string = "民數記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("民","民數記", chap_string, "1","nstrunv")
							default:
								print_string = Bible_print_string("民","民數記", chap_string, sec_string,"nstrunv")
						}
				}
			case "Deut","Deuteronomy","申","申命","申命記","De","de","신명기","Phục-truyền Luật-lệ Ký","Второзаконие":
				bible_short_name = "申"
				switch chap_string {
					case "":
						print_string = "申命記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("申","申命記", chap_string, "1","nstrunv")
							default:
								print_string = Bible_print_string("申","申命記", chap_string, sec_string,"nstrunv")
						}
				}
			case "Josh","Joshua","約書亞","約書亞記","Jos","jos","여호수아","ヨシュア記","ヨシュア","Giô-suê","Книга Иисуса Навина":
				bible_short_name = "書"
				switch chap_string {
					case "":
						print_string = "約書亞記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("書","約書亞記", chap_string, "1","nstrunv")
							default:
								print_string = Bible_print_string("書","約書亞記", chap_string, sec_string,"nstrunv")
						}
				}
			case "Judg","Judges","士","士師","士師記","Jud","jud","jdg","Jdg","Книга Судей израилевых","Các Quan Xét","사사기":
				bible_short_name = "士"
				switch chap_string {
					case "":
						print_string = "士師記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("士","士師記", chap_string, "1","nstrunv")
							default:
								print_string = Bible_print_string("士","士師記", chap_string, sec_string,"nstrunv")
						}
				}
			case "Ruth","路得","路得記","Ru","ru","Rut","rut","룻기","ルツ","ルツ記","Ru-tơ","Книга Руфи":
				bible_short_name = "得"
				switch chap_string {
					case "":
						print_string = "路得記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("得","路得記", chap_string, "1","nstrunv")
							default:
								print_string = Bible_print_string("得","路得記", chap_string, sec_string,"nstrunv")
						}
				}
			case "1 Sam","First Samuel","撒上","撒母耳記上","1Sa","1sa","サムエル記上","サムエル上","サム上","사무엘상","1 Sa-mu-ên","Первая книга Царств":
				bible_short_name = "撒上"
				switch chap_string {
					case "":
						print_string = "撒母耳記上"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("撒上","撒母耳記上", chap_string, "1","nstrunv")
							default:
								print_string = Bible_print_string("撒上","撒母耳記上", chap_string, sec_string,"nstrunv")
						}
				}
			case "2 Sam","Second Samuel","撒下","撒母耳記下","2Sa","2sa","사무엘하","サムエル記下","サムエル下","サム下","2 Sa-mu-ên","Вторая книга Царств":
				bible_short_name = "撒下"
				switch chap_string {
					case "":
						print_string = "撒母耳記下"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("撒下","撒母耳記下", chap_string, "1","nstrunv")
							default:
								print_string = Bible_print_string("撒下","撒母耳記下", chap_string, sec_string,"nstrunv")
						}
				}
			case "1 Kin","First Kings","王上","列王上","列王紀上","列王記上","1Ki","1ki","열왕기상","Третья книга Царств","1 Các Vua":
				bible_short_name = "王上"
				switch chap_string {
					case "":
						print_string = "列王紀上"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("王上","列王紀上", chap_string, "1","nstrunv")
							default:
								print_string = Bible_print_string("王上","列王紀上", chap_string, sec_string,"nstrunv")
						}
				}
			case "2 Kin","Second Kings","王下","列王下","列王記下","列王紀下","2Ki","2ki","열왕기하","2 Các Vua","Четвертая книга Царств":
				bible_short_name = "王下"
				switch chap_string {
					case "":
						print_string = "列王紀下"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("王下","列王紀下", chap_string, "1","nstrunv")
							default:
								print_string = Bible_print_string("王下","列王紀下", chap_string, sec_string,"nstrunv")
						}
				}
			case "1 Chr","First Chronicles","歷上","代上","歷代志上","歷代上","1Ch","1ch","Первая книга Паралипоменон","1 Sử-ký","歴上","歴代上","歴代志上","역대상":
				bible_short_name = "代上"
				switch chap_string {
					case "":
						print_string = "歷代志上"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("代上","歷代志上", chap_string, "1","nstrunv")
							default:
								print_string = Bible_print_string("代上","歷代志上", chap_string, sec_string,"nstrunv")
						}
				}
			case "2 Chr","Second Chronicles","代下","歷下","歷代下","歷代志下","2Ch","2ch","역대하","歴代志下","歴代下","歴下","2 Sử-ký","Вторая книга Паралипоменон":
				bible_short_name = "代下"
				switch chap_string {
					case "":
						print_string = "歷代志下"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("代下","歷代志下", chap_string, "1","nstrunv")
							default:
								print_string = Bible_print_string("代下","歷代志下", chap_string, sec_string,"nstrunv")
						}
				}
			case "Ezra","拉","以斯拉","以斯拉記","Ezr","ezr","Первая книга Ездры","Ê-xơ-ra","エズラ","エズラ記","에스라":
				bible_short_name = "拉"
				switch chap_string {
					case "":
						print_string = "以斯拉記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("拉","以斯拉記", chap_string, "1","nstrunv")
							default:
								print_string = Bible_print_string("拉","以斯拉記", chap_string, sec_string,"nstrunv")
						}
				}
			case "Neh","Nehemiah","尼","尼希米","尼希米記","Ne","ne","느헤미야","ネヘミヤ書","ネヘミヤ","Nê-hê-mi","Книга Неемии":
				bible_short_name = "尼"
				switch chap_string {
					case "":
						print_string = "尼希米記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("尼","尼希米記", chap_string, "1","nstrunv")
							default:
								print_string = Bible_print_string("尼","尼希米記", chap_string, sec_string,"nstrunv")
						}
				}
			case "Esth","Esther","斯","以斯帖","以斯帖記","Es","est","Есфирь","Ê-xơ-tê","エステル","エステル記","에스더":
				bible_short_name = "斯"
				switch chap_string {
					case "":
						print_string = "以斯帖記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("斯","以斯帖記", chap_string, "1","nstrunv")
							default:
								print_string = Bible_print_string("斯","以斯帖記", chap_string, sec_string,"nstrunv")
						}
				}
			case "Job","job","伯","約伯","約伯記","Книга Иова","Gióp","ヨブ","ヨブ記","욥기":
				bible_short_name = "伯"
				switch chap_string {
					case "":
						print_string = "約伯記"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("伯","約伯記", chap_string, "1","nstrunv")
							default:
								print_string = Bible_print_string("伯","約伯記", chap_string, sec_string,"nstrunv")
						}
				}
			case "Ps","Psalms","詩","詩篇","ps","시편","Thi-thiên","Псалтирь":
				bible_short_name = "詩"
				switch chap_string {
					case "":
						print_string = "詩篇"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("詩","詩篇", chap_string, "1","nstrunv")
							default:
								print_string = Bible_print_string("詩","詩篇", chap_string, sec_string,"nstrunv")
						}
				}
			case "Prov","Proverbs","箴","箴言","Pr","pr","Притчи Соломона","Châm-ngôn","잠언":
				bible_short_name = "箴"
				switch chap_string {
					case "":
						print_string = "箴言"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("箴","箴言", chap_string, "1","nstrunv")
							default:
								print_string = Bible_print_string("箴","箴言", chap_string, sec_string,"nstrunv")
						}
				}
			case "Eccl","Ecclesiastes","傳","傳道","傳道書","Ec","ec","Книга Екклезиаста","или Проповедника","Truyền-đạo","伝道の書","伝道","伝","伝道書","전도서":
				bible_short_name = "傳"
				switch chap_string {
					case "":
						print_string = "傳道書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("傳","傳道書", chap_string, "1","nstrunv")
							default:
								print_string = Bible_print_string("傳","傳道書", chap_string, sec_string,"nstrunv")
						}
				}
			case "Song","Song of Solomon","歌","雅歌","So","so","sng","Sng","Песнь песней Соломона","Nhã-ca","아가":
				bible_short_name = "歌"
				switch chap_string {
					case "":
						print_string = "雅歌"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("歌","雅歌", chap_string, "1","nstrunv")
							default:
								print_string = Bible_print_string("歌","雅歌", chap_string, sec_string,"nstrunv")
						}
				}
			case "Is","Isaiah","賽","以賽","以賽亞","以賽亞書","Isa","isa","Книга пророка Исаии","Ê-sai","イザヤ書","イザヤ","이사야":
				bible_short_name = "賽"
				switch chap_string {
					case "":
						print_string = "以賽亞書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("賽","以賽亞書", chap_string, "1","nstrunv")
							default:
								print_string = Bible_print_string("賽","以賽亞書", chap_string, sec_string,"nstrunv")
						}
				}
			case "Jer","Jeremiah","耶","耶利米","耶利米書","jer","예레미야","エレミヤ","エレミヤ書","Giê-rê-mi","Книга пророка Иеремии":
				bible_short_name = "耶"
				switch chap_string {
					case "":
						print_string = "耶利米書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("耶","耶利米書", chap_string, "1","nstrunv")
							default:
								print_string = Bible_print_string("耶","耶利米書", chap_string, sec_string,"nstrunv")
						}
				}
			case "Lam","Lamentations","哀","哀歌","耶利米哀歌","La","lam","예레미야애가","Ca-thương","Плач Иеремии":
				bible_short_name = "哀"
				switch chap_string {
					case "":
						print_string = "耶利米哀歌"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("哀","耶利米哀歌", chap_string, "1","nstrunv")
							default:
								print_string = Bible_print_string("哀","耶利米哀歌", chap_string, sec_string,"nstrunv")
						}
				}
			case "Ezek","Ezekiel","結","以西結","以西結書","Eze","eze","에스겔","エゼキエル書","エゼキエル","Ê-xê-chi-ên","Книга пророка Иезекииля":
				bible_short_name = "結"
				switch chap_string {
					case "":
						print_string = "以西結書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("結","以西結書", chap_string, "1","nstrunv")
							default:
								print_string = Bible_print_string("結","以西結書", chap_string, sec_string,"nstrunv")
						}
				}
			case "Dan","Daniel","但","但以理","但以理書","Da","da","Книга пророка Даниила","Đa-ni-ên","ダニエル書","ダニエル","다니엘":
				bible_short_name = "但"
				switch chap_string {
					case "":
						print_string = "但以理書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("但","但以理書", chap_string, "1","nstrunv")
							default:
								print_string = Bible_print_string("但","但以理書", chap_string, sec_string,"nstrunv")
						}
				}
			case "Hos","Hosea","何","何西","何西阿","何西阿書","Ho","ho","Книга пророка Осии","Ô-sê","ホセア書","ホセア","호세아":
				bible_short_name = "何"
				switch chap_string {
					case "":
						print_string = "何西阿書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("何","何西阿書", chap_string, "1","nstrunv")
							default:
								print_string = Bible_print_string("何","何西阿書", chap_string, sec_string,"nstrunv")
						}
				}
			case "Joel","珥","約珥","約珥書","Joe","joe","Книга пророка Иоиля","Giô-ên","ヨエル書","ヨエル","요엘":
				bible_short_name = "珥"
				switch chap_string {
					case "":
						print_string = "約珥書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("珥","約珥書", chap_string, "1","nstrunv")
							default:
								print_string = Bible_print_string("珥","約珥書", chap_string, sec_string,"nstrunv")
						}
				}
			case "Amos","摩","阿摩司書","Am","am","Книга пророка Амоса","A-mốt","アモス書","アモス","아모스":
				bible_short_name = "摩"
				switch chap_string {
					case "":
						print_string = "阿摩司書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("摩","阿摩司書", chap_string, "1","nstrunv")
							default:
								print_string = Bible_print_string("摩","阿摩司書", chap_string, sec_string,"nstrunv")
						}
				}
			case "Obad","Obadiah","俄","俄巴底亞","俄巴底亞書","Ob","ob","오바댜","オバデヤ書","オバデヤ","Áp-đia","Книга пророка Авдия":
				bible_short_name = "俄"
				switch chap_string {
					case "":
						print_string = "俄巴底亞書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("俄","俄巴底亞書", chap_string, "1","nstrunv")
							default:
								print_string = Bible_print_string("俄","俄巴底亞書", chap_string, sec_string,"nstrunv")
						}
				}
			case "Jon","Jonah","拿","約拿","約拿書","jon","요나","ヨナ書","ヨナ","Giô-na","Книга пророка Ионы":
				bible_short_name = "拿"
				switch chap_string {
					case "":
						print_string = "約拿書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("拿","約拿書", chap_string, "1","nstrunv")
							default:
								print_string = Bible_print_string("拿","約拿書", chap_string, sec_string,"nstrunv")
						}
				}
			case "Micah","彌","彌迦","彌迦書","Mic","mic","Книга пророка Михея","Mi-chê","ミカ書","ミカ","미가":
				bible_short_name = "彌"
				switch chap_string {
					case "":
						print_string = "彌迦書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("彌","彌迦書", chap_string, "1","nstrunv")
							default:
								print_string = Bible_print_string("彌","彌迦書", chap_string, sec_string,"nstrunv")
						}
				}
			case "Nah","Nahum","鴻","那鴻","那鴻書","Na","na","Книга пророка Наума","Na-hum","ナホム書","ナホム","나훔":
				bible_short_name = "鴻"
				switch chap_string {
					case "":
						print_string = "那鴻書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("鴻","那鴻書", chap_string, "1","nstrunv")
							default:
								print_string = Bible_print_string("鴻","那鴻書", chap_string, sec_string,"nstrunv")
						}
				}
			case "Habakkuk","哈","哈巴","哈巴谷","哈巴谷書","Hab","hab","Книга пророка Аввакума","Ha-ba-cúc","ハバクク書","ハバクク","ハバ","クク","ハバ書","하박국":
				bible_short_name = "哈"
				switch chap_string {
					case "":
						print_string = "哈巴谷書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("哈","哈巴谷書", chap_string, "1","nstrunv")
							default:
								print_string = Bible_print_string("哈","哈巴谷書", chap_string, sec_string,"nstrunv")
						}
				}
			case "Zeph","Zephaniah","番","西番雅","西番雅書","Zep","zep","스바냐","ゼパニヤ書","ゼパニヤ","Sô-phô-ni","Книга пророка Софонии":
				bible_short_name = "番"
				switch chap_string {
					case "":
						print_string = "西番雅書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("番","西番雅書", chap_string, "1","nstrunv")
							default:
								print_string = Bible_print_string("番","西番雅書", chap_string, sec_string,"nstrunv")
						}
				}
			case "Haggai","該","哈該","哈該書","Hag","hag","학개","ハガイ書","ハガイ","A-ghê","Книга пророка Аггея":
				bible_short_name = "該"
				switch chap_string {
					case "":
						print_string = "哈該書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("該","哈該書", chap_string, "1","nstrunv")
							default:
								print_string = Bible_print_string("該","哈該書", chap_string, sec_string,"nstrunv")
						}
				}
			case "Zech","Zechariah","亞","撒迦利亞","撒迦利亞書","Zec","zec","Книга пророка Захарии","Xa-cha-ri","스가랴","ゼカリヤ書","ゼカリヤ":
				bible_short_name = "亞"
				switch chap_string {
					case "":
						print_string = "撒迦利亞書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("亞","撒迦利亞書", chap_string, "1","nstrunv")
							default:
								print_string = Bible_print_string("亞","撒迦利亞書", chap_string, sec_string,"nstrunv")
						}
				}
			case "Malachi","瑪","瑪拉","瑪拉基","瑪拉基書","Mal","mal","말라기","マラキ書","マラキ","Ma-la-chi","Книга пророка Малахии":
				bible_short_name = "瑪"
				switch chap_string {
					case "":
						print_string = "瑪拉基書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("瑪","瑪拉基書", chap_string, "1","nstrunv")
							default:
								print_string = Bible_print_string("瑪","瑪拉基書", chap_string, sec_string,"nstrunv")
						}
				}
			case "Matt","Matthew","太","馬太","馬太福音","Mt","mt","마태복음","マタイによる福音書","マタイ","マタイによる","Ma-thi-ơ","От Матфея святое благовествование":
				bible_short_name = "太"
				switch chap_string {
					case "":
						print_string = "馬太福音"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("太","馬太福音", chap_string, "1","nstrunv")
							default:
								print_string = Bible_print_string("太","馬太福音", chap_string, sec_string,"nstrunv")
						}
				}
			case "Mark","可","馬可","馬可福音","Mr","mr","マルコによる福音書","マルコ","マルコによる","마가복음","Mác","От Марка святое благовествование":
				bible_short_name = "可"
				switch chap_string {
					case "":
						print_string = "馬可福音"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("可","馬可福音", chap_string, "1","nstrunv")
							default:
								print_string = Bible_print_string("可","馬可福音", chap_string, sec_string,"nstrunv")
						}
				}
			case "Luke","路","路加","路加福音","Lu","lu","От Луки святое благовествование","Lu-ca","ルカによる福音書","ルカ","ルカによる","누가복음":
				bible_short_name = "路"
				switch chap_string {
					case "":
						print_string = "路加福音"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("路","路加福音", chap_string, "1","nstrunv")
							default:
								print_string = Bible_print_string("路","路加福音", chap_string, sec_string,"nstrunv")
						}
				}
			case "John","約","約翰","約翰福音","Joh","joh","От Иоанна святое благовествование","Giăng","ヨハネによる福音書","ヨハネ","ヨハネによる","요한복음":
				bible_short_name = "約"
				switch chap_string {
					case "":
						print_string = "約翰福音"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("約","約翰福音", chap_string, "1","nstrunv")
							default:
								print_string = Bible_print_string("約","約翰福音", chap_string, sec_string,"nstrunv")
						}
				}
			case "Acts","徒","使徒","使徒行傳","Ac","ac","Деяния святых апостолов","Công-vụ Các Sứ-đồ","使徒行伝","사도행전":
				bible_short_name = "徒"
				switch chap_string {
					case "":
						print_string = "使徒行傳"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("徒","使徒行傳", chap_string, "1","nstrunv")
							default:
								print_string = Bible_print_string("徒","使徒行傳", chap_string, sec_string,"nstrunv")
						}
				}
			case "Rom","Romans","羅","羅馬","羅馬書","Ro","ro","Послание к Римлянам","Rô-ma","ローマ","ローマ人への手紙","로마서":
				bible_short_name = "羅"
				switch chap_string {
					case "":
						print_string = "羅馬書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("羅","羅馬書", chap_string, "1","nstrunv")
							default:
								print_string = Bible_print_string("羅","羅馬書", chap_string, sec_string,"nstrunv")
						}
				}
			case "1 Cor","First Corinthians","林前","哥林多前","哥林多前書","1Co","1co","Первое послание к Коринфянам","1 Cô-rinh-tô","コリント人への第一の手紙","コリント一","コリント人への第一","고린도전서":
				bible_short_name = "林前"
				switch chap_string {
					case "":
						print_string = "哥林多前書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("林前","哥林多前書", chap_string, "1","nstrunv")
							default:
								print_string = Bible_print_string("林前","哥林多前書", chap_string, sec_string,"nstrunv")
						}
				}
			case "2 Cor","Second Corinthians","林後","哥林多後","哥林多後書","2Co","2co","Второе послание к Коринфянам","2 Cô-rinh-tô","コリント人への第二の手紙","コリント二","コリント人への第二の","고린도후서":
				bible_short_name = "林後"
				switch chap_string {
					case "":
						print_string = "哥林多後書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("林後","哥林多後書", chap_string, "1","nstrunv")
							default:
								print_string = Bible_print_string("林後","哥林多後書", chap_string, sec_string,"nstrunv")
						}
				}
			case "Gal","Galatians","加","加拉太","加拉太書","Ga","ga","Послание к Галатам","Ga-la-ti","ガラテヤ","ガラテヤ人への手紙","갈라디아서":
				bible_short_name = "加"
				switch chap_string {
					case "":
						print_string = "加拉太書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("加","加拉太書", chap_string, "1","nstrunv")
							default:
								print_string = Bible_print_string("加","加拉太書", chap_string, sec_string,"nstrunv")
						}
				}
			case "Ephesians","弗","以弗所","以弗所書","Eph","eph","Послание к Ефесянам","Ê-phê-sô","エペソ人への手紙","エペソ","エペソ人","エペソ人の手紙","에베소서":
				bible_short_name = "弗"
				switch chap_string {
					case "":
						print_string = "以弗所書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("弗","以弗所書", chap_string, "1","nstrunv")
							default:
								print_string = Bible_print_string("弗","以弗所書", chap_string, sec_string,"nstrunv")
						}
				}
			case "Phil","Philippians","腓","腓立","腓立比","腓立比書","Php","php","빌립보서","ピリピ","ピリピ人.ピリピ人への手紙","Послание к Филиппийцам","Phi-líp":
				bible_short_name = "腓"
				switch chap_string {
					case "":
						print_string = "腓立比書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("腓","腓立比書", chap_string, "1","nstrunv")
							default:
								print_string = Bible_print_string("腓","腓立比書", chap_string, sec_string,"nstrunv")
						}
				}
			case "Col","col","Colossians","西","歌羅西","歌羅","歌羅西書","Послание к Колоссянам","Cô-lô-se","コロサイ人への手紙","コロサイ","コロ","골로새서":
				bible_short_name = "西"
				switch chap_string {
					case "":
						print_string = "歌羅西書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("西","歌羅西書", chap_string, "1","nstrunv")
							default:
								print_string = Bible_print_string("西","歌羅西書", chap_string, sec_string,"nstrunv")
						}
				}
			case "1 Thess","First Thessalonians","帖前","帖撒羅尼迦前","帖撒羅尼迦前書","1Th","1th","데살로니가전서","テサロニケ人への第一の手紙","テサ一","テサロニケ一","1 Tê-sa-lô-ni-ca","Первое послание к Фессалоникийцам (Солунянам)":
				bible_short_name = "帖前"
				switch chap_string {
					case "":
						print_string = "帖撒羅尼迦前書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("帖前","帖撒羅尼迦前書", chap_string, "1","nstrunv")
							default:
								print_string = Bible_print_string("帖前","帖撒羅尼迦前書", chap_string, sec_string,"nstrunv")
						}
				}
			case "2 Thess","Second Thessalonians","帖後","帖撒羅尼迦後","帖撒羅尼迦後書","2Th","2th","데살로니가후서","テサロニケ人への第二の手紙","テサ二","テサロニケ二","2 Tê-sa-lô-ni-ca","Второе послание к Фессалоникийцам (Солунянам)":
				bible_short_name = "帖後"
				switch chap_string {
					case "":
						print_string = "帖撒羅尼迦後書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("帖後","帖撒羅尼迦後書", chap_string, "1","nstrunv")
							default:
								print_string = Bible_print_string("帖後","帖撒羅尼迦後書", chap_string, sec_string,"nstrunv")
						}
				}
			case "1 Tim","First Timothy","提前","提摩太前","提摩太前書","1Ti","1ti","Первое послание к Тимофею","1 Ti-mô-thê","テモテヘの第一の手紙","テモテ一","디모데전서":
				bible_short_name = "提前"
				switch chap_string {
					case "":
						print_string = "提摩太前書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("提前","提摩太前書", chap_string, "1","nstrunv")
							default:
								print_string = Bible_print_string("提前","提摩太前書", chap_string, sec_string,"nstrunv")
						}
				}
			case "2 Tim","Second Timothy","提後","提摩太後","提摩太後書","2Ti","2ti","Второе послание к Тимофею","2 Ti-mô-thê","テモテヘの第二の手紙","テモテ二","디모데후서":
				bible_short_name = "提後"
				switch chap_string {
					case "":
						print_string = "提摩太後書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("提後","提摩太後書", chap_string, "1","nstrunv")
							default:
								print_string = Bible_print_string("提後","提摩太後書", chap_string, sec_string,"nstrunv")
						}
				}
			case "Titus","多","提多","提多書","Tit","tit","Послание к Титу","Tít","テトスヘの手紙","テトス","디도서":
				bible_short_name = "多"
				switch chap_string {
					case "":
						print_string = "提多書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("多","提多書", chap_string, "1","nstrunv")
							default:
								print_string = Bible_print_string("多","提多書", chap_string, sec_string,"nstrunv")
						}
				}
			case "Philem","Philemon","門","腓利","腓利門","腓利門書","Phm","phm","Послание к Филимону","Phi-lê-môn","ピレモンヘの手紙","ピレモン","빌레몬서":
				bible_short_name = "門"
				switch chap_string {
					case "":
						print_string = "腓利門書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("門","腓利門書", chap_string, "1","nstrunv")
							default:
								print_string = Bible_print_string("門","腓利門書", chap_string, sec_string,"nstrunv")
						}
				}
			case "Heb","Hebrews","來","希伯來","希伯來書","heb","Послание к Евреям","Hê-bơ-rơ","ヘブル人への手紙","ヘブル","히브리서":
				bible_short_name = "來"
				switch chap_string {
					case "":
						print_string = "希伯來書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("來","希伯來書", chap_string, "1","nstrunv")
							default:
								print_string = Bible_print_string("來","希伯來書", chap_string, sec_string,"nstrunv")
						}
				}
			case "James","雅","雅各","雅各書","Jas","jas","Послание Иакова","Gia-cơ","ヤコブの手紙","ヤコブ","야고보서":
				bible_short_name = "雅"
				switch chap_string {
					case "":
						print_string = "雅各書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("雅","雅各書", chap_string, "1","nstrunv")
							default:
								print_string = Bible_print_string("雅","雅各書", chap_string, sec_string,"nstrunv")
						}
				}
			case "1 Pet","First Peter","彼前","彼得前","彼得前書","1Pe","1pe","Первое послание Петра","1 Phi-e-rơ","ペテロの第一の手紙","ペテロ一","베드로전서":
				bible_short_name = "彼前"
				switch chap_string {
					case "":
						print_string = "彼得前書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("彼前","彼得前書", chap_string, "1","nstrunv")
							default:
								print_string = Bible_print_string("彼前","彼得前書", chap_string, sec_string,"nstrunv")
						}
				}
			case "2 Pet","Second Peter","彼後","彼得後","彼得後書","2Pe","2pe","Второе послание Петра","2 Phi-e-rơ","ペテロの第二の手紙","ペテロ","베드로후서":
				bible_short_name = "彼後"
				switch chap_string {
					case "":
						print_string = "彼得後書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("彼後","彼得後書", chap_string, "1","nstrunv")
							default:
								print_string = Bible_print_string("彼後","彼得後書", chap_string, sec_string,"nstrunv")
						}
				}
			case "1 John","First John","約一","約翰一書","約翰1","約翰1書","1Jo","1jo","Первое послание Иоанна","1 Giăng","ヨハネの第一の手紙","ヨハネ一","요한일서":
				bible_short_name = "約一"
				switch chap_string {
					case "":
						print_string = "約翰一書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("約一","約翰一書", chap_string, "1","nstrunv")
							default:
								print_string = Bible_print_string("約一","約翰一書", chap_string, sec_string,"nstrunv")
						}
				}
			case "2 John","second John","約二","約翰二書","約翰2","約翰2書","2Jo","Второе послание Иоанна","2 Giăng","ヨハネの第二の手紙","ヨハネ二","요한2서":
				bible_short_name = "約二"
				switch chap_string {
					case "":
						print_string = "約翰二書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("約二","約翰二書", chap_string, "1","nstrunv")
							default:
								print_string = Bible_print_string("約二","約翰二書", chap_string, sec_string,"nstrunv")
						}
				}
			case "3 John","Third John","約三","約翰三書","約翰3","約翰3書","3Jo","3jo","Третье послание Иоанна","3 Giăng","ヨハネの第三の手紙","ヨハネ三","요한3서":
				bible_short_name = "約三"
				switch chap_string {
					case "":
						print_string = "約翰三書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("約三","約翰三書", chap_string, "1","nstrunv")
							default:
								print_string = Bible_print_string("約三","約翰三書", chap_string, sec_string,"nstrunv")
						}
				}
			case "Jude","猶","猶大","猶大書","jude","Послание Иуды","Giu-đe","ユダの手紙","ユダ","유다서":
				bible_short_name = "猶"
				switch chap_string {
					case "":
						print_string = "猶大書"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("猶","猶大書", chap_string, "1","nstrunv")
							default:
								print_string = Bible_print_string("猶","猶大書", chap_string, sec_string,"nstrunv")
						}
				}
			case "Rev","Revelation","啟","啟示","啟示錄","Re","re","ｒｅ","Ｒｅ","rev","Откровение ап. Иоанна Богослова (Апокалипсис)","Khải-huyền","ヨハネの黙示録","黙示録","요한계시록":
				bible_short_name = "啟"
				switch chap_string {
					case "":
						print_string = "啟示錄"  //不知章節的時候 //用來等觸發 UI 及特別說明文字
					default:
						switch sec_string {
							case "":	//不知節的時候，知章
								print_string = Bible_print_string("啟","啟示錄", chap_string, "1","nstrunv")
							default:
								print_string = Bible_print_string("啟","啟示錄", chap_string, sec_string,"nstrunv")
						}
				}
			case "聖經","bible","Bible":
				print_string = "聖經"
			default:
				//print_string = "你是要找 " +  reg.ReplaceAllString(text, "$3") + " 對嗎？\n對不起，我還沒學呢...\n"
				print_string = "聖經"
		}
	default:
		if reply_mode!="" {
			print_string = "HI～\n如有其他建議或想討論，請對我輸入「開發者」進行聯絡。"
		} else {
            print_string = "" //安靜模式
		}
	}
	log.Print("Return 前的 print_string = " + print_string)
	return print_string, chap_string, sec_string, bible_short_name
}


