package logic

import (
	"context"
	"database/sql"

	"myapi/article/internal/svc"
	"myapi/article/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ArticleUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewArticleUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ArticleUpdateLogic {
	return &ArticleUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ArticleUpdateLogic) ArticleUpdate(req *types.ArticleUpdateRequest) (resp *types.ArticleUpdateResponse, err error) {
	// 查询文章内容
	article, err := l.svcCtx.ArticleModel.FindOne(l.ctx, req.Id)
	if err != nil {
		return nil, err
	}
	article.Title = sql.NullString{String: req.Title, Valid: true}
	article.Content = sql.NullString{String: req.Content, Valid: true}

	// 更新文章内容
	err = l.svcCtx.ArticleModel.Update(l.ctx, article)
	if err != nil {
		return nil, err
	}

	return &types.ArticleUpdateResponse{
		Status: 0,
		Msg:    "更新成功",
		Data: types.ArticleItems{
			Id:        article.Id,
			Title:     article.Title.String,
			Content:   article.Content.String,
			CreatedAt: article.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: article.CreatedAt.Format("2006-01-02 15:04:05"),
		},
	}, nil
}
