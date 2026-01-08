package config

import "time"

type AuthConfig struct {
	JWTsecret    string
	AccessToken  time.Duration
	RefreshToken time.Duration
}
