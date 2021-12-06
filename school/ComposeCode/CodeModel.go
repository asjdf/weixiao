package ComposeCode

type campusCodePostBody struct {
	AppKey     string `json:"app_key"`     //应用分配的app_key
	Timestamp  int64  `json:"timestamp"`   //当前unix时间戳（秒）
	Nonce      string `json:"nonce"`       //32位随机字符串
	Signature  string `json:"signature"`   //使用签名算法计算出来的数字签名
	Scene      int    `json:"scene"`       //1.门禁 2.消费 3.签到 4.其它
	DeviceNo   string `json:"device_no"`   //设备编号
	Location   string `json:"location"`    //扫码地点
	AuthCode   string `json:"auth_code"`   //扫码获得的动态码
	SchoolCode string `json:"school_code"` //扫码的学校编码
}

type campusCodeResponse struct {
	Code    int64 `json:"code"`
	Offline int64 `json:"offline"`
	User    User  `json:"user"`
}

type User struct {
	CardNumber         string  `json:"card_number"`
	Entrusts           string  `json:"entrusts"`
	Rick               string  `json:"rick"`
	Name               string  `json:"name"`
	IdentityType       string  `json:"identity_type"`
	Grade              string  `json:"grade"`
	College            string  `json:"college"`
	Profession         string  `json:"profession"`
	Class              string  `json:"class"`
	IdentityTitle      string  `json:"identity_title"`
	Gender             int64   `json:"gender"`
	Organization       []int64 `json:"organization"`
	Campus             string  `json:"campus"`
	DormNumber         string  `json:"dorm_number"`
	PhysicalChipNumber string  `json:"physical_chip_number"`
	PhysicalCardNumber string  `json:"physical_card_number"`
	Nation             string  `json:"nation"`
	Birthday           string  `json:"birthday"`
	OriginPlace        string  `json:"origin_place"`
	GraduatedSchool    string  `json:"graduated_school"`
	Address            string  `json:"address"`
	ContactPerson      string  `json:"contact_person"`
	ContactPhone       string  `json:"contact_phone"`
	Email              string  `json:"email"`
	IDCard             string  `json:"id_card"`
	Telephone          string  `json:"telephone"`
	StartAt            string  `json:"start_at"`
	ExpireAt           string  `json:"expire_at"`
	UpdatedAt          string  `json:"updated_at"`
}
