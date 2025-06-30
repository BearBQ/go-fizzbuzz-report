package output

import "fmt"

func PrintFizzBuzz(incoming [][3]string) error { //функция печати
	if incoming == nil {
		return fmt.Errorf("пустой вывод данных")
	}
	fmt.Println(incoming)
	return nil
}

func PrintFizzBuzzNumb([4]int) error { //функция печати

}
