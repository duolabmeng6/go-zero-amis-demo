package logic

import (
	"context"

	"myapi/article/internal/svc"
	"myapi/article/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ArticleEditLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewArticleEditLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ArticleEditLogic {
	return &ArticleEditLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ArticleEditLogic) ArticleEdit(req *types.ArticleEditRequest) (resp *types.ArticleEditResponse, err error) {
	// 查询文章
	article, err := l.svcCtx.ArticleModel.FindOne(l.ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &types.ArticleEditResponse{
		Status: 0,
		Msg:    "",
		Data: types.ArticleItems{
			Id:        article.Id,
			Title:     article.Title.String,
			Content:   article.Content.String,
			CreatedAt: article.CreatedAt.Time.String(),
			UpdatedAt: article.CreatedAt.Time.String(),
		},
	}, nil
}
