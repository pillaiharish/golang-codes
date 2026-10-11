Before any interviews, I want these rules almost automatic for me:
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


# The mental models matter more than memorizing syntax
For slices, always draw:
```
slice
┌─────────────┐
│ ptr         │───────► backing array
│ len         │
│ cap         │
└─────────────┘
```

# For interfaces, always draw:
```
interface
 ┌───────────────┐
 │ dynamic type  │
 │ dynamic value │
 └───────────────┘
```

These two pictures will solve a surprising number of tricky Go questions.