package user

import (
	"context"
	"myapi/common/errorx"
	"myapi/model"
	"myapi/user/internal/logic"
	"myapi/user/internal/svc"
	"myapi/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterResp, err error) {
	// 查询手机号是否已经注册
	_, err = l.svcCtx.UserModel.FindOneByMobile(l.ctx, req.Mobile)
	if err == nil {
		return nil, errorx.NewCodeError(1001, "手机号已经注册")
	}
	// 注册账号 手机号和密码写入数据库
	insertResult, err := l.svcCtx.UserModel.Insert(l.ctx, &model.Users{
		Name:     req.Mobile,
		Password: req.Password,
		Mobile:   req.Mobile,
	})
	if err != nil {
		return nil, err
	}
	// 获取用户id
	userId, err := insertResult.LastInsertId()
	if err != nil {
		return nil, err
	}
	// 生成 AccessToken
	generateTokenLogic := logic.NewGenerateTokenLogic(l.ctx, l.svcCtx)
	tokenResp, err := generateTokenLogic.GenerateToken(userId)
	if err != nil {
		return nil, err
	}
	return &types.RegisterResp{
		AccessToken:  tokenResp.AccessToken,
		AccessExpire: tokenResp.AccessExpire,
		RefreshAfter: tokenResp.RefreshAfter,
	}, nil
}
