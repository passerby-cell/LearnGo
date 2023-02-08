# Map
## 定义
```go
m1 := map[string]string{
    "name": "go",
    "version": "1.20",
}
//定义空Map
m2 := make(map[string]int)  //m2 == empty map 推荐
var m3 map[string]int  //m3 == nil
fmt.Println(m1)
fmt.Println(m2)
fmt.Println(m3)
```
## 遍历
```go
for k, v := range m1 {
	fmt.Println(k, v)
}
```
## 判断key是否存在
```go
if val, isExist := m1["nama"]; isExist {
	fmt.Println(val)
} else {
	fmt.Println("not exist")
}
```
## 删除元素
```go
delete(m1, "name")
```
- 创建：make(map[string]string)
- 获取元素：m[key]
- key不存在时，获得Value类型的初始值
- 用value,ok:=m[key]来判断是否存在key

## map的key
- map使用哈希表，必须可以比较相等
- 除了slice，map，function的内建类型都可以作为key
- Struct类型不包含上述slice，map，function字段，也可作为key