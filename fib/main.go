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

func getFib(fibNum int)  (fibValue int){
	if fibNum == 0 {
		fibValue = 0
		return
	}
	if fibNum == 1 {
		fibValue = 1
		return
	}
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

