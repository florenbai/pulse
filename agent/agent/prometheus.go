package agent

import (
	"agent/core"
	"agent/kitex_gen/prometheus"
	"agent/utils"
	"context"
	"errors"
	"fmt"
	"google.golang.org/protobuf/types/known/emptypb"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

type PrometheusServiceImpl struct{}

type RuleGroupYaml struct {
	Groups []RuleGroup `yaml:"groups"`
}

type RuleGroup struct {
	Name  string `yaml:"name"`
	Rules []Rule `yaml:"rules"`
}

type Rule struct {
	Alert       string            `yaml:"alert,omitempty"`
	Expr        string            `yaml:"expr,omitempty"`
	For         string            `yaml:"for,omitempty"`
	Labels      map[string]string `yaml:"labels,omitempty""`
	Annotations map[string]string `yaml:"annotations,omitempty""`
}

func (s *PrometheusServiceImpl) CreateAlertGroup(ctx context.Context, req *prometheus.AlertGroupRequest) (res *emptypb.Empty, err error) {
	ext := filepath.Ext(req.GetFileName())
	if ext != ".rules" {
		return nil, errors.New("prometheus规则文件必须使用后缀 .rules")
	}
	_, err = os.Stat(req.GetFileName())
	if err != nil && !os.IsNotExist(err) {
		return nil, err
	}
	if os.IsNotExist(err) {
		file, err := os.Create(req.GetFileName())
		if err != nil {
			return nil, err
		}
		defer file.Close()
	}

	return nil, nil
}

func (s *PrometheusServiceImpl) DeleteAlertGroup(ctx context.Context, req *prometheus.AlertGroupRequest) (res *emptypb.Empty, err error) {
	_, err = os.Stat(req.GetFileName())
	if err != nil {
		return nil, err
	}
	err = utils.CopyFile(req.GetFileName(), fmt.Sprintf("%s.bak", req.GetFileName()), false)
	if err != nil {
		return nil, err
	}
	os.Remove(req.GetFileName())
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (s *PrometheusServiceImpl) CreateAlertRules(ctx context.Context, req *prometheus.AlertRules) (res *emptypb.Empty, err error) {
	var groupYaml RuleGroupYaml
	data, err := os.ReadFile(req.GetFilename())
	if err != nil {
		return nil, err
	}
	err = utils.CopyFile(req.GetFilename(), fmt.Sprintf("%s.bak", req.GetFilename()), false)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(data, &groupYaml)
	if err != nil {
		return nil, err
	}
	var groupNames []string
	for _, v := range groupYaml.Groups {
		groupNames = append(groupNames, v.Name)
	}
	var rules []Rule
	for _, v := range req.Rules {
		rule := Rule{
			Alert:       v.Alert,
			Expr:        v.Expr,
			For:         v.For,
			Labels:      v.Labels,
			Annotations: v.Annotations,
		}
		rules = append(rules, rule)
	}

	if utils.InOfT(req.GetGroupName(), groupNames) {
		for k, v := range groupYaml.Groups {
			if v.Name == req.GetGroupName() {
				groupYaml.Groups[k].Rules = append(groupYaml.Groups[k].Rules, rules...)
			}
		}
	} else {
		newGroup := RuleGroup{
			Name:  req.GetGroupName(),
			Rules: rules,
		}
		groupYaml.Groups = append(groupYaml.Groups, newGroup)
	}
	yamlStr, err := yaml.Marshal(groupYaml)
	if err != nil {
		return nil, err
	}
	file, err := os.Create(req.GetFilename())
	if err != nil {
		return nil, err
	}
	defer file.Close()

	_, err = file.WriteString(string(yamlStr))
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (s *PrometheusServiceImpl) UpdateAlertRule(ctx context.Context, req *prometheus.UpdateAlertRuleRequest) (res *emptypb.Empty, err error) {
	var groupYaml RuleGroupYaml
	ext := filepath.Ext(req.GetFilename())
	if ext != ".rules" {
		return nil, errors.New("prometheus规则文件必须使用后缀 .rules")
	}
	data, err := os.ReadFile(req.GetFilename())
	if err != nil {
		return nil, err
	}
	err = utils.CopyFile(req.GetFilename(), fmt.Sprintf("%s.bak", req.GetFilename()), false)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(data, &groupYaml)
	if err != nil {
		return nil, err
	}
	for gk, group := range groupYaml.Groups {
		if group.Name == req.GetGroupName() {
			for k, v := range group.Rules {
				if v.Alert == req.GetUpdateAlertName() {
					groupYaml.Groups[gk].Rules[k] = Rule{
						Alert:       req.GetAlertRule().GetAlert(),
						Expr:        req.GetAlertRule().GetExpr(),
						For:         req.GetAlertRule().GetFor(),
						Labels:      req.GetAlertRule().GetLabels(),
						Annotations: req.GetAlertRule().GetAnnotations(),
					}
				}
			}
		}
	}
	file, err := os.Create(req.GetFilename())
	if err != nil {
		return nil, err
	}
	defer file.Close()
	yamlStr, err := yaml.Marshal(groupYaml)
	if err != nil {
		return nil, err
	}
	_, err = file.WriteString(string(yamlStr))
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (s *PrometheusServiceImpl) DeleteAlertRule(ctx context.Context, req *prometheus.AlertRules) (res *emptypb.Empty, err error) {
	var groupYaml RuleGroupYaml
	data, err := os.ReadFile(req.GetFilename())
	if err != nil {
		return nil, err
	}
	err = utils.CopyFile(req.GetFilename(), fmt.Sprintf("%s.bak", req.GetFilename()), false)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(data, &groupYaml)
	if err != nil {
		return nil, err
	}
	var deleteName []string
	for _, v := range req.Rules {
		deleteName = append(deleteName, v.GetAlert())
	}
	for i, g := range groupYaml.Groups {
		if req.GetGroupName() == g.Name {
			newRules := g.Rules
			for k, v := range g.Rules {
				if utils.InOfT(v.Alert, deleteName) {
					newRules = removeElement(newRules, k)
				}
			}
			groupYaml.Groups[i].Rules = newRules
		}
	}

	yamlStr, err := yaml.Marshal(groupYaml)
	if err != nil {
		return nil, err
	}
	file, err := os.Create(req.GetFilename())
	if err != nil {
		return nil, err
	}
	defer file.Close()

	_, err = file.WriteString(string(yamlStr))
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (s *PrometheusServiceImpl) Reload(ctx context.Context, req *emptypb.Empty) (res *emptypb.Empty, err error) {
	err = core.ReloadPrometheus()
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func removeElement(slice []Rule, index int) []Rule {
	// 检查索引是否越界
	if index < 0 || index >= len(slice) {
		return nil
	}

	newSlice := make([]Rule, len(slice)-1)

	copy(newSlice, slice[:index])

	copy(newSlice[index:], slice[index+1:])

	return newSlice
}
