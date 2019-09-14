package dto

import "github.com/zhaoche27/colago/common/extension"

type Commander interface {
	BizScenarioInfo() *extension.BizScenario
	Desc() string
}

type Command struct {
	BizScenario *extension.BizScenario
}

func (c *Command) BizScenarioInfo() *extension.BizScenario {
	return c.BizScenario
}

func (c *Command) Desc() string {
	return ""
}
