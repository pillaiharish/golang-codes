Answer:

```
prints 0
then panics
```

Because:
```
reading nil map → allowed
writing nil map → panic
```