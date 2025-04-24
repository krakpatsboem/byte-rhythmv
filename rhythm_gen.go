package main

import (
    "fmt"
    "math/rand"
    "time"
)

func generatePattern(steps int) []int {
    rand.Seed(time.Now().UnixNano())
    pattern := make([]int, steps)
    for i := range pattern {
        if rand.Float32() > 0.5 {
            pattern[i] = 1
        } else {
            pattern[i] = 0
        }
    }
    return pattern
}

func main() {
    pattern := generatePattern(16)
    fmt.Println("Generated Rhythm Pattern:")
    fmt.Println(pattern)
}
