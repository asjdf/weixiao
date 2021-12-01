// 用户卡面信息项主动更新
// https://wiki.weixiao.qq.com/api/school/cardInfoUpdate.html

package school

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/parnurzeal/gorequest"
	"strconv"
	"time"
	"weixiao/util"
)

const CardInfoUpdateApi = "https://weixiao.qq.com/apps/v1/data/update-sync-fields"

type ReqCardInfoUpdate struct {
	AppKey    string       `json:"app_key"`    // uni分配的应用APP_KEY
	NonceStr  string       `json:"nonce_str"`  // 32位随机字符串
	Timestamp string       `json:"timestamp"`  // 当前unix时间戳
	CardField string       `json:"card_field"` // 需要更新的信息项: 余额balance,图书借阅borrow,补贴subsidy,餐次mealtimes
	CardInfo  []UpdateInfo `json:"card_info"`  // 对应卡片字段的更新信息项
	Sign      string       `json:"sign"`       // 按照签名算法生成的签名（参考签名算法）
}

type UpdateInfo struct {
	CardNum string `json:"card_number"` // 学生卡的卡号
	Value   string `json:"value"`       // 要更新字段的值
}

type ReqRawData struct {
	RawData string `json:"raw_data"` // 加密后得到的数据
	AppKey  string `json:"app_key"`
}

type ResMsg struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// NewCardInfoUpdateReq 生成更新卡信息请求 cardField为需要更新的信息项: 余额balance,图书借阅borrow,补贴subsidy,餐次mealtimes
func (s School) NewCardInfoUpdateReq(cardField string) *ReqCardInfoUpdate {
	return &ReqCardInfoUpdate{
		AppKey:    s.conf.AppKey,
		CardField: cardField,
		CardInfo:  make([]UpdateInfo, 0),
	}
}

func (r *ReqCardInfoUpdate) AddData(cardNum, value string) *ReqCardInfoUpdate {
	r.CardInfo = append(r.CardInfo, UpdateInfo{
		CardNum: cardNum,
		Value:   value,
	})
	return r
}

// CardInfoUpdate 发送主动更新卡数据的请求 todo：剥离请求部分
func (s *School) CardInfoUpdate(r *ReqCardInfoUpdate) (err error) {
	r.NonceStr = util.RandStr(32)
	r.Timestamp = strconv.FormatInt(time.Now().Unix(), 10)

	rByte, _ := json.Marshal(r)
	fmt.Println(string(rByte))
	encrypt := s.CBCEncrypter(rByte)
	if err != nil {
		return
	}
	req := &ReqRawData{
		RawData: encrypt,
		AppKey:  s.conf.AppKey,
	}
	res := ResMsg{}
	_, _, errs := gorequest.New().Post(CardInfoUpdateApi).Send(req).EndStruct(&res)
	if errs != nil {
		err = errs[0]
		return
	}
	if res.Errcode != 0 {
		return errors.New(fmt.Sprintf("%v - %v", res.Errcode, res.Errmsg))
	}
	return
}
