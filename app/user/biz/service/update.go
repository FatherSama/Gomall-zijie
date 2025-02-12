package service

import (
	"context"

	"github.com/cloudwego/biz-demo/gomall/app/user/biz/dal/mysql"
	"github.com/cloudwego/biz-demo/gomall/app/user/biz/model"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/user"
	"golang.org/x/crypto/bcrypt"
)

type UpdateService struct {
	ctx context.Context
} // NewUpdateService new UpdateService
func NewUpdateService(ctx context.Context) *UpdateService {
	return &UpdateService{ctx: ctx}
}

// Run create note info
func (s *UpdateService) Run(req *user.UpdateReq) (resp *user.UpdateResp, err error) {
	// Finish your business logic.
	//找到用户，确认原密码，修改新密码
	userRow, err := model.GetByEmail(mysql.DB, s.ctx, req.Email)
	if err != nil {
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(userRow.PasswordHashed), []byte(req.Password))
	if err != nil {
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return
	}
	userRow.PasswordHashed = string(hashedPassword)
	err = model.Update(mysql.DB, s.ctx, userRow)
	if err != nil {
		return
	}

	return &user.UpdateResp{
		Success: true,
	}, nil
}
