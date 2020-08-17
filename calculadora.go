package mycalculator

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func (Calculadora) Operate(operator string, n1 int, n2 int) int {
	switch operator {
	case "+":
		return n1 + n2
	case "-":
		return n1 - n2
	case "*":
		return n1 * n2
	case "/":
		return n1 / n2
	default:
		return 0
	}
}

func GetNumber() int {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Introduzca un numero")
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

func GetOperator() string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Qué operación desea realizar?")
	fmt.Println("Introduzca: \n+ para sumar \n- para restar \n* para multiplicar \n/ para dividir")
	scanner.Scan()
	return scanner.Text()
}

type Calculadora struct {
}
