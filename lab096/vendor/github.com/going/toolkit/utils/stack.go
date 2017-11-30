/**
 * Author:        Tony.Shao
 * Email:         xiocode@gmail.com
 * Github:        github.com/xiocode
 * File:          stack.go
 * Description:   Print Panic Stack
 */

package utils

import (
	"github.com/going/toolkit/log"
	"runtime"
)

func PrintPanicStack() {
	if x := recover(); x != nil {
		log.Errorf("%v", x)
		for i := 0; i < 10; i++ {
			funcName, file, line, ok := runtime.Caller(i)
			if ok {
				log.Errorf("Panic %v:[func:%v,file:%v,line:%v]\n", i, runtime.FuncForPC(funcName).Name(), file, line)
			}
		}
	}
}
