package service

import (
	"context"
	"time"

	"github.com/cloudwego/biz-demo/gomall/app/auth/biz/pkg/jwt"
	auth "github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/auth"
)

type DeliverTokenByRPCService struct {
	ctx context.Context
} // NewDeliverTokenByRPCService new DeliverTokenByRPCService
func NewDeliverTokenByRPCService(ctx context.Context) *DeliverTokenByRPCService {
	return &DeliverTokenByRPCService{ctx: ctx}
}

// Run create note info
func (s *DeliverTokenByRPCService) Run(req *auth.DeliverTokenReq) (resp *auth.DeliveryResp, err error) {
	// 设置token过期时间为4小时
	expireTime := time.Now().Add(4 * time.Hour).Unix()

	// 生成JWT token
	token, err := jwt.GenerateToken(req.UserId, "", "user", expireTime)
	if err != nil {
		return nil, err
	}

	return &auth.DeliveryResp{
		Token: token,
	}, nil
}
