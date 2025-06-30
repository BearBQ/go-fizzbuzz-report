package main

import (
	"fmt"
	"go-fizzbuzz-report/fizzbuzz"
)

func main() {
	for {
		err := fizzbuzz.RunFizzBuzz()
		if err != nil {
			fmt.Println(err)
		}
	}
}
