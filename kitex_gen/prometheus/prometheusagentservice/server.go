// Code generated by Kitex v0.9.0. DO NOT EDIT.
package prometheusagentservice

import (
	server "github.com/cloudwego/kitex/server"
	prometheus "darkhawk/kitex_gen/prometheus"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler prometheus.PrometheusAgentService, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}

func RegisterService(svr server.Server, handler prometheus.PrometheusAgentService, opts ...server.RegisterOption) error {
	return svr.RegisterService(serviceInfo(), handler, opts...)
}
