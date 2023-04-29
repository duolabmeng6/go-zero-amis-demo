package user

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	"myapi/user/internal/logic"
	"myapi/user/internal/svc"
	"myapi/user/internal/types"

	"github.com/jinzhu/copier"
)

type DetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DetailLogic) Detail(req *types.UserInfoReq) (resp *types.UserInfoResp, err error) {
	//获取jwt的用户id
	jwtUserId := logic.GetUidFromCtx(l.ctx)
	logx.Infof("??????", jwtUserId)

	// 获取用户信息
	userInfoResp, err := l.svcCtx.UserModel.FindOne(l.ctx, jwtUserId)
	if err != nil {
		return nil, errors.New("未查询到")
	}
	userInfo := types.User{}
	_ = copier.Copy(&userInfo, userInfoResp)

	return &types.UserInfoResp{UserInfo: userInfo}, nil

}
