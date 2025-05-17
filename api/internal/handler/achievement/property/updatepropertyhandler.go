package property

import (
	"net/http"

	"github.com/bearllfleed/scholar-track/api/internal/logic/achievement/property"
	"github.com/bearllfleed/scholar-track/api/internal/svc"
	"github.com/bearllfleed/scholar-track/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UpdatePropertyHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdatePropertyReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := property.NewUpdatePropertyLogic(r.Context(), svcCtx)
		resp, err := l.UpdateProperty(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
