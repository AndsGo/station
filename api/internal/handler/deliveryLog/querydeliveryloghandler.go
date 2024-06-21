package deliveryLog

import (
	"net/http"

	"api/internal/logic/deliveryLog"
	"api/internal/svc"
	"api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 查询DeliveryLog
func QueryDeliveryLogHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeliveryLogReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := deliveryLog.NewQueryDeliveryLogLogic(r.Context(), svcCtx)
		resp, err := l.QueryDeliveryLog(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
