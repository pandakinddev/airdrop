package utils

import (
	"math/rand"
	"time"

	"golang.org/x/oauth2"
)

// Retry ...
func Retry(job func() error, max int) error {
	var err error
	jitter := rand.Intn(5)
	for i := 1; i <= max; i++ {
		if err = job(); err != nil {
			time.Sleep(time.Duration(jitter*100*i) * time.Millisecond)
			continue
		}
		return nil
	}
	return err
}

// FromRefreshToken ...
func FromRefreshToken(token string) *oauth2.Token {
	return &oauth2.Token{
		RefreshToken: token,
	}
}

// FromAccessToken ...
func FromAccessToken(token string) *oauth2.Token {
	return &oauth2.Token{
		AccessToken: token,
	}
}
