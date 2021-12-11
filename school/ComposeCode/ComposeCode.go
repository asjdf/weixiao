package ComposeCode

import (
	"github.com/parnurzeal/gorequest"
	"net/url"
	"time"
	"weixiao/school"
)

type ComposeCode struct {
	*school.School
}

func (s *ComposeCode) GetUserDataByCampusCode(scene int, nonce, deviceNo, location, authCode string) {
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
