package main

func 𝝅() func() {
	defer func() { print(recover()) }()
	defer func() { panic(recover().(string) + "oops!") }()
	panic("oh no!")
	return func() {}
}

func main() {
	𝝅()()
}
