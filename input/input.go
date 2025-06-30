package input

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Функция ввода данных реализации fizzbuzz, возвращает число
func InputNumber() (int, error) {
	var number int
	fmt.Printf("Введите число для реализации программы FizzBuzz\n")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	number, err := strconv.Atoi(input)
	if err != nil {
		return 0, fmt.Errorf("введено не целое число")
	}
	if number <= 0 {
		return 0, fmt.Errorf("число не может быть отрицательным или равным нулю")
	}

	return number, nil
}
