# 资源管理与错误处理
## defer调用
- 确保调用在该函数结束时发生
- 参数在defer语句时计算
- defer列表为后进先出

## 何时调用defer
- Open/Close
- Lock/Unlock
- PringHeader/PrintFooter