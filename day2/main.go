package main

import (
	"bufio"
	"log"
	"os"
)

type Move byte

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
		totalPointage += getPoints(turns[i])
	}

	println("Total pointage :", totalPointage)

}

func getPoints(turn Turn) int16 {
	var baseLinePoints int16 = 0
	switch turn.playerMove {
	case RockPlayer:
		baseLinePoints += 1
		// Draw case
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
