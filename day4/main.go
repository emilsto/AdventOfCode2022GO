package main

import (
	"bufio"
	"os"
)

type shift []byte

type shiftData struct {
	shift_start shift
	shift_end   shift
}

const (
	seperator     = ','
	shift_breaker = '-'
)

type elfPair struct {
	elf_one_shift shiftData
	elf_two_shift shiftData
}

func main() {
	var buffer []shift
	readFileToBuf("input.txt", &buffer)

	// Assign elfpairs
	elf_pair := generateElfPair(buffer[0])

	// Check overlap

}

func generateElfPair(line_of_shifts shift) *elfPair {

	var pair = new(elfPair)

	indexOfSeperator := 0
	indexOfFirstBreaker := 0
	indexOfSecondBreaker := 0

	println(string(line_of_shifts))

	for i, byt := range line_of_shifts {
		if byt == seperator {
			indexOfSeperator = i
		}
		if byt == shift_breaker && indexOfFirstBreaker == 0 {
			indexOfFirstBreaker = i
		}
		if indexOfFirstBreaker != 0 && byt == shift_breaker {
			indexOfSecondBreaker = i
		}
	}

	// Assign elf one shift
	pair.elf_one_shift.shift_start = line_of_shifts[:indexOfFirstBreaker]
	pair.elf_one_shift.shift_end = line_of_shifts[indexOfFirstBreaker+1 : indexOfSeperator]

	// Assign elf two shift
	pair.elf_two_shift.shift_start = line_of_shifts[indexOfSeperator+1 : indexOfSecondBreaker]
	pair.elf_two_shift.shift_end = line_of_shifts[indexOfSecondBreaker+1:]

	println(string(pair.elf_one_shift.shift_start))
	println(string(pair.elf_one_shift.shift_end))

	println(string(pair.elf_two_shift.shift_start))
	println(string(pair.elf_two_shift.shift_end))

	return pair
}

func checkOverlap(shifted_elfpair elfPair) {

}

func readFileToBuf(path string, buf *[]shift) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Bytes()
		*buf = append(*buf, line)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
