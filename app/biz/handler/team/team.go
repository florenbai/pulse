package team

import (
	"context"
	"darkhawk/app/biz/model"
	"darkhawk/app/biz/service"
	pb "darkhawk/hertz_gen/team"
	"darkhawk/pkg/response"
	"github.com/cloudwego/hertz/pkg/app"
	"strconv"
	"strings"
)

func CreateTeam(ctx context.Context, c *app.RequestContext) {
	var req pb.TeamRequest
	if err := c.BindAndValidate(&req); err != nil {
		response.SendBaseResp(c, err)
		return
	}
	err := service.NewTeamService(ctx, c).CreateTeam(req)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}

	response.SendBaseResp(c, response.Success)
}

func UpdateTeam(ctx context.Context, c *app.RequestContext) {
	var req pb.TeamRequest
	if err := c.BindAndValidate(&req); err != nil {
		response.SendBaseResp(c, err)
		return
	}
	idStr := c.Params.ByName("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	err = service.NewTeamService(ctx, c).UpdateTeam(id, req)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}

	response.SendBaseResp(c, response.Success)
}

func TeamDetail(ctx context.Context, c *app.RequestContext) {
	idStr := c.Params.ByName("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	teamsModel := new(model.Teams)
	data, err := teamsModel.GetTeamDetail(ctx, id)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendDataResp(c, response.Success, data)
}

func List(ctx context.Context, c *app.RequestContext) {
	var err error
	var req pb.TeamQuery
	if err = c.Bind(&req); err != nil {
		response.SendBaseResp(c, err)
		return
	}

	resp, err := service.NewTeamService(ctx, c).GetTeamList(req)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}

	response.SendDataResp(c, response.Success, resp)
}

func TeamDelete(ctx context.Context, c *app.RequestContext) {
	idStr := c.Params.ByName("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	err = service.NewTeamService(ctx, c).DeleteTeam(id)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendBaseResp(c, response.Success)
}

func ChangeTeamStatus(ctx context.Context, c *app.RequestContext) {
	idStr := c.Params.ByName("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	err = service.NewTeamService(ctx, c).ChangeTeamStatus(id)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}

	response.SendBaseResp(c, response.Success)
}

func All(ctx context.Context, c *app.RequestContext) {
	teamsModel := new(model.Teams)
	data, err := teamsModel.All(ctx)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendDataResp(c, response.Success, data)
}

func TeamUserAll(ctx context.Context, c *app.RequestContext) {
	var err error
	teamStr := c.Params.ByName("team")
	if teamStr == "" {
		response.SendBaseResp(c, response.TeamNotFoundErr)
		return
	}
	teamArr := strings.Split(teamStr, ",")
	teamsModel := new(model.Teams)
	data, err := teamsModel.AllUser(ctx, teamArr)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendDataResp(c, response.Success, data)
}

func Authorization(ctx context.Context, c *app.RequestContext) {
	idStr := c.Params.ByName("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	var req pb.AuthorizationRequest
	if err = c.Bind(&req); err != nil {
		response.SendBaseResp(c, err)
		return
	}

	err = service.NewTeamService(ctx, c).AuthorizationTeam(id, req)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendBaseResp(c, response.Success)
}

func AuthorizationInfo(ctx context.Context, c *app.RequestContext) {
	idStr := c.Params.ByName("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	resp, err := service.NewTeamService(ctx, c).AuthorizationInfo(id)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendDataResp(c, response.Success, resp)
}
