# 直接存struct

## NOTICE
 - 需要实现`encoding.BinaryMarshaler`和`encoding.BinaryUnmarshaler`

## 例子
```
func (u *UUIDLogin) MarshalBinary() (data []byte, err error) {
	return json.Marshal(u)
}
func (u *UUIDLogin) UnmarshalBinary(data []byte) error {
	err := json.Unmarshal(data, &u)
	return err
}
```