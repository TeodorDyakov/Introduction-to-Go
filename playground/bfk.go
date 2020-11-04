package main

import (
    "bufio"
    "fmt"
    "os"
)

func main(){
	var arr[30000]byte 
	ptr, insPtr := 0, 0
	in := bufio.NewReader(os.Stdin)
	prog := []byte(os.Args[1])

	jump := func(step int, do bool){
		for brackets := 1; do && brackets != 0; {
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
			case '[' : jump(1, arr[ptr] == 0)
			case ']' : jump(-1, arr[ptr] != 0)
			case ',' : 
				input, _ := in.ReadString('\n')
				arr[ptr] = ([]byte(input))[0]
			case '.' : fmt.Printf(string(arr[ptr]))
		}
	}
}