# Go Pets Package with Encapsulation

This implementation demonstrates Go's unique approach to encapsulation through a `Dog` struct with controlled state management. Unlike traditional OOP languages, Go achieves encapsulation through package-level visibility rules and constructor patterns.

---

## Key Encapsulation Features
- **Unexported State**: `lastSlept` field hidden from external packages [2][3]
- **Private Methods**: Internal `needsSleep()` and `sleep()` logic encapsulation [1][4]
- **Controlled Construction**: `NewDog()` factory enforces valid initialization [2][5]
- **State Validation**: Automatic sleep requirement checks during interactions [4]

---

## Encapsulation Breakdown

```
type Dog struct {
// Unexported fields (package-private)
lastSlept time.Time // 🛡️ Hidden state
}

// Unexported methods (internal logic)
func (d Dog) needsSleep() bool { ... } // ⚙️ State validation
func (d Dog) sleep() { ... } // 🔒 State mutation

// Exported API surface
func NewDog(...) Dog { ... } // 🏭 Safe construction
func (d Dog) PlayWithDog(...) string // 🎮 Public interaction

```

---

## Why This Matters?
1. **Data Integrity**: 
   - Internal sleep state cannot be modified directly [4][6]
   ```
   // pet.lastSlept = time.Now() // Compile error!
   ```

2. **Behavior Control**:
- Sleep requirements enforced through unexported method [3]

    ```
    // pet.needsSleep() // Compile error! [2]
    ```

3. **Evolution Safety**:
- Internal implementation can change without breaking consumers [5]

---


## OOP vs Go Encapsulation
| Feature          | Traditional OOP       | Go Approach             |
|-------------------|-----------------------|-------------------------|
| Access Modifiers  | `private/protected`   | Unexported identifiers [1][3] |
| Construction      | Constructors          | Factory functions [2]   |
| State Validation  | Setter methods        | Method encapsulation [4] |
| Inheritance       | Class hierarchy       | Composition [5][6]      |


---

## Practical Usage

```
// Mandatory proper initialization
sleepTime := time.Now().Add(-5*time.Hour)
pet := pets.NewDog("Buddy", "Brown", "Labrador", sleepTime)

// Public interaction methods
fmt.Println(pet.PlayWithDog("fetch")) // Auto-handles sleep state
```

---


## Best Practices (From Search Insights)
1. **Minimal Exports**: Only expose essential API surface [3][4]
2. **Value Receivers**: Use `(d Dog)` for immutable operations [5]
3. **Time Handling**: Leverage `time.Duration` for state checks [2]
4. **Documented Construction**: Clearly specify factory requirements [1]

---

## FAQ 
**Q: Why no classes in Go?**  
A: Go uses structs + packages for encapsulation rather than class hierarchies

**Q: How to handle inheritance?**  
A: Prefer composition over inheritance (see `NewDog` pattern)

**Q: What about protected visibility?**  
A: Go uses package boundaries - same package sees all identifiers

**Q: When to use pointer receivers?**  
A: When modifying internal state (not shown here, but common pattern)

---

## Encapsulation usage guidelines
- Maintain unexported fields/methods unless absolutely necessary
- Add new state through factory functions
- Preserve existing public method signatures


---
## Key Encapsulation Concepts

This package demonstrates three crucial encapsulation techniques:

1. **Unexported Fields**  
   `lastSlept` is private to the package (lowercase name)
```
type Dog struct {
lastSlept time.Time // inaccessible outside pets package
}
```

2. **Unexported Methods**  
Internal state management methods:
```
func (d Dog) needsSleep() bool { ... } // Private helper
func (d Dog) sleep() { ... } // Private state modifier
```

3. **Controlled Construction**  
Enforced via factory function:
```
func NewDog(...) Dog { ... } // Mandatory initialization
```

---

## Usage

### 1. Create Dogs Properly

```
// Mandatory: Use constructor with sleep time
pet := pets.NewDog(
"Bubzee",
"Black",
"Rottweiler",
time.Now().Add(-5*time.Hour), // Last slept 5 hrs ago
)
```


### 2. Interact Through Public API

```
fmt.Println(pet.Feed("Kibble")) // Public method
fmt.Println(pet.PlayWithDog("fetch")) // Public method
```

### 3. Forbidden Actions (Compile Errors)
```
pet.lastSlept = time.Now() // Error: unexported field
pet.needsSleep() // Error: unexported method
```

---

### Key Implementation Details

1. **Automatic Sleep Tracking**  
   The `lastSlept` field is automatically managed through the `sleep()` method when energy thresholds are reached.

2. **Immutable State**  
   Methods use value receivers to prevent accidental mutation:
```
func (d Dog) PlayWithDog(...) { ... } // Safe for concurrent use
```


3. **Time Calculations**  
Uses precise `time.Duration` checks:

```
time.Now().Sub(d.lastSlept) > 4*time.Hour
```

---

## Resources

- [1] Encapsulation and access control in Go - Mastering privacy with naming conventions: [https://codesignal.com/learn/courses/go-structs-basics-revision/lessons/go-encapsulation-and-access-control-mastering-privacy-with-naming-conventions](https://codesignal.com/learn/courses/go-structs-basics-revision/lessons/go-encapsulation-and-access-control-mastering-privacy-with-naming-conventions)
- [2] Encapsulation in Go - Safeguarding data with package-level visibility : [https://codesignal.com/learn/courses/clean-coding-with-go-structures/lessons/encapsulation-in-go-safeguarding-data-with-package-level-visibility](https://codesignal.com/learn/courses/clean-coding-with-go-structures/lessons/encapsulation-in-go-safeguarding-data-with-package-level-visibility)
- [3] Encapsulation in Go by Example: [https://golangbyexample.com/encapsulation-in-go/](https://golangbyexample.com/encapsulation-in-go/)
- [4] Object-Oriented Programming (OOP) in Go: [https://www.developer.com/languages/oop-go/](https://www.developer.com/languages/oop-go/)
- [5] Youtube video on Object Oriented Programming in Go by Jake Wright: [https://www.youtube.com/watch?v=dCI145NR-Fc](https://www.youtube.com/watch?v=dCI145NR-Fc)
- [6] Go Documentation: [https://go.dev/doc/](https://go.dev/doc/)