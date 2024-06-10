//Fibonacci in Go
package main

import "fmt"

// fib function calculates the nth Fibonacci number
func fib(n int) int {
    if n == 0 {
        return 0
    }
    a, b := 0, 1
    for i := 2; i <= n; i++ {
        a, b = b, a + b
    }
    return b
}

func main() {
    n := 9
    fmt.Println(fib(n))
}

//Ali Barzegari Dahaj
