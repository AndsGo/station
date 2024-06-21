package station

import (
	"net/http"

	"api/internal/logic/station"
	"api/internal/svc"
	"api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 新增站点
func AddStationHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.StationInfo
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := station.NewAddStationLogic(r.Context(), svcCtx)
		resp, err := l.AddStation(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
