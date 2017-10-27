package main

import (
	"fmt"
	"log"

	// "net/http"
	// "os"
	// "regexp"

	// "strconv"
	"strings"
	
	// "github.com/line/line-bot-sdk-go/linebot"

	// "bytes"

	// "io/ioutil"

	// "image/jpeg"

    // "crypto/md5"
    // "encoding/hex"

    "encoding/json"
    // "github.com/bitly/go-simplejson"
)

type Bible_record struct {
	Engs 		string	`json:"engs"`
	Chineses	string	`json:"chineses"`	
	Chap		float64	`json:"chap"`
	Sec			float64	`json:"sec"`
	Bible_text	string	`json:"bible_text"`		
}

type Bibles_fhl_net_json struct {
	Record			[]Bible_record	`json:"record"`
	Prev			Bible_record	`json:"prev"`
	Next			Bible_record	`json:"next"`
	Version			string			`json:"version"`
	Status			string 			`json:"status"`
	Record_count	float64 		`json:"record_count"`
	Proc			float64			`json:"proc"`
	
}

func GetJson_bible(json_text string) (string){
	var find_bible Bibles_fhl_net_json
	json.Unmarshal([]byte(json_text), &find_bible)
	//return find_bible.Record[0].Bible_text

	out_string := ""	//find_bible.Record[0].Bible_text //取出一項的時候這樣做

	if int(find_bible.Record_count) > 1 {
		//https://trello.com/c/IJ4gwpUU/1105-json-the-go-playground
		//範圍經節的關鍵
		for i := 0; i < int(find_bible.Record_count); i++ {
			//https://trello.com/c/44jd5Afa/1108-the-go-playground
			out_string += strings.Replace(fmt.Sprintf("%f", find_bible.Record[i].Sec),`.000000`, "", -1) + ". " + find_bible.Record[i].Bible_text + "\n"
		}	
	}else{
		if (find_bible.Record_count==0){
			out_string = ""
		}else{
			out_string = find_bible.Record[0].Bible_text
		}
	}

	out_string = strings.Replace(out_string,"<br/>", "\n", -1)
	out_string = strings.Replace(out_string,`<div class="quote">`, "", -1)
	out_string = strings.Replace(out_string,`</div>`, "", -1)
	out_string = strings.Replace(out_string,`<h2>`, "", -1)
	out_string = strings.Replace(out_string,`</h2>`, "", -1)
	out_string = strings.Replace(out_string,`<h3>`, "", -1)
	out_string = strings.Replace(out_string,`</h3>`, "", -1)
	out_string = strings.Replace(out_string,`<h4>`, "", -1)
	out_string = strings.Replace(out_string,`</h4>`, "", -1)
	out_string = strings.Replace(out_string,`<i>`, "", -1)
	out_string = strings.Replace(out_string,`</i>`, "", -1)
	return out_string


	

	// return "123"

    // js, err := NewJson([]byte(json_text))
    // js, err := NewJson([]byte(`{
    //     "test": {
    //         "array": [1, "2", 3],
    //         "int": 10,
    //         "float": 5.150,
    //         "bignum": 9223372036854775807,
    //         "string": "simplejson",
    //         "bool": true
    //     }
    // }`))
    // arr, _ := js.Get("test").Get("array").Array()
    // i, _ := js.Get("test").Get("int").Int()
    // ms := js.Get("test").Get("string").MustString()
    // return js.Get("record").Get("bible_text").MustString()

	// `{
	// 	"status":"success",
	// 	"record_count":1,
	// 	"v_name":null,
	// 	"version":"nstrunv",
	// 	"proc":0,
	// 	"record":[
	// 		{
	// 			"engs":"Gen",
	// 			"chineses":"\u5275",
	// 			"chap":1,
	// 			"sec":5,
	// 			"bible_text":"\u795e\u7a31\u5149\u70ba\u300c\u665d\u300d\uff0c\u7a31\u6697\u70ba\u300c\u591c\u300d\u3002\u6709\u665a\u4e0a\uff0c\u6709\u65e9\u6668\uff0c\u9019\u662f\u982d\u4e00\u65e5\u3002"
	// 		}
	// 	],
	// 	"prev":{
	// 		"chineses":"\u5275",
	// 		"engs":"Gen",
	// 		"chap":1,
	// 		"sec":4
	// 	} ,
	// 	"next":{
	// 		"chineses":"\u5275",
	// 		"engs":"Gen",
	// 		"chap":1,
	// 		"sec":6
	// 	}
	// }`
}

func Bible_print_all_string(short_name, output_name, chap_string, sec_string string) (string) {
	//多語言，每個語言用一本
	//nstrunv,jp,kjv(英文),korean,russian(俄文),vietnamese(越南)

	bible_json_string := HttpGET_("http://bible.fhl.net/json/qb.php?chineses=" + short_name + "&chap=" + chap_string + "&sec=" + sec_string + "&version=nstrunv")
	bible_text_string := ""
	//是否有 GET 到回應
	if bible_json_string!="" {
		//是否有查到紀錄
		if(GetJson_bible(bible_json_string)!=""){
			bible_text_string = "【和合本】\n" + GetJson_bible(bible_json_string) + "\n" //把 JSON 丟過去分離出要的內容
			//}

				// log.Print("GET = ")
				// log.Print("bible_json_string = " + bible_json_string)
				// log.Print("GetJson_bible = " + bible_text_string)

			//多版本、多國語言串聯
			//https://www.bible.com/zh-TW/audio-bible-app-versions/


			// bible_json_string = HttpGET_("http://bible.fhl.net/json/qb.php?chineses=" + short_name + "&chap=" + chap_string + "&sec=" + sec_string + "&version=rcuv")
			// if (bible_json_string!="") && (GetJson_bible(bible_json_string)!="") {
			// 	bible_text_string += "\n【和合本修訂版 2010】\n" + GetJson_bible(bible_json_string) + "\n" //把 JSON 丟過去分離出要的內容++
			// }

			// bible_json_string = HttpGET_("http://bible.fhl.net/json/qb.php?chineses=" + short_name + "&chap=" + chap_string + "&sec=" + sec_string + "&version=ncv")
			// if (bible_json_string!="") && (GetJson_bible(bible_json_string)!="") {
			// 	bible_text_string += "\n【中文新譯本 2010】\n" + GetJson_bible(bible_json_string) + "\n" //把 JSON 丟過去分離出要的內容
			// }

			// bible_json_string = HttpGET_("http://bible.fhl.net/json/qb.php?chineses=" + short_name + "&chap=" + chap_string + "&sec=" + sec_string + "&version=tcv")
			// if (bible_json_string!="") && (GetJson_bible(bible_json_string)!="") {
			// 	bible_text_string += "\n【現代中文譯本修訂版 1997】\n" + GetJson_bible(bible_json_string) + "\n" //把 JSON 丟過去分離出要的內容
			// }

			// bible_json_string = HttpGET_("http://bible.fhl.net/json/qb.php?chineses=" + short_name + "&chap=" + chap_string + "&sec=" + sec_string + "&version=wlunv")
			// if (bible_json_string!="") && (GetJson_bible(bible_json_string)!="") {
			// 	bible_text_string += "\n【文言文（深文理和合本）】\n" + GetJson_bible(bible_json_string) + "\n" //把 JSON 丟過去分離出要的內容
			// }

			bible_json_string = HttpGET_("http://bible.fhl.net/json/qb.php?chineses=" + short_name + "&chap=" + chap_string + "&sec=" + sec_string + "&version=sgebklhl")
			if (bible_json_string!="") && (GetJson_bible(bible_json_string)!="") {
				bible_text_string += "\n【台語（全民台語聖經漢羅）】\n" + GetJson_bible(bible_json_string) + "\n" //把 JSON 丟過去分離出要的內容
			}

			// bible_json_string = HttpGET_("http://bible.fhl.net/json/qb.php?chineses=" + short_name + "&chap=" + chap_string + "&sec=" + sec_string + "&version=sgebklcl")
			// if (bible_json_string!="") && (GetJson_bible(bible_json_string)!="") {
			// 	bible_text_string += "\n【台語（全民台語聖經全羅）】\n" + GetJson_bible(bible_json_string) + "\n" //把 JSON 丟過去分離出要的內容
			// }

			// bible_json_string = HttpGET_("http://bible.fhl.net/json/qb.php?chineses=" + short_name + "&chap=" + chap_string + "&sec=" + sec_string + "&version=bklhl")
			// if (bible_json_string!="") && (GetJson_bible(bible_json_string)!="") {
			// 	bible_text_string += "\n【台語（巴克禮漢羅）】\n" + GetJson_bible(bible_json_string) + "\n" //把 JSON 丟過去分離出要的內容
			// }

			// bible_json_string = HttpGET_("http://bible.fhl.net/json/qb.php?chineses=" + short_name + "&chap=" + chap_string + "&sec=" + sec_string + "&version=bklcl")
			// if (bible_json_string!="") && (GetJson_bible(bible_json_string)!="") {
			// 	bible_text_string += "\n【台語（巴克禮全羅）】\n" + GetJson_bible(bible_json_string) + "\n" //把 JSON 丟過去分離出要的內容
			// }

			// bible_json_string = HttpGET_("http://bible.fhl.net/json/qb.php?chineses=" + short_name + "&chap=" + chap_string + "&sec=" + sec_string + "&version=prebklhl")
			// if (bible_json_string!="") && (GetJson_bible(bible_json_string)!="") {
			// 	bible_text_string += "\n【台語（馬雅各漢羅）】\n" + GetJson_bible(bible_json_string) + "\n" //把 JSON 丟過去分離出要的內容
			// }

			// bible_json_string = HttpGET_("http://bible.fhl.net/json/qb.php?chineses=" + short_name + "&chap=" + chap_string + "&sec=" + sec_string + "&version=prebklcl")
			// if (bible_json_string!="") && (GetJson_bible(bible_json_string)!="") {
			// 	bible_text_string += "\n【台語（馬雅各全羅）】\n" + GetJson_bible(bible_json_string) + "\n" //把 JSON 丟過去分離出要的內容
			// }

			bible_json_string = HttpGET_("http://bible.fhl.net/json/qb.php?chineses=" + short_name + "&chap=" + chap_string + "&sec=" + sec_string + "&version=hakka")
			if (bible_json_string!="") && (GetJson_bible(bible_json_string)!="") {
				bible_text_string += "\n【客家話（新約）】\n" + GetJson_bible(bible_json_string) + "\n" //把 JSON 丟過去分離出要的內容
			}

			//https://www.bible.com/zh-TW/audio-bible-app-versions/1-kjv-king-james-version
			bible_json_string = HttpGET_("http://bible.fhl.net/json/qb.php?chineses=" + short_name + "&chap=" + chap_string + "&sec=" + sec_string + "&version=kjv")
			if (bible_json_string!="") && (GetJson_bible(bible_json_string)!="") {
				bible_text_string += "\n【英文 KJV（King James Version 英王欽定版本）】\n" + GetJson_bible(bible_json_string) + "\n" //把 JSON 丟過去分離出要的內容
			}

			// bible_json_string = HttpGET_("http://bible.fhl.net/json/qb.php?chineses=" + short_name + "&chap=" + chap_string + "&sec=" + sec_string + "&version=bbe")
			// if (bible_json_string!="") && (GetJson_bible(bible_json_string)!="") {
			// 	bible_text_string += "\n【英文 BBE（簡易英文譯本）】\n" + GetJson_bible(bible_json_string) + "\n" //把 JSON 丟過去分離出要的內容
			// }

			// //https://www.bible.com/zh-TW/audio-bible-app-versions/206-web-world-english-bible
			// bible_json_string = HttpGET_("http://bible.fhl.net/json/qb.php?chineses=" + short_name + "&chap=" + chap_string + "&sec=" + sec_string + "&version=web")
			// if (bible_json_string!="") && (GetJson_bible(bible_json_string)!="") {
			// 	bible_text_string += "\n【英文 WEB（環球英文聖經）】\n" + GetJson_bible(bible_json_string) + "\n" //把 JSON 丟過去分離出要的內容
			// }

			// //https://www.bible.com/zh-TW/audio-bible-app-versions/12-asv-american-standard-version
			// //http://www.freeasv.org/?page=listen
			// bible_json_string = HttpGET_("http://bible.fhl.net/json/qb.php?chineses=" + short_name + "&chap=" + chap_string + "&sec=" + sec_string + "&version=asv")
			// if (bible_json_string!="") && (GetJson_bible(bible_json_string)!="") {
			// 	bible_text_string += "\n【英文 ASV（美國標準譯本）】\n" + GetJson_bible(bible_json_string) + "\n" //把 JSON 丟過去分離出要的內容
			// }

			// bible_json_string = HttpGET_("http://bible.fhl.net/json/qb.php?chineses=" + short_name + "&chap=" + chap_string + "&sec=" + sec_string + "&version=darby")
			// if (bible_json_string!="") && (GetJson_bible(bible_json_string)!="") {
			// 	bible_text_string += "\n【英文 Darby 1890】\n" + GetJson_bible(bible_json_string) + "\n" //把 JSON 丟過去分離出要的內容
			// }

			// //https://bible.fhl.net/annouce/annouce426.html
			// bible_json_string = HttpGET_("http://bible.fhl.net/json/qb.php?chineses=" + short_name + "&chap=" + chap_string + "&sec=" + sec_string + "&version=erv")
			// if (bible_json_string!="") && (GetJson_bible(bible_json_string)!="") {
			// 	bible_text_string += "\n【英文 ERV（English Revised Version 英國修訂譯本）】\n" + GetJson_bible(bible_json_string) + "\n" //把 JSON 丟過去分離出要的內容
			// }

			//http://bible.salterrae.net/
			//http://bible.salterrae.net/kougo/html/
			bible_json_string = HttpGET_("http://bible.fhl.net/json/qb.php?chineses=" + short_name + "&chap=" + chap_string + "&sec=" + sec_string + "&version=jp")
			if (bible_json_string!="") && (GetJson_bible(bible_json_string)!="") {
				bible_text_string += "\n【日文口語訳】\n" + GetJson_bible(bible_json_string) + "\n" //把 JSON 丟過去分離出要的內容
			}

			bible_json_string = HttpGET_("http://bible.fhl.net/json/qb.php?chineses=" + short_name + "&chap=" + chap_string + "&sec=" + sec_string + "&version=korean")
			if (bible_json_string!="") && (GetJson_bible(bible_json_string)!="") {
				bible_text_string += "\n【韓文 KRV】\n" + GetJson_bible(bible_json_string) + "\n" //把 JSON 丟過去分離出要的內容
			}

			bible_json_string = HttpGET_("http://bible.fhl.net/json/qb.php?chineses=" + short_name + "&chap=" + chap_string + "&sec=" + sec_string + "&version=vietnamese")
			if (bible_json_string!="") && (GetJson_bible(bible_json_string)!="") {
				bible_text_string += "\n【越南 VI1934】\n" + GetJson_bible(bible_json_string) + "\n" //把 JSON 丟過去分離出要的內容
			}

			bible_json_string = HttpGET_("http://bible.fhl.net/json/qb.php?chineses=" + short_name + "&chap=" + chap_string + "&sec=" + sec_string + "&version=russian")
			if (bible_json_string!="") && (GetJson_bible(bible_json_string)!="") {
				bible_text_string += "\n【俄文 SYNOD】\n" + GetJson_bible(bible_json_string) + "\n" //把 JSON 丟過去分離出要的內容
			}

			// //http://c.ibibles.net/c02.php?1
			// bible_json_string = HttpGET_("http://bible.fhl.net/json/qb.php?chineses=" + short_name + "&chap=" + chap_string + "&sec=" + sec_string + "&version=lxx")
			// if (bible_json_string!="") && (GetJson_bible(bible_json_string)!="") {
			// 	bible_text_string += "\n【舊約 希臘古譯 Εβδομήκοντα εκδοχή 七十士本 LXX】\n" + GetJson_bible(bible_json_string) + "\n" //把 JSON 丟過去分離出要的內容
			// }

			// bible_json_string = HttpGET_("http://bible.fhl.net/json/qb.php?chineses=" + short_name + "&chap=" + chap_string + "&sec=" + sec_string + "&version=bhs")
			// if (bible_json_string!="") && (GetJson_bible(bible_json_string)!="") {
			// 	bible_text_string += "\n【舊約 Hebraica Stuttgartensia הבראיקה 馬索拉原文 BHS】\n" + GetJson_bible(bible_json_string) + "\n" //把 JSON 丟過去分離出要的內容
			// }

		}else{
			return "查詢章節超過聖經範圍，有可能指定查詢的節超過範圍。"		
		}
	}

	if bible_text_string != ""{
		return "[" + output_name + " " + chap_string + " : " +  sec_string + "]\n" + bible_text_string
	}else{
		return ""
	}
}

func Bible_print_all_var_string(short_name, output_name, chap_string, sec_string string) (string) {

	//nstrunv,jp,kjv(英文),korean,russian(俄文),vietnamese(越南)

	bible_json_string := HttpGET_("http://bible.fhl.net/json/qb.php?chineses=" + short_name + "&chap=" + chap_string + "&sec=" + sec_string + "&version=nstrunv")
	bible_text_string := ""
	//是否有 GET 到回應
	if bible_json_string!="" {
		//是否有查到紀錄
		if(GetJson_bible(bible_json_string)!=""){
			bible_text_string = "【和合本】\n" + GetJson_bible(bible_json_string) + "\n" //把 JSON 丟過去分離出要的內容
			//}

				// log.Print("GET = ")
				// log.Print("bible_json_string = " + bible_json_string)
				// log.Print("GetJson_bible = " + bible_text_string)

			//多版本、多國語言串聯
			//https://www.bible.com/zh-TW/audio-bible-app-versions/ http://sharinglearner.blogspot.tw/2008/11/blog-post.html


			bible_json_string = HttpGET_("http://bible.fhl.net/json/qb.php?chineses=" + short_name + "&chap=" + chap_string + "&sec=" + sec_string + "&version=rcuv")
			if (bible_json_string!="") && (GetJson_bible(bible_json_string)!="") {
				bible_text_string += "\n【和合本修訂版 2010】\n" + GetJson_bible(bible_json_string) + "\n" //把 JSON 丟過去分離出要的內容+++
			}

			bible_json_string = HttpGET_("http://bible.fhl.net/json/qb.php?chineses=" + short_name + "&chap=" + chap_string + "&sec=" + sec_string + "&version=ncv")
			if (bible_json_string!="") && (GetJson_bible(bible_json_string)!="") {
				bible_text_string += "\n【中文新譯本 2010】\n" + GetJson_bible(bible_json_string) + "\n" //把 JSON 丟過去分離出要的內容
			}

			bible_json_string = HttpGET_("http://bible.fhl.net/json/qb.php?chineses=" + short_name + "&chap=" + chap_string + "&sec=" + sec_string + "&version=tcv")
			if (bible_json_string!="") && (GetJson_bible(bible_json_string)!="") {
				bible_text_string += "\n【現代中文譯本修訂版 1997】\n" + GetJson_bible(bible_json_string) + "\n" //把 JSON 丟過去分離出要的內容
			}

			bible_json_string = HttpGET_("http://bible.fhl.net/json/qb.php?chineses=" + short_name + "&chap=" + chap_string + "&sec=" + sec_string + "&version=wlunv")
			if (bible_json_string!="") && (GetJson_bible(bible_json_string)!="") {
				bible_text_string += "\n【文言文（深文理和合本）】\n" + GetJson_bible(bible_json_string) + "\n" //把 JSON 丟過去分離出要的內容
			}

			bible_json_string = HttpGET_("http://bible.fhl.net/json/qb.php?chineses=" + short_name + "&chap=" + chap_string + "&sec=" + sec_string + "&version=sgebklhl")
			if (bible_json_string!="") && (GetJson_bible(bible_json_string)!="") {
				bible_text_string += "\n【台語（全民台語聖經漢羅）】\n" + GetJson_bible(bible_json_string) + "\n" //把 JSON 丟過去分離出要的內容
			}

			bible_json_string = HttpGET_("http://bible.fhl.net/json/qb.php?chineses=" + short_name + "&chap=" + chap_string + "&sec=" + sec_string + "&version=sgebklcl")
			if (bible_json_string!="") && (GetJson_bible(bible_json_string)!="") {
				bible_text_string += "\n【台語（全民台語聖經全羅）】\n" + GetJson_bible(bible_json_string) + "\n" //把 JSON 丟過去分離出要的內容
			}

			bible_json_string = HttpGET_("http://bible.fhl.net/json/qb.php?chineses=" + short_name + "&chap=" + chap_string + "&sec=" + sec_string + "&version=bklhl")
			if (bible_json_string!="") && (GetJson_bible(bible_json_string)!="") {
				bible_text_string += "\n【台語（巴克禮漢羅）】\n" + GetJson_bible(bible_json_string) + "\n" //把 JSON 丟過去分離出要的內容
			}

			bible_json_string = HttpGET_("http://bible.fhl.net/json/qb.php?chineses=" + short_name + "&chap=" + chap_string + "&sec=" + sec_string + "&version=bklcl")
			if (bible_json_string!="") && (GetJson_bible(bible_json_string)!="") {
				bible_text_string += "\n【台語（巴克禮全羅）】\n" + GetJson_bible(bible_json_string) + "\n" //把 JSON 丟過去分離出要的內容
			}

			bible_json_string = HttpGET_("http://bible.fhl.net/json/qb.php?chineses=" + short_name + "&chap=" + chap_string + "&sec=" + sec_string + "&version=prebklhl")
			if (bible_json_string!="") && (GetJson_bible(bible_json_string)!="") {
				bible_text_string += "\n【台語（馬雅各漢羅）】\n" + GetJson_bible(bible_json_string) + "\n" //把 JSON 丟過去分離出要的內容
			}

			bible_json_string = HttpGET_("http://bible.fhl.net/json/qb.php?chineses=" + short_name + "&chap=" + chap_string + "&sec=" + sec_string + "&version=prebklcl")
			if (bible_json_string!="") && (GetJson_bible(bible_json_string)!="") {
				bible_text_string += "\n【台語（馬雅各全羅）】\n" + GetJson_bible(bible_json_string) + "\n" //把 JSON 丟過去分離出要的內容
			}

			bible_json_string = HttpGET_("http://bible.fhl.net/json/qb.php?chineses=" + short_name + "&chap=" + chap_string + "&sec=" + sec_string + "&version=hakka")
			if (bible_json_string!="") && (GetJson_bible(bible_json_string)!="") {
				bible_text_string += "\n【客家話（新約）】\n" + GetJson_bible(bible_json_string) + "\n" //把 JSON 丟過去分離出要的內容
			}

			bible_json_string = HttpGET_("http://bible.fhl.net/json/qb.php?chineses=" + short_name + "&chap=" + chap_string + "&sec=" + sec_string + "&version=rukai")
			if (bible_json_string!="") && (GetJson_bible(bible_json_string)!="") {
				bible_text_string += "\n【魯凱語聖經】\n" + GetJson_bible(bible_json_string) + "\n" //把 JSON 丟過去分離出要的內容
			}

			bible_json_string = HttpGET_("http://bible.fhl.net/json/qb.php?chineses=" + short_name + "&chap=" + chap_string + "&sec=" + sec_string + "&version=tsou")
			if (bible_json_string!="") && (GetJson_bible(bible_json_string)!="") {
				bible_text_string += "\n【鄒語聖經（新約）】\n" + GetJson_bible(bible_json_string) + "\n" //把 JSON 丟過去分離出要的內容
			}

			//https://www.bible.com/zh-TW/audio-bible-app-versions/1-kjv-king-james-version  //NRSV
			bible_json_string = HttpGET_("http://bible.fhl.net/json/qb.php?chineses=" + short_name + "&chap=" + chap_string + "&sec=" + sec_string + "&version=kjv")
			if (bible_json_string!="") && (GetJson_bible(bible_json_string)!="") {
				bible_text_string += "\n【英文 KJV（King James Version 英王欽定版本）】\n" + GetJson_bible(bible_json_string) + "\n" //把 JSON 丟過去分離出要的內容
			}

			bible_json_string = HttpGET_("http://bible.fhl.net/json/qb.php?chineses=" + short_name + "&chap=" + chap_string + "&sec=" + sec_string + "&version=bbe")
			if (bible_json_string!="") && (GetJson_bible(bible_json_string)!="") {
				bible_text_string += "\n【英文 BBE（Bible in Basic English 簡易英文譯本）】\n" + GetJson_bible(bible_json_string) + "\n" //把 JSON 丟過去分離出要的內容
			}

			//https://www.bible.com/zh-TW/audio-bible-app-versions/206-web-world-english-bible
			bible_json_string = HttpGET_("http://bible.fhl.net/json/qb.php?chineses=" + short_name + "&chap=" + chap_string + "&sec=" + sec_string + "&version=web")
			if (bible_json_string!="") && (GetJson_bible(bible_json_string)!="") {
				bible_text_string += "\n【英文 WEB（環球英文聖經）】\n" + GetJson_bible(bible_json_string) + "\n" //把 JSON 丟過去分離出要的內容
			}

			//https://www.bible.com/zh-TW/audio-bible-app-versions/12-asv-american-standard-version
			//http://www.freeasv.org/?page=listen
			bible_json_string = HttpGET_("http://bible.fhl.net/json/qb.php?chineses=" + short_name + "&chap=" + chap_string + "&sec=" + sec_string + "&version=asv")
			if (bible_json_string!="") && (GetJson_bible(bible_json_string)!="") {
				bible_text_string += "\n【英文 ASV（American Standard Version 美國標準譯本）】\n" + GetJson_bible(bible_json_string) + "\n" //把 JSON 丟過去分離出要的內容
			}

			bible_json_string = HttpGET_("http://bible.fhl.net/json/qb.php?chineses=" + short_name + "&chap=" + chap_string + "&sec=" + sec_string + "&version=darby")
			if (bible_json_string!="") && (GetJson_bible(bible_json_string)!="") {
				bible_text_string += "\n【英文 Darby 1890】\n" + GetJson_bible(bible_json_string) + "\n" //把 JSON 丟過去分離出要的內容
			}

			//https://bible.fhl.net/annouce/annouce426.html
			bible_json_string = HttpGET_("http://bible.fhl.net/json/qb.php?chineses=" + short_name + "&chap=" + chap_string + "&sec=" + sec_string + "&version=erv")
			if (bible_json_string!="") && (GetJson_bible(bible_json_string)!="") {
				bible_text_string += "\n【英文 ERV（English Revised Version 英國修訂譯本）】\n" + GetJson_bible(bible_json_string) + "\n" //把 JSON 丟過去分離出要的內容
			}

			//http://bible.salterrae.net/
			//http://bible.salterrae.net/kougo/html/
			bible_json_string = HttpGET_("http://bible.fhl.net/json/qb.php?chineses=" + short_name + "&chap=" + chap_string + "&sec=" + sec_string + "&version=jp")
			if (bible_json_string!="") && (GetJson_bible(bible_json_string)!="") {
				bible_text_string += "\n【日文（口語訳 Colloquil Japanese Version 1955）】\n" + GetJson_bible(bible_json_string) + "\n" //把 JSON 丟過去分離出要的內容
			}

			//https://www.bible.com/zh-TW/bible/88/2ki.1	88
			bible_json_string = HttpGET_("http://bible.fhl.net/json/qb.php?chineses=" + short_name + "&chap=" + chap_string + "&sec=" + sec_string + "&version=korean")
			if (bible_json_string!="") && (GetJson_bible(bible_json_string)!="") {
				bible_text_string += "\n【韓文 KRV（개역한글）】\n" + GetJson_bible(bible_json_string) + "\n" //把 JSON 丟過去分離出要的內容
			}

			//https://www.bible.com/zh-TW/bible/193/2ki.1	193
			bible_json_string = HttpGET_("http://bible.fhl.net/json/qb.php?chineses=" + short_name + "&chap=" + chap_string + "&sec=" + sec_string + "&version=vietnamese")
			if (bible_json_string!="") && (GetJson_bible(bible_json_string)!="") {
				bible_text_string += "\n【越南 VI1934（1934 Vietnamese Bible）】\n" + GetJson_bible(bible_json_string) + "\n" //把 JSON 丟過去分離出要的內容
			}

			//https://www.bible.com/zh-TW/bible/167/1ch.1	167
			bible_json_string = HttpGET_("http://bible.fhl.net/json/qb.php?chineses=" + short_name + "&chap=" + chap_string + "&sec=" + sec_string + "&version=russian")
			if (bible_json_string!="") && (GetJson_bible(bible_json_string)!="") {
				bible_text_string += "\n【俄文 SYNOD（Синодальный перевод）】\n" + GetJson_bible(bible_json_string) + "\n" //把 JSON 丟過去分離出要的內容
			}

			//http://c.ibibles.net/c02.php?1
			bible_json_string = HttpGET_("http://bible.fhl.net/json/qb.php?chineses=" + short_name + "&chap=" + chap_string + "&sec=" + sec_string + "&version=lxx")
			if (bible_json_string!="") && (GetJson_bible(bible_json_string)!="") {
				bible_text_string += "\n【舊約 希臘古譯 Εβδομήκοντα εκδοχή 七十士本 LXX】\n" + GetJson_bible(bible_json_string) + "\n" //把 JSON 丟過去分離出要的內容
			}

			bible_json_string = HttpGET_("http://bible.fhl.net/json/qb.php?chineses=" + short_name + "&chap=" + chap_string + "&sec=" + sec_string + "&version=bhs")
			if (bible_json_string!="") && (GetJson_bible(bible_json_string)!="") {
				bible_text_string += "\n【舊約 Hebraica Stuttgartensia הבראיקה 馬索拉原文 BHS】\n" + GetJson_bible(bible_json_string) + "\n" //把 JSON 丟過去分離出要的內容
			}

		}else{
			return "查詢章節超過聖經範圍，有可能指定查詢的節超過範圍。"		
		}
	}

	if bible_text_string != ""{
		return "[" + output_name + " " + chap_string + " : " +  sec_string + "]\n" + bible_text_string
	}else{
		return ""
	}
}

func Bible_print_string(short_name, output_name, chap_string, sec_string, ver_ string) (string) {
	bible_json_string := HttpGET_("http://bible.fhl.net/json/qb.php?chineses=" + short_name + "&chap=" + chap_string + "&sec=" + sec_string + "&version=" + ver_)
	bible_text_string := ""
	if bible_json_string!="" {
		if(GetJson_bible(bible_json_string)!=""){
			bible_text_string = GetJson_bible(bible_json_string)
		}else{
			return "查詢章節超過聖經範圍，有可能指定查詢的節超過範圍。"		
		}		
	}
			log.Print("GET = ")
			log.Print("bible_json_string = " + bible_json_string)
			log.Print("GetJson_bible = " + bible_text_string)

	if bible_text_string != ""{
		return "【聖經 " + output_name + " " + chap_string + "：" +  sec_string + "】\n" + bible_text_string
	}else{
		return ""
	}
}