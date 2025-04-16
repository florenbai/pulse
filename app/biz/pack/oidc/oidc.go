package oidc

import (
	"context"
	"darkhawk/conf"
	"encoding/base64"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/coreos/go-oidc/v3/oidc"
	"golang.org/x/oauth2"
)

type Oidc struct {
	Provider     *oidc.Provider
	Issuer       string
	ClientId     string
	ClientSecret string
	CallbackPath string
	Scope        string
	AuthURL      string
	TokenURL     string
}

const (
	CallbackPath = "/api/v1/oidc/callback"
)

func NewOIDC(ctx context.Context, issuer, clientId, clientSecret, scope, callbackPath string) *Oidc {
	ctx = oidc.InsecureIssuerURLContext(ctx, issuer)
	pro, err := oidc.NewProvider(ctx, issuer)
	if err != nil {
		panic(err)
	}
	return &Oidc{Provider: pro, Issuer: issuer, ClientSecret: clientSecret,
		ClientId: clientId, CallbackPath: callbackPath, Scope: scope}
}

func InitProvider(ctx context.Context, issuer string) (*oidc.Provider, error) {
	var err error
	ctx = oidc.InsecureIssuerURLContext(ctx, issuer)
	providerObj, err := oidc.NewProvider(ctx, issuer)
	return providerObj, err
}

// GetCallbackUri 获取本机url
func GetCallbackUri(c *app.RequestContext, withIdp string) string {
	var xSchema = string(c.GetHeader(conf.SchemeHeaderValue))
	var schema = string(c.Request.Scheme())
	if len(xSchema) > 0 {
		schema = xSchema
	}
	hlog.Debugf("X-Scheme: %s", xSchema)
	path := CallbackPath
	if len(withIdp) > 0 {
		path = fmt.Sprintf("%s/%s", path, withIdp)
	}
	host := fmt.Sprintf("%s://%s%s", schema, string(c.Request.Host()), path)
	return host
}

func (o *Oidc) InitProvider(ctx context.Context) (*oidc.Provider, error) {
	var err error
	o.Provider, err = InitProvider(ctx, o.Issuer)
	return o.Provider, err
}

// OidcCallbackWithCode sso code类型 客户端回调接口方法
func (o *Oidc) OidcCallbackWithCode(ctx context.Context, redirectURL, code string) (*oidc.UserInfo, error) {
	var err error
	o.Provider, err = o.InitProvider(ctx)
	if err != nil {
		return nil, err
	}

	hlog.CtxDebugf(ctx, "get oidc provider is success")
	// 1. 配置OAuth 2.0客户端
	oauth2Config := &oauth2.Config{
		ClientID:     o.ClientId,
		ClientSecret: o.ClientSecret,
		RedirectURL:  o.CallbackPath,
		Endpoint:     o.Provider.Endpoint(),
		Scopes:       []string{o.Scope},
	}
	if o.TokenURL != "" {
		oauth2Config.Endpoint.TokenURL = o.TokenURL
	}
	if o.AuthURL != "" {
		oauth2Config.Endpoint.AuthURL = o.AuthURL
	}
	oauth2Token, err := oauth2Config.Exchange(ctx, code)
	if err != nil {
		hlog.CtxErrorf(ctx, "get oauth2 access_token id_token err, %s", err)
		return nil, err
	}
	hlog.CtxDebugf(ctx, "oidc client auth code is success")
	// Extract the ID Token from OAuth2 token.
	rawIDToken, ok := oauth2Token.Extra("id_token").(string)
	if !ok {
		hlog.CtxErrorf(ctx, "handle missing id token")
		return nil, fmt.Errorf("handle missing id token")
	}
	var verifier = o.Provider.Verifier(&oidc.Config{ClientID: oauth2Config.ClientID})
	// Parse and verify ID Token payload.
	idToken, err := verifier.Verify(ctx, rawIDToken)
	if err != nil {
		fmt.Println("idToken:", idToken)
		hlog.CtxErrorf(ctx, "verifier id_token err, %s", err)
		return nil, err
	}
	userInfo, err := o.Provider.UserInfo(ctx, oauth2.StaticTokenSource(oauth2Token))
	if err != nil {
		hlog.CtxErrorf(ctx, "%s", err.Error())
		data, er := base64.StdEncoding.DecodeString(err.Error())
		if er != nil {
			return nil, err
		}
		return nil, fmt.Errorf(string(data))
	}
	hlog.CtxDebugf(ctx, "oidc client get userinfo is success")
	return userInfo, nil
}
