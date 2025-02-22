# listplus

## Description

Lists written based on golang `list` that support generics and sorting

## Installation

```bash
go get github.com/lqxhub/listplus
```

### Usage

Same as golang's `list` usage, except that generics and sorting are supported

```go
list := New[int]()
for _, v := range rand.Perm(1000) {
    list.PushBack(v)
}
list.Sort(func (front, back int) bool {
    return front > back
})
```
