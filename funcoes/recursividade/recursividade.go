package main

import "fmt"

func fatorial(n uint64) uint64 {
	if n == 0 {
		return 1
	}

	return n * fatorial(n-1)
}

func main() {
	resultado := fatorial(5)
	fmt.Println(resultado)

}
