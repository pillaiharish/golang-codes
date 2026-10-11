# GO Traps
I’d learn this in four passes:

1. **Go syntax traps**
2. **Slices**
3. **Maps**
4. **Interfaces**
5. Then finish with **15–20 output/bug-finding questions** combining all four.

---

# 1. Go syntax traps

## 1.1 `:=` vs `=`

`:=` declares variables.

```go
x := 10
```

`=` assigns to existing variables.

```go
x = 20
```

The important trap:

```go
x := 10

if true {
    x := 20
    fmt.Println(x)
}

fmt.Println(x)
```

Output:

```text
20
10
```

The inner `x` is a **new variable** because of lexical scope.

This is called **shadowing**.

A very common production bug:

```go
var err error

if condition {
    result, err := doSomething()
    fmt.Println(result)
}

if err != nil {
    // ...
}
```

The `err` inside the `if` may be a different variable.

### But this is valid:

```go
x := 10

x, y := 20, 30
```

Why?

Because at least **one variable (`y`) is new**.

Equivalent conceptual result:

```go
x = 20
y := 30
```

Interview rule:

> `:=` requires at least one new variable in the same scope.

---

# 2. Unused variables are compile errors

This doesn't compile:

```go
func main() {
    x := 10
}
```

Go requires local variables to be used.

Same for imports.

```go
import "strings"
```

If unused:

```text
"strings" imported and not used
```

You can intentionally discard values using `_`.

```go
value, _ := strconv.Atoi("123")
```

`_` is the **blank identifier**.

---

# 3. Multiple return values

Extremely common in Go:

```go
func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("divide by zero")
    }

    return a / b, nil
}
```

Usage:

```go
result, err := divide(10, 2)
```

Ignore one:

```go
result, _ := divide(10, 2)
```

---

# 4. `if` can have initialization

Very Go-specific syntax:

```go
if err := doSomething(); err != nil {
    fmt.Println(err)
}
```

`err` only exists inside the `if` / `else`.

```go
if value := getValue(); value > 10 {
    fmt.Println(value)
}

fmt.Println(value) // compile error
```

Interviewers love this scope question.

---

# 5. `for` is Go's only loop

There is no:

```text
while
do while
```

Classic loop:

```go
for i := 0; i < 10; i++ {
}
```

Equivalent to while:

```go
for x < 10 {
    x++
}
```

Infinite:

```go
for {
}
```

Range:

```go
for i, v := range arr {
    fmt.Println(i, v)
}
```

Ignore index:

```go
for _, v := range arr {
}
```

Only index:

```go
for i := range arr {
}
```

---

# 6. `switch` does NOT fall through by default

In C/C++:

```c
case 1:
case 2:
```

can fall through.

Go doesn't.

```go
x := 1

switch x {
case 1:
    fmt.Println("one")
case 2:
    fmt.Println("two")
}
```

Prints:

```text
one
```

No `break` required.

Explicit fallthrough exists:

```go
switch x {
case 1:
    fmt.Println("one")
    fallthrough
case 2:
    fmt.Println("two")
}
```

Output:

```text
one
two
```

---

# 7. Arrays are values in Go

This surprises people coming from C/C++.

```go
a := [3]int{1, 2, 3}
b := a

b[0] = 100

fmt.Println(a)
fmt.Println(b)
```

Output:

```text
[1 2 3]
[100 2 3]
```

`b := a` copies the entire array.

Also:

```go
[3]int
[4]int
```

are **different types**.

---

# Now the important part: Slices

A slice is not an array.

```go
a := [5]int{1, 2, 3, 4, 5}
```

Array.

```go
s := []int{1, 2, 3, 4, 5}
```

Slice.

Think of a slice approximately as:

```text
+--------------------+
| pointer to array   |
| length             |
| capacity           |
+--------------------+
```

Conceptually:

```go
type slice struct {
    ptr *T
    len int
    cap int
}
```

This mental model explains almost every slice interview question.

---

# 8. `len` vs `cap`

Example:

```go
s := make([]int, 3, 5)

fmt.Println(len(s))
fmt.Println(cap(s))
```

Output:

```text
3
5
```

Underlying storage conceptually:

```text
[0 0 0 _ _]
 ↑
 len = 3
 cap = 5
```

You can access:

```go
s[0]
s[1]
s[2]
```

But:

```go
s[3]
```

panics even though capacity is 5.

Why?

Because indexing is limited by **length**, not capacity.

---

# 9. `make([]int, 3)` does NOT create an empty slice

Very common trap.

```go
s := make([]int, 3)
```

Produces:

```text
[0 0 0]
```

Then:

```go
s = append(s, 10)
```

Result:

```text
[0 0 0 10]
```

People sometimes intend:

```go
s := make([]int, 0, 3)
```

Then:

```go
s = append(s, 10)
```

Result:

```text
[10]
```

Memorize the difference:

```go
make([]int, 3)
```

means:

```text
len=3 cap=3
```

while

```go
make([]int, 0, 3)
```

means:

```text
len=0 cap=3
```

---

# 10. Slice shares underlying arrays

This is one of the most important Go concepts.

```go
a := []int{10, 20, 30, 40}

b := a[1:3]

b[0] = 999

fmt.Println(a)
```

Output:

```text
[10 999 30 40]
```

Why?

```text
a
        ┌────┬────┬────┬────┐
        │ 10 │ 20 │ 30 │ 40 │
        └────┴────┴────┴────┘
              ↑
              b
```

`b` did not copy the values.

It references the same backing array.

---

# 11. Slicing syntax

```go
s[a:b]
```

includes `a`, excludes `b`.

Example:

```go
s := []int{10, 20, 30, 40, 50}

fmt.Println(s[1:4])
```

Output:

```text
[20 30 40]
```

Same idea as:

```text
[1,4)
```

---

# 12. Slice capacity after slicing

Consider:

```go
s := []int{10, 20, 30, 40, 50}

x := s[1:3]

fmt.Println(len(x))
fmt.Println(cap(x))
```

Length:

```text
2
```

Capacity:

```text
4
```

Why?

`x` starts at original index `1`.

Backing array remaining:

```text
20 30 40 50
```

So:

```text
len = 2
cap = 4
```

This is interview-question territory.

---

# 13. `append` may or may not allocate

This is probably the most important slice trap.

Consider:

```go
a := make([]int, 2, 4)

a[0] = 10
a[1] = 20

b := append(a, 30)
```

There is available capacity.

So `b` probably uses the **same backing array**.

Conceptually:

```text
a -> [10 20 _ _]
b -> [10 20 30 _]
```

Now:

```go
b[0] = 999
fmt.Println(a[0])
```

prints:

```text
999
```

because backing storage is shared.

But once capacity is exhausted:

```go
a := []int{1, 2}

b := append(a, 3)
```

Go may allocate a **new backing array**.

Then modifications to `b` may no longer affect `a`.

Important rule:

> `append()` can change where a slice points.

Therefore always capture its return value:

```go
s = append(s, value)
```

Don't write:

```go
append(s, value)
```

In fact, Go won't let you discard it like that as a standalone call.

---

# 14. Slice function trap

Look carefully:

```go
func add(s []int) {
    s = append(s, 100)
}

func main() {
    s := []int{1, 2}
    add(s)

    fmt.Println(s)
}
```

Output remains:

```text
[1 2]
```

Why?

The **slice header itself is passed by value**.

Inside `add`:

```go
s = append(...)
```

changes the local slice variable.

Correct:

```go
func add(s []int) []int {
    return append(s, 100)
}

s = add(s)
```

However:

```go
func modify(s []int) {
    s[0] = 999
}
```

does affect the caller.

Why?

The copied slice header still points to the same backing array.

This distinction is extremely important.

---

# 15. `nil` slice vs empty slice

```go
var a []int

b := []int{}
```

Both:

```go
len(a) == 0
len(b) == 0
```

But:

```go
a == nil // true
b == nil // false
```

Both can be appended to:

```go
a = append(a, 10)
b = append(b, 10)
```

Usually both work similarly, but serialization can differ.

For example JSON commonly distinguishes them:

```text
nil slice   -> null
empty slice -> []
```

---

# 16. You cannot compare slices

This is illegal:

```go
a := []int{1, 2}
b := []int{1, 2}

fmt.Println(a == b)
```

Compile error.

Slices can only be compared to:

```go
nil
```

For modern Go, you can use:

```go
slices.Equal(a, b)
```

or manually compare.

---

# MAPS

Now maps.

```go
m := map[string]int{
    "alice": 10,
    "bob":   20,
}
```

Or:

```go
m := make(map[string]int)
```

---

# 17. Reading nonexistent map keys doesn't panic

Example:

```go
m := map[string]int{}

fmt.Println(m["hello"])
```

Output:

```text
0
```

Because the zero value of `int` is `0`.

For:

```go
map[string]string
```

missing key returns:

```text
""
```

For:

```go
map[string]bool
```

missing key returns:

```text
false
```

This causes an important ambiguity.

Suppose:

```go
m := map[string]int{
    "alice": 0,
}
```

Then:

```go
m["alice"]
m["bob"]
```

both return:

```text
0
```

How do we distinguish them?

---

# 18. Comma-ok idiom

```go
value, ok := m["alice"]
```

If key exists:

```text
value = value stored
ok = true
```

If not:

```text
value = zero value
ok = false
```

Example:

```go
if value, ok := m["alice"]; ok {
    fmt.Println("found:", value)
}
```

You'll see the same `value, ok` pattern elsewhere in Go:

```go
v, ok := map[key]
```

and later:

```go
v, ok := x.(SomeType)
```

for interfaces.

---

# 19. Nil map: read okay, write panic

This is a classic assessment question.

```go
var m map[string]int
```

`m` is nil.

Reading:

```go
fmt.Println(m["hello"])
```

works.

Returns:

```text
0
```

But:

```go
m["hello"] = 10
```

panics:

```text
assignment to entry in nil map
```

Fix:

```go
m = make(map[string]int)
```

---

# 20. Map iteration order is unspecified

Never depend on this:

```go
for k, v := range m {
    fmt.Println(k, v)
}
```

Given:

```go
m := map[string]int{
    "a": 1,
    "b": 2,
    "c": 3,
}
```

You cannot say the output will necessarily be:

```text
a 1
b 2
c 3
```

Map iteration order is intentionally unspecified.

If ordering matters:

```go
keys := make([]string, 0, len(m))

for k := range m {
    keys = append(keys, k)
}

sort.Strings(keys)

for _, k := range keys {
    fmt.Println(k, m[k])
}
```

---

# 21. Delete doesn't complain if key doesn't exist

```go
delete(m, "abc")
```

Safe even if `"abc"` doesn't exist.

---

# 22. Maps aren't safe for unsynchronized concurrent writes

Later when we cover concurrency, remember:

```go
go func() {
    m["a"] = 1
}()

go func() {
    m["b"] = 2
}()
```

is unsafe.

You'd normally protect it using:

```go
sync.Mutex
```

or use:

```go
sync.Map
```

for suitable use cases.

We'll cover that properly in your `Mutex/RWMutex` section.

---

# INTERFACES

This part is essential because it distinguishes somebody who merely writes Go from somebody who understands Go.

---

# 23. Interfaces are implemented implicitly

Suppose:

```go
type Speaker interface {
    Speak() string
}
```

Now:

```go
type Dog struct{}

func (Dog) Speak() string {
    return "woof"
}
```

There is no:

```text
implements Speaker
```

keyword.

Because `Dog` has:

```go
Speak() string
```

it automatically satisfies `Speaker`.

Usage:

```go
func makeSound(s Speaker) {
    fmt.Println(s.Speak())
}

makeSound(Dog{})
```

Output:

```text
woof
```

This implicit satisfaction is central to Go's interface design.

---

# 24. Pointer receiver vs value receiver

Suppose:

```go
type Dog struct{}

func (d *Dog) Speak() string {
    return "woof"
}
```

Then:

```go
var s Speaker
```

This works:

```go
d := Dog{}
s = &d
```

But this does not:

```go
s = d
```

Why?

Because `Speak()` belongs to the method set of:

```text
*Dog
```

not:

```text
Dog
```

Important rule:

If methods use value receiver:

```go
func (Dog) Speak()
```

both:

```text
Dog
*Dog
```

satisfy the interface.

If method uses pointer receiver:

```go
func (*Dog) Speak()
```

only:

```text
*Dog
```

satisfies it.

This is a common interview question.

---

# 25. `any`

Previously Go used:

```go
interface{}
```

For example:

```go
func printAnything(v interface{}) {
}
```

Modern Go has the alias:

```go
any
```

So:

```go
func printAnything(v any) {
}
```

means essentially:

```go
func printAnything(v interface{}) {
}
```

Because an empty interface has zero required methods, virtually every type satisfies it.

---

# 26. Type assertion

Suppose:

```go
var x any = "hello"
```

You can retrieve the concrete value:

```go
s := x.(string)
```

Now:

```text
s = "hello"
```

But this:

```go
n := x.(int)
```

panics.

Safer form:

```go
n, ok := x.(int)

if !ok {
    fmt.Println("not an int")
}
```

Again we get:

```text
value, ok
```

Very idiomatic Go.

---

# 27. Type switch

```go
func describe(v any) {
    switch x := v.(type) {
    case int:
        fmt.Println("integer:", x)

    case string:
        fmt.Println("string:", x)

    case bool:
        fmt.Println("boolean:", x)

    default:
        fmt.Println("unknown")
    }
}
```

Example:

```go
describe(10)
describe("hello")
```

Output:

```text
integer: 10
string: hello
```

---

# 28. The infamous nil-interface trap

This one is senior-interview material.

Look at:

```go
type MyError struct{}

func (e *MyError) Error() string {
    return "something failed"
}

func getError() error {
    var e *MyError = nil
    return e
}

func main() {
    err := getError()

    fmt.Println(err == nil)
}
```

What prints?

Not:

```text
true
```

It prints:

```text
false
```

Why?

Conceptually an interface contains:

```text
(dynamic type, dynamic value)
```

A truly nil interface is:

```text
(nil, nil)
```

But `err` contains:

```text
(*MyError, nil)
```

Therefore the interface itself isn't nil.

Think:

```text
err
+------------------------+
| dynamic type: *MyError |
| dynamic value: nil     |
+------------------------+
```

versus:

```text
nil interface
+--------------------+
| type: nil          |
| value: nil         |
+--------------------+
```

So:

```go
err != nil
```

This bug appears in real Go programs.

Correct pattern:

```go
func getError() error {
    // ...
    return nil
}
```

Don't wrap a typed nil pointer inside an interface.

---

# 29. Interface design: accept interfaces, return concrete types

A useful Go design guideline is generally:

```go
func NewService() *Service
```

rather than:

```go
func NewService() ServiceInterface
```

and consumers accept the smallest interface they need.

For example:

```go
type Reader interface {
    Read([]byte) (int, error)
}
```

This is one reason Go's standard library composes so well.

You'll see tiny interfaces like:

```go
io.Reader
io.Writer
error
fmt.Stringer
```

For example `error` itself is simply:

```go
type error interface {
    Error() string
}
```

That's why your custom type becomes an error by implementing:

```go
Error() string
```

This connects directly with the error handling you just learned.

---

# 30. Four questions I would absolutely expect you to solve

### Question 1

What prints?

```go
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

Answer:

```text
[100 20]
[100 20 30]
```

Same backing array.

---

### Question 2

What happens?

```go
func main() {
    var m map[string]int

    fmt.Println(m["a"])

    m["a"] = 10
}
```

Answer:

```text
prints 0
then panics
```

Because:

```text
reading nil map → allowed
writing nil map → panic
```

---

### Question 3

What prints?

```go
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

Answer:

```text
[100 2 3]
```

The element mutation changes the shared backing array.

The append only changes the local slice header, and because this original slice likely has no spare capacity, it needs new backing storage.

---

### Question 4

Does this compile?

```go
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

No.

Because:

```go
Run()
```

is defined on:

```go
*Job
```

So this works:

```go
r = &j
```

---

# Assessment cheat sheet

Before the test, I want these rules almost automatic for you:

| Topic | Rule |
|---|---|
| `:=` | at least one variable must be new in current scope |
| Shadowing | inner `:=` may create different variable |
| Array | value type; assignment copies it |
| Slice | descriptor pointing to backing array |
| `len(slice)` | accessible elements |
| `cap(slice)` | room from slice start to backing-array end |
| `append` | may allocate new backing array |
| slice parameter | slice header passed by value |
| `s[i] = x` | can modify shared backing array |
| nil slice | append is valid |
| slice comparison | only against `nil` |
| missing map key | returns zero value |
| map comma-ok | distinguishes missing key |
| nil map read | works |
| nil map write | panic |
| map iteration | unspecified order |
| interface implementation | implicit |
| pointer receiver | usually requires pointer for interface satisfaction |
| type assertion | `v, ok := x.(T)` |
| nil interface | `(nil,nil)` |
| typed nil interface | `(type,nil)` and therefore `!= nil` |

## The mental models matter more than memorizing syntax

For slices, always draw:

```text
slice
 ┌─────────────┐
 │ ptr         │───────► backing array
 │ len         │
 │ cap         │
 └─────────────┘
```

For interfaces, always draw:

```text
interface
 ┌───────────────┐
 │ dynamic type  │
 │ dynamic value │
 └───────────────┘
```

Those two pictures will solve a surprising number of tricky Go questions.

Your next immediate practice should be **20 output-prediction questions on slices/maps/interfaces**, because understanding the theory isn't enough for a timed 75-minute assessment. After that, we can go straight into **goroutines → channels → select**, which is probably the highest-value Go-specific area remaining for the test.