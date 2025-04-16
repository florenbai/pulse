package response

import (
	"errors"
	"fmt"
)

type ErrNo struct {
	ErrCode uint32
	ErrMsg  string
}

const (
	Err_Success             = 20000
	Err_Unauthenticated     = 10000
	Err_Unauthorized        = 10001
	Err_AuthorizeFail       = 10002
	Err_NoLic               = 10003
	Err_BadRequest          = 40000
	Err_ParamsErr           = 40001
	Err_DataNotFoundErr     = 40002
	Err_ServerNotFound      = 40003
	Err_RequestServerFail   = 40004
	Err_BindAndValidateFail = 40006
	Err_ServerHandleFail    = 50001
	Err_ServerInternalErr   = 50002
	Err_RPCAuthSrvErr       = 50003
	Err_AuthSrvErr          = 50004
	Err_RPCBlobSrvErr       = 50005
	Err_BlobSrvErr          = 50006
	Err_RPCCarSrvErr        = 60000
	Err_CarSrvErr           = 60001
	Err_RPCProfileSrvErr    = 60002
	Err_ProfileSrvErr       = 60003
	Err_RPCTripSrvErr       = 60004
	Err_TripSrvErr          = 70000
	Err_RecordAlreadyExist  = 70001
	Err_DirtyData           = 70003
)

var (
	Success                   = NewErrNo(Err_Success, "操作成功")
	ServiceErr                = NewErrNo(Err_ServerInternalErr, "服务没有正常启动")
	ParamErr                  = NewErrNo(Err_ParamsErr, "参数错误")
	Err                       = NewErrNo(Err_ServerInternalErr, "服务器内部错误")
	OidcErr                   = NewErrNo(Err_Unauthenticated, "未从统一认证服务器获取用户信息")
	OidcCodeErr               = NewErrNo(Err_Unauthenticated, "统一认证服务器返回参数不正确")
	LoginExceedErr            = NewErrNo(Err_Unauthenticated, "您的登录凭证已过期")
	DataNotFoundErr           = NewErrNo(Err_DataNotFoundErr, "您访问的数据不存在")
	UnauthorizedErr           = NewErrNo(Err_Unauthorized, "您没有权限执行该操作")
	AuthorizeFailErr          = NewErrNo(Err_AuthorizeFail, "认证失败")
	RedirectNotFoundErr       = NewErrNo(Err_ParamsErr, "重新定向Url参数不存在")
	NameAlreadyExistErr       = NewErrNo(Err_ParamsErr, "您输入的名称已存在")
	LevelNotFoundErr          = NewErrNo(Err_ParamsErr, "您输入的告警等级不存在")
	AlertLevelAlreadyExistErr = NewErrNo(Err_ParamsErr, "您输入的告警等级已存在")
	ChannelStatusCloseErr     = NewErrNo(Err_ServerHandleFail, "渠道已被禁用 ")
	JsonParamErr              = NewErrNo(Err_ParamsErr, "这不是一个Json格式数据")
	NodeNotFound              = NewErrNo(Err_ParamsErr, "节点不存在")

	WorkspaceExistErr = NewErrNo(Err_BadRequest, "工作区已经存在")
	ScheduleExistErr  = NewErrNo(Err_BadRequest, "值班名称已经存在")
	TeamExistErr      = NewErrNo(Err_BadRequest, "团队名称已经存在")
	TeamNotFoundErr   = NewErrNo(Err_BadRequest, "团队不存在")
)

func (e ErrNo) Error() string {
	return fmt.Sprintf("err_code=%d, err_msg=%s", e.ErrCode, e.ErrMsg)
}

func (e ErrNo) WithMessage(msg string) ErrNo {
	e.ErrMsg = msg
	return e
}

func NewErrNo(code uint32, msg string) ErrNo {
	return ErrNo{code, msg}
}

// ConvertErr convert error to Errno
func ConvertErr(err error) ErrNo {
	Err := ErrNo{}
	if errors.As(err, &Err) {
		return Err
	}

	s := ServiceErr
	s.ErrMsg = err.Error()
	return s
}

func (e ErrNo) GetErrMsg() string {
	return e.ErrMsg
}

func (e ErrNo) GetErrCode() uint32 {
	return e.ErrCode
}
