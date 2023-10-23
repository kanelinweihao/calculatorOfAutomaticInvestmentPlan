package env

import (
	"embed"
	"fmt"
	"runtime"
	// "io/fs"
	"go.lwh.com/linweihao/calculatorOfAutomaticInvestmentPlan/config/version"
	// "go.lwh.com/linweihao/calculatorOfAutomaticInvestmentPlan/utils/err"
	"go.lwh.com/linweihao/calculatorOfAutomaticInvestmentPlan/utils/file"
	"go.lwh.com/linweihao/calculatorOfAutomaticInvestmentPlan/utils/rfl"
	// "go.lwh.com/linweihao/calculatorOfAutomaticInvestmentPlan/utils/time"
)

var (
	ModuleNameCn string = "定投计算器"
	ModuleNameEn string = "calculatorOfAutomaticInvestmentPlan"
	ParamsHtml          = rfl.Params{
		"Host": "127.0.0.1",
		"Port": "8080",
	}
	FilesPublic embed.FS
)

/*
OS
*/

func GetOSName() (osName string) {
	osName = runtime.GOOS
	// fmt.Println(osName)
	/*osArch := runtime.GOARCH
	  fmt.Println(osArch)
	  pathRoot := runtime.GOROOT()
	  fmt.Println(pathRoot)*/
	return osName
}

func EnbaleCPU() {
	/*numGoroutine := runtime.NumGoroutine()
	  fmt.Println(numGoroutine)*/
	numCPU := runtime.NumCPU()
	// fmt.Println(numCPU)
	runtime.GOMAXPROCS(numCPU)
	// numCPUEnabled := runtime.GOMAXPROCS(numCPU)
	// fmt.Println(numCPUEnabled)
	return
}

/*
Set
*/

func SetFilesPublic(FilesPublicFromEmbed embed.FS) {
	FilesPublic = FilesPublicFromEmbed
	return
}

/*
Get
*/

/*
Domain
*/

func GetDomain() (domain string) {
	domain = fmt.Sprintf(
		"%s:%s",
		ParamsHtml["Host"],
		ParamsHtml["Port"])
	// fmt.Println(domain)
	return domain
}

/*
Params
*/

func GetParamsTmpl() (paramsTmpl rfl.Params) {
	ModTitle := ModuleNameCn
	Version := version.GetVersion()
	paramsTmpl = rfl.Params{
		"ModTitle": ModTitle,
		"Version":  Version,
	}
	return paramsTmpl
}

/*
File
*/

func GetFSOfTmpl(pathTmpl string) (fsPublic embed.FS, patternsTmpl string) {
	// fmt.Println(pathTmpl)
	// pathTmplRel := file.GetFilePathRel(pathTmpl)
	// fmt.Println(pathTmplRel)
	fsPublic = FilesPublic
	patternsTmpl = file.GetFilePathEmbed(pathTmpl)
	return fsPublic, patternsTmpl
	/*var fsTmpl fs.File
	  fsTmpl,errFsOpen := FilesPublic.Open(pathTmplRel)
	  err.ErrCheck(errFsOpen)
	  // rfl.ShowType(fsTmpl)
	  return fsTmpl*/
}
