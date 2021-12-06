package composeCard

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"testing"
	"weixiao/school"
)

// 把cards 转换到cardinfo过程自己写，这里只是个example
func handler(cards []string) []CardInfo {
	cardInfos := make([]CardInfo, 0)
	for k, v := range cards {
		cardInfo := new(CardInfo)
		cardInfo.CardNumber = v
		cardInfo.Value = float64(k)
		cardInfos = append(cardInfos, *cardInfo)
	}
	return cardInfos
}

func TestBalance(t *testing.T) {
	cfg := school.Config{
		AppKey:    "1234567890123456",
		AppSecret: "0123456789123456",
		Ocode:     "",
	}
	c := ComposeCard{school.NewSchool(&cfg)}
	CbcCardNum(c.School)
	CBCDecrypterCardNum(c.School)
	f := c.balanceReqHandle(handler)
	if f == nil {
		return
	}
	initRouter(f)

}

func initRouter(f func(c *gin.Context)) {
	r := gin.Default()
	r.POST("/balance", f)
	err := r.Run()
	if err != nil {
		return
	}
}

// AesCbcCardNum 得到加密后的字符串
func CbcCardNum(sch *school.School) string {

	var cardNum = []string{"123", "231", "41251"}
	marshalCard, err := json.Marshal(cardNum)
	if err != nil {
		return ""
	}
	fmt.Println(sch.CBCEncrypter(marshalCard))
	return sch.CBCEncrypter(marshalCard)

}

func CBCDecrypterCardNum(sch *school.School) {
	var text string
	text = "224a00a526bc619739f8048c33dd6c9465e1973885099c7b687a3ae66f35d645b348f5106c368b02c20d089e61505f107359be2c31233dc80b17b8522a5f982a009e8080c23b17382230dae69a774f2124018235f62f4cfd23aa92f118b4eae0c5853eca47006652c099f5d79e0423e7"
	cardInfo, err := sch.CBCDecrypter(text)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(cardInfo)
}
