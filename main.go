package main

import "fmt"

var a []int

func fib() func(int) int {
	return func(x int) int {
		if x >= 2 {
			xf := a[x-1] + a[x-2]
			a = append(a, xf)
			return xf
		}

		a = append(a, x)
		return (a[x])
	}
}

func main() {
	fmt.Println("Hi, this is a dummy project for testing GPG signing in GH :)")

	fmt.Println("Here you have some fibonacci :D\n")
	f := fib()
	for i := 0; i < 25; i++ {
		fmt.Println(i, "-", f(i))
	}
}
