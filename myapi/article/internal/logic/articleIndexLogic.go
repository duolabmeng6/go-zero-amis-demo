package logic

import (
	"context"
	"myapi/article/internal/svc"
	"myapi/article/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ArticleIndexLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewArticleIndexLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ArticleIndexLogic {
	return &ArticleIndexLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ArticleIndexLogic) ArticleIndex(req *types.ArticleIndexRequest) (resp *types.ArticleIndexResponse, err error) {
	articleItems, err := l.svcCtx.ArticleModel.Index(l.ctx, req.Keywords, req.PerPage, req.Page, req.OrderBy, req.OrderDir)
	if err != nil {
		return nil, err
	}

	// 构造数据返回
	var articles []types.ArticleItems
	for _, item := range *articleItems {
		articles = append(articles, types.ArticleItems{
			Id:        item.Id,
			Title:     item.Title.String,
			Content:   item.Content.String,
			CreatedAt: item.CreatedAt.Time.Format("2006-01-02 15:04:05"),
			UpdatedAt: item.CreatedAt.Time.Format("2006-01-02 15:04:05"),
		})
	}
	//获取总数量
	total, err := l.svcCtx.ArticleModel.Count(l.ctx, req.Keywords)

	return &types.ArticleIndexResponse{
		Items: articles,
		Total: total,
	}, nil
}
