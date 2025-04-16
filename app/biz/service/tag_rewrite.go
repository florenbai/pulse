package service

import (
	"context"
	"darkhawk/app/biz/model"
	pb "darkhawk/hertz_gen/integration"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/sessions"
	"regexp"
)

type TagRewriteService struct {
	ctx context.Context
	c   *app.RequestContext
}

func NewTagRewriteService(ctx context.Context, c *app.RequestContext) *TagRewriteService {
	return &TagRewriteService{
		ctx: ctx,
		c:   c,
	}
}

func (service *TagRewriteService) CreateTagRewrite(id uint64, req pb.TagRewriteRequest) error {
	integrationModel := new(model.Integration)
	err := model.GetOneById(service.ctx, integrationModel.TableName(), id, &integrationModel)
	if err != nil {
		return err
	}
	if req.GetRewriteItem().Expression != nil {
		_, err := regexp.Compile(req.GetRewriteItem().GetExpression())
		if err != nil {
			return fmt.Errorf("您的表达式不是有效正则表达式：%s", err.Error())
		}
	}
	var filters []model.TagFilter
	for _, f := range req.GetFilters() {
		filter := model.TagFilter{
			Tag:       f.GetTag(),
			Operation: f.GetOperation(),
			Value:     f.GetValue(),
		}
		filters = append(filters, filter)
	}
	session := sessions.Default(service.c)
	uid := session.Get("uid")

	tagData := model.TagRewrite{
		Filters:       filters,
		IntegrationId: id,
		RewriteItem: model.RewriteItem{
			OldTag:     req.GetRewriteItem().GetOldTag(),
			Expression: req.GetRewriteItem().GetExpression(),
			NewTag:     req.GetRewriteItem().GetNewTag(),
			Value:      req.GetRewriteItem().GetValue(),
			DeleteTag:  req.GetRewriteItem().GetDeleteTag(),
		},
		RewriteType: model.RewriteType(req.RewriteType),
		Uid:         uid.(uint64),
	}
	return model.Create(service.ctx, tagData.TableName(), &tagData)
}

func (service *TagRewriteService) UpdateTagRewrite(id uint64, req pb.TagRewriteRequest) error {
	tagModel := new(model.TagRewrite)
	err := model.GetOneById(service.ctx, tagModel.TableName(), id, &tagModel)
	if err != nil {
		return err
	}
	if req.GetRewriteItem().Expression != nil {
		_, err := regexp.Compile(req.GetRewriteItem().GetExpression())
		if err != nil {
			return fmt.Errorf("您的表达式不是有效正则表达式：%s", err.Error())
		}
	}
	var filters []model.TagFilter
	for _, f := range req.GetFilters() {
		filter := model.TagFilter{
			Tag:       f.GetTag(),
			Operation: f.GetOperation(),
			Value:     f.GetValue(),
		}
		filters = append(filters, filter)
	}
	session := sessions.Default(service.c)
	uid := session.Get("uid")

	tagModel.Filters = filters
	tagModel.RewriteItem = model.RewriteItem{
		OldTag:     req.GetRewriteItem().GetOldTag(),
		Expression: req.GetRewriteItem().GetExpression(),
		NewTag:     req.GetRewriteItem().GetNewTag(),
		Value:      req.GetRewriteItem().GetValue(),
		DeleteTag:  req.GetRewriteItem().GetDeleteTag(),
	}
	tagModel.RewriteType = model.RewriteType(req.RewriteType)
	tagModel.Uid = uid.(uint64)
	return model.EditOneById(service.ctx, tagModel.TableName(), id, &tagModel)
}

func (service *TagRewriteService) TagRewriteList(id uint64) ([]model.TagRewrite, error) {
	var data []model.TagRewrite
	tagModel := new(model.TagRewrite)
	err := model.GetAll(service.ctx, tagModel.TableName(), &data, model.WhereWithScope("integration_id", id))
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (service *TagRewriteService) DeleteTagRewrite(id uint64) error {
	tagModel := new(model.TagRewrite)
	err := model.GetOneById(service.ctx, tagModel.TableName(), id, &tagModel)
	if err != nil {
		return err
	}
	return model.DeleteOneById(service.ctx, tagModel.TableName(), id, model.TagRewrite{})
}
