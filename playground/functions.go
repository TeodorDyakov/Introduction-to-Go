package main

import "fmt"

func fac(n int) int{
	if n == 1{
		return 1
	}
	return n * fac(n-1)
}

func op(n string) func(int, int) int {
	if n == "+" {
		return func(a, b int) int{
			return a + b
		}
	}
	if n == "*"	{
		return func(a, b int) int{
			return a * b
		}
	}
	return nil
}

func seq() func() int{
	var n int = 0
	return func() int{
		n++
		return n
	}
}

func main() {
	var f = op("+")
	fmt.Println(f(2,3))
	//5
    fmt.Println(op("*")(2,3))
    //6
    var p = seq()
    fmt.Println(p())
    //1
    fmt.Println(p())
    //2
    fmt.Println(p())
    //3
    p = seq()
    fmt.Println(p())
    //1
    fmt.Println(p())
    //2
    fmt.Println(p())
    //3
}