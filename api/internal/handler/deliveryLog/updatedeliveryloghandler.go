package deliveryLog

import (
	"net/http"

	"api/internal/logic/deliveryLog"
	"api/internal/svc"
	"api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 修改DeliveryLog
func UpdateDeliveryLogHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeliveryLogInfo
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := deliveryLog.NewUpdateDeliveryLogLogic(r.Context(), svcCtx)
		resp, err := l.UpdateDeliveryLog(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
