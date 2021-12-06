package composeCard

import (
	"fmt"
	"testing"
	"weixiao/school"
)

// 主动更新校园卡信息
func TestSchool_CardInfoUpdate(t *testing.T) {
	cfg := school.Config{
		AppKey:    "1234567890123456",
		AppSecret: "0123456789123456",
		Ocode:     "",
	}
	c := ComposeCard{school.NewSchool(&cfg)}

	req := c.NewCardInfoUpdateReq("borrow").AddData("20322230", "1")
	err := c.CardInfoUpdate(req)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
