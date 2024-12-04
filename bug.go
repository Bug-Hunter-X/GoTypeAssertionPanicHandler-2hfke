```go
func main() {
    var i interface{} = 10
    j := i.(int32)
    fmt.Println(j) // this will panic if i is not int32
}
```