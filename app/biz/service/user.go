package service

import (
	"context"
	"crypto/md5"
	"darkhawk/app/biz/model"
	"darkhawk/conf"
	"darkhawk/hertz_gen/base"
	pb "darkhawk/hertz_gen/user"
	"darkhawk/pkg/jwt"
	"darkhawk/pkg/response"
	"darkhawk/pkg/utils"
	"encoding/hex"
	"encoding/json"
	"github.com/cloudwego/hertz/pkg/app"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/structpb"
)

type UserService struct {
	ctx context.Context
	c   *app.RequestContext
}

func NewUserService(ctx context.Context, c *app.RequestContext) *UserService {
	return &UserService{
		ctx: ctx,
		c:   c,
	}
}

func (service *UserService) VerifyUserInfo(username string, password string) (string, *model.User, error) {
	// 计算密码MD5
	h := md5.New()
	h.Write([]byte(password))
	encryptedPassword := hex.EncodeToString(h.Sum(nil))

	userModel := new(model.User)
	ok, err := userModel.VerifyUser(service.ctx, username, encryptedPassword)
	if err != nil {
		return "", nil, err
	}
	if !ok {
		return "", nil, response.AuthorizeFailErr
	}
	token, err := jwt.GenerateToken(userModel.Id, userModel.Username, []byte(conf.GetConf().Authentication.AuthSecret))
	if err != nil {
		return "", nil, err
	}
	return token, userModel, nil
}

func (service *UserService) CreateUser(req pb.UserRequest) error {
	userModel := new(model.User)
	ok, err := model.ExistName(service.ctx, userModel.TableName(), "username", req.GetUsername(), nil)
	if err != nil {
		return err
	}
	if ok {
		return response.NameAlreadyExistErr
	}
	if err = utils.AnyToAny(req, &userModel); err != nil {
		return err
	}
	return model.Create(service.ctx, userModel.TableName(), &userModel)
}

func (service *UserService) UpdateUser(id uint64, req pb.UserRequest) error {
	userModel := new(model.User)
	ok, err := model.ExistName(service.ctx, userModel.TableName(), "username", req.GetUsername(), &id)
	if err != nil {
		return err
	}
	if ok {
		return response.NameAlreadyExistErr
	}
	if err = utils.AnyToAny(req, &userModel); err != nil {
		return err
	}
	return model.EditOneById(service.ctx, userModel.TableName(), id, &userModel)
}

func (service *UserService) GetUserList(req pb.UserQuery) (*base.BaseListResp, error) {
	search := map[string]string{
		"nickname": req.GetNickName(),
	}
	associate := map[string]string{
		"nickname": "user.username LIKE ?",
	}
	scopes := model.ParamWithScope(search, associate, nil, false)
	userModel := new(model.User)
	data, i, err := userModel.List(service.ctx, req.GetPage(), req.GetPageSize(), scopes)
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

func (service *UserService) DeleteUser(id uint64) error {
	userModel := new(model.User)
	err := model.GetOneById(service.ctx, userModel.TableName(), id, &userModel)
	if err != nil {
		return err
	}
	return userModel.Delete(service.ctx, id)
}

func (service *UserService) UserLogin(nickname string, username string, email string) (*model.User, error) {
	userModel := new(model.User)
	err := model.GetOneWithScope(service.ctx, userModel.TableName(), &userModel, model.WhereWithScope("username", username))
	if err != nil {
		return nil, err
	}
	if userModel.Id == 0 {
		userModel.Email = email
		userModel.Nickname = nickname
		userModel.Username = username
		err = model.Create(service.ctx, userModel.TableName(), &userModel)
		if err != nil {
			return nil, err
		}
	}

	return userModel, nil
}
