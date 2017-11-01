### 示例1

#### 注意点
 - 运行`go-bindata`的路径，就是在代码里面读取内容的根目录。比如这个实验里面直接写的是test.json，
 因为是直接在data文件夹里面执行的`go-bindata`
 - `go-bindata`生成的go文件的package name都是main，这个可能要自己改下。