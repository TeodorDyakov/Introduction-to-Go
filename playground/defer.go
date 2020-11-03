package main
import "fmt"

 func main(){

 	//функциите се изпълняват в обратен ред на ивикване
 	for i := 0; i < 10; i++{
 		defer fmt.Println(i)
 	}
	
	for i := 0; i < 5; i++ {
		defer func() {
			fmt.Printf(" %v", i)
		}()
	}
	
	for i := 0; i < 5; i++ {
		defer func(i int) {
			fmt.Printf(" %v", i)
		}(i)
	}

 }