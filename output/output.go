package output

import (
	"fmt"
	"io"
	"os"
)

var Stdout io.Writer = os.Stdout

// PrintFizzBuzz - выводит на печать последовательность fizzbuzz
func PrintFizzBuzz(incoming [][3]string) error {

	if incoming == nil {
		return fmt.Errorf("пустой вывод данных")
	}
	for _, row := range incoming {
		fmt.Fprintf(Stdout, "%s      %s      %s\n", row[0], row[1], row[2])
	}
	return nil
}

// PrintFizzBuzzNumb - выводит на печать количество элементов по типам
func PrintFizzBuzzNumb(incoming [4]int) error {
	if incoming == [4]int{0, 0, 0, 0} {
		return fmt.Errorf("на печать пришли нулевые значения")
	}
	fmt.Fprintf(Stdout, "Количество FizzBuzz - %v, Количество Fizz - %v, Количество Buzz - %v, Количество числовых значений - %v \n", incoming[0], incoming[1], incoming[2], incoming[3])
	return nil
}
