package test

import (
	"gorm.io/gorm"
	"time"
)

type RepositoryTest struct {
	gorm.Model   `json:"-"`
	AccessToken  string `json:"access_token"`
	User         User
	RefreshToken string    `json:"refresh_token"`
	ExpiresAt    time.Time `json:"expires_at"`
	UserID       int       `json:"user_id"`
	IP           string    `json:"ip"`
}
