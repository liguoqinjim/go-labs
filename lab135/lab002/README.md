### instance_example

#### 注意点
 - 注意，patch这种方法的时候`func (c *Client) Get(url string) (resp *Response, err error) {...}`，
 第一个参数要注意是`*Client`，如：`func(_ *http.Client, url string) (*http.Response, error)`

#### 运行结果
![Imgur](https://i.imgur.com/PCOVImF.png)