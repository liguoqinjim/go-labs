# [selenium](https://github.com/tebeka/selenium)

|实验|简介|说明|
|---|---|---|
|lab001|example|调用firefox|
|lab002|demo|调用chrome|

## NOTICE
 - 需要把`selenium-server-standalone-3.xx.jar`和相关驱动如`geckodriver.exe`放到根目录下
 - 现在这个库和selenium的3.14版本有点冲突，所有改了点代码才能用。之后会提交pr，修改了`func (s *Service) start(port int) error`方法