package main
import "fmt"
import "math"
var n int = 4;
var b = 1;
func square(){
	n++;
	n*=b;
}

func f(){
	defer square()
	b = 2;
	fmt.Println("value of n before end of f: ", n)
}

 func main(){
 	f();
 	fmt.Println("value of n after f: ", n)
 	fmt.Println(real(3i+4))
 	//функциите се изпълняват в обратен ред на ивикване
 	for i := 0; i < 10; i++{
 		defer fmt.Println(i)
 	}
 }