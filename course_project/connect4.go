package main

import (
	"fmt"
	"math/rand"
	"net"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"time"
)

var clear map[string]func() //create a map for storing clear funcs
var board [][]string
var col []int = make([]int, BOARD_WIDTH)

const PORT string = ":12345"
const BOARD_WIDTH int = 7
const BOARD_HEIGHT int = 6
const EMPTY_SPOT string = "_"
const PLAYER_ONE_COLOR string = "○"
const PLAYER_TWO_COLOR string = "◙"

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

	//initialize the connect 4 board
	for i := 0; i < BOARD_HEIGHT; i++ {
		row := make([]string, BOARD_WIDTH)

		for i := 0; i < len(row); i++ {
			row[i] = EMPTY_SPOT
		}
		board = append(board, row)
	}
}

func clearConsole() {
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}

func print(board [][]string) {
	for i := 0; i < len(board[0]); i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			fmt.Print(board[i][j] + " ")
		}
		fmt.Println()
	}
}

func drop(board [][]string, column int, player string) bool {
	if column < len(board[0]) && col[column] < len(board) {
		board[5-col[column]][column] = player
		col[column]++
		return true
	}
	return false
}

func areFourConnected(board [][]string, player string) bool {
	// horizontalCheck
	for j := 0; j < len(board[0])-3; j++ {
		for i := 0; i < len(board); i++ {
			if board[i][j] == player &&
				board[i][j+1] == player &&
				board[i][j+2] == player &&
				board[i][j+3] == player {
				return true
			}
		}
	}
	// verticalCheck
	for i := 0; i < len(board)-3; i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == player &&
				board[i+1][j] == player &&
				board[i+2][j] == player &&
				board[i+3][j] == player {
				return true
			}
		}
	}
	// ascendingDiagonalCheck
	for i := 3; i < len(board); i++ {
		for j := 0; j < len(board[0])-3; j++ {
			if board[i][j] == player &&
				board[i-1][j+1] == player &&
				board[i-2][j+2] == player &&
				board[i-3][j+3] == player {
				return true
			}
		}
	}
	// descendingDiagonalCheck
	for i := 3; i < len(board); i++ {
		for j := 3; j < len(board[0]); j++ {
			if board[i][j] == player &&
				board[i-1][j-1] == player &&
				board[i-2][j-2] == player &&
				board[i-3][j-3] == player {
				return true
			}
		}
	}
	return false
}

var playerToColor map[int]string = map[int]string{
   	1 : PLAYER_ONE_COLOR,
    -1 : PLAYER_TWO_COLOR,
}

func minimax(board[][] string, color int, depth int) (int, int){
	if areFourConnected(board, PLAYER_TWO_COLOR){
		// fmt.Println("danger noodle")
		return 1000-depth, -1
	} else if areFourConnected(board, PLAYER_ONE_COLOR){
		return -1000+depth, -1
	}else if depth == 5{
		return 0, -1
	}

	if(color == -1){
		value := -1000000
		bestMove := 0

		r := rand.Perm(7)

		for _, i := range r{
			if col[i] < 5 {
				drop(board, i, playerToColor[color])
				val, _ := minimax(board, -color, depth + 1)
				if(value < val){
					bestMove = i
					value = val
				}
				col[i]--
				board[5 - col[i]][i] = EMPTY_SPOT
			}
		}
		return value, bestMove
	}
	

	value := 1000000
	bestMove := 0

	for i := 0; i < len(col); i++{
		if col[i] < 5 {
			drop(board, i, playerToColor[color])
			val, _ := minimax(board, -color, depth + 1)
			if(value > val){
				bestMove = i
				value = val
			}
			col[i]--
			board[5 - col[i]][i] = EMPTY_SPOT
		}
	}
	return value, bestMove
}

func playAgainstAi(){
	humanColor := PLAYER_ONE_COLOR
	aiColor := PLAYER_TWO_COLOR
	waiting := false

	for !areFourConnected(board, humanColor) && !areFourConnected(board, aiColor) {

		clearConsole()
		print(board)

		if waiting {
			fmt.Println("waiting for oponent move...\n")
			_, bestMove := minimax(board, -1, 0)
			drop(board, bestMove, aiColor)
			waiting = false
		} else {
			for {
				fmt.Printf("Enter column to drop: ")

				var column int
				fmt.Scan(&column)

				if column >= len(board[0]) || column < 0 || !drop(board, column, humanColor) {
					fmt.Println("You cant place here! Try another column")
				} else {
					waiting = true
					break
				}
			}
		}
	}

	clearConsole()
	print(board)
	if areFourConnected(board, humanColor) {
		fmt.Println("You won!")
	} else {
		fmt.Println("You lost.")
	}
}

func main() {

	fmt.Println("Hello! Welcome to connect four CMD!\n" +
		"To connect to a friend ip: Enter [1]\n" +
		"To wait for friend to connect to you enter [2]\n" +
		"To play against AI, enter [3]")

	var option string
	fmt.Scan(&option)

	for !(option == "1" || option == "2" || option == "3") {
		fmt.Println("Unknown command! Try again:")
		fmt.Scan(&option)
	}

	if option == "3" {
		playAgainstAi()
		return
	}

	var conn net.Conn
	var color string
	var opponentColor string
	waiting := true

	if option == "1" {
		color = PLAYER_ONE_COLOR
		opponentColor = PLAYER_TWO_COLOR

		fmt.Println("Enter friends IP:")

		var ip string
		fmt.Scan(&ip)

		conn, _ = net.Dial("tcp", ip + PORT)
		waiting = false

	} else if option == "2" {
		color = PLAYER_TWO_COLOR
		opponentColor = PLAYER_ONE_COLOR

		ln, _ := net.Listen("tcp", PORT)

		fmt.Println("Waiting for a friend...")

		var err error
		conn, err = ln.Accept()

		if err == nil {
			fmt.Print("A friend has connected! The game will begin soon!")
			time.Sleep(1 * time.Second)
		}
	}

	for !areFourConnected(board, color) && !areFourConnected(board, opponentColor) {

		clearConsole()
		print(board)

		if waiting {
			fmt.Println("waiting for oponent move...\n")
			var message string
			fmt.Fscan(conn, &message)
			otherPlayerColumn, _ := strconv.Atoi(message)
			drop(board, otherPlayerColumn, opponentColor)
			waiting = false
		} else {
			for {
				fmt.Printf("Enter column to drop: ")

				var column int
				fmt.Scan(&column)

				if column >= len(board[0]) || column < 0 || !drop(board, column, color) {
					fmt.Println("You cant place here! Try another column")
				} else {
					fmt.Fprintf(conn, "%d\n", column)
					waiting = true
					break
				}
			}
		}
	}

	clearConsole()
	print(board)
	if areFourConnected(board, color) {
		fmt.Println("You won!")
	} else {
		fmt.Println("You lost.")
	}

}
