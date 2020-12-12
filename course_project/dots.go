package main

import(
	"fmt"
    "os"
    "os/exec"
 	"runtime"
 	"net"
 	"bufio"
 	"strconv"
 )

var clear map[string]func() //create a map for storing clear funcs

func init() {
    clear = make(map[string]func()) //Initialize it
    clear["linux"] = func() { 
        cmd := exec.Command("clear") //Linux example, its tested
        cmd.Stdout = os.Stdout
        cmd.Run()
    }
    clear["windows"] = func() {
        cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested 
        cmd.Stdout = os.Stdout
        cmd.Run()
    }
}

func CallClear() {
    value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
    if ok { //if we defined a clear func for that platform:
        value()  //we execute it
    } else { //unsupported platform
        panic("Your platform is unsupported! I can't clear terminal screen :(")
    }
}

func print(arr[][] string){
	for i := 0; i < 6; i++ {
		for j := 0; j < 7; j++ {
			fmt.Print(arr[i][j]+" ")
		}
		fmt.Println()
	}
}

var col[]int = make([]int, 7)

func drop(arr[][] string, column int, player string)(bool){
	if column < 7 && col[column] < 6{
		arr[5 - col[column]][column] = player
		col[column]++
		return true;
	 }
	 return false
}

func areFourConnected(board[][] string, player string)(bool){
    // horizontalCheck 
    for j := 0; j< 7 - 3 ; j++{
        for i := 0; i < 6; i++{
            if (board[i][j] == player && board[i][j+1] == player && board[i][j+2] == player && board[i][j+3] == player){
                return true;
            }           
        }
    }
    // verticalCheck
    for i := 0; i < 6-3 ; i++  {
        for j := 0; j < 7; j++ {
            if (board[i][j] == player && board[i+1][j] == player && board[i+2][j] == player && board[i+3][j] == player){
                return true;
            }           
        }
    }
    // ascendingDiagonalCheck 
    for i := 3; i < 6; i++ {
        for j := 0; j < 7 - 3; j++ {
            if board[i][j] == player && board[i-1][j+1] == player && board[i-2][j+2] == player && board[i-3][j+3] == player{
                return true;
            }
        }
    }
    // descendingDiagonalCheck
    for i := 3; i < 6; i++ {
        for j := 3; j < 7; j++ {
            if (board[i][j] == player && board[i-1][j-1] == player && board[i-2][j-2] == player && board[i-3][j-3] == player){
                return true;
            }
        }
    }
    return false;
}


func main(){

	var conn net.Conn

	fmt.Println("Hello! Wlcome to connect four CMD!\nTo connect to a friend ip: Enter [1]\n To wait for friend to connect to you enter [2]\n")

	var option int
	fmt.Scan(&option)

	if option == 1{
		fmt.Println("Enter friends IP:\n")
		reader := bufio.NewReader(os.Stdin)
		ip, _ := reader.ReadString('\n')
		conn, _ = net.Dial("tcp", ip[:len(ip) - 1]+":12345")

	}
	if option == 2 {
		ln, _ := net.Listen("tcp", ":12345")

		var err error
 		conn, err = ln.Accept()

 		if err == nil {
 			fmt.Print("A frined has connected!")
 		}
	}


 	board := [][]string{}

    for i := 0; i < 6; i++ {
    	row := make([]string, 7)

    	for i := 0; i < 7; i++{
			row[i] = "_"
		}
    	board = append(board, row)
    }

	var	player string = "○"
	turn := 0
	
	waiting := true

	for !areFourConnected(board, player){
		
		if turn % 2 == 0 {
			player = "○"
		} else {
			player = "◙"
		}

		var message string

		if waiting{
		    message, _ = bufio.NewReader(conn).ReadString('\n')	
		}

		fmt.Println("message received:" + message + "hui")

		otherPlayerColumn, _ := strconv.Atoi(message[:len(message)-1])

		// fmt.Printf("to int:%d", c)

		drop(board, otherPlayerColumn, "*")

		print(board)

		for ;; {
			fmt.Printf("Player %d, enter column to drop: ", turn % 2 + 1)
			column := 0

			fmt.Scan(&column)

			if column > 6 || column < 0 || !drop(board, column, player) {
				fmt.Println("You cant place here! Try another column")
			}else {
				break;
			}			
		}

		turn++
		CallClear()
	}

	print(board)
	fmt.Println("Winner is:" + player)

}