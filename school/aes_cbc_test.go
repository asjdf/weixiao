package school

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestAes(t *testing.T) {
	cfg := Config{
		AppKey:    "777777777777",
		AppSecret: "12345678901234567890123456789012",
		Ocode:     "1234567890",
	}
	sch := NewSchool(&cfg)

	// 生成主动更新卡数据的请求
	req := sch.NewCardInfoUpdateReq("balance").AddData("cardNum", "1")
	// 主动更新卡数据
	b, _ := json.Marshal(&req)
	encrypt := sch.CBCEncrypter(b)
	fmt.Println(encrypt)

	decrypt, _ := sch.CBCDecrypter(encrypt)
	fmt.Println(decrypt)
}
