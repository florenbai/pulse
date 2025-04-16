package response

import (
	"darkhawk/hertz_gen/base"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
)

type Response struct {
	Code    uint32      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// SendBaseResp build baseResp from error
func SendBaseResp(c *app.RequestContext, err error) {
	if err == nil {
		c.JSON(http.StatusOK, baseResp(Success))
		return
	}
	e := ErrNo{}
	if errors.As(err, &e) {
		c.JSON(http.StatusOK, baseResp(e))
		return
	}
	s := ServiceErr.WithMessage(err.Error())
	c.JSON(http.StatusOK, baseResp(s))
	return
}

// SendDataResp pack response
func SendDataResp(c *app.RequestContext, err error, data interface{}) {
	Err := ConvertErr(err)
	c.JSON(http.StatusOK, Response{
		Code:    Err.ErrCode,
		Message: Err.ErrMsg,
		Data:    data,
	})
	return
}

func baseResp(err ErrNo) *base.BaseResp {
	return &base.BaseResp{Code: err.GetErrCode(), Message: err.GetErrMsg()}
}
