package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	operation := scanner.Text()
	fmt.Println(operation)
	operator := "p"
	valores := strings.Split(operation, operator) //lista de strings; cada indice son strings en operation sin el + ya que se usa como separado
	fmt.Println(valores[0] + valores[1])          //concatenación de strings en valores
	factor1, err1 := strconv.Atoi(valores[0])     //string valores[0] convertido a entero
	if err1 != nil {
		fmt.Println(err1)
	}
	factor2, err2 := strconv.Atoi(valores[1]) //string valores[1] convertido a entero
	if err2 != nil {
		fmt.Println(err2)
	}
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
	fmt.Println(result)
}
