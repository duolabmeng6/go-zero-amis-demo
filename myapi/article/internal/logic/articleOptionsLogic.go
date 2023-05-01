package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"myapi/article/internal/svc"
)

type ArticleOptionsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewArticleOptionsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ArticleOptionsLogic {
	return &ArticleOptionsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ArticleOptionsLogic) ArticleOptions() error {
	// todo: add your logic here and delete this line

	return nil
}
