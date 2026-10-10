# defer, panic and recover examples

These three features work together to handle **cleanup** and **truly exceptional, unrecoverable situations**. 

> ⚠️ **Golden Rule**: Do not use `panic` and `recover` as a replacement for standard `error` returns. Use them only for catastrophic failures that the program cannot reasonably handle (e.g., a web server catching a crash in one request so it doesn't take down the entire server).

---

### 1. `defer`: The Cleanup Crew
The `defer` keyword schedules a function call to be executed **right before the surrounding function returns**, no matter how it returns (whether it finishes normally, returns early, or panics).

**Common uses**: Closing files, unlocking mutexes, closing database connections, or printing final logs.

```go
package main

import "fmt"

func main() {
    fmt.Println("1. Starting function")
    
    // This is scheduled to run LAST, just before main() exits.
    defer fmt.Println("3. Deferred: Cleaning up resources")
    
    fmt.Println("2. Doing some work")
    
    // The deferred function runs automatically here, at the end of main().
}
```
**Output:**
```text
1. Starting function
2. Doing some work
3. Deferred: Cleaning up resources
```

#### Two Crucial Rules of `defer`:
1. **LIFO Order**: If you have multiple `defer` statements, they execute in **Last-In, First-Out** order (like a stack).
   ```go
   defer fmt.Println("First")
   defer fmt.Println("Second")
   // Output: "Second", then "First"
   ```
2. **Arguments are evaluated immediately**: The *values* passed to a deferred function are locked in at the moment `defer` is called, not when the function actually executes.
   ```go
   i := 1
   defer fmt.Println("Deferred value:", i) // Locks in the value 1
   i = 2
   // Output: "Deferred value: 1"
   ```

---

### 2. `panic`: The Emergency Brake
A `panic` immediately stops the normal execution of the current function. It begins "unwinding" the call stack, running any deferred functions along the way, until the program crashes and prints a stack trace.

**When to use it**: Only when the program is in an unrecoverable state (e.g., a critical configuration file is missing, or a database connection fails at startup). 

```go
package main

import "os"

func main() {
    file, err := os.Open("non_existent_file.txt")
    if err != nil {
        // We don't know how to handle this, so we fail fast.
        panic("CRITICAL: Cannot start without the config file!")
    }
    defer file.Close()
    
    // This line will NEVER be reached if the file doesn't exist.
    println("File opened successfully")
}
```

---

### 3. `recover`: The Safety Net
If a `panic` occurs, the program will crash. However, `recover` allows you to intercept that panic, stop the crash, and regain control. 

**The Absolute Rule of `recover`**: It **only** works when called *inside* a `defer` function. If you call `recover()` in normal execution flow, it does nothing and returns `nil`.

---

### 4. Putting It All Together: The Classic Pattern
Here is how `defer`, `panic`, and `recover` work together in a real-world scenario. Imagine a math function that should never divide by zero, but we want to catch the mistake gracefully instead of crashing the whole app.

```go
package main

import "fmt"

// safeDivision returns an int. We name the return variable "result" 
// so we can modify it inside the defer block.
func safeDivision(a, b int) (result int) {
    
    // 1. DEFER: This anonymous function will run when safeDivision returns.
    defer func() {
        // 3. RECOVER: Catches the panic if one occurred.
        // 'r' will hold the value passed to panic() (e.g., "cannot divide by zero")
        if r := recover(); r != nil {
            fmt.Println("⚠️ Recovered from panic:", r)
            
            // We can set a safe default value before the function exits
            result = 0 
        }
    }()

    // 2. PANIC: If something truly invalid happens, trigger the emergency brake.
    if b == 0 {
        panic("cannot divide by zero")
    }

    // Normal execution path
    return a / b
}

func main() {
    fmt.Println("--- Test 1: Normal Execution ---")
    ans1 := safeDivision(10, 2)
    fmt.Println("Result:", ans1)

    fmt.Println("\n--- Test 2: Panic and Recover ---")
    ans2 := safeDivision(10, 0)
    fmt.Println("Result after recovery:", ans2)

    fmt.Println("\n--- Program continues normally! ---")
}
```

**Output:**
```text
--- Test 1: Normal Execution ---
Result: 5

--- Test 2: Panic and Recover ---
⚠️ Recovered from panic: cannot divide by zero
Result after recovery: 0

--- Program continues normally! ---
```

### How the Execution Flow Works in Test 2:
1. `safeDivision(10, 0)` is called.
2. The `defer` statement is registered.
3. `b == 0` is true, so `panic("cannot divide by zero")` is triggered.
4. Normal execution stops immediately. The `return a / b` line is **skipped**.
5. Go looks for deferred functions. It finds the anonymous `defer func()`.
6. Inside the defer, `recover()` catches the panic. `r` becomes the string `"cannot divide by zero"`.
7. Because `r != nil`, we print the warning and set the named return variable `result = 0`.
8. The function exits gracefully, returning `0`.
9. `main()` continues to the next line as if nothing catastrophic happened.

### Summary Cheat Sheet
| Keyword | Purpose | Analogy |
| :--- | :--- | :--- |
| **`defer`** | Schedules a function to run at the end of the current function. | "I'll take out the trash *right before* I leave the house." |
| **`panic`** | Aborts normal execution and starts unwinding the stack. | Pulling the emergency brake on a train. |
| **`recover`** | Catches a panic inside a `defer` block and stops the crash. | The airbag deploying when the emergency brake is pulled. |

