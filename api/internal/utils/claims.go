package utils

import (
	"github.com/bearllflee/scholar-track/api/internal/config"
	"github.com/bearllflee/scholar-track/api/internal/types"
	"github.com/zeromicro/go-zero/core/logx"
	"net"
	"net/http"
)

func ClearToken(w http.ResponseWriter, r *http.Request) {
	// 增加cookie Authorization 向来源的web添加
	host, _, err := net.SplitHostPort(r.Host)
	if err != nil {
		host = r.Host
	}

	if net.ParseIP(host) != nil {
		http.SetCookie(w, &http.Cookie{Name: "Authorization", Value: "", MaxAge: -1, Path: "/", HttpOnly: false, Secure: false})
	} else {
		http.SetCookie(w, &http.Cookie{Name: "Authorization", Value: "", MaxAge: -1, Path: "/", Domain: host, HttpOnly: false, Secure: false})
	}
}

func SetToken(w http.ResponseWriter, r *http.Request, token string, maxAge int) {
	// 增加cookie Authorization 向来源的web添加
	host, _, err := net.SplitHostPort(r.Host)
	if err != nil {
		host = r.Host
	}

	if net.ParseIP(host) != nil {
		http.SetCookie(w, &http.Cookie{Name: "Authorization", Value: token, MaxAge: maxAge, Path: "/", HttpOnly: false, Secure: false})
	} else {
		http.SetCookie(w, &http.Cookie{Name: "Authorization", Value: token, MaxAge: maxAge, Path: "/", Domain: host, HttpOnly: false, Secure: false})
	}
}

func GetToken(r *http.Request) string {
	token, err := r.Cookie("Authorization")
	if err == nil {
		return token.Value
	}
	return r.Header.Get("Authorization")
}

func GetClaims(r *http.Request) (*types.CustomClaims, error) {
	token := GetToken(r)
	jwt := NewJwt([]byte(config.GlobalConfig.JwtConf.SecretKey), config.GlobalConfig.JwtConf.Expire, config.GlobalConfig.JwtConf.Buffer, config.GlobalConfig.JwtConf.Issuer, config.GlobalConfig.JwtConf.Audience)
	claims, err := jwt.ParseToken(token)
	if err != nil {
		logx.Error("从请求中获取jwt解析信息失败, 请检查请求头是否存在Authorization且claims是否为规定结构")
	}
	return claims, err
}

func GetUserId(r *http.Request) uint64 {
	if claims := r.Context().Value("claims"); claims == nil {
		if cl, err := GetClaims(r); err == nil {
			return cl.BaseClaims.ID
		}
		return 0
	} else {
		waitUse := claims.(*types.CustomClaims)
		return waitUse.BaseClaims.ID
	}
}
