
# listplus

## 描述

基于 golang `list` 修改，支持泛型和排序的 `list`

## 安装

```bash
go get github.com/lqxhub/listplus
```

### 使用

与 golang 的 `list` 用法相同，只是支持泛型和排序

```go
list := New[int]()
for _, v := range rand.Perm(1000) {
    list.PushBack(v)
}
list.Sort(func (front, back int) bool {
    return front > back
})
```
