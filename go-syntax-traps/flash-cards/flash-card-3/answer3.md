Answer:

```
[100 2 3]
```

The element mutation changes the shared backing array.
The append only changes the local slice header, and because this original slice likely has no spare capacity, it needs new backing storage.