package model

import (
	"context"
	"darkhawk/app/biz/dal/mysql"
	"gorm.io/gorm"
)

type User struct {
	Model
	Nickname string `gorm:"type:varchar(128);column:nickname;comment:用户显示名称" json:"nickname"`
	Username string `gorm:"type:varchar(128);column:username;comment:用户名称" json:"username"`
	Password string `gorm:"type:varchar(128);column:password;comment:用户密码" json:"password"`
	Empno    string `gorm:"type:varchar(128);column:empno;comment:工号" json:"empno"`
	UserId   string `gorm:"type:varchar(128);column:userid;comment:微信工号" json:"userid"`
	Email    string `gorm:"type:varchar(128);column:email;comment:邮箱" json:"email"`
	Phone    string `gorm:"type:varchar(128);column:phone;comment:电话" json:"phone"`
	UserType string `gorm:"type:varchar(128);column:user_type;comment:电话" json:"userType"`
}

func (m *User) TableName() string {
	return mysql.UserTable
}

func (m *User) All(ctx context.Context) ([]User, error) {
	var data []User
	err := GetAll(ctx, m.TableName(), &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (m *User) GetUsersByUserIds(ctx context.Context, ids []uint64) ([]User, error) {
	var data []User
	err := mysql.DB.WithContext(ctx).Table(m.TableName()).Where("id IN ?", ids).Find(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

// GetUserWxIds 获取微信Id
func (m *User) GetUserWxIds(ctx context.Context, ids []uint64) ([]string, error) {
	var wxData []string
	data, err := m.GetUsersByUserIds(ctx, ids)
	if err != nil {
		return nil, err
	}
	for _, v := range data {
		wxData = append(wxData, v.UserId)
	}
	return wxData, nil
}

// GetUserEmails 获取用户EMAIL
func (m *User) GetUserEmails(ctx context.Context, ids []uint64) ([]string, error) {
	var emailData []string
	data, err := m.GetUsersByUserIds(ctx, ids)
	if err != nil {
		return nil, err
	}
	for _, v := range data {
		emailData = append(emailData, v.Email)
	}
	return emailData, nil
}

// GetUserPhone 获取用户手机号码
func (m *User) GetUserPhone(ctx context.Context, ids []uint64) ([]string, error) {
	var phoneData []string
	data, err := m.GetUsersByUserIds(ctx, ids)
	if err != nil {
		return nil, err
	}
	for _, v := range data {
		phoneData = append(phoneData, v.Phone)
	}
	return phoneData, nil
}

// List 用户列表
func (m *User) List(ctx context.Context, page uint64, pageSize uint64, scopes ...func(*gorm.DB) *gorm.DB) ([]User, int64, error) {
	var err error
	var data []User

	var i int64
	err = mysql.DB.WithContext(ctx).Table(m.TableName()).Scopes(scopes...).Count(&i).Error
	if err != nil {
		return nil, 0, err
	}
	scopes = append(scopes, OrderScope("id DESC"), Paginate(page, pageSize))
	err = mysql.DB.WithContext(ctx).Table(m.TableName()).Scopes(scopes...).Find(&data).Error
	if err != nil {
		return nil, 0, err
	}
	return data, i, nil
}

// Delete 删除用户
func (m *User) Delete(ctx context.Context, id uint64) error {
	err := mysql.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Table(m.TableName()).Where("id", id).Delete(User{}).Error; err != nil {
			return err
		}
		if err := tx.Table(mysql.TeamUserTable).Where("user_id", id).Delete(TeamUser{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

func (m *User) VerifyUser(ctx context.Context, username string, password string) (bool, error) {
	err := mysql.DB.WithContext(ctx).Table(m.TableName()).Where("username = ? AND password = ?", username, password).First(&m).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
