package main

import (
	"fmt"
	"net/url"
)

func main() {
	webpage := "http://mywebpage.com/thumbify"
	image := "http://images.com/cat.png"
	fmt.Println(webpage + "?image=" + url.QueryEscape(image))

	a := "opA%2BQIxz40y7oK%2Fdr%2BCGM%2F4VtSefmfPBe9gR%2BUpZ8nWEawWQQYSMrHgRy4M1Kk55S7iNjQcf4cNw0pLjEP4%2B54LW4PxwN%2Brpk14YawAFcihyCcJRNA0QXZRT8mnNTx9xDzbDUJHe1siqPmIQQ97vpL2y4S5Shm5sR4JuTYvNfeI2MiAhlHujmrjjE4LY1S%2Fy"
	fmt.Println(url.QueryUnescape(a))

	s := "上海"
	fmt.Println(url.QueryEscape(s))
}
