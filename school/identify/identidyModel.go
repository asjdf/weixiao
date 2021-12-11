package identify

type PeopleDetail struct {
	CardNumber         string `json:"card_number"`
	Gender             string `json:"gender"`
	HeadImage          string `json:"head_image"`
	Grade              string `json:"grade"`
	College            string `json:"college"`
	Profession         string `json:"profession"`
	Class              string `json:"class"`
	IdentityType       string `json:"identity_type"`
	IdentityTitle      string `json:"identity_title"`
	CardType           string `json:"card_type"`
	IdCard             string `json:"id_card"`
	Telephone          string `json:"telephone"`
	Organization       string `json:"organization"`
	ExpireAt           string `json:"expire_at"`
	StartAt            string `json:"start_at"`
	Campus             string `json:"campus"`
	Employer           string `json:"employer"`
	DormNumber         string `json:"dorm_number"`
	Remark             string `json:"remark"`
	PhysicalChipNumber string `json:"physical_chip_number"`
	PhysicalCardNumber string `json:"physical_card_number"`
	Email              string `json:"email"`
	Qq                 string `json:"qq"`
	OriginPlace        string `json:"origin_place"`
	GraduatedSchool    string `json:"graduated_school"`
	Address            string `json:"address"`
}
