package lab001

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSpec(t *testing.T) {
	// Only pass t into top-level Convey calls
	Convey("Given some integer with a starting value", t, func() {
		x := 1

		Convey("When the integer is incremented", func() {
			x++

			Convey("The value should be greater by one", func() {
				//So(x, ShouldEqual, 2) //测试成功
				So(x, ShouldEqual, 3) //测试失败
			})
		})
	})
}
