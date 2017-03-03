
package main

import (
	"fmt"
	"log"

	"net/http"
	"os"
	"regexp"

	"strconv"
	"strings"
	
	"github.com/line/line-bot-sdk-go/linebot"

	// "bytes"

	"io/ioutil"

	"image/jpeg"

    "crypto/md5"
    "encoding/hex"

    // "encoding/json"
    	// "github.com/bitly/go-simplejson"
	"math/rand"
	"time"

)

var bot *linebot.Client

func main() {
	var err error
	bot, err = linebot.New(os.Getenv("ChannelSecret"), os.Getenv("ChannelAccessToken"))
	log.Println("Bot:", bot, " err:", err)
	http.HandleFunc("/callback", callbackHandler)
	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)
	http.ListenAndServe(addr, nil)
}

func random() int {  
    rnd := rand.New(rand.NewSource(time.Now().UnixNano()))  
    //for i := 0; i < 1; i++ {  
        //fmt.Println(rnd.Int())  
    //}
    return rnd.Int();
}

//https://gist.github.com/synr/d3d68d42b12204d981b39203a0b16762
func GetMD5Hash(text string) string {
    hasher := md5.New()
    hasher.Write([]byte(text))
    return hex.EncodeToString(hasher.Sum(nil))
}

func real_num(text string) string {
	text = strings.Replace(text, "１", "1", -1)
	text = strings.Replace(text, "２", "2", -1)
	text = strings.Replace(text, "３", "3", -1)
	text = strings.Replace(text, "４", "4", -1)
	text = strings.Replace(text, "５", "5", -1)
	text = strings.Replace(text, "６", "6", -1)
	text = strings.Replace(text, "７", "7", -1)
	text = strings.Replace(text, "８", "8", -1)
	text = strings.Replace(text, "９", "9", -1)
	text = strings.Replace(text, "０", "0", -1)
	return text
}



//http://qiita.com/koki_cheese/items/66980888d7e8755d01ec
// func handleTask(w http.ResponseWriter, r *http.Request) {
// }

	//修改時主要參考官方文件以及：
	// https://github.com/line/line-bot-sdk-go/blob/master/examples/kitchensink/server.go
		// KEY = handleText 等
	// https://github.com/dongri/line-bot-sdk-go
		// KEY = linebot.NewAudioMessage(originalContentURL, duration)
func callbackHandler(w http.ResponseWriter, r *http.Request) {
		// allow cross domain AJAX requests
		// http://stackoverflow.com/questions/12830095/setting-http-headers-in-golang/
		//	https://developer.mozilla.org/ja/docs/Web/HTTP/HTTP_access_control
	// w.Header().Set("Access-Control-Allow-Origin", "*")
	// w.Header().Set( "Access-Control-Allow-Methods","GET, POST, PUT, DELETE, OPTIONS" )
		//http://qiita.com/futosu/items/b49f7d9e28101daaa99e
		//https://play.golang.org/p/xHp44c_pJm
	// w.Header().Set("Access-Control-Allow-Headers","Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		// https://groups.google.com/forum/#!topic/golang-nuts/-Sh616lXNRE

	//-----------------------------------------------

	// log.Print("r")
	// log.Print(r)

	events, err := bot.ParseRequest(r)
												//c := appengine.NewContext(r)

	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}
	
	for _, event := range events {

		//2016.12.23+ 統一基本資訊集中
		//2016.12.24+ 嘗試抓使用者資訊 https://github.com/line/line-bot-sdk-go/blob/master/examples/kitchensink/server.go
		target_id_code := event.Source.UserID + event.Source.GroupID + event.Source.RoomID//target_id_code := ""
 		log.Print("event.Source.UserID = " + event.Source.UserID)
		log.Print("event.Source.GroupID = " + event.Source.GroupID)
		log.Print("event.Source.RoomID = " + event.Source.RoomID)
		log.Print("target_id_code = " + target_id_code)
		target_item := ""
		if event.Source.UserID!="" {
			target_item = "好友"
		}
		if event.Source.GroupID!="" {
			target_item = "群組對話"
		}
		if event.Source.RoomID!="" {
			target_item = "房間"
		}
		log.Print("target_item = " + target_item)

		username := ""
		userStatus := ""
		userImageUrl := ""
																				//userLogo_url := ""
		switch target_id_code{
			case "U77931c3187565323bcd8838498dcf651":
				username = "LL"
			case "U0a8152d2cea8c981aa2436a0ab596bca":
				username = "♪"
			case "C259ba6d5ace1c4932361612bf55394b3":
				username = "群組封測人員"
			case "Ceb12b46125861074c3b84106a5dcf750":
				username = "會友群組"
			case "C516948ebf27789543c92971e5ecf4a9b":
				username = "青少年"
			case "C0f81a7e0d9053cd305d2b27641e56090":
				username = "群組封測人員2"
			case "Ce82f808ee9915b3b58ca6955c2e1da61":
				username = "前任"
		}
		log.Print("username = " + username)

		//如果是群組會出錯，只能 1 對 1的時候。
		//if target_item == "好友"{
		if event.Source.UserID!="" {
			//2016.12.24+ 嘗試抓使用者資訊 https://github.com/line/line-bot-sdk-go/blob/master/examples/kitchensink/server.go
			profile, err := bot.GetProfile(event.Source.UserID).Do()
			if err != nil {
				log.Print(1162)
			    log.Print(err)
			}
			log.Print("profile.DisplayName = " + profile.DisplayName)			// println(res.Displayname)
			log.Print("profile.StatusMessage " + profile.StatusMessage)			// println(res.StatusMessage)
			log.Print("profile.PictureURL = " + profile.PictureURL)

														// log.Print("userLogo_url = " +  userLogo_url)
			//如果不是認識的 ID，就取得對方的名
			if username == ""{
				username = profile.DisplayName
			}
			userStatus = profile.StatusMessage
			userImageUrl = profile.PictureURL

			log.Print("username = " + username)
			log.Print("userStatus = " + userStatus)
			log.Print("userImageUrl = " + userImageUrl)

		}

		user_talk := ""
		if username == ""{
			user_talk = "【" + target_item + "】 " + target_id_code
		}else{
			user_talk = username
		}
		log.Print("※ user_talk = " + user_talk)

		//2016.12.27+

		// https://devdocs.line.me/en/#template-messages
		// HTTPS
		// JPEG or PNG
		// Aspect ratio: 1:1.51
		// Max width: 1024px
		// Max: 1 MB

		// 1024 = 1.51x
		// X = 678.145

		// 300 = 1.51x
		// x = 300 / 1.51 = 長的 / 1.51 = 198
		// 300 * 1.51 = 453 (用 300 當短)

		SystemImageURL := "https://trello-attachments.s3.amazonaws.com/52ff05f27a3c676c046c37f9/5831e5e304f9fac88ac50a23/d390ae079971c82074b5174c98899e9e/2017.png"
		//imageURL := SystemImageURL
		imageURL := "https://trello-attachments.s3.amazonaws.com/52ff05f27a3c676c046c37f9/585e3fb981c1240b4df88d73/34fba56ed5cbb5d7f0a9d7d2543ff238/C02018kong-koan_8481.JPG" //單位圖
		Bible_imageURL := "https://trello-attachments.s3.amazonaws.com/52ff05f27a3c676c046c37f9/585e3fa894a8bb9c02203e5a/e92ebf1eb7711389210725a8dc07261f/1024x678.jpg"

		//共用模板
		LineTemplate_chat := linebot.NewURITemplateAction("線上與開發者聊天", "http://www.smartsuppchat.com/widget?key=77b943aeaffa11a51bb483a816f552c70e322417&vid=" + target_id_code + "&lang=tw&pageTitle=%E9%80%99%E6%98%AF%E4%BE%86%E8%87%AA%20LINE%40%20%E9%80%B2%E4%BE%86%E7%9A%84%E5%8D%B3%E6%99%82%E9%80%9A%E8%A8%8A")
		LineTemplate_addme := linebot.NewURITemplateAction("加開發者 LINE", "https://line.me/R/ti/p/@uwk0684z")
		LineTemplate_download_app := linebot.NewURITemplateAction("下載 GOODTV APP", "http://www.goodtv.tv/app/")

		LineTemplate_CarouselColumn_feedback := linebot.NewCarouselColumn(
			SystemImageURL, "意見回饋 feedback", "你可以透過此功能\n對 LINE 機器人的 開發者 提出建議\n或錯誤回報、其他提案。",
			LineTemplate_addme,
			LineTemplate_chat,
			//linebot.NewMessageTemplateAction("聯絡 LINE 機器人開發者", "開發者"),
			linebot.NewPostbackTemplateAction("發訊息給教會", "取得發訊息給教會的提示",""),
		)

		LineTemplate_CarouselColumn_bible_list := linebot.NewCarouselColumn(
			Bible_imageURL, "聖經", "聖經查詢功能",
			linebot.NewPostbackTemplateAction("如何用我查聖經","如何查詢聖經","聖經"),
			linebot.NewPostbackTemplateAction("舊約列表","舊約列表","舊約列表"),
			linebot.NewPostbackTemplateAction("新約列表","新約列表","新約列表"),
		)

		LineTemplate_CarouselColumn_bible_one := linebot.NewCarouselColumn(
			Bible_imageURL, "主題經文", "依照主題隨機抽取經文！",
			linebot.NewURITemplateAction("隨機主題", "http://tool.ccnda.net/qr/view.jsp?ID=0"),
			linebot.NewURITemplateAction("解決問題的經文", "http://tool.ccnda.net/qr/view.jsp?ID=1779"),
			linebot.NewURITemplateAction("更多主題選擇", "http://tool.ccnda.net/qr/index.jsp"),
		)

		LineTemplate_firstinfo := linebot.NewCarouselTemplate(
			linebot.NewCarouselColumn(
				imageURL, "我是公館教會的小天使", "我可以幫大家取得教會資訊。\n可以邀我進群組方便更多人使用。這是一種資訊整合的便捷應用。",
				linebot.NewPostbackTemplateAction("本週週報 & 聚會時間", "週報 POST","週報"),
				linebot.NewPostbackTemplateAction("交通資訊","地圖 POST", "教會地圖"),
				linebot.NewPostbackTemplateAction("聯絡資訊","聯絡資訊 POST", "聯絡資訊"),
			),
			LineTemplate_CarouselColumn_bible_list,
			linebot.NewCarouselColumn(
				Bible_imageURL, "聖經查詢方法", "以下是示範。\n也可以手動輸入試試看各種組合。",
				linebot.NewPostbackTemplateAction("聖經 創世紀 1：1","聖經 創世紀 1：1","聖經 創世紀 1：1"),
				linebot.NewPostbackTemplateAction("英文聖經 出埃及 1：4-5","英文聖經 出埃及 1：4-5","英文聖經 出埃及 1：4-5"),
				linebot.NewPostbackTemplateAction("多國語言聖經 創世紀 1：1","多國語言聖經 創世紀 1：1","多國語言聖經 創世紀 1：1"),
			),
			LineTemplate_CarouselColumn_bible_one,
			// LineTemplate_other_example,
			// LineTemplate_other,
			LineTemplate_CarouselColumn_feedback,
		)
		    		//ImageURL_week_1 := "https://trello-attachments.s3.amazonaws.com/52ff05f27a3c676c046c37f9/585e3fb981c1240b4df88d73/3d66ad0e506a2fcc000b2d8aa229e3bf/wphhlgfd.jpg"
		    		//ImageURL_week_2 := "https://2.bp.blogspot.com/-klbjCCYRTQU/V8T5wT6AgMI/AAAAAAABeqA/B18lDyGcf00yA-k_rhC3m0iDj7IqRv3_ACLcB/s1600/%25E6%2588%2590%25E4%25B8%25BB%25E7%25A7%258B%25E5%25AD%25A3%25E7%258F%25AD%25E6%25B5%25B7%25E5%25A0%25B1.jpg"
		    		//ImageURL_week_3 := "https://trello-attachments.s3.amazonaws.com/52ff05f27a3c676c046c37f9/585e3fb981c1240b4df88d73/0b3cb02d676bf97e63654e3b43d0280d/Tatung.jpg"
		    		//obj_message_week_1 := linebot.NewImageMessage(ImageURL_week_1, ImageURL_week_1) //後面是預覽
		    		//obj_message_week_2 := linebot.NewImageMessage(ImageURL_week_2, ImageURL_week_2) //後面是預覽
		    		//obj_message_week_3 := linebot.NewImageMessage(ImageURL_week_3, ImageURL_week_3) //後面是預覽

					family_focus_text := `228 和平紀念日暫停一次`

		LineTemplate_family_focus := linebot.NewCarouselTemplate(

			linebot.NewCarouselColumn(
				imageURL, "家庭禮拜","二二八和平紀念日暫停一次",
				// linebot.NewURITemplateAction("幫你播打電話", "tel:0800550880"),
				//linebot.NewURITemplateAction("Google 地圖", "https://goo.gl/maps/fWophZAFGVL2"),
				//linebot.NewMessageTemplateAction("聯絡教會","聯絡資訊"),
				//linebot.NewURITemplateAction("詳細交通資訊", "http://www.kycare.com.tw/index.php/contacts/"),
				//linebot.NewPostbackTemplateAction("與教會聯絡","聯絡資訊 POST", "聯絡資訊"),
				linebot.NewPostbackTemplateAction("228 和平紀念日暫停一次","聯絡資訊 POST", "聯絡資訊"),
				linebot.NewPostbackTemplateAction("228 和平紀念日暫停一次","聯絡資訊 POST", "聯絡資訊"),
				linebot.NewURITemplateAction("報名家庭禮拜","http://blog.kkcpct.org/#contact"),
			),
			linebot.NewCarouselColumn(
				imageURL, "教會資訊", "我可以幫大家取得教會資訊。",
				linebot.NewPostbackTemplateAction("本週週報 & 聚會時間", "週報 POST","週報"),
				// linebot.NewPostbackTemplateAction("交通資訊","地圖 POST", "教會地圖"),
				linebot.NewPostbackTemplateAction("聯絡 & 交通資訊","聯絡資訊 POST", "聯絡資訊"),
				linebot.NewPostbackTemplateAction("網站資訊","網站資訊 POST", "網站資訊"),
			),
			//LineTemplate_CarouselColumn_bible_one,
			// LineTemplate_other_example,
			// LineTemplate_other,
			LineTemplate_CarouselColumn_feedback,
		)
		t_family_focus := "想報名家庭禮拜或有相關疑問，請線上報名：http://blog.kkcpct.org/#contact\n或 聯絡 台北公館教會 諮詢：\n\n電話：02-29327941\n傳真：02-29345003\n電子郵件：kkcpct@ms29.hinet.net\n通訊地址：11677 台北市汀州路四段 85 巷 2 號\n網站留言：http://blog.kkcpct.org/#contact\n提示：\n這部分在最新版本 LINE APP 會以按鈕呈現，\n可幫助快速撥打電話等通訊聯絡。"
		obj_message_family_focus := linebot.NewTemplateMessage(t_family_focus, LineTemplate_family_focus)

					//this_week_text_1 :=`從缺，本週週報資料還沒送進來。`

					ImageURL_week_1 := "https://trello-attachments.s3.amazonaws.com/5897c415be467e461197bfbd/58b93ca932ee2886e0caf993/553faf9dcaae371565cceaee26f50627/20170305.index_mini.jpg"
					obj_message_week_1 := linebot.NewImageMessage(ImageURL_week_1, ImageURL_week_1)

					this_week_text_2 := `本週週報（2017/03/05）：
https://s.kkcpct.org/2mi0jch

近期週報收藏區：
https://week.kkcpct.org

歡迎使用電子郵件訂閱週報：
https://news.kkcpct.org`


		weektime_msg := "台北公館教會的聚會資訊：\n\n" +
						"主日聚會時間：\n上午 09:00（台語）\n上午 11:00（華語）\n\n" +
						"早天祈禱會：\n每週日 上午 07:30\n\n" +
						"青少年團契：\n每週日 上午 09:00\n\n" +
						"兒童主日學：\n每週日 上午 09:00\n\n" +
						"松年團契：\n每週日 上午 11:00\n\n" +
						"成人主日學：\n每週日 上午 12:30\n\n" +
						"婦女團契：\n每週二 上午 09:00\n\n" +
						"查經 / 祈禱會：\n每週三 晚上08:00\n\n" +
						"聖歌隊：\n每週五 晚上08:00\n\n" +
						"社青團契：\n每週六 晚上07:00\n\n" +
						"學青團契：\n每週六 晚上07:00"

		next_week_msg := "本週主日（2017/03/05）" +
					"\n" +
					"預定行事：\n婦女事工紀念主日\n" +
					"靈修會籌備會\n" +
					"兒童營籌備會\n" +
					"\n" +
					"台語禮拜：吳慧馨 牧師\n" +
					"主裡尊貴的你：詩篇 8:1-9\n" +
					"\n" +
					"♪ 聖詩 ♪：196、396\n" +
					"啟應文：第 23 篇\n" +
					"\n" +
					"獻詩：婦女團契\n" +
					"司禮：李麗珠\n" +
					"敬拜讚美團：雅歌（李君振）\n" +
					"司獻：黃美俐、陳勝雄、黃全業\n" +
					"司琴：王孟治\n" +
					"招待：李碧珠、陳勝彥、吳欽英\n" +
					"音控：李武智\n" +
					"輪值長老：李君振、李妍蓉\n" +
					"獻花者：楊光華\n" +
					"\n" +
					"----\n" +
					"\n" +
					"華語禮拜：吳慧馨 牧師\n" +
					"返璞歸真：馬太福音 4:1-11\n" +
					"\n" +
					"♪：只要有祢在我左右（讚美之泉）\n" +
					"\n" +
					"獻詩：婦女團契\n" +
					"司禮、敬拜讚美團、司獻、司琴、音控：少年團契\n" +
					"招待：白壽雄、陳嘉崇、吳秀鳳\n" +
					"輪值長老：李君振、李妍蓉\n" +
					"獻花者：楊光華\n" +
					"\n" +
					"其他請參閱本季行事表：\nhttps://calendar.kkcpct.org/table\n\n"+
					"教會 Google 行事曆：\nhttps://calendar.kkcpct.org\nhttps://s.kkcpct.org/2kBvGyR\n\n" + 
					"Google 月曆訂閱 ics 檔：\nhttps://calendar.kkcpct.org/ics\n\n" +
					"iCalendar Feeds 使用方法請參照：https://s.kkcpct.org/2mbZYER"

		LineTemplate_nextweek_review := linebot.NewCarouselTemplate(

			linebot.NewCarouselColumn(
				Bible_imageURL, "主日禮拜經文預習", "以下是下週（2017/03/05）的經文",
				linebot.NewPostbackTemplateAction("台語禮拜 詩篇 8:1-9","第一場預習","中文聖經新譯本 詩篇 8:1-9"),
				linebot.NewPostbackTemplateAction("華語禮拜 馬太福音 4:1-11","第一場預習","中文聖經 馬太福音 4:1-11"),
				// linebot.NewPostbackTemplateAction("上週華語禮拜 行在水面上","第二場預習","行在水面上"),
				// linebot.NewMessageTemplateAction("瞭解聖經查詢方法","聖經"),
				linebot.NewPostbackTemplateAction("本週禮拜內容預告","下週預告 POST", "下週預告"),
				//linebot.NewPostbackTemplateAction("本週新眼光","新眼光 POST", "新眼光"),
			),
			linebot.NewCarouselColumn(
				Bible_imageURL, "聖經查詢方法", "以下是示範。\n也可以手動輸入試試看各種組合。",
				linebot.NewPostbackTemplateAction("本週新眼光","新眼光 POST", "新眼光"),
				//linebot.NewPostbackTemplateAction("聖經 創世紀 1：1","聖經 創世紀 1：1","聖經 創世紀 1：1"),
				linebot.NewPostbackTemplateAction("英文聖經 出埃及 1：4","英文聖經 出埃及 1：4","英文聖經 出埃及 1：4"),
				linebot.NewPostbackTemplateAction("多國語言聖經 創世紀 1：1","多國語言聖經 創世紀 1：1","多國語言聖經 創世紀 1：1"),
			),
			linebot.NewCarouselColumn(
				imageURL, "教會資訊", "其他可以幫大家取得的教會資訊",
				//linebot.NewPostbackTemplateAction("本週週報 & 聚會時間", "週報 POST","週報"),
				linebot.NewURITemplateAction("Mail 訂閱週報","https://news.kkcpct.org"),
				linebot.NewPostbackTemplateAction("交通資訊","地圖 POST", "教會地圖"),
				linebot.NewPostbackTemplateAction("聯絡資訊","聯絡資訊 POST", "聯絡資訊"),
			),
			LineTemplate_CarouselColumn_bible_one,
			// LineTemplate_other_example,
			// LineTemplate_other,
			LineTemplate_CarouselColumn_feedback,
		)
		t_nextweek_review := "台北公館教會的聯絡資訊：\n\n電話：02-29327941\n傳真：02-29345003\n電子郵件：kkcpct@ms29.hinet.net\n通訊地址：11677 台北市汀州路四段 85 巷 2 號\n\n提示：\n這部分在最新版本 LINE APP 會以預習經文按鈕呈現，\n可幫助會眾快速熟悉下週經文。"
		obj_message_nextweek_review := linebot.NewTemplateMessage(t_nextweek_review, LineTemplate_nextweek_review)


		LineTemplate_nextweek_read := linebot.NewCarouselTemplate(
			linebot.NewCarouselColumn(
				Bible_imageURL, "本週新眼光讀經", "2017/02/26～2017/02/28",
				linebot.NewPostbackTemplateAction("26 日 俄巴底亞書 1:1-14","新眼光查經","聖經 俄巴底亞書 1:1-14"),
				linebot.NewPostbackTemplateAction("27 日 俄巴底亞書 1:15-16","新眼光查經","聖經 俄巴底亞書 1:15-16"),
				linebot.NewPostbackTemplateAction("28 日 俄巴底亞書 1:17-21","新眼光查經","聖經 俄巴底亞書 1:17-21"),
			),
			linebot.NewCarouselColumn(
				Bible_imageURL, "本週新眼光讀經", "2017/03/01～2017/03/03",
				linebot.NewPostbackTemplateAction("1 日 約拿書 1:1-3","新眼光查經","聖經 約拿書 1:1-3"),
				linebot.NewPostbackTemplateAction("2 日 約拿書 1:4-17","新眼光查經","聖經 約拿書 1:4-17"),
				linebot.NewPostbackTemplateAction("3 日 約拿書 2:1-10","新眼光查經","聖經 約拿書 2:1-10"),
			),
			linebot.NewCarouselColumn(
				Bible_imageURL, "本週新眼光讀經", "2017/03/04",
				linebot.NewPostbackTemplateAction("4 日 約拿書 3:1-10","新眼光查經","聖經 約拿書 3:1-10"),
				// linebot.NewPostbackTemplateAction("25 日 阿摩司書 9:11-15","新眼光查經","聖經 阿摩司書 9:11-15"),
				linebot.NewMessageTemplateAction("瞭解聖經查詢方法","聖經"),
				linebot.NewMessageTemplateAction("本季行事曆","行事曆"),
			),
			linebot.NewCarouselColumn(
				Bible_imageURL, "聖經查詢方法", "以下是示範。\n也可以手動輸入試試看各種組合。",
				linebot.NewPostbackTemplateAction("聖經 創世紀 1：1","聖經 創世紀 1：1","聖經 創世紀 1：1"),
				linebot.NewPostbackTemplateAction("英文聖經 出埃及 1：4","英文聖經 出埃及 1：4","英文聖經 出埃及 1：4"),
				linebot.NewPostbackTemplateAction("多國語言聖經 創世紀 1：1","多國語言聖經 創世紀 1：1","多國語言聖經 創世紀 1：1"),
			),
			linebot.NewCarouselColumn(
				imageURL, "教會資訊", "我可以幫大家取得教會資訊。",
				linebot.NewPostbackTemplateAction("本週週報 & 聚會時間", "週報 POST","週報"),
				linebot.NewPostbackTemplateAction("交通資訊","地圖 POST", "教會地圖"),
				linebot.NewPostbackTemplateAction("聯絡資訊","聯絡資訊 POST", "聯絡資訊"),
			),
			//LineTemplate_CarouselColumn_bible_one,
			// LineTemplate_other_example,
			// LineTemplate_other,
			//LineTemplate_CarouselColumn_feedback,
		)
		t_nextweek_read := `本週新眼光讀經

26 日 俄巴底亞書 1:1-14
27 日 俄巴底亞書 1:15-16
28 日 俄巴底亞書 1:17-21
3 月 1 日 約拿書 1:1-3
2 日 約拿書 1:4-17
3 日 約拿書 2:1-10
4 日 約拿書 3:1-10` + "\n\n提示：\n這部分在最新版本 LINE APP 會以預習經文按鈕呈現，\n可幫助會眾快速熟悉經文。"
		obj_message_nextweek_read := linebot.NewTemplateMessage(t_nextweek_read, LineTemplate_nextweek_read)

		//正題

		//只會抓到透過按鈕按下去的東西。方便做新的觸發點。(缺點是沒有 UI 介面的時候會無法使用)
		if event.Type == linebot.EventTypePostback {
				//這裡用來設計按下某按鈕後要做什麼事情
				log.Print("觸發 Postback 功能（不讓使用者察覺的程式利用）")
				log.Print("event.Postback.Data = " + event.Postback.Data)
				HttpPost_JANDI(target_item + " " + "[" + user_talk + "](" + userImageUrl + ") 觸發了按鈕並呼了 event.Postback.Data = " + event.Postback.Data + `\n` + userStatus, "brown" , "LINE 程式觀察",target_id_code)
				HttpPost_IFTTT(target_item + " " + user_talk + " 觸發了按鈕並呼了 event.Postback.Data = " + event.Postback.Data + `\n<br>` + userImageUrl + `\n<br>` + userStatus , "LINE 程式觀察" ,target_id_code)
				HttpPost_Zapier(target_item + " " + "[" + user_talk + "](" + userImageUrl + ") 觸發了按鈕並呼了 event.Postback.Data = " + event.Postback.Data + `\n` + userStatus, "LINE 程式觀察" ,target_id_code,user_talk)

				// if event.Postback.Data == "測試"{

				// }

				if event.Postback.Data == "取得發訊息給教會的提示"{
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("請手動輸入「給教會：」開頭，後面接上你想給教會的話即可。\n\n發出後會立即送出。\n如需要回覆，\n請在訊息文字中也附上您的聯絡方式，謝謝！")).Do(); err != nil {
							log.Print(507)
							log.Print(err)
					}
				}

				

				if event.Postback.Data == "週報"{
					

					//if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(),obj_message_week_2,obj_message_week_3,linebot.NewTextMessage(next_week_msg),obj_message_nextweek_review).Do(); err != nil {
					//if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(this_week_text_1),linebot.NewTextMessage(this_week_text_2),obj_message_nextweek_read,linebot.NewTextMessage(next_week_msg),obj_message_nextweek_review).Do(); err != nil {
					//if _, err = bot.ReplyMessage(event.ReplyToken, obj_message_nextweek_read,linebot.NewTextMessage(next_week_msg),obj_message_week_1,linebot.NewTextMessage(this_week_text_2),obj_message_nextweek_review).Do(); err != nil {
					//if _, err = bot.ReplyMessage(event.ReplyToken, obj_message_nextweek_read,linebot.NewTextMessage(next_week_msg),obj_message_week_1,linebot.NewTextMessage(this_week_text_2),obj_message_nextweek_review).Do(); err != nil {
					if _, err = bot.ReplyMessage(event.ReplyToken, obj_message_nextweek_review,obj_message_week_1,linebot.NewTextMessage(this_week_text_2)).Do(); err != nil {
							log.Print(486)
							log.Print(err)
					}
				}

				if event.Postback.Data == "聚會時間"{
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(weektime_msg)).Do(); err != nil {
							log.Print(493)
							log.Print(err)
					}
				}

				if event.Postback.Data == "家庭禮拜"{
					//obj_message := linebot.NewStickerMessage("2", "514") //https://devdocs.line.me/en/?go#send-message-object
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(family_focus_text),obj_message_family_focus).Do(); err != nil {
							log.Print(500)
							log.Print(err)
					}
				}

				if event.Postback.Data == "電子郵件"{
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("教會 E-mail 信箱地址：\nkkcpct@ms29.hinet.net")).Do(); err != nil {
							log.Print(507)
							log.Print(err)
					}
				}





				if event.Postback.Data == "test"{


					// https://devdocs.line.me/en/#imagemap-message
					// "x": 0,
     				//	"y": 0,
		   			// "width": 520,
		   			// "height": 1040

		   			//log.Print("test MD5 = " + GetMD5Hash(event.Postback.Data))

		   			//測試圖片地圖
					obj_message := linebot.NewImagemapMessage(
							"https://synr.github.io/test",
							"Imagemap alt text",
							linebot.ImagemapBaseSize{1040, 1040},
							linebot.NewURIImagemapAction("https://store.line.me/family/manga/en", linebot.ImagemapArea{0, 0, 520, 520}),
							linebot.NewURIImagemapAction("https://store.line.me/family/music/en", linebot.ImagemapArea{520, 0, 520, 520}),
							linebot.NewURIImagemapAction("https://store.line.me/family/play/en", linebot.ImagemapArea{0, 520, 520, 520}),
							linebot.NewMessageImagemapAction("URANAI!", linebot.ImagemapArea{520, 520, 520, 520}),	//上限 400 字
					)

					if _, err := bot.ReplyMessage(event.ReplyToken,obj_message).Do(); err != nil {
						log.Print(1586)
						log.Print(err)
					}
				}

				if event.Postback.Data == "開啟管理者選單"{
					switch target_id_code {
						case "U6f738a70b63c5900aa2c0cbbe0af91c4":
							imageURL = SystemImageURL
							LineTemplate_test := linebot.NewCarouselTemplate(
								linebot.NewCarouselColumn(
									imageURL, "管理模式", "測試中",
									linebot.NewURITemplateAction("本季行事曆","https://docs.google.com/spreadsheets/d/1RYchaiPFyPNzCC7paUzg4tfcJ6Y2XlsASer4V5K4_eU/pubhtml"),
									linebot.NewPostbackTemplateAction("無效選項","admin", ""),
									linebot.NewPostbackTemplateAction("登出","登出管理者", ""),
								),
								// LineTemplate_other_example,
								// LineTemplate_other,
								LineTemplate_CarouselColumn_feedback,
							)
							no_temp_msg := "請更新使用最新版本的 LINE APP 才能查看內容 。"
							obj_message := linebot.NewTemplateMessage(no_temp_msg, LineTemplate_test)
							if _, err = bot.ReplyMessage(event.ReplyToken, obj_message).Do(); err != nil {
									log.Print(572)
									log.Print(err)
							}
						default:
					}
				}

				if event.Postback.Data == "passcheck"{
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("請輸入暗號驗證管理者身分")).Do(); err != nil {
							log.Print(1929)
							log.Print(err)
					}
				}

				//2017.01.03+
				if event.Postback.Data == "admin"{
					switch target_id_code {
						case "U6f738a70b63c5900aa2c0cbbe0af91c4":
							imageURL = SystemImageURL
							LineTemplate_test := linebot.NewCarouselTemplate(
								linebot.NewCarouselColumn(
									imageURL, "管理模式", "For ADMIN mode.",
									linebot.NewPostbackTemplateAction("核對「暗號」","passcheck", ""),
									linebot.NewPostbackTemplateAction("管理模式","admin", ""),
									linebot.NewPostbackTemplateAction("測試圖片地圖","test", ""),
								),
								// LineTemplate_other_example,
								// LineTemplate_other,
								LineTemplate_CarouselColumn_feedback,
							)
							no_temp_msg := "請更新使用最新版本的 LINE APP 才能查看內容 。"
							obj_message := linebot.NewTemplateMessage(no_temp_msg, LineTemplate_test)
							if _, err = bot.ReplyMessage(event.ReplyToken, obj_message).Do(); err != nil {
									log.Print(605)
									log.Print(err)
							}
						default:
							if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("您無法使用此功能。")).Do(); err != nil {
									log.Print(1955)
									log.Print(err)
							}
					}
				}

				if event.Postback.Data == "登出管理者"{
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("你已登出管理模式")).Do(); err != nil {
						log.Print(1965)
						log.Print(err)
					}
				}






				if event.Postback.Data == "取消離開群組"{
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("你已經取消請我離開 :)")).Do(); err != nil {
						log.Print(1207)
						log.Print(err)
					}
				}

				//2016.12.26+
				if event.Postback.Data == "按下確定離開群組對話"{
					template := linebot.NewCarouselTemplate(
						linebot.NewCarouselColumn(
							SystemImageURL, "請機器人離開群組", "你確定要請我離開嗎QAQ？\n如果確定請按下方按鈕 QQ",
							linebot.NewPostbackTemplateAction("請機器人離開群組","離開群組", "機器人已經自動離開。\n如要加回來請找：\nhttps://line.me/R/ti/p/@rgp6918s\n如要聯絡開發者請找：\nhttps://line.me/R/ti/p/@uwk0684z"),
							//linebot.NewPostbackTemplateAction("請機器人離開群組","離開群組", "機器人已經自動離開。\n如要加回來請找：\nhttps://line.me/R/ti/p/@sjk2434l\n如要聯絡開發者請找：\nhttps://line.me/R/ti/p/@uwk0684z"),
							LineTemplate_addme,
							LineTemplate_chat,
						),
					)
					obj_message := linebot.NewTemplateMessage("這是命令機器人自己離開群組的方法。\n這功能只支援 APP 使用。\n請用 APP 端查看下一步。", template)
					if _, err = bot.ReplyMessage(event.ReplyToken, obj_message).Do(); err != nil {
						log.Print(1225)
						log.Print(err)
					}
				}

				if event.Postback.Data == "離開群組"{
					if target_item == "群組對話" {
						if _, err := bot.LeaveGroup(target_id_code).Do(); err != nil {
							log.Print(1233)
						    log.Print(err)
						}
						//HttpPost_JANDI("自動離開 "  + user_talk , "gray" , "LINE 離開群組",target_id_code)
						//HttpPost_IFTTT("自動離開 "  + user_talk , "LINE 離開群組",target_id_code)
						HttpPost_Zapier("自動離開 "  + user_talk , "LINE 離開群組",target_id_code,user_talk)
						log.Print("觸發自動離開 " + user_talk +  " 群組")
					}
				}
		}
		//觸發加入好友
		if event.Type == linebot.EventTypeFollow {
				HttpPost_JANDI("有新的好朋友：["  + user_talk + "](" + userImageUrl  + ")" + `\n` + userStatus, "blue" , "LINE 新好友",target_id_code)
				HttpPost_IFTTT("有新的好朋友："  + user_talk  + `\n<br>` + userImageUrl + `\n<br>` + userStatus, "LINE 新好友" ,target_id_code)
				HttpPost_Zapier("有新的好朋友：["  + user_talk + "](" + userImageUrl  + ")" + `\n` + userStatus, "LINE 新好友" ,target_id_code,user_talk)
				//target_id_code := event.Source.UserID + event.Source.GroupID + event.Source.RoomID	//target_id_code := ""
				log.Print("觸發與 " + user_talk + " 加入好友")

			    imageURL = SystemImageURL
				//template := LineTemplate_firstinfo
				t_msg := "請用最新版本的 LINE APP 觀看效果。\n如有其他建議或想討論，請對我輸入「開發者」進行聯絡。"
				obj_message := linebot.NewTemplateMessage(t_msg, LineTemplate_firstinfo)

				// username := ""
				// if target_id_code == "U6f738a70b63c5900aa2c0cbbe0af91c4"{
				// 	username = "LL"
				// }
				// if target_id_code == "Uf150a9f2763f5c6e18ce4d706681af7f"{
				// 	username = "包包"
				// }
				//reply 的寫法
				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("你好啊！" + username + "～\n想知道我的功能，可以說：「簡介」\n單獨輸入「聖經」可以知道查詢方法。\n\nPS：最新版的 LINE APP 上可以看到不一樣的內容喔！"),obj_message).Do(); err != nil {
						log.Print(1288)
						log.Print(err)
				}
		}
		//觸發解除好友
		if event.Type == linebot.EventTypeUnfollow {
				HttpPost_JANDI("與 ["  + user_talk + "](" + userImageUrl + ") 解除好友" + `\n` + userStatus, "gray" , "LINE 被解除好友",target_id_code)
				HttpPost_IFTTT("與 "  + user_talk + " 解除好友" + `\n<br>` + userImageUrl + `\n<br>` + userStatus , "LINE 被解除好友" ,target_id_code)
				HttpPost_Zapier("與 ["  + user_talk + "](" + userImageUrl + ") 解除好友" + `\n` + userStatus , "LINE 被解除好友" ,target_id_code,user_talk)
				log.Print("觸發與 " + user_talk + " 解除好友")
		}
		//觸發加入群組聊天
		if event.Type == linebot.EventTypeJoin {
				HttpPost_JANDI("加入了 "  + user_talk , "blue" , "LINE 已加入群組",target_id_code)
				HttpPost_IFTTT("加入了 "  + user_talk , "LINE 已加入群組" ,target_id_code)
				HttpPost_Zapier("加入了 "  + user_talk , "LINE 已加入群組" ,target_id_code,user_talk)
				log.Print("觸發加入" + user_talk)
 				//source := event.Source
 				//log.Print("觸發加入群組聊天事件 = " + source.GroupID)

 				push_string := "很高興你邀請我進來這裡聊天！"

				switch username{
					case "群組封測人員2","青少年","前任":
						push_string = "很高興你邀請我進來這裡聊天！\n" + "我有針對這個群組才有的特殊功能！\n\n詳細用法請詢問羅傳道！"
				}

				//if source.GroupID == "Ca78bf89fa33b777e54b4c13695818f81"{
				if target_id_code == "Ca78bf89fa33b777e54b4c13695818f81"{
					push_string += "\n你好，" + user_talk + "。"
				}
					//push 的寫法
					// 				if _, err = bot.PushMessage(source.GroupID, linebot.NewTextMessage(push_string)).Do(); err != nil {
					// 					log.Print(err)
					// 				}
					// 				if _, err = bot.PushMessage("Ca78bf89fa33b777e54b4c13695818f81", linebot.NewTextMessage("這裡純測試對嗎？\n只發於測試聊天室「test」")).Do(); err != nil {
					// 					log.Print(err)
					// 				}
					//target_id_code := event.Source.UserID + event.Source.GroupID + event.Source.RoomID	//target_id_code := ""
			    imageURL = SystemImageURL
				//template := LineTemplate_firstinfo
				t_msg := "請用最新版本的 LINE APP 觀看效果。\n如有其他建議或想討論，請對我輸入「開發者」進行聯絡。"
				obj_message := linebot.NewTemplateMessage(t_msg, LineTemplate_firstinfo)

				//reply 的寫法
				switch username{
					case "群組封測人員2","青少年","前任":
						if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("群組聊天的各位大家好哇～！\n" + push_string),linebot.NewTextMessage("這是一種資訊整合的便捷應用，效果類似於自動回話小助理。\n\n概念上最基本的應用類似於遊戲 NPC 或 0800 電話總機，會根據指示自動回覆相關基本資訊。\n也可做其他延伸應用，像是聖經查詢 或 留言給意見...等等。\n\n目前除了教會相關資訊外，還可查詢 24 本聖經。\n支援 10 種語言、24 種聖經版本的精準經節查詢機能。\n並支援範圍查詢的寫法。（例如：聖經 創世紀 1:1-10）\n\n詳細說明可輸入「聖經」，有完整的使用說明介紹。"),obj_message).Do(); err != nil {
								log.Print(1351)
								log.Print(err)
						}
					default:
						if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("群組聊天的各位大家好哇～！\n" + push_string + "\n\n想知道我的功能，請說：「簡介」\n單獨輸入「聖經」可以知道查詢方法。"),linebot.NewTextMessage("這是一種資訊整合的便捷應用，效果類似於自動回話小助理。\n\n概念上最基本的應用類似於遊戲 NPC 或 0800 電話總機，會根據指示自動回覆相關基本資訊。\n也可做其他延伸應用，像是聖經查詢 或 留言給意見...等等。\n\n目前除了教會相關資訊外，還可查詢 24 本聖經。\n支援 10 種語言、24 種聖經版本的精準經節查詢機能。\n並支援範圍查詢的寫法。（例如：聖經 創世紀 1:1-10）\n\n詳細說明可輸入「聖經」，有完整的使用說明介紹。"),obj_message).Do(); err != nil {
								log.Print(1351)
								log.Print(err)
						}
				}
		}
		//觸發離開群組聊天
		if event.Type == linebot.EventTypeLeave {
				HttpPost_JANDI("被請離開 "  + user_talk , "gray" , "LINE 離開群組",target_id_code)
				HttpPost_IFTTT("被請離開 "  + user_talk , "LINE 離開群組",target_id_code)
				HttpPost_Zapier("被請離開 "  + user_talk , "LINE 離開群組",target_id_code,user_talk)
				log.Print("觸發被踢出 " + user_talk +  " 群組")
		}
		//？？？？？
			//https://admin-official.line.me/beacon/register
			//https://devdocs.line.me/en/#line-beacon
			//https://devdocs.line.me/ja/#line-beacon
			//這個應用要有硬體
		if event.Type == linebot.EventTypeBeacon {
			HttpPost_JANDI(target_item + " " + "[" + user_talk + "](" + userImageUrl + ") 觸發 Beacon（啥鬼）" + `\n` + userStatus, "yellow" , "LINE 對話同步",target_id_code)
			HttpPost_IFTTT(target_item + " " + user_talk + " 觸發 Beacon（啥鬼）" + `\n<br>` + userImageUrl + `\n<br>` + userStatus, "LINE 對話同步",target_id_code)
			HttpPost_Zapier(target_item + " " + "[" + user_talk + "](" + userImageUrl + ") 觸發 Beacon（啥鬼）" + `\n` + userStatus, "LINE 對話同步",target_id_code,user_talk)
			log.Print(user_talk + " 觸發 Beacon（啥鬼）")
		}
		//觸發收到訊息
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
	 			//message.ID
				//message.Text
				// log.Print(message.ID)
				// log.Print(message.Text)
				//bot_msg := "你是說 " + message.Text + " 嗎？\n\n我看看喔...等我一下..."

				// if target_id_code == "Ceb12b46125861074c3b84106a5dcf750"{
				// 	if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("感謝您呼叫我！\n但為了不要過度打擾群組對話，在此群組暫時不支援此功能。\n請移步到 http://line.me/R/ti/g/X8_PRxyGL7 對話群組進行查詢。\n或個人+我（https://line.me/R/ti/p/@rgp6918s）做好友查詢。").Do(); err != nil {
				// 			log.Print(24483)
				// 			log.Print(err)
				// 	}
				// 	return
				// }

					bot_msg_lo := lo(message.Text)
					if ((bot_msg_lo != "") && (random()>1))  {
						switch username{
							case "群組封測人員2","青少年","前任":
								switch bot_msg_lo{
									case "5fed8bfd031fa698e125567b128d1024":
								// if username == "LL" {
								// 	if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("登入成功！")).Do(); err != nil {
								// 		log.Print(2162)
								// 		log.Print(err)
								// 	}
								// }

									//if target_id_code == "U6f738a70b63c5900aa2c0cbbe0af91c4"{
										//imageURL := "https://i2.bahamut.com.tw/anime/crazy_logo.png" //單位圖

										LineTemplate_family_focus := linebot.NewCarouselTemplate(
										    linebot.NewCarouselColumn(
										        "https://trello-attachments.s3.amazonaws.com/5897c415be467e461197bfbd/58a397cbf644e94a218a1ff0/2521b59aba6ceecd77177b55f50e512c/QT.jpg", "每日靈修", "每日靈修",
										        //linebot.NewURITemplateAction("♪ 詩歌 ♪", "tel:02"),
										        linebot.NewURITemplateAction("2017 每日靈修短文", "https://www.dropbox.com/sh/djnixzfkp36k4tn/AADatyCaOJRvLvtBbNn2SzTza?dl=0"),
										        //linebot.NewMessageTemplateAction("聯絡教會","聯絡資訊"),
										    ),
										)
										t_family_focus := "請使用最新版的 LINE APP 查看此處"
										obj_message_family_focus := linebot.NewTemplateMessage(t_family_focus, LineTemplate_family_focus)

										if _, err = bot.ReplyMessage(event.ReplyToken,obj_message_family_focus,linebot.NewTextMessage("此內容僅限本群組使用，內容如有缺漏請回報羅傳道。")).Do(); err != nil {
										    log.Print(1042)
										    log.Print(err)
										}
										return
									//}
									case "546eabd81e99ec08b8b0301af80310d9":
										LineTemplate_family_focus := linebot.NewCarouselTemplate(
										    linebot.NewCarouselColumn(
										        "https://trello-attachments.s3.amazonaws.com/5897c415be467e461197bfbd/58a397cbf644e94a218a1ff0/b986aa3d94787dc232b57b4b794e44c8/worship.jpg", "♪ 詩歌 ♪", "詩歌、敬拜讚美...等。",
										        linebot.NewURITemplateAction("♪ 詩歌 ♪", "https://www.dropbox.com/sh/0yfeslxsksnambx/AABZbqumA3vS_Wa9nnr_2IDOa?dl=0"),
										        //linebot.NewURITemplateAction("2017 每日靈修短文", "https://goo.gl/maps"),
										        //linebot.NewMessageTemplateAction("聯絡教會","聯絡資訊"),
										    ),
										)
										t_family_focus := "請使用最新版的 LINE APP 查看此處"
										obj_message_family_focus := linebot.NewTemplateMessage(t_family_focus, LineTemplate_family_focus)

										if _, err = bot.ReplyMessage(event.ReplyToken,obj_message_family_focus,linebot.NewTextMessage("此內容僅限本群組使用，內容如有缺漏請回報羅傳道。")).Do(); err != nil {
										    log.Print(1060)
										    log.Print(err)
										}		
										return		
								}
								return
							default:
								log.Print(target_id_code)
								log.Print("下面有東西嗎")
								log.Print(random())
								log.Print(random()>1860081517322633183)
								return
						}
					}


				//安靜名單
				notalk := "Ceb12b46125861074c3b84106a5dcf750"//"C259ba6d5ace1c4932361612bf55394b3"//"Ceb12b46125861074c3b84106a5dcf750" //

				//給教會
				bot_msg := send_to_JANDI(message.Text, target_item, user_talk, userImageUrl, userStatus, target_id_code)
				log.Print("給完教會後是多少？ bot_msg = " + bot_msg)
				bible_chap := ""
				bible_sec := ""
				bible_short_name := ""

				if bot_msg != "已經傳送給教會"{
					//bible()
					//bot_msg, bible_chap, bible_sec, bible_short_name := bible(message.Text,target_id_code,"")//bot_msg = bible(message.Text,message.ID,"")
					bot_msg, bible_chap, bible_sec, bible_short_name = bible(bot_msg,target_id_code,"")//bot_msg = bible(message.Text,message.ID,"")
				}

				bible_id := "0"
				//bible_com_text :=""
				switch bible_short_name {
					case "創":
						bible_id = "1"
						//bible_com_text = "gen"
					case "出":
						bible_id = "2"
						//bible_com_text = "ex"
					case "利":
						bible_id = "3"
						//bible_com_text = "lev"
					case "民":
						bible_id = "4"
						//bible_com_text = "num"
					case "申":
						bible_id = "5"
						//bible_com_text = "deu"
					case "書":
						bible_id = "6"
						//bible_com_text = "jos"
					case "士":
						bible_id = "7"
						//bible_com_text = "jdg"
					case "得":
						bible_id = "8"
						//bible_com_text = "rut"
					case "撒上":
						bible_id = "9"
						//bible_com_text = "1sa"
					case "撒下":
						bible_id = "10"
						//bible_com_text = "2sa"
					case "王上":
						bible_id = "11"
						//bible_com_text = "1ki"
					case "王下":
						bible_id = "12"
						//bible_com_text = "2ki"
					case "代上":
						bible_id = "13"
						//bible_com_text = "1ch"
					case "代下":
						bible_id = "14"
						//bible_com_text = "2ch"
					case "拉":
						bible_id = "15"
						//bible_com_text = "ezr"
					case "尼":
						bible_id = "16"
						//bible_com_text = "neh"
					case "斯":
						bible_id = "17"
						//bible_com_text = "est"
					case "伯":
						bible_id = "18"
						//bible_com_text = "job"
					case "詩":
						bible_id = "19"
						//bible_com_text = "psa"
					case "箴":
						bible_id = "20"
						//bible_com_text = "pro"
					case "傳":
						bible_id = "21"
						//bible_com_text = "ecc"
					case "歌":
						bible_id = "22"
						//bible_com_text = "sng"
					case "賽":
						bible_id = "23"
						//bible_com_text = "isa"
					case "耶":
						bible_id = "24"
						//bible_com_text = "jer"
					case "哀":
						bible_id = "25"
						//bible_com_text = "lam" //---------------------------
					case "結":
						bible_id = "26"
						//bible_com_text = "eze"
					case "但":
						bible_id = "27"
						//bible_com_text = "gen"
					case "何":
						bible_id = "28"
						//bible_com_text = "gen"
					case "珥":
						bible_id = "29"
						//bible_com_text = "gen"
					case "摩":
						bible_id = "30"
						//bible_com_text = "gen"
					case "俄":
						bible_id = "31"
						//bible_com_text = "gen"
					case "拿":
						bible_id = "32"
						//bible_com_text = "gen"
					case "彌":
						bible_id = "33"
						//bible_com_text = "gen"
					case "鴻":
						bible_id = "34"
						//bible_com_text = "gen"
					case "哈":
						bible_id = "35"
						//bible_com_text = "gen"
					case "番":
						bible_id = "36"
						//bible_com_text = "gen"
					case "該":
						bible_id = "37"
						//bible_com_text = "gen"
					case "亞":
						bible_id = "38"
						//bible_com_text = "gen"
					case "瑪":
						bible_id = "39"
						//bible_com_text = "gen"
					case "太":
						//bible_com_text = "gen"
					case "可":
						bible_id = "41"
						//bible_com_text = "gen"
					case "路":
						bible_id = "42"
						//bible_com_text = "gen"
					case "約":
						bible_id = "43"
						//bible_com_text = "gen"
					case "徒":
						bible_id = "44"
						//bible_com_text = "gen"
					case "羅":
						bible_id = "45"
						//bible_com_text = "gen"
					case "林前":
						bible_id = "46"
						//bible_com_text = "gen"
					case "林後":
						bible_id = "47"
						//bible_com_text = "gen"
					case "加":
						bible_id = "48"
						//bible_com_text = "gen"
					case "弗":
						bible_id = "49"
						//bible_com_text = "gen"
					case "腓":
						bible_id = "50"
						//bible_com_text = "gen"
					case "西":
						bible_id = "51"
						//bible_com_text = "gen"
					case "帖前":
						bible_id = "52"
						//bible_com_text = "gen"
					case "帖後":
						bible_id = "53"
						//bible_com_text = "gen"
					case "提前":
						bible_id = "54"
						//bible_com_text = "gen"
					case "提後":
						bible_id = "55"
						//bible_com_text = "gen"
					case "多":
						bible_id = "56"
						//bible_com_text = "gen"
					case "門":
						bible_id = "57"
						//bible_com_text = "gen"
					case "來":
						bible_id = "58"
						//bible_com_text = "gen"
					case "雅":
						bible_id = "59"
						//bible_com_text = "gen"
					case "彼前":
						bible_id = "60"
						//bible_com_text = "gen"
					case "彼後":
						bible_id = "61"
						//bible_com_text = "gen"
					case "約一":
						bible_id = "62"
						//bible_com_text = "gen"
					case "約二":
						bible_id = "63"
						//bible_com_text = "gen"
					case "約三":
						bible_id = "64"
						//bible_com_text = "gen"
					case "猶":
						bible_id = "65"
						//bible_com_text = "gen"
					case "啟":
						bible_id = "66"
						//bible_com_text = "gen"
				}

				log.Print("根據 bible() function 匹配到的回應內容：" + bot_msg)
				log.Print("分析所得的章 = bible_chap = " + bible_chap)
				log.Print("分析所得的節 = bible_sec = " + bible_sec)
				
								//增加到這
					//				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.ID+":"+message.Text+" OK!")).Do(); err != nil {
					// 				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(bot_msg)).Do(); err != nil {
					// 					log.Print(err)
					// 				}
								//https://devdocs.line.me/en/?go#send-message-object
				

				//沒辦法建立 anime function 直接在裡面操作（因為用途不一樣當然不能）。 只好先用加法，從下游進行正則分析處理 reg  //https://play.golang.org/p/cjO5La2cKR
					//bible() 只是負責處理字串，理所當然裡面無法做任何的發言動作。（除非把可以發言的相關物件傳進去？）
				// reg := regexp.MustCompile("^.*(有喔！有喔！你在找這個對吧！？)\\n(https?.*)(\\n*.*)$")
				// log.Print("--抓取［" + bot_msg + "］分析觀察--")
				// log.Print("anime 後的 1 = " + reg.ReplaceAllString(bot_msg, "$1"))
				// log.Print("anime 後的 2 = " + reg.ReplaceAllString(bot_msg, "$2")) //URL
				// log.Print("完結篇廢話 = 3 = " + reg.ReplaceAllString(bot_msg, "$3")) //完結篇的廢話




				// //anime url get //2016.12.22+
				// anime_url := reg.ReplaceAllString(bot_msg, "$2")

				// //判斷得到的 $2 是不是 http 開頭字串
				// reg_http := regexp.MustCompile("^(http)s?.*") 

				// if reg_http.ReplaceAllString(anime_url,"$1") != "http"{
				// 	log.Print("anime_url = " + anime_url)
				// 	anime_url = ""
				// }

				//判斷是不是找不到
				//reg_nofind := regexp.MustCompile("^你是要找.*\\n.*\\n.*\\n.*\\n.*\\n.*(才會增加比較慢XD）)$") 

				//這是從字串結果來判斷的方式，但發現有其他方式判斷（直接 bot_msg==開發者）所以這個暫時不用				
				//reg_loking_for_admin := regexp.MustCompile("^(你找我主人？OK！).*\\n.*\\n.*\\n.*\\n.*\\n.*\\n.*\\n.*\\n.*\\n.*\\n.*\\n.*\\n.*") 
					//reg_loking_for_admin := regexp.MustCompile("^(你找我主人？OK！).*") 

				//2016.12.26:這裡的 bot_msg 已經是下游，經過 bible() 處理過了，沒有匹配的發言內容都會被濾掉。


				reg_nofind := regexp.MustCompile("^.*\\n.*對不起，(我還沒學呢...)\\n$")
				
				log.Print("--抓取分析觀察--")
				log.Print("找不到的 $1 = " + reg_nofind.ReplaceAllString(bot_msg, "$1"))
				log.Print("判斷是不是沒有匹配到內容（true = 沒找到）= ")
				log.Print(reg_nofind.ReplaceAllString(bot_msg, "$1")=="我還沒學呢...")
				
				if bot_msg != ""{
					//2016.12.20+ for test	
					switch bot_msg{
						case "c38b3100b02ef42411a99b7975e4ff47":
							// if username == "LL" {
							// 	if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("登入成功！")).Do(); err != nil {
							// 		log.Print(2162)
							// 		log.Print(err)
							// 	}
							// }
							switch username{
							case "LL":
								if target_id_code == "U6f738a70b63c5900aa2c0cbbe0af91c4"{
									// imageURL = SystemImageURL
									template := linebot.NewConfirmTemplate(
										"驗證成功！\n要現在進入管理介面嗎？",
										linebot.NewPostbackTemplateAction("是","開啟管理者選單", ""),
										linebot.NewPostbackTemplateAction("否","登出管理者", ""),
									)
									obj_message := linebot.NewTemplateMessage("這功能只支援最新版本 APP 使用。\n請用 APP 端查看下一步。", template)
									if _, err = bot.ReplyMessage(event.ReplyToken,obj_message).Do(); err != nil {
										log.Print(2162)
										log.Print(err)
									}
								}
							}
							return
						case "測試":
							if target_id_code == "U6f738a70b63c5900aa2c0cbbe0af91c4"{
								imageURL = SystemImageURL
								LineTemplate_test := linebot.NewCarouselTemplate(
									linebot.NewCarouselColumn(
										imageURL, "test", "For test mode.",
										linebot.NewPostbackTemplateAction("測試","test", ""),
										linebot.NewPostbackTemplateAction("管理模式","admin", ""),
										linebot.NewPostbackTemplateAction("申請使用管理者","開發者", "開發者"),
									),
									// LineTemplate_other_example,
									// LineTemplate_other,
									//LineTemplate_CarouselColumn_feedback,
								)
								no_temp_msg := "你已觸發測試模式，請更新最新版本的 LINE 查看內容 。"
								obj_message := linebot.NewTemplateMessage(no_temp_msg, LineTemplate_test)
								if _, err = bot.ReplyMessage(event.ReplyToken, obj_message).Do(); err != nil {
										log.Print(847)
										log.Print(err)
								}
							}
							return
						case "地圖":
							imageURL = "https://trello-attachments.s3.amazonaws.com/52ff05f27a3c676c046c37f9/585e3fb981c1240b4df88d73/34fba56ed5cbb5d7f0a9d7d2543ff238/C02018kong-koan_8481.JPG"
							LineTemplate_test := linebot.NewCarouselTemplate(
								linebot.NewCarouselColumn(
									imageURL, "找教會？我們在這裡！", "也能找其他教會",
									linebot.NewURITemplateAction("Google Map 公館教會","https://goo.gl/maps/h6s5ccdXrL52"),
									linebot.NewURITemplateAction("最推薦的教會地圖系統","https://church.oursweb.net/lite/"),
									linebot.NewURITemplateAction("長老會查詢系統","http://www.pct.org.tw/look4church.aspx"),
								),
								// LineTemplate_other_example,
								// LineTemplate_other,
								linebot.NewCarouselColumn(
									imageURL, "其他功能", "各種這間教會的資訊",
									linebot.NewMessageTemplateAction("週報", "週報"),
									linebot.NewMessageTemplateAction("聯絡資訊", "聯絡資訊"),
									linebot.NewMessageTemplateAction("網站資訊", "網站資訊"),
								),
								LineTemplate_CarouselColumn_feedback,
							)
							no_temp_msg := "如果你要尋找身邊附近的教會，\n推薦使用這個地圖系統尋找：\nhttps://church.oursweb.net/lite/\n\n如果要找我們教會就在這裡！"
							obj_message := linebot.NewTemplateMessage(no_temp_msg, LineTemplate_test)

							obj_message_map := linebot.NewLocationMessage("台北公館教會", "11677 台北市汀州路四段 85 巷 2 號", 25.007408,121.537688) //台北市信義區富陽街46號
							if _, err = bot.ReplyMessage(event.ReplyToken, obj_message, obj_message_map).Do(); err != nil {
								log.Print(1876)
								log.Print(err)
							}
							return
						case "開發者":
							if target_id_code == notalk{
								if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("感謝您呼叫我！\n但為了不要過度打擾群組對話，在此群組暫時不支援此功能。\n請移步到 http://line.me/R/ti/g/X8_PRxyGL7 對話群組進行查詢。\n或個人+我（https://line.me/R/ti/p/@rgp6918s）好友，進行個人查詢。")).Do(); err != nil {
										log.Print(24483)
										log.Print(err)
								}
								return
							}
							bot_msg = "你找我的製造者？OK！\n我跟你講我的夥伴喵在哪，你去加他。\n他跟製造者很親近的，跟他說的話製造者都會看到。\nhttps://line.me/R/ti/p/%40uwk0684z\n\n\n你也可以從下面這個連結直接去找開發者線上對話。\n\n如果他不在線上一樣可以留言給他，\n他會收到的！\n這跟手機、電腦桌面軟體都有同步連線。" +
							"\n\nhttp://www.smartsuppchat.com/widget?key=77b943aeaffa11a51bb483a816f552c70e322417&vid=" + target_id_code +
							"&lang=tw&pageTitle=%E9%80%99%E6%98%AF%E4%BE%86%E8%87%AA%20LINE%40%20%E9%80%B2%E4%BE%86%E7%9A%84%E5%8D%B3%E6%99%82%E9%80%9A%E8%A8%8A"
							log.Print("觸發找製造者")
							template := linebot.NewCarouselTemplate(
								linebot.NewCarouselColumn(
									SystemImageURL, "開發者相關資訊", "你可以透過此功能\n聯絡 開發者",
									LineTemplate_addme,
									LineTemplate_chat,
									linebot.NewPostbackTemplateAction("聯絡 LINE 機器人開發者", "開發者", "開發者"),
								),
							)
							obj_message := linebot.NewTemplateMessage("上面這些都是聯絡開發者的相關方法。", template)
							if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(bot_msg),obj_message).Do(); err != nil {
								log.Print(1672)
								log.Print(err)
							}
							//HttpPost_JANDI(target_item + " [" + user_talk + "](" + userImageUrl + ")：" + message.Text + `\n` + userStatus, "yellow" , "LINE 同步：執行找開發者",target_id_code)
							//HttpPost_IFTTT(target_item + " " + user_talk + "：" + message.Text + `\n<br>` + userImageUrl + `\n<br>` + userStatus, "LINE 同步：執行找開發者",target_id_code)
							HttpPost_Zapier(target_item + " [" + user_talk + "](" + userImageUrl + ")：" + message.Text + `\n` + userStatus, "LINE 同步：執行找開發者",target_id_code,user_talk)
							return
						case "GOTEST":
							//簡單說模板有三種（Y/N[1~2動]、Bottons[最多4個動作]、carousel[3個動作 && 並排最多五個(每個動作數量要一致)]），動作也有三種（操作使用者發言、POST兼使用者發言(使用者發言可為空)、URI 可連網址或 tel: 等協定）
								//bot_msg = "HI～ 我最近很喜歡看巴哈姆特動畫瘋。\nhttp://ani.gamer.com.tw/\n\n你也可以問我動畫，我可以帶你去看！\n要問我動畫的話可以這樣問：\n動畫 動畫名稱 集數\n\n例如：\n動畫 美術社 12\nアニメ 美術社大有問題 12\nanime 美術社 １\n巴哈姆特 美術社 12\n以上這些都可以\n\n但中間要用空白或冒號、分號隔開喔！\n不然我會看不懂 ＞A＜\n\nPS：目前這只提供查詢動畫的功能。\n如有其他建議或想討論，請對我輸入「開發者」進行聯絡。"
								//bot_msg = "有喔！有喔！你在找這個對吧！？\n" + "https://ani.gamer.com.tw/animeVideo.php?sn=5863" + "\n\n等等！這是最後一話！？"

								//2016.12.22+ free POST
								//func //HttpPost_JANDI(body, connectColor, title, --url--) error  
								//http://nipponcolors.com/#matsuba
								// //HttpPost_JANDI("test for LINE BOT", "#42602D" , "test")
								////HttpPost_IFTTT("test for line bot", "純測試",target_id_code) //2016.12.22+ 成功！！！
								//HttpPost_LINE_notify("test")
								
								// "http://ani.gamer.com.tw/animeVideo.php?sn=6878",
								//  第？話",
								//  "https://p2.bahamut.com.tw/B/2KU/33/0001485933.PNG",
								//  "查詢結果",
								//  "動畫名稱 ",
								// bot_msg 

								//log.Print("完結篇廢話 = 3 = " + reg.ReplaceAllString(bot_msg, "$3")) //完結篇的廢話

								//Create message
								//https://github.com/line/line-bot-sdk-go
								//https://github.com/line/line-bot-sdk-go/blob/master/linebot/message.go

								//模板成功  //官方範例 https://github.com/line/line-bot-sdk-go/blob/master/examples/kitchensink/server.go
									//linebot.NewTemplateMessage
									// //1 confirm 純是否類型的問法
									// //.NewConfirmTemplate 模板，裡面最多只能有兩個動作，按鈕只能左右
									// //.NewMessageTemplateAction 發言動作

									// template := linebot.NewConfirmTemplate(
									// 	"Do it?",
									// 	linebot.NewMessageTemplateAction("Yes", "Yes!"),
									// 	linebot.NewMessageTemplateAction("No", "No!"),
									// )

			 					//     leftBtn := linebot.NewMessageTemplateAction("left", "left clicked")// 後面的參數 "left clicked" = 在使用者按下後，自動幫使用者發訊息
			 					//     rightBtn := linebot.NewMessageTemplateAction("right", "right clicked")// 後面的參數 "right clicked" = 在使用者按下後，自動幫使用者發訊息
								 //    //.NewMessageTemplateAction("字面按鈕", "設定讓使用者按下後發送內容") 會讓使用者發送那樣的內容給系統
			 					//     template := linebot.NewConfirmTemplate("Hello World", leftBtn, rightBtn)
			 					//     //.NewConfirmTemplate

								//linebot.NewTemplateMessage
		 					    //2 buttons
		 					    //.NewButtonsTemplate 模板，裡面最多只能有四個動作
		 					    //.NewURITemplateAction 開啟指定網址的動作
		 					    //.NewPostbackTemplateAction ？？動作
		 					    //						第二參數可以讓她　ＰＯＳＴ指定內容（但還不會處理．．．）	第三參數類似於 .NewMessageTemplateAction 的效果
			 					//     imageURL := "https://images.gamme.com.tw/news2/2016/51/39/paCYoqCXkqSarqSZ.jpg"
									// template := linebot.NewButtonsTemplate(
									// 	imageURL, "你好歡迎光臨", "這是內文",							//這前三個 分別是圖片(必須https)、標題、內文
									// 	linebot.NewURITemplateAction("來我的網站", "https://synr.github.io"),
									// 	linebot.NewPostbackTemplateAction("目錄查詢", "目錄", "目錄"),
									// 	linebot.NewPostbackTemplateAction("開發者", "開發者", "開發者"),
									// 	linebot.NewMessageTemplateAction("Say message", "Rice=米"),
									// )

									//linebot.NewTemplateMessage
									//3 carousel .NewCarouselTemplate  最多可以並排五個「.NewCarouselColumn」的樣板，
									//「.NewCarouselColumn」裡面最多只能有三個動作按鈕，但並列的其他項目也要一致數量才能。2016.12.22+
									//圖片可以是 PNG
									// imageURL := "https://images.gamme.com.tw/news2/2016/51/39/paCYoqCXkqSarqSZ.jpg"
									// template := linebot.NewCarouselTemplate(
									// 	linebot.NewCarouselColumn(
									// 		"https://p2.bahamut.com.tw/B/2KU/33/0001485933.PNG", "hoge", "fuga",
									// 		linebot.NewURITemplateAction("測試看動畫", "http://ani.gamer.com.tw/animeVideo.php?sn=6878"),
									// 		linebot.NewPostbackTemplateAction("Say hello1", "hello こんにちは", ""),
									// 		linebot.NewURITemplateAction("Go to line.me", "https://line.me"),
									// 	),
									// 	linebot.NewCarouselColumn(
									// 		"https://p2.bahamut.com.tw/B/2KU/18/0001484818.PNG", "hoge", "fuga",
									// 		linebot.NewURITemplateAction("Go to line.me", "https://line.me"),
									// 		linebot.NewPostbackTemplateAction("言 hello2", "hello こんにちは", "hello こんにちは"),
									// 		linebot.NewMessageTemplateAction("Say message", "Rice=米"),
									// 	),
									// 	linebot.NewCarouselColumn(
									// 		imageURL, "hoge", "fuga",
									// 		linebot.NewPostbackTemplateAction("言 hello2", "hello こんにちは", "hello こんにちは"),
									// 		linebot.NewURITemplateAction("Go to line.me", "https://line.me"),
									// 		linebot.NewURITemplateAction("Go to line.me", "https://line.me"),
									// 	),
									// 	linebot.NewCarouselColumn(
									// 		imageURL, "hoge", "fuga",
									// 		linebot.NewURITemplateAction("Go to line.me", "https://line.me"),
									// 		linebot.NewPostbackTemplateAction("言 hello2", "hello こんにちは", "hello こんにちは"),
									// 		linebot.NewMessageTemplateAction("Say message", "Rice=米"),
									// 	),
									// 	linebot.NewCarouselColumn(
									// 		imageURL, "hoge", "fuga",
									// 		linebot.NewURITemplateAction("Go to line.me", "https://line.me"),
									// 		linebot.NewPostbackTemplateAction("言 hello2", "hello こんにちは", "hello こんにちは"),
									// 		linebot.NewURITemplateAction("Go to line.me", "https://line.me"),
									// 	),
									// )
									//所以有三種樣板，有三種動作按鈕。兩個樣板可以放圖片，一個單純只能兩個按鈕。


			 					    //obj_message := linebot.NewTemplateMessage("HI～ 我最近很喜歡看巴哈姆特動畫瘋。\nhttp://ani.gamer.com.tw/\n\n你也可以問我動畫，我可以帶你去看！\n要問我動畫的話可以這樣問：\n動畫 動畫名稱 集數\n\n例如：\n動畫 美術社 12\nアニメ 美術社大有問題 12\nanime 美術社 １\n巴哈姆特 美術社 12\n以上這些都可以\n\n但中間要用空白或冒號、分號隔開喔！\n不然我會看不懂 ＞A＜\n\nPS：目前這只提供查詢動畫的功能。\n如有其他建議或想討論，請對我輸入「開發者」進行聯絡。", template)//messgage := linebot.NewTemplateMessage("請使用更新 APP 或使用手機 APP 才能看到這個功能。", template)
									//obj_message := linebot.NewTemplateMessage(bot_msg, template)
			 					    //.NewTemplateMessage("無法支援按鈕模式時要發出的訊息",Template 物件)

										// 						if _, err = bot.ReplyMessage(event.ReplyToken, message).Do(); err != nil {
										// 							log.Print(err)
										// 						}


									//https://devdocs.line.me/en/?go#send-message-object


								//++ https://github.com/dongri/line-bot-sdk-go KEY:linebot.NewImageMessage

								//.NewImageMessage 發圖片成功
								//originalContentURL := "https://avatars0.githubusercontent.com/u/5731891?v=3&s=96"
		    					//previewImageURL := "https://avatars0.githubusercontent.com/u/5731891?v=3&s=96"
		    					//obj_message := linebot.NewImageMessage(originalContentURL, previewImageURL)


								//.NewStickerMessage 發貼貼圖成功	 //https://devdocs.line.me/files/sticker_list.pdf					
								//obj_message := linebot.NewStickerMessage("1", "1") //https://devdocs.line.me/en/?go#send-message-object

								//這是個謎
								//https://devdocs.line.me/en/?go#imagemap-message
								//https://github.com/line/line-bot-sdk-go/blob/master/examples/kitchensink/server.go
								// obj_message := linebot.NewImagemapMessage(
								// 	"https://synr.github.io/img/index.jpg",
								// 	"Imagemap alt text",
								// 	linebot.ImagemapBaseSize{1040, 1040},
								// 	linebot.NewURIImagemapAction("https://store.line.me/family/manga/en", linebot.ImagemapArea{0, 0, 520, 520}),
								// 	linebot.NewURIImagemapAction("https://store.line.me/family/music/en", linebot.ImagemapArea{520, 0, 520, 520}),
								// 	linebot.NewURIImagemapAction("https://store.line.me/family/play/en", linebot.ImagemapArea{0, 520, 520, 520}),
								// 	linebot.NewMessageImagemapAction("URANAI!", linebot.ImagemapArea{520, 520, 520, 520}),
								// )
								//func NewImagemapMessage
								//https://github.com/line/line-bot-sdk-go/blob/master/linebot/message.go > Actions:  actions
								//看起來好像可以有動作

								//Audio //https://github.com/dongri/line-bot-sdk-go
							    // originalContentURL := "https://dl.dropboxusercontent.com/u/358152/linebot/resource/ok.m4a"
							    // duration := 1000
							    // obj_message := linebot.NewAudioMessage(originalContentURL, duration)

		 					    //接收各種 message object
								//if _, err = bot.ReplyMessage(event.ReplyToken, obj_message,obj_message,obj_message,obj_message,obj_message).Do(); err != nil { //五聯發
								// if _, err = bot.ReplyMessage(event.ReplyToken, obj_message).Do(); err != nil { 
								//if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewStickerMessage("1", "1"),linebot.NewStickerMessage("1", "2"),linebot.NewStickerMessage("2", "19"),linebot.NewStickerMessage("2", "20"),linebot.NewStickerMessage("1", "3")).Do(); err != nil {
								// 	log.Print(err)
								// }
							return
						case "行事曆":
						    imageURL = "https://trello-attachments.s3.amazonaws.com/52ff05f27a3c676c046c37f9/585e3fb981c1240b4df88d73/34fba56ed5cbb5d7f0a9d7d2543ff238/C02018kong-koan_8481.JPG"
							template := linebot.NewCarouselTemplate(
								linebot.NewCarouselColumn(
									imageURL, "教會行事曆", "預覽教會行事曆",
									linebot.NewURITemplateAction("線上行事曆","https://calendar.kkcpct.org"),
									linebot.NewURITemplateAction("本季行事表","https://calendar.kkcpct.org/table"),
									linebot.NewURITemplateAction("訂閱月曆 Feeds ics","https://calendar.kkcpct.org/ics"),
									//linebot.NewPostbackTemplateAction("本週週報", "週報", ""),
									//linebot.NewPostbackTemplateAction("聚會時間", "聚會時間", ""),
									// linebot.NewPostbackTemplateAction("其他本週公告", "其他本週公告", ""),
								),
								linebot.NewCarouselColumn(
									imageURL, "其他功能", "各種這間教會的資訊",
									linebot.NewMessageTemplateAction("聯絡資訊", "聯絡資訊"),
									linebot.NewMessageTemplateAction("交通資訊", "教會地圖"),
									linebot.NewMessageTemplateAction("網站資訊", "網站資訊"),
								),
								LineTemplate_CarouselColumn_feedback,
							)
							t_msg :="本季行事曆表：https://goo.gl/2V5sbN\n\n"+
									"教會 Google 行事曆（按右下角可以加到你的帳號）：\nhttps://calendar.kkcpct.org\nhttps://goo.gl/ufFxB2\n\n" + 
									"Google 月曆 ics 檔：\nhttps://calendar.kkcpct.org/ics\n\n" +
									"iCalendar Feeds 使用方法請參照：https://goo.gl/KNfQx0"
							obj_message := linebot.NewTemplateMessage(t_msg, template)
							if _, err = bot.ReplyMessage(event.ReplyToken, obj_message).Do(); err != nil {
									log.Print(1630)
									log.Print(err)
							}
							return
						case "聚會時間":
						    imageURL = "https://trello-attachments.s3.amazonaws.com/52ff05f27a3c676c046c37f9/585e3fb981c1240b4df88d73/34fba56ed5cbb5d7f0a9d7d2543ff238/C02018kong-koan_8481.JPG"
							template := linebot.NewCarouselTemplate(
								linebot.NewCarouselColumn(
									imageURL, "主日聚會時間", "上午 09:00（台語）\n上午 11:00（華語）",
									linebot.NewPostbackTemplateAction("本週週報", "週報", ""),
									// linebot.NewPostbackTemplateAction("聚會時間", "聚會時間", ""),
									linebot.NewMessageTemplateAction("本季行事曆","行事曆"),
									linebot.NewPostbackTemplateAction("家庭禮拜", "家庭禮拜", ""),
								),
								linebot.NewCarouselColumn(
									imageURL, "其他功能", "各種這間教會的資訊",
									linebot.NewMessageTemplateAction("聯絡資訊", "聯絡資訊"),
									linebot.NewMessageTemplateAction("交通資訊", "教會地圖"),
									linebot.NewMessageTemplateAction("網站資訊", "網站資訊"),
								),
								LineTemplate_CarouselColumn_feedback,
							)
							t_msg := weektime_msg
							obj_message := linebot.NewTemplateMessage(t_msg, template)
							if _, err = bot.ReplyMessage(event.ReplyToken, obj_message).Do(); err != nil {
									log.Print(1630)
									log.Print(err)
							}
							return
						case "圖書查詢":
							if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("你想找教會相關的圖書是嗎？\n\n這裡推薦使用以下綜合查詢系統查詢各教會相關單位館藏！\nhttp://ttlib.fhl.net")).Do(); err != nil {
									log.Print(15083)
									log.Print(err)
							}
							return
						case "查詢可用簡寫":
							if target_id_code == notalk{
								if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("感謝您呼叫我！\n但為了不要過度打擾群組對話，在此群組暫時不支援此功能。\n請移步到 http://line.me/R/ti/g/X8_PRxyGL7 對話群組進行查詢。\n或個人+我（https://line.me/R/ti/p/@rgp6918s）好友，進行個人查詢。")).Do(); err != nil {
										log.Print(24483)
										log.Print(err)
								}
								return
							}
							log.Print("有走進 查詢可用簡寫")
						    imageURL = "https://trello-attachments.s3.amazonaws.com/52ff05f27a3c676c046c37f9/585e3fb981c1240b4df88d73/34fba56ed5cbb5d7f0a9d7d2543ff238/C02018kong-koan_8481.JPG"
							template := linebot.NewCarouselTemplate(
								linebot.NewCarouselColumn(
									imageURL, "示範短寫查法", "示範如何簡短查聖經",
									linebot.NewMessageTemplateAction("聖經 創 1:1", "聖經 創 1:1"),
									linebot.NewMessageTemplateAction("聖經 馬太 1:1-20", "聖經 馬太 1:1-20"),
									linebot.NewMessageTemplateAction("聖經 詩 1:1", "聖經 詩 1:1"),
								),
								linebot.NewCarouselColumn(
									imageURL, "其他功能", "各種這間教會的資訊",
									linebot.NewMessageTemplateAction("聯絡資訊", "聯絡資訊"),
									linebot.NewMessageTemplateAction("交通資訊", "教會地圖"),
									linebot.NewMessageTemplateAction("網站資訊", "網站資訊"),
								),
								LineTemplate_CarouselColumn_feedback,
							)
							s_text_list := `創世記 = Gen = Genesis = 創 = 創世 = 創世紀 = 創世記 = Ge = ge = gen = 창세기
出埃及記 = ex = Ex = Exodus = 埃及 = 出 = 出埃及 = 出埃及記 = 출애굽기 = エジプト = 出エジプト = 出エジプト記
利未記 = Lev = Leviticus = 利 = 利未 = 利未記 = Le = le = レビ記 = レビ = 레위기
民數記 = Num = Numbers = 民 = 民數 = 民數記 = Nu = nu = 민수기 = 民数 = 民数記
申 = 申命記 = Deut = Deuteronomy = 申 = 申命 = 申命記 = De = de = 신명기
書 = 約書亞記 = Josh = Joshua = 約書亞 = 約書亞記 = Jos = jos = 여호수아 = ヨシュア記 = ヨシュア
士 = 士師記 = Judg = Judges = 士 = 士師 = 士師記 = Jud = jud = jdg = Jdg = 사사기
得 = 路得記 = Ruth = 路得 = 路得記 = Ru = ru = Rut = rut = 룻기 = ルツ = ルツ記
撒上 = 撒母耳記上 = 1 Sam = First Samuel = 撒上 = 撒母耳記上 = 1Sa = 1sa = サムエル記上 = サムエル上 = サム上 = 사무엘상
撒下 = 撒母耳記下 = 2 Sam = Second Samuel = 撒下 = 撒母耳記下 = 2Sa = 2sa = 사무엘하 = サムエル記下 = サムエル下 = サム下

王上 = 列王紀上 = 1 Kin = First Kings = 王上 = 列王上 = 列王紀上 = 列王記上 = 1Ki = 1ki = 열왕기상
王下 = 列王紀下 = 2 Kin = Second Kings = 王下 = 列王下 = 列王記下 = 列王紀下 = 2Ki = 2ki = 열왕기하
代上 = 歷代志上 = 1 Chr = First Chronicles = 歷上 = 代上 = 歷代志上 = 歷代上 = 1Ch = 1ch = 歴上 = 歴代上 = 歴代志上 = 역대상
代下 = 歷代志下 = 2 Chr = Second Chronicles = 代下 = 歷下 = 歷代下 = 歷代志下 = 2Ch = 2ch = 역대하 = 歴代志下 = 歴代下 = 歴下
拉 = 以斯拉記 = Ezra = 拉 = 以斯拉 = 以斯拉記 = Ezr = ezr = エズラ = エズラ記 = 에스라
尼 = 尼希米記 = Neh = Nehemiah = 尼 = 尼希米 = 尼希米記 = Ne = ne = 느헤미야 = ネヘミヤ書 = ネヘミヤ
斯 = 以斯帖記 = Esth = Esther = 斯 = 以斯帖 = 以斯帖記 = Es = est = Есфирь = Ê-xơ-tê = エステル = エステル記 = 에스더
伯 = 約伯記 = Job = job = 伯 = 約伯 = 約伯記 = Книга Иова = Gióp = ヨブ = ヨブ記 = 욥기
詩 = 詩篇 = Ps = Psalms = 詩 = 詩篇 = ps = 시편 = Thi-thiên = Псалтирь
箴 = 箴言 = Prov = Proverbs = 箴 = 箴言 = Pr = pr = Притчи Соломона = Châm-ngôn = 잠언

傳 = 傳道書 = Eccl = Ecclesiastes = 傳 = 傳道 = 傳道書 = Ec = ec = Книга Екклезиаста = или Проповедника = Truyền-đạo = 伝道の書 = 伝道 = 伝 = 伝道書 = 전도서
歌 = 雅歌 = Song = Song of Solomon = 歌 = 雅歌 = So = so = sng = Sng = Песнь песней Соломона = Nhã-ca = 아가
賽 = 以賽亞書 = Is = Isaiah = 賽 = 以賽 = 以賽亞 = 以賽亞書 = Isa = isa = Книга пророка Исаии = Ê-sai = イザヤ書 = イザヤ = 이사야
耶 = 耶利米書 = Jer = Jeremiah = 耶 = 耶利米 = 耶利米書 = jer = 예레미야 = エレミヤ = エレミヤ書 = Giê-rê-mi = Книга пророка Иеремии
哀 = 耶利米哀歌 = Lam = Lamentations = 哀 = 哀歌 = 耶利米哀歌 = La = lam = 예레미야애가 = Ca-thương = Плач Иеремии
結 = 以西結書 = Ezek = Ezekiel = 結 = 以西結 = 以西結書 = Eze = eze = 에스겔 = エゼキエル書 = エゼキエル = Ê-xê-chi-ên = Книга пророка Иезекииля
但 = 但以理書 = Dan = Daniel = 但 = 但以理 = 但以理書 = Da = da = Книга пророка Даниила = Đa-ni-ên = ダニエル書 = ダニエル = 다니엘
何 = 何西阿書 = Hos = Hosea = 何 = 何西 = 何西阿 = 何西阿書 = Ho = ho = Книга пророка Осии = Ô-sê = ホセア書 = ホセア = 호세아
珥 = 約珥書 = Joel = 珥 = 約珥 = 約珥書 = Joe = joe = Книга пророка Иоиля = Giô-ên = ヨエル書 = ヨエル = 요엘
摩 = 阿摩司書 = Amos = 摩 = 阿摩司書 = Am = am = Книга пророка Амоса = A-mốt = アモス書 = アモス = 아모스

俄 = 俄巴底亞書 = Obad = Obadiah = 俄 = 俄巴底亞 = 俄巴底亞書 = Ob = ob = 오바댜 = オバデヤ書 = オバデヤ = Áp-đia = Книга пророка Авдия
拿 = 約拿書 = Jon = Jonah = 拿 = 約拿 = 約拿書 = jon = 요나 = ヨナ書 = ヨナ = Giô-na = Книга пророка Ионы
彌 = 彌迦書 = Micah = 彌 = 彌迦 = 彌迦書 = Mic = mic = Книга пророка Михея = Mi-chê = ミカ書 = ミカ = 미가
鴻 = 那鴻書 = Nah = Nahum = 鴻 = 那鴻 = 那鴻書 = Na = na = Книга пророка Наума = Na-hum = ナホム書 = ナホム = 나훔
哈 = 哈巴谷書 = Habakkuk = 哈 = 哈巴 = 哈巴谷 = 哈巴谷書 = Hab = hab = Книга пророка Аввакума = Ha-ba-cúc = ハバクク書 = ハバクク = ハバ = クク = ハバ書 = 하박국
番 = 西番雅書 = Zeph = Zephaniah = 番 = 西番雅 = 西番雅書 = Zep = zep = 스바냐 = ゼパニヤ書 = ゼパニヤ = Sô-phô-ni = Книга пророка Софонии
該 = 哈該書 = Haggai = 該 = 哈該 = 哈該書 = Hag = hag = 학개 = ハガイ書 = ハガイ = A-ghê = Книга пророка Аггея
亞 = 撒迦利亞書 = Zech = Zechariah = 亞 = 撒迦利亞 = 撒迦利亞書 = Zec = zec = Книга пророка Захарии = Xa-cha-ri = 스가랴 = ゼカリヤ書 = ゼカリヤ
瑪 = 瑪拉基書 = Malachi = 瑪 =  = 瑪拉 = 瑪拉基 = 瑪拉基書 = Mal = mal = 말라기 = マラキ書 = マラキ = Ma-la-chi = Книга пророка Малахии
太 = 馬太福音 = Matt = Matthew = 太 = 馬太 = 馬太福音 = Mt = mt = 마태복음 = マタイによる福音書 = マタイ = マタイによる = Ma-thi-ơ = От Матфея святое благовествование

可 = 馬可福音 = Mark = 可 = 馬可 = 馬可福音 = Mr = mr = マルコによる福音書 = マルコ = マルコによる = 마가복음 = Mác = От Марка святое благовествование
路 = 路加福音 = Luke = 路 = 路加 = 路加福音 = Lu = lu = От Луки святое благовествование = Lu-ca = ルカによる福音書 = ルカ = ルカによる = 누가복음
約 = 約翰福音 = John = 約 = 約翰 = 約翰福音 = Joh = joh = От Иоанна святое благовествование = Giăng = ヨハネによる福音書 = ヨハネ = ヨハネによる = 요한복음
徒 = 使徒行傳 = Acts = 徒 = 使徒 = 使徒行傳 = Ac = ac = Деяния святых апостолов = Công-vụ Các Sứ-đồ = 使徒行伝 = 사도행전
羅 = 羅馬書 = Rom = Romans = 羅 = 羅馬 = 羅馬書 = Ro = ro = Послание к Римлянам = Rô-ma = ローマ = ローマ人への手紙 = 로마서
林前 = 哥林多前書 = 1 Cor = First Corinthians = 林前 = 哥林多前 = 哥林多前書 = 1Co = 1co = Первое послание к Коринфянам = 1 Cô-rinh-tô = コリント人への第一の手紙 = コリント一 = コリント人への第一 = 고린도전서
林後 = 哥林多後書 = 2 Cor = Second Corinthians = 林後 = 哥林多後 = 哥林多後書 = 2Co = 2co = Второе послание к Коринфянам = 2 Cô-rinh-tô = コリント人への第二の手紙 = コリント二 = コリント人への第二の = 고린도후서
加 = 加拉太書 = Gal = Galatians = 加 = 加拉太 = 加拉太書 = Ga = ga = Послание к Галатам = Ga-la-ti = ガラテヤ = ガラテヤ人への手紙 = 갈라디아서
弗 = 以弗所書 = Ephesians = 弗 = 以弗所 = 以弗所書 = Eph = eph = Послание к Ефесянам = Ê-phê-sô = エペソ人への手紙 = エペソ = エペソ人 = エペソ人の手紙 = 에베소서
腓 = 腓立比書 = Phil = Philippians = 腓 = 腓立 = 腓立比 = 腓立比書 = Php = php = 빌립보서 = ピリピ = ピリピ人.ピリピ人への手紙 = Послание к Филиппийцам = Phi-líp

西 = 歌羅西書 = Col = col = Colossians = 西 = 歌羅西 = 歌羅 = 歌羅西書 = Послание к Колоссянам = Cô-lô-se = コロサイ人への手紙 = コロサイ = コロ = 골로새서
帖前 = 帖撒羅尼迦前書 = 1 Thess = First Thessalonians = 帖前 = 帖撒羅尼迦前 = 帖撒羅尼迦前書 = 1Th = 1th = 데살로니가전서 = テサロニケ人への第一の手紙 = テサ一 = テサロニケ一 = 1 Tê-sa-lô-ni-ca = Первое послание к Фессалоникийцам (Солунянам)
帖後 = 帖撒羅尼迦後書 = 2 Thess = Second Thessalonians = 帖後 = 帖撒羅尼迦後 = 帖撒羅尼迦後書 = 2Th = 2th = 데살로니가후서 = テサロニケ人への第二の手紙 = テサ二 = テサロニケ二 = 2 Tê-sa-lô-ni-ca = Второе послание к Фессалоникийцам (Солунянам)
提前 = 提摩太前書 = 1 Tim = First Timothy = 提前 = 提摩太前 = 提摩太前書 = 1Ti = 1ti = Первое послание к Тимофею = 1 Ti-mô-thê = テモテヘの第一の手紙 = テモテ一 = 디모데전서
提後 = 提摩太後書 = 2 Tim = Second Timothy = 提後 = 提摩太後 = 提摩太後書 = 2Ti = 2ti = Второе послание к Тимофею = 2 Ti-mô-thê = テモテヘの第二の手紙 = テモテ二 = 디모데후서
多 = 提多書 = Titus = 多 = 提多 = 提多書 = Tit = tit = Послание к Титу = Tít = テトスヘの手紙 = テトス = 디도서
門 = 腓利門書 = Philem = Philemon = 門 = 腓利 = 腓利門 = 腓利門書 = Phm = phm = Послание к Филимону = Phi-lê-môn = ピレモンヘの手紙 = ピレモン = 빌레몬서
來 = 希伯來書 = Heb = Hebrews = 來 = 希伯來 = 希伯來書 = heb = Послание к Евреям = Hê-bơ-rơ = ヘブル人への手紙 = ヘブル = 히브리서
雅 = 雅各書 = James = 雅 = 雅各 = 雅各書 = Jas = jas = Послание Иакова = Gia-cơ = ヤコブの手紙 = ヤコブ = 야고보서
彼前 = 彼得前書 = 1 Pet = First Peter = 彼前 = 彼得前 = 彼得前書 = 1Pe = 1pe = Первое послание Петра = 1 Phi-e-rơ = ペテロの第一の手紙 = ペテロ一 = 베드로전서

彼後 = 彼得後書 = 2 Pet = Second Peter = 彼後 = 彼得後 = 彼得後書 = 2Pe = 2pe = Второе послание Петра = 2 Phi-e-rơ = ペテロの第二の手紙 = ペテロ = 베드로후서
約一 = 約翰一書 = 1 John = First John = 約一 = 約翰一書 = 約翰1 = 約翰1書 = 1Jo = 1jo = Первое послание Иоанна = 1 Giăng = ヨハネの第一の手紙 = ヨハネ一 = 요한일서
約二 = 約翰二書 = 2 John = second John = 約二 = 約翰二書 = 約翰2 = 約翰2書 = 2Jo = Второе послание Иоанна = 2 Giăng = ヨハネの第二の手紙 = ヨハネ二 = 요한2서
約三 = 約翰三書 = 3 John = Third John = 約三 = 約翰三書 = 約翰3 = 約翰3書 = 3Jo = 3jo = Третье послание Иоанна = 3 Giăng = ヨハネの第三の手紙 = ヨハネ三 = 요한3서
猶 = 猶大書 = Jude = 猶 = 猶大 = 猶大書 = jude = Послание Иуды = Giu-đe = ユダの手紙 = ユダ = 유다서
啟 = 啟示錄 = Rev = Revelation = 啟 = 啟示 = 啟示錄 = Re = re = ｒｅ = Ｒｅ = rev = Откровение ап. Иоанна Богослова (Апокалипсис) = Khải-huyền = ヨハネの黙示録 = 黙示録 = 요한계시록`
							t_msg := "這裡是 APP 專用按鈕體驗短查法區域，請使用最新版本的 LINE APP 進行最佳體驗。"
							obj_message := linebot.NewTemplateMessage(t_msg, template)
							if _, err = bot.ReplyMessage(event.ReplyToken,
								linebot.NewTextMessage(string([]rune(s_text_list)[0:2000])),
								linebot.NewTextMessage(string([]rune(s_text_list)[1990:3990])),
								linebot.NewTextMessage(string([]rune(s_text_list)[3980:5980])),
								linebot.NewTextMessage(string([]rune(s_text_list)[5960:len([]rune(s_text_list))])),
								obj_message,
							).Do(); err != nil {
									log.Print(15185)
									log.Print(err)
							}
							return
						case "週報":
							// if target_id_code == notalk{
							// 	if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("感謝您呼叫我！\n但為了不要過度打擾群組對話，在此群組暫時不支援此功能。\n請移步到 http://line.me/R/ti/g/X8_PRxyGL7 對話群組進行查詢。\n或個人+我（https://line.me/R/ti/p/@rgp6918s）好友，進行個人查詢。")).Do(); err != nil {
							// 			log.Print(24483)
							// 			log.Print(err)
							// 	}
							// 	return
							// }
//							if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("本週教會電子週報從缺\n這只是示範效果可以如何展示"),obj_message_week_2,obj_message_week_3,linebot.NewTextMessage(next_week_msg),obj_message_nextweek_review).Do(); err != nil {
							//if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(this_week_text_1),linebot.NewTextMessage(this_week_text_2),obj_message_nextweek_read,linebot.NewTextMessage(next_week_msg),obj_message_nextweek_review).Do(); err != nil {
							if _, err = bot.ReplyMessage(event.ReplyToken, obj_message_nextweek_review,obj_message_week_1,linebot.NewTextMessage(this_week_text_2)).Do(); err != nil {
									log.Print(7161)
									log.Print(err)
							}
							return
						case "新眼光":
							// if target_id_code == notalk{
							// 	if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("感謝您呼叫我！\n但為了不要過度打擾群組對話，在此群組暫時不支援此功能。\n請移步到 http://line.me/R/ti/g/X8_PRxyGL7 對話群組進行查詢。\n或個人+我（https://line.me/R/ti/p/@rgp6918s）好友，進行個人查詢。")).Do(); err != nil {
							// 			log.Print(24483)
							// 			log.Print(err)
							// 	}
							// 	return
							// }
							if _, err = bot.ReplyMessage(event.ReplyToken, obj_message_nextweek_read).Do(); err != nil {
									log.Print(1443)
									log.Print(err)
							}
							return
						case "下週預告":
							// if target_id_code == notalk{
							// 	if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("感謝您呼叫我！\n但為了不要過度打擾群組對話，在此群組暫時不支援此功能。\n請移步到 http://line.me/R/ti/g/X8_PRxyGL7 對話群組進行查詢。\n或個人+我（https://line.me/R/ti/p/@rgp6918s）好友，進行個人查詢。")).Do(); err != nil {
							// 			log.Print(24483)
							// 			log.Print(err)
							// 	}
							// 	return
							// }
							if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(next_week_msg)).Do(); err != nil {
									log.Print(1443)
									log.Print(err)
							}
							return
						case "聯絡資訊":
						    imageURL = "https://trello-attachments.s3.amazonaws.com/52ff05f27a3c676c046c37f9/585e3fb981c1240b4df88d73/34fba56ed5cbb5d7f0a9d7d2543ff238/C02018kong-koan_8481.JPG"
							template := linebot.NewCarouselTemplate(
								linebot.NewCarouselColumn(
									imageURL, "教會聯絡資訊", "電話：（02）2932-7941\n傳真：（02）2934-5003",
									linebot.NewURITemplateAction("電話：(02)2932-7941", "tel:+886229327941"),
									linebot.NewPostbackTemplateAction("電子郵件", "電子郵件", ""),
									linebot.NewMessageTemplateAction("通訊地址", "教會地圖"),
								),
								linebot.NewCarouselColumn(
									imageURL, "聯絡牧師", "王豐榮 牧師：0933-007-724\n吳慧馨 牧師：0933-007-504\n羅滋嶸 傳道：0912-145-239",
									linebot.NewURITemplateAction("王豐榮 牧師：0933-007-724", "tel:0933007724"),
									linebot.NewURITemplateAction("吳慧馨 牧師：0933-007-504", "tel:0933007504"),
									linebot.NewURITemplateAction("羅滋嶸 傳道：0912-145-239", "tel:0912145239"),
								),
								linebot.NewCarouselColumn(
									imageURL, "其他功能", "各種這間教會的資訊",
									//linebot.NewMessageTemplateAction("聚會時間 & 週報", "週報"),TPEKKC
									linebot.NewURITemplateAction("臉書快傳","https://m.me/TPEKKC"),
									linebot.NewMessageTemplateAction("交通資訊", "教會地圖"),
									linebot.NewMessageTemplateAction("網站資訊", "官方網站"),
								),
								LineTemplate_CarouselColumn_feedback,
							)
							t_msg := "台北公館教會的聯絡資訊：\n\n電話：02-29327941\n傳真：02-29345003\n電子郵件：kkcpct@ms29.hinet.net\n\n聯絡牧者\n王豐榮 牧師：0933-007-724\n吳慧馨 牧師：0933-007-504\n羅滋嶸 傳道：0912-145-239"
							obj_message := linebot.NewTemplateMessage(t_msg, template)
							if _, err = bot.ReplyMessage(event.ReplyToken, obj_message).Do(); err != nil {
									log.Print(1630)
									log.Print(err)
							}
							return
						case "網站資訊":
						    imageURL = "https://trello-attachments.s3.amazonaws.com/52ff05f27a3c676c046c37f9/585e3fb981c1240b4df88d73/34fba56ed5cbb5d7f0a9d7d2543ff238/C02018kong-koan_8481.JPG"
							template := linebot.NewCarouselTemplate(
								linebot.NewCarouselColumn(
									imageURL, "網站資訊", "教會網站資訊",
									linebot.NewURITemplateAction("台北公館教會的 YouTube", "https://m.youtube.com/user/kkcpct"),
									linebot.NewURITemplateAction("台北公館教會的 Facebook", "https://m.facebook.com/TPEKKC"),
									linebot.NewURITemplateAction("台北公館教會的 Blog", "http://blog.kkcpct.org"),
									//linebot.NewURITemplateAction("教會機構資料", "http://www.pct.org.tw/churchdata.aspx?strOrgNo=C02018"), //http://www.kkcpct.org/ //RSS https://www.youtube.com/feeds/videos.xml?channel_id=UCQsLuQJupY5RiwMpEpFaswQ
								),
								linebot.NewCarouselColumn(
									imageURL, "其他相關網站", "其他福音網站",
									// linebot.NewURITemplateAction("佳音電台", "http://www.goodnews.org.tw/gnfm909.php"),
									LineTemplate_download_app,
									linebot.NewURITemplateAction("讚美之泉", "http://store.sop.org/product/mp3_mmo_search/"),
									linebot.NewURITemplateAction("天韻", "http://www.heavenlymelody.com.tw/videos/"),
								),
								linebot.NewCarouselColumn(
									imageURL, "其他功能", "各種這間教會的資訊",
									linebot.NewMessageTemplateAction("聚會時間 & 週報", "週報"),
									linebot.NewMessageTemplateAction("交通資訊", "教會地圖"),
									linebot.NewMessageTemplateAction("聯絡資訊", "聯絡資訊"),
								),
								LineTemplate_CarouselColumn_feedback,
							)
							t_msg := "台北公館教會的 YouTube：\nhttps://www.youtube.com/user/kkcpct\n\n台北公館教會的 Facebook\nhttps://m.facebook.com/TPEKKC\n\n台北公館教會的 Blog\nhttp://blog.kkcpct.org"
							obj_message := linebot.NewTemplateMessage(t_msg, template)
							if _, err = bot.ReplyMessage(event.ReplyToken, obj_message).Do(); err != nil {
									log.Print(1630)
									log.Print(err)
							}
							return
						case "選單":
							// if target_id_code == notalk{
							// 	if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("感謝您呼叫我！\n但為了不要過度打擾群組對話，在此群組暫時不支援此功能。\n請移步到 http://line.me/R/ti/g/X8_PRxyGL7 對話群組進行查詢。\n或個人+我（https://line.me/R/ti/p/@rgp6918s）好友，進行個人查詢。")).Do(); err != nil {
							// 			log.Print(24483)
							// 			log.Print(err)
							// 	}
							// 	return
							// }
						    imageURL = SystemImageURL
							//template := LineTemplate_firstinfo
							t_msg := "建議使用最新版本的 LINE APP 以獲得最佳互動介面體驗。\n" +
									"以下的內容你對我說，就會有相關的效果回應給你。\n" +
									"\n" +
									"【教會】\n" +
									"\n" +
									"輸入「聚會時間」可查聚會時間。\n" +
									"輸入「週報」可查最新電子週報以及下週預告。（互動介面有預習經文的功能）\n" +
									"輸入「教會地圖」可查各種教會。\n" +
									"輸入「聯絡資訊」可查電話、傳真、E-mail。（互動介面能直接撥打電話）\n" +
									"輸入「網站資訊」可查到現有的相關網站，包含 YouTube、Facebook...等。\n" +
									"輸入「圖書查詢」將引導你找到機構圖書。\n" +
									"\n" +
									"【推薦我】\n" +
									"\n" +
									"輸入「轉傳」會出現「我」的連結，方便轉傳連結讓更多朋友使用。\n" +
									"\n" +
									"【聖經】\n" +
									"\n" +
									"輸入「聖經」、「bible」或「Bible」\n" +
									"會有進一步的使用說明教你查詢。\n" +
									"基本支援六種語言、連續範圍查詢與呈現。\n" +
									"\n" +
									"查詢的文字結果也可以直接當作一般訊息按轉傳，\n" +
									"分享聖經節錄給你的朋友。\n" +
									"\n" +
									"如有其他建議，輸入「開發者」可進行聯絡。"
							obj_message := linebot.NewTemplateMessage(t_msg, LineTemplate_firstinfo)

							if target_id_code == notalk{
								if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("這是一種資訊整合的便捷應用，效果類似於自動回話小助理。\n\n概念上最基本的應用類似於遊戲 NPC 或 0800 電話總機，會根據指示自動回覆相關基本資訊。\n也可做其他延伸應用，像是聖經查詢 或 留言給意見...等等。\n\n目前除了教會相關資訊外，還可查詢 24 本聖經。\n支援 10 種語言、24 種聖經版本的精準經節查詢機能。\n並支援範圍查詢的寫法。（例如：聖經 創世紀 1:1-10）\n\n詳細說明可輸入「聖經」，有完整的使用說明介紹。") , obj_message).Do(); err != nil {
										log.Print(1639)
										log.Print(err)
								}
							}else{
								if _, err = bot.ReplyMessage(event.ReplyToken, obj_message).Do(); err != nil {
										log.Print(1640)
										log.Print(err)
								}
							}

							return
						case "家庭禮拜":
							//obj_message := linebot.NewStickerMessage("2", "514") //https://devdocs.line.me/en/?go#send-message-object
							if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(family_focus_text),obj_message_family_focus).Do(); err != nil {
									log.Print(500)
									log.Print(err)
							}
							return
						case "轉傳":
							if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("歡迎大家介紹 台北公館教會小天使 給更多人使用！\nhttps://line.me/R/ti/p/@rgp6918s\n\n你可以將這則文字訊息轉傳給其他對象。"),linebot.NewTextMessage("這是一種資訊整合的便捷應用，效果類似於自動回話小助理。\n\n概念上最基本的應用類似於遊戲 NPC 或 0800 電話總機，會根據指示自動回覆相關基本資訊。\n也可做其他延伸應用，像是聖經查詢 或 留言給意見...等等。\n\n目前除了教會相關資訊外，還可查詢 24 本聖經。\n支援 10 種語言、24 種聖經版本的精準經節查詢機能。\n並支援範圍查詢的寫法。（例如：聖經 創世紀 1:1-10）\n\n詳細說明可輸入「聖經」，有完整的使用說明介紹。")).Do(); err != nil {
									log.Print(7285)
									log.Print(err)
							}
							return
						case "已經傳送給教會":
							if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("已經幫你把意見傳送給教會囉！\n如需教會對您進行回覆，\n請補上個人聯絡資訊（mail 或其他）再發送一次！謝謝！"),linebot.NewStickerMessage("2", "514")).Do(); err != nil {
									log.Print(15391)
									log.Print(err)
							}
							return
						case "機器人88":
							if target_item == "群組對話" {
								log.Print("觸發離開群組，APP 限定")
								//post KEY = 離開群組
								template := linebot.NewConfirmTemplate(
									"你確定要請我離開嗎QAQ？",
									//.NewPostbackTemplateAction(按鈕字面,post,替使用者發言)
									linebot.NewPostbackTemplateAction("是","按下確定離開群組對話", ""),
									linebot.NewPostbackTemplateAction("否", "取消離開群組",""),
								)
								obj_message := linebot.NewTemplateMessage("你確定要請我離開嗎QAQ？\n這功能只支援 APP 使用。\n請用 APP 端查看下一步。", template)
								if _, err = bot.ReplyMessage(event.ReplyToken, obj_message).Do(); err != nil {
									log.Print(1654)
									log.Print(err)
								}
							}
							return
						case "新約列表":
							if target_id_code == notalk{
								if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("感謝您呼叫我！\n但為了不要過度打擾群組對話，在此群組暫時不支援此功能。\n請移步到 http://line.me/R/ti/g/X8_PRxyGL7 對話群組進行查詢。\n或個人+我（https://line.me/R/ti/p/@rgp6918s）好友，進行個人查詢。")).Do(); err != nil {
										log.Print(24483)
										log.Print(err)
								}
								return
							}
								// new_list := "【福音書】\n" +
								// 			"\n" +
								// 			"馬太福音\n" +
								// 			"馬可福音\n" +
								// 			"路加福音\n" +
								// 			"約翰福音\n" +
								// 			"\n" +
								// 			"【新約歷史書】\n" +
								// 			"\n" +
								// 			"使徒行傳\n" +
								// 			"\n" +
								// 			"【保羅書信】\n" +
								// 			"\n" +
								// 			"羅馬書\n" +
								// 			"哥林多前書\n" +
								// 			"哥林多後書\n" +
								// 			"加拉太書\n" +
								// 			"以弗所書\n" +
								// 			"腓立比書\n" +
								// 			"歌羅西書\n" +
								// 			"帖撒羅尼迦前書\n" +
								// 			"帖撒羅尼迦後書\n" +
								// 			"提摩太前書\n" +
								// 			"提摩太後書\n" +
								// 			"提多書\n" +
								// 			"腓利門書\n" +
								// 			"\n" +
								// 			"【其他書信】\n" +
								// 			"\n" +
								// 			"希伯來書\n" +
								// 			"雅各書\n" +
								// 			"彼得前書\n" +
								// 			"彼得後書\n" +
								// 			"約翰一書\n" +
								// 			"約翰二書\n" +
								// 			"約翰三書\n" +
								// 			"猶大書\n" +
								// 			"\n" +
								// 			"【預言書】\n" +
								// 			"\n" +
								// 			"啟示錄"
								imageURL = Bible_imageURL
								LineTemplate_old1 := linebot.NewCarouselTemplate(
									linebot.NewCarouselColumn(
										imageURL, "新約", "福音書",
										linebot.NewPostbackTemplateAction("馬太福音", "馬太福音", "聖經 馬太福音 1:1"),
										linebot.NewPostbackTemplateAction("馬可福音", "馬可福音", "聖經 馬可福音 1:1"),
										linebot.NewPostbackTemplateAction("路加福音", "路加福音", "聖經 路加福音 1:1"),
									),
									linebot.NewCarouselColumn(
										imageURL, "新約", "福音書(1) + 新約歷史書(1) + 保羅書信",
										linebot.NewPostbackTemplateAction("約翰福音", "約翰福音", "聖經 約翰福音 1:1"),
										linebot.NewPostbackTemplateAction("使徒行傳", "使徒行傳", "聖經 使徒行傳 1:1"),
										linebot.NewPostbackTemplateAction("羅馬書", "羅馬書", "聖經 羅馬書 1:1"),
									),
									linebot.NewCarouselColumn(
										imageURL, "新約", "保羅書信",
										linebot.NewPostbackTemplateAction("哥林多前書", "哥林多前書", "聖經 哥林多前書 1:1"),
										linebot.NewPostbackTemplateAction("哥林多後書", "哥林多後書", "聖經 哥林多後書 1:1"),
										linebot.NewPostbackTemplateAction("加拉太書", "加拉太書", "聖經 加拉太書 1:1"),
									),
									LineTemplate_CarouselColumn_feedback,
								)
								temp_msg := "【福音書】\n" +
											"\n" +
											"馬太福音\n" +
											"馬可福音\n" +
											"路加福音\n" +
											"約翰福音\n" +
											"\n" +
											"【新約歷史書】\n" +
											"\n" +
											"使徒行傳"
								obj_message1 := linebot.NewTemplateMessage(temp_msg, LineTemplate_old1)

								imageURL = Bible_imageURL
								LineTemplate_old2 := linebot.NewCarouselTemplate(
									linebot.NewCarouselColumn(
										imageURL, "新約", "保羅書信",
										linebot.NewPostbackTemplateAction("以弗所書", "以弗所書", "聖經 以弗所書 1:1"),
										linebot.NewPostbackTemplateAction("腓立比書", "腓立比書", "聖經 腓立比書 1:1"),
										linebot.NewPostbackTemplateAction("歌羅西書", "歌羅西書", "聖經 歌羅西書 1:1"),
									),
									linebot.NewCarouselColumn(
										imageURL, "新約", "保羅書信",
										linebot.NewPostbackTemplateAction("帖撒羅尼迦前書", "帖撒羅尼迦前書", "聖經 帖撒羅尼迦前書 1:1"),
										linebot.NewPostbackTemplateAction("帖撒羅尼迦後書", "帖撒羅尼迦後書", "聖經 帖撒羅尼迦後書 1:1"),
										linebot.NewPostbackTemplateAction("提摩太前書", "提摩太前書", "聖經 提摩太前書 1:1"),
									),
									linebot.NewCarouselColumn(
										imageURL, "新約", "保羅書信",
										linebot.NewPostbackTemplateAction("提摩太後書", "提摩太後書", "聖經 提摩太後書 1:1"),
										linebot.NewPostbackTemplateAction("提多書", "提多書", "聖經 提多書 1:1"),
										linebot.NewPostbackTemplateAction("腓利門書", "腓利門書", "聖經 腓利門書 1:1"),
									),
								)
								temp_msg = "【保羅書信】\n" +
											"\n" +
											"羅馬書\n" +
											"哥林多前書\n" +
											"哥林多後書\n" +
											"加拉太書\n" +
											"以弗所書\n" +
											"腓立比書\n" +
											"歌羅西書\n" +
											"帖撒羅尼迦前書\n" +
											"帖撒羅尼迦後書\n" +
											"提摩太前書\n" +
											"提摩太後書\n" +
											"提多書\n" +
											"腓利門書"
								obj_message2 := linebot.NewTemplateMessage(temp_msg, LineTemplate_old2)

								imageURL = Bible_imageURL
								LineTemplate_old3 := linebot.NewCarouselTemplate(
									linebot.NewCarouselColumn(
										imageURL, "新約", "其他書信",
										linebot.NewPostbackTemplateAction("希伯來書", "希伯來書", "聖經 希伯來書 1:1"),
										linebot.NewPostbackTemplateAction("雅各書", "雅各書", "聖經 雅各書 1:1"),
										linebot.NewPostbackTemplateAction("彼得前書", "彼得前書", "聖經 彼得前書 1:1"),
									),
									linebot.NewCarouselColumn(
										imageURL, "新約", "其他書信",
										linebot.NewPostbackTemplateAction("彼得後書", "彼得後書", "聖經 彼得後書 1:1"),
										linebot.NewPostbackTemplateAction("約翰一書", "約翰一書", "聖經 約翰一書 1:1"),
										linebot.NewPostbackTemplateAction("約翰二書", "約翰二書", "聖經 約翰二書 1:1"),
									),
									linebot.NewCarouselColumn(
										imageURL, "新約", "其他書信",
										linebot.NewPostbackTemplateAction("約翰三書", "約翰三書", "聖經 約翰三書 1:1"),
										linebot.NewPostbackTemplateAction("猶大書", "猶大書", "聖經 猶大書 1:1"),
										linebot.NewPostbackTemplateAction("啟示錄", "啟示錄", "聖經 啟示錄 1:1"),
									),
									LineTemplate_CarouselColumn_bible_one,
									LineTemplate_CarouselColumn_bible_list,
								)
								temp_msg = "【其他書信】\n" +
											"\n" +
											"希伯來書\n" +
											"雅各書\n" +
											"彼得前書\n" +
											"彼得後書\n" +
											"約翰一書\n" +
											"約翰二書\n" +
											"約翰三書\n" +
											"猶大書\n" +
											"\n" +
											"【預言書】\n" +
											"\n" +
											"啟示錄"
								obj_message3 := linebot.NewTemplateMessage(temp_msg, LineTemplate_old3)

								//if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(new_list)).Do(); err != nil {
								if _, err = bot.ReplyMessage(event.ReplyToken, obj_message1, obj_message2,obj_message3).Do(); err != nil {
										log.Print(7484)
										log.Print(err)
								}
							return
						case "舊約列表":
							if target_id_code == notalk{
								if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("感謝您呼叫我！\n但為了不要過度打擾群組對話，在此群組暫時不支援此功能。\n請移步到 http://line.me/R/ti/g/X8_PRxyGL7 對話群組進行查詢。\n或個人+我（https://line.me/R/ti/p/@rgp6918s）好友，進行個人查詢。")).Do(); err != nil {
										log.Print(24483)
										log.Print(err)
								}
								return
							}
								// old_list := "【摩西五經】\n" +
								// 			"\n" +
								// 			"創世紀\n" +
								// 			"出埃及記\n" +
								// 			"利未記\n" +
								// 			"民數記\n" +
								// 			"申命記\n" +
								// 			"\n" +
								// 			"【舊約歷史書】\n" +
								// 			"\n" +
								// 			"約書亞記\n" +
								// 			"士師記\n" +
								// 			"路得記\n" +
								// 			"撒母耳記上\n" +
								// 			"撒母耳記下\n" +
								// 			"列王紀上\n" +
								// 			"列王紀下\n" +
								// 			"歷代志上\n" +
								// 			"歷代志下\n" +
								// 			"以斯拉記\n" +
								// 			"尼希米記\n" +
								// 			"以斯帖記\n" +
								// 			"\n" +
								// 			"【詩歌智慧書】\n" +
								// 			"\n" +
								// 			"約伯記\n" +
								// 			"詩篇\n" +
								// 			"箴言\n" +
								// 			"傳道書\n" +
								// 			"雅歌\n" +
								// 			"\n" +
								// 			"【大先知書】\n" +
								// 			"\n" +
								// 			"以賽亞書\n" +
								// 			"耶利米書\n" +
								// 			"耶利米哀歌\n" +
								// 			"以西結書\n" +
								// 			"但以理書\n" +
								// 			"\n" +
								// 			"【小先知書】\n" +
								// 			"\n" +
								// 			"何西阿書\n" +
								// 			"約珥書\n" +
								// 			"阿摩司書\n" +
								// 			"俄巴底亞書\n" +
								// 			"約拿書\n" +
								// 			"彌迦書\n" +
								// 			"那鴻書\n" +
								// 			"哈巴谷書\n" +
								// 			"西番雅書\n" +
								// 			"哈該書\n" +
								// 			"撒迦利亞書\n" +
								// 			"瑪拉基書"
								// 																												linebot.NewTextMessage(old_list),
								imageURL = Bible_imageURL
								LineTemplate_old1 := linebot.NewCarouselTemplate(
									linebot.NewCarouselColumn(
										imageURL, "舊約", "摩西五經",
										linebot.NewPostbackTemplateAction("創世紀", "創世紀", "聖經 創世紀 1:1"),
										linebot.NewPostbackTemplateAction("出埃及記", "出埃及記", "聖經 出埃及記 1:1"),
										linebot.NewPostbackTemplateAction("利未記", "利未記", "聖經 利未記 1:1"),
									),
									linebot.NewCarouselColumn(
										imageURL, "舊約", "摩西五經(2) + 舊約歷史書(1)",
										linebot.NewPostbackTemplateAction("民數記", "民數記", "聖經 民數記 1:1"),
										linebot.NewPostbackTemplateAction("申命記", "申命記", "聖經 申命記 1:1"),
										linebot.NewPostbackTemplateAction("約書亞記", "約書亞記", "聖經 約書亞記 1:1"),
									),
									linebot.NewCarouselColumn(
										imageURL, "舊約", "舊約歷史書",
										linebot.NewPostbackTemplateAction("士師記", "士師記", "聖經 士師記 1:1"),
										linebot.NewPostbackTemplateAction("路得記", "路得記", "聖經 路得記 1:1"),
										linebot.NewPostbackTemplateAction("撒母耳記上", "撒母耳記上", "聖經 撒母耳記上 1:1"),
									),
									LineTemplate_CarouselColumn_bible_one,
									LineTemplate_CarouselColumn_feedback,
								)
								temp_msg := "【摩西五經】\n" +
											"\n" +
											"創世紀\n" +
											"出埃及記\n" +
											"利未記\n" +
											"民數記\n" +
											"申命記\n" +
											"\n" +
											"【舊約歷史書】\n" +
											"\n" +
											"約書亞記\n" +
											"士師記\n" +
											"路得記\n" +
											"撒母耳記上\n" +
											"撒母耳記下\n" +
											"列王紀上\n" +
											"列王紀下\n" +
											"歷代志上\n" +
											"歷代志下\n" +
											"以斯拉記\n" +
											"尼希米記\n" +
											"以斯帖記"
								obj_message1 := linebot.NewTemplateMessage(temp_msg, LineTemplate_old1)

								LineTemplate_old2 := linebot.NewCarouselTemplate(
									linebot.NewCarouselColumn(
										imageURL, "舊約", "舊約歷史書",
										linebot.NewPostbackTemplateAction("撒母耳記下", "撒母耳記下", "聖經 撒母耳記下 1:1"),
										linebot.NewPostbackTemplateAction("列王紀上", "列王紀上", "聖經 列王紀上 1:1"),
										linebot.NewPostbackTemplateAction("列王紀下", "列王紀下", "聖經 列王紀下 1:1"),
									),
									linebot.NewCarouselColumn(
										imageURL, "舊約", "舊約歷史書",
										linebot.NewPostbackTemplateAction("歷代志上", "歷代志上", "聖經 歷代志上 1:1"),
										linebot.NewPostbackTemplateAction("歷代志下", "歷代志下", "聖經 歷代志下 1:1"),
										linebot.NewPostbackTemplateAction("以斯拉記", "以斯拉記", "聖經 以斯拉記 1:1"),
									),
									linebot.NewCarouselColumn(
										imageURL, "舊約", "舊約歷史書(2) + 詩歌智慧書(1)",
										linebot.NewPostbackTemplateAction("尼希米記", "尼希米記", "聖經 尼希米記 1:1"),
										linebot.NewPostbackTemplateAction("以斯帖記", "以斯帖記", "聖經 以斯帖記 1:1"),
										linebot.NewPostbackTemplateAction("約伯記", "約伯記", "聖經 約伯記 1:1"),
									),
								)
								temp_msg = "【詩歌智慧書】\n" +
											"\n" +
											"約伯記\n" +
											"詩篇\n" +
											"箴言\n" +
											"傳道書\n" +
											"雅歌"
								obj_message2 := linebot.NewTemplateMessage(temp_msg, LineTemplate_old2)

								LineTemplate_old3 := linebot.NewCarouselTemplate(
									linebot.NewCarouselColumn(
										imageURL, "舊約", "詩歌智慧書",
										linebot.NewPostbackTemplateAction("詩篇", "詩篇", "聖經 詩篇 1:1"),
										linebot.NewPostbackTemplateAction("箴言", "箴言", "聖經 箴言 1:1"),
										linebot.NewPostbackTemplateAction("傳道書", "傳道書", "聖經 傳道書 1:1"),
									),
									linebot.NewCarouselColumn(
										imageURL, "舊約", "詩歌智慧書(1) + 大先知書(2)",
										linebot.NewPostbackTemplateAction("雅歌", "雅歌", "聖經 雅歌 1:1"),
										linebot.NewPostbackTemplateAction("以賽亞書", "以賽亞書", "聖經 以賽亞書 1:1"),
										linebot.NewPostbackTemplateAction("耶利米書", "耶利米書", "聖經 耶利米書 1:1"),
									),
									linebot.NewCarouselColumn(
										imageURL, "舊約", "大先知書",
										linebot.NewPostbackTemplateAction("耶利米哀歌", "耶利米哀歌", "聖經 耶利米哀歌 1:1"),
										linebot.NewPostbackTemplateAction("以西結書", "以西結書", "聖經 以西結書 1:1"),
										linebot.NewPostbackTemplateAction("但以理書", "但以理書", "聖經 但以理書 1:1"),
									),
								)
								temp_msg =  "【大先知書】\n" +
											"\n" +
											"以賽亞書\n" +
											"耶利米書\n" +
											"耶利米哀歌\n" +
											"以西結書\n" +
											"但以理書"
								obj_message3 := linebot.NewTemplateMessage(temp_msg, LineTemplate_old3)

								LineTemplate_old4 := linebot.NewCarouselTemplate(
									linebot.NewCarouselColumn(
										imageURL, "舊約", "小先知書",
										linebot.NewPostbackTemplateAction("何西阿書", "何西阿書", "聖經 何西阿書 1:1"),
										linebot.NewPostbackTemplateAction("約珥書", "約珥書", "聖經 約珥書 1:1"),
										linebot.NewPostbackTemplateAction("阿摩司書", "阿摩司書", "聖經 阿摩司書 1:1"),
									),
									linebot.NewCarouselColumn(
										imageURL, "舊約", "小先知書",
										linebot.NewPostbackTemplateAction("俄巴底亞書", "俄巴底亞書", "聖經 俄巴底亞書 1:1"),
										linebot.NewPostbackTemplateAction("約拿書", "約拿書", "聖經 約拿書 1:1"),
										linebot.NewPostbackTemplateAction("彌迦書", "彌迦書", "聖經 彌迦書 1:1"),
									),
									linebot.NewCarouselColumn(
										imageURL, "舊約", "小先知書",
										linebot.NewPostbackTemplateAction("那鴻書", "那鴻書", "聖經 那鴻書 1:1"),
										linebot.NewPostbackTemplateAction("哈巴谷書", "哈巴谷書", "聖經 哈巴谷書 1:1"),
										linebot.NewPostbackTemplateAction("西番雅書", "西番雅書", "聖經 西番雅書 1:1"),
									),
									linebot.NewCarouselColumn(
										imageURL, "舊約", "小先知書",
										linebot.NewPostbackTemplateAction("哈該書", "哈該書", "聖經 哈該書 1:1"),
										linebot.NewPostbackTemplateAction("撒迦利亞書", "撒迦利亞書", "聖經 撒迦利亞書 1:1"),
										linebot.NewPostbackTemplateAction("瑪拉基書", "瑪拉基書", "聖經 瑪拉基書 1:1"),
									),
									LineTemplate_CarouselColumn_bible_list,
								)
								temp_msg = "【小先知書】\n" +
											"\n" +
											"何西阿書\n" +
											"約珥書\n" +
											"阿摩司書\n" +
											"俄巴底亞書\n" +
											"約拿書\n" +
											"彌迦書\n" +
											"那鴻書\n" +
											"哈巴谷書\n" +
											"西番雅書\n" +
											"哈該書\n" +
											"撒迦利亞書\n" +
											"瑪拉基書"
								obj_message4 := linebot.NewTemplateMessage(temp_msg + "\n\n你也可以用最新版本的 LINE APP 查看，\n有很多便捷的介面按鈕可以更快速查詢。", LineTemplate_old4)

								//if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(old_list)).Do(); err != nil {
								if _, err = bot.ReplyMessage(event.ReplyToken, obj_message1,obj_message2,obj_message3,obj_message4).Do(); err != nil {
										log.Print(1286)
										log.Print(err)
								}
							return
						case (message.Text + "？\n抱歉目前找不到\n"):
							//相當於 reg_nofind.ReplaceAllString(bot_msg, "$1")=="我還沒學呢..."
								if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(bot_msg)).Do(); err != nil {
										log.Print(1192)
										log.Print(err)
								}
							return
						// case "查詢章節超過聖經範圍，有可能指定查詢的節超過範圍。":

						// 	return
						default: //查詢成功的內容(暫時，更嚴謹的話要在這之前分析...)
					 		// if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(bot_msg)).Do(); err != nil {
					 		// 	log.Print(1771)
					 		// 	log.Print(err)
					 		// }
							if target_id_code == notalk{
								if bot_msg=="聖經" {
									if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("感謝您呼叫我！\n但為了不要過度打擾群組對話，在此群組暫時不支援此功能。\n請移步到 http://line.me/R/ti/g/X8_PRxyGL7 對話群組進行查詢。\n或個人+我（https://line.me/R/ti/p/@rgp6918s）好友，進行個人查詢。")).Do(); err != nil {
											log.Print(24483)
											log.Print(err)
									}
									return
								}else{
									//放行聖經查詢功能
									//return									
								}
							}

							//觀察
							log.Print(`觀察 message.Text + "？\n抱歉目前找不到" = `)
							log.Print(bot_msg==message.Text + "？\n抱歉目前找不到")

							if (reg_nofind.ReplaceAllString(bot_msg, "$1")!="我還沒學呢..."){
								//查詢成功

								if bible_short_name=="" {
									//連書都沒有但「有聖經的時候」

									imageURL = Bible_imageURL
									LineTemplate_bible_info := linebot.NewCarouselTemplate(
										linebot.NewCarouselColumn(
											imageURL, "聖經查詢方法", "以下是示範。\n也可以手動輸入試試看各種組合。",
											linebot.NewPostbackTemplateAction("聖經 何西阿書 14：1-3","聖經 何西阿書 14：1-3","聖經 何西阿書 14：1-3"),
											linebot.NewPostbackTemplateAction("英文聖經 出埃及 1：4-5","英文聖經 出埃及 1：4-5","英文聖經 出埃及 1：4-5"),
											linebot.NewPostbackTemplateAction("多國語言聖經 創世紀 5：1","多國語言聖經 創世紀 5：1","多國語言聖經 創世紀 5：1"),
										),
										linebot.NewCarouselColumn(
											imageURL, "更多示範", "以下是示範。\n也可以手動輸入試試看各種組合。",
											linebot.NewPostbackTemplateAction("日文聖經 啟示錄 6：5-6","日文聖經 啟示錄 6：5-6","日文聖經 啟示錄 6：5-6"),
											linebot.NewPostbackTemplateAction("韓文聖經 創 ５：５－７","韓文聖經：創世紀：５：５－７","韓文聖經：創世紀：５：５－７"),
											linebot.NewPostbackTemplateAction("越南聖經；創世紀；5；5-9","越南聖經；創世紀；5；5-9","越南聖經；創世紀；5；5-9"),
										),
										linebot.NewCarouselColumn(
											imageURL, "更多速查示範", "以下是示範。\n也可以手動輸入試試看各種組合。",
											linebot.NewPostbackTemplateAction("俄文聖經 創 5 1-10","俄文聖經 創 5 1-10","俄文聖經 創 5 1-10"),
											linebot.NewPostbackTemplateAction("台語聖經 Gen 1：1-7","台語聖經 Gen 1：1-7","台語聖經 Gen 1：1-7"),
											linebot.NewPostbackTemplateAction("All bible Gen 5:5","All bible Gen 5:5","All bible Gen 5:5"),
										),
										LineTemplate_CarouselColumn_bible_list,
										LineTemplate_CarouselColumn_feedback,
									)
									//temp_msg := bot_msg
									obj_message := linebot.NewTemplateMessage("你也可以開最新版本的 LINE APP 直接查看試用相關使用範例。", LineTemplate_bible_info)

									bible_cmd_info_string := "我可以幫你精準查詢聖經章節！\n\n" +
											"【使用方法】\n\n一本聖經 + (分隔符) + 某本書 + (分隔符) + 篇 + (分隔符) + 節\n※ 分隔符 = 空白、冒號或分號。\n「節」可以用 1-10 的方式列出範圍經節。\n\n" +
											"目前可以查詢六種語言跟同時多國語言版本的聖經。\n\n" +
											"例如：\n" +
											"聖經 創世紀 1：1\n" +
											"英文聖經 出埃及 1：4-5\n" +
											"日文聖經 啟示錄 6：5-6\n" +
											"韓文聖經 創世紀 1：1-7\n" +
											"越南聖經 創世紀 1：1-8\n" +
											"俄文聖經 創世紀 1：1\n" +
											"多國語聖經 創世紀 1：1\n" +
											"\n" +
											"如果你習慣其他語言，也可以用英文、日文、韓文搜尋。\n" +
											"例如：\n" +
											"聖經 Joshua 1:1-3\n" +
											"聖經 jos 1:1-3\n" +
											"聖經 ヨシュア記 1:1-3\n" +
											"聖書 出エジプト記 1:1-5\n" +
											"聖經 여호수아 1:1-3\n\n" +
											"另有其他短稱的寫法查詢，請另外輸入「查詢可用簡寫」來查看所有可用關鍵字。\n" +
											"\n" +
											"除了多語言同步呈現的聖經之外，\n" +
											"還有另一種「研究聖經」的模式，或稱「多版聖經」。\n" +
											"可以同時呈現各語言中的其他版本。\n" +
											"例如：\n" + 
											"會同時呈現台語全羅跟漢羅以及文言文、希臘文古譯、馬索拉原文。\n" +
											"\n" +
											"使用方式：\n" +
											"「研究聖經 創世紀 1:1-2」或「多版本聖經 啟示錄 1:1-2」之類的寫法。\n" +
											"\n" +
											"他會一次查詢所有版本。\n" +
											"目前支援 24 種版本的聖經，所以他會同時查詢 24 種版本的聖經。\n" +
											"查詢時間需要等待三十秒以上是正常的，請耐心等候。\n" +
											"\n" +
											"但請特別注意！不建議用此方法查詢太大範圍的節。\n" +
											"因為 LINE 有限制單則訊息只能容納 2000 字。\n" +
											"雖然開發者有做技術突破自動切割超過的內容，另外連發訊息做彌補。\n" +
											"但最多只能容納三發訊息，\n" +
											"也就是總共最多只能回傳 6000 字。\n" +
											"因此建議不要查詢超過 3 節，\n" +
											"或改用「多國聖經」的方式查詢語言對比結果。" 
									bible_ver_list_string := `以下介紹單獨使用的各版本聖經的觸發關鍵字
「=」代表同樣功能，會觸發查詢同一本聖經。
都用於查詢聖經查詢功能的「開頭」，將依照不同開頭查詢不同聖經。

共支援六種主流語言及其他四種語言，總計支援 24 個聖經版本。

【中文聖經系列】

※ 目前預設「聖經」為中文和合本，觸發字有：聖經 = bible = Bible = ｂｉｂｌｅ = Ｂｉｂｌｅ

中文聖經 = 中文聖經和合本修訂版 = Rcuv = rcuv = ｒｃｕｖ = Ｒｃｕｖ
中文聖經新譯本 = ncv = Ncv = Ｎｃｖ = ｎｃｖ
中文聖經譯本修訂版 = tcv = TCV = Ｔｃｖ = ＴＣＶ
文言文聖經 = 深文理和合本

【台語聖經系列】

台語聖經 = 閩南語聖經 = 台語聖經漢羅 = 全民台語聖經漢羅
台語聖經全羅 = 全民台語聖經全羅

台語聖經馬雅各漢羅
台語聖經馬雅各全羅

台語聖經巴克禮漢羅
台語聖經巴克禮全羅

【客家】※ 只供查詢新約

客家聖經

【英文聖經系列】

※ 目前預設「英文聖經」為 KJV 版本，觸發字有：
英文聖經 = 英語聖書 = Kjv = kjv = Ｋｊｖ = ｋｊｖ = Eng bible = ENG Bible = English bible

英文聖經ERV = erv = ERV = Erv = ＥＲＶ = Ｅｒｖ = ｅｒｖ
英文聖經Darby = darby = DARBY = Ｄａｒｂｙ = ＤＡＲＢＹ = ｄａｒｂｙ
英文聖經ASV = ASV = Asv = asv = ＡＳＶ = Ａｓｖ = ａｓｖ
英文聖經WEB = WEB = Web = web = ＷＥＢ = Ｗｅｂ = ｗｅｂ
英文聖經BBE = BBE = Bbe = bbe = ＢＢＥ = Ｂｂｅ = ｂｂｅ

【其他外語聖經】

日文聖經 = 聖書 = 日本語聖書 = JP bible = JP Bible = Jp bible
韓文聖經 = KR bible = Korean = korean = Kr Bible = Kr bible
越南聖經
俄文聖經

【古譯文】※ 只供查詢舊約

馬索拉聖經 = bhs = Bhs = BHS = ＢＨＳ = Ｂｈｓ = ｂｈｓ
希臘聖經 = lxx = LXX = Lxx = ＬＸＸ = Ｌｘｘ = ｌｘｘ

【同時查詢多本的功能】

多國聖經：只會列出各語言其中一本聖經，不包含古譯文。
多版聖經：會列出所有可查詢的所有版本聖經。

多國聖經 = 多語聖經 = 多語言聖經 = 多國語聖經 = 多國語言聖經 = allbible = all bible = All bible = All Bible
研究聖經 = 總和聖經 = 綜合聖經 = 聖經研究 = 多版聖經 = 多版本聖經 = Allbible

`
									if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(bible_cmd_info_string),linebot.NewTextMessage(bible_ver_list_string),obj_message).Do(); err != nil {
											log.Print(7387)
											log.Print(err)
									}
									return
								}

								if bible_chap=="" {
									bible_chap = "1"
								}

								if bible_sec=="" {
									bible_sec = "1"
								}

								if bot_msg == "查詢章節超過聖經範圍，有可能指定查詢的節超過範圍。"{
									bible_chap = "1"
									bible_sec = "1"
								}

								taiwan_mp3_ver := "1"
								bible_id_jp := "01"
								bible_chap_jp := "001"
								//https://golang.org/pkg/strconv/#example_Atoi
								if bible_id_int, err := strconv.Atoi(bible_id); err == nil {
									//成功轉型 int
									if bible_id_int >= 40 {
										taiwan_mp3_ver = "5"	//新約 (bible_id=40 以上) 使用 5 ver
									}
									if bible_id_int < 10 {
										bible_id_jp = "0" + bible_id
									}else{
										bible_id_jp = bible_id
									}
								}

								if bible_chap_int, err := strconv.Atoi(bible_chap); err == nil {
									//成功轉型 int
									if bible_chap_int < 10 {
										bible_chap_jp = "00" + bible_chap
									}else{
										if bible_chap_int < 100 {
											bible_chap_jp = "0" + bible_chap
										}else{
											bible_chap_jp = bible_chap
										}
									}
								}

								imageURL = Bible_imageURL
								LineTemplate_find := linebot.NewCarouselTemplate(
									linebot.NewCarouselColumn(
										imageURL, "查看整篇前後文", "找到 " + message.Text + "！\n可按按鈕去看前後文。\n輸入「聖經」可以知道查詢方法。",
										//linebot.NewURITemplateAction("中文","http://bible.fhl.net/new/read.php?chineses=" + bible_short_name + "&nodic=1&chap=" + bible_chap + "&TABFLAG=1&sec=" + bible_sec + "&VERSION1=unv&VERSION2=kjv"),
										linebot.NewURITemplateAction("中文和合本","http://bible.fhl.net/new/read.php?chineses=" + bible_short_name + "&nodic=1&chap=" + bible_chap + "&TABFLAG=1&VERSION1=unv"),
										linebot.NewURITemplateAction("英文（KJV）","http://bible.fhl.net/new/read.php?chineses=" + bible_short_name + "&nodic=1&chap=" + bible_chap + "&TABFLAG=1&VERSION1=kjv"),
										linebot.NewURITemplateAction("日文","http://bible.fhl.net/new/read.php?chineses=" + bible_short_name + "&nodic=1&chap=" + bible_chap + "&TABFLAG=1&VERSION1=jp"),
									),
									linebot.NewCarouselColumn(
										imageURL, "♪ 有聲聖經 ♪", "以下是你查詢的章節，\n請選擇你喜歡的語言收聽 ♪\n※ 這播放時會耗費上網流量",
										linebot.NewURITemplateAction("中文和合本","http://bible.fhl.net/new/listenhb.php?version=0&bid=" + bible_id + "&chap=" + bible_chap + "&auto=1"),
										linebot.NewURITemplateAction("台語","http://bible.fhl.net/new/listenhb.php?version=" + taiwan_mp3_ver + "&bid=" + bible_id + "&chap=" + bible_chap + "&auto=1"),
										//linebot.NewURITemplateAction("客家話","http://bible.fhl.net/new/listenhb.php?version=2&bid=" + bible_id + "&chap=" + bible_chap + "&auto=1"),
										linebot.NewURITemplateAction("日文","http://bible.salterrae.net/kougo/gtalk/" + bible_id_jp + "/" + bible_id_jp + "_" + bible_chap_jp + ".mp3"),
									),
									linebot.NewCarouselColumn(
										imageURL, "♪ 有聲聖經 ♪", "以下是你查詢的章節，\n請選擇你喜歡的語言收聽 ♪\n※ 這播放時會耗費上網流量",
										//linebot.NewURITemplateAction("英文","https://www.bible.com/zh-TW/bible/1/" + bible_com_text + "." + bible_chap),

										//linebot.NewURITemplateAction("俄文","https://www.bible.com/zh-TW/bible/143/" + bible_com_text + "." + bible_chap),
										linebot.NewURITemplateAction("福州話","http://bible.fhl.net/new/listenhb.php?version=8&bid=" + bible_id + "&chap=" + bible_chap + "&auto=1"),
										linebot.NewURITemplateAction("希伯來文讀經","http://bible.fhl.net/new/listenhb.php?version=7&bid=" + bible_id + "&chap=" + bible_chap + "&auto=1"),
										linebot.NewURITemplateAction("關於有聲聖經","http://bible.fhl.net/new/audio.html"),
									),
									// linebot.NewCarouselColumn(
									// 	imageURL, "♪ 有聲聖經 ♪", "以下是你查詢的章節，\n請選擇你喜歡的語言收聽 ♪\n※ 這播放時會耗費上網流量",
									// 	linebot.NewURITemplateAction("中文和合本","http://bible.fhl.net/new/listenhb.php?version=0&bid=" + bible_id + "&chap=" + bible_chap + "&auto=1"),
									// 	linebot.NewURITemplateAction("台語","http://bible.fhl.net/new/listenhb.php?version=" + taiwan_mp3_ver + "&bid=" + bible_id + "&chap=" + bible_chap + "&auto=1"),
									// 	linebot.NewURITemplateAction("客家話","http://bible.fhl.net/new/listenhb.php?version=2&bid=" + bible_id + "&chap=" + bible_chap + "&auto=1"),
									// ),
									LineTemplate_CarouselColumn_bible_one,
									LineTemplate_CarouselColumn_bible_list,
								)
								wordmode_str := "【查看整篇前後文】\n\n" +
										"中文和合本\n" + "http://bible.fhl.net/new/read.php?chineses=" + bible_short_name + "&nodic=1&chap=" + bible_chap + "&TABFLAG=1&VERSION1=unv\n\n" +
										"英文（KJV）\nhttp://bible.fhl.net/new/read.php?chineses=" + bible_short_name + "&nodic=1&chap=" + bible_chap + "&TABFLAG=1&VERSION1=kjv\n\n" +
										"日文\nhttp://bible.fhl.net/new/read.php?chineses=" + bible_short_name + "&nodic=1&chap=" + bible_chap + "&TABFLAG=1&VERSION1=jp\n\n"
										// "【♪ 有聲聖經 ♪】\n\n" +
										// "中文和合本\nhttp://bible.fhl.net/new/listenhb.php?version=0&bid=" + bible_id + "&chap=" + bible_chap + "&auto=1\n\n" +
										// "台語\nhttp://bible.fhl.net/new/listenhb.php?version=" + taiwan_mp3_ver + "&bid=" + bible_id + "&chap=" + bible_chap + "&auto=1\n\n" +
										// "客家話\nhttp://bible.fhl.net/new/listenhb.php?version=2&bid=" + bible_id + "&chap=" + bible_chap + "&auto=1\n\n" +
										//"輸入「聖經」會出現如何查詢聖經。\n輸入「舊約列表」出現舊約\n輸入「舊約列表」出現新約"
								LineTemplate_find2 := linebot.NewCarouselTemplate(
									linebot.NewCarouselColumn(
										imageURL, "查詢其他語言", "找到 " + message.Text + "！\n查詢其他語言",
										//linebot.NewURITemplateAction("中文","http://bible.fhl.net/new/read.php?chineses=" + bible_short_name + "&nodic=1&chap=" + bible_chap + "&TABFLAG=1&sec=" + bible_sec + "&VERSION1=unv&VERSION2=kjv"),
										linebot.NewPostbackTemplateAction("英文（KJV）", "英文（KJV）", "英文聖經：" + bible_short_name + " " + bible_chap + "：" + bible_sec),
										linebot.NewPostbackTemplateAction("日文聖經", "日文聖經", "日文聖經：" + bible_short_name + " " + bible_chap + "：" + bible_sec),
										linebot.NewPostbackTemplateAction("多國語言並列", "多國語言並列", "多國語聖經：" + bible_short_name + " " + bible_chap + "：" + bible_sec),
										// linebot.NewURITemplateAction("和合本","http://bible.fhl.net/new/read.php?chineses=" + bible_short_name + "&nodic=1&chap=" + bible_chap + "&TABFLAG=1&VERSION1=unv"),
										// linebot.NewURITemplateAction("英文（KJV）","http://bible.fhl.net/new/read.php?chineses=" + bible_short_name + "&nodic=1&chap=" + bible_chap + "&TABFLAG=1&VERSION1=kjv"),
										// linebot.NewURITemplateAction("日文","http://bible.fhl.net/new/read.php?chineses=" + bible_short_name + "&nodic=1&chap=" + bible_chap + "&TABFLAG=1&VERSION1=jp"),
									),
									linebot.NewCarouselColumn(
										imageURL, "查詢其他語言", "共支援六種語言查詢！\n中、英、日、韓、越、俄，與多國聖經。",
										//linebot.NewURITemplateAction("中文","http://bible.fhl.net/new/read.php?chineses=" + bible_short_name + "&nodic=1&chap=" + bible_chap + "&TABFLAG=1&sec=" + bible_sec + "&VERSION1=unv&VERSION2=kjv"),
										linebot.NewPostbackTemplateAction("韓文聖經", "韓文聖經", "韓文聖經：" + bible_short_name + " " + bible_chap + "：" + bible_sec),
										linebot.NewPostbackTemplateAction("越南聖經", "越南聖經", "越南聖經：" + bible_short_name + " " + bible_chap + "：" + bible_sec),
										linebot.NewPostbackTemplateAction("俄文聖經", "俄文聖經", "俄文聖經：" + bible_short_name + " " + bible_chap + "：" + bible_sec),
									),
									linebot.NewCarouselColumn(
										imageURL, "其他語言查看整篇", "可按按鈕去看。",
										//linebot.NewURITemplateAction("中文","http://bible.fhl.net/new/read.php?chineses=" + bible_short_name + "&nodic=1&chap=" + bible_chap + "&TABFLAG=1&sec=" + bible_sec + "&VERSION1=unv&VERSION2=kjv"),
										linebot.NewURITemplateAction("台語","http://bible.fhl.net/new/read.php?chineses=" + bible_short_name + "&nodic=1&chap=" + bible_chap + "&TABFLAG=1&VERSION1=unv&VERSION2=gebklhlruby"),
										linebot.NewURITemplateAction("中英對照","http://bible.fhl.net/new/read.php?chineses=" + bible_short_name + "&nodic=1&chap=" + bible_chap + "&TABFLAG=1&VERSION1=unv&VERSION2=kjv"),
										linebot.NewURITemplateAction("中日對照","http://bible.fhl.net/new/read.php?chineses=" + bible_short_name + "&nodic=1&chap=" + bible_chap + "&TABFLAG=1&VERSION1=unv&VERSION2=jp"),
									),
									// LineTemplate_other_example,
									// LineTemplate_other,
									LineTemplate_CarouselColumn_feedback,
								)
								wordmode_str2 := //"台語\nhttp://bible.fhl.net/new/read.php?chineses=" + bible_short_name + "&nodic=1&chap=" + bible_chap + "&TABFLAG=1&VERSION1=gebklhlruby\n\n" +
										"中文/台語對照\nhttp://bible.fhl.net/new/read.php?chineses=" + bible_short_name + "&nodic=1&chap=" + bible_chap + "&TABFLAG=1&VERSION1=unv&VERSION2=gebklhlruby\n\n" +
										"中英對照\nhttp://bible.fhl.net/new/read.php?chineses=" + bible_short_name + "&nodic=1&chap=" + bible_chap + "&TABFLAG=1&VERSION1=unv&VERSION2=kjv\n\n"+
										"中日對照\nhttp://bible.fhl.net/new/read.php?chineses=" + bible_short_name + "&nodic=1&chap=" + bible_chap + "&TABFLAG=1&VERSION1=unv&VERSION2=jp"
										// "中韓對照\nhttp://bible.fhl.net/new/read.php?chineses=" + bible_short_name + "&nodic=1&chap=" + bible_chap + "&TABFLAG=1&VERSION1=unv&VERSION2=korean\n\n" +
										// "中越對照\nhttp://bible.fhl.net/new/read.php?chineses=" + bible_short_name + "&nodic=1&chap=" + bible_chap + "&TABFLAG=1&VERSION1=unv&VERSION2=vietnamese\n\n" +
										// "中俄對照\nhttp://bible.fhl.net/new/read.php?chineses=" + bible_short_name + "&nodic=1&chap=" + bible_chap + "&TABFLAG=1&VERSION1=unv&VERSION2=russian"
								//temp_msg := bot_msg
								obj_message := linebot.NewTemplateMessage(wordmode_str2, LineTemplate_find)
								obj_message2 := linebot.NewTemplateMessage("你可以開最新版本的 LINE APP 有方便的按鈕可以使用。\n單獨輸入「聖經」可以知道查詢方法。\n\n" + wordmode_str, LineTemplate_find2)
								//2017.02.14+ 閹割按鈕版本 notalk 的時候不使用
								if target_id_code != notalk{
									if _, err = bot.ReplyMessage(event.ReplyToken, obj_message2,obj_message,linebot.NewTextMessage(bot_msg)).Do(); err != nil {
											log.Print(7557)
											log.Print(err)
											log.Print("linebot: APIError 400 The request body has 1 error(s)\n[messages[2].text] Length must be between 0 and 2000"==err.Error()) //bot_msg 位子換了所以變成 [2]
											// //HttpPost_JANDI(target_item + "[" + user_talk + "](" + userImageUrl + ")：" + message.Text + `\n` + userStatus, "red" , "查詢失敗" + `\n` + err.Error(),target_id_code)
											// //HttpPost_IFTTT(target_item + " " + user_talk + "：" + message.Text + `\n<br>` + userImageUrl + `\n<br>` + userStatus , "LINE 同步：查詢失敗" + `\n<br>` +  err.Error(),target_id_code)

											// if(err.Error()=="linebot: APIError 400 The request body has 1 error(s)\n[messages[2].text] Length must be between 0 and 2000"){
											// 	if _, err = bot.ReplyMessage(event.ReplyToken,linebot.NewStickerMessage("2", "152"),linebot.NewTextMessage("【查詢 "  + message.Text +  " 發生錯誤】\n\n查詢得到的回應超過 2000 字，\n超過 LINE 的承受上限。\n\n請減少查詢節數再查一次。")).Do(); err != nil {
											// 		log.Print(7489)
											// 		log.Print(err.Error())
											// 	}
											// }
											switch err.Error(){
												case "linebot: APIError 400 The request body has 1 error(s)\n[messages[0].text] Length must be between 0 and 2000","linebot: APIError 400 The request body has 1 error(s)\n[messages[1].text] Length must be between 0 and 2000","linebot: APIError 400 The request body has 1 error(s)\n[messages[2].text] Length must be between 0 and 2000":
													log.Print("！！！！！有走進來 case error msg2 must be between 0 and 2000，總數 = ")
													//log.Print(len(bot_msg)) //7121 bytes
													log.Print(len([]rune(bot_msg))) //https://play.golang.org/p/yikJz-BKOW //utf8.RuneLen

													//log.Print("測試 = ")
													//log.Print(len(bot_msg[0:3000]))
													//log.Print(len([]rune(string([]rune(bot_msg)[0:2000]))))
														//部分文字列を取り出す	http://ashitani.jp/golangtips/tips_string.html#string_Replace
													//string([]rune(test_string)[0:len([]rune(test_string))])	https://play.golang.org/p/ivzNYS711B //UTF-8 的顯示部分字串的方法 結合 len([]rune(test_string)) 就得到 UTF-8 版本的字數
													//http://stackoverflow.com/questions/15018545/how-to-index-characters-in-a-golang-string
													if _, err = bot.ReplyMessage(event.ReplyToken,
														obj_message2,
														obj_message,
														linebot.NewTextMessage(string([]rune(bot_msg)[0:2000])),
														linebot.NewTextMessage(string([]rune(bot_msg)[1990:len([]rune(bot_msg))])),
													).Do(); err != nil {
														//第一次挽救：分割成兩個還是失敗，一定是後面那個太胖。
														log.Print(7587)
														log.Print(err.Error())
														if _, err = bot.ReplyMessage(event.ReplyToken,
															obj_message2,
															obj_message,
															linebot.NewTextMessage(string([]rune(bot_msg)[0:2000])),
															linebot.NewTextMessage(string([]rune(bot_msg)[1990:3990])),
															linebot.NewTextMessage(string([]rune(bot_msg)[3980:len([]rune(bot_msg))])),
														).Do(); err != nil {
														//if _, err = bot.ReplyMessage(event.ReplyToken,obj_message2,obj_message,linebot.NewTextMessage(bot_msg[0:3000])).Do(); err != nil {
														//if _, err = bot.ReplyMessage(event.ReplyToken,linebot.NewTextMessage(len(bot_msg)).Do(); err != nil {
															//HttpPost_JANDI(target_item + "[" + user_talk + "](" + userImageUrl + ")：" + message.Text + `\n` + userStatus, "red" , "查詢失敗" + `\n` + err.Error(),target_id_code)
															//HttpPost_IFTTT(target_item + " " + user_talk + "：" + message.Text + `\n<br>` + userImageUrl + `\n<br>` + userStatus , "LINE 同步：查詢失敗" + `\n<br>` +  err.Error(),target_id_code)
															log.Print(7600)
															log.Print(err.Error())
															//連續三次發現失敗(原始內容、分成2、分成3)，勸退XD
															if _, err = bot.ReplyMessage(event.ReplyToken,
																linebot.NewTextMessage(string([]rune(bot_msg)[0:2000])),
																linebot.NewTextMessage(string([]rune(bot_msg)[1990:3990])),
																linebot.NewTextMessage(string([]rune(bot_msg)[3980:5980])),
																linebot.NewStickerMessage("2", "152"),
																linebot.NewTextMessage("【查詢 "  + message.Text +  " 發生錯誤】\n\n查詢得到的回應超過 6000 字，\n超過 LINE 可以傳輸的上限。\n因此內容最後有遺漏。\n\n請減少查詢節的數量，重新再查一次。"),
															).Do(); err != nil {
																log.Print(7604)
																log.Print(err.Error())
															}
														}else{
															// 第二次分割成三個才成功
															send_color := "yellow"
															send_title := "查詢成功"
															if bot_msg=="查詢章節超過聖經範圍，有可能指定查詢的節超過範圍。"{
																send_color = "orange"
																send_title = "查詢失敗，範圍超過。"
															}
															HttpPost_JANDI(target_item + " [" + user_talk + "](" + userImageUrl + ")：" + message.Text + `\n` + userStatus, send_color , "LINE 同步：" + send_title + `\n` + bot_msg,target_id_code)
															HttpPost_IFTTT(target_item + " " + user_talk + "：" + message.Text + `\n<br>` + userImageUrl + `\n<br>` + userStatus, "LINE 同步：" + send_title + `\n` + strings.Replace(bot_msg,"\n", `\n<br/>`, -1),target_id_code)
															HttpPost_Zapier(target_item + " [" + user_talk + "](" + userImageUrl + ")：" + message.Text + `\n` + userStatus, "LINE 同步：" + send_title + `\n` + bot_msg,target_id_code,user_talk)
														}
													}else{
														// 第一次分割成兩個就成功了
														send_color := "yellow"
														send_title := "查詢成功"
														if bot_msg=="查詢章節超過聖經範圍，有可能指定查詢的節超過範圍。"{
															send_color = "orange"
															send_title = "查詢失敗，範圍超過。"
														}
														HttpPost_JANDI(target_item + " [" + user_talk + "](" + userImageUrl + ")：" + message.Text + `\n` + userStatus, send_color , "LINE 同步：" + send_title + `\n` + bot_msg,target_id_code)
														HttpPost_IFTTT(target_item + " " + user_talk + "：" + message.Text + `\n<br>` + userImageUrl + `\n<br>` + userStatus, "LINE 同步：" + send_title + `\n` + strings.Replace(bot_msg,"\n", `\n<br/>`, -1),target_id_code)
														HttpPost_Zapier(target_item + " [" + user_talk + "](" + userImageUrl + ")：" + message.Text + `\n` + userStatus, "LINE 同步：" + send_title + `\n` + bot_msg,target_id_code,user_talk)
													}
											}
									}else{
										log.Print("bot_msg = " + bot_msg)
										send_color := "yellow"
										send_title := "查詢成功"
										if bot_msg=="查詢章節超過聖經範圍，有可能指定查詢的節超過範圍。"{
											send_color = "orange"
											send_title = "查詢失敗，範圍超過。"
										}
										HttpPost_JANDI(target_item + " [" + user_talk + "](" + userImageUrl + ")：" + message.Text + `\n` + userStatus, send_color , "LINE 同步：" + send_title + `\n` + bot_msg,target_id_code)
										HttpPost_IFTTT(target_item + " " + user_talk + "：" + message.Text + `\n<br>` + userImageUrl + `\n<br>` + userStatus, "LINE 同步：" + send_title + `\n` + strings.Replace(bot_msg,"\n", `\n<br/>`, -1),target_id_code)
										HttpPost_Zapier(target_item + " [" + user_talk + "](" + userImageUrl + ")：" + message.Text + `\n` + userStatus, "LINE 同步：" + send_title + `\n` + bot_msg,target_id_code,user_talk)
									}
									//2017.01.22
									// HttpPost_JANDI(target_item + "[" + user_talk + "](" + userImageUrl + ")：" + message.Text + `\n` + userStatus + `查詢結果：\n` + bot_msg, "yellow" , "查詢成功",target_id_code)
									// HttpPost_IFTTT(target_item + " " + user_talk + "：" + message.Text + `\n<br>` + userImageUrl + `\n<br>` + userStatus + `查詢結果：\n` + bot_msg , "LINE 同步：查詢成功" ,target_id_code)
									// HttpPost_Zapier(target_item + "[" + user_talk + "](" + userImageUrl + ")" + message.Text + `\n` + userStatus, "LINE 程式觀察" ,target_id_code,user_talk)									
								}else{
									if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(bot_msg)).Do(); err != nil {
											log.Print(7557)
											log.Print(err)
											log.Print("linebot: APIError 400 The request body has 1 error(s)\n[messages[2].text] Length must be between 0 and 2000"==err.Error()) //bot_msg 位子換了所以變成 [2]
											// //HttpPost_JANDI(target_item + "[" + user_talk + "](" + userImageUrl + ")：" + message.Text + `\n` + userStatus, "red" , "查詢失敗" + `\n` + err.Error(),target_id_code)
											// //HttpPost_IFTTT(target_item + " " + user_talk + "：" + message.Text + `\n<br>` + userImageUrl + `\n<br>` + userStatus , "LINE 同步：查詢失敗" + `\n<br>` +  err.Error(),target_id_code)

											// if(err.Error()=="linebot: APIError 400 The request body has 1 error(s)\n[messages[2].text] Length must be between 0 and 2000"){
											// 	if _, err = bot.ReplyMessage(event.ReplyToken,linebot.NewStickerMessage("2", "152"),linebot.NewTextMessage("【查詢 "  + message.Text +  " 發生錯誤】\n\n查詢得到的回應超過 2000 字，\n超過 LINE 的承受上限。\n\n請減少查詢節數再查一次。")).Do(); err != nil {
											// 		log.Print(7489)
											// 		log.Print(err.Error())
											// 	}
											// }
											switch err.Error(){
												case "linebot: APIError 400 The request body has 1 error(s)\n[messages[0].text] Length must be between 0 and 2000","linebot: APIError 400 The request body has 1 error(s)\n[messages[1].text] Length must be between 0 and 2000","linebot: APIError 400 The request body has 1 error(s)\n[messages[2].text] Length must be between 0 and 2000":
													log.Print("！！！！！有走進來 case error msg2 must be between 0 and 2000，總數 = ")
													//log.Print(len(bot_msg)) //7121 bytes
													log.Print(len([]rune(bot_msg))) //https://play.golang.org/p/yikJz-BKOW //utf8.RuneLen

													//log.Print("測試 = ")
													//log.Print(len(bot_msg[0:3000]))
													//log.Print(len([]rune(string([]rune(bot_msg)[0:2000]))))
														//部分文字列を取り出す	http://ashitani.jp/golangtips/tips_string.html#string_Replace
													//string([]rune(test_string)[0:len([]rune(test_string))])	https://play.golang.org/p/ivzNYS711B //UTF-8 的顯示部分字串的方法 結合 len([]rune(test_string)) 就得到 UTF-8 版本的字數
													//http://stackoverflow.com/questions/15018545/how-to-index-characters-in-a-golang-string
													if _, err = bot.ReplyMessage(event.ReplyToken,
														linebot.NewTextMessage(string([]rune(bot_msg)[0:2000])),
														linebot.NewTextMessage(string([]rune(bot_msg)[1990:len([]rune(bot_msg))])),
													).Do(); err != nil {
														//第一次挽救：分割成兩個還是失敗，一定是後面那個太胖。
														log.Print(7587)
														log.Print(err.Error())
														if _, err = bot.ReplyMessage(event.ReplyToken,
															linebot.NewTextMessage(string([]rune(bot_msg)[0:2000])),
															linebot.NewTextMessage(string([]rune(bot_msg)[1990:3990])),
															linebot.NewTextMessage(string([]rune(bot_msg)[3980:len([]rune(bot_msg))])),
														).Do(); err != nil {
														//if _, err = bot.ReplyMessage(event.ReplyToken,obj_message2,obj_message,linebot.NewTextMessage(bot_msg[0:3000])).Do(); err != nil {
														//if _, err = bot.ReplyMessage(event.ReplyToken,linebot.NewTextMessage(len(bot_msg)).Do(); err != nil {
															//HttpPost_JANDI(target_item + "[" + user_talk + "](" + userImageUrl + ")：" + message.Text + `\n` + userStatus, "red" , "查詢失敗" + `\n` + err.Error(),target_id_code)
															//HttpPost_IFTTT(target_item + " " + user_talk + "：" + message.Text + `\n<br>` + userImageUrl + `\n<br>` + userStatus , "LINE 同步：查詢失敗" + `\n<br>` +  err.Error(),target_id_code)
															log.Print(7600)
															log.Print(err.Error())
															//連續三次發現失敗(原始內容、分成2、分成3)，勸退XD
															if _, err = bot.ReplyMessage(event.ReplyToken,
																linebot.NewTextMessage(string([]rune(bot_msg)[0:2000])),
																linebot.NewTextMessage(string([]rune(bot_msg)[1990:3990])),
																linebot.NewTextMessage(string([]rune(bot_msg)[3980:5980])),
																linebot.NewStickerMessage("2", "152"),
																linebot.NewTextMessage("【查詢 "  + message.Text +  " 發生錯誤】\n\n查詢得到的回應超過 6000 字，\n超過 LINE 可以傳輸的上限。\n因此內容最後有遺漏。\n\n請減少查詢節的數量，重新再查一次。"),
															).Do(); err != nil {
																log.Print(7604)
																log.Print(err.Error())
															}
														}else{
															// 第二次分割成三個才成功
															send_color := "yellow"
															send_title := "查詢成功"
															if bot_msg=="查詢章節超過聖經範圍，有可能指定查詢的節超過範圍。"{
																send_color = "orange"
																send_title = "查詢失敗，範圍超過。"
															}
															HttpPost_JANDI(target_item + " [" + user_talk + "](" + userImageUrl + ")：" + message.Text + `\n` + userStatus, send_color , "LINE 同步：" + send_title + `\n` + bot_msg,target_id_code)
															HttpPost_IFTTT(target_item + " " + user_talk + "：" + message.Text + `\n<br>` + userImageUrl + `\n<br>` + userStatus, "LINE 同步：" + send_title + `\n` + strings.Replace(bot_msg,"\n", `\n<br/>`, -1),target_id_code)
															HttpPost_Zapier(target_item + " [" + user_talk + "](" + userImageUrl + ")：" + message.Text + `\n` + userStatus, "LINE 同步：" + send_title + `\n` + bot_msg,target_id_code,user_talk)
														}
													}else{
														// 第一次分割成兩個就成功了
														send_color := "yellow"
														send_title := "查詢成功"
														if bot_msg=="查詢章節超過聖經範圍，有可能指定查詢的節超過範圍。"{
															send_color = "orange"
															send_title = "查詢失敗，範圍超過。"
														}
														HttpPost_JANDI(target_item + " [" + user_talk + "](" + userImageUrl + ")：" + message.Text + `\n` + userStatus, send_color , "LINE 同步：" + send_title + `\n` + bot_msg,target_id_code)
														HttpPost_IFTTT(target_item + " " + user_talk + "：" + message.Text + `\n<br>` + userImageUrl + `\n<br>` + userStatus, "LINE 同步：" + send_title + `\n` + strings.Replace(bot_msg,"\n", `\n<br/>`, -1),target_id_code)
														HttpPost_Zapier(target_item + " [" + user_talk + "](" + userImageUrl + ")：" + message.Text + `\n` + userStatus, "LINE 同步：" + send_title + `\n` + bot_msg,target_id_code,user_talk)
													}
											}
									}else{
										log.Print("bot_msg = " + bot_msg)
										send_color := "yellow"
										send_title := "查詢成功"
										if bot_msg=="查詢章節超過聖經範圍，有可能指定查詢的節超過範圍。"{
											send_color = "orange"
											send_title = "查詢失敗，範圍超過。"
										}
										HttpPost_JANDI(target_item + " [" + user_talk + "](" + userImageUrl + ")：" + message.Text + `\n` + userStatus, send_color , "LINE 同步：" + send_title + `\n` + bot_msg,target_id_code)
										HttpPost_IFTTT(target_item + " " + user_talk + "：" + message.Text + `\n<br>` + userImageUrl + `\n<br>` + userStatus, "LINE 同步：" + send_title + `\n` + strings.Replace(bot_msg,"\n", `\n<br/>`, -1),target_id_code)
										HttpPost_Zapier(target_item + " [" + user_talk + "](" + userImageUrl + ")：" + message.Text + `\n` + userStatus, "LINE 同步：" + send_title + `\n` + bot_msg,target_id_code,user_talk)
									}
									//2017.01.22
									// HttpPost_JANDI(target_item + "[" + user_talk + "](" + userImageUrl + ")：" + message.Text + `\n` + userStatus + `查詢結果：\n` + bot_msg, "yellow" , "查詢成功",target_id_code)
									// HttpPost_IFTTT(target_item + " " + user_talk + "：" + message.Text + `\n<br>` + userImageUrl + `\n<br>` + userStatus + `查詢結果：\n` + bot_msg , "LINE 同步：查詢成功" ,target_id_code)
									// HttpPost_Zapier(target_item + "[" + user_talk + "](" + userImageUrl + ")" + message.Text + `\n` + userStatus, "LINE 程式觀察" ,target_id_code,user_talk)									
								}
							}else{
								//沒找到 reg_nofind.ReplaceAllString(bot_msg, "$1")=="我還沒學呢..."
								if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(bot_msg)).Do(); err != nil {
										log.Print(7650)
										log.Print(err)
								}
								HttpPost_JANDI(target_item + " [" + user_talk + "](" + userImageUrl + ")：" + message.Text + `\n` + userStatus, "orange" , "LINE 同步：查詢失敗",target_id_code)
								HttpPost_IFTTT(target_item + " " + user_talk + "：" + message.Text + `\n<br>` + userImageUrl + `\n<br>` + userStatus, "LINE 同步：查詢失敗",target_id_code)
								HttpPost_Zapier(target_item + " [" + user_talk + "](" + userImageUrl + ")：" + message.Text + `\n` + userStatus, "LINE 同步：查詢失敗",target_id_code,user_talk)

								//HttpPost_JANDI(target_item + "[" + user_talk + "](" + userImageUrl + ")：" + message.Text + `\n` + userStatus, "red" , "查詢失敗",target_id_code)
								//HttpPost_IFTTT(target_item + " " + user_talk + "：" + message.Text + `\n<br>` + userImageUrl + `\n<br>` + userStatus , "LINE 同步：查詢失敗" ,target_id_code)
								//HttpPost_Zapier("[" + user_talk + "](" + userImageUrl + ") 觸發了按鈕並呼了 event.Postback.Data = " + message.Text + `\n` + userStatus, "LINE 程式觀察" ,target_id_code,user_talk)
							}
					}
					//2016.12.22+ 利用正則分析字串結果，來設置觸發找開發者的時候要 + 的 UI  //不能用 bot_msg == 開發者，因為 bot_msg 早就被改寫成一串廢話。
					// if reg_loking_for_admin.ReplaceAllString(bot_msg,"$1") == "你找我的製造者？OK！"{

					// }





					//因為 bot_msg==GOTEST 的時候，不可能會找到 anime_url。所以不用在 else 裡面。
					// if anime_url!=""{
					// 	//找到的時候的 UI
					//     imageURL = "https://i2.bahamut.com.tw/anime/FB_anime.png"
					// 	template := linebot.NewCarouselTemplate(
					// 		linebot.NewCarouselColumn(
					// 			imageURL, "動畫搜尋結果", "在找" + message.Text + "對吧！？\n建議可以直接在巴哈姆特動畫瘋 APP 裡面播放！",							
					// 			linebot.NewURITemplateAction("點此播放找到的動畫", anime_url),
					// 			LineTemplate_download_app,
					// 			linebot.NewMessageTemplateAction("查詢其他動畫", "目錄"),
					// 		),
					// 		LineTemplate_CarouselColumn_feedback,
					// 		// LineTemplate_other_example,
					// 		// LineTemplate_other,
					// 	)
					// 	obj_message := linebot.NewTemplateMessage(bot_msg, template)

								// 	originalContentURL_1 := "https://trello-attachments.s3.amazonaws.com/52ff05f27a3c676c046c37f9/5831e5e304f9fac88ac50a23/722268f159dc640ed1639ffd31b4dd0d/94455.jpg"
				   	// 				previewImageURL_1 := "https://trello-attachments.s3.amazonaws.com/52ff05f27a3c676c046c37f9/5831e5e304f9fac88ac50a23/722268f159dc640ed1639ffd31b4dd0d/94455.jpg"
				   	// 				obj_message_img_1 := linebot.NewImageMessage(originalContentURL_1, previewImageURL_1)

								// 	originalContentURL_2 := "https://trello-attachments.s3.amazonaws.com/52ff05f27a3c676c046c37f9/5831e5e304f9fac88ac50a23/f7e158cdc3f1e9640a5f5cf188c33b13/94454.jpg"
				   	// 				previewImageURL_2 := "https://trello-attachments.s3.amazonaws.com/52ff05f27a3c676c046c37f9/5831e5e304f9fac88ac50a23/f7e158cdc3f1e9640a5f5cf188c33b13/94454.jpg"
				   	// 				obj_message_img_2 := linebot.NewImageMessage(originalContentURL_2, previewImageURL_2)

					// 	if _, err = bot.ReplyMessage(event.ReplyToken,linebot.NewTextMessage("可參考以下圖例操作讓搜尋到的影片，直接在巴哈姆特動畫瘋 APP 進行播放。"),obj_message_img_1,obj_message_img_2,obj_message).Do(); err != nil {
					// 		log.Print(1724)
					// 		log.Print(err)
					// 	}
					// 	//HttpPost_JANDI(target_item + " [" + user_talk + "](" + userImageUrl + ")：" + message.Text + `\n` + userStatus, "yellow" , "LINE 同步：查詢成功" + `\n` + anime_url,target_id_code)
					// 	//HttpPost_IFTTT(target_item + " " + user_talk + "：" + message.Text + `\n<br>` + userImageUrl + `\n<br>` + userStatus, "LINE 同步：查詢成功" + `\n` + anime_url,target_id_code)
					// 	HttpPost_Zapier(target_item + " [" + user_talk + "](" + userImageUrl + ")：" + message.Text + `\n` + userStatus, "LINE 同步：查詢成功" + `\n` + anime_url,target_id_code,user_talk)
					// 	log.Print("target_id_code +  anime_url = " + target_id_code + "\n" + anime_url)
					// }else{
					// 	//2016.12.22+ 利用正則分析字串結果，來設置觸發找不到的時候要 + 的 UI
					// 	if reg_nofind.ReplaceAllString(bot_msg,"$1") == "才會增加比較慢XD）"{
					// 		//找不到的時候
					// 		imageURL = "https://i2.bahamut.com.tw/anime/FB_anime.png"
					// 		template := linebot.NewCarouselTemplate(
					// 			linebot.NewCarouselColumn(
					// 				imageURL, "找不到 "  +  message.Text   +   " 耶", "有可能打錯字或這真的沒有收錄，\n才會找不到。",							
					// 				linebot.NewMessageTemplateAction("查看新番", "新番"),
					// 				linebot.NewMessageTemplateAction("可查詢的其他動畫目錄", "目錄"),
					// 				LineTemplate_download_app,
					// 			),
					// 			LineTemplate_CarouselColumn_feedback,
					// 			// LineTemplate_other_example,
					// 			// LineTemplate_other,
					// 		)
					// 		obj_message := linebot.NewTemplateMessage("除了「目錄」以外，\n你也可以輸入「新番」查詢近期的動畫。", template)
					// 		if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(bot_msg),obj_message).Do(); err != nil {
					// 			log.Print(1763)
					// 			log.Print(err)
					// 		}
					// 		//HttpPost_JANDI(target_item + " [" + user_talk + "](" + userImageUrl + ")：" + message.Text + `\n` + userStatus, "orange" , "LINE 同步：查詢失敗",target_id_code)
					// 		//HttpPost_IFTTT(target_item + " " + user_talk + "：" + message.Text + `\n<br>` + userImageUrl + `\n<br>` + userStatus, "LINE 同步：查詢失敗",target_id_code)
					// 		HttpPost_Zapier(target_item + " [" + user_talk + "](" + userImageUrl + ")：" + message.Text + `\n` + userStatus, "LINE 同步：查詢失敗",target_id_code,user_talk)
					// 	}else{
					// 		//這是最原始的動作部分，還沒改寫 UI 模式的時候就靠這裡直接回傳結果就好。至於要傳什麼內容已經在 bible() 裡面處理好了。
					 		// if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(bot_msg)).Do(); err != nil {
					 		// 	log.Print(1771)
					 		// 	log.Print(err)
					 		// }
					// 	}
					// }
				}
					// 				m := linebot.NewTextMessage("ok")
					// 				    if _, err = bot.ReplyMessage(event.ReplyToken, m).Do(); err != nil {

					// 				    }
									
									//----------PushMessage-----------這段可以跟 ReplyMessage 同時有效，但是只會在 1 對 1 有效。群組無效。---------
									//------開發者測試方案有效(好友最多50人/訊息無上限)，免費版(好友不限人數/訊息限制1000)、入門版無效，旗艦版、專業版有效。
									
									//http://muzigram.muzigen.net/2016/09/linebot-golang-linebot-heroku.html
									//https://github.com/mogeta/lbot/blob/master/main.go
					 		// source := event.Source
					 		// log.Print("source.UserID = " + source.UserID)
					 		// log.Print("target_id_code = " + target_id_code)
									//2016.12.20+//push_string := ""
					// 				if source.UserID == "U6f738a70b63c5900aa2c0cbbe0af91c4"{
					// 					push_string = "你好，主人。（PUSH_MESSAGE 才可以發）"
					// 				}
					// 				if source.UserID == "Uf150a9f2763f5c6e18ce4d706681af7f"{
					// 					push_string = "唉呦，你是包包吼"
					// 				}
					//2016.12.20+ close push
					// 					if source.Type == linebot.EventSourceTypeUser {
					// 						if _, err = bot.PushMessage(source.UserID, linebot.NewTextMessage(push_string)).Do(); err != nil {
					// 							log.Print(err)
					// 						}
					// 					}
					// 					if source.Type == linebot.EventSourceTypeUser {
					// 						if _, err = bot.PushMessage(source.UserID, linebot.NewTextMessage(push_string)).Do(); err != nil {
					// 							log.Print(err)
					// 						}
					// 					}
						//上面重覆兩段 push 用來證明 push 才可以連發訊息框，re 只能一個框
					//---------------------這段可以跟 ReplyMessage 同時有效，但是只會在 1 對 1 有效。群組無效。---------
			case *linebot.ImageMessage:
				// 				_, err := bot.SendText([]string{event.RawContent.Params[0]}, "Hi~\n歡迎加入 Delicious!\n\n想查詢附近或各地美食都可以LINE我呦！\n\n請問你想吃什麼?\nex:義大利麵\n\n想不到吃什麼，也可以直接'傳送目前位置訊息'")
				// 				var img = "http://imageshack.com/a/img921/318/DC21al.png"
				// 				_, err = bot.SendImage([]string{content.From}, img, img)
				// 				if err != nil {
				// 					log.Println(err)
				// 				}
									
				// 				if err := bot.handleImage(message, event.ReplyToken); err != nil {
				// 					log.Print(err)
				// 				}
									//https://devdocs.line.me/en/#webhook-event-object
				log.Print("對方丟圖片 message.ID = " + message.ID)

				//log.Print("對方丟圖片 linebot.EventSource = " + linebot.EventSource

				//----------------------------------------------------------------取得使用者資訊的寫法
				// source := event.Source

				// userID := event.Source.UserID
				// groupID := event.Source.GroupID
				// RoomID := event.Source.RoomID
				// markID := userID + groupID + RoomID
				
				// log.Print(source.UserID)
				//----------------------------------------------------------------取得使用者資訊的寫法

				// username := ""
				// if markID == "U6f738a70b63c5900aa2c0cbbe0af91c4"{//if source.UserID == "U6f738a70b63c5900aa2c0cbbe0af91c4"{
				// 	username = "LL = " + userID + groupID + RoomID //2016.12.20+
				// }
				// if markID == "Uf150a9f2763f5c6e18ce4d706681af7f"{
				// 	username = "包包"
				// }

				//https://devdocs.line.me/en/#get-content
				//[GAE/GoでLineBotをつくったよ〜 - ベーコンの裏](http://sun-bacon.hatenablog.com/entry/2016/10/10/233520)
				content, err := bot.GetMessageContent(message.ID).Do()
				if err != nil {
					log.Print(2141)
					log.Print(err)
				}
				defer content.Content.Close()
				log.Print("content.ContentType = " + content.ContentType)
				log.Print("content.ContentLength = ")
				log.Print(content.ContentLength) //檔案大小??
				log.Print("content.Content = ")
				log.Print(content.Content)

				//https://github.com/line/line-bot-sdk-go/blob/master/linebot/get_content_test.go
				//ContentLength
				//https://golang.org/pkg/image/jpeg/

				//目標是把 content.Content 存起來

                image, err := jpeg.Decode(content.Content)
                if err != nil {
                	log.Print(2167)
                    log.Print(err)
                }
                log.Printf("image %v", image.Bounds())
                //http://ithelp.ithome.com.tw/articles/10161612
                //https://webcache.googleusercontent.com/search?q=cache:cLTwZS5RNmMJ:https://libraries.io/go/github.com%252Fline%252Fline-bot-sdk-go%252Flinebot+&cd=6&hl=zh-TW&ct=clnk&gl=tw



				var imgByte []byte
				imgByte, err = ioutil.ReadAll(content.Content)
				if err != nil {
					log.Print(err)
				}

				log.Print(imgByte)

                //暫時放棄 = =

									// file, err := ioutil.TempFile("temp.jpg", "")
									// if err != nil {
									// 	log.Print(2175)
									// 	log.Print(err)
									// }
									// defer file.Close()
									
									// _, err = ioutil.WriteFile("temp.jpg", []byte(image.Bounds()), 0600)//io.Copy(file, content.Content)
									// if err != nil {
									// 	log.Print(2182)
									// 	log.Print(err)
									// }
									// log.Printf("Saved %s", file.Name())


                //可以
				// if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("這圖片是？\n\n" + username + "你丟給我圖片幹嘛！\n我眼睛還沒長好看不懂XD")).Do(); err != nil {
				// 	log.Print(1845)
				// 	log.Print(err)
				// }
			case *linebot.VideoMessage:
				// //https://github.com/dongri/line-bot-sdk-go
			 //    originalContentURL := "https://dl.dropboxusercontent.com/u/358152/linebot/resource/video-original.mp4"
			 //    previewImageURL := "https://dl.dropboxusercontent.com/u/358152/linebot/resource/video-preview.png"
			 //    obj_message := linebot.NewVideoMessage(originalContentURL, previewImageURL)
 			// 	if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("這影片是？\n我也給你影片吧！\n\n這只是測試功能"),obj_message).Do(); err != nil {
 			// 		log.Print(1854)
 			// 		log.Print(err)
 			// 	}
			case *linebot.AudioMessage:
				// //下面都是 OK 的寫法，但是還是沒辦法取得...........
				// //另外因為現在這個專案不適合這樣玩
				// originalContentURL := "https://dl.dropboxusercontent.com/u/358152/linebot/resource/ok.m4a"
				// duration := 1000
				// obj_message := linebot.NewAudioMessage(originalContentURL, duration)
 			// 	if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("這是什麼聲音？"),obj_message).Do(); err != nil {
 			// 		log.Print(1862)
 			// 		log.Print(err)
 			// 	}
			case *linebot.LocationMessage:
				log.Print("message.Title = " + message.Title)
				log.Print("message.Address = " + message.Address)
				log.Print("message.Latitude = ")
				log.Print(message.Latitude)
				log.Print("message.Longitude = ")
				log.Print(message.Longitude)

	
				target_x_num := 25.007408
				target_y_num := 121.537688
				//obj_message := linebot.NewLocationMessage(message.Title, message.Address, message.Latitude, message.Longitude)
				obj_message_map := linebot.NewLocationMessage("台北公館教會", "11677 台北市汀州路四段85巷2號", target_x_num, target_y_num) //台北市信義區富陽街46號

				target_x := "25.007408"
				target_y := "121.537688"
				//obj_message_1 := linebot.NewLocationMessage("歡迎光臨", "地球", 25.022413, 121.556427) //台北市信義區富陽街46號
					//obj_message_2 := linebot.NewLocationMessage("歡迎光臨", "哪個近", 25.022463, 121.556454) //這個遠

				//if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("你好！小天使來幫你帶路！\n你在這裡？\n要看看附近的教會嗎？\nhttps://www.google.com/maps/search/%E6%95%99%E6%9C%83/@" + fmt.Sprintf("%f",message.Latitude) + "%2C" + fmt.Sprintf("%f",message.Longitude) + ",15z\n以下是我剛剛收到的 GPS 定位地址！"),obj_message,linebot.NewTextMessage("我們教會在這裡～\n為您預備導航路線圖：\nhttp://maps.google.com/maps?f=d&saddr=" + fmt.Sprintf("%f",message.Latitude) + "," + fmt.Sprintf("%f",message.Longitude) + "&daddr=" + target_x + "," + target_y + "&hl=zh-tw&dirflg=&sort=num&mrsp=0&doflg=ptk&ttype=now\n以下是我們教會的確切地址！"),obj_message_map,linebot.NewStickerMessage("2", "514")).Do(); err != nil {
				if target_item == "群組對話"{
					LineTemplate_address := linebot.NewCarouselTemplate(
						linebot.NewCarouselColumn(
							imageURL, "我來幫你帶路！", "你在「" + message.Address + "」？",
//							linebot.NewURITemplateAction("看看附近的教會？", "https://www.google.com/maps/search/%E6%95%99%E6%9C%83/@" + fmt.Sprintf("%f",message.Latitude) + "%2C" + fmt.Sprintf("%f",message.Longitude) + ",15z"),
							linebot.NewURITemplateAction("帶你去公館教會","http://maps.google.com/maps?f=d&saddr=" + fmt.Sprintf("%f",message.Latitude) + "," + fmt.Sprintf("%f",message.Longitude) + "&daddr=" + target_x + "," + target_y + "&hl=zh-tw&dirflg=&sort=num&mrsp=0&doflg=ptk&ttype=now"),
							linebot.NewPostbackTemplateAction("公館教會地圖","公館教會地圖 POST", "教會地圖"),
						),
						// LineTemplate_other_example,
						// LineTemplate_other,
						LineTemplate_CarouselColumn_feedback,
					)
					t_address := "帶你去公館教會！\nhttp://maps.google.com/maps?f=d&saddr=" + fmt.Sprintf("%f",message.Latitude) + "," + fmt.Sprintf("%f",message.Longitude) + "&daddr=" + target_x + "," + target_y + "&hl=zh-tw&dirflg=&sort=num&mrsp=0&doflg=ptk&ttype=now"
					obj_message_address := linebot.NewTemplateMessage(t_address, LineTemplate_address)
					if _, err = bot.ReplyMessage(
						event.ReplyToken, 
						//linebot.NewStickerMessage("2", "514"),
						//linebot.NewTextMessage("你好！小天使來幫你帶路！\n你在「" + message.Address + "」？\n要看看附近的教會嗎？\nhttps://www.google.com/maps/search/%E6%95%99%E6%9C%83/@" + fmt.Sprintf("%f",message.Latitude) + "%2C" + fmt.Sprintf("%f",message.Longitude) + ",15z"),
						// obj_message,
						//linebot.NewTextMessage("我們教會在這裡～\n為您預備導航路線圖：\nhttp://maps.google.com/maps?f=d&saddr=" + fmt.Sprintf("%f",message.Latitude) + "," + fmt.Sprintf("%f",message.Longitude) + "&daddr=" + target_x + "," + target_y + "&hl=zh-tw&dirflg=&sort=num&mrsp=0&doflg=ptk&ttype=now\n下面是我們教會的地圖！"),
						//obj_message_map,
						obj_message_address,
					).Do(); err != nil {
						log.Print(18766)
						log.Print(err)
					}
				}else{
					if _, err = bot.ReplyMessage(
						event.ReplyToken, 
						linebot.NewStickerMessage("2", "514"),
						linebot.NewTextMessage("你好！小天使來幫你帶路！\n你在「" + message.Address + "」？\n要看看附近的教會嗎？\nhttps://www.google.com/maps/search/%E6%95%99%E6%9C%83/@" + fmt.Sprintf("%f",message.Latitude) + "%2C" + fmt.Sprintf("%f",message.Longitude) + ",15z"),
						// obj_message,
						linebot.NewTextMessage("我們教會在這裡～\n為您預備導航路線圖：\nhttp://maps.google.com/maps?f=d&saddr=" + fmt.Sprintf("%f",message.Latitude) + "," + fmt.Sprintf("%f",message.Longitude) + "&daddr=" + target_x + "," + target_y + "&hl=zh-tw&dirflg=&sort=num&mrsp=0&doflg=ptk&ttype=now\n下面是我們教會的地圖！"),
						obj_message_map,
					).Do(); err != nil {
						log.Print(1876)
						log.Print(err)
					}
				}
			case *linebot.StickerMessage:
				log.Print("message.PackageID = " + message.PackageID)
				log.Print("message.StickerID = " + message.StickerID)
					//https://github.com/line/line-bot-sdk-go/blob/master/examples/kitchensink/server.go handleSticker
					//message.PackageID, message.StickerID
				//丟跟對方一樣的貼圖回他
			// obj_message_moto := linebot.NewStickerMessage(message.PackageID, message.StickerID)
					//https://github.com/line/line-bot-sdk-go/blob/master/examples/kitchensink/server.go
					//2016.12.20+ 多次框框的方式成功！（最多可以五個）
					//.NewStickerMessage 發貼貼圖成功	 //https://devdocs.line.me/files/sticker_list.pdf			
				obj_message := linebot.NewStickerMessage("2", "514") //https://devdocs.line.me/en/?go#send-message-object
 				//if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("OU<"),linebot.NewTextMessage("0.0"),linebot.NewTextMessage("．ω．"),linebot.NewTextMessage("．ω．")).Do(); err != nil {

				PackageID_int := 0
				StickerID_int := 0
				if PackageID_int, err = strconv.Atoi(message.PackageID); err != nil {
					log.Print("7793 字串轉整數失敗")
					log.Print(PackageID_int)
					log.Print(err.Error())
				}

				if StickerID_int, err = strconv.Atoi(message.StickerID); err != nil {
					log.Print("7798 字串轉整數失敗")
					log.Print(StickerID_int)
					log.Print(err.Error())
				}

				//特別處理過貼圖範圍外的貼圖
				if (PackageID_int!=0) && (PackageID_int<=4){
					// if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("．ω．"),obj_message_moto,obj_message).Do(); err != nil {
					// 	log.Print(7806)
					// 	log.Print(err)
					// }
					if _, err = bot.ReplyMessage(event.ReplyToken, obj_message).Do(); err != nil {
						log.Print(7806)
						log.Print(err)
					}
				}else{
					// if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("．ω．"),obj_message).Do(); err != nil {
					// 	log.Print(7811)
					// 	log.Print(err)
					// }
					if(username=="LL"){
						//https://store.line.me/stickershop/product/1021884/zh-Hant
						//if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("PackageID_int = " + strconv.Itoa(PackageID_int) + "\nStickerID_int = " + strconv.Itoa(StickerID_int)), obj_message).Do(); err != nil {
						if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("https://store.line.me/stickershop/product/" + strconv.Itoa(PackageID_int) + "/zh-Hant"), obj_message).Do(); err != nil {	
							log.Print(7806)
							log.Print(err)
						}
					}else{

					}
				}

					// if _, err = bot.ReplyMessage(event.ReplyToken, obj_message).Do(); err != nil {
					// 	log.Print(7806)
					// 	log.Print(err)
					// }
			}
		}
	}
}
