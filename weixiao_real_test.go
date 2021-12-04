package weixiao

import (
	"fmt"
	"testing"
	"weixiao/school"
)


func TestSchool(t *testing.T) {
	cfg := school.Config{
		AppKey:    "7F7B3C6A55220EE8",
		AppSecret: "0123456789123456",
		Ocode:     "",
	}
	sch := school.NewSchool(&cfg)

	// 生成主动更新卡数据的请求
	req := sch.NewCardInfoUpdateReq("balance").AddData("cardNum", "1")
	// 主动更新卡数据
	err := sch.CardInfoUpdate(req)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}