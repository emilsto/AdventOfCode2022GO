package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {

	cals := make([]int, 0)
	highBatches := make([]int, 0, 3)
	highBatches = append(highBatches, 0, 0, 0)
	high := 0

	file, err := os.Open("calories.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() != "" {
			val, err := strconv.Atoi(scanner.Text())
			if err != nil {
				log.Println("Failure to convert to int: ", scanner.Text())
				// no reason to continue if failure
				panic("exiting..")
			}
			cals = append(cals, val)
		}
		if scanner.Text() == "" {
			temp := 0
			for _, val := range cals {
				temp += val
			}
			if temp > high {
				high = temp
			}
			highBatches = toOrder(highBatches)
			for i := len(highBatches) - 1; i > 0; i-- {
				if highBatches[i] < temp {
					println("appending new high batch", temp, highBatches[i])
					highBatches[i] = temp
					break
				}
			}
			cals = nil
		}
	}

	println(high)
	printSlice(&highBatches)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func printSlice(slice *[]int) {
	var sum = 0
	for i, v := range *slice {
		println(i, v)
		sum += v
	}
	println(sum)
}

func toOrder(slice []int) []int {
	for i := range slice {
		for j := range slice {
			if slice[i] > slice[j] {
				swap(&slice[i], &slice[j])
			}
		}
	}
	return slice
}

func swap(a *int, b *int) {
	var temp int = *a
	*a = *b
	*b = temp
}
