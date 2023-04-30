package logic

import (
	"context"

	"myapi/article/internal/svc"
	"myapi/article/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ArticleDestroyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewArticleDestroyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ArticleDestroyLogic {
	return &ArticleDestroyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ArticleDestroyLogic) ArticleDestroy(req *types.ArticleDestroyRequest) (resp *types.ArticleDestroyResponse, err error) {
	// 删除文章
	err = l.svcCtx.ArticleModel.Delete(l.ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &types.ArticleDestroyResponse{
		Status: 0,
		Msg:    "删除成功",
	}, nil
}
