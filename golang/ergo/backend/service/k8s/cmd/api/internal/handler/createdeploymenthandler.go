package handler

import (
	"net/http"

	"backend/service/k8s/cmd/api/internal/logic"
	"backend/service/k8s/cmd/api/internal/svc"
	"backend/service/k8s/cmd/api/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func CreateDeploymentHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RequestCreateDeployment
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewCreateDeploymentLogic(r.Context(), ctx)
		resp, err := l.CreateDeployment(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
