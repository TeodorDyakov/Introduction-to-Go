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

var n int = 5

func global() func() int{
	return func() int{
		return n
	}
}

func main() {
	var f func(int, int) int = op("+")
	fmt.Println(f(2,3))
	//5
    fmt.Println(op("*")(2,3))
    //6
    
    var p func() int = seq()
    fmt.Println(p())
    //1
    fmt.Println(p())
    //2
    
    p = seq()
    fmt.Println(p())
    //1
    fmt.Println(p())
    //2
  	x := p
  	fmt.Println(x())
    //3  

    g := global()
    fmt.Println(g())
    //5
    n = 7
    fmt.Println(g())
    //7
}




