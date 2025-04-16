package menu

import (
	"context"
	"darkhawk/app/biz/service"
	pb "darkhawk/hertz_gen/menu"
	"darkhawk/pkg/response"
	"github.com/cloudwego/hertz/pkg/app"
	"strconv"
)

func Create(ctx context.Context, c *app.RequestContext) {
	var req pb.MenuRequest
	if err := c.Bind(&req); err != nil {
		response.SendBaseResp(c, err)
		return
	}
	err := service.NewMenuService(ctx, c).CreateMenu(req)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendBaseResp(c, response.Success)
}

func Edit(ctx context.Context, c *app.RequestContext) {
	var req pb.MenuRequest
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
	err = service.NewMenuService(ctx, c).EditMenu(req, id)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendBaseResp(c, response.Success)
}

func List(ctx context.Context, c *app.RequestContext) {
	resp, err := service.NewMenuService(ctx, c).GetMenuList()
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendDataResp(c, response.Success, resp)
}

func Delete(ctx context.Context, c *app.RequestContext) {
	idStr := c.Params.ByName("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	err = service.NewMenuService(ctx, c).DeleteMenu(id)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendBaseResp(c, response.Success)
}

func ChangeSort(ctx context.Context, c *app.RequestContext) {
	var req []uint64
	if err := c.Bind(&req); err != nil {
		response.SendBaseResp(c, err)
		return
	}
	err := service.NewMenuService(ctx, c).ChangeMenuSort(req)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendBaseResp(c, response.Success)
}

func NameList(ctx context.Context, c *app.RequestContext) {
	resp, err := service.NewMenuService(ctx, c).GetMenuNameList()
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendDataResp(c, response.Success, resp)
}

func ChangeStatus(ctx context.Context, c *app.RequestContext) {
	idStr := c.Params.ByName("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	err = service.NewMenuService(ctx, c).ChangeMenuStatus(id)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendBaseResp(c, response.Success)
}
