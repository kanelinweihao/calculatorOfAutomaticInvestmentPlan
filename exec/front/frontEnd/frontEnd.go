package frontEnd

import (
	// "fmt"
	"go.lwh.com/linweihao/calculatorOfAutomaticInvestmentPlan/utils/err"
	// "go.lwh.com/linweihao/calculatorOfAutomaticInvestmentPlan/utils/time"
	"go.lwh.com/linweihao/calculatorOfAutomaticInvestmentPlan/exec/front/htmlSet"
)

func ExecFrontEnd() {
	defer err.ThrowError()
	htmlSet.SetHtml()
	return
}
