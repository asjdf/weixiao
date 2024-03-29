package util

import (
	"fmt"
	"net/url"
	"testing"
)

func TestSign(t *testing.T) {
	values := url.Values{}
	values.Add("app_key", "1234567890123456")
	values.Add("timestamp", "123456")

	values.Add("school_code", "41330103566")
	sign := Sign("A5B59272FD2E95B21D02B9C9888104F0", values)
	fmt.Println(sign)
}
