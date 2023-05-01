package logic

import (
	"context"
	"strconv"
	"strings"

	"myapi/article/internal/svc"
	"myapi/article/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ArticleBulkDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewArticleBulkDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ArticleBulkDeleteLogic {
	return &ArticleBulkDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ArticleBulkDeleteLogic) ArticleBulkDelete(req *types.ArticleBulkDeleteRequest) (resp *types.ArticleBulkDeleteResponse, err error) {
	// 批量删除 ids 参数类似于 1,2,3 需要分割为,然后一个一个删除
	ids := req.Ids
	// 分割 ids
	idsArr := strings.Split(ids, ",")
	// 遍历删除
	for _, id := range idsArr {
		// 删除
		idint, _ := strconv.ParseInt(id, 10, 64)
		err = l.svcCtx.ArticleModel.Delete(l.ctx, idint)
		if err != nil {
			return nil, err
		}
	}

	return &types.ArticleBulkDeleteResponse{
		Status: 0,
		Msg:    "删除成功",
	}, nil
}
