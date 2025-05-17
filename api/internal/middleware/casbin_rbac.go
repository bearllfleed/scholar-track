package middleware

import (
	"github.com/bearllfleed/scholar-track/pkg/response"
	"net/http"
	"strconv"

	"github.com/bearllfleed/scholar-track/api/internal/utils"
)

type CasbinRpcService interface {
	Enforce(r *http.Request, role string, path string, method string) (bool, error)
}

type CasbinMiddleware struct {
	CasbinRpc CasbinRpcService
}

func NewCasbinMiddleware(casbinRpc CasbinRpcService) *CasbinMiddleware {
	return &CasbinMiddleware{
		CasbinRpc: casbinRpc,
	}
}

func (m *CasbinMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		claims, err := utils.GetClaims(r)
		if err != nil {
			response.ErrWithMessage(r.Context(), w, "获取jwt信息失败")
			return
		}
		// 获取当前用户角色
		role := claims.AuthorityId
		// 获取当前请求方法
		method := r.Method
		// 获取当前请求路径
		path := r.URL.Path
		// 判断是否具有权限
		ok, err := m.CasbinRpc.Enforce(r, strconv.Itoa(int(role)), path, method)
		if err != nil {
			response.ErrWithMessage(r.Context(), w, "权限验证失败")
			return
		}
		if !ok {
			response.ErrWithMessage(r.Context(), w, "权限不足")
			return
		}
		next(w, r)
	}
}
