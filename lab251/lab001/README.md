# lab001 
    
## mongodb->go struct
```bson
# mongodb命令行返回，db.Tag.find().limit(1).pretty()
{
	"_id" : ObjectId("5ffd6f5abbb9163e3085b551"),
	"uid" : 15,
	"mid" : 1955,
	"tag" : "dentist",
	"timestamp" : 1193435061
}    
```

```go
type Tag struct {
	ID        primitive.ObjectID `bson:"_id"`
	UID       int                `bson:"uid"`
	Mid       int                `bson:"mid"`
	Tag       string             `bson:"tag"`
	Timestamp int                `bson:"timestamp"`
}

```