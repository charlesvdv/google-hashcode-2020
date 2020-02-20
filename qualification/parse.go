package main

import (
	"bufio"
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

	librairies := []Library{}
	for scanner.Scan() {
		library := Library{}
		line = scanner.Text()
		lineValues := strings.Split(line, " ")

		library.BookCount = toInt(lineValues[0])
		library.SignupTime = toInt(lineValues[1])
		library.MaxBookPerDay = toInt(lineValues[2])

		scanner.Scan()
		line = scanner.Text()
		for _, book := range strings.Split(line, " ") {
			library.BookIDS = append(library.BookIDS, toInt(book))
		}

		librairies = append(librairies, library)
	}

	if len(librairies) != libraryCount {
		log.Fatal("library don't match")
	}

	if len(bookScores) != bookCount {
		log.Fatal("book don't match")
	}

	return numOfDays, bookScores, librairies
}
