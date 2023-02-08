# 接口
## 接口的定义
```go
type Receiver interface {
	GetRes(s string) string
}
```
## 接口的实现
```go
type Receiver struct {
	Content string
}

func (receiver Receiver) GetRes(s string) string {
	return receiver.Content
}
```
```go
type Receiver struct {
}

func (receiver Receiver) GetRes(s string) string {
	response, err := http.Get(s)
	if err != nil {
		panic(err)
		return ""
	}
	dumpResponse, err := httputil.DumpResponse(response, true)
	response.Body.Close()
	if err != nil {
		panic(err)
	}
	return string(dumpResponse)
}
```
- 首字母大写才能向外暴露

## 接口的组合
```go
type ReceiverPoster interface {
	GetRes(s string) string
	GetResPoster(s string) string
}
```
## 常用系统接口
- Stringer
- Reader/Writer