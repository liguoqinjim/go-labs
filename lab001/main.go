package main

import (
	"errors"
	"flag"
	"fmt"
	"strings"
	"time"
)

//Example 1 解析一个参数，a.png是default value
var fileName = flag.String("fileName", "a.png", "config file name")

//Example 2 两个参数解析到一个变量，在参数简写的时候可以用到
var gopherType string

func init() {
	const (
		defaultGopher = "pocket"
		usage         = "the variety of gopher"
	)
	flag.StringVar(&gopherType, "gopher_type", defaultGopher, usage)
	flag.StringVar(&gopherType, "g", defaultGopher, usage+" (shorthand)")
}

//Example 3 自定义参数类型
type interval []time.Duration

func (i *interval) String() string {
	return fmt.Sprint(*i)
}

func (i *interval) Set(value string) error {
	if len(*i) > 0 {
		return errors.New("interval flag already set")
	}

	for _, dt := range strings.Split(value, ",") {
		duration, err := time.ParseDuration(dt)
		if err != nil {
			return err
		}
		*i = append(*i, duration)
	}
	return nil
}

var intervalFlag interval

func init() { //可以有多个init方法
	flag.Var(&intervalFlag, "deltaT", "comma-separated list of intervals to use between events")
}

func main() {
	flag.Parse() //解析参数

	//Example 1
	fmt.Println("fileName =", *fileName)

	//Example 2
	fmt.Println("gopherType =", gopherType)

	//Example 3
	fmt.Println("intervalFlag =", intervalFlag)
}
