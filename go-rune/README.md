
# What is a rune

 A rune in Go represents a Unicode code point and is an alias for int32. This type of slice is often used for processing Unicode characters of strings.

# Character vs. rune

Each rune actually refers to only one Unicode code point, meaning one rune represents a single character.

Furthermore, Go does not have a char data type, so all variables initialized with a character would automatically
be typecasted into int32, which, in this case, represents a rune.

Here, it might seem like rune is the char alternative for Go, but that is not actually true. In Go, strings are
actually made up of sequences of bytes, not a sequence of rune.

In this regard, we can consider rune to be different from the usual char data type we see in other programming languages.
Even when we need to print multiple rune in a string, we must first typecast them into rune arrays.
