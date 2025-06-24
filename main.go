package main

func main() {
	var numb int = 44
	newFizz := NewFizzBuzz(numb)
	CalcFizzBuzz(newFizz, numb)
	PrintFizzBuzz(newFizz)
	PrintFizzBuzzNumb((newFizz))
}
