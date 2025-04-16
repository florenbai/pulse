package service

import (
	"context"
	"darkhawk/app/biz/model"
	"darkhawk/hertz_gen/base"
	pb "darkhawk/hertz_gen/team"
	"darkhawk/pkg/response"
	"darkhawk/pkg/utils"
	"encoding/json"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/sessions"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/structpb"
)

type TeamService struct {
	ctx context.Context
	c   *app.RequestContext
}

func NewTeamService(ctx context.Context, c *app.RequestContext) *TeamService {
	return &TeamService{
		ctx: ctx,
		c:   c,
	}
}

func (service *TeamService) CreateTeam(req pb.TeamRequest) error {
	teamModel := new(model.Teams)
	ok, err := model.ExistName(service.ctx, teamModel.TableName(), "team_name", req.GetTeamName(), nil)
	if err != nil {
		return err
	}
	if ok {
		return response.TeamExistErr
	}
	if err = utils.AnyToAny(req, &teamModel); err != nil {
		return err
	}

	session := sessions.Default(service.c)
	uid := session.Get("uid").(uint64)
	teamModel.UpdatedBy = uid
	teamModel.DataPermission = 1
	return teamModel.Create(service.ctx, req.GetTeamMembers())
}

func (service *TeamService) UpdateTeam(id uint64, req pb.TeamRequest) error {
	teamModel := new(model.Teams)
	err := model.GetOneById(service.ctx, teamModel.TableName(), id, &teamModel)
	if err != nil {
		return err
	}
	ok, err := model.ExistName(service.ctx, teamModel.TableName(), "team_name", req.GetTeamName(), &id)
	if err != nil {
		return err
	}
	if ok {
		return response.TeamExistErr
	}
	if err = utils.AnyToAny(req, &teamModel); err != nil {
		return err
	}

	session := sessions.Default(service.c)
	uid := session.Get("uid").(uint64)
	teamModel.UpdatedBy = uid
	return teamModel.Update(service.ctx, req.GetTeamMembers())
}

func (service *TeamService) GetTeamList(req pb.TeamQuery) (*base.BaseListResp, error) {
	search := map[string]string{
		"team_name": req.GetTeamName(),
	}
	associate := map[string]string{
		"team_name": "teams.team_name LIKE ?",
	}
	scopes := model.ParamWithScope(search, associate, nil, false)
	teamModel := new(model.Teams)
	data, i, err := teamModel.List(service.ctx, req.GetPage(), req.GetPageSize(), scopes)
	if err != nil {
		return nil, err
	}

	listValue := new(structpb.ListValue)
	b, _ := json.Marshal(data)
	err = protojson.Unmarshal(b, listValue)
	if err != nil {
		return nil, err
	}
	return &base.BaseListResp{
		Total: uint64(i),
		List:  listValue,
	}, nil
}

func (service *TeamService) DeleteTeam(id uint64) error {
	teamModel := new(model.Teams)
	err := model.GetOneById(service.ctx, teamModel.TableName(), id, &teamModel)
	if err != nil {
		return err
	}
	return teamModel.Delete(service.ctx, id)
}

func (service *TeamService) ChangeTeamStatus(id uint64) error {
	teamModel := new(model.Teams)
	err := model.GetOneById(service.ctx, teamModel.TableName(), id, &teamModel)
	if err != nil {
		return err
	}
	return teamModel.ChangeStatus(service.ctx, id)
}

func (service *TeamService) AuthorizationTeam(id uint64, req pb.AuthorizationRequest) error {
	teamModel := new(model.Teams)
	err := model.GetOneById(service.ctx, teamModel.TableName(), id, &teamModel)
	if err != nil {
		return err
	}
	teamModel.DataPermission = req.GetWorkspace()
	err = model.EditOneById(service.ctx, teamModel.TableName(), id, &teamModel)
	if err != nil {
		return err
	}

	teamRuleModel := new(model.TeamRule)
	return teamRuleModel.AuthorizationMenu(service.ctx, id, req.GetMenus())
}

func (service *TeamService) AuthorizationInfo(id uint64) (*pb.AuthorizationRequest, error) {
	var data pb.AuthorizationRequest
	teamModel := new(model.Teams)
	teamRuleModel := new(model.TeamRule)
	err := model.GetOneById(service.ctx, teamModel.TableName(), id, &teamModel)
	if err != nil {
		return nil, err
	}
	err = model.GetPluck(service.ctx, teamRuleModel.TableName(), "menu_id", &data.Menus, model.WhereWithScope("team_id", id))
	if err != nil {
		return nil, err
	}
	data.Workspace = teamModel.DataPermission
	return &data, nil
}
