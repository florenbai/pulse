package user

import (
	"context"
	"darkhawk/app/biz/model"
	"darkhawk/app/biz/model/request"
	"darkhawk/app/biz/service"
	pb "darkhawk/hertz_gen/user"
	"darkhawk/pkg/response"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/sessions"
	"strconv"
)

func CreateUser(ctx context.Context, c *app.RequestContext) {
	var req pb.UserRequest
	if err := c.Bind(&req); err != nil {
		response.SendBaseResp(c, err)
		return
	}
	err := service.NewUserService(ctx, c).CreateUser(req)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}

	response.SendBaseResp(c, response.Success)
}

func UpdateUser(ctx context.Context, c *app.RequestContext) {
	var req pb.UserRequest
	if err := c.Bind(&req); err != nil {
		response.SendBaseResp(c, err)
		return
	}
	idStr := c.Params.ByName("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	err = service.NewUserService(ctx, c).UpdateUser(id, req)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}

	response.SendBaseResp(c, response.Success)
}

func UserDetail(ctx context.Context, c *app.RequestContext) {
	idStr := c.Params.ByName("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	userModel := new(model.User)
	err = model.GetOneById(ctx, userModel.TableName(), id, &userModel)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}

	response.SendDataResp(c, response.Success, userModel)
}

func List(ctx context.Context, c *app.RequestContext) {
	var err error
	var req pb.UserQuery
	if err = c.Bind(&req); err != nil {
		response.SendBaseResp(c, err)
		return
	}

	resp, err := service.NewUserService(ctx, c).GetUserList(req)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}

	response.SendDataResp(c, response.Success, resp)
}

func Menu(ctx context.Context, c *app.RequestContext) {
	resp, err := service.NewMenuService(ctx, c).GetUserMenu()
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}

	response.SendDataResp(c, response.Success, resp)
}

func All(ctx context.Context, c *app.RequestContext) {
	userModel := new(model.User)
	data, err := userModel.All(ctx)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendDataResp(c, response.Success, data)
}

func Delete(ctx context.Context, c *app.RequestContext) {
	idStr := c.Params.ByName("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	err = service.NewUserService(ctx, c).DeleteUser(id)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendBaseResp(c, nil)
}

func Logout(ctx context.Context, c *app.RequestContext) {
	session := sessions.Default(c)
	session.Clear()
	_ = session.Save()
	response.SendBaseResp(c, nil)
}

func Login(ctx context.Context, c *app.RequestContext) {
	var req request.LoginRequest
	if err := c.BindJSON(&req); err != nil {
		response.SendBaseResp(c, err)
		return
	}
	token, user, err := service.NewUserService(ctx, c).VerifyUserInfo(req.Username, req.Password)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}

	data := map[string]interface{}{
		"token": token,
		"user": map[string]string{
			"username": user.Username,
			"nickname": user.Nickname,
		},
	}
	response.SendDataResp(c, response.Success, data)
}
