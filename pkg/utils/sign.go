package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"errors"
	"math/rand"
	"strings"
	"sync"
	"time"
)

// VerifyHMACSign ...
func VerifyHMACSign(data, mac, secret []byte) bool {
	compSign := hmac.New(sha256.New, secret)

	_, err := compSign.Write(data)
	if err != nil {
		return false
	}
	expectedMac := compSign.Sum(nil)

	return hmac.Equal(mac, expectedMac)
}

// SignWith ...
func SignWith(data, secret []byte) ([]byte, error) {
	compSign := hmac.New(sha256.New, secret)
	_, err := compSign.Write(data)
	if err != nil {
		return nil, errors.New("failed to generate sign")
	}
	return compSign.Sum(nil), nil
}

const (
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var (
	m   sync.Mutex
	src = rand.NewSource(time.Now().UnixNano())
)

// RandBytes ...
func RandBytes(n int) ([]byte, error) {
	d := make([]byte, n)
	m.Lock()
	defer m.Unlock()
	_, err := rand.Read(d)
	if err != nil {
		return nil, err
	}
	return d, nil
}

// RandString ...
func RandString(n int) string {
	sb := strings.Builder{}
	sb.Grow(n)
	m.Lock()
	defer m.Unlock()
	// A rsc.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			sb.WriteByte(letterBytes[idx])
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return sb.String()
}
