package main

import "fmt"

func Filter(p func(int) bool) func(...int) []int {
	return func(nums ...int) []int {
		slice := []int{}
		for _, num := range nums {
			if p(num) {
				slice = append(slice, num)
			}
		}
		return slice
	}
}

func Mapper(f func(int) int) func(...int) []int {
	return func(nums ...int) []int {
		for idx, num := range nums {
			nums[idx] = f(num)
		}
		return nums
	}
}

func Reducer(initial int, f func(int, int) int) func(...int) int {
	init := initial
	return func(nums ...int) int {
		for _, num := range nums {
			init = f(init, num)
		}
		return init
	}
}

func MapReducer(initial int, mapper func(int) int, reducer func(int, int) int) func(...int) int {
	red := Reducer(initial, reducer)
	return func(nums ...int) int {
		return red(Mapper(mapper)(nums...)...)
	}
}

func main() {
	odds := Filter(func(x int) bool { return x%2 == 1 })
	evens := Filter(func(x int) bool { return x%2 == 0 })
	fmt.Println(odds(1, 2, 3, 4, 5), evens(1, 2, 3, 4, 5))
	double := Mapper(func(a int) int { return 2 * a })

	fmt.Println(double(1, 2, 3)) // [2, 4, 6]
	fmt.Println(double(4, 5, 6)) // [8, 10, 12]

	sum := Reducer(0, func(a, b int) int { return a + b })
	fmt.Println(sum(1, 2, 3))       // 6
	fmt.Println(sum(5))             // 11
	fmt.Println(sum(100, 101, 102)) // 314

	powerSum := MapReducer(
		0,
		func(v int) int { return v * v },
		func(a, v int) int { return a + v },
	)

	fmt.Println(powerSum(1, 2, 3, 4)) // 30
	fmt.Println(powerSum(1, 2, 3, 4)) // 60
	fmt.Println(powerSum())           // 60
}