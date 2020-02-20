package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
)

func toInt(str string) int {
	val, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return val
}

func parseInput(input io.Reader) (int, []int, []Library) {
	scanner := bufio.NewScanner(input)

	// Fucking scanner
	const maxCapacity = 2048 * 1024
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)

	scanner.Scan()
	headerLine := scanner.Text()
	rawHeaderLine := strings.Split(headerLine, " ")

	bookCount := toInt(rawHeaderLine[0])
	libraryCount := toInt(rawHeaderLine[1])
	numOfDays := toInt(rawHeaderLine[2])

	scanner.Scan()
	line := strings.TrimSpace(scanner.Text())
	bookScores := []int{}
	for _, rawVal := range strings.Split(line, " ") {
		bookScores = append(bookScores, toInt(rawVal))
	}

	libraries := []Library{}
	libraryID := 0
	for scanner.Scan() {
		library := Library{}
		line = scanner.Text()
		lineValues := strings.Split(line, " ")

		library.ID = libraryID
		library.BookCount = toInt(lineValues[0])
		library.SignupTime = toInt(lineValues[1])
		library.MaxBookPerDay = toInt(lineValues[2])

		scanner.Scan()
		line = scanner.Text()
		for _, book := range strings.Split(line, " ") {
			library.BookIDs = append(library.BookIDs, toInt(book))
		}

		libraries = append(libraries, library)
		libraryID += 1
	}

	if len(libraries) != libraryCount {
		log.Fatal("library don't match")
	}

	if len(bookScores) != bookCount {
		log.Fatal("book don't match")
	}

	return numOfDays, bookScores, libraries
}

func format(libRegistered []int, booksByLib map[int][]int) {
	goodLibs := libRegistered[:0]
	for _, x := range libRegistered {
		if len(booksByLib[x]) == 0 {
			continue
		}
		goodLibs = append(goodLibs, x)
	}

	fmt.Printf("%v\n", len(goodLibs))

	for _, lib := range goodLibs {
		books := booksByLib[lib]
		if len(books) == 0 {
			continue
		}
		fmt.Printf("%v %v\n", lib, len(books))

		for i, book := range books {
			if i != 0 {
				fmt.Printf(" ")
			}
			fmt.Printf("%v", book)
		}
		fmt.Printf("\n")
	}
}
