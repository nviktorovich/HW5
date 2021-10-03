package main

import (
	"errors"
	"fmt"
	"os"
	"time"
)

func userInput() (userInt int, err error) {
	fmt.Println("Введите порядковый номер числа Фибоначчи")
	_, err = fmt.Scan(&userInt)
	if err != nil || userInt < 0 {
		err = errors.New("необходимо ввести положительное целое число")
		return
	}
	return
}

var memory = map[int]int {0: 0, 1: 1}

func getFib(fibNum int)  (fibValue int){

	for k := range memory {
		if fibNum == k {
			fibValue = memory[fibNum]
			return
		}
	}
	memory[fibNum - 1] = getFib(fibNum - 1)
	return getFib(fibNum - 1) + getFib(fibNum - 2)
}

func main()  {
	fibNum, err := userInput()
	start := time.Now()
	if err != nil{
		fmt.Printf("ошибка ввода. %s", err)
		os.Exit(1)
	}
	fmt.Println(getFib(fibNum))
	stop := time.Now()
	fmt.Println(stop.Sub(start))
}
