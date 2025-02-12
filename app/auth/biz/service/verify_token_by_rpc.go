package service

import (
	"context"

	"github.com/cloudwego/biz-demo/gomall/app/auth/biz/pkg/jwt"
	auth "github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/auth"
)

type VerifyTokenByRPCService struct {
	ctx context.Context
} // NewVerifyTokenByRPCService new VerifyTokenByRPCService
func NewVerifyTokenByRPCService(ctx context.Context) *VerifyTokenByRPCService {
	return &VerifyTokenByRPCService{ctx: ctx}
}

// Run create note info
func (s *VerifyTokenByRPCService) Run(req *auth.VerifyTokenReq) (resp *auth.VerifyResp, err error) {
	// 验证JWT token
	_, err = jwt.ParseToken(req.Token)
	if err != nil {
		return &auth.VerifyResp{
			Res: false,
		}, nil
	}

	return &auth.VerifyResp{
		Res: true,
	}, nil
}
