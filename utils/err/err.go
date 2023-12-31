package err

import (
	"errors"
	"fmt"
	"path"
	"runtime"
)

func ErrCheck(err error) {
	if err == nil {
		return
	}
	panic(err)
}

func ErrPanic(msgError string) {
	err := errors.New(msgError)
	panic(err)
}

func ThrowError() {
	err := recover()
	if err == nil {
		return
	}
	fileName, funcName, codeLine := getLocationOfErr()
	fmt.Println("----------")
	fmt.Print("fileName = ")
	fmt.Println(fileName)
	fmt.Print("funcName = ")
	fmt.Println(funcName)
	fmt.Print("codeLine = ")
	fmt.Println(codeLine)
	fmt.Print("msgError = ")
	fmt.Println(err)
	fmt.Println("----------")
	fmt.Println("\n")
	fmt.Println("\n")
}

func getLocationOfErr() (fileName string, funcName string, codeLine int) {
	skip := 4
	pc, file, line, ok := runtime.Caller(skip)
	if !ok {
		return
	}
	fileName = path.Base(file)
	funcName = runtime.FuncForPC(pc).Name()
	codeLine = line
	return fileName, funcName, codeLine
}
