package model

import (
	"context"
	"darkhawk/app/biz/dal/mysql"
)

type Menu struct {
	Model
	Name          string `json:"name"`
	Desc          string `json:"desc"`
	Type          int8   `json:"type"`
	Pid           uint64 `json:"pid"`
	Sort          int64  `json:"sort"`
	Icon          string `json:"icon"`
	Locale        string `json:"locale"`
	RoutePath     string `json:"routePath"`
	RouteName     string `json:"routeName"`
	Redirect      string `json:"redirect"`
	Permission    string `json:"permission"`
	RequiresAuth  *bool  `json:"requiresAuth"`
	HideMenu      *bool  `json:"hideMenu"`
	HideChildMenu *bool  `json:"hideChildMenu"`
	IsLink        *bool  `json:"isLink"`
	ActiveMenu    *bool  `json:"activeMenu"`
	Affix         *bool  `json:"affix"`
	Cache         *bool  `json:"cache"`
	Enabled       *bool  `json:"enabled"`
}

type MenuList struct {
	Menu
	Children []MenuList `json:"children,omitempty"`
}

type MenuWithChildren struct {
	Key uint64 `json:"key"`
	MenuItem
	Children []MenuWithChildren `json:"children,omitempty"`
}

type MenuItem struct {
	Name    string `json:"name"`
	Path    string `json:"path"`
	Meta    Meta   `json:"meta"`
	Sort    int64  `json:"sort"`
	Enabled *bool  `json:"enabled"`
}

type Meta struct {
	Roles              []string `json:"roles"`
	Icon               string   `json:"icon"`
	Id                 uint64   `json:"id"`
	Locale             string   `json:"locale,omitempty"`
	IgnoreCache        bool     `json:"ignoreCache,omitempty"`
	RequiresAuth       bool     `json:"requiresAuth,omitempty"`
	HideInMenu         *bool    `json:"hideInMenu,omitempty"`
	HideChildrenInMenu *bool    `json:"hideChildrenInMenu,omitempty"`
	ActiveMenu         *bool    `json:"activeMenu"`
	Order              int64    `json:"order,omitempty"`
	NoAffix            *bool    `json:"noAffix,omitempty"`
}

func (m *Menu) TableName() string {
	return mysql.MenuTable
}

func (m *Menu) DeleteMenuAndChildren(ctx context.Context, id uint64) error {

	deleteId, err := m.GetMenuIdAndChildren(ctx, []uint64{id})
	if err != nil {
		return err
	}
	deleteId = append(deleteId, id)
	err = mysql.DB.WithContext(ctx).Table(m.TableName()).Where("id IN ?", deleteId).Delete(Menu{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (m *Menu) GetMenuIdAndChildren(ctx context.Context, ids []uint64) ([]uint64, error) {
	var deleteIds []uint64
	var cids []uint64
	err := mysql.DB.WithContext(ctx).Table(m.TableName()).Where("pid IN ?", ids).Pluck("id", &cids).Error
	if err != nil {
		return nil, err
	}
	deleteIds = append(deleteIds, cids...)
	if len(cids) > 0 {
		pids, err := m.GetMenuIdAndChildren(ctx, cids)
		if err != nil {
			return nil, err
		}
		deleteIds = append(deleteIds, pids...)
	}
	return deleteIds, nil

}
