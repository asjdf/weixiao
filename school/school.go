package school

type Config struct {
	AppKey    string
	AppSecret string
	Ocode     string
}

type School struct {
	conf *Config
	AccessTokenHandle
}

func NewSchool(cfg *Config) *School {
	defaultAkHandle := NewAccessTokenHandle(cfg)
	return &School{conf: cfg, AccessTokenHandle: defaultAkHandle}
}

