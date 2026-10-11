What prints?
``` go
func change(s []int) {
    s[0] = 100
    s = append(s, 200)
}

func main() {
    s := []int{1, 2, 3}

    change(s)

    fmt.Println(s)
}
```