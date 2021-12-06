package util

type ErrorReturn struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type SuccessReturn struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	RawData string `json:"raw_data"`
	AppKey  string `json:"app_key"`
}

func MakeSuccessReturn(data interface{}, appKey string) (int, interface{}) {
	return 200, SuccessReturn{
		Code:    0,
		Message: "success",
		RawData: data.(string),
		AppKey:  appKey,
	}

}

func MakeErrorReturn(msg string, msgCode int) (int, interface{}) {
	return 400, ErrorReturn{
		Code:    msgCode,
		Message: msg,
	}
}
