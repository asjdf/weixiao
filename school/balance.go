package school

/*
error code
40001 bind error
40002 appkey is not equal
40003 CBCDecrypter error
40004 Unmarshal error
40005 marshal error

*/
import (
	"encoding/json"
	"fmt"
	"weixiao/util"
)
import "github.com/gin-gonic/gin"

type balanceReq struct {
	RawData string `json:"raw_data"`
	AppKey  string `json:"app_key"`
}

type CardInfo struct {
	CardNumber string  `json:"card_number"`
	Value      float64 `json:"value"`
}

func (s School) balanceReqHandle(f func(cards []string) []CardInfo) func(c *gin.Context) {
	return func(c *gin.Context) {
		var balReq balanceReq
		cardNum := make([]string, 0)
		cardInfo := make([]CardInfo, 0)

		// json data is ok
		err := c.ShouldBind(&balReq)
		//err := c.Bind(&balReq)
		if err != nil {
			c.JSON(util.MakeErrorReturn("bind error", 40001))
			fmt.Println(err)
			return
		}
		// check the appKey
		if balReq.AppKey != s.conf.AppKey {
			c.JSON(util.MakeErrorReturn("appkey is not equal", 40002))
			return
		}

		//get the cardNum
		dStr, err := s.CBCDecrypter(balReq.RawData)
		fmt.Println(dStr)
		if err != nil {
			fmt.Println(err)
			c.JSON(util.MakeErrorReturn(err.Error(), 40003))
			return
		}
		err = json.Unmarshal([]byte(dStr), &cardNum)
		if err != nil {
			c.JSON(util.MakeErrorReturn(err.Error(), 40004))
			return
		}

		// get the cardInfo and encryption it
		cardInfo = f(cardNum)
		cardInfoMarshal, err := json.Marshal(cardInfo)
		if err != nil {
			c.JSON(util.MakeErrorReturn(err.Error(), 40005))
			return
		}
		cardInfoCBC := s.CBCEncrypter(cardInfoMarshal)
		c.JSON(util.MakeSuccessReturn(cardInfoCBC, balReq.AppKey))
	}
}
