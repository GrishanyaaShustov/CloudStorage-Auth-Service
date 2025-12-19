package config

type CookiesConfig struct {
	AccessCookieName string `yaml:"access_cookie_name"`
	SecureCookies    bool   `yaml:"secure_cookies"`
}
