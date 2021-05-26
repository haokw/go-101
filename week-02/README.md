# week 02 作业

我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应 该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码?

应该抛给上层，由调用方来判断 sql.ErrNoRows 是否符合预期，是否应该处理

```go
type notfound interface {
    NotFound() bool
}

func IsNotFound(err error) bool {
    nfe, ok := err.(notfound)
    return ok && nfe.NotFound()
}

func ReadName() error {
    var (
        id int
        username string
    )
    err := db.QueryRow(query, 1).Scan(&id, &username)
    return errors.Wrap(err, "open failed")
}

func main() {
    err := ReadName()
    if err != nil {
        if IsNotFound(err) {
            fmt.Printf("not found")
        } else {
            fmt.Printf("db error")
        }
        fmt.Printf("original error: %T %v", errors.Cause(err), errors.Cause(err))
        fmt.Printf("stack trace:\n%+v\n", err)
        os.Exit(1)
    }
}
```

## comment

虽然基本逻辑的最终含义差不多，但是有更优雅的实现方式，另外在错误码的定义上，会更建议传自定义的业务错误码出去。

对 not found 和其他错误进行转换成自定义错误常量，这样上层业务可以与sql层解耦。
使用 wrapf 对错误进行包装, 上层用 errors.Is 判断错误.

可参考：
dao:
```go
return errors.Wrapf(code.NotFound, fmt.Sprintf("sql: %s error: %v", sql, err))
```

如果dao是其他错误

```go
return errors.Wrapf(code.Internal, fmt.Sprintf("sql: %s error: %v", sql, err))
```

biz:

```go
if errors.Is(err, code.NotFound} {

}
```

## 答疑直播

参考 kratos error.go 实现

业务层不应该依赖底层的错误类型，包装为 自定义 notfind 类型
