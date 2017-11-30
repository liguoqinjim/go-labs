/**
 * Author:        Tony.Shao
 * Email:         xiocode@gmail.com
 * Github:        github.com/xiocode
 * File:          random.go
 * Description:   random generator
 */

package random

import (
	"math/rand"
	"time"
)

func RandomInt64(min, max int64) int64 {
	rand.Seed(time.Now().Unix())
	return rand.Int63n(max-min) + min
}
