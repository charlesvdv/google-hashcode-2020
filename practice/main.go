package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func parseInput(reader io.Reader) (int, []int) {
	scanner := bufio.NewScanner(reader)

	scanner.Scan()
	rawHeader := tokenize(scanner.Text())
	scanner.Scan()
	data := tokenize(scanner.Text())

	if rawHeader[1] != len(data) {
		log.Fatalf("Expected number of pizzas: '%v', got '%v'", rawHeader[1], len(data))
	}

	return rawHeader[0], data
}

func tokenize(data string) []int {
	result := []int{}
	for _, val := range strings.Split(data, " ") {
		number, err := strconv.Atoi(val)
		if err != nil {
			log.Fatalf("Input with value '%v' is invalid", val)
		}
		result = append(result, number)
	}
	return result
}

func optimizePizzasOrder(slicesToOrder int, pizzas []int) (int, []int) {
	bestSum := 0
	bestSolution := []int{}
	for i := len(pizzas); i > 1; i-- {
		sum, solution := bestFromPizzaSubset(slicesToOrder, pizzas[0:i])

		if bestSum < sum {
			bestSum = sum
			bestSolution = solution
		}
	}
	return bestSum, bestSolution
}

func bestFromPizzaSubset(slicesToOrder int, pizzas []int) (int, []int) {
	currentSum := 0
	currentSolution := []int{}
	for i := len(pizzas) - 1; i >= 0; i-- {
		currentSlice := pizzas[i]
		if currentSum+currentSlice < slicesToOrder {
			currentSum = currentSum + currentSlice
			currentSolution = append(currentSolution, i)
		}
	}

	sort.Ints(currentSolution)

	return currentSum, currentSolution
}

func formatOutput(writer io.Writer, pizzasPicked []int) {
	io.WriteString(writer, strconv.Itoa(len(pizzasPicked))+"\n")
	for i, pizza := range pizzasPicked {
		if i != 0 {
			io.WriteString(writer, " ")
		}
		io.WriteString(writer, strconv.Itoa(pizza))
	}
}

func main() {
	slicesToOrder, pizzas := parseInput(bufio.NewReader(os.Stdin))

	_, bestPizzas := optimizePizzasOrder(slicesToOrder, pizzas)

	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	formatOutput(writer, bestPizzas)
}
