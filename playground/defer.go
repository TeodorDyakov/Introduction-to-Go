package main
import "fmt"

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
 }