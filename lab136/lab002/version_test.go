package main

import (
	"github.com/golang/mock/gomock"
	"lab136/lab002/spider"
	"testing"
)

//这个测试会错误，因为我们还没有写CreateGoVersionSpider
func TestGetGoVersion(t *testing.T) {
	v := GetGoVersion(spider.CreateGoVersionSpider())

	if v != "go1.8.3" {
		t.Errorf("got '%s' want 'go1.8.3'", v)
	}
}

//使用mock来测试
func TestGetGoVersion2(t *testing.T) {
	mockCtl := gomock.NewController(t)
	mockSpider := spider.NewMockSpider(mockCtl)
	mockSpider.EXPECT().GetBody().Return("go1.8.3")
	goVer := GetGoVersion(mockSpider)

	if goVer != "go1.8.3" {
		t.Errorf("got '%s' want 'go1.8.3'", goVer)
	}
}
