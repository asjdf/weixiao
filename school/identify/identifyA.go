package identify

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"weixiao/util"
)

type IdentifyInfo struct {
	CardNumber string `json:"card_number"`
	Password   string `json:"password"`
	AppKey     string `json:"app_key"`
	NonceStr   string `json:"nonce_str"`
	Timestamp  int    `json:"timestamp"`
	Sign       string `json:"sign"`
}
type identifyReq struct {
	RawData string `json:"raw_data"`
	AppKey  string `json:"app_key"`
}

func (i Identify) identifyReqHandle(f func(info IdentifyInfo) PeopleDetail) func(c *gin.Context) {
	// https://wiki.weixiao.qq.com/api/school/identityA.html
	return func(c *gin.Context) {
		var (
			balReq       identifyReq
			info         IdentifyInfo
			peopleDetail PeopleDetail
		)

		// json data is ok
		err := c.ShouldBind(&balReq)
		//err := c.Bind(&balReq)
		if err != nil {
			c.JSON(util.MakeErrorReturn("bind error", 40001))
			fmt.Println(err)
			return
		}
		// check the appKey
		if balReq.AppKey != i.Conf.AppKey {
			c.JSON(util.MakeErrorReturn("appkey is not equal", 40002))
			return
		}

		//get the cardNum
		dStr, err := i.CBCDecrypter(balReq.RawData)
		fmt.Println(dStr)
		if err != nil {
			fmt.Println(err)
			c.JSON(util.MakeErrorReturn(err.Error(), 40003))
			return
		}
		err = json.Unmarshal([]byte(dStr), &info)
		if err != nil {
			c.JSON(util.MakeErrorReturn(err.Error(), 40004))
			return
		}

		// get the cardInfo and encryption it
		peopleDetail = f(info)
		cardInfoMarshal, err := json.Marshal(peopleDetail)
		if err != nil {
			c.JSON(util.MakeErrorReturn(err.Error(), 40005))
			return
		}
		cardInfoCBC := i.CBCEncrypter(cardInfoMarshal)
		c.JSON(util.MakeSuccessReturn(cardInfoCBC, balReq.AppKey))
	}
}
