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
