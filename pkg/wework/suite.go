package wework

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"time"
)

// UpdateSuiteTicket 用于接收回调后更新sdk实例的suite ticket
func (ww *weWork) UpdateSuiteTicket(ticket string) {
	ww.suiteTicket = ticket
	ww.cache.conn.Set(context.TODO(), "suiteTicket", ticket, time.Second*600)
	ww.requestSuiteToken()
}

type suiteTokenResponse struct {
	BizResponse
	SuiteAccessToken string `json:"suite_access_token"`
	ExpiresIn        int    `json:"expires_in"`
}

func (ww *weWork) requestSuiteToken() (resp suiteTokenResponse) {
	if ww.suiteTicket == "" {
		resp.ErrCode = 400
		resp.ErrorMsg = "suite ticket 未推送"
		logger.Sugar().Error("suite ticket 未推送")
		return
	}
	apiUrl := "/cgi-bin/service/get_suite_token"
	h := H{}
	h["suite_id"] = ww.suiteId
	h["suite_secret"] = ww.suiteSecret
	h["suite_ticket"] = ww.suiteTicket
	_, err := ww.httpClient.R().SetBody(h).SetResult(&resp).Post(apiUrl)
	if err != nil {
		resp.ErrCode = 400
		resp.ErrorMsg = err.Error()
		logger.Sugar().Error(err)
		return
	}
	return
}

func (ww *weWork) getSuiteAccessToken() (token string) {
	var err error
	token, err = ww.cache.conn.Get(context.TODO(), "suiteToken").Result()
	if err != nil {
		resp := ww.requestSuiteToken()
		if resp.ErrCode != 0 {
			logger.Sugar().Error(resp.ErrorMsg)
			return ""
		}
		token = resp.SuiteAccessToken
		err = ww.cache.conn.Set(context.TODO(), "suiteToken", token, time.Second*time.Duration(resp.ExpiresIn)).Err()
		if err != nil {
			logger.Sugar().Error(err.Error())
			return ""
		}
	}
	return token
}

type GetPreAuthCodeResponse struct {
	BizResponse
	PreAuthCode string `json:"pre_auth_code"`
	ExpiresIn   int    `json:"expires_in"`
}

// GetPreAuthCode 获取预授权码
// https://open.work.weixin.qq.com/api/doc/90001/90143/90601
func (ww *weWork) GetPreAuthCode() (resp GetPreAuthCodeResponse) {
	_, err := ww.httpClient.R().SetQueryParam("suite_access_token", ww.getSuiteAccessToken()).
		SetResult(&resp).Get("/cgi-bin/service/get_pre_auth_code")
	if err != nil {
		logger.Sugar().Error(err)
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
		return
	}
	return
}

type DealerCorpInfo struct {
	Corpid   string `json:"corpid"`
	CorpName string `json:"corp_name"`
}
type AuthCorpInfo struct {
	CorpId            string `json:"corpid"`
	CorpName          string `json:"corp_name"`
	CorpType          string `json:"corp_type"`
	CorpSquareLogoURL string `json:"corp_square_logo_url"`
	CorpUserMax       int    `json:"corp_user_max"`
	CorpAgentMax      int    `json:"corp_agent_max"`
	CorpFullName      string `json:"corp_full_name"`
	VerifiedEndTime   int    `json:"verified_end_time"`
	SubjectType       int    `json:"subject_type"`
	CorpWxqrcode      string `json:"corp_wxqrcode"`
	CorpScale         string `json:"corp_scale"`
	CorpIndustry      string `json:"corp_industry"`
	CorpSubIndustry   string `json:"corp_sub_industry"`
}
type AuthUserInfo struct {
	UserId     string `json:"userid"`
	OpenUserId string `json:"open_userid"`
	Name       string `json:"name"`
	Avatar     string `json:"avatar"`
}
type RegisterCodeInfo struct {
	RegisterCode string `json:"register_code"`
	TemplateID   string `json:"template_id"`
	State        string `json:"state"`
}
type Agent struct {
	AgentId         int    `json:"agentid"`
	Name            string `json:"name"`
	RoundLogoURL    string `json:"round_logo_url"`
	SquareLogoURL   string `json:"square_logo_url"`
	Appid           int    `json:"appid"`
	AuthMode        int    `json:"auth_mode,omitempty"`
	IsCustomizedApp bool   `json:"is_customized_app,omitempty"`
	Privilege       struct {
		Level      int      `json:"level"`
		AllowParty []int    `json:"allow_party"`
		AllowUser  []string `json:"allow_user"`
		AllowTag   []int    `json:"allow_tag"`
		ExtraParty []int    `json:"extra_party"`
		ExtraUser  []string `json:"extra_user"`
		ExtraTag   []int    `json:"extra_tag"`
	} `json:"privilege,omitempty"`
	SharedFrom struct {
		Corpid string `json:"corpid"`
	} `json:"shared_from"`
}
type GetPermanentCodeResponse struct {
	BizResponse
	AccessToken    string         `json:"access_token"`
	ExpiresIn      int            `json:"expires_in"`
	PermanentCode  string         `json:"permanent_code"`
	DealerCorpInfo DealerCorpInfo `json:"dealer_corp_info"`
	AuthCorpInfo   AuthCorpInfo   `json:"auth_corp_info"`
	AuthInfo       struct {
		Agent []Agent `json:"agent"`
	} `json:"auth_info"`
	AuthUserInfo     AuthUserInfo     `json:"auth_user_info"`
	RegisterCodeInfo RegisterCodeInfo `json:"register_code_info"`
	State            string           `json:"state"`
}

// GetPermanentCode 获取企业永久授权码
// https://open.work.weixin.qq.com/api/doc/90001/90143/90603
func (ww *weWork) GetPermanentCode(authCode string) (resp GetPermanentCodeResponse) {
	h := H{}
	h["auth_code"] = authCode
	_, err := ww.httpClient.R().SetQueryParam("suite_access_token", ww.getSuiteAccessToken()).
		SetBody(h).SetResult(&resp).Post("/cgi-bin/service/get_permanent_code")
	if err != nil {
		resp.ErrCode = 400
		resp.ErrorMsg = err.Error()
		logger.Sugar().Error(err)
	}
	return
}

type GetAuthInfoResponse struct {
	BizResponse
	DealerCorpInfo DealerCorpInfo `json:"dealer_corp_info"`
	AuthCorpInfo   AuthCorpInfo   `json:"auth_corp_info"`
	AuthInfo       struct {
		Agent []Agent `json:"agent"`
	} `json:"auth_info"`
}

// GetAuthInfo 获取企业授权信息
// https://open.work.weixin.qq.com/api/doc/90001/90143/90604
func (ww *weWork) GetAuthInfo(authCorpId, permanentCode string) (resp GetAuthInfoResponse) {
	h := H{}
	h["auth_corpid"] = authCorpId
	h["permanent_code"] = permanentCode
	_, err := ww.httpClient.R().SetQueryParam("suite_access_token", ww.getSuiteAccessToken()).
		SetBody(h).SetResult(&resp).Post("/cgi-bin/service/get_auth_info")
	if err != nil {
		resp.ErrCode = 400
		resp.ErrorMsg = err.Error()
		logger.Sugar().Error(err)
	}
	return
}

type getCorpTokenResponse struct {
	BizResponse
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

// 从数据库查询永久授权码和授权企业的企业微信id，获取对应的access token
func (ww *weWork) requestCorpToken(corpId uint) (resp getCorpTokenResponse) {
	var err error
	var corpid, secret string
	var customizedApp bool
	if ww.getAppSecretFunc != nil {
		corpid, secret, customizedApp = ww.getAppSecretFunc(corpId)
	}
	// 兼容代开发应用/自建应用/三方应用的token获取
	if !customizedApp {
		h := H{}
		h["auth_corpid"] = corpid
		h["permanent_code"] = secret
		_, err = ww.httpClient.R().SetQueryParam("suite_access_token", ww.getSuiteAccessToken()).
			SetBody(h).SetResult(&resp).Post("/cgi-bin/service/get_corp_token")
	} else {
		_, err = ww.httpClient.R().SetQueryParam("corpid", corpid).
			SetQueryParam("corpsecret", secret).
			SetResult(&resp).Get("/cgi-bin/gettoken")
	}
	if err != nil {
		logger.Sugar().Error(err)
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

func (ww *weWork) SetAppSecretFunc(f func(corpId uint) (corpid string, secret string, customizedApp bool)) {
	ww.getAppSecretFunc = f
}

func (ww *weWork) SetAgentIdFunc(f func(corpId uint) (agentId int)) {
	ww.getAgentIdFunc = f
}

func (ww *weWork) getCorpToken(corpId uint) (token string) {
	var err error
	token, err = ww.cache.conn.Get(context.TODO(), fmt.Sprintf("corpToken-%v", corpId)).Result()
	if err != nil {
		resp := ww.requestCorpToken(corpId)
		if resp.ErrCode != 0 {
			logger.Sugar().Error(resp.ErrorMsg)
			return ""
		} else {
			token = resp.AccessToken
			err = ww.cache.conn.Set(context.TODO(), fmt.Sprintf("corpToken-%v", corpId), token, time.Second*time.Duration(resp.ExpiresIn)).Err()
			if err != nil {
				logger.Sugar().Error(err.Error())
				return ""
			}
		}
	}
	return token
}

type GetUserInfo3rdResponse struct {
	BizResponse
	CorpId     string `json:"CorpId"`
	UserId     string `json:"UserId"`
	DeviceId   string `json:"DeviceId"`
	UserTicket string `json:"user_ticket"`
	ExpiresIn  int    `json:"expires_in"`
	OpenUserId string `json:"open_userid"`
}

// GetUserInfo3rd 获取访问用户身份
// https://open.work.weixin.qq.com/api/doc/90001/90143/91121
func (ww *weWork) GetUserInfo3rd(code string) (resp GetUserInfo3rdResponse) {
	_, err := ww.httpClient.R().
		SetQueryParam("suite_access_token", ww.getSuiteAccessToken()).
		SetQueryParam("code", code).
		SetResult(&resp).Get("/cgi-bin/service/auth/getuserinfo3rd")
	if err != nil {
		logger.Sugar().Error(err)
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

type GetUserInfoDetail3rdResponse struct {
	BizResponse
	CorpId string `json:"corpid"`
	UserId string `json:"userid"`
	Name   string `json:"name"`
	Gender string `json:"gender"`
	Avatar string `json:"avatar"`
	QrCode string `json:"qr_code"`
}

// GetUserInfoDetail3rd 获取访问用户敏感信息
// https://open.work.weixin.qq.com/api/doc/90001/90143/91122
func (ww *weWork) GetUserInfoDetail3rd(userTicket string) (resp GetUserInfoDetail3rdResponse) {
	h := H{}
	h["user_ticket"] = userTicket
	_, err := ww.httpClient.R().
		SetQueryParam("suite_access_token", ww.getSuiteAccessToken()).
		SetBody(h).
		SetResult(&resp).Post("/cgi-bin/service/auth/getuserdetail3rd")
	if err != nil {
		logger.Sugar().Error(err)
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

type GetUserInfoResponse struct {
	BizResponse
	UserId         string `json:"UserId,omitempty"`
	DeviceId       string `json:"DeviceId,omitempty"`
	UserTicket     string `json:"user_ticket,omitempty"`
	OpenId         string `json:"OpenId,omitempty"`
	ExternalUserId string `json:"external_userid,omitempty"`
}

// GetUserInfo
// https://developer.work.weixin.qq.com/document/path/91023
func (ww *weWork) GetUserInfo(corpId uint, code string) (resp GetUserInfoResponse) {
	_, err := ww.getRequest(corpId).SetQueryParam("code", code).
		SetResult(&resp).Get("/cgi-bin/auth/getuserinfo")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

type GetUserDetailResponse struct {
	BizResponse
	Userid  string `json:"userid"`
	Gender  string `json:"gender"`
	Avatar  string `json:"avatar"`
	QrCode  string `json:"qr_code"`
	Mobile  string `json:"mobile"`
	Email   string `json:"email"`
	BizMail string `json:"biz_mail"`
	Address string `json:"address"`
}

func (ww *weWork) GetUserDetail(corpId uint, userTicket string) (resp GetUserDetailResponse) {
	p := H{"user_ticket": userTicket}
	_, err := ww.getRequest(corpId).SetBody(p).
		SetResult(&resp).Post("/cgi-bin/auth/getuserdetail")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

type GetAppQrCodeRequest struct {
	SuiteID    string `json:"suite_id"`
	Appid      int    `json:"appid,omitempty"`
	State      string `json:"state,omitempty"`
	Style      int    `json:"style,omitempty" validate:"omitempty,oneof=0 1 2 3 4"`
	ResultType int    `json:"result_type" validate:"required,oneof=2"`
}

type GetAppQrCodeResponse struct {
	BizResponse
	QrCode string `json:"qrcode"`
}

// GetAppQrCode 获取应用二维码 仅支持二维码地址返回
// https://developer.work.weixin.qq.com/document/path/95430#36592
func (ww *weWork) GetAppQrCode(request GetAppQrCodeRequest) (resp GetAppQrCodeResponse) {
	if ok := validate.Validate(request); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
	_, err := ww.httpClient.R().SetQueryParam("suite_access_token", ww.GetSuiteToken()).
		SetBody(request).SetResult(&resp).
		Post("/cgi-bin/service/get_app_qrcode")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

type GetAdminListRequest struct {
	AuthCorpId string `json:"auth_corpid" validate:"required"`
	AgentId    uint   `json:"agentid" validate:"required"`
}

type GetAdminListResponse struct {
	BizResponse
	Admin []struct {
		Userid     string `json:"userid"`
		OpenUserid string `json:"open_userid"`
		AuthType   int    `json:"auth_type"`
	} `json:"admin"`
}

func (ww *weWork) GetAdminList(request GetAdminListRequest) (resp GetAdminListResponse) {
	if ok := validate.Validate(request); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
	_, err := ww.httpClient.R().SetQueryParam("suite_access_token", ww.GetSuiteToken()).
		SetBody(request).SetResult(&resp).
		Post("/cgi-bin/service/get_admin_list")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

// ExecuteCorpApi
// 　apiUrl 需要带有 /cgi-bin
// 　GET请求时data传入nil即可
// Deprecated:
func (ww *weWork) ExecuteCorpApi(corpId uint, apiUrl string, query url.Values, data H) (body []byte, err error) {
	query.Add("access_token", ww.getCorpToken(corpId))
	if os.Getenv("debug") != "" {
		query.Add("debug", "1")
	}
	if len(data) != 0 {
		return HttpPost(fmt.Sprintf("%s?%s", apiUrl, query.Encode()), data)
	} else {
		return HttpGet(fmt.Sprintf("%s?%s", apiUrl, query.Encode()))
	}
}
