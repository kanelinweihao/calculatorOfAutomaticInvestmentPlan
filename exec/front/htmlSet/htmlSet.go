package htmlSet

import (
	// "fmt"
	"go.lwh.com/linweihao/calculatorOfAutomaticInvestmentPlan/exec/back/backEnd"
	"go.lwh.com/linweihao/calculatorOfAutomaticInvestmentPlan/utils/rfl"
	"go.lwh.com/linweihao/calculatorOfAutomaticInvestmentPlan/utils/tmpl"
)

var paramsRoute = rfl.BoxParams[string]{
	"/": {
		"pathTmpl": {
			"pathTmpl": "./public/tmpl/calculator.tmpl",
		},
		"paramsOutDefault": {
			"ModTitle":    "未知标题",
			"Version":     "v1.0.0",
			"SumTimes":    "2",
			"AmountBegin": "10000",
			"AmountAdd":   "1000",
			"RateReal":    "0.01",
			"RateShow":    "0",
			"MsgOut":      "",
		},
		"paramsInDefault": {
			"SumTimes":    "1",
			"AmountBegin": "0",
			"AmountAdd":   "0",
			"RateReal":    "0",
		},
	},
	"/index": {
		"pathTmpl": {
			"pathTmpl": "./public/tmpl/calculator.tmpl",
		},
		"paramsOutDefault": {
			"ModTitle":    "未知标题",
			"Version":     "v1.0.0",
			"SumTimes":    "2",
			"AmountBegin": "10000",
			"AmountAdd":   "1000",
			"RateReal":    "0.01",
			"RateShow":    "0",
			"MsgOut":      "",
		},
		"paramsInDefault": {
			"SumTimes":    "1",
			"AmountBegin": "0",
			"AmountAdd":   "0",
			"RateReal":    "0",
		},
	},
}

type EntityBackEnd struct{}

func (self *EntityBackEnd) ExecBackEnd(paramsIn rfl.Params) (paramsFromBackend rfl.Params) {
	paramsFromBackend = backEnd.ExecBackEnd(paramsIn)
	return paramsFromBackend
}

func GetParamsRoute() (paramsRoute rfl.BoxParams[string]) {
	return paramsRoute
}

func SetHtml() {
	entityBackEnd := &EntityBackEnd{}
	tmpl.SetTmplAndServer(paramsRoute, entityBackEnd)
	return
}
