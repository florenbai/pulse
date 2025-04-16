package alert_rule

import (
	h "darkhawk/pkg/http"
	"encoding/json"
	"fmt"
	"net/http"
)

type PrometheusRule struct {
	Annotations map[string]string `json:"annotations"`
	Duration    int64             `json:"duration"`
	Health      string            `json:"health"`
	Labels      map[string]string `json:"labels"`
	Name        string            `json:"name"`
	Query       string            `json:"query"`
}

type PrometheusGroup struct {
	Name  string           `json:"name"`
	File  string           `json:"file"`
	Rules []PrometheusRule `json:"rules"`
}

type PrometheusGroups struct {
	Groups []PrometheusGroup `json:"groups"`
}

type PrometheusResponse struct {
	Status string           `json:"status"`
	Data   PrometheusGroups `json:"data"`
}

func GetRemotePrometheusRule(addr string) (PrometheusGroups, error) {
	var data PrometheusResponse
	resp, err := h.Get(fmt.Sprintf("%s/api/v1/rules", addr), nil, nil)
	if err != nil {
		return PrometheusGroups{}, err
	}
	if resp.StatusCode() != http.StatusOK {
		return PrometheusGroups{}, fmt.Errorf("%v", resp.Error())
	}

	err = json.Unmarshal(resp.Body(), &data)
	if resp.StatusCode() != http.StatusOK {
		return PrometheusGroups{}, nil
	}
	return data.Data, nil
}
