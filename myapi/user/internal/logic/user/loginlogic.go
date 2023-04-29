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

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {

	userId, err := l.loginByMobile(req.Mobile, req.Password)
	if err != nil {
		return nil, err
	}
	generateTokenLogic := logic.NewGenerateTokenLogic(l.ctx, l.svcCtx)
	tokenResp, err := generateTokenLogic.GenerateToken(userId)
	return &types.LoginResp{
		AccessToken:  tokenResp.AccessToken,
		AccessExpire: tokenResp.AccessExpire,
		RefreshAfter: tokenResp.RefreshAfter,
	}, nil
}

func (l *LoginLogic) loginByMobile(mobile, password string) (int64, error) {

	user, err := l.svcCtx.UserModel.FindOneByMobile(l.ctx, mobile)
	if err != nil && err != model.ErrNotFound {
		return 0, errorx.NewCodeError(1001, "根据手机号查询用户信息失败 "+mobile)
	}
	if user == nil {
		return 0, errorx.NewCodeError(1002, "手机号不存在 "+mobile)
	}

	if !(password == user.Password) {
		return 0, errorx.NewCodeError(1003, "密码匹配出错")
	}

	return user.Id, nil
}

func (l *LoginLogic) loginBySmallWx() error {
	return nil
}
