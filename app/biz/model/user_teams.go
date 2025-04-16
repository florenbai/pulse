package model

import (
	"context"
	"darkhawk/app/biz/dal/mysql"
)

type TeamUser struct {
	TeamId uint64 `gorm:"type:int(11);column:team_id;comment:团队编号" json:"teamId" form:"teamId"`
	UserId uint64 `gorm:"type:int(11);column:user_id;comment:用户编号" json:"userId" form:"userId"`
}

func (m *TeamUser) TableName() string {
	return mysql.TeamUserTable
}

func (m *TeamUser) IsAllWorkspace(ctx context.Context, uid uint64) bool {
	var i int64
	mysql.DB.WithContext(ctx).Table(m.TableName()).Joins("LEFT JOIN teams ON teams.id = user_teams.team_id").Where("user_teams.user_id", uid).Where("data_permission", 2).Count(&i)
	if i > 0 {
		return true
	}
	return false
}

func (m *TeamUser) GetUserWorkspace(ctx context.Context, uid uint64) ([]interface{}, error) {
	var workspaces []interface{}
	err := mysql.DB.WithContext(ctx).Table(m.TableName()).Joins("LEFT JOIN workspace_team ON workspace_team.team_id = user_teams.team_id").Where("user_teams.user_id", uid).Pluck("workspace_team.workspace_id", &workspaces).Error
	if err != nil {
		return nil, err
	}
	return workspaces, nil
}
