package a

func main() { // want "a.go:3"
	f()
}

func f() {} // want "a.go:4 a.go:7"
