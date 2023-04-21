# April Gophers 2023

## What will this program do?

1. Panic
2. Will Not Compile
3. Print "True"
4. Print "False"
---
```go
package main

func foo() (b bool) {
    b = false
    return
}

func main() {
    switch foo()
    {
        case false: println("False")
        case true: println("True")
    }
}
```
---

## Answer

1. ~~Panic~~
2. ~~Will Not Compile~~
3. **Print "True"**
4. ~~Print "False"~~

Take another look at the switch statement.

```go
    switch foo()
    {
        case false: println("False")
        case true: println("True")
    }
```

It's missing a switch expression!
```go
    switch foo() ; /* a missing switch expression! */ {

        case false: println("False")
        case true: println("True")
    }
```

"The switch expression may be preceded by a simple statement, which executes before the expression is evaluated."

"A missing switch expression is equivalent to the boolean value true."