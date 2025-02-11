package service

import (
	"context"

	"github.com/cloudwego/biz-demo/gomall/app/user/biz/dal/mysql"
	"github.com/cloudwego/biz-demo/gomall/app/user/biz/model"
	user "github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/user"
)

type DeleteService struct {
	ctx context.Context
} // NewDeleteService new DeleteService
func NewDeleteService(ctx context.Context) *DeleteService {
	return &DeleteService{ctx: ctx}
}

// Run create note info
func (s *DeleteService) Run(req *user.DeleteReq) (resp *user.DeleteResp, err error) {
	// Finish your business logic.
	//通过传入user_id删除用户，要先根据req.UserId查询用户，然后删除用户
	if err = model.Delete(mysql.DB, s.ctx, &model.User{Base: model.Base{ID: int(req.UserId)}}); err != nil {
		return
	}
	return &user.DeleteResp{Success: true}, nil
}
