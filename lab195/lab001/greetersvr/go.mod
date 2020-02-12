module li.com/greetersvr

go 1.13

require (
	github.com/micro/go-micro/v2 v2.0.0
	github.com/micro/go-plugins/registry/consul/v2 v2.0.2 // indirect
	li.com/greetersvc v0.0.0-00010101000000-000000000000
)

replace li.com/greetersvc => ../greetersvc
