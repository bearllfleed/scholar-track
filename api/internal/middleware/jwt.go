package middleware

import (
	"context"
	"errors"
	"github.com/bearllflee/scholar-track/api/internal/config"
	"github.com/bearllflee/scholar-track/pkg/response"
	"net/http"

	"github.com/bearllflee/scholar-track/api/internal/utils"
)

type JwtMiddleware struct {
	Config config.Config
}

func NewJwtMiddleware(c config.Config) *JwtMiddleware {
	return &JwtMiddleware{
		Config: c,
	}
}

func (m *JwtMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := utils.GetToken(r)
		if token == "" {
			response.ErrWithMessage(r.Context(), w, "未登录或非法访问")
			return
		}
		jwt := utils.NewJwt([]byte(m.Config.JwtConf.SecretKey), m.Config.JwtConf.Expire, m.Config.JwtConf.Buffer, m.Config.JwtConf.Issuer, m.Config.JwtConf.Audience)
		claims, err := jwt.ParseToken(token)
		if err != nil {
			if errors.Is(err, utils.ErrTokenExpired) {
				response.ErrWithNoAuth(r.Context(), w, "授权已过期")
				utils.ClearToken(w, r)
				return
			}
			response.ErrWithNoAuth(r.Context(), w, err.Error())
			utils.ClearToken(w, r)
			return
		}
		// 将解析后的用户信息存入上下文
		ctx := context.WithValue(r.Context(), "claims", claims)
		r = r.WithContext(ctx)
		next(w, r)
	}
}
