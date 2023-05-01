package logic

import (
	"context"
	"database/sql"
	"myapi/model"

	"myapi/article/internal/svc"
	"myapi/article/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ArticleStoreLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewArticleStoreLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ArticleStoreLogic {
	return &ArticleStoreLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ArticleStoreLogic) ArticleStore(req *types.ArticleStoreRequest) (resp *types.ArticleStoreResponse, err error) {
	// 插入数据库
	articleData := new(model.Articles)
	articleData.Title = sql.NullString{String: req.Title, Valid: true}
	articleData.Content = sql.NullString{String: req.Content, Valid: true}
	Result, err := l.svcCtx.ArticleModel.Insert(l.ctx, articleData)
	if err != nil {
		return nil, err
	}
	id, _ := Result.LastInsertId()
	article, err := l.svcCtx.ArticleModel.FindOne(l.ctx, id)
	if err != nil {
		return nil, err
	}
	return &types.ArticleStoreResponse{
		Id:        article.Id,
		Title:     article.Title.String,
		Content:   article.Content.String,
		CreatedAt: article.CreatedAt.Time.Format("2006-01-02 15:04:05"),
		UpdatedAt: article.CreatedAt.Time.Format("2006-01-02 15:04:05"),
	}, nil
}
