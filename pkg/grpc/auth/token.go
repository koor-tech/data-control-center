package auth

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	TokenExpireTime  = 26 * time.Hour
	TokenRenewalTime = 10 * time.Hour
	TokenMaxRenews   = 7
)

type UserInfoClaims struct {
	AccID    uint64 `json:"accid"`
	Username string `json:"usr"`

	jwt.RegisteredClaims
}

type TokenMgr struct {
	jwtSigningKey []byte
}

func NewTokenMgr(jwtSecret string) *TokenMgr {
	return &TokenMgr{
		jwtSigningKey: []byte(strings.TrimSpace(jwtSecret)),
	}
}

func (t *TokenMgr) NewWithClaims(claims *UserInfoClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(t.jwtSigningKey)
}

func (t *TokenMgr) ParseWithClaims(tokenString string) (*UserInfoClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserInfoClaims{}, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return "", errors.New("failed to verify jwt token method")
		}
		return t.jwtSigningKey, nil
	})
	if err != nil {
		return nil, errors.New("failed to parse jwt token")
	}

	claims, ok := token.Claims.(*UserInfoClaims)
	if ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("failed to parse token claims")
}

func BuildTokenClaimsFromAccount(accountID uint64, username string) *UserInfoClaims {
	claims := &UserInfoClaims{
		AccID:    accountID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:   "data-control-center",
			Subject:  strconv.FormatUint(accountID, 10),
			ID:       strconv.FormatUint(accountID, 10),
			Audience: []string{"data-control-center"},
		},
	}
	SetTokenClaimsTimes(claims)

	return claims
}

func SetTokenClaimsTimes(claims *UserInfoClaims) {
	now := time.Now()
	// A usual scenario is to set the expiration time relative to the current time
	claims.RegisteredClaims.ExpiresAt = jwt.NewNumericDate(now.Add(TokenExpireTime))
	claims.RegisteredClaims.IssuedAt = jwt.NewNumericDate(now)
	claims.RegisteredClaims.NotBefore = jwt.NewNumericDate(now)
}
