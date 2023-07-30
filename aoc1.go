package main

import (
	"fmt"
	"strconv"
	"strings"

	"golang.design/x/clipboard"
)

func main() {

	err := clipboard.Init()
	if err != nil {
		panic(err)
	}
	data()
}

func data() {
	largest := 0
	total := 0
	content := clipboard.Read(clipboard.FmtText)
	var keys []string
	if content == nil {
		panic("No clipboard content")
	}
	sets := string(content)
	setsMap := strings.Split(sets, "\r\n\r\n")
	elves := make(map[string][]int)

	for i, numbers := range setsMap {
		lSet := strings.Split(numbers, "\r\n")
		var intset []int
		for _, numberStr := range lSet {
			number, err := strconv.Atoi(strings.TrimSpace(numberStr))
			if err != nil {
				//fmt.Println("Error converting number:", err)
				continue
			}
			intset = append(intset, number)
		}
		key := fmt.Sprintf("Elf %d", i+1)
		elves[key] = intset
		keys = append(keys, key)
	}

	for _, key := range keys {
		values := elves[key]
		subTotal := 0
		for _, value := range values {
			subTotal += value
		}
		if subTotal > largest {
			largest = subTotal
		}
		total += subTotal
		fmt.Println(key, " Calories:", subTotal)
	}
	fmt.Println("Number of Elves:", len(elves))
	fmt.Println("Largest:", largest)
	fmt.Println("Total:", total)

}
