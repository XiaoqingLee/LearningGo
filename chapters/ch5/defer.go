package ch5

import (
	"fmt"
	"runtime"
)

func GetInvokingFunctionName() string {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	return fmt.Sprintf("%s:%d %s\n", frame.File, frame.Line, frame.Function)
}

func returnOperator() int {
	fmt.Printf("calling %s\n", GetInvokingFunctionName())
	return 33
}

func deferFuncParam() int {
	fmt.Printf("calling %s\n", GetInvokingFunctionName())
	return 22
}
func FunctionContainingDefer() (returnedValue int) {
	returnedValue = 11
	fmt.Printf("calling %s\n", GetInvokingFunctionName())
	defer func(param int) {
		returnedValue = param
		fmt.Printf("calling %s\n", GetInvokingFunctionName())
	}(deferFuncParam())
	fmt.Printf("calling %s\n", GetInvokingFunctionName())
	return returnOperator()
}

//在遇到defer语句的时候会对defer后函数调用的参数进行计算.
//参数计算和函数调用不发生在同一时间
func Defer() {
	fmt.Printf("calling %s\n", GetInvokingFunctionName())
	returned := FunctionContainingDefer()
	fmt.Printf("calling %s\n", GetInvokingFunctionName())
	fmt.Printf("returnedValue: %d\n", returned)
}

/*
calling C:/Users/xiaoqing/GoLandProjects/LearningGo/chapters/ch5/defer.go:37 github.com/xiaoqingLee/learning-go/chapters/ch5.Defer

calling C:/Users/xiaoqing/GoLandProjects/LearningGo/chapters/ch5/defer.go:27 github.com/xiaoqingLee/learning-go/chapters/ch5.FunctionContainingDefer

calling C:/Users/xiaoqing/GoLandProjects/LearningGo/chapters/ch5/defer.go:22 github.com/xiaoqingLee/learning-go/chapters/ch5.deferFuncParam

calling C:/Users/xiaoqing/GoLandProjects/LearningGo/chapters/ch5/defer.go:32 github.com/xiaoqingLee/learning-go/chapters/ch5.FunctionContainingDefer

calling C:/Users/xiaoqing/GoLandProjects/LearningGo/chapters/ch5/defer.go:17 github.com/xiaoqingLee/learning-go/chapters/ch5.returnOperator

calling C:/Users/xiaoqing/GoLandProjects/LearningGo/chapters/ch5/defer.go:30 github.com/xiaoqingLee/learning-go/chapters/ch5.FunctionContainingDefer.func1

calling C:/Users/xiaoqing/GoLandProjects/LearningGo/chapters/ch5/defer.go:39 github.com/xiaoqingLee/learning-go/chapters/ch5.Defer

returnedValue: 22

Process finished with the exit code 0
*/

//return 和 defer 的陷阱
func bareReturn() (returnedValue int) { //返回 1
	returnedValue = 1
	return
}

func bareReturn2() (returnedValue int) { //返回 42
	returnedValue = 1
	return 42 //使用bare return的函数return语句后面如果有值, 那么这个值不会被忽略!
}

func ReturnAndDefer() (returnedValue int) { //返回 1
	defer func() { //在defer中可以修改返回值, 当且仅当返回值使用命名写法
		returnedValue = 1
	}()
	return returnedValue
}

func ReturnAndDefer2() int { //返回 0, 因为返回值列表并没有使用变量名, 所以最终返回的是在执行defer调用之前就已经计算完毕的return语句后面的操作数的值
	returnedValue := 0
	defer func() {
		returnedValue = 1
	}()
	return returnedValue
}

func ReturnAndDefer3() (returnedValue int) { //返回 1
	defer func() {
		returnedValue = 1
	}()
	return 42
}

func ReturnAndDefer4() (returnedValue int) { //返回 1
	defer func() {
		returnedValue = 1
	}()
	return
}

func TestReturnAndDefer() {
	fmt.Println(bareReturn())
	fmt.Println(bareReturn2())
	fmt.Println(ReturnAndDefer())
	fmt.Println(ReturnAndDefer2())
	fmt.Println(ReturnAndDefer3())
	fmt.Println(ReturnAndDefer4())
}
