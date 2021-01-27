package main

//  _________  _______   ________  ________
// |\___   ___\\  ___ \ |\   ___ \|\   __  \
// \|___ \  \_\ \   __/|\ \  \_|\ \ \  \|\  \
//      \ \  \ \ \  \_|/_\ \  \ \\ \ \  \\\  \
//       \ \  \ \ \  \_|\ \ \  \_\\ \ \  \\\  \
//        \ \__\ \ \_______\ \_______\ \_______\
//         \|__|  \|_______|\|_______|\|_______|
//
//                          Written on 19/01/2021

import (
	"fmt"
	"sync"
	"time"
)

type Task interface {
	Execute(int) (int, error)
}

func Pipeline(tasks ...Task) Task {
	return pipelineTask{tasks}
}

type pipelineTask struct {
	tasks []Task
}

func (pt pipelineTask) Execute(arg int) (res int, err error) {
	if pt.tasks == nil {
		return -1, fmt.Errorf("no tasks")
	}

	res = arg

	for _, task := range pt.tasks {
		res, err = task.Execute(res)
		if err != nil {
			return res, err
		}
	}

	return
}

func Fastest(tasks ...Task) Task {
	return fastestTask{tasks}
}

type fastestTask struct {
	tasks []Task
}

type returnVals struct {
	res int
	err error
}

func (ft fastestTask) Execute(arg int) (int, error) {
	if ft.tasks == nil {
		return -1, fmt.Errorf("no tasks")
	}

	var doOnce sync.Once
	c := make(chan returnVals)

	for _, task := range ft.tasks {
		go func(t Task) {
			res, err := t.Execute(arg)
			doOnce.Do(func() {
				c <- returnVals{res, err}
			})
			}(task)
	}

	taskReturnValues := <-c

	return taskReturnValues.res, taskReturnValues.err
}

type timedTask struct {
	task    Task
	timeout time.Duration
}

func (tt timedTask) Execute(arg int) (int, error) {
	// buffered channel so we dont leave the goroutine 'hanging' on send
	done := make(chan returnVals, 1)

	go func() {
		res, err := tt.task.Execute(arg)
		done <- returnVals{res, err}
	}()

	select {
	case taskReturnValues := <-done:
		return taskReturnValues.res, taskReturnValues.err
	case <-time.After(tt.timeout):
		return 0, fmt.Errorf("couldnt finish in given time")
	}
}

func Timed(task Task, timeout time.Duration) Task {
	return timedTask{task, timeout}
}

type mapReduceTask struct {
	tasks  []Task
	reduce func([]int) int
}

func (mrt mapReduceTask) Execute(arg int) (int, error) {
	if mrt.tasks == nil {
		return -1, fmt.Errorf("no tasks")
	}
	//bufffered channel so goroutines dont block on send
	c := make(chan returnVals, len(mrt.tasks))

	for _, task := range mrt.tasks {
		go func(t Task) {
			res, err := t.Execute(arg)
			c <- returnVals{res, err}
		}(task)
	}

	slice := make([]int, len(mrt.tasks))

	for idx, _ := range mrt.tasks {
		vals := <-c
		res, err := vals.res, vals.err
		if err != nil {
			return 0, err
		}
		slice[idx] = res
	}

	return mrt.reduce(slice), nil
}

func ConcurrentMapReduce(reduce func(results []int) int, tasks ...Task) Task {
	return mapReduceTask{tasks, reduce}
}

type greatestSearcherTask struct {
	errorLimit int
	tasks      <-chan Task
}

func (gst greatestSearcherTask) Execute(arg int) (int, error) {
	c := make(chan returnVals)
	var wg sync.WaitGroup

	go func() {
		for task := range gst.tasks {
			wg.Add(1)
			go func(t Task) {
				res, err := t.Execute(arg)
				c <- returnVals{res, err}
				wg.Done()
			}(task)
		}
		wg.Wait()
		close(c)
	}()

	//hacker's way to get min value of 32 bit int
	max := -1 << 31

	numberOfErrors := 0
	cnt := 0
	
	var err error 
	
	for retVal := range c {
		cnt++
		if retVal.err != nil {
			numberOfErrors++
		} else if max < retVal.res {
			max = retVal.res
		}
		if numberOfErrors > gst.errorLimit {
			max = -1
			err = fmt.Errorf("more errors than error limit")
		}
	}

	if cnt == 0 {
		return -1, fmt.Errorf("No tasks received!")
	}

	return max, err
}

func GreatestSearcher(errorLimit int, tasks <-chan Task) Task {
	return greatestSearcherTask{errorLimit, tasks}
}

type lazyAdder struct {
	adder
	delay time.Duration
}

func (la lazyAdder) Execute(addend int) (int, error) {
	time.Sleep(la.delay * time.Millisecond)
	return la.adder.Execute(addend)
}

type adder struct {
	augend int
}

func (a adder) Execute(addend int) (int, error) {
	result := a.augend + addend
	if result > 127 {
		return 0, fmt.Errorf("Result %d exceeds the adder threshold", a)
	}
	return result, nil
}

func main() {
	if res, err := Pipeline(adder{50}, adder{60}).Execute(10); err != nil {
		fmt.Printf("The pipeline returned an error\n")
	} else {
		fmt.Printf("The pipeline returned %d\n", res)
	}

	if res, err := Pipeline(adder{30}, adder{-50}).Execute(100); err != nil {
		fmt.Printf("The pipeline returned an error\n")
	} else {
		fmt.Printf("The pipeline returned %d\n", res)
	}

	f := Fastest(
		lazyAdder{adder{20}, 500},
		lazyAdder{adder{50}, 300},
		adder{41},
	)
	fmt.Println(f.Execute(1))

	_, e1 := Timed(lazyAdder{adder{20}, 50}, 2*time.Millisecond).Execute(2)
	r2, e2 := Timed(lazyAdder{adder{20}, 50}, 300*time.Millisecond).Execute(2)
	fmt.Println(e1, r2, e2)

	fmt.Println(Pipeline().Execute(10))

	reduce := func(results []int) int {
		smallest := 128
		for _, v := range results {
			if v < smallest {
				smallest = v
			}
		}
		return smallest
	}

	mr := ConcurrentMapReduce(reduce, adder{30}, adder{50}, adder{20})
	if res, err := mr.Execute(5); err != nil {
		fmt.Printf("We got an error!\n")
	} else {
		fmt.Printf("The ConcurrentMapReduce returned %d\n", res)
	}

	tasks := make(chan Task)
	gs := GreatestSearcher(2, tasks) // Приемаме 2 грешки

	go func() {
		tasks <- adder{4}
		tasks <- lazyAdder{adder{22}, 20}
		tasks <- adder{125} // Това е първата "допустима" грешка (защото 125+10 > 127)
		time.Sleep(50 * time.Millisecond)
		tasks <- adder{32} // Това би трябвало да "спечели"

		// Това би трябвало да timeout-не и да е втората "допустима" грешка
		tasks <- Timed(lazyAdder{adder{100}, 2000}, 20*time.Millisecond)

		// Ако разкоментираме това, gs.Execute() трябва да върне грешка
		// tasks <- adder{127} // трета (и недопустима) грешка
		close(tasks)
	}()
	result, err := gs.Execute(10)
	fmt.Println(result, err)
}
