package main

import (
	"bufio"
	"log"
	"os"
)

type Move byte

const (
	Win  Move = 'Z'
	Draw Move = 'Y'
	Lose Move = 'X'
)

// Complimentary moves to play by the cheatbook
const (
	RockWin  = PaperPlayer
	RockDraw = RockPlayer
	RockLose = ScissorsPlayer

	PaperWin  = ScissorsPlayer
	PaperDraw = PaperPlayer
	PaperLose = RockPlayer

	ScissorWin  = RockPlayer
	ScissorDraw = ScissorsPlayer
	ScissorLose = PaperPlayer
)

const (
	RockOpponent     Move = 'A'
	PaperOpponent    Move = 'B'
	ScissorsOpponent Move = 'C'
	RockPlayer       Move = 'X'
	PaperPlayer      Move = 'Y'
	ScissorsPlayer   Move = 'Z'
)

type Turn struct {
	playerMove   Move
	opponentMove Move
}

func readFileToBuf(path string, buf *[]Turn) error {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Bytes()

		// We trust that the moves are 1. Player and 2. Opponent and that
		// They are space separated
		if len(line) == 3 {
			oMove := Move(line[0])
			pMove := Move(line[2])
			turn := Turn{
				playerMove:   pMove,
				opponentMove: oMove,
			}
			*buf = append(*buf, turn)
		}
		// If the line is not len of 3, then it is fucked
		continue
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil

}

func main() {

	var turns []Turn

	err := readFileToBuf("moves.txt", &turns)

	if err != nil {
		log.Fatal("Error in reading file")
	}

	var totalPointage int16 = 0

	for i := range turns {
		//cheatMove(&turns[i])
		totalPointage += getPoints(&turns[i])
	}

	println("Total pointage :", totalPointage)

}

func getPoints(turn *Turn) int16 {
	var baseLinePoints int16 = 0
	switch turn.playerMove {
	case RockPlayer:
		baseLinePoints += 1
		if turn.opponentMove == RockOpponent {
			baseLinePoints += 3
		} else if turn.opponentMove == ScissorsOpponent {
			baseLinePoints += 6
		} else {
			baseLinePoints += 0
		}
		return baseLinePoints
	case PaperPlayer:
		baseLinePoints += 2
		if turn.opponentMove == PaperOpponent {
			baseLinePoints += 3
		} else if turn.opponentMove == RockOpponent {
			baseLinePoints += 6
		} else {
			baseLinePoints += 0
		}
	case ScissorsPlayer:
		baseLinePoints += 3
		if turn.opponentMove == ScissorsOpponent {
			baseLinePoints += 3
		} else if turn.opponentMove == PaperOpponent {
			baseLinePoints += 6
		} else {
			baseLinePoints += 0
		}
	}
	return baseLinePoints
}

func cheatMove(turn *Turn) {
	switch turn.playerMove {
	case Win:
		if turn.opponentMove == RockOpponent {
			// This is excplicitly dereferrencing the turn, but Go does that automagically
			// So there is no need to do it. Fun!
			(*turn).playerMove = RockWin
		}
		if turn.opponentMove == PaperOpponent {
			turn.playerMove = PaperWin
		}
		if turn.opponentMove == ScissorsOpponent {
			turn.playerMove = ScissorWin
		}
	case Draw:
		if turn.opponentMove == RockOpponent {
			turn.playerMove = RockDraw
		}
		if turn.opponentMove == PaperOpponent {
			turn.playerMove = PaperDraw
		}
		if turn.opponentMove == ScissorsOpponent {
			turn.playerMove = ScissorDraw
		}

	case Lose:
		if turn.opponentMove == RockOpponent {
			turn.playerMove = RockLose
		}
		if turn.opponentMove == PaperOpponent {
			turn.playerMove = PaperLose
		}
		if turn.opponentMove == ScissorsOpponent {
			turn.playerMove = ScissorLose
		}

	}
}
