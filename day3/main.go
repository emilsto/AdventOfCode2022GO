package main

import (
	"bufio"
	"log"
	"os"
)

type compartment string

type Rucksack struct {
	first  compartment
	second compartment
}

func readFileToBuf(path string, buf *[]string) error {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		*buf = append(*buf, line)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil

}

func main() {
	var buf []string
	readFileToBuf("input.txt", &buf)

	scoremap := generateScoreMap()

	var score int = 0
	for i := 0; i < len(buf); i++ {
		sack := createRucksack(&buf[i])
		points := matchItems(&sack)
		score += (scoreMatches(&points, &scoremap))
	}

	print(score)

}

func generateScoreMap() map[string]int {

	scoremap := make(map[string]int)
	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for i := 0; i < len(letters); i++ {
		scoremap[string(letters[i])] = i + 1
	}

	return scoremap
}

func createRucksack(buf *string) Rucksack {
	var rucksack Rucksack
	// We know that all rucksacks are devisible by 2
	halfpoint := (len(*buf) / 2)
	rucksack.first = compartment((*buf)[0:halfpoint])
	rucksack.second = compartment((*buf)[halfpoint:])

	return rucksack
}

func matchItems(rucksack *Rucksack) map[string]int {
	match := make(map[string]int)

	for i := range rucksack.first {
		for j := range rucksack.second {
			if rucksack.first[i] == rucksack.second[j] {
				_, found := match[string(rucksack.first[i])]
				if !found {
					match[string(rucksack.first[i])] = i
				}
			}
		}
	}
	return match
}

func scoreMatches(match *map[string]int, scores *map[string]int) int {

	var result int
	for key := range *match {
		result = (*scores)[key]
	}

	return result

}
