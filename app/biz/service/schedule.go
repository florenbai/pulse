package service

import (
	"context"
	"darkhawk/app/biz/model"
	basepb "darkhawk/hertz_gen/base"
	pb "darkhawk/hertz_gen/schedule"
	"darkhawk/pkg/response"
	"encoding/json"
	"github.com/cloudwego/hertz/pkg/app"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/structpb"
	"strconv"
	"time"
)

type ScheduleService struct {
	ctx context.Context
	c   *app.RequestContext
}

func NewScheduleService(ctx context.Context, c *app.RequestContext) *ScheduleService {
	return &ScheduleService{
		ctx: ctx,
		c:   c,
	}
}

// CreateSchedule 创建值班
func (service *ScheduleService) CreateSchedule(req pb.PostScheduleRequest) error {
	scheduleModel := new(model.Schedule)
	ok, err := model.ExistName(service.ctx, scheduleModel.TableName(), "schedule_name", req.GetScheduleName(), nil)
	if err != nil {
		return err
	}
	if ok {
		return response.ScheduleExistErr
	}

	teamModel := new(model.Teams)
	userids, err := teamModel.GetUsersByTeam(service.ctx, []uint64{req.GetTeamId()})
	if err != nil {
		return err
	}
	t := model.LocalTime(time.Now())
	p, err := strconv.Atoi(req.GetSchedulePeriod()[0])
	if err != nil {
		return err
	}
	scheduleModel.ScheduleName = req.GetScheduleName()
	scheduleModel.ScheduleDesc = req.GetScheduleDesc()
	scheduleModel.TeamId = req.GetTeamId()
	scheduleModel.StartRange = req.GetScheduleTimeRange()[0]
	scheduleModel.EndRange = req.GetScheduleTimeRange()[1]
	scheduleModel.PeriodType = req.GetSchedulePeriod()[1]
	scheduleModel.Period = uint8(p)
	if len(req.GetUsers()) > 0 {
		scheduleModel.Users = req.GetUsers()
	} else {
		scheduleModel.Users = userids
	}

	scheduleModel.CreatedAt = &t

	err = model.Create(service.ctx, scheduleModel.TableName(), &scheduleModel)
	if err != nil {
		return err
	}
	if len(req.GetUsers()) > 0 {
		currentMouth := time.Now().Format("2006-01")
		planService := NewSchedulePlanService(service.ctx, service.c)
		return planService.InitUsersPlan(scheduleModel, currentMouth)
	}
	return nil
}

// GetScheduleList 值班列表
func (service *ScheduleService) GetScheduleList(req basepb.PageInfoReq) (*basepb.BaseListResp, error) {
	scheduleModel := new(model.Schedule)
	data, i, err := scheduleModel.List(service.ctx, req.GetPage(), req.GetPageSize(), model.NoneScope())

	listValue := new(structpb.ListValue)
	b, _ := json.Marshal(data)
	err = protojson.Unmarshal(b, listValue)
	if err != nil {
		return nil, err
	}
	return &basepb.BaseListResp{
		Total: uint64(i),
		List:  listValue,
	}, nil
}

// DeleteSchedule 删除值班表
func (service *ScheduleService) DeleteSchedule(id uint64) error {
	scheduleModel := new(model.Schedule)
	schedulePlanModel := new(model.SchedulePlan)
	err := model.DeleteOneById(service.ctx, scheduleModel.TableName(), id, model.Schedule{})
	if err != nil {
		return err
	}
	return model.DeleteByScope(service.ctx, schedulePlanModel.TableName(), model.SchedulePlan{}, model.WhereWithScope("schedule_id", id))
}
