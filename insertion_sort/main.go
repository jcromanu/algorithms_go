package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//array := []int{1, 5, 6, 664, 53322, 1}
	result := InsertionSort(ReadIntArray())
	fmt.Println(result)
}

func ReadIntArray() []int {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSuffix(text, "\n")
	array := strings.Split(text, " ")
	intArray := make([]int, len(array))
	for i, val := range array {
		if inte, err := strconv.Atoi(val); err == nil {
			intArray[i] = inte
		} else {
			panic(err)
		}
	}
	return intArray
}

func InsertionSort(array []int) []int {
	for i, _ := range array {
		j := i
		for j > 0 {
			if array[j] < array[j-1] {
				temporal := array[j]
				array[j] = array[j-1]
				array[j-1] = temporal
			}
			j--
		}
	}
	return array
}
