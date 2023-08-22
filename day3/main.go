package main

import (
	"bufio"
	"log"
	"os"
)

type compartment string

type Rucksack struct {
	first_compartment  compartment
	second_compartment compartment
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
		mathched_compartment_types := matchItems(&sack)
		score += (scoreMatches(&mathched_compartment_types, &scoremap))
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
	// We know (belive) that all rucksacks are devisible by 2
	halfpoint := (len(*buf) / 2)
	rucksack.first_compartment = compartment((*buf)[:halfpoint])
	rucksack.second_compartment = compartment((*buf)[halfpoint:])

	return rucksack
}

func matchItems(rucksack *Rucksack) map[string]int {
	matched_types := make(map[string]int)

	// Making two maps would be a way to make this O(n) instead of O(n^2)
	for i := range rucksack.first_compartment {
		for j := range rucksack.second_compartment {
			if rucksack.first_compartment[i] == rucksack.second_compartment[j] {
				_, found := matched_types[string(rucksack.first_compartment[i])]
				if !found {
					matched_types[string(rucksack.first_compartment[i])] = i
				}
			}
		}
	}
	return matched_types
}

func scoreMatches(mathed_types *map[string]int, scores *map[string]int) int {

	var result int
	for key := range *mathed_types {
		result = (*scores)[key]
	}

	return result

}
