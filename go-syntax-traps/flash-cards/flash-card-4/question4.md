What prints?

```
type Runner interface {
    Run()
}

type Job struct{}

func (j *Job) Run() {}

func main() {
    var r Runner

    j := Job{}

    r = j
}
```