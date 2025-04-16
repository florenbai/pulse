package model

import (
	"context"
	"darkhawk/app/biz/dal/mysql"
	"gorm.io/gorm"
)

type Teams struct {
	Model
	TeamName       string `gorm:"type:varchar(100);column:team_name;comment:团队名称" json:"teamName" form:"teamName"`
	TeamDesc       string `gorm:"type:text;column:team_desc;comment:团队描述" json:"teamDesc" form:"teamDesc"`
	UpdatedBy      uint64 `gorm:"type:int(11);column:updated_by;comment:团队描述" json:"updatedBy" form:"updatedBy"`
	Status         bool   `gorm:"type:tinyint(1);default:true;comment:状态" json:"status" form:"status"`
	DataPermission uint64 `json:"dataPermission"`
}

type TeamsList struct {
	Teams
	NickName string `gorm:"column:nickname;comment:更新用户" json:"nickname" form:"nickname"`
}

type TeamDetail struct {
	Teams
	TeamMembers []uint64 `json:"teamMembers" form:"teamMembers"`
}

type TeamUserGroup struct {
	TeamName string `gorm:"type:varchar(100);column:team_name;comment:团队名称" json:"teamName" form:"teamName"`
	UserId   uint64 `gorm:"type:int(11);column:user_id;comment:用户编号" json:"userId" form:"userId"`
	NickName string `gorm:"column:nickname;comment:更新用户" json:"nickname" form:"nickname"`
}

func (m *Teams) TableName() string {
	return mysql.TeamsTable
}

// Create 创建团队
func (m *Teams) Create(ctx context.Context, teamMembers []uint64) error {

	return mysql.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Table(m.TableName()).Create(&m).Error; err != nil {
			return err
		}
		var teamUsers []TeamUser
		for _, v := range teamMembers {
			teamUser := TeamUser{
				TeamId: m.Id,
				UserId: v,
			}
			teamUsers = append(teamUsers, teamUser)
		}
		if err := tx.Table(mysql.TeamUserTable).Create(&teamUsers).Error; err != nil {
			return err
		}
		return nil
	})
}

// Update 编辑团队
func (m *Teams) Update(ctx context.Context, teamMembers []uint64) error {

	return mysql.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Table(m.TableName()).Where("id", m.Id).Updates(&m).Error; err != nil {
			return err
		}
		var teamUsers []TeamUser
		for _, v := range teamMembers {
			teamUser := TeamUser{
				TeamId: m.Id,
				UserId: v,
			}
			teamUsers = append(teamUsers, teamUser)
		}
		if err := tx.Table(mysql.TeamUserTable).Where("team_id", m.Id).Delete(Teams{}).Error; err != nil {
			return err
		}
		if err := tx.Table(mysql.TeamUserTable).Create(&teamUsers).Error; err != nil {
			return err
		}
		return nil
	})
}

// List 团队列表
func (m *Teams) List(ctx context.Context, page uint64, pageSize uint64, scopes ...func(*gorm.DB) *gorm.DB) ([]TeamsList, int64, error) {
	var err error
	var data []TeamsList

	var i int64
	err = mysql.DB.WithContext(ctx).Table(m.TableName()).Scopes(scopes...).Count(&i).Error
	if err != nil {
		return nil, 0, err
	}
	scopes = append(scopes, OrderScope("id DESC"), Paginate(page, pageSize))
	err = mysql.DB.WithContext(ctx).Select("teams.*,user.nickname").Table(m.TableName()).Joins("LEFT JOIN user ON user.id = teams.updated_by").Scopes(scopes...).Find(&data).Error
	if err != nil {
		return nil, 0, err
	}
	return data, i, nil
}

func (m *Teams) GetTeamDetail(ctx context.Context, id uint64) (*TeamDetail, error) {
	var users []uint64
	var data TeamDetail
	err := mysql.DB.WithContext(ctx).Table(m.TableName()).Where("id", id).First(&m).Error
	if err != nil {
		return nil, err
	}
	err = mysql.DB.WithContext(ctx).Table(mysql.TeamUserTable).Where("team_id", id).Pluck("user_id", &users).Error
	if err != nil {
		return nil, err
	}
	data.TeamMembers = users
	data.Teams = *m
	return &data, nil
}

// Delete 删除团队
func (m *Teams) Delete(ctx context.Context, id uint64) error {
	err := mysql.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Table(m.TableName()).Where("id", id).Delete(Teams{}).Error; err != nil {
			return err
		}
		if err := tx.Table(mysql.TeamUserTable).Where("team_id", id).Delete(TeamUser{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

func (m *Teams) ChangeStatus(ctx context.Context, id uint64) error {
	var status = false
	if m.Status == false {
		status = true
	}
	return mysql.DB.WithContext(ctx).Table(m.TableName()).Where("id", id).Update("status", status).Error
}

func (m *Teams) All(ctx context.Context) ([]Teams, error) {
	var data []Teams
	err := GetAll(ctx, m.TableName(), &data, WhereWithScope("status", true))
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (m *Teams) AllUser(ctx context.Context, id []string) ([]TeamUserGroup, error) {
	var data []TeamUserGroup
	err := mysql.DB.WithContext(ctx).Table(mysql.TeamUserTable).Select("user.nickname,user_teams.user_id,teams.team_name").
		Joins("LEFT JOIN user ON user.id = user_teams.user_id").
		Joins("LEFT JOIN teams ON teams.id = user_teams.team_id").
		Where("user_teams.team_id IN ?", id).Find(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (m *Teams) GetUsersByTeam(ctx context.Context, id []uint64) ([]uint64, error) {
	var data []uint64
	err := mysql.DB.WithContext(ctx).Table(mysql.TeamUserTable).Where("user_teams.team_id IN ?", id).Pluck("user_id", &data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}
