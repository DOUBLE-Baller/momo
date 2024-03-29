package logic

import (
	"backend/common/errorx"
	"backend/service/k8s/cmd/api/internal/svc"
	"backend/service/k8s/cmd/api/internal/types"
	"context"
	"github.com/tal-tech/go-zero/core/logx"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"log"
)

type DeleteIngressLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteIngressLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeleteIngressLogic {
	return DeleteIngressLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteIngressLogic) DeleteIngress(req types.Request) (*types.Response, error) {
	configPath := "k8s/kube.config"

	config, err := clientcmd.BuildConfigFromFlags("", configPath)
	if err != nil {
		log.Fatal(err)
	}
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}
	//检查创建人，创建人才可以删除自己的数据
	//系统数据不能删除
	err = clientSet.ExtensionsV1beta1().Ingresses(req.NameSpace).Delete(context.TODO(), req.Name, metav1.DeleteOptions{})
	return nil, errorx.NewCodeError(200, "成功", err)
}
