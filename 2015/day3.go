package main

import (
    "fmt"
    "os"
)

var p = fmt.Println
var pf = fmt.Printf

func main() {
    b, err := os.ReadFile("test.in")
    if err != nil {
        fmt.Println("Error reading input.in")
        os.Exit(1)
    }
    s := string(b)
    for _, r := range s {
        r
    }
}
