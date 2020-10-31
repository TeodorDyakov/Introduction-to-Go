package main
import "fmt"

//Масивите във Го са стойности, а не пойнтъри. Това значи че когато са аргумент на ф-я
//се предават по стойност, тоест се копират.
//операторът = копира сотйностите на масива. Те са типизирани от дължината на масива и типът на данните.
//Масивите се ползват рядко защото не са удобни за работа. Често вместо тя се ползва slice
func main(){
	
	var arr1[3] string = [3]string{"dog", "dog", "dog"}
	arr1[1] = "cat"
	fmt.Println(arr1)
	
	//ok
	arr1 = [3]string{"cat", "cat", "cat"}
	fmt.Println(arr1)
	
	arr2 := [3]int{1,2,3}
	fmt.Println(arr2)

	arr3 := [...]int{1,2,4}
	fmt.Println(arr3)

	fmt.Println(arr3 == arr2)
	//false - елементите са различни

	arr3 = arr2
	
	fmt.Println(&arr3 == &arr2)
	//false
	fmt.Println(arr3 == arr2)
	//true 

	//частично задаване на стойности, останалите се допълват до 0
	var y [5]int = [5]int{10, 20, 30}
	
	fmt.Println(y)
	fmt.Println(len(y))

}