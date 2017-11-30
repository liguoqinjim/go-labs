/**
 * Author:        Tony.Shao
 * Email:         xiocode@gmail.com
 * Github:        github.com/xiocode
 * File:          waitgroup_wrapper.go
 * Description:   WaitGroup Wrapper
 */

package utils

import (
	"sync"
)

type WaitGroupWrapper struct {
	sync.WaitGroup
}

func (wg *WaitGroupWrapper) Wrap(function func()) {
	wg.Add(1)
	go func() {
		function()
		wg.Done()
	}()
}
