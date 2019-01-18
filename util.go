package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetUserSelection(selections []string) (num int) {
	for i, match := range selections {
		fmt.Printf("%d: %v\n", i, match)
	}
	var err error
	for {
		fmt.Printf("Please enter the number corresponding to the desired entry (%d-%d) or -1 if none match: ",
			0, len(selections)-1)
		var selection string
		fmt.Scanln(&selection)
		num, err = strconv.Atoi(selection)
		if err == nil && num < len(selections) {
			return
		}
	}
}

func ReadCustomList(filename string) (names []string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		names = append(names, strings.TrimSpace(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return
}
