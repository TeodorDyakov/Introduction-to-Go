package main

import(
	"math/rand"
)

func alphabeta(b *Board, maximizer bool, depth, alpha, beta, max_depth int) (int, int) {
	if b.areFourConnected(PLAYER_TWO_COLOR) {
		return BIG - depth, - 1
	} else if b.areFourConnected(PLAYER_ONE_COLOR) {
		return SMALL + depth, - 1
	} else if depth == max_depth {
		return 0, -1
	}

	var value int
	var bestMove int
	shuffledColumns := rand.Perm(7)

	if maximizer {
		value = SMALL
		for _, column := range shuffledColumns {
			if b.drop(column, PLAYER_TWO_COLOR) {
				val, _ := alphabeta(b, false, depth + 1, alpha, beta, max_depth)
				if value < val {
					bestMove = column
					value = val
				}
				b.undoDrop(column)
				if(alpha < value){
					alpha = val
				}
				if(alpha >= beta){
					break;
				}
			}
		}
	} else {
		value = BIG
		for _, column := range shuffledColumns {
			if b.drop(column, PLAYER_ONE_COLOR) {
				val, _ := alphabeta(b, true, depth + 1, alpha, beta, max_depth)
				if value > val {
					bestMove = column
					value = val
				}
				b.undoDrop(column)
				if(beta > value){
					beta = val
				}
				if(alpha >= beta){
					break;
				}
			}
		}
	}
	return value, bestMove
}