package ComposeCode

import (
	"github.com/parnurzeal/gorequest"
	"net/url"
	"time"
)

type campusCodePostBody struct {
	AppKey     string `json:"app_key"`
	Timestamp  int64  `json:"timestamp"`
	Nonce      string `json:"nonce"`
	Signature  string `json:"signature"`
	SchoolCode string `json:"school_code"`
	AuthCode   string `json:"auth_code"`
	Scene      int    `json:"scene"`
	DeviceNo   string `json:"device_no"`
	Location   string `json:"location"`
}

func (s school.School) GetUserDataByCampusCode(scene int, nonce, deviceNo, location, authCode string) {
	//https://wiki.weixiao.qq.com/api/school/campuscode.html
	requestUrl := url.URL{
		Scheme: "https",
		Host:   "weixiao.qq.com",
		Path:   "/apps/school-api/campus-code",
	}
	postBody := campusCodePostBody{
		AppKey:     s.Conf.AppKey,
		Timestamp:  time.Now().Unix(),
		Nonce:      nonce,
		Scene:      scene,
		DeviceNo:   deviceNo,
		Location:   location,
		AuthCode:   authCode,
		SchoolCode: s.Conf.Ocode,
	}

	resp := new(campusCodeResponse)
	gorequest.New().
		Post(requestUrl.String()).
		Send(postBody).
		EndStruct(resp)
}
