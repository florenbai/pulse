package model

import (
	"context"
	"darkhawk/app/biz/dal/mysql"
	"gorm.io/gorm"
)

type Workspace struct {
	Model
	WorkspaceName  string `gorm:"type:varchar(100);column:workspace_name;comment:工作区名称" json:"workspaceName" form:"workspaceName"`
	WorkspaceDesc  string `gorm:"type:varchar(500);column:workspace_desc;comment:工作区描述" json:"workspaceDesc" form:"workspaceDesc"`
	SystemStrategy uint64 `gorm:"type:int(11);column:system_strategy;comment:全局告警策略" json:"systemStrategy" form:"systemStrategy"`
	Enable         bool   `gorm:"type:tinyint(1);default:true;comment:状态" json:"enable" form:"enable"`
}

type WorkspaceTeam struct {
	TeamId      uint64 `gorm:"type:int(11);column:team_id;comment:团队编号" json:"teamId" form:"teamId"`
	WorkspaceId uint64 `gorm:"type:int(11);column:workspace_id;comment:工作区编号" json:"workspaceId" form:"workspaceId"`
}

type WorkspaceList struct {
	Id            uint64 `gorm:"column:id;primary_key;comment:主键id" json:"id" form:"id"` // 主键
	WorkspaceName string `gorm:"type:varchar(100);column:workspace_name;comment:工作区名称" json:"workspaceName" form:"workspaceName"`
	WorkspaceDesc string `gorm:"type:varchar(500);column:workspace_desc;comment:工作区描述" json:"workspaceDesc" form:"workspaceDesc"`
	Enable        bool   `gorm:"type:tinyint(1);default:true;comment:状态" json:"enable" form:"enable"`
	Pending       int    `json:"pending" form:"pending"`
	Processing    int    `json:"processing" form:"processing"`
	SourceCount   int    `json:"sourceCount" gorm:"source_count"`
}

type WorkspaceBaseInfo struct {
	WorkspaceName  string `gorm:"type:varchar(100);column:workspace_name;comment:工作区名称" json:"workspaceName" form:"workspaceName"`
	WorkspaceDesc  string `gorm:"type:varchar(500);column:workspace_desc;comment:工作区描述" json:"workspaceDesc" form:"workspaceDesc"`
	Teams          string `json:"teams"`
	SystemStrategy uint64 `gorm:"type:int(11);column:system_strategy;comment:全局告警策略" json:"systemStrategy" form:"systemStrategy"`
}

func (m *Workspace) TableName() string {
	return mysql.WorkspaceTable
}

// Create 创建工作区
func (m *Workspace) Create(ctx context.Context, team []uint64) error {
	return mysql.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Table(m.TableName()).Create(&m).Error; err != nil {
			return err
		}
		var workspaceTeams []WorkspaceTeam
		for _, v := range team {
			workspaceTeam := WorkspaceTeam{
				TeamId:      v,
				WorkspaceId: m.Id,
			}
			workspaceTeams = append(workspaceTeams, workspaceTeam)
		}
		if err := tx.Table(mysql.WorkspaceTeamTable).Create(&workspaceTeams).Error; err != nil {
			return err
		}

		return nil
	})
}

// Update 更新工作区基础信息
func (m *Workspace) Update(ctx context.Context, team []uint64) error {
	return mysql.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Table(m.TableName()).Save(&m).Error; err != nil {
			return err
		}
		var workspaceTeams []WorkspaceTeam
		for _, v := range team {
			workspaceTeam := WorkspaceTeam{
				TeamId:      v,
				WorkspaceId: m.Id,
			}
			workspaceTeams = append(workspaceTeams, workspaceTeam)
		}
		err := tx.Table(mysql.WorkspaceTeamTable).Where("workspace_id", m.Id).Delete(WorkspaceTeam{}).Error
		if err != nil {
			return err
		}
		if err := tx.Table(mysql.WorkspaceTeamTable).Create(&workspaceTeams).Error; err != nil {
			return err
		}

		return nil
	})
}

// List 工作区列表
func (m *Workspace) List(ctx context.Context, page uint64, pageSize uint64, scopes ...func(*gorm.DB) *gorm.DB) ([]WorkspaceList, int64, error) {
	var i int64
	var data []WorkspaceList

	err := mysql.DB.WithContext(ctx).Table(m.TableName()).Scopes(scopes...).Count(&i).Error
	if err != nil {
		return nil, 0, err
	}
	scopes = append(scopes, OrderScope("workspace.id DESC"), Paginate(page, pageSize))
	err = mysql.DB.WithContext(ctx).Table(m.TableName()).
		Select("workspace.*,t1.pending,t2.processing").
		Joins("LEFT JOIN (SELECT COUNT(*) AS pending,alert_workspace.workspace_id FROM `alert_event` LEFT JOIN alert_workspace ON alert_workspace.alert_id = alert_event.id WHERE progress = 1 GROUP BY alert_workspace.workspace_id) AS t1 ON t1.workspace_id = workspace.id").
		Joins("LEFT JOIN (SELECT COUNT(*) AS processing,alert_workspace.workspace_id FROM `alert_event` LEFT JOIN alert_workspace ON alert_workspace.alert_id = alert_event.id  WHERE progress = 2 GROUP BY alert_workspace.workspace_id) AS t2 ON t2.workspace_id = workspace.id").
		Scopes(scopes...).Find(&data).Error
	if err != nil {
		return nil, 0, err
	}

	return data, i, nil
}

// BaseInfo 获取基础信息
func (m *Workspace) BaseInfo(ctx context.Context, id uint64) (*WorkspaceBaseInfo, error) {
	var bl WorkspaceBaseInfo
	err := mysql.DB.WithContext(ctx).Table(m.TableName()).Select("workspace.*,GROUP_CONCAT(workspace_team.team_id) AS teams").Joins("LEFT JOIN workspace_team ON workspace_team.workspace_id = workspace.id").Where("workspace.id", id).Group("workspace.id").Take(&bl).Error
	if err != nil {
		return nil, err
	}
	return &bl, nil
}

// Delete 删除工作区
func (m *Workspace) Delete(ctx context.Context) error {
	return mysql.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Table(mysql.WorkspaceSourceTable).Where("workspace_id", m.Id).Delete(WorkspaceSource{}).Error; err != nil {
			return err
		}
		if err := tx.Table(mysql.WorkspaceTeamTable).Where("workspace_id", m.Id).Delete(WorkspaceTeam{}).Error; err != nil {
			return err
		}
		if err := tx.Table(mysql.AlertAggregationTable).Where("workspace_id", m.Id).Delete(AlertAggregation{}).Error; err != nil {
			return err
		}
		if err := tx.Table(mysql.AlertStrategyTable).Where("workspace_id", m.Id).Delete(AlertStrategy{}).Error; err != nil {
			return err
		}
		if err := tx.Table(mysql.LevelSourceMappingTable).Where("workspace_id", m.Id).Delete(LevelSourceMapping{}).Error; err != nil {
			return err
		}
		if err := tx.Table(m.TableName()).Delete(&m).Error; err != nil {
			return err
		}

		return nil
	})
}

func (m *Workspace) IdMap(ctx context.Context) (map[uint64]string, error) {
	var data = make(map[uint64]string)
	var all []Workspace
	err := GetAll(ctx, m.TableName(), &all)
	if err != nil {
		return nil, err
	}
	for _, v := range all {
		data[v.Id] = v.WorkspaceName
	}
	return data, nil
}

func (m *Workspace) GetEnableWorkspace(ctx context.Context, workspaces []uint64) ([]uint64, error) {
	var data []uint64
	err := mysql.DB.WithContext(ctx).Table(m.TableName()).Where("enable", 1).Where("id IN ?", workspaces).Pluck("id", &data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}
