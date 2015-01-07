package main

import (
        "fmt"
        "math"
)

// Use Newton's method n times.
func SqrtFixed(x float64, n int) float64 {
        z := float64(1)

        for i := 0; i < n; i++ {
                z -= (z*z - x) / 2 * z
        }

        return z
}

// Use Newton's method until delta is less than n.
func Sqrt(x, n float64) float64 {
        z := float64(1)
        iterations := 0
        delta := n

        for math.Abs(delta) >= n {
                iterations++
                delta = (z*z - x) / 2 * z
                z -= delta
        }

        fmt.Printf("iterations: %v\n", iterations)
        fmt.Printf("difference from math.Sqrt: %v\n", math.Sqrt(x)-z)

        return z
}

func main() {
        fmt.Println(SqrtFixed(2, 10))
        fmt.Println(Sqrt(2, 0.1))
}
