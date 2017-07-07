### 测试message中的field名称大小写、下划线的影响

#### protoc命令
`protoc -I=. --go_out=. .\params.proto`

#### `.proto`中的field名称和go文件中的变量名的对比
|`.proto`|`.go`|
|---|---|
|id|Id|
|NAME|NAME|
|MATH_SCORE|MATH_SCORE|
|ChineseSocre|ChineseSocre|
|englishscore|Englishscore|
|musicSocre|MusicSocre|