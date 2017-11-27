package log4go

import (
	"io"
	"log"
)

type MyLog struct {
	Logger       *log.Logger
	DebugEnabled *bool
}

func New(out io.Writer) *MyLog {
	mylog := new(MyLog)
	mylog.Logger = log.New(out, "", log.LstdFlags)
	t := false
	mylog.DebugEnabled = &t
	return mylog
}

func (l *MyLog) Debug(arg ...interface{}) {
	if !*l.DebugEnabled {
		return
	}
	arg = append([]interface{}{"[debug]"}, arg...)
	l.Logger.Println(arg...)
}

func (l *MyLog) Debugf(format string, arg ...interface{}) {
	if !*l.DebugEnabled {
		return
	}
	l.Logger.Printf("[debug] "+format, arg...)
}

func (l *MyLog) Info(arg ...interface{}) {
	arg = append([]interface{}{"[info]"}, arg...)
	l.Logger.Println(arg...)
}

func (l *MyLog) Error(arg ...interface{}) {
	arg = append([]interface{}{"[error]"}, arg...)
	l.Logger.Println(arg...)
}

func (l *MyLog) Errorf(format string, arg ...interface{}) {
	l.Logger.Printf("[error] "+format, arg...)
}
