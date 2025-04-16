package schedule

import (
	"context"
	"darkhawk/app/biz/service"
	pb "darkhawk/hertz_gen/schedule"
	"darkhawk/pkg/response"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"strconv"
	"time"
)

func PlanList(ctx context.Context, c *app.RequestContext) {
	idStr := c.Params.ByName("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	var req pb.SchedulePlanListRequest
	if err := c.Bind(&req); err != nil {
		response.SendBaseResp(c, err)
		return
	}
	givenMonth, err := time.Parse("2006-01", req.GetMonth())
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	nextMonth := time.Now().AddDate(0, 1, 0)
	if givenMonth.After(nextMonth) {
		response.SendBaseResp(c, errors.New("您只能查看不大于当前时间1个月的排班"))
		return
	}
	resp, err := service.NewSchedulePlanService(ctx, c).GetSchedulePlanList(id, req)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}

	response.SendDataResp(c, response.Success, resp)
}

func PlanIn(ctx context.Context, c *app.RequestContext) {
	idStr := c.Params.ByName("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	var req pb.SchedulePlanSetRequest
	if err := c.Bind(&req); err != nil {
		response.SendBaseResp(c, err)
		return
	}
	err = service.NewSchedulePlanService(ctx, c).SetPlanItem(id, req)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendBaseResp(c, response.Success)
}

func PlanOut(ctx context.Context, c *app.RequestContext) {
	idStr := c.Params.ByName("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	var req pb.SchedulePlanSetRequest
	if err := c.Bind(&req); err != nil {
		response.SendBaseResp(c, err)
		return
	}
	err = service.NewSchedulePlanService(ctx, c).DeletePlanItem(id, req)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendBaseResp(c, response.Success)
}
