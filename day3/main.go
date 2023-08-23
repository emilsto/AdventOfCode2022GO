package main

import (
	"bufio"
	"log"
	"os"
	"strings"
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

	// Im lazy
	scoremap := generateScoreMap()

	security_group := make([]Rucksack, 0, 100)

	var score int = 0
	var badge_score int = 0
	for i := 0; i < len(buf); i++ {
		sack := createRucksack(&buf[i])
		security_group = append(security_group, sack)
		matched_compartment_types := matchItems(&[]string{string(sack.first_compartment), string(sack.second_compartment)})
		score += (scoreMatches(&matched_compartment_types, &scoremap))

		// Create a new group every 3 rucksacks
		if (i+1)%3 == 0 {
			secretkey := getSecretKey(&security_group)
			println("Group: ", i/3+1, " has a secret key of :", secretkey)
			badge_score += (scoreMatches(&secretkey, &scoremap))
			// Reset the group
			security_group = nil
		}
	}

	println("Badge score: ", badge_score)
	print(score)

}

func getSecretKey(security_group *[]Rucksack) string {

	var contents []string
	for i := range *security_group {
		content := getRuckSackContent(&(*security_group)[i])
		contents = append(contents, content)
	}

	secretKey := matchItems(&contents)

	return secretKey
}

func getRuckSackContent(sack *Rucksack) string {
	content := string(sack.first_compartment) + string(sack.second_compartment)
	return content
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

func matchItems(str *[]string) string {
	if len(*str) == 0 {
		return ""
	}

	firstStr := (*str)[0]

	for _, char := range firstStr {
		common := true

		for i := 1; i < len(*str); i++ {
			if !strings.ContainsRune((*str)[i], char) {
				common = false
				break
			}
		}

		if common {
			return string(char)
		}
	}

	return ""
}

func scoreMatches(match *string, scores *map[string]int) int {

	return (*scores)[*match]

}
