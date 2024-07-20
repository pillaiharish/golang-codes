# Gin Binding
Gin binding is a de-serialization library. 

It supports JSON, XML, query parameter, and more out of the box and comes with a built-in validation framework.

Gin bindings are used to serialize JSON, XML, path parameters, form data, etc. to structs and maps. 

Gin supports various formats by providing struct tags. 

BindJSON reads the body buffer to de-serialize it to a struct. 

BindJSON cannot be called on the same context twice because it flushes the body buffer.

If we want to de-serialize the body to two different structs, use ShouldBindBodyWith to copy the body buffer and add it to context.


```bash
harish $ go run main.go
[GIN-debug] Listening and serving HTTP on :3000
Entire Request:  map[city:Bengaluru name:Harish]
Only body structure:  {Harish}
```