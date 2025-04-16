package model

import (
	"context"
	"darkhawk/app/biz/dal/mysql"
	"darkhawk/pkg/response"
	"darkhawk/pkg/utils"
	"database/sql/driver"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"strings"
	"time"
)

type Model struct {
	Id        uint64     `gorm:"column:id;primary_key;comment:主键id" json:"id" form:"id"` // 主键
	CreatedAt *LocalTime `gorm:"column:created_at;not null;type:datetime;default:current_timestamp;comment:创建时间" json:"createdAt" form:"createdAt" copier:"must"`
	UpdatedAt *LocalTime `gorm:"column:updated_at;not null;type:datetime;default:current_timestamp on update CURRENT_TIMESTAMP;comment:更新时间" json:"updatedAt"  form:"updatedAt"`
}

type LocalTime time.Time
type LocalDate time.Time

func (t *LocalDate) MarshalJSON() ([]byte, error) {
	tTime := time.Time(*t)
	return []byte(fmt.Sprintf("\"%v\"", tTime.Format("2006-01-02"))), nil
}

func (t *LocalDate) Value() (driver.Value, error) {
	if t == nil {
		return nil, nil
	}
	var zeroTime time.Time
	tlt := time.Time(*t)
	// 判断给定时间是否和默认零时间的时间戳相同
	if tlt.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return tlt, nil
}

func (t *LocalDate) Scan(v interface{}) error {
	if value, ok := v.(time.Time); ok {
		*t = LocalDate(value)
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

func (t *LocalTime) MarshalJSON() ([]byte, error) {
	tTime := time.Time(*t)
	return []byte(fmt.Sprintf("\"%v\"", tTime.Format("2006-01-02 15:04:05"))), nil
}

func (t *LocalTime) Value() (driver.Value, error) {
	if t == nil {
		return nil, nil
	}
	var zeroTime time.Time
	tlt := time.Time(*t)
	// 判断给定时间是否和默认零时间的时间戳相同
	if tlt.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return tlt, nil
}

func (t *LocalTime) Scan(v interface{}) error {
	if value, ok := v.(time.Time); ok {
		*t = LocalTime(value)
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

func (t *LocalTime) Format() string {
	return fmt.Sprintf("%s", time.Time(*t).Format("2006-01-02 15:04:05"))
}

func (t *LocalTime) ToTime() time.Time {
	return time.Time(*t)
}

func (t *LocalTime) String() string {
	return fmt.Sprintf("%s", time.Time(*t).String())
}

func NewScopeContainer() []func(db *gorm.DB) *gorm.DB {
	return []func(db *gorm.DB) *gorm.DB{}
}

func GetDB() *gorm.DB {
	return mysql.DB
}

// DebugScope debug scope
func DebugScope() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Debug()
	}
}

// SelectScope 查询函数
func SelectScope(sel string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Select(sel)
	}
}

func GroupScope(group string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Group(group)
	}
}

func JoinScope(query string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Joins(query)
	}
}

// Paginate 分页处理函数
func Paginate(page uint64, pageSize uint64) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page <= 0 {
			page = 1
		}
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (int(page) - 1) * int(pageSize)
		return db.Offset(offset).Limit(int(pageSize))
	}
}

// ParamWithScope 参数配置生成 scope
// associate 映射（如无映射则以param map key为查询字段）
// exclude 排除
// emptyIn 空字符串是否包含查询
func ParamWithScope(paramMap map[string]string, associate map[string]string, exclude []string, emptyIn bool) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		for k, v := range paramMap {
			if len(exclude) > 0 && utils.InOfT(k, exclude) {
				continue
			}
			if v == "" && !emptyIn {
				continue
			}
			if i, ok := associate[k]; ok {
				if strings.Contains(i, "LIKE") {
					db.Where(i, "%"+v+"%")
				} else {
					db.Where(i, v)
				}

			} else {
				db.Where(k, v)
			}
		}
		return db
	}
}

// NoneScope 空scope，避免无scope nil的问题
func NoneScope() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db
	}
}

// WhereWithScope 根据条件生成并行scope
func WhereWithScope(key string, value any) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(key, value)
	}
}

func TimeRangeScope(key string, value []string) func(db *gorm.DB) *gorm.DB {
	if len(value) != 2 {
		return func(db *gorm.DB) *gorm.DB {
			return db
		}
	}
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(key+" >= ? AND "+key+" <= ?", value[0], value[1])
	}
}

// InWithScope in操作
func InWithScope(key string, value any) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(key+" IN ?", value)
	}
}

// LikeWithScope 根据条件生成Like scope
func LikeWithScope(key string, value any) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(key+" LIKE ?", "%"+fmt.Sprintf("%v", value)+"%")
	}
}

// OrWithScope 根据条件生成或scope
func OrWithScope(key string, value any) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Or(key, value)
	}
}

// OrderScope 排序scope
func OrderScope(value interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Order(value)
	}
}

// GetOneWithScope 获取一条数据并支持scope
func GetOneWithScope(ctx context.Context, table string, out interface{}, scopes ...func(*gorm.DB) *gorm.DB) error {
	err := mysql.DB.WithContext(ctx).Table(table).Scopes(scopes...).First(out).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return err
	}
	return nil
}

// GetOneById 根据id获取数据
func GetOneById(ctx context.Context, table string, id uint64, out interface{}) error {
	err := mysql.DB.WithContext(ctx).Table(table).Where("id = ?", id).First(out).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.DataNotFoundErr
		}
		return err
	}
	return nil
}

// EditOneById 根据id编辑数据
func EditOneById(ctx context.Context, table string, id uint64, updates interface{}) error {
	err := mysql.DB.WithContext(ctx).Table(table).Where("id = ?", id).Updates(updates).Error
	if err != nil {
		return err
	}
	return nil
}

// EditByScopes 根据scope编辑数据
func EditByScopes(ctx context.Context, table string, updates interface{}, scopes ...func(*gorm.DB) *gorm.DB) error {
	err := mysql.DB.WithContext(ctx).Table(table).Scopes(scopes...).Updates(updates).Error
	if err != nil {
		return err
	}
	return nil
}

// Create 创建数据
func Create(ctx context.Context, table string, out interface{}) error {
	err := mysql.DB.Debug().WithContext(ctx).Table(table).Create(out).Error
	if err != nil {
		return err
	}
	return nil

}

// Transaction 事务
func Transaction(ctx context.Context, f func(tx *gorm.DB) error) error {
	return mysql.DB.WithContext(ctx).Transaction(f)
}

// DeleteByScope 根据scope删除数据
func DeleteByScope(ctx context.Context, table string, out interface{}, scopes ...func(*gorm.DB) *gorm.DB) error {
	err := mysql.DB.WithContext(ctx).Table(table).Scopes(scopes...).Delete(out).Error
	if err != nil {
		return err
	}
	return nil
}

// DeleteOneById 根据id删除一条数据
func DeleteOneById(ctx context.Context, table string, id uint64, model interface{}) error {
	return mysql.DB.WithContext(ctx).Table(table).Where("id = ?", id).Delete(&model).Error
}

// GetPageList 获取分页列表
func GetPageList(ctx context.Context, table string, page uint64, pageSize uint64, out interface{}, scopes ...func(*gorm.DB) *gorm.DB) (int64, error) {
	var i int64
	err := mysql.DB.WithContext(ctx).Table(table).Scopes(scopes...).Count(&i).Error
	if err != nil {
		return 0, err
	}
	scopes = append(scopes, Paginate(page, pageSize))
	err = mysql.DB.WithContext(ctx).Table(table).Scopes(scopes...).Find(out).Error
	if err != nil {
		return i, err
	}
	return i, nil
}

// GetAll 获取所有数据
func GetAll(ctx context.Context, table string, out interface{}, scopes ...func(*gorm.DB) *gorm.DB) error {
	return mysql.DB.WithContext(ctx).Table(table).Scopes(scopes...).Find(out).Error
}

// GetPluck 获取数组数据
func GetPluck(ctx context.Context, table string, column string, out interface{}, scopes ...func(*gorm.DB) *gorm.DB) error {
	return mysql.DB.WithContext(ctx).Table(table).Scopes(scopes...).Pluck(column, out).Error
}

// GetCountWithScope 获取数量并支持scope
func GetCountWithScope(ctx context.Context, table string, scopes ...func(*gorm.DB) *gorm.DB) (int64, error) {
	var i int64
	err := mysql.DB.WithContext(ctx).Table(table).Scopes(scopes...).Count(&i).Error
	if err != nil {
		return i, err
	}
	return i, nil
}

func Error(err error, nullError bool) bool {
	notFound := errors.Is(err, gorm.ErrRecordNotFound)
	if notFound && nullError == false {
		return false
	} else {
		return true
	}
}

// ExistName 名称是否存在
func ExistName(ctx context.Context, tableName string, field string, name string, id *uint64) (bool, error) {
	var i int64
	var err error
	if id == nil {
		err = mysql.DB.WithContext(ctx).Table(tableName).Where(field, name).Count(&i).Error
	} else {
		err = mysql.DB.WithContext(ctx).Table(tableName).Where(field, name).Where("id != ?", *id).Count(&i).Error
	}

	if err != nil {
		return true, err
	}
	if i > 0 {
		return true, nil
	}
	return false, nil
}
