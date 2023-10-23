package backEnd

import (
	// "fmt"
	"go.lwh.com/linweihao/calculatorOfAutomaticInvestmentPlan/exec/back/dataGet"
	"go.lwh.com/linweihao/calculatorOfAutomaticInvestmentPlan/utils/rfl"
	"go.lwh.com/linweihao/calculatorOfAutomaticInvestmentPlan/utils/time"
)

func ExecBackEnd(paramsIn rfl.Params) (paramsOut rfl.Params) {
	paramsOut = dataGet.GetData(paramsIn)
	time.ShowTimeAndMsg("OK")
	return paramsOut
}
