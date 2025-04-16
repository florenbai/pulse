package schedule

import (
	"context"
	"darkhawk/app/biz/model"
	"darkhawk/app/biz/service"
	basepb "darkhawk/hertz_gen/base"
	pb "darkhawk/hertz_gen/schedule"
	"darkhawk/pkg/response"
	"github.com/cloudwego/hertz/pkg/app"
	"strconv"
)

func CreatedSchedule(ctx context.Context, c *app.RequestContext) {
	var req pb.PostScheduleRequest
	if err := c.Bind(&req); err != nil {
		response.SendBaseResp(c, err)
		return
	}
	err := service.NewScheduleService(ctx, c).CreateSchedule(req)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendBaseResp(c, nil)
	return
}

func List(ctx context.Context, c *app.RequestContext) {
	var req basepb.PageInfoReq
	if err := c.Bind(&req); err != nil {
		response.SendBaseResp(c, err)
		return
	}
	resp, err := service.NewScheduleService(ctx, c).GetScheduleList(req)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}

	response.SendDataResp(c, response.Success, resp)
}

func All(ctx context.Context, c *app.RequestContext) {
	scheduleModel := new(model.Schedule)
	data, err := scheduleModel.All(ctx)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendDataResp(c, response.Success, data)
}

func DeleteSchedule(ctx context.Context, c *app.RequestContext) {
	idStr := c.Params.ByName("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	err = service.NewScheduleService(ctx, c).DeleteSchedule(id)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendBaseResp(c, response.Success)
}
