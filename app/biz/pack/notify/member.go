package notify

import (
	"context"
	"darkhawk/app/biz/model"
	"darkhawk/pkg/utils"
	"time"
)

type Notify struct {
	Teams     []uint64
	Members   []uint64
	Schedules []uint64
}

func (n *Notify) GetNotifyMembers(ctx context.Context) ([]uint64, error) {
	var members []uint64
	// 排班人员获取
	scheduleMembers, err := n.GetSchedulesMembers(ctx)
	if err != nil {
		return nil, err
	}
	members = append(members, scheduleMembers...)
	// 团队
	teamModel := new(model.Teams)
	teamUser, err := teamModel.GetUsersByTeam(ctx, n.Teams)
	if err != nil {
		return nil, err
	}
	members = append(members, teamUser...)
	// 个人
	members = append(members, n.Members...)

	return utils.RemoveDuplicate(members), nil

}

// GetSchedulesMembers 获取排班用户
func (n *Notify) GetSchedulesMembers(ctx context.Context) ([]uint64, error) {
	var members []uint64
	if len(n.Schedules) == 0 {
		return nil, nil
	}
	schedulePlanModel := new(model.SchedulePlan)
	plans, err := schedulePlanModel.GetCurrentMonthSchedulePlanByIds(ctx, n.Schedules)
	if err != nil {
		return nil, err
	}
	today := time.Now().Day()
	for _, planItem := range plans {
		for _, v := range planItem.Plan {
			if v.Day == today && v.IsSchedule && !utils.InOfT(planItem.Uid, members) {
				members = append(members, planItem.Uid)
			}
		}
	}
	return members, nil
}
