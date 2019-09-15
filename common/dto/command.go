package dto

import (
	"github.com/zhaoche27/colago/common/extension"
)

type Commander interface {
	BizScenarioInfo() extension.BizScenario
	Desc() string
}

type Command struct {
	BizScenario extension.BizScenario `json:"biz_scenario"`
}

func (c Command) BizScenarioInfo() extension.BizScenario {
	if extension.ZeroBizScenario == c.BizScenario {
		return extension.DefaultBizScenario
	}
	return c.BizScenario
}

func (c Command) Desc() string {
	return ""
}
