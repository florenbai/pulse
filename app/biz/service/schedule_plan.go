package service

import (
	"context"
	"darkhawk/app/biz/model"
	pb "darkhawk/hertz_gen/schedule"
	"darkhawk/pkg/utils"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"time"
)

type SchedulePlanService struct {
	ctx context.Context
	c   *app.RequestContext
}

func NewSchedulePlanService(ctx context.Context, c *app.RequestContext) *SchedulePlanService {
	return &SchedulePlanService{
		ctx: ctx,
		c:   c,
	}
}

// GetSchedulePlanList 获取用度值班列表
func (service *SchedulePlanService) GetSchedulePlanList(id uint64, req pb.SchedulePlanListRequest) ([]model.SchedulePlanList, error) {
	spModel := new(model.SchedulePlan)
	scheduleModel := new(model.Schedule)
	err := model.GetOneById(service.ctx, scheduleModel.TableName(), id, &scheduleModel)
	if err != nil {
		return nil, err
	}
	// 是否存在当前月份计划，不存在则初始化月份数据
	givenMonth, err := time.Parse("2006-01", req.GetMonth())
	if err != nil {
		return nil, err
	}
	currentTime := time.Now()
	count, err := model.GetCountWithScope(service.ctx, spModel.TableName(), model.WhereWithScope("schedule_id", id), model.WhereWithScope("month", req.GetMonth()))
	if count == 0 && len(scheduleModel.Users) > 0 {
		if currentTime.Year() > givenMonth.Year() || (currentTime.Year() == givenMonth.Year() && currentTime.Month() < givenMonth.Month()) {
			err = service.InitTeamPlan(id, req.GetMonth())
			if err != nil {
				return nil, err
			}
		} /*
			else {
				err = service.NextTeamPlan(*scheduleModel, req.GetMonth())
				if err != nil {
					return nil, err
				}
			}
		*/
	}
	data, err := spModel.GetSchedulePlanList(service.ctx, model.WhereWithScope("schedule_id", id), model.WhereWithScope("month", req.GetMonth()))
	if err != nil {
		return nil, err
	}
	return data, nil
}

// SetPlanItem 设置值班项
func (service *SchedulePlanService) SetPlanItem(id uint64, req pb.SchedulePlanSetRequest) error {
	spModel := new(model.SchedulePlan)
	err := model.GetOneById(service.ctx, spModel.TableName(), id, &spModel)
	if err != nil {
		return err
	}
	if len(spModel.Plan) == 0 {
		plan, err := service.InitMonthPlan(spModel.Month)
		if err != nil {
			return err
		}
		spModel.Plan = plan
	}
	for _, v := range req.Index {
		spModel.Plan[v].IsSchedule = true
	}

	return model.EditOneById(service.ctx, spModel.TableName(), id, &spModel)
}

// DeletePlanItem 删除值班项
func (service *SchedulePlanService) DeletePlanItem(id uint64, req pb.SchedulePlanSetRequest) error {
	spModel := new(model.SchedulePlan)
	err := model.GetOneById(service.ctx, spModel.TableName(), id, &spModel)
	if err != nil {
		return err
	}
	if len(spModel.Plan) == 0 {
		return nil
	}
	for _, v := range req.Index {
		if (len(spModel.Plan) - 1) >= int(v) {
			spModel.Plan[v].IsSchedule = false
		} else {
			return nil
		}
	}
	return model.EditOneById(service.ctx, spModel.TableName(), id, &spModel)
}

// InitMonthPlan 初始化月度值班项
func (service *SchedulePlanService) InitMonthPlan(month string) ([]model.PlanItem, error) {
	var plan []model.PlanItem
	days, err := utils.GetMonthDays(month)
	if err != nil {
		return nil, err
	}
	var i = 1
	for {
		if i > days {
			break
		}
		plan = append(plan, model.PlanItem{
			Day:        i,
			IsSchedule: false,
		})
		i++
	}
	return plan, nil
}

// InitTeamPlan 初始化团队值班
func (service *SchedulePlanService) InitTeamPlan(id uint64, month string) error {
	scheduleModel := new(model.Schedule)
	spModel := new(model.SchedulePlan)
	users, err := scheduleModel.GetScheduleUserById(service.ctx, id)
	if err != nil {
		return err
	}
	var plans []model.SchedulePlan
	for _, v := range users {
		p, err := service.InitMonthPlan(month)
		if err != nil {
			return err
		}
		plans = append(plans, model.SchedulePlan{
			ScheduleId: id,
			Month:      month,
			Plan:       p,
			Uid:        v,
		})
	}
	return model.Create(service.ctx, spModel.TableName(), &plans)
}

func (service *SchedulePlanService) NextTeamPlan(schedule model.Schedule, month string) error {
	spModel := new(model.SchedulePlan)
	c, err := model.GetCountWithScope(service.ctx, spModel.TableName(), model.WhereWithScope("schedule_id", schedule.Id), model.WhereWithScope("month", month))
	if err != nil {
		return err
	}
	if c > 0 {
		return err
	}
	plans, err := service.SchedulePlanPeriodTypeDay(schedule.Id, schedule.Users, month, schedule.CreatedAt.Format(), int(schedule.Period))
	if err != nil {
		return err
	}
	return model.Create(service.ctx, spModel.TableName(), &plans)
}

func (service *SchedulePlanService) InitUsersPlan(schedule *model.Schedule, month string) error {
	var schedulePlans []model.SchedulePlan
	spModel := new(model.SchedulePlan)

	schedulePlans, err := service.SchedulePlanPeriodTypeDay(schedule.Id, schedule.Users, month, schedule.CreatedAt.Format(), int(schedule.Period))
	if err != nil {
		return err
	}
	return model.Create(service.ctx, spModel.TableName(), &schedulePlans)
}

func (service *SchedulePlanService) SchedulePlanPeriodTypeDay(id uint64, users []uint64, month string, initialDay string, period int) ([]model.SchedulePlan, error) {
	var newUserSlice []uint64
	var scheduleDays int
	var startType string
	var schedulePlans []model.SchedulePlan
	initialTime, err := time.Parse("2006-01-02 15:04:05", initialDay)
	initialMonth := initialTime.Format("2006-01")
	if err != nil {
		return nil, err
	}
	// 排班当前月份为初始月份
	if month == initialMonth {
		days, err := utils.GetMonthDays(month)
		if err != nil {
			return nil, err
		}
		scheduleDays = days - initialTime.Day()
		startType = "current"
		// 排班为之后月份
	} else {
		targetDay, err := time.Parse("2006-01-02", fmt.Sprintf("%s-01", month))
		if err != nil {
			return nil, err
		}
		diff := targetDay.Sub(initialTime)
		scheduleDays = int(diff.Hours()/24) - 1
	}
	fmt.Println(scheduleDays)
	userCount := len(users)
	if scheduleDays == 0 {
		scheduleDays = 1
	}
	if userCount < scheduleDays {
		position := scheduleDays % userCount
		newUserSlice = append(users[position:], users[:position]...)
	} else {
		newUserSlice = users[:scheduleDays]
	}
	fmt.Println(newUserSlice)
	for k, v := range newUserSlice {
		plans, err := service.InitMonthPlan(month)
		fmt.Println(plans)
		if err != nil {
			return nil, err
		}
		userPlan := SchedulePlanByDay(k, plans, period, len(newUserSlice), startType)
		schedulePlans = append(schedulePlans, model.SchedulePlan{
			ScheduleId: id,
			Month:      month,
			Plan:       userPlan,
			Uid:        v,
		})
	}
	return schedulePlans, nil
}

func SchedulePlanByDay(userPosition int, plans []model.PlanItem, period int, userCount int, startType string) []model.PlanItem {
	var start int
	u := userCount * period
	if startType == "current" {
		// 当前用户初始日期
		if userPosition == 0 {
			start = time.Now().Day() + userPosition
		} else {
			start = time.Now().Day() + userPosition*period
		}
	} else {
		if userPosition == 0 {
			start = 1
		} else {
			start = 1 + userPosition*period
		}
	}
	i := 0
	for k, v := range plans {
		if v.Day < start {
			continue
		}
		if period-i > 0 {
			plans[k].IsSchedule = true
		}
		i++
		if u-i == 0 {
			i = 0
		}
	}
	return plans
}
