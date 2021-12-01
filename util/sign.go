package util

import (
	"crypto/md5"
	"encoding/hex"
	"net/url"
	"sort"
	"strings"
)

// Sign 微信支付签名算法 用作签名密钥的key参数是APP_SECRET，并非APP_KEY
func Sign(secret string, params url.Values) string {
	paramsStr := params.Encode() + "&key=" + secret

	h := md5.New()
	h.Write([]byte(paramsStr))
	return strings.ToUpper(hex.EncodeToString(h.Sum([]byte(""))))
}

func SignMap(secret string, param map[string]string) string {
	var paramKey []string
	for k := range param {
		paramKey = append(paramKey, k)
	}
	sort.Strings(paramKey)
	paramArray := url.Values{}
	for _, v := range paramKey {
		paramArray.Add(v, param[v])
		// 如果参数的值为空不参与签名 但是为了保证那种空值的存在 在sign中不做判断
	}

	paramsStr := paramArray.Encode() + "&key=" + secret

	h := md5.New()
	h.Write([]byte(paramsStr))
	return strings.ToUpper(hex.EncodeToString(h.Sum([]byte(""))))
}
