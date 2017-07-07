### 测试package的自动生成

#### 生成golang的时候
要是在`.proto`文件中不写`package`的时候，自动生成的go文件的package会是`.proto`文件的文件名。
比如这个例子中，go文件的中的`package`就是params