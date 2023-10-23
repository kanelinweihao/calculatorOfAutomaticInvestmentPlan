package main

/*commandSetEnv*/
// go env -w GO111MODULE=on
// go env -w GOPROXY=https://goproxy.cn,direct
// go env -w CGO_ENABLED=0
// go env -w GOOS=windows
// go env -w GOARCH=amd64
/*commandInitMod*/
// mkdir calculatorOfAutomaticInvestmentPlan
// cd ./calculatorOfAutomaticInvestmentPlan
// go mod init go.lwh.com/linweihao/calculatorOfAutomaticInvestmentPlan
/*commandImportMod*/
// go get -u github.com/shopspring/decimal
// go get -u github.com/akavel/rsrc@latest
/*commandBuild*/
// rsrc -manifest ./build/main.manifest -ico ./build/icon_go.ico -o rsrc.syso
// go fmt ./...
// go mod tidy
// GOOS=windows go build -o 定投计算器.exe
/*commandRun*/
// go mod edit -replace lwhFrameGo/utils/time=../lwhFrameGo/utils/time
// go run main.go

import (
	"embed"
	"fmt"
	"go.lwh.com/linweihao/calculatorOfAutomaticInvestmentPlan/config/env"
	"go.lwh.com/linweihao/calculatorOfAutomaticInvestmentPlan/config/version"
	"go.lwh.com/linweihao/calculatorOfAutomaticInvestmentPlan/exec/front/frontEnd"
	"go.lwh.com/linweihao/calculatorOfAutomaticInvestmentPlan/utils/err"
	"go.lwh.com/linweihao/calculatorOfAutomaticInvestmentPlan/utils/time"
	// t "lwhFrameGo/utils/time"
)

//go:embed public
var FilesPublic embed.FS

func init() {
	fmt.Println("\n")
	msgVersion := version.GetMsgVersion()
	time.ShowTimeAndMsg(msgVersion)
	fmt.Println("\n")
	defer err.ThrowError()
	env.EnbaleCPU()
	env.SetFilesPublic(FilesPublic)
	return
}

func main() {
	doIt()
	return
}

func doIt() {
	frontEnd.ExecFrontEnd()
	return
}
