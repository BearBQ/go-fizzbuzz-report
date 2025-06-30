package output

import "fmt"

func PrintFizzBuzz(incoming [][3]string) error { //функция печати
	if incoming == nil {
		return fmt.Errorf("пустой вывод данных")
	}
	for n, _ := range incoming {
		fmt.Printf("%s      %s      %s\n", incoming[n][0], incoming[n][1], incoming[n][2])
	}
	return nil
}

func PrintFizzBuzzNumb(incoming [4]int) error { //функция печати
	if incoming == [4]int{0, 0, 0, 0} {
		return fmt.Errorf("на печать пришли нулевые значения")
	}
	fmt.Printf("Количество FizzBuzz - %v, Количество Fizz - %v, Количество Buzz - %v, Количество числовых значений - %v \n", incoming[0], incoming[1], incoming[2], incoming[3])
	return nil
}
