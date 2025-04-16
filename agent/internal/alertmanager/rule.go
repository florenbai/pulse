package alertmanager

import (
	"gopkg.in/yaml.v3"
	"os"
)

type Rule struct {
	Alert       string            `yaml:"alert"`
	Expr        string            `yaml:"expr"`
	For         string            `yaml:"for"`
	Labels      map[string]string `yaml:"labels"`
	Annotations map[string]string `yaml:"annotations"`
}

type RuleGroup struct {
	Name  string `yaml:"name"`
	Rules []Rule `yaml:"rules"`
}

type RuleFile struct {
	Groups []RuleGroup `yaml:"groups"`
}

func ReadRuleFile(filePath string) (*RuleFile, error) {
	var ruleFile RuleFile
	// 读取文件内容
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	// 解析YAML内容到结构体
	err = yaml.Unmarshal(data, &ruleFile)
	if err != nil {
		return nil, err
	}

	return &ruleFile, nil
}
