package input

import "fmt"

//Функция ввода данных реализации fizzbuzz, возвращает число
func InputNumber() (int, error) {
	var number int
	fmt.Printf("Введите число для реализации программы FizzBuzz\n")
	fmt.Scanf("%d\n", &number)
	if number <= 0 {
		return 0, fmt.Errorf("число не может быть отрицательным или равным нулю")
	}
	return number, nil
}
