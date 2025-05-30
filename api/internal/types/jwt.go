package types

import "github.com/golang-jwt/jwt/v4"

type CustomClaims struct {
	BaseClaims
	BufferTime int64
	jwt.RegisteredClaims
}

type BaseClaims struct {
	ID          uint64
	Username    string
	AuthorityId uint64
}

