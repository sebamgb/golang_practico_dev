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
	valores := strings.Split(operation, "+") //lista de strings; cada indice son strings en operation sin el + ya que se usa como separado
	fmt.Println(valores[0] + valores[1])     //concatenación de strings en valores
	operator1, _ := strconv.Atoi(valores[0]) //string valores[0] convertido a entero
	operator2, _ := strconv.Atoi(valores[1]) //string valores[1] convertido a entero
	suma := operator1 + operator2
	fmt.Println(suma)
}
