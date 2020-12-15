package main

import (
	"fmt"
	"math/rand"
	"net"
	"os"
	"strconv"
	"time"
)

const(
	connHost = "localhost"
	connPort = "12345"
	connType = "tcp"
	BIG = 100000
	SMALL = -BIG
	PLAYER_ONE_COLOR = "○"
	PLAYER_TWO_COLOR = "◙"
	MIN_DIFFICULTY = 1
	MAX_DIFFICULTY = 7
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func minimax(board [][]string, maximizer bool, depth, max_depth int) (int, int) {
	if areFourConnected(board, PLAYER_TWO_COLOR) {
		return BIG - depth, -1
	} else if areFourConnected(board, PLAYER_ONE_COLOR) {
		return SMALL + depth, -1
	} else if depth == max_depth {
		return 0, -1
	}

	var value int
	var bestMove int
	shuffledColumns := rand.Perm(7)

	if maximizer {
		value = SMALL
		for _, i := range shuffledColumns {
			if drop(board, i, PLAYER_TWO_COLOR) {
				val, _ := minimax(board, false, depth + 1, max_depth)
				if value < val {
					bestMove = i
					value = val
				}
				// undo the move(backtracking)
				col[i]--
				board[5-col[i]][i] = EMPTY_SPOT
			}
		}
		return value, bestMove
	}else {
		value = BIG
		for _, i := range shuffledColumns {
			if drop(board, i, PLAYER_ONE_COLOR) {
				val, _ := minimax(board, true, depth + 1, max_depth)
				if value > val {
					bestMove = i
					value = val
				}
				//undo the move(backtracking)
				col[i]--
				board[5-col[i]][i] = EMPTY_SPOT
			}
		}
	}
	return value, bestMove
}

func playAgainstAi() {

	fmt.Printf("Choose difficulty (number between 1 and 7), %d - easy, %d - hard\n", MIN_DIFFICULTY, MAX_DIFFICULTY)
	var option string
	fmt.Scan(&option)
	
	difficulty, err := strconv.Atoi(option)

	for err != nil || difficulty < MIN_DIFFICULTY || difficulty > MAX_DIFFICULTY {
		fmt.Println("Invalid input! Try again:")
		fmt.Scan(&option)
		difficulty, err = strconv.Atoi(option)
	}

	humanColor := PLAYER_ONE_COLOR
	aiColor := PLAYER_TWO_COLOR
	waiting := false

	for !areFourConnected(board, humanColor) && !areFourConnected(board, aiColor) {

		clearConsole()
		printBoard(board)

		if waiting {
			fmt.Println("waiting for oponent move...\n")
			_, bestMove := minimax(board, true, 0, difficulty)
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
	printBoard(board)
	if areFourConnected(board, humanColor) {
		fmt.Println("You won!")
	} else {
		fmt.Println("You lost.")
	}
}

func playMultiplayer(){
	var conn net.Conn
	var color string
	var opponentColor string
	
	waiting := true

	fmt.Println("Connecting to", connType, "server", connHost+":"+connPort)
	conn, err := net.Dial(connType, connHost+":"+connPort)
	if err != nil {
		fmt.Println("Error connecting:", err.Error())
		os.Exit(1)
	}
	fmt.Println("searching for opponent...")

	var msg string
	fmt.Fscan(conn, &msg)
	// fmt.Println(msg)

	if msg == "go"{
		color = PLAYER_ONE_COLOR
		opponentColor = PLAYER_TWO_COLOR
		waiting = false
	}else{
		color = PLAYER_TWO_COLOR
		opponentColor = PLAYER_ONE_COLOR
		waiting = true
	}

	for !areFourConnected(board, color) && !areFourConnected(board, opponentColor) {

		clearConsole()
		printBoard(board)

		if waiting {
			fmt.Println("waiting for oponent move...\n")
			
			c1 := make(chan string, 1)
				
			go func(){
				var message string
				fmt.Fscan(conn, &message)
				c1 <- message	
			}()
			
			var colString string
			select {
		    case colString = <-c1:
		    	otherPlayerColumn, _ := strconv.Atoi(colString)
				drop(board, otherPlayerColumn, opponentColor)
				waiting = false
		    case <-time.After(60 * time.Second):
		        fmt.Println("timeout Opponent failed to make a move in 60 seconds")
		        return
		    }

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
	
	fmt.Fprintf(conn, "end")

	clearConsole()
	printBoard(board)
	if areFourConnected(board, color) {
		fmt.Println("You won!")
	} else {
		fmt.Println("You lost.")
	}
}

func main() {

	fmt.Println("Hello! Welcome to connect four CMD!\n" +
		"To enter multiplayer lobby press [1]\n" + "To play against AI press [2]\n")

	var option string
	fmt.Scan(&option)

	for !(option == "1" || option == "2") {
		fmt.Println("Unknown command! Try again:")
		fmt.Scan(&option)
	}

	if option == "2" {
		playAgainstAi()
		return
	} else{
		playMultiplayer()
	}

}
