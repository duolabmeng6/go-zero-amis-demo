package model

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ArticlesModel = (*customArticlesModel)(nil)

type (
	// ArticlesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customArticlesModel.
	ArticlesModel interface {
		articlesModel
		Index(ctx context.Context, keywords string, perPage int64, page int64) (*[]Articles, error)
		Count(ctx context.Context, keywords string) (int64, error)
	}

	customArticlesModel struct {
		*defaultArticlesModel
	}
)

// NewArticlesModel returns a model for the database table.
func NewArticlesModel(conn sqlx.SqlConn) ArticlesModel {
	return &customArticlesModel{
		defaultArticlesModel: newArticlesModel(conn),
	}
}

// Index 查询文章列表 perPage 每页显示数量 page 页码 keywords 搜索关键字可能为空
func (m *defaultArticlesModel) Index(ctx context.Context, keywords string, perPage int64, page int64) (*[]Articles, error) {
	var resp []Articles
	query := "select " + articlesRows + " from " + m.table + " LIMIT ?,?"
	logx.Error("????????")
	logx.Error(keywords)

	err := error(nil)
	if keywords != "" {
		query = "select " + articlesRows + " from " + m.table + " where title like ? LIMIT ?,?"
		keywords = "%" + keywords + "%"
		err = m.conn.QueryRowsCtx(ctx, &resp, query, keywords, (page-1)*perPage, perPage)
	} else {
		err = m.conn.QueryRowsCtx(ctx, &resp, query, (page-1)*perPage, perPage)
	}
	if err != nil {
		logx.Error(err)
		return nil, err
	}
	return &resp, nil

}

// 查询总数量
func (m *defaultArticlesModel) Count(ctx context.Context, keywords string) (int64, error) {
	query := ""
	if keywords != "" {
		keywords = "%" + keywords + "%"
		query = "select count(*) from " + m.table + " where title like ?"
	} else {
		query = "select count(*) from " + m.table
	}
	var total int64
	err := m.conn.QueryRowCtx(ctx, &total, query)
	if err != nil {
		logx.Error(err)
		return 0, err
	}
	return total, nil
}
