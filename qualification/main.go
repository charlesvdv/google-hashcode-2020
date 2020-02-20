package main

import (
	"bufio"
	"fmt"
	"os"
)

type Library struct {
	BookCount     int
	SignupTime    int
	MaxBookPerDay int
	BookIDS       []int
}

type Process struct {
	BookProcessed map[int]bool
}

func main() {
	days, books, librairies := parseInput(bufio.NewReader(os.Stdin))

	fmt.Printf("%v\n%v\n%v\n", days, books, librairies)
}
