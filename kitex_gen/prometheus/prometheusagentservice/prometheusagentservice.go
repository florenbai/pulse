// Code generated by Kitex v0.9.0. DO NOT EDIT.

package prometheusagentservice

import (
	"context"
	"errors"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	proto "google.golang.org/protobuf/proto"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	prometheus "darkhawk/kitex_gen/prometheus"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"CreateAlertGroup": kitex.NewMethodInfo(
		createAlertGroupHandler,
		newCreateAlertGroupArgs,
		newCreateAlertGroupResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"DeleteAlertGroup": kitex.NewMethodInfo(
		deleteAlertGroupHandler,
		newDeleteAlertGroupArgs,
		newDeleteAlertGroupResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"CreateAlertRules": kitex.NewMethodInfo(
		createAlertRulesHandler,
		newCreateAlertRulesArgs,
		newCreateAlertRulesResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"UpdateAlertRule": kitex.NewMethodInfo(
		updateAlertRuleHandler,
		newUpdateAlertRuleArgs,
		newUpdateAlertRuleResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"DeleteAlertRule": kitex.NewMethodInfo(
		deleteAlertRuleHandler,
		newDeleteAlertRuleArgs,
		newDeleteAlertRuleResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"Reload": kitex.NewMethodInfo(
		reloadHandler,
		newReloadArgs,
		newReloadResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
}

var (
	prometheusAgentServiceServiceInfo                = NewServiceInfo()
	prometheusAgentServiceServiceInfoForClient       = NewServiceInfoForClient()
	prometheusAgentServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return prometheusAgentServiceServiceInfo
}

// for client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return prometheusAgentServiceServiceInfoForStreamClient
}

// for stream client
func serviceInfoForClient() *kitex.ServiceInfo {
	return prometheusAgentServiceServiceInfoForClient
}

// NewServiceInfo creates a new ServiceInfo containing all methods
func NewServiceInfo() *kitex.ServiceInfo {
	return newServiceInfo(false, true, true)
}

// NewServiceInfo creates a new ServiceInfo containing non-streaming methods
func NewServiceInfoForClient() *kitex.ServiceInfo {
	return newServiceInfo(false, false, true)
}
func NewServiceInfoForStreamClient() *kitex.ServiceInfo {
	return newServiceInfo(true, true, false)
}

func newServiceInfo(hasStreaming bool, keepStreamingMethods bool, keepNonStreamingMethods bool) *kitex.ServiceInfo {
	serviceName := "PrometheusAgentService"
	handlerType := (*prometheus.PrometheusAgentService)(nil)
	methods := map[string]kitex.MethodInfo{}
	for name, m := range serviceMethods {
		if m.IsStreaming() && !keepStreamingMethods {
			continue
		}
		if !m.IsStreaming() && !keepNonStreamingMethods {
			continue
		}
		methods[name] = m
	}
	extra := map[string]interface{}{
		"PackageName": "prometheus",
	}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.9.0",
		Extra:           extra,
	}
	return svcInfo
}

func createAlertGroupHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(prometheus.AlertGroupRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(prometheus.PrometheusAgentService).CreateAlertGroup(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *CreateAlertGroupArgs:
		success, err := handler.(prometheus.PrometheusAgentService).CreateAlertGroup(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*CreateAlertGroupResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newCreateAlertGroupArgs() interface{} {
	return &CreateAlertGroupArgs{}
}

func newCreateAlertGroupResult() interface{} {
	return &CreateAlertGroupResult{}
}

type CreateAlertGroupArgs struct {
	Req *prometheus.AlertGroupRequest
}

func (p *CreateAlertGroupArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *CreateAlertGroupArgs) Unmarshal(in []byte) error {
	msg := new(prometheus.AlertGroupRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var CreateAlertGroupArgs_Req_DEFAULT *prometheus.AlertGroupRequest

func (p *CreateAlertGroupArgs) GetReq() *prometheus.AlertGroupRequest {
	if !p.IsSetReq() {
		return CreateAlertGroupArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *CreateAlertGroupArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *CreateAlertGroupArgs) GetFirstArgument() interface{} {
	return p.Req
}

type CreateAlertGroupResult struct {
	Success *emptypb.Empty
}

var CreateAlertGroupResult_Success_DEFAULT *emptypb.Empty

func (p *CreateAlertGroupResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *CreateAlertGroupResult) Unmarshal(in []byte) error {
	msg := new(emptypb.Empty)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *CreateAlertGroupResult) GetSuccess() *emptypb.Empty {
	if !p.IsSetSuccess() {
		return CreateAlertGroupResult_Success_DEFAULT
	}
	return p.Success
}

func (p *CreateAlertGroupResult) SetSuccess(x interface{}) {
	p.Success = x.(*emptypb.Empty)
}

func (p *CreateAlertGroupResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *CreateAlertGroupResult) GetResult() interface{} {
	return p.Success
}

func deleteAlertGroupHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(prometheus.AlertGroupRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(prometheus.PrometheusAgentService).DeleteAlertGroup(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *DeleteAlertGroupArgs:
		success, err := handler.(prometheus.PrometheusAgentService).DeleteAlertGroup(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*DeleteAlertGroupResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newDeleteAlertGroupArgs() interface{} {
	return &DeleteAlertGroupArgs{}
}

func newDeleteAlertGroupResult() interface{} {
	return &DeleteAlertGroupResult{}
}

type DeleteAlertGroupArgs struct {
	Req *prometheus.AlertGroupRequest
}

func (p *DeleteAlertGroupArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *DeleteAlertGroupArgs) Unmarshal(in []byte) error {
	msg := new(prometheus.AlertGroupRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var DeleteAlertGroupArgs_Req_DEFAULT *prometheus.AlertGroupRequest

func (p *DeleteAlertGroupArgs) GetReq() *prometheus.AlertGroupRequest {
	if !p.IsSetReq() {
		return DeleteAlertGroupArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *DeleteAlertGroupArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *DeleteAlertGroupArgs) GetFirstArgument() interface{} {
	return p.Req
}

type DeleteAlertGroupResult struct {
	Success *emptypb.Empty
}

var DeleteAlertGroupResult_Success_DEFAULT *emptypb.Empty

func (p *DeleteAlertGroupResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *DeleteAlertGroupResult) Unmarshal(in []byte) error {
	msg := new(emptypb.Empty)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *DeleteAlertGroupResult) GetSuccess() *emptypb.Empty {
	if !p.IsSetSuccess() {
		return DeleteAlertGroupResult_Success_DEFAULT
	}
	return p.Success
}

func (p *DeleteAlertGroupResult) SetSuccess(x interface{}) {
	p.Success = x.(*emptypb.Empty)
}

func (p *DeleteAlertGroupResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *DeleteAlertGroupResult) GetResult() interface{} {
	return p.Success
}

func createAlertRulesHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(prometheus.AlertRules)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(prometheus.PrometheusAgentService).CreateAlertRules(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *CreateAlertRulesArgs:
		success, err := handler.(prometheus.PrometheusAgentService).CreateAlertRules(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*CreateAlertRulesResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newCreateAlertRulesArgs() interface{} {
	return &CreateAlertRulesArgs{}
}

func newCreateAlertRulesResult() interface{} {
	return &CreateAlertRulesResult{}
}

type CreateAlertRulesArgs struct {
	Req *prometheus.AlertRules
}

func (p *CreateAlertRulesArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *CreateAlertRulesArgs) Unmarshal(in []byte) error {
	msg := new(prometheus.AlertRules)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var CreateAlertRulesArgs_Req_DEFAULT *prometheus.AlertRules

func (p *CreateAlertRulesArgs) GetReq() *prometheus.AlertRules {
	if !p.IsSetReq() {
		return CreateAlertRulesArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *CreateAlertRulesArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *CreateAlertRulesArgs) GetFirstArgument() interface{} {
	return p.Req
}

type CreateAlertRulesResult struct {
	Success *emptypb.Empty
}

var CreateAlertRulesResult_Success_DEFAULT *emptypb.Empty

func (p *CreateAlertRulesResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *CreateAlertRulesResult) Unmarshal(in []byte) error {
	msg := new(emptypb.Empty)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *CreateAlertRulesResult) GetSuccess() *emptypb.Empty {
	if !p.IsSetSuccess() {
		return CreateAlertRulesResult_Success_DEFAULT
	}
	return p.Success
}

func (p *CreateAlertRulesResult) SetSuccess(x interface{}) {
	p.Success = x.(*emptypb.Empty)
}

func (p *CreateAlertRulesResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *CreateAlertRulesResult) GetResult() interface{} {
	return p.Success
}

func updateAlertRuleHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(prometheus.UpdateAlertRuleRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(prometheus.PrometheusAgentService).UpdateAlertRule(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *UpdateAlertRuleArgs:
		success, err := handler.(prometheus.PrometheusAgentService).UpdateAlertRule(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*UpdateAlertRuleResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newUpdateAlertRuleArgs() interface{} {
	return &UpdateAlertRuleArgs{}
}

func newUpdateAlertRuleResult() interface{} {
	return &UpdateAlertRuleResult{}
}

type UpdateAlertRuleArgs struct {
	Req *prometheus.UpdateAlertRuleRequest
}

func (p *UpdateAlertRuleArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *UpdateAlertRuleArgs) Unmarshal(in []byte) error {
	msg := new(prometheus.UpdateAlertRuleRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var UpdateAlertRuleArgs_Req_DEFAULT *prometheus.UpdateAlertRuleRequest

func (p *UpdateAlertRuleArgs) GetReq() *prometheus.UpdateAlertRuleRequest {
	if !p.IsSetReq() {
		return UpdateAlertRuleArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UpdateAlertRuleArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *UpdateAlertRuleArgs) GetFirstArgument() interface{} {
	return p.Req
}

type UpdateAlertRuleResult struct {
	Success *emptypb.Empty
}

var UpdateAlertRuleResult_Success_DEFAULT *emptypb.Empty

func (p *UpdateAlertRuleResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *UpdateAlertRuleResult) Unmarshal(in []byte) error {
	msg := new(emptypb.Empty)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UpdateAlertRuleResult) GetSuccess() *emptypb.Empty {
	if !p.IsSetSuccess() {
		return UpdateAlertRuleResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UpdateAlertRuleResult) SetSuccess(x interface{}) {
	p.Success = x.(*emptypb.Empty)
}

func (p *UpdateAlertRuleResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *UpdateAlertRuleResult) GetResult() interface{} {
	return p.Success
}

func deleteAlertRuleHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(prometheus.AlertRules)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(prometheus.PrometheusAgentService).DeleteAlertRule(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *DeleteAlertRuleArgs:
		success, err := handler.(prometheus.PrometheusAgentService).DeleteAlertRule(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*DeleteAlertRuleResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newDeleteAlertRuleArgs() interface{} {
	return &DeleteAlertRuleArgs{}
}

func newDeleteAlertRuleResult() interface{} {
	return &DeleteAlertRuleResult{}
}

type DeleteAlertRuleArgs struct {
	Req *prometheus.AlertRules
}

func (p *DeleteAlertRuleArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *DeleteAlertRuleArgs) Unmarshal(in []byte) error {
	msg := new(prometheus.AlertRules)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var DeleteAlertRuleArgs_Req_DEFAULT *prometheus.AlertRules

func (p *DeleteAlertRuleArgs) GetReq() *prometheus.AlertRules {
	if !p.IsSetReq() {
		return DeleteAlertRuleArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *DeleteAlertRuleArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *DeleteAlertRuleArgs) GetFirstArgument() interface{} {
	return p.Req
}

type DeleteAlertRuleResult struct {
	Success *emptypb.Empty
}

var DeleteAlertRuleResult_Success_DEFAULT *emptypb.Empty

func (p *DeleteAlertRuleResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *DeleteAlertRuleResult) Unmarshal(in []byte) error {
	msg := new(emptypb.Empty)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *DeleteAlertRuleResult) GetSuccess() *emptypb.Empty {
	if !p.IsSetSuccess() {
		return DeleteAlertRuleResult_Success_DEFAULT
	}
	return p.Success
}

func (p *DeleteAlertRuleResult) SetSuccess(x interface{}) {
	p.Success = x.(*emptypb.Empty)
}

func (p *DeleteAlertRuleResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *DeleteAlertRuleResult) GetResult() interface{} {
	return p.Success
}

func reloadHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(emptypb.Empty)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(prometheus.PrometheusAgentService).Reload(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *ReloadArgs:
		success, err := handler.(prometheus.PrometheusAgentService).Reload(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*ReloadResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newReloadArgs() interface{} {
	return &ReloadArgs{}
}

func newReloadResult() interface{} {
	return &ReloadResult{}
}

type ReloadArgs struct {
	Req *emptypb.Empty
}

func (p *ReloadArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *ReloadArgs) Unmarshal(in []byte) error {
	msg := new(emptypb.Empty)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var ReloadArgs_Req_DEFAULT *emptypb.Empty

func (p *ReloadArgs) GetReq() *emptypb.Empty {
	if !p.IsSetReq() {
		return ReloadArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ReloadArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *ReloadArgs) GetFirstArgument() interface{} {
	return p.Req
}

type ReloadResult struct {
	Success *emptypb.Empty
}

var ReloadResult_Success_DEFAULT *emptypb.Empty

func (p *ReloadResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *ReloadResult) Unmarshal(in []byte) error {
	msg := new(emptypb.Empty)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ReloadResult) GetSuccess() *emptypb.Empty {
	if !p.IsSetSuccess() {
		return ReloadResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ReloadResult) SetSuccess(x interface{}) {
	p.Success = x.(*emptypb.Empty)
}

func (p *ReloadResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ReloadResult) GetResult() interface{} {
	return p.Success
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) CreateAlertGroup(ctx context.Context, Req *prometheus.AlertGroupRequest) (r *emptypb.Empty, err error) {
	var _args CreateAlertGroupArgs
	_args.Req = Req
	var _result CreateAlertGroupResult
	if err = p.c.Call(ctx, "CreateAlertGroup", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) DeleteAlertGroup(ctx context.Context, Req *prometheus.AlertGroupRequest) (r *emptypb.Empty, err error) {
	var _args DeleteAlertGroupArgs
	_args.Req = Req
	var _result DeleteAlertGroupResult
	if err = p.c.Call(ctx, "DeleteAlertGroup", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CreateAlertRules(ctx context.Context, Req *prometheus.AlertRules) (r *emptypb.Empty, err error) {
	var _args CreateAlertRulesArgs
	_args.Req = Req
	var _result CreateAlertRulesResult
	if err = p.c.Call(ctx, "CreateAlertRules", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UpdateAlertRule(ctx context.Context, Req *prometheus.UpdateAlertRuleRequest) (r *emptypb.Empty, err error) {
	var _args UpdateAlertRuleArgs
	_args.Req = Req
	var _result UpdateAlertRuleResult
	if err = p.c.Call(ctx, "UpdateAlertRule", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) DeleteAlertRule(ctx context.Context, Req *prometheus.AlertRules) (r *emptypb.Empty, err error) {
	var _args DeleteAlertRuleArgs
	_args.Req = Req
	var _result DeleteAlertRuleResult
	if err = p.c.Call(ctx, "DeleteAlertRule", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Reload(ctx context.Context, Req *emptypb.Empty) (r *emptypb.Empty, err error) {
	var _args ReloadArgs
	_args.Req = Req
	var _result ReloadResult
	if err = p.c.Call(ctx, "Reload", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
