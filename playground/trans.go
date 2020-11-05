package main

import (
    "fmt"
    "os"
)

func main(){
	code := []byte(os.Args[1])

	res := "package main\nimport (\n\"bufio\"\n\"fmt\"\n\"os\")\nfunc main(){\nvar mem[30000]byte\nptr := 0\nin := bufio.NewReader(os.Stdin)\n_=in\n"
	
	m :=map[byte]string{
	'>' : "ptr++",
	'<' : "ptr--",
	'+' : "mem[ptr]++",
	'-' : "mem[ptr]--",
	'[' : "for mem[ptr] != 0 {",
	']' : "}",
	',' : "input, _ := in.ReadString('\\n')\nmem[ptr] = ([]byte(input))[0]",
	'.' : "fmt.Printf(string(mem[ptr]))"}

	for _, v := range code {
		res += m[v]+"\n"
	}
	fmt.Println(res+"}")
}