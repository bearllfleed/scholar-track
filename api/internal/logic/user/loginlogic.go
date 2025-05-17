package user

import (
	"context"
	"github.com/bearllflee/scholar-track/api/internal/svc"
	"github.com/bearllflee/scholar-track/api/internal/types"
	"github.com/bearllflee/scholar-track/api/internal/utils"
	"github.com/bearllflee/scholar-track/pkg/cerror"
	"github.com/bearllflee/scholar-track/rpc/system/client/user"

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

func (l *LoginLogic) Login(req *types.LoginReq) (*types.LoginResp, error) {
	resp, err := l.svcCtx.User.FindUserByUsername(l.ctx, &user.FindUserByUsernameReq{
		Username: req.Username,
	})
	if err != nil {
		return nil, err
	}
	// 验证密码
	if !utils.BcryptCheck(req.Password, resp.Password) {
		return nil, cerror.ErrPasswordWrong
	}
	// 生成token
	jwt := utils.NewJwt([]byte(l.svcCtx.Config.JwtConf.SecretKey), l.svcCtx.Config.JwtConf.Expire, l.svcCtx.Config.JwtConf.Buffer, l.svcCtx.Config.JwtConf.Issuer, l.svcCtx.Config.JwtConf.Audience)
	claims := jwt.CreateClaims(types.BaseClaims{
		ID:          resp.Id,
		AuthorityId: resp.RoleId,
		Username:    resp.Username,
	})
	token, err := jwt.GenerateToken(&claims)
	if err != nil {
		return nil, err
	}
	return &types.LoginResp{
		Token: token,
	}, nil
}
