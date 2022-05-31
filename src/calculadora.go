package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type calcu struct{}

func (calcu) operate(entry, operator string) int {
	cleanEntry := strings.Split(entry, operator)
	factor1 := parseTo(cleanEntry[0])
	factor2 := parseTo(cleanEntry[1])
	var result int
	switch operator {
	case "+":
		result = factor1 + factor2
	case "-":
		result = factor1 - factor2
	case "*":
		result = factor1 * factor2
	case "/":
		result = factor1 / factor2
	default:
		fmt.Printf("invalid operator '%v'\n", operator)
	}
	return result
}

func parseTo(entry string) int {
	factor, _ := strconv.Atoi(entry) //string entrada convertido a entero
	return factor
}

func readEntry() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func main() {
	entry := readEntry()
	operator := readEntry()
	c := calcu{}
	result := c.operate(entry, operator)
	fmt.Println(result)
}
