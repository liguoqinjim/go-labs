package spider

type Spider interface {
	GetBody() string
}

//还没有完成就要测试，这时候可以用mock
func CreateGoVersionSpider() Spider {
	return nil
}
