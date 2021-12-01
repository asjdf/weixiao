package school

import (
	"fmt"
	"github.com/parnurzeal/gorequest"
	"sync"
	"time"
	"weixiao/cache"
)

const (
	AccessTokenApi       = "https://open.wecard.qq.com/cgi-bin/oauth2/token"
	CacheKeySchoolPrefix = "weixiao_school_"
)

// AccessTokenHandle 管理AccessToken
type AccessTokenHandle interface {
	Get(forceUpdate bool) (token string, err error)  // 用于获取token（优先读取缓存）
	GetFromServer() (ResT ResAccessToken, err error) // 从服务器获取token（忽略缓存）
}

type ReqAccessToken struct {
	AppKey    string `json:"app_key"`
	AppSecret string `json:"app_secret"`
	GrantType string `json:"grant_type"`
	Scope     string `json:"scope"`
	Ocode     string `json:"ocode"`
}

type ResAccessToken struct {
	ErrorDescription string `json:"error_description"`
	Error            string `json:"error"`

	TokenType   string `json:"token_type"`
	AccessToken string `json:"access_token"`
	Ocode       string `json:"ocode"`      // 学校代号
	ExpiresIn   int    `json:"expires_in"` // 有效期 单位s
}


// DefaultAccessToken 默认AccessToken 获取
type DefaultAccessToken struct {
	appKey          string
	appSecret       string
	ocode           string
	cache           cache.Cache
	accessTokenLock *sync.Mutex
}

// Get 用于获取token（优先读取缓存）
func (t *DefaultAccessToken) Get(force bool) (token string, err error) {
	accessTokenCacheKey := fmt.Sprintf("%s_access_token_%s", CacheKeySchoolPrefix, t.appKey)

	if force == false {
		// 先从cache中取
		if val := t.cache.Get(accessTokenCacheKey); val != nil {
			return val.(string), nil
		}
	}

	// 加上lock，是为了防止在并发获取token时，cache刚好失效，导致从服务器上获取到不同token
	t.accessTokenLock.Lock()
	defer t.accessTokenLock.Unlock()

	// 双检，防止重复从服务器获取
	if val := t.cache.Get(accessTokenCacheKey); val != nil {
		return val.(string), nil
	}

	// cache失效，从微信服务器获取
	var resAccessToken ResAccessToken
	resAccessToken, err = t.GetFromServer()
	if err != nil {
		return
	}

	expires := resAccessToken.ExpiresIn - 1500
	err = t.cache.Set(accessTokenCacheKey, resAccessToken.AccessToken, time.Duration(expires)*time.Second)
	if err != nil {
		return
	}
	token = resAccessToken.AccessToken
	return
}

// GetFromServer 从服务器获取token（忽略缓存）
func (t *DefaultAccessToken) GetFromServer() (resT ResAccessToken, err error) {
	req := ReqAccessToken{
		AppKey:    t.appKey,
		AppSecret: t.appSecret,
		GrantType: "client_credentials",
		Scope:     "base",
		Ocode:     t.ocode,
	}

	_, _, errs := gorequest.New().Post(AccessTokenApi).Send(req).EndStruct(&resT)
	if errs != nil {
		err = errs[0]
		return
	}
	return
}

// NewAccessTokenHandle 快速生成默认 AccessTokenHandle
func NewAccessTokenHandle(conf *Config) AccessTokenHandle {
	memCache := cache.NewMemcache()
	return NewCustomAccessTokenHandle(conf, memCache)
}

// NewCustomAccessTokenHandle 构造 AccessTokenHandle
func NewCustomAccessTokenHandle(conf *Config, cache cache.Cache) AccessTokenHandle {
	if cache == nil {
		panic("cache is needed")
	}
	return &DefaultAccessToken{
		appKey:          conf.AppKey,
		appSecret:       conf.AppSecret,
		ocode:           conf.Ocode,
		cache:           cache,
		accessTokenLock: new(sync.Mutex),
	}
}
