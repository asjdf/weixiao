package identify

type IdentifyInfo struct {
	CardNumber string `json:"card_number"`
	Password   string `json:"password"`
	AppKey     string `json:"app_key"`
	NonceStr   string `json:"nonce_str"`
	Timestamp  int    `json:"timestamp"`
	Sign       string `json:"sign"`
}
