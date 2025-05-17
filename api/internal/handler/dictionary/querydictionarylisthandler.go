package dictionary

import (
	"net/http"

	"github.com/bearllflee/scholar-track/api/internal/logic/dictionary"
	"github.com/bearllflee/scholar-track/api/internal/svc"
	"github.com/bearllflee/scholar-track/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func QueryDictionaryListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.QueryDictionaryListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := dictionary.NewQueryDictionaryListLogic(r.Context(), svcCtx)
		resp, err := l.QueryDictionaryList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
