package main

import "lab136/lab002/spider"

func GetGoVersion(s spider.Spider) string {
	body := s.GetBody()
	return body
}
