package school

import (
	"fmt"
	"testing"
)

func TestAes(t *testing.T) {
	cfg := Config{
		AppKey:    "1234567890123456",
		AppSecret: "0123456789123456",
		Ocode:     "",
	}
	sch := NewSchool(&cfg)

	encrypt := sch.CBCEncrypter([]byte("12345"))
	fmt.Println(encrypt)

	decrypt, _ := sch.CBCDecrypter("90b2c15e84cb78e5161f42867807c4bc")
	fmt.Println(decrypt)
}
