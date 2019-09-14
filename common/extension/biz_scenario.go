package extension

import "strings"

const (
	DefaultBizId    = "defaultBizId"
	DefaultUseCase  = "defaultUseCase"
	DefaultScenario = "defaultScenario"
	DotSeparator    = "."
)

type BizScenario struct {
	BizId    string
	UseCase  string
	Scenario string
}

func NewDefaultBizScenario() *BizScenario {
	return &BizScenario{BizId: DefaultBizId, UseCase: DefaultUseCase, Scenario: DefaultScenario}
}

func NewBizScenarioByBizId(bizId string) *BizScenario {
	return NewBizScenario(bizId, DefaultUseCase, DefaultScenario)
}

func NewBizScenarioWithoutBizId(useCase, scenario string) *BizScenario {
	return NewBizScenario(DefaultBizId, useCase, scenario)
}

func NewBizScenario(bizId, useCase, scenario string) *BizScenario {
	return &BizScenario{BizId: bizId, UseCase: useCase, Scenario: scenario}
}

func (bs BizScenario) String() string {
	return strings.Join([]string{bs.BizId, bs.UseCase, bs.Scenario}, DotSeparator)
}
