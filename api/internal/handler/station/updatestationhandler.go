package station

import (
	"net/http"

	"api/internal/logic/station"
	"api/internal/svc"
	"api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 修改站点
func UpdateStationHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.StationInfo
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := station.NewUpdateStationLogic(r.Context(), svcCtx)
		resp, err := l.UpdateStation(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
