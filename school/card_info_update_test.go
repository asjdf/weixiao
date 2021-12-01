package school

import (
	"fmt"
	"testing"
)

// 主动更新校园卡信息
func TestSchool_CardInfoUpdate(t *testing.T) {
	cfg := Config{
		AppKey:    "1234567890123456",
		AppSecret: "0123456789123456",
		Ocode:     "",
	}
	sch := NewSchool(&cfg)

	req := sch.NewCardInfoUpdateReq("borrow").AddData("20322230", "1")
	err := sch.CardInfoUpdate(req)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
