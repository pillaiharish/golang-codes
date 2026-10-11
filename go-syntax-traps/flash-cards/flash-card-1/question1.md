What prints?
``` go
func main() {
    s := make([]int, 2, 4)

    s[0] = 10
    s[1] = 20

    x := append(s, 30)

    x[0] = 100

    fmt.Println(s)
    fmt.Println(x)
}
```