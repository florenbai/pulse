// Code generated by Kitex v0.5.1. DO NOT EDIT.

package userservice

import (
	server "github.com/cloudwego/kitex/server"
	sso "minos/kitex_gen/sso"
)

// NewInvoker creates a server.Invoker with the given handler and options.
func NewInvoker(handler sso.UserService, opts ...server.Option) server.Invoker {
	var options []server.Option

	options = append(options, opts...)

	s := server.NewInvoker(options...)
	if err := s.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	if err := s.Init(); err != nil {
		panic(err)
	}
	return s
}
