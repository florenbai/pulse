// Code generated by Kitex v0.9.0. DO NOT EDIT.
package owlagentservice

import (
	server "github.com/cloudwego/kitex/server"
	owl_agent "agent/kitex_gen/owl_agent"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler owl_agent.OwlAgentService, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}

func RegisterService(svr server.Server, handler owl_agent.OwlAgentService, opts ...server.RegisterOption) error {
	return svr.RegisterService(serviceInfo(), handler, opts...)
}
