package main

import (
	"bufio"
	"fmt"
	"os"
)

// example of working with defer
func ReadFile(filename string) error {
	fmt.Println("opened file")
	// read file
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	
	// defer closing file
	defer fmt.Println("closed file")
	defer file.Close()

	// read line by line
	scanner := bufio.NewScanner(file)
	
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
	
	
	return nil
}

func main() {
	/*
	defer - отложить
	defer is a keyword as a func or var or others
	idea of defer is to postpone a call of a particular function to manage the resources
	отложить выполнение указанной функции до момента завершения окружающей функции. 
	То есть defer позволяет управлять ресурсами и гарантировать выполнение определённых действий независимо от того, 
	как будет завершена текущая функция — будь то успешное выполнение или неожиданный выход из-за ошибки.

	data structure used behind is stack aka LIFO queue

	order
		start
		end
		3rd
		2nd
		1st

	use cases:
		- close files, connections, channels
		- unlock mutexes
		- log function exit points
		- handle panics and recover from them
	*/

	// fmt.Println("start")

	// defer fmt.Println("1st call")
	// defer fmt.Println("2nd call")
	// defer fmt.Println("3rd call")

	// fmt.Println("end")

	ReadFile("/Users/macbook/Desktop/go.txt")
}
