package school

type Config struct {
	AppKey    string
	AppSecret string
	Ocode     string
}

type School struct {
	Conf *Config
	AccessTokenHandle
}

func NewSchool(cfg *Config) *School {
	defaultAkHandle := NewAccessTokenHandle(cfg)
	return &School{Conf: cfg, AccessTokenHandle: defaultAkHandle}
}

