# Layer 1 God Prep
## What Layer 1 should mean

Layer 1 is **not** interview depth.  
Layer 1 is:

> “I can read, write, run, and understand simple Go programs without constantly looking up syntax.”

At Layer 1, your goal is **familiarity**, not mastery.

You should avoid these in Layer 1:

- deep slice internals
- deep concurrency internals
- garbage collection details
- system design
- advanced interface internals
- performance tuning
- large project architecture

Those come later.

Layer 1 should make you comfortable with:

- Go tooling
- program structure
- variables
- basic types
- operators
- `if`
- `for`
- `switch`
- functions
- basic errors
- basic collections
- basic strings
- basic packages/modules
- simple input/output
- very basic testing

---

# Layer 1 master topic list

Here is the high-level list first:

```text
1. Environment and tooling
2. Anatomy of a Go program
3. Printing and formatting output
4. Variables, constants, and zero values
5. Basic types and type conversion
6. Operators and expressions
7. Conditional statements: if and switch
8. Loops: for and range
9. Functions
10. Basic error handling
11. Basic collections: arrays, slices, maps
12. Strings and runes basics
13. Structs and pointers preview
14. Packages, modules, and imports
15. Basic input and command-line arguments
16. Basic testing
17. Debugging and common beginner compile errors
18. Layer 1 practice exercises
```

Now let’s go deeply into each one.

---

# 1. Environment and tooling

This topic is important because many beginners struggle not with Go syntax, but with running and organizing Go code.

---

## 1.1 Installing and verifying Go

### 1.1.1 Check Go installation

Example:

```bash
go version
```

Expected output similar to:

```text
go version go1.23.0 darwin/arm64
```

Explanation:

This confirms Go is installed and available in your terminal.

---

### 1.1.2 Check Go environment

Example:

```bash
go env GOPATH GOMODCACHE
```

Explanation:

- `GOPATH` is your Go workspace directory.
- `GOMODCACHE` is where downloaded modules are cached.

For Layer 1, you do not need deep understanding. Just know Go uses modules and caches dependencies.

---

## 1.2 Creating a Go project

### 1.2.1 Create a folder

Example:

```bash
mkdir hello-go
cd hello-go
```

Explanation:

This creates your project directory.

---

### 1.2.2 Initialize a module

Example:

```bash
go mod init example.com/hello
```

This creates:

```text
go.mod
```

Example `go.mod`:

```go
module example.com/hello

go 1.23
```

Explanation:

A Go module is a unit of Go code with dependencies.

For Layer 1, think of `go.mod` as:

> “This folder is a Go project.”

---

## 1.3 Running Go programs

### 1.3.1 Create `main.go`

Example:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, Go")
}
```

Explanation:

- `package main` means this is an executable program.
- `import "fmt"` imports the formatting package.
- `func main` is the entry point.
- `fmt.Println` prints text.

---

### 1.3.2 Run with `go run`

Example:

```bash
go run main.go
```

Or:

```bash
go run .
```

Explanation:

- `go run main.go` runs one file.
- `go run .` runs the current package.

Use `go run .` once your project has multiple files.

---

### 1.3.3 Build an executable

Example:

```bash
go build
```

Then run:

```bash
./hello-go
```

Explanation:

`go build` compiles your program into a binary.

For Layer 1, `go run` is enough, but you should know `go build` exists.

---

## 1.4 Formatting and basic tooling

### 1.4.1 Format code

Example:

```bash
gofmt -w main.go
```

Or:

```bash
go fmt ./...
```

Explanation:

Go has standard formatting.

You should not argue about formatting in Go. You format the code.

---

### 1.4.2 Use `go vet`

Example:

```bash
go vet ./...
```

Explanation:

`go vet` finds suspicious code.

For Layer 1, just know it exists.

---

### 1.4.3 Read documentation

Example:

```bash
go doc fmt.Println
```

Explanation:

This shows documentation for `fmt.Println`.

You should learn to read Go docs early.

---

# 2. Anatomy of a Go program

This is the foundation.

---

## 2.1 Package declaration

### 2.1.1 Every Go file belongs to a package

Example:

```go
package main
```

Explanation:

Every Go file starts with a package declaration.

For executable programs, the main package must be:

```go
package main
```

---

### 2.1.2 Package names are usually short

Example:

```go
package mathutil
```

Explanation:

Go package names are usually short and lowercase.

Examples:

```go
package http
package json
package user
package auth
```

---

## 2.2 Imports

### 2.2.1 Single import

Example:

```go
import "fmt"
```

Explanation:

This imports the standard library `fmt` package.

---

### 2.2.2 Multiple imports

Example:

```go
import (
    "fmt"
    "math"
)
```

Explanation:

Use grouped imports when importing multiple packages.

---

### 2.2.3 Unused imports are compile errors

Bad example:

```go
import "fmt"

func main() {
}
```

Error:

```text
imported and not used
```

Explanation:

Go does not allow unused imports.

---

## 2.3 The `main` function

### 2.3.1 Entry point of executable programs

Example:

```go
package main

import "fmt"

func main() {
    fmt.Println("Program started")
}
```

Explanation:

When you run a Go executable program, Go calls:

```go
func main()
```

The `main` function:

- has no parameters
- has no return value

---

## 2.4 Comments

### 2.4.1 Single-line comments

Example:

```go
// This prints a greeting
fmt.Println("Hello")
```

Explanation:

Use `//` for single-line comments.

---

### 2.4.2 Multi-line comments

Example:

```go
/*
This is a multi-line comment.
It can span multiple lines.
*/
```

Explanation:

Use `/* */` for block comments, but in Go, single-line comments are more common.

---

# 3. Printing and formatting output

This is essential for learning and debugging.

---

## 3.1 Basic printing

### 3.1.1 `fmt.Println`

Example:

```go
fmt.Println("Hello")
fmt.Println("World")
```

Output:

```text
Hello
World
```

Explanation:

`Println` adds a newline automatically.

---

### 3.1.2 `fmt.Print`

Example:

```go
fmt.Print("Hello ")
fmt.Print("World")
```

Output:

```text
Hello World
```

Explanation:

`Print` does not add a newline automatically.

---

## 3.2 Formatted printing

### 3.2.1 `fmt.Printf`

Example:

```go
name := "Go"
version := 1.23

fmt.Printf("Language: %s, Version: %.2f\n", name, version)
```

Output:

```text
Language: Go, Version: 1.23
```

Explanation:

- `%s` prints a string
- `%.2f` prints a float with 2 decimal places
- `\n` adds a newline

---

## 3.3 Common format verbs

### 3.3.1 `%v` and `%T`

Example:

```go
x := 42

fmt.Printf("%v\n", x)
fmt.Printf("%T\n", x)
```

Output:

```text
42
int
```

Explanation:

- `%v` prints value
- `%T` prints type

---

### 3.3.2 Common verbs table

Example:

```go
fmt.Printf("%d\n", 10)
fmt.Printf("%f\n", 3.14)
fmt.Printf("%s\n", "hello")
fmt.Printf("%q\n", "hello")
fmt.Printf("%t\n", true)
fmt.Printf("%v\n", []int{1, 2, 3})
```

Output:

```text
10
3.140000
hello
"hello"
true
[1 2 3]
```

Explanation:

- `%d`: integer
- `%f`: float
- `%s`: string
- `%q`: quoted string
- `%t`: boolean
- `%v`: default value format

---

## 3.4 Creating strings with formatting

### 3.4.1 `fmt.Sprintf`

Example:

```go
message := fmt.Sprintf("Hello %s", "Go")
fmt.Println(message)
```

Output:

```text
Hello Go
```

Explanation:

`Sprintf` returns a string instead of printing it.

---

# 4. Variables, constants, and zero values

This is core syntax.

---

## 4.1 Variable declaration

### 4.1.1 Declare with type and value

Example:

```go
var age int = 30
```

Explanation:

This creates a variable named `age` of type `int` with value `30`.

---

### 4.1.2 Declare with type inference

Example:

```go
var city = "Berlin"
```

Explanation:

Go infers that `city` is a `string`.

---

### 4.1.3 Short variable declaration

Example:

```go
score := 95
```

Explanation:

`:=` declares and initializes a variable.

This is very common inside functions.

---

## 4.2 Multiple variables

### 4.2.1 Declare multiple variables

Example:

```go
var (
    host string
    port int
)
```

Explanation:

This declares two variables:

- `host` as string
- `port` as int

---

### 4.2.2 Multiple assignment

Example:

```go
a, b := 10, 20
```

Explanation:

This creates two variables in one line.

---

### 4.2.3 Swap values

Example:

```go
a := 1
b := 2

a, b = b, a

fmt.Println(a, b)
```

Output:

```text
2 1
```

Explanation:

Go supports multiple assignment, so swapping is simple.

---

## 4.3 Zero values

### 4.3.1 Variables have default zero values

Example:

```go
var i int
var f float64
var s string
var ok bool
var p *int
var sl []int
var m map[string]int
```

Explanation:

Their zero values are:

```text
int       0
float64   0.0
string    ""
bool      false
pointer   nil
slice     nil
map       nil
```

This is very important in Go.

---

## 4.4 Constants

### 4.4.1 Basic constant

Example:

```go
const pi = 3.14159
```

Explanation:

Constants cannot be reassigned.

---

### 4.4.2 Typed constant

Example:

```go
const maxUsers int = 100
```

Explanation:

This constant is explicitly typed as `int`.

---

### 4.4.3 Constant group

Example:

```go
const (
    StatusOK       = 200
    StatusNotFound = 404
)
```

Explanation:

Use grouped constants for related values.

---

### 4.4.4 `iota`

Example:

```go
const (
    Sunday = iota
    Monday
    Tuesday
)
```

Explanation:

`iota` creates incrementing constants.

Values are:

```text
Sunday = 0
Monday = 1
Tuesday = 2
```

For Layer 1, just understand the basic idea.

---

# 5. Basic types and type conversion

Go is statically typed.

---

## 5.1 Integer types

### 5.1.1 Common integer types

Example:

```go
var a int = 10
var b int8 = 127
var c int64 = 9999999999
var d uint = 20
```

Explanation:

- `int` is the default integer type.
- `int8`, `int16`, `int32`, `int64` are sized integers.
- `uint` is unsigned integer.

For Layer 1, prefer `int` unless you have a reason.

---

## 5.2 Floating-point types

### 5.2.1 `float32` and `float64`

Example:

```go
var x float32 = 3.14
var y float64 = 3.1415926535
```

Explanation:

Use `float64` by default unless you need `float32`.

---

## 5.3 Boolean type

### 5.3.1 Bool values

Example:

```go
var isActive bool = true
var isDeleted bool = false
```

Explanation:

Booleans are either `true` or `false`.

Go does not treat numbers as booleans.

This is invalid:

```go
if 1 {
}
```

Correct:

```go
if true {
}
```

---

## 5.4 String type

### 5.4.1 Basic string

Example:

```go
var name string = "Go"
```

Explanation:

Strings are immutable sequences of bytes.

For Layer 1, just know strings are text values.

---

## 5.5 Byte and rune

### 5.5.1 Byte

Example:

```go
var b byte = 'A'
fmt.Println(b)
```

Output:

```text
65
```

Explanation:

A `byte` is an alias for `uint8`.

---

### 5.5.2 Rune

Example:

```go
var r rune = '🙂'
fmt.Println(r)
```

Explanation:

A `rune` represents a Unicode code point.

For Layer 1, just know:

- `byte` is one byte
- `rune` is one Unicode code point

---

## 5.6 Type conversion

### 5.6.1 Go does not convert implicitly

Bad example:

```go
var i int = 10
var f float64 = i
```

This fails.

Correct:

```go
var i int = 10
var f float64 = float64(i)
```

Explanation:

Go requires explicit conversion.

---

### 5.6.2 Numeric conversion

Example:

```go
x := 10
y := float64(x)
z := int(y)
```

Explanation:

- `float64(x)` converts int to float64
- `int(y)` converts float64 to int

---

### 5.6.3 String conversion

Example:

```go
s := "hello"
b := []byte(s)
s2 := string(b)
```

Explanation:

Strings can be converted to byte slices and back.

For Layer 1, just know the syntax.

---

## 5.7 Parsing and formatting numbers

### 5.7.1 String to int

Example:

```go
import "strconv"

n, err := strconv.Atoi("42")
if err != nil {
    fmt.Println("invalid number")
}
```

Explanation:

`strconv.Atoi` converts string to int.

---

### 5.7.2 Int to string

Example:

```go
s := strconv.Itoa(42)
```

Explanation:

`strconv.Itoa` converts int to string.

---

# 6. Operators and expressions

---

## 6.1 Arithmetic operators

### 6.1.1 Basic arithmetic

Example:

```go
a := 10
b := 3

fmt.Println(a + b)
fmt.Println(a - b)
fmt.Println(a * b)
fmt.Println(a / b)
fmt.Println(a % b)
```

Output:

```text
13
7
30
3
1
```

Explanation:

- `+` addition
- `-` subtraction
- `*` multiplication
- `/` division
- `%` remainder

Important: integer division truncates.

```go
10 / 3 = 3
```

---

## 6.2 Assignment operators

### 6.2.1 Basic assignment

Example:

```go
x := 10
```

---

### 6.2.2 Compound assignment

Example:

```go
x := 10
x += 5
x -= 2
x *= 3
x /= 2
```

Explanation:

`x += 5` means:

```go
x = x + 5
```

---

### 6.2.3 Increment and decrement

Example:

```go
x := 1
x++
x--
```

Explanation:

In Go, `x++` is a statement, not an expression.

This is invalid:

```go
y := x++
```

Correct:

```go
x++
y := x
```

---

## 6.3 Comparison operators

### 6.3.1 Basic comparisons

Example:

```go
a := 5
b := 10

fmt.Println(a == b)
fmt.Println(a != b)
fmt.Println(a < b)
fmt.Println(a <= b)
fmt.Println(a > b)
fmt.Println(a >= b)
```

Output:

```text
false
true
true
true
false
false
```

Explanation:

Comparison operators produce boolean values.

---

## 6.4 Logical operators

### 6.4.1 AND, OR, NOT

Example:

```go
age := 25
hasID := true

if age >= 18 && hasID {
    fmt.Println("allowed")
}
```

Explanation:

- `&&` means AND
- `||` means OR
- `!` means NOT

Example:

```go
if !hasID {
    fmt.Println("ID required")
}
```

---

## 6.5 Bitwise operators

For Layer 1, this can be light.

### 6.5.1 Basic bitwise example

Example:

```go
x := 6 // 110
y := 3 // 011

fmt.Println(x & y)  // 2
fmt.Println(x | y)  // 7
fmt.Println(x ^ y)  // 5
fmt.Println(x << 1) // 12
fmt.Println(x >> 1) // 3
```

Explanation:

- `&` AND
- `|` OR
- `^` XOR
- `<<` left shift
- `>>` right shift

For Layer 1, awareness is enough.

---

# 7. Conditional statements: `if` and `switch`

This is exactly the kind of thing Layer 1 must cover well.

---

## 7.1 Basic `if`

### 7.1.1 Simple condition

Example:

```go
age := 20

if age >= 18 {
    fmt.Println("adult")
}
```

Explanation:

If the condition is true, the block runs.

---

## 7.2 `if` with `else`

### 7.2.1 If-else example

Example:

```go
score := 70

if score >= 90 {
    fmt.Println("A")
} else {
    fmt.Println("not A")
}
```

Explanation:

If the first condition is false, the `else` block runs.

---

## 7.3 `else if`

### 7.3.1 Multiple conditions

Example:

```go
score := 82

if score >= 90 {
    fmt.Println("A")
} else if score >= 80 {
    fmt.Println("B")
} else if score >= 70 {
    fmt.Println("C")
} else {
    fmt.Println("F")
}
```

Output:

```text
B
```

Explanation:

Go checks conditions from top to bottom.

---

## 7.4 `if` with initialization statement

### 7.4.1 Declare variable inside `if`

Example:

```go
nums := []int{1, 2, 3}

if n := len(nums); n == 0 {
    fmt.Println("empty")
} else {
    fmt.Println("not empty")
}
```

Explanation:

This line:

```go
n := len(nums)
```

creates `n` only for the `if` scope.

This is a very common Go pattern.

---

## 7.5 Go has no ternary operator

### 7.5.1 No `? :`

Bad mindset from other languages:

```text
result = condition ? a : b
```

Go equivalent:

```go
var result int

if condition {
    result = a
} else {
    result = b
}
```

Explanation:

Go intentionally avoids the ternary operator.

---

## 7.6 Basic `switch`

### 7.6.1 Switch on a value

Example:

```go
status := 200

switch status {
case 200:
    fmt.Println("OK")
case 404:
    fmt.Println("Not Found")
default:
    fmt.Println("Unknown")
}
```

Output:

```text
OK
```

Explanation:

`switch` is often cleaner than many `else if` statements.

---

### 7.6.2 Multiple values in one case

Example:

```go
status := 204

switch status {
case 200, 204:
    fmt.Println("Success")
case 500:
    fmt.Println("Server error")
default:
    fmt.Println("Other")
}
```

Output:

```text
Success
```

Explanation:

One case can match multiple values.

---

## 7.7 Switch without condition

### 7.7.1 Switch as cleaner if-else

Example:

```go
score := 85

switch {
case score >= 90:
    fmt.Println("A")
case score >= 80:
    fmt.Println("B")
default:
    fmt.Println("C")
}
```

Output:

```text
B
```

Explanation:

A `switch` with no expression behaves like an if-else chain.

---

## 7.8 Switch with initialization

### 7.8.1 Declare variable in switch

Example:

```go
nums := []int{1, 2, 3}

switch n := len(nums); {
case n == 0:
    fmt.Println("empty")
case n < 5:
    fmt.Println("small")
default:
    fmt.Println("large")
}
```

Output:

```text
small
```

Explanation:

`n := len(nums)` is scoped to the switch.

---

## 7.9 Fallthrough

### 7.9.1 Go does not fall through by default

Example:

```go
x := 1

switch x {
case 1:
    fmt.Println("one")
case 2:
    fmt.Println("two")
}
```

Output:

```text
one
```

Explanation:

Only the matched case runs.

If you want fallthrough, use:

```go
fallthrough
```

But this is rare.

---

# 8. Loops: `for` and `range`

This is extremely important.

Go only has one loop keyword:

```go
for
```

But it can behave in multiple ways.

---

## 8.1 Classic counted loop

### 8.1.1 Basic for loop

Example:

```go
for i := 0; i < 5; i++ {
    fmt.Println(i)
}
```

Output:

```text
0
1
2
3
4
```

Explanation:

Three parts:

```go
init; condition; post
```

---

## 8.2 Condition-only loop

### 8.2.1 While-style loop

Example:

```go
count := 3

for count > 0 {
    fmt.Println(count)
    count--
}
```

Output:

```text
3
2
1
```

Explanation:

This behaves like a `while` loop in other languages.

---

## 8.3 Infinite loop

### 8.3.1 Loop forever

Example:

```go
for {
    fmt.Println("running")
    break
}
```

Explanation:

A `for` with no condition runs forever until `break`, `return`, or panic.

---

## 8.4 Looping over slices

### 8.4.1 `range` over slice

Example:

```go
nums := []int{10, 20, 30}

for i, v := range nums {
    fmt.Println(i, v)
}
```

Output:

```text
0 10
1 20
2 30
```

Explanation:

`range` gives:

- index
- value

---

### 8.4.2 Ignore index

Example:

```go
for _, v := range nums {
    fmt.Println(v)
}
```

Explanation:

Use `_` when you do not need the index.

---

### 8.4.3 Ignore value

Example:

```go
for i := range nums {
    fmt.Println(i)
}
```

Explanation:

This gives only indexes.

---

## 8.5 Looping over maps

### 8.5.1 `range` over map

Example:

```go
ages := map[string]int{
    "alice": 30,
    "bob":   25,
}

for name, age := range ages {
    fmt.Println(name, age)
}
```

Explanation:

`range` over map gives:

- key
- value

Important:

Map iteration order is random.

---

## 8.6 Looping over strings

### 8.6.1 `range` over string

Example:

```go
s := "hello"

for i, r := range s {
    fmt.Println(i, r)
}
```

Explanation:

`range` over string gives:

- byte index
- rune

---

## 8.7 `break` and `continue`

### 8.7.1 Stop loop early

Example:

```go
for i := 0; i < 10; i++ {
    if i == 3 {
        break
    }
    fmt.Println(i)
}
```

Output:

```text
0
1
2
```

Explanation:

`break` exits the loop.

---

### 8.7.2 Skip iteration

Example:

```go
for i := 0; i < 5; i++ {
    if i == 2 {
        continue
    }
    fmt.Println(i)
}
```

Output:

```text
0
1
3
4
```

Explanation:

`continue` skips to the next iteration.

---

## 8.8 Nested loops

### 8.8.1 Loop inside loop

Example:

```go
for i := 0; i < 3; i++ {
    for j := 0; j < 3; j++ {
        fmt.Println(i, j)
    }
}
```

Explanation:

Inner loop runs fully for each outer loop iteration.

---

## 8.9 Loop labels

### 8.9.1 Break outer loop

Example:

```go
outer:
for i := 0; i < 3; i++ {
    for j := 0; j < 3; j++ {
        if i == 1 && j == 1 {
            break outer
        }
        fmt.Println(i, j)
    }
}
```

Explanation:

The label `outer` lets `break outer` exit the outer loop.

For Layer 1, just know labels exist.

---

## 8.10 Common loop patterns

These are important for Layer 1.

---

### 8.10.1 Sum elements

Example:

```go
nums := []int{1, 2, 3, 4}

total := 0

for _, n := range nums {
    total += n
}

fmt.Println(total)
```

Output:

```text
10
```

Explanation:

Initialize accumulator, loop, add.

---

### 8.10.2 Count matching elements

Example:

```go
nums := []int{1, 2, 2, 3}

count := 0

for _, n := range nums {
    if n == 2 {
        count++
    }
}

fmt.Println(count)
```

Output:

```text
2
```

---

### 8.10.3 Find maximum

Example:

```go
nums := []int{3, 9, 1}

max := nums[0]

for _, n := range nums[1:] {
    if n > max {
        max = n
    }
}

fmt.Println(max)
```

Output:

```text
9
```

Explanation:

Assume first element is max, then compare.

---

### 8.10.4 Find element

Example:

```go
nums := []int{10, 20, 30}

found := false

for _, n := range nums {
    if n == 20 {
        found = true
        break
    }
}

fmt.Println(found)
```

Output:

```text
true
```

---

# 9. Functions

Layer 1 should cover function syntax deeply enough to write small programs.

---

## 9.1 Basic function declaration

### 9.1.1 Function with no parameters

Example:

```go
func greet() {
    fmt.Println("Hello")
}
```

Call:

```go
greet()
```

Explanation:

This function takes no input and returns nothing.

---

## 9.2 Function with parameters

### 9.2.1 One parameter

Example:

```go
func greet(name string) {
    fmt.Println("Hello", name)
}
```

Call:

```go
greet("Go")
```

Explanation:

The function expects a `string`.

---

### 9.2.2 Multiple parameters

Example:

```go
func add(a int, b int) int {
    return a + b
}
```

Call:

```go
result := add(2, 3)
fmt.Println(result)
```

Output:

```text
5
```

Explanation:

Parameters are typed individually.

You can also write:

```go
func add(a, b int) int {
    return a + b
}
```

---

## 9.3 Return values

### 9.3.1 Single return

Example:

```go
func square(x int) int {
    return x * x
}
```

---

### 9.3.2 Multiple return values

Example:

```go
func divide(a float64, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("division by zero")
    }
    return a / b, nil
}
```

Call:

```go
result, err := divide(10, 2)
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(result)
}
```

Explanation:

Go commonly returns:

```go
(value, error)
```

---

### 9.3.3 Ignore return value

Example:

```go
result, _ := divide(10, 2)
```

Explanation:

Use `_` to ignore a return value.

In real code, do not ignore errors unless you have a reason.

---

## 9.4 Named return values

### 9.4.1 Basic named returns

Example:

```go
func fullName(first string, last string) (name string) {
    name = first + " " + last
    return
}
```

Explanation:

The return variable is named `name`.

A bare `return` returns the named values.

For Layer 1, know syntax, but do not overuse named returns.

---

## 9.5 Variadic functions

### 9.5.1 Accept variable number of arguments

Example:

```go
func sum(nums ...int) int {
    total := 0

    for _, n := range nums {
        total += n
    }

    return total
}
```

Call:

```go
fmt.Println(sum(1, 2, 3))
fmt.Println(sum(1, 2, 3, 4, 5))
```

Output:

```text
6
15
```

Explanation:

`nums ...int` means the function can receive zero or more `int` values.

Inside the function, `nums` behaves like a slice.

---

## 9.6 `defer`

### 9.6.1 Basic defer

Example:

```go
func process() {
    defer fmt.Println("done")
    fmt.Println("working")
}
```

Output:

```text
working
done
```

Explanation:

Deferred statements run when the surrounding function returns.

For Layer 1, just know `defer` delays execution until function exit.

---

## 9.7 Functions as values

For Layer 1, keep this light.

### 9.7.1 Assign function to variable

Example:

```go
add := func(a int, b int) int {
    return a + b
}

fmt.Println(add(2, 3))
```

Output:

```text
5
```

Explanation:

Functions are values in Go.

---

## 9.8 Callbacks

### 9.8.1 Pass function as argument

Example:

```go
func apply(nums []int, f func(int) int) []int {
    result := make([]int, len(nums))

    for i, n := range nums {
        result[i] = f(n)
    }

    return result
}
```

Call:

```go
doubled := apply([]int{1, 2, 3}, func(n int) int {
    return n * 2
})

fmt.Println(doubled)
```

Output:

```text
[2 4 6]
```

Explanation:

This is a preview of functional patterns.

For Layer 1, this is optional but useful.

---

# 10. Basic error handling

Layer 1 should cover errors only at a basic level.

---

## 10.1 The `error` value

### 10.1.1 Functions often return errors

Example:

```go
func findUser(id int) (string, error) {
    if id <= 0 {
        return "", fmt.Errorf("invalid id: %d", id)
    }
    return "Alice", nil
}
```

Explanation:

`nil` error means success.

Non-nil error means failure.

---

## 10.2 Checking errors

### 10.2.1 Standard pattern

Example:

```go
name, err := findUser(1)
if err != nil {
    fmt.Println("error:", err)
    return
}

fmt.Println("user:", name)
```

Explanation:

This is the most common Go pattern:

```go
if err != nil {
    // handle error
}
```

---

## 10.3 Creating errors

### 10.3.1 `errors.New`

Example:

```go
import "errors"

func fail() error {
    return errors.New("something failed")
}
```

Explanation:

Creates a simple error.

---

### 10.3.2 `fmt.Errorf`

Example:

```go
func fail(id int) error {
    return fmt.Errorf("failed for id %d", id)
}
```

Explanation:

Creates formatted error message.

---

## 10.4 Panic awareness

### 10.4.1 Panic is not normal error handling

Example:

```go
panic("something terrible happened")
```

Explanation:

Panic is for serious unrecoverable situations.

In normal Go code, prefer returning errors.

For Layer 1, just know panic exists but should not be your default error mechanism.

---

# 11. Basic collections: arrays, slices, maps

This is where Layer 1 touches Layer 2. Keep it surface-level.

---

## 11.1 Arrays

### 11.1.1 Declare array

Example:

```go
var nums [3]int
```

Explanation:

This creates an array of length 3.

All values are zero initially:

```text
[0 0 0]
```

---

### 11.1.2 Initialize array

Example:

```go
nums := [3]int{1, 2, 3}
```

Explanation:

Array length is fixed at 3.

---

### 11.1.3 Access and update

Example:

```go
nums := [3]int{1, 2, 3}

fmt.Println(nums[0])

nums[0] = 10

fmt.Println(nums)
```

Output:

```text
1
[10 2 3]
```

Explanation:

Indexes start at 0.

---

### 11.1.4 Array length

Example:

```go
nums := [3]int{1, 2, 3}
fmt.Println(len(nums))
```

Output:

```text
3
```

---

## 11.2 Slices

### 11.2.1 Slice literal

Example:

```go
nums := []int{1, 2, 3}
```

Explanation:

This looks like an array, but has no fixed size in the type.

---

### 11.2.2 Append to slice

Example:

```go
nums := []int{1, 2}
nums = append(nums, 3)
```

Explanation:

`append` returns the updated slice.

Always assign it back:

```go
nums = append(nums, 3)
```

---

### 11.2.3 Create slice with `make`

Example:

```go
nums := make([]int, 5)
```

Explanation:

Creates slice with length 5.

Values are zero.

---

### 11.2.4 Length and capacity

Example:

```go
nums := make([]int, 0, 10)

fmt.Println(len(nums))
fmt.Println(cap(nums))
```

Output:

```text
0
10
```

Explanation:

- `len` is number of elements currently in slice
- `cap` is underlying capacity

For Layer 1, just know these exist.

---

### 11.2.5 Loop over slice

Example:

```go
nums := []int{1, 2, 3}

for i, n := range nums {
    fmt.Println(i, n)
}
```

---

## 11.3 Maps

### 11.3.1 Map literal

Example:

```go
ages := map[string]int{
    "alice": 30,
    "bob":   25,
}
```

Explanation:

This creates a map from string to int.

---

### 11.3.2 Set and get values

Example:

```go
ages := map[string]int{}

ages["alice"] = 30

fmt.Println(ages["alice"])
```

Output:

```text
30
```

---

### 11.3.3 Check if key exists

Example:

```go
age, ok := ages["bob"]

if ok {
    fmt.Println("bob exists:", age)
} else {
    fmt.Println("bob not found")
}
```

Explanation:

This is called the comma-ok pattern.

---

### 11.3.4 Delete key

Example:

```go
delete(ages, "bob")
```

Explanation:

Removes key `"bob"` from map.

---

### 11.3.5 Loop over map

Example:

```go
for name, age := range ages {
    fmt.Println(name, age)
}
```

Explanation:

Map iteration order is random.

---

# 12. Strings and runes basics

Strings are very common, so Layer 1 should cover basics.

---

## 12.1 String literals

### 12.1.1 Interpreted string

Example:

```go
s := "hello\nworld"
```

Explanation:

Escape sequences like `\n` work.

---

### 12.1.2 Raw string

Example:

```go
s := `hello
world`
```

Explanation:

Backticks create raw strings.

Escape sequences are not processed.

Useful for:

- JSON examples
- SQL queries
- multi-line text

---

## 12.2 String length

### 12.2.1 `len` counts bytes

Example:

```go
s := "hello"
fmt.Println(len(s))
```

Output:

```text
5
```

Explanation:

`len(s)` returns number of bytes, not necessarily number of visible characters.

---

## 12.3 String indexing

### 12.3.1 Index gives byte

Example:

```go
s := "hello"
fmt.Println(s[0])
```

Output:

```text
104
```

Explanation:

`s[0]` is a byte, not a character.

To print character:

```go
fmt.Printf("%c\n", s[0])
```

Output:

```text
h
```

---

## 12.4 Range over string

### 12.4.1 Iterate runes

Example:

```go
s := "héllo"

for i, r := range s {
    fmt.Println(i, r, string(r))
}
```

Explanation:

`range` decodes runes.

This is different from byte indexing.

---

## 12.5 String concatenation

### 12.5.1 Basic concatenation

Example:

```go
first := "Hello"
second := "Go"

message := first + " " + second
fmt.Println(message)
```

Output:

```text
Hello Go
```

---

### 12.5.2 Use `strings.Join`

Example:

```go
import "strings"

parts := []string{"go", "is", "fun"}
result := strings.Join(parts, " ")

fmt.Println(result)
```

Output:

```text
go is fun
```

---

## 12.6 Common string functions

### 12.6.1 Contains

Example:

```go
import "strings"

fmt.Println(strings.Contains("hello", "ell"))
```

Output:

```text
true
```

---

### 12.6.2 Split

Example:

```go
parts := strings.Split("a,b,c", ",")
fmt.Println(parts)
```

Output:

```text
[a b c]
```

---

### 12.6.3 ToUpper / ToLower

Example:

```go
fmt.Println(strings.ToUpper("hello"))
fmt.Println(strings.ToLower("HELLO"))
```

Output:

```text
HELLO
hello
```

---

## 12.7 String conversion

### 12.7.1 String to rune slice

Example:

```go
s := "hello"
runes := []rune(s)

fmt.Println(runes[0])
```

Explanation:

Converts string to slice of runes.

---

### 12.7.2 Rune slice to string

Example:

```go
runes := []rune{'h', 'i'}
s := string(runes)

fmt.Println(s)
```

Output:

```text
hi
```

---

# 13. Structs and pointers preview

This is technically early Layer 2, but for Layer 1, you should at least know the syntax.

---

## 13.1 Struct declaration

### 13.1.1 Define struct

Example:

```go
type User struct {
    Name string
    Age  int
}
```

Explanation:

A struct groups related fields.

---

## 13.2 Create struct

### 13.2.1 Struct literal

Example:

```go
u := User{
    Name: "Alice",
    Age:  30,
}
```

Explanation:

Creates a `User` value.

---

## 13.3 Access fields

### 13.3.1 Read field

Example:

```go
fmt.Println(u.Name)
```

---

### 13.3.2 Update field

Example:

```go
u.Age = 31
```

---

## 13.4 Zero value struct

### 13.4.1 Struct zero value

Example:

```go
var u User
```

Explanation:

Fields are zero values:

```go
u.Name == ""
u.Age == 0
```

---

## 13.5 Pointers basics

### 13.5.1 Address-of operator

Example:

```go
x := 10
p := &x

fmt.Println(p)
```

Explanation:

`&x` gives the memory address of `x`.

---

### 13.5.2 Dereference pointer

Example:

```go
x := 10
p := &x

fmt.Println(*p)

*p = 20

fmt.Println(x)
```

Output:

```text
10
20
```

Explanation:

`*p` accesses the value pointed to by `p`.

---

## 13.6 Why pointers matter

Example:

```go
func increment(n *int) {
    *n++
}
```

Call:

```go
x := 1
increment(&x)
fmt.Println(x)
```

Output:

```text
2
```

Explanation:

Pointers allow functions to modify original values.

For Layer 1, just understand:

- `&` gets address
- `*` reads/writes value at address

---

# 14. Packages, modules, and imports

This is essential for real Go code.

---

## 14.1 Package declaration

### 14.1.1 Files belong to packages

Example:

```go
package main
```

Explanation:

All Go files declare their package.

---

## 14.2 Exported names

### 14.2.1 Capital letters are exported

Example:

```go
package mathutil

func Add(a int, b int) int {
    return a + b
}

func subtract(a int, b int) int {
    return a - b
}
```

Explanation:

- `Add` is exported because it starts with capital `A`.
- `subtract` is not exported.

---

## 14.3 Local package example

Project structure:

```text
hello/
  go.mod
  main.go
  mathutil/
    add.go
```

`go.mod`:

```go
module example.com/hello

go 1.23
```

`mathutil/add.go`:

```go
package mathutil

func Add(a int, b int) int {
    return a + b
}
```

`main.go`:

```go
package main

import (
    "fmt"
    "example.com/hello/mathutil"
)

func main() {
    result := mathutil.Add(2, 3)
    fmt.Println(result)
}
```

Run:

```bash
go run .
```

Output:

```text
5
```

Explanation:

- `mathutil` is a local package.
- `Add` is exported.
- `main` imports it using module path.

---

## 14.4 Import standard library packages

### 14.4.1 Common packages

Example:

```go
import (
    "fmt"
    "strings"
    "strconv"
    "time"
)
```

Explanation:

For Layer 1, common packages to know:

```text
fmt
strings
strconv
errors
time
os
```

---

## 14.5 External modules

### 14.5.1 Add dependency

Example:

```bash
go get example.com/some/module
```

Explanation:

Downloads module and updates `go.mod`.

For Layer 1, standard library is enough, but know `go get` exists.

---

# 15. Basic input and command-line arguments

---

## 15.1 Reading input with `fmt.Scanln`

### 15.1.1 Read one value

Example:

```go
package main

import "fmt"

func main() {
    var name string

    fmt.Print("Enter your name: ")
    fmt.Scanln(&name)

    fmt.Println("Hello", name)
}
```

Explanation:

`&name` passes the address of `name` so `Scanln` can write into it.

---

## 15.2 Reading full lines

### 15.2.1 Use `bufio.Scanner`

Example:

```go
package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    fmt.Print("Enter text: ")
    scanner.Scan()

    text := scanner.Text()
    fmt.Println("You entered:", text)
}
```

Explanation:

`bufio.Scanner` is better for reading full lines.

---

## 15.3 Command-line arguments

### 15.3.1 Use `os.Args`

Example:

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    args := os.Args[1:]

    fmt.Println("Arguments:", args)
}
```

Run:

```bash
go run main.go hello world
```

Output:

```text
Arguments: [hello world]
```

Explanation:

- `os.Args[0]` is program name
- `os.Args[1:]` contains user arguments

---

## 15.4 Parse command-line number

### 15.4.1 Use `strconv.Atoi`

Example:

```go
package main

import (
    "fmt"
    "os"
    "strconv"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("usage: go run main.go <number>")
        return
    }

    n, err := strconv.Atoi(os.Args[1])
    if err != nil {
        fmt.Println("invalid number")
        return
    }

    fmt.Println("You entered:", n)
}
```

Explanation:

Converts command-line string to integer.

---

# 16. Basic testing

Layer 1 should include the smallest possible testing habit.

---

## 16.1 Test file naming

### 16.1.1 File must end with `_test.go`

Example:

```text
add.go
add_test.go
```

Explanation:

Go test files must end with `_test.go`.

---

## 16.2 Simple function

Example:

```go
package main

func Add(a int, b int) int {
    return a + b
}
```

---

## 16.3 Simple test

Example:

```go
package main

import "testing"

func TestAdd(t *testing.T) {
    got := Add(2, 3)
    want := 5

    if got != want {
        t.Fatalf("expected %d, got %d", want, got)
    }
}
```

Explanation:

- Test function starts with `Test`
- It receives `*testing.T`
- Use `t.Fatalf` or `t.Errorf` when test fails

---

## 16.4 Run tests

Example:

```bash
go test
```

Explanation:

Runs all tests in current package.

For verbose output:

```bash
go test -v
```

---

# 17. Debugging and common beginner compile errors

This should be part of Layer 1 study material.

---

## 17.1 Unused variable

Bad example:

```go
func main() {
    x := 10
}
```

Error:

```text
x declared and not used
```

Explanation:

Go does not allow unused local variables.

Fix:

```go
func main() {
    x := 10
    fmt.Println(x)
}
```

---

## 17.2 Unused import

Bad example:

```go
import "fmt"

func main() {
}
```

Error:

```text
imported and not used
```

Fix:

Remove unused import.

---

## 17.3 Undefined variable

Bad example:

```go
fmt.Println(message)
```

Error:

```text
undefined: message
```

Explanation:

Variable was not declared.

---

## 17.4 Wrong assignment operator

Bad example:

```go
x := 10
x := 20
```

Error:

```text
no new variables on left side of :=
```

Correct:

```go
x := 10
x = 20
```

Explanation:

Use `:=` for declaration.

Use `=` for assignment.

---

## 17.5 Type mismatch

Bad example:

```go
var age int = "30"
```

Error:

```text
cannot use "30" as type int
```

Correct:

```go
var age int = 30
```

---

## 17.6 Index out of range

Bad example:

```go
nums := []int{1, 2, 3}
fmt.Println(nums[3])
```

Runtime panic:

```text
index out of range
```

Explanation:

Valid indexes are:

```text
0, 1, 2
```

---

## 17.7 Printing debug values

### 17.7.1 Use `%v` and `%#v`

Example:

```go
user := struct {
    Name string
    Age  int
}{
    Name: "Alice",
    Age:  30,
}

fmt.Printf("%v\n", user)
fmt.Printf("%#v\n", user)
```

Explanation:

`%#v` prints more detailed Go syntax representation.

Useful for debugging.

---

# 18. Layer 1 practice exercises

Your study material should include exercises.

---

## Exercise 1: Hello program

Goal:

Practice program structure and printing.

Task:

Write a program that prints:

```text
Hello, Go!
```

Example solution:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, Go!")
}
```

---

## Exercise 2: Grade classifier

Goal:

Practice `if`, `else if`.

Task:

Given score, print grade.

Example:

```go
score := 85

if score >= 90 {
    fmt.Println("A")
} else if score >= 80 {
    fmt.Println("B")
} else if score >= 70 {
    fmt.Println("C")
} else {
    fmt.Println("F")
}
```

---

## Exercise 3: Sum slice

Goal:

Practice `for` and slices.

Task:

Sum all numbers in slice.

Example:

```go
nums := []int{1, 2, 3, 4}
total := 0

for _, n := range nums {
    total += n
}

fmt.Println(total)
```

---

## Exercise 4: Count words

Goal:

Practice maps and strings.

Task:

Count word frequencies.

Example:

```go
words := []string{"go", "go", "rust", "go"}

counts := map[string]int{}

for _, word := range words {
    counts[word]++
}

fmt.Println(counts)
```

Output:

```text
map[go:3 rust:1]
```

---

## Exercise 5: Simple calculator function

Goal:

Practice functions and errors.

Task:

Create `divide` function returning error on zero.

Example:

```go
func divide(a float64, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("division by zero")
    }
    return a / b, nil
}
```

---

## Exercise 6: Reverse slice

Goal:

Practice loops and slices.

Example:

```go
nums := []int{1, 2, 3, 4, 5}

for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
    nums[i], nums[j] = nums[j], nums[i]
}

fmt.Println(nums)
```

Output:

```text
[5 4 3 2 1]
```

---

## Exercise 7: Command-line greeting

Goal:

Practice `os.Args`.

Task:

Print greeting using first argument.

Example:

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("usage: go run main.go <name>")
        return
    }

    fmt.Println("Hello", os.Args[1])
}
```

---

# Layer 1 definition of done

You can consider Layer 1 complete when you can do all of these without notes:

```text
1. Create and run a Go program.
2. Use fmt.Println and fmt.Printf.
3. Declare variables using var and :=.
4. Explain zero values.
5. Use constants.
6. Convert between basic types.
7. Use arithmetic, comparison, and logical operators.
8. Write if / else if / else.
9. Write switch statements.
10. Write classic for loops.
11. Use range over slices and maps.
12. Use break and continue.
13. Write functions with parameters and return values.
14. Write functions returning error.
15. Use arrays, slices, and maps at a basic level.
16. Use basic string functions.
17. Use struct and pointer syntax at a basic level.
18. Create and import local packages.
19. Read simple input or command-line arguments.
20. Write one basic test.
```

---

# What NOT to include deeply in Layer 1

Avoid going too deep into these at this stage:

```text
slice internals
map internals
interface internals
goroutine scheduler
channels in depth
context in depth
memory escape analysis
garbage collection
pprof
concurrency patterns
system design
LLD/HLD
advanced testing patterns
performance optimization
```

You can mention them, but do not build deep material for them in Layer 1.

---

# Suggested Layer 1 study material structure

For each topic, create one note like this:

```text
Topic:
Why it matters:
Syntax:
Minimal example:
Common variations:
Common mistakes:
Small exercise:
Flashcards:
```

Example for `if`:

```text
Topic:
if statement

Why it matters:
Used for conditional execution.

Syntax:
if condition {
}

Minimal example:
if age >= 18 {
    fmt.Println("adult")
}

Common variations:
if/else
else if
if with short statement

Common mistakes:
Using non-boolean condition
Forgetting braces
Expecting ternary operator

Small exercise:
Write a program that prints "even" if number is divisible by 2.

Flashcards:
Q: Does Go allow if with initialization?
A: Yes, example: if x := len(s); x > 0 {}

Q: Does Go have ternary operator?
A: No.
```

---

# Best order to study Layer 1

I would recommend this order:

```text
1. Tooling and running programs
2. Program anatomy
3. Printing
4. Variables and constants
5. Basic types
6. Operators
7. if and switch
8. for loops
9. Functions
10. Errors basics
11. Arrays/slices/maps basics
12. Strings basics
13. Structs/pointers preview
14. Packages/modules
15. Input
16. Testing basics
17. Common errors
18. Exercises
```

---

# Final advice

Layer 1 should feel like:

```text
I can write simple Go code.
I can run it.
I can fix basic compile errors.
I can explain basic syntax.
```

It should not feel like:

```text
I know everything about Go.
```

If you make Layer 1 material well, later layers become much easier because you will not struggle with basic syntax while trying to learn deeper concepts.