# Go Template

常用数据结构的 Go 泛型实现

## 安装

```bash
go get github.com/Repeater11/go-template
```

## 快速开始

```go
// 导入所需的包
import "github.com/Repeater11/go-template/structure/vector"

// 使用 Vector
v := vector.NewVector(1, 2, 3)
v.PushBack(4)
vector.Sort(v)
```

更多数据结构的使用示例，请查看各子包文档。

## 已实现

| 模块                         | 说明     | 文档                                                        |
| ---------------------------- | -------- | ----------------------------------------------------------- |
| [vector](./structure/vector) | 动态数组 | `go doc github.com/Repeater11/go-template/structure/vector` |
| [deque](./structure/deque)   | 双端队列 | `go doc github.com/Repeater11/go-template/structure/deque`  |

## 计划实现

- [ ] list - 双向链表
- [ ] stack - 栈
- [ ] queue - 队列
- [ ] set - 集合
- [ ] map - 映射

## 文档

查看各模块文档：

```bash
# Vector
go doc github.com/Repeater11/go-template/structure/vector

# Deque
go doc github.com/Repeater11/go-template/structure/deque

# 将来的其他模块...
# go doc github.com/Repeater11/go-template/structure/list
```

## 测试

```bash
# 测试所有模块
go test ./...

# 测试特定模块
go test ./structure/vector/...
go test ./structure/deque/...

# 测试覆盖率
go test -cover ./...
```

## 许可证

MIT
