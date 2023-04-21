# March Gophers 2023

## What will this program do?

1. 2
2. 0
3. Code Doesn't Compile
4. Code Panics
5. Last Number Will Be Random
---
```go
func main() {
	var i int
	defer fmt.Println(i)

	var wg sync.WaitGroup
	for i < 2 {
		wg.Add(1)
		go func() {
			defer wg.Done()

			i++
			fmt.Println(i)
		}()
	}
	wg.Wait()
}
```
---
## Answer

1. ~~2~~
2. 0
3. ~~Code Doesn't Compile~~
4. ~~Code Panics~~
5. ~~Last Number Will Be Random~~

In this puzzle, this entire section...

```go
	var wg sync.WaitGroup
	for i < 2 {
		wg.Add(1)
		go func() {
			defer wg.Done()

			i++
			fmt.Println(i)
		}()
	}
	wg.Wait()
```
Is a red herring. What's actually important is the first two lines
```go
var i int
defer fmt.Println(i)
```

"The defer function evaluates the value of its parameters at the function call, even though it only gets 
invoked after the function returns!"

"This is why the correct answer to the question is 0"

~Submitted by Berni Varga
https://github.com/BerniVarga/go-ing-well/wiki/Go-Conundrum-2

