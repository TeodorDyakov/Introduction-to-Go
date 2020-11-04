package main

import (
    "bufio"
    "fmt"
    "os"
)

func main(){
	var arr[30000]byte 
	var ptr, insPtr int
	reader := bufio.NewReader(os.Stdin)
	prog := []byte(os.Args[1])

	jump := func(step int){
		for brackets := 1; brackets != 0; {
			insPtr += step
			switch (prog[insPtr]) {
				case '[' : brackets += step
				case ']' : brackets -= step
			}	
		}			
	}
	for ; insPtr < len(prog); insPtr++ {
		switch prog[insPtr] {
			case '>' : ptr++
			case '<' : ptr--
			case '+' : arr[ptr]++
			case '-' : arr[ptr]--
			case '[' :
				if arr[ptr] == 0{
					jump(1)
				}
			case ']' :
				if arr[ptr] != 0{
					jump(-1)
				}
			case ',' : 
				input, _ := reader.ReadString('\n')
				arr[ptr] = ([]byte(input))[0]
			case '.' : fmt.Printf(string(arr[ptr]))
		}
	}
}