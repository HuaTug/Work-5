package utils

import (
	"errors"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
)

var SecretKey = []byte("NigTusg")

type Claims struct {
	Uid       int64  `json:"uid"`
	UserName  string `json:"user_name"`
	ExpiresAt int64  `json:"exp"`
	jwt.StandardClaims
}

// GenerateToken 签发給用户token
func GenerateToken(uid int64, username string) (accessToken, refreshToken string, err error) {
	acExpireTime := time.Hour * 24 * 30
	reExpireTime := time.Hour
	claims := Claims{
		Uid:      uid,
		UserName: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: int64(acExpireTime),
		},
	}

	accessToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(SecretKey)
	if err != nil {
		return "", "", err
	}

	refreshToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: int64(reExpireTime),
	}).SignedString(SecretKey)
	if err != nil {
		return "", "", err
	}
	return
}

// ParseToken 验证用户token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token,
		&Claims{}, func(t *jwt.Token) (interface{}, error) { return SecretKey, nil })
	if tokenClaims != nil {
		if Claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return Claims, nil
		}
	}
	return nil, err
}

// ParseRefreshToken 验证用户token
func ParseRefreshToken(aToken, rToken string) (newAToken, newRToken string,
	err error) {
	accessClaim, err := ParseToken(aToken)
	if err != nil {
		return "", "", nil
	}

	refreshClaim, err := ParseToken(rToken)
	if err != nil {
		return "", "", nil
	}

	if accessClaim.ExpiresAt > time.Now().Unix() {
		return GenerateToken(accessClaim.Uid, accessClaim.UserName)
	}

	if refreshClaim.ExpiresAt > time.Now().Unix() {
		return GenerateToken(accessClaim.Uid, accessClaim.UserName)
	}

	return "", "", errors.New("身份过期，请重新登陆")

}
