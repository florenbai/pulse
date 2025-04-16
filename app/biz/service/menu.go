package service

import (
	"context"
	"darkhawk/app/biz/model"
	"darkhawk/conf"
	pb "darkhawk/hertz_gen/menu"
	"darkhawk/pkg/response"
	"darkhawk/pkg/utils"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/sessions"
)

type MenuService struct {
	ctx context.Context
	c   *app.RequestContext
}

func NewMenuService(ctx context.Context, c *app.RequestContext) *MenuService {
	return &MenuService{
		ctx: ctx,
		c:   c,
	}
}

func (service *MenuService) CreateMenu(req pb.MenuRequest) error {
	menuModel := new(model.Menu)
	ok, err := model.ExistName(service.ctx, menuModel.TableName(), "name", req.GetName(), nil)
	if err != nil {
		return err
	}
	if ok {
		return response.NameAlreadyExistErr
	}
	if err = utils.AnyToAny(req, &menuModel); err != nil {
		return err
	}

	err = model.Create(service.ctx, menuModel.TableName(), &menuModel)
	if err != nil {
		return err
	}
	enabled := true
	menuModel.Sort = int64(menuModel.Id)
	menuModel.Enabled = &enabled
	return model.EditOneById(service.ctx, menuModel.TableName(), menuModel.Id, &menuModel)
}

func (service *MenuService) EditMenu(req pb.MenuRequest, id uint64) error {
	menuModel := new(model.Menu)
	err := model.GetOneById(service.ctx, menuModel.TableName(), id, &menuModel)
	if err != nil {
		return err
	}
	ok, err := model.ExistName(service.ctx, menuModel.TableName(), "name", req.GetName(), &id)
	if err != nil {
		return err
	}
	if ok {
		return response.NameAlreadyExistErr
	}
	if err = utils.AnyToAny(req, &menuModel); err != nil {
		return err
	}

	return model.EditOneById(service.ctx, menuModel.TableName(), menuModel.Id, &menuModel)
}

func (service *MenuService) DeleteMenu(id uint64) error {
	menuModel := new(model.Menu)
	err := model.GetOneById(service.ctx, menuModel.TableName(), id, &menuModel)
	if err != nil {
		return err
	}
	return menuModel.DeleteMenuAndChildren(service.ctx, id)
}

func (service *MenuService) GetUserMenu() ([]model.MenuWithChildren, error) {
	var menus []model.Menu
	menuModel := new(model.Menu)
	session := sessions.Default(service.c)
	uid := session.Get("uid").(uint64)

	var teams []uint64
	userTeamModel := new(model.TeamUser)
	err := model.GetPluck(service.ctx, userTeamModel.TableName(), "team_id", &teams, model.WhereWithScope("user_id", uid))
	if err != nil {
		return nil, err
	}

	var enableMenus []uint64
	teamRuleModel := new(model.TeamRule)
	err = model.GetPluck(service.ctx, teamRuleModel.TableName(), "menu_id", &enableMenus, model.InWithScope("team_id", teams))
	if err != nil {
		return nil, err
	}
	if conf.GetConf().Server.RemotePrometheus {
		err = model.GetAll(service.ctx, menuModel.TableName(), &menus, model.OrderScope("sort ASC"), model.WhereWithScope("enabled", true), model.InWithScope("id", enableMenus))
	} else {
		err = model.GetAll(service.ctx, menuModel.TableName(), &menus, model.OrderScope("sort ASC"), model.WhereWithScope("enabled", true), model.InWithScope("id", enableMenus), model.WhereWithScope("route_name != ?", "alert-rule-list"))
	}

	if err != nil {
		return nil, err
	}

	data := service.HandleUserMenu(menus, 0)
	return data, nil
}

func (service *MenuService) GetMenuNameList() ([]model.MenuWithChildren, error) {
	var menus []model.Menu
	menuModel := new(model.Menu)
	err := model.GetAll(service.ctx, menuModel.TableName(), &menus, model.OrderScope("sort ASC"))
	if err != nil {
		return nil, err
	}

	data := service.HandleMenuName(menus, 0)
	return data, nil
}

func (service *MenuService) GetMenuList() ([]model.MenuList, error) {
	var menus []model.Menu
	menuModel := new(model.Menu)
	err := model.GetAll(service.ctx, menuModel.TableName(), &menus, model.OrderScope("sort ASC"))
	if err != nil {
		return nil, err
	}

	data := service.HandleMenuList(menus, 0)
	return data, nil
}

func (service *MenuService) HandleMenuList(menus []model.Menu, parent uint64) []model.MenuList {
	var data []model.MenuList
	for _, v := range menus {
		if parent == v.Pid {
			m := model.MenuList{
				Menu:     v,
				Children: []model.MenuList{},
			}

			children := service.HandleMenuList(menus, v.Id)
			if len(children) > 0 {
				m.Children = children
			}
			data = append(data, m)
		}
	}
	return data
}

func (service *MenuService) HandleMenuName(menus []model.Menu, parent uint64) []model.MenuWithChildren {
	var data []model.MenuWithChildren
	for _, v := range menus {
		if parent == v.Pid {
			m := model.MenuWithChildren{
				Key: v.Id,
				MenuItem: model.MenuItem{
					Name: v.Name,
				},
			}

			children := service.HandleMenuName(menus, v.Id)
			if len(children) > 0 {
				m.Children = children
			}
			data = append(data, m)
		}

	}
	return data
}

func (service *MenuService) HandleUserMenu(menus []model.Menu, parent uint64) []model.MenuWithChildren {
	var data []model.MenuWithChildren
	for _, v := range menus {
		if parent == v.Pid {
			m := model.MenuWithChildren{
				Key: v.Id,
				MenuItem: model.MenuItem{
					Name:    v.RouteName,
					Path:    v.RoutePath,
					Sort:    v.Sort,
					Enabled: v.Enabled,
					Meta: model.Meta{
						Roles:              []string{"*"},
						Icon:               v.Icon,
						Locale:             v.Locale,
						IgnoreCache:        *v.Cache,
						RequiresAuth:       *v.RequiresAuth,
						HideInMenu:         v.HideMenu,
						HideChildrenInMenu: v.HideChildMenu,
						ActiveMenu:         v.ActiveMenu,
						Order:              v.Sort,
						NoAffix:            v.Affix,
					},
				},
			}

			children := service.HandleUserMenu(menus, v.Id)
			if len(children) > 0 {
				m.Children = children
			}
			data = append(data, m)
		}

	}
	return data
}

func (service *MenuService) ChangeMenuSort(req []uint64) error {
	menuModel := new(model.Menu)
	for k, v := range req {
		err := model.EditOneById(service.ctx, menuModel.TableName(), v, model.Menu{Sort: int64(k + 1)})
		if err != nil {
			return err
		}
	}
	return nil
}

func (service *MenuService) ChangeMenuStatus(id uint64) error {
	var t bool
	menuModel := new(model.Menu)
	err := model.GetOneById(service.ctx, menuModel.TableName(), id, &menuModel)
	if err != nil {
		return err
	}
	if *menuModel.Enabled {
		t = false
	} else {
		t = true
	}
	return model.EditOneById(service.ctx, menuModel.TableName(), menuModel.Id, model.Menu{Enabled: &t})
}
