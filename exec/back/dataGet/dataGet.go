package dataGet

import (
	"fmt"
	"go.lwh.com/linweihao/calculatorOfAutomaticInvestmentPlan/utils/calc"
	"go.lwh.com/linweihao/calculatorOfAutomaticInvestmentPlan/utils/err"
	"go.lwh.com/linweihao/calculatorOfAutomaticInvestmentPlan/utils/rfl"
	"go.lwh.com/linweihao/calculatorOfAutomaticInvestmentPlan/utils/time"
)

type EntityCalculator struct {
	SumTimes        string
	AmountBegin     string
	AmountAdd       string
	RateReal        string
	SumTimesInt     int
	RateToMul       string
	SumAmountAdd    string
	SumAmountCost   string
	SumAmountEarn   string
	SumAmountLast   string
	SumAmountAll    string
	AmountBeginLoop string
	RateShow        string
	MsgOut          string
}

func (self *EntityCalculator) Init() {
	self.RateToMul = calc.Add("1", self.RateReal)
	self.SumAmountAdd = "0"
	self.SumAmountCost = "0"
	self.SumAmountEarn = "0"
	self.SumAmountLast = self.AmountBegin
	self.SumAmountAll = self.AmountBegin
	self.AmountBeginLoop = "0"
	self.SumTimesInt = rfl.StrToInt(self.SumTimes)
	return
}

func (self *EntityCalculator) Calc() {
	self.Init()
	for times := 1; times <= self.SumTimesInt; times++ {
		self.SumAmountLast = self.SumAmountAll
		self.AmountBeginLoop = calc.Add(self.SumAmountLast, self.AmountAdd)
		self.SumAmountAll = calc.Mul(self.AmountBeginLoop, self.RateToMul)
		self.SumAmountAdd = calc.Add(self.SumAmountAdd, self.AmountAdd)
	}
	self.SumAmountCost = calc.Add(self.AmountBegin, self.SumAmountAdd)
	self.SumAmountEarn = calc.Sub(self.SumAmountAll, self.SumAmountCost)
	if self.SumAmountCost == "0" {
		self.RateShow = "0"
	} else {
		self.RateShow = calc.Div(self.SumAmountEarn, self.SumAmountCost)
	}
	// fmt.Println(self.RateShow)
	/*setStructValues*/
	self.MsgOut = fmt.Sprintf(
		"\n总成本为|%s|,\n总收益为|%s|,\n最终价值为|%s|,\n名义收益率为|%s|",
		self.SumAmountCost,
		self.SumAmountEarn,
		self.SumAmountAll,
		self.RateShow)
	return
}

func GetData(paramsIn rfl.Params) (paramsOut rfl.Params) {
	defer err.ThrowError()
	// fmt.Println(paramsIn)
	entityCalculator := &EntityCalculator{}
	// rfl.ShowType(entityCalculator)
	rfl.ToEntityFromAttr(paramsIn, entityCalculator)
	// fmt.Printf("%#v\n\n", entityCalculator)
	entityCalculator.Calc()
	// fmt.Printf("%#v\n\n", entityCalculator)
	paramsOut = rfl.ToAttrFromEntity(*entityCalculator)
	// fmt.Println(paramsOut)
	time.ShowTimeAndMsg("Data get success")
	return paramsOut
}
