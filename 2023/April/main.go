package main

func foo() (b bool) {
	b = false
	return
}

func main() {
	switch foo(); {
	case false:
		println("False")
	case true:
		println("True")
	}
}
