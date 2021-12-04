package school

type balanceReq struct {
	RawData string `json:"raw_data"`
	AppKey  string `json:"app_key"`
}

type xxxx struct {
	CardNumber string  `json:"card_number"`
	Value      float64 `json:"value"`
}

func (s School) balanceReqHandle(func(cards []string) []xxxx) {

}
