package handler

import (
	"backend/common/errorx"
	"backend/service/backend/cmd/api/internal/logic/systemDepartment"
	"backend/service/backend/cmd/api/internal/svc"
	"backend/service/backend/cmd/api/internal/types"
	"fmt"
	"github.com/tal-tech/go-zero/rest/httpx"
	"net/http"
)
// 部门管理 list
func SystemDepartmentListHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SystemDepartmentListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, errorx.NewDefaultError(fmt.Sprintf("%v", err), ""))
			return
		}

		l := logic.NewSystemDepartmentListLogic(r.Context(), ctx)
		resp, err := l.SystemDepartmentList(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
