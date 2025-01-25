package app

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/suisbuds/miao/global"
	"github.com/suisbuds/miao/pkg/errcode"
	"github.com/suisbuds/miao/pkg/util"
)

type Claims struct {
	AppKey    string `json:"app_key"`
	AppSecret string `json:"app_secret"`
	jwt.RegisteredClaims
}

// 获得密钥
func GetJWTSecret() []byte {
	return []byte(global.JWTSetting.Secret)
}

// 生成 token
func GenerateToken(appKey, appSecret string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(global.JWTSetting.Timeout)
	claims := Claims{
		AppKey:    util.EncodeSHA256(appKey),
		AppSecret: util.EncodeSHA256(appSecret),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime),
			Issuer:    global.JWTSetting.Issuer,
			IssuedAt:  jwt.NewNumericDate(nowTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(GetJWTSecret())
	return tokenString, err
}

// 解析和校验 token
func ParseToken(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// HMAC 签名: HS256
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errcode.UnauthorizedTokenInvalidSigningMethod
		}
		return GetJWTSecret(), nil
	})
	if err != nil {
		return nil, err
	}
	if token.Valid {
		// token 有效, 返回解析后的 claims
		return claims, nil
	}
	return nil, errcode.UnauthorizedTokenError
}
