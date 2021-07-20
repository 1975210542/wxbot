package wx

import (
	"encoding/json"
	"fmt"
	"github.com/1975210542/wxbot/utils"
	"runtime"
)

type AccessToken struct {
	Token     string
	ExpiresIn int64
}

//获取Token
func GetAccessToken() string {
	return ""
}

type WechatApplication struct {
	ToUser                 string        `json:"touser"`
	ToParty                string        `json:"toparty"`
	ToTag                  string        `json:"totag"`
	MsgType                string        `json:"msgtype"`
	AgentId                int           `json:"agentid"`
	Text                   NotifyContent `json:"text"`
	Safe                   int           `json:"safe"`
	EnableIdTrans          int           `json:"enable_id_trans"`
	EnableDuplicateCheck   int           `json:"enable_duplicate_check"`
	DuplicateCheckInterval int           `json:"duplicate_check_interval"`

	Callers []string `json:"-"`
}
type NotifyContent struct {
	Content string `json:"content"`
}

const Token = "WHIsi4KzmQviCQ0Odcu7fvPij_Wtv1IWIdWqtmolx2Ei-BhmIkj_G9cg8Y8wjDlRwvvYa2P1chyx92MLk-Mc2mLN2WOfeWIOTRM0H7CDmRois_zUv88c2K9l4ISpvexWgYWKDQ4WD5G8p2bPhKSXQy0nEvjjA_BWZtg7cZnK9KIe3OKSdRP8C_MnzE6SzCScvsj5Ps7IZO980GYiYy-CQw"

func NewWechatApplication(toUser string) *WechatApplication {
	//agentId, _ := strconv.Atoi(global.WxSetting.AgentId)
	agentId := 1000002
	data := &WechatApplication{
		ToUser:                 toUser,
		MsgType:                "text",
		AgentId:                agentId,
		Safe:                   0,
		EnableIdTrans:          0,
		EnableDuplicateCheck:   0,
		DuplicateCheckInterval: 0,
	}
	return data
}

//设置当前的整个调用栈信息
func (w *WechatApplication) WithCallersFrames() *WechatApplication {
	maxCallerDepth := 25
	minCallerDepth := 1
	callers := []string{}
	pcs := make([]uintptr, maxCallerDepth)
	depth := runtime.Callers(minCallerDepth, pcs)
	frames := runtime.CallersFrames(pcs[:depth])
	for fram, more := frames.Next(); more; fram, more = frames.Next() {
		callers = append(callers, fmt.Sprintf("%s: %d %s", fram.File, fram.Line, fram.Function))
		if !more {
			break
		}
	}
	w.Callers = callers
	return w
}

func (w *WechatApplication) SendTextMessage(text string) string {
	//url := global.WxSetting.Url
	url := "https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token="

	url += Token
	contentType := "application/json"
	var callers []string
	callers = append(callers, text)
	callers = append(callers, w.Callers...)
	jsonStr, _ := json.Marshal(callers)
	w.Text = NotifyContent{
		Content: string(jsonStr),
	}

	rest, err := utils.HttpPostRequest(url, w, contentType)
	if err != nil {
		fmt.Println("报错:", err)
	}
	fmt.Println("报错:", err)
	return rest
}
