# 数组
## 定义
```go
var arr1 [5]int
arr2:=[3]int{1,2,3}
arr3:=[...]int{1,2,3}
var grid [4][5]bool
```
## 遍历
```go
for _, v := range arr3 {
fmt.Println(v)
}

for i, v := range arr3 {
fmt.Println(i, v)
}
```
- 可获取数组的下表和值
- 可通过_省略变量
- 不仅range，任何地方都可以通过_省略变量

## 为什么使用range
- 意义明确，美观
- C++没有类似能力
- Java/Python：只能for each value，原生不能同时获取下标和值

## 数组是值类型
- [10]int和[20]int是不同的类型
- 调用func f(arr [10]int)会拷贝数组
- 在Go中一般不直接使用数组