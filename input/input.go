package input

import "fmt"

//Функция ввода данных реализации fizzbuzz, возвращает число
func InputNumber() int {
	var number int
	fmt.Printf("Введите число для реализации программы FizzBuzz\n")
	fmt.Scanf("%d\n", &number)
	return number
}

//Функция ввода данных команды, возвращает имя команды, string
func CommandName() (string, error) {
	var command string
	fmt.Printf("Введите команду:\nAdd - добавляет слово в словарь\nGet - узнать определение слова\nDelete - удалить слово\nList - просмотреть весь словарь\n")
	_, err := fmt.Scanf("%s\n", &command)
	if err != nil {
		return "", fmt.Errorf("ошибка при чтении данных: %v", err)
	}
	for {
		switch command {
		case "Add", "Get", "Delete", "List":
			return command, nil
		default:
			fmt.Println("Неверная команда")

		}
	}
}
