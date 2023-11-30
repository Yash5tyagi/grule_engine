package rule_engine

import (
	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/builder"
	"github.com/hyperjumptech/grule-rule-engine/engine"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
)

var knowledgeLibrary = *ast.NewKnowledgeLibrary()

type RuleInput interface {
	DataKey() string
}

type RuleOutput interface {
	DataKey() string
}

type RuleConfig interface {
	RuleName() string
	RuleInput() RuleInput
	RuleOutput() RuleOutput
}

type RuleEngineSvc struct {
}

func NewRuleEngineSvc() *RuleEngineSvc {
	buildRuleEngine()
	return &RuleEngineSvc{}
}

func buildRuleEngine() {
	ruleBuilder := builder.NewRuleBuilder(&knowledgeLibrary)
	ruleFile := pkg.NewFileResource("rules.grl")
	err := ruleBuilder.BuildRuleFromResource("Rules", "0.0.1", ruleFile)
	if err != nil {
		panic(err)
	}
}

func (svc *RuleEngineSvc) Execute(ruleConf RuleConfig) error {
	knowledgeBase, _ := knowledgeLibrary.NewKnowledgeBaseInstance("Rules", "0.0.1")
	dataCtx := ast.NewDataContext()

	err := dataCtx.Add(ruleConf.RuleInput().DataKey(), ruleConf.RuleInput())
	if err != nil {
		return err
	}

	err = dataCtx.Add(ruleConf.RuleOutput().DataKey(), ruleConf.RuleOutput())
	if err != nil {
		return err
	}

	ruleEngine := engine.NewGruleEngine()
	err = ruleEngine.Execute(dataCtx, knowledgeBase)
	if err != nil {
		return err
	}
	return nil
}
