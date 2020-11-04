package main

import (
    "bufio"
    "fmt"
    "os"
)

func main(){
	var mem[30000]byte 
	ptr, insPtr := 0, 0
	in := bufio.NewReader(os.Stdin)
	code := []byte(os.Args[1])

	jump := func(step int){
		for brackets := 1; brackets != 0; {
			insPtr += step
			switch (code[insPtr]) {
				case '[' : brackets += step
				case ']' : brackets -= step
			}	
		}		
	}

	for ; insPtr < len(code); insPtr++ {
		switch code[insPtr] {
			case '>' : ptr++
			case '<' : ptr--
			case '+' : mem[ptr]++
			case '-' : mem[ptr]--
			case '[' : if mem[ptr] == 0 {jump(1)}
			case ']' : if mem[ptr] != 0 {jump(-1)}
			case ',' : input, _ := in.ReadString('\n')
				mem[ptr] = ([]byte(input))[0]
			case '.' : fmt.Printf(string(mem[ptr]))
		}
	}
}