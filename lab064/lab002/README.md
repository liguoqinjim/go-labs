### 增删改查
在mongodb中做一些增删改查的操作。

#### 注意点
##### ObjectId
我们有的时候想直接用objectId来操作document，这个时候我们就要在映射对象里面加上id字段。如下：
```
type Game struct {
	ID           bson.ObjectId `bson:"_id,omitempty"` //_id这样可以收到mongodb的id，omitempty可以在insert时候不插入这个值，而是由mongodb自动生成
	Winner       string        `bson:"winner"`
	OfficialGame bool          `bson:"official_game"`
	Location     string        `bson:"location"`
	StartTime    time.Time     `bson:"start"`
	EndTime      time.Time     `bson:"end"`
	Players      []Player      `bson:"players"`
}
```

#### 运行结果
![Imgur](http://i.imgur.com/2F4FgI8.png)

#### 参考资料
1. https://objectrocket.com/docs/mongodb_go_examples.html#example-document
2. http://www.open-open.com/lib/view/open1453851461605.html