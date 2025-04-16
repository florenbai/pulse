// Code generated by Kitex v0.9.0. DO NOT EDIT.

package agentservice

import (
	"context"
	"errors"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	proto "google.golang.org/protobuf/proto"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	agent "darkhawk/kitex_gen/agent"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"Ping": kitex.NewMethodInfo(
		pingHandler,
		newPingArgs,
		newPingResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"UpdateConfig": kitex.NewMethodInfo(
		updateConfigHandler,
		newUpdateConfigArgs,
		newUpdateConfigResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
}

var (
	agentServiceServiceInfo                = NewServiceInfo()
	agentServiceServiceInfoForClient       = NewServiceInfoForClient()
	agentServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return agentServiceServiceInfo
}

// for client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return agentServiceServiceInfoForStreamClient
}

// for stream client
func serviceInfoForClient() *kitex.ServiceInfo {
	return agentServiceServiceInfoForClient
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
	serviceName := "AgentService"
	handlerType := (*agent.AgentService)(nil)
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
		"PackageName": "agent",
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

func pingHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(agent.PingRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(agent.AgentService).Ping(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *PingArgs:
		success, err := handler.(agent.AgentService).Ping(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*PingResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newPingArgs() interface{} {
	return &PingArgs{}
}

func newPingResult() interface{} {
	return &PingResult{}
}

type PingArgs struct {
	Req *agent.PingRequest
}

func (p *PingArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *PingArgs) Unmarshal(in []byte) error {
	msg := new(agent.PingRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var PingArgs_Req_DEFAULT *agent.PingRequest

func (p *PingArgs) GetReq() *agent.PingRequest {
	if !p.IsSetReq() {
		return PingArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *PingArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *PingArgs) GetFirstArgument() interface{} {
	return p.Req
}

type PingResult struct {
	Success *emptypb.Empty
}

var PingResult_Success_DEFAULT *emptypb.Empty

func (p *PingResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *PingResult) Unmarshal(in []byte) error {
	msg := new(emptypb.Empty)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *PingResult) GetSuccess() *emptypb.Empty {
	if !p.IsSetSuccess() {
		return PingResult_Success_DEFAULT
	}
	return p.Success
}

func (p *PingResult) SetSuccess(x interface{}) {
	p.Success = x.(*emptypb.Empty)
}

func (p *PingResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *PingResult) GetResult() interface{} {
	return p.Success
}

func updateConfigHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(agent.PingRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(agent.AgentService).UpdateConfig(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *UpdateConfigArgs:
		success, err := handler.(agent.AgentService).UpdateConfig(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*UpdateConfigResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newUpdateConfigArgs() interface{} {
	return &UpdateConfigArgs{}
}

func newUpdateConfigResult() interface{} {
	return &UpdateConfigResult{}
}

type UpdateConfigArgs struct {
	Req *agent.PingRequest
}

func (p *UpdateConfigArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *UpdateConfigArgs) Unmarshal(in []byte) error {
	msg := new(agent.PingRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var UpdateConfigArgs_Req_DEFAULT *agent.PingRequest

func (p *UpdateConfigArgs) GetReq() *agent.PingRequest {
	if !p.IsSetReq() {
		return UpdateConfigArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UpdateConfigArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *UpdateConfigArgs) GetFirstArgument() interface{} {
	return p.Req
}

type UpdateConfigResult struct {
	Success *emptypb.Empty
}

var UpdateConfigResult_Success_DEFAULT *emptypb.Empty

func (p *UpdateConfigResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *UpdateConfigResult) Unmarshal(in []byte) error {
	msg := new(emptypb.Empty)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UpdateConfigResult) GetSuccess() *emptypb.Empty {
	if !p.IsSetSuccess() {
		return UpdateConfigResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UpdateConfigResult) SetSuccess(x interface{}) {
	p.Success = x.(*emptypb.Empty)
}

func (p *UpdateConfigResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *UpdateConfigResult) GetResult() interface{} {
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

func (p *kClient) Ping(ctx context.Context, Req *agent.PingRequest) (r *emptypb.Empty, err error) {
	var _args PingArgs
	_args.Req = Req
	var _result PingResult
	if err = p.c.Call(ctx, "Ping", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UpdateConfig(ctx context.Context, Req *agent.PingRequest) (r *emptypb.Empty, err error) {
	var _args UpdateConfigArgs
	_args.Req = Req
	var _result UpdateConfigResult
	if err = p.c.Call(ctx, "UpdateConfig", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
