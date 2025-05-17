package utils

import (
	"errors"
	"time"

	"github.com/bearllflee/scholar-track/api/internal/types"
	"github.com/golang-jwt/jwt/v4"
)

type JWT struct {
	SecretKey []byte
	Expire    string
	Buffer    string
	Issuer    string
	Audience  string
}

var (
	ErrTokenExpired     = errors.New("token is expired")
	ErrTokenNotValidYet = errors.New("token not active yet")
	ErrTokenMalformed   = errors.New("that's not even a token")
	ErrTokenInvalid     = errors.New("couldn't handle this token")
)

func NewJwt(secretKey []byte, expire string, buffer string, issuer string, audience string) *JWT {
	return &JWT{SecretKey: secretKey, Expire: expire, Buffer: buffer, Issuer: issuer, Audience: audience}
}

func (j *JWT) CreateClaims(baseClaims types.BaseClaims) types.CustomClaims {
	bf, _ := ParseDuration(j.Buffer)
	ep, _ := ParseDuration(j.Expire)
	claims := types.CustomClaims{
		BaseClaims: baseClaims,
		BufferTime: int64(bf / time.Second), // 缓冲时间1天 缓冲时间内会获得新的token刷新令牌 此时一个用户会存在两个有效令牌 但是前端只留一个 另一个会丢失
		RegisteredClaims: jwt.RegisteredClaims{
			Audience:  jwt.ClaimStrings{j.Audience},              // 受众
			NotBefore: jwt.NewNumericDate(time.Now().Add(-1000)), // 签名生效时间
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ep)),    // 过期时间 7天  配置文件
			Issuer:    j.Issuer,                                  // 签名的发行者
		},
	}
	return claims
}

func (j *JWT) GenerateToken(claims *types.CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SecretKey)
}

// 解析 token
func (j *JWT) ParseToken(tokenString string) (*types.CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &types.CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SecretKey, nil
	})
	if err != nil {
		var ve *jwt.ValidationError
		if errors.As(err, &ve) {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, ErrTokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, ErrTokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, ErrTokenNotValidYet
			} else {
				return nil, ErrTokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*types.CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, ErrTokenInvalid

	} else {
		return nil, ErrTokenInvalid
	}
}
