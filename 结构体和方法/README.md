# 结构体和方法
## 面向对象
- go语言仅支持封装，不支持继承和多态

## 结构的定义
```go
type treeNode struct {
	value int
	left  *treeNode
	right *treeNode
}

func (node treeNode) getValue() int {
	return node.value
}
func (node *treeNode) setValue(value int) {
	node.value = value
}
func main() {
	var root treeNode
	root.value = 1
	root.left = &treeNode{}
	root.right = &treeNode{}
	root.right.left = &treeNode{2, nil, nil}
	fmt.Println(root.getValue())
	root.setValue(4)
	fmt.Println(root.getValue())
}
```
- 要改变内容必须使用指针接收者
- 结构过大也考虑使用指针接收者
- 一致性：如有指针接收者，最好都用指针接收者
- 值接收者是go语言特有