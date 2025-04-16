package agent

import (
	"agent/core"
	"agent/rpc"
	"context"
	"time"
)

func HeartBeat() {
	go func() {
		ticker := time.NewTicker(1 * time.Minute)
		for {
			select {
			case <-ticker.C:
				rpc.Ping(context.TODO())
			}
		}
	}()
}

func ConfigWatch() {
	go func() {
		ticker := time.NewTicker(1 * time.Minute)
		var configReloadTime time.Time
		for {
			select {
			case <-ticker.C:
				reloadTime := core.PrometheusLastConfigTime()
				if reloadTime.After(configReloadTime) {
					rpc.UpdateConfig(context.TODO())
				}
			}
		}
	}()
}
