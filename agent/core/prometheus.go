package core

import (
	"agent/cliconfig"
	"agent/utils"
	"fmt"
	"github.com/tidwall/gjson"
	"time"
)

func PrometheusLastConfigTime() time.Time {
	resp, _ := utils.Get(fmt.Sprintf("%sapi/v1/status/runtimeinfo", cliconfig.GetConf().Prometheus.Endpoint), nil, nil)
	if resp.IsSuccess() {
		return gjson.Parse(resp.String()).Get("data.lastConfigTime").Time()
	}
	return time.Time{}
}

func ReloadPrometheus() error {
	_, err := utils.Post(fmt.Sprintf("%s-/reload", cliconfig.GetConf().Prometheus.Endpoint), nil, nil)
	if err != nil {
		return err
	}
	return nil
}
