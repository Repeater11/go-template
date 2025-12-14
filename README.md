# Go Template

常用数据结构的 Go 泛型实现

## 安装

```bash
go get github.com/Repeater11/go-template
```

## 快速开始

```go
import "github.com/Repeater11/go-template/structure"

// Vector示例
v := structure.NewVector(1, 2, 3)
v.PushBack(4, 5)
structure.Sort(v)

// 更多示例和API文档，请查看各模块的文档
```

## 已实现

- ✅ **Vector** - 动态数组，类似 C++ `std::vector`

## 计划实现

- [ ] List - 双向链表
- [ ] Stack - 栈
- [ ] Queue - 队列
- [ ] Set - 集合
- [ ] Map - 映射

## 文档

查看完整文档：

```bash
go doc github.com/Repeater11/go-template/structure
```

## 测试

```bash
cd structure
go test -v
```

## 许可证

MIT
