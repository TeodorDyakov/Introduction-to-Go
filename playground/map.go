package main

import "fmt"

func main(){
    m := map[string]func(float64, float64)float64{
    "+" : func(a, b float64)float64{return a + b},
    "*" : func(a, b float64)float64{return a - b},
    "-" : func(a, b float64)float64{return a * b},
    "/" : func(a, b float64)float64{return a / b}}
    
    fmt.Println(m["+"](2.0, 3.0))
}
