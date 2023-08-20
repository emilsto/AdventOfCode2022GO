package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const (
	A int = 0
	B     = 1
	C     = 2
	X     = 3
	Y     = 4
	Z     = 5
)

func readFileToScanner(path string, buf *[]string) error {
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

	var buffer []string

	err := readFileToScanner("moves.txt", &buffer)

	if err != nil {
		log.Fatal("Error in reading file")
	}

	fmt.Println(len(buffer))

}
