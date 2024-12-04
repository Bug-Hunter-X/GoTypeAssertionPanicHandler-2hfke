```go
func main() {
    var i interface{} = 10
    switch j := i.(type) {
    case int32:
        fmt.Println("int32", j)
    case int:
        fmt.Println("int", j)
    default:
        fmt.Println("Unknown type")
    }
}
```