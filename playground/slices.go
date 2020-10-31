package main

import "fmt"

func avg(arr []float64)(avg float64){
	for _, v := range arr{
		avg += v
	}
	avg /= float64(len(arr))
	return
}
 
 //пресмята moving average - res[i] е средното на arr[i: + n]
func movingAvg(arr []float64,  n int)(res []float64){

	res = make([]float64, len(arr) - n + 1)

	for i := 0; i <= len(arr) - n; i++ {
		res[i] = avg(arr[i: i + n])
	}

	return res
}

func main() {

    s := make([]string, 3)
    fmt.Println(len(s))
   	//3 
    s[2] = "cat"
    s[0] = "rat"
    s[1] = "dog"

    s = append(s, "pig")
    s = append(s, "fish")
    fmt.Println(s)
    fmt.Println(len(s))


    fmt.Println(s[4:],' ', s[:3], s[:])

    //5

  	s = s[2:3]
    fmt.Println(s)

    //странно ? Може би
    s = s[0:3]
    fmt.Println(s)

    f :=[...]float64{5.0, 2.0, 10.0, 16.0, 2.0, 5.0}
    //fmt.Println(avg(f)) - не работи
    //Защо ? Заради несъвместими типове: f е от тип [6]int, а пък функцията е с параметър []int -
    // за да работи какво правим? Ползваме slice
    // слайсове се правят с полуотворен интервал f[l, r] - f[r] - влиза в слайса, f[r] - не
    slice := f[0:6]
    fmt.Println(slice, '\n', avg(slice))
    fmt.Println(movingAvg(slice, 3))
    _ = append(f[1:3], 9999.0)
    slice = append(f[:], 666.0)
    fmt.Println(f)
    fmt.Println(slice)
    slice[0] = 420.0
    fmt.Println(slice)
    fmt.Println(f[0])
    }

    //слайсовете са готин feature във Go. Те предствляват тип който всъшност е нещо като пойнтър към масив и има дължина
    //Докато писах на Java съм се замислял че, нещо подобно ня тях би улеснилo пресмятанията със масиви,
    // които често са досадни или неудобни.