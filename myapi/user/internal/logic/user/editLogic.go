package user

import (
	"context"
	"github.com/jinzhu/copier"
	"myapi/common/errorx"
	"myapi/user/internal/logic"

	"myapi/user/internal/svc"
	"myapi/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EditLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEditLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EditLogic {
	return &EditLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EditLogic) Edit(req *types.UserEditReq) (resp *types.UserEditResp, err error) {
	// 获取用户id
	userId := logic.GetUidFromCtx(l.ctx)
	// 获取用户信息
	userInfoResp, err := l.svcCtx.UserModel.FindOne(l.ctx, userId)
	if err != nil {
		return nil, errorx.NewCodeError(1001, "未查询到")
	}
	// 修改用户信息
	userInfoResp.Name = req.Name

	// 写入数据库
	err = l.svcCtx.UserModel.Update(l.ctx, userInfoResp)
	if err != nil {
		return nil, err
	}

	userInfo := types.User{}
	_ = copier.Copy(&userInfo, userInfoResp)
	return &types.UserEditResp{UserInfo: userInfo}, nil

}
