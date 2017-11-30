/**
 * Author:        Tony.Shao
 * Email:         xiocode@gmail.com
 * Github:        github.com/xiocode
 * File:          strftime.go
 * Description:   Go实现的Python strftime
 */
package strftime

import (
	"strings"
	"time"
)

var conversion = map[string]string{
	"B": "January", //月 英文 完整
	"b": "Jan",     //月 英文 缩写
	"m": "01",      //月 数字
	"A": "Monday",  //周 英文 完整
	"a": "Mon",     //周 缩写 完整
	"d": "02",      //日 数字
	"H": "15",      //时 24小时制 数字
	"I": "03",      //时 12小时制 数字
	"M": "04",      //分 数字
	"S": "05",      //秒 数字
	"Y": "2006",    //年 完整 数字
	"y": "06",      //年 缩写 数字
	"p": "PM",      //12小时制 上下午 AM PM
	"Z": "MST",     //时区
	"z": "-0700",   //时区 数字
}

// Go的时间格式好坑爹，还是按Python的来吧！！
func Format(format string, t time.Time) string {
	layout := layoutParser(format)
	return t.Format(layout)
}

// Go的时间格式好坑爹，还是按Python的来吧！！
func Parse(format, value string) (time.Time, error) {
	layout := layoutParser(format)
	return time.Parse(layout, value)
}

func layoutParser(format string) string {
	formatChunks := strings.Split(format, "%")
	var layout []string
	for _, chunk := range formatChunks {
		if len(chunk) == 0 {
			continue
		}
		if layoutCmd, ok := conversion[chunk[0:1]]; ok {
			layout = append(layout, layoutCmd)
			if len(chunk) > 1 {
				layout = append(layout, chunk[1:])
			}
		} else {
			layout = append(layout, "%", chunk)
		}
	}
	return strings.Join(layout, "")
}
