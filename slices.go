package main

import "fmt"

/*
Слайс (срез) — это "обёртка" вокруг массива, которая предоставляет более удобный и гибкий интерфейс
для работы с последовательностями данных.

Технически срез состоит из трех компонентов:
	1. Указатель на базовый массив
	2. Длина (length) — текущее количество элементов
	3. Ёмкость (capacity) — максимальное количество элементов без переаллокации

Основные способы создания
// 1. Литерал среза (аналогично массиву, но без указания длины)
slice1 := []int{1, 2, 3}

// 2. Создание из массива
arr := [5]int{1, 2, 3, 4, 5}
slice2 := arr[1:4] // [2, 3, 4]

// 3. Использование make()
slice3 := make([]int, 3)      // длина 3, ёмкость 3
slice4 := make([]int, 3, 5)   // длина 3, ёмкость 5

// 4. Пустой срез
var slice5 []int             // nil-срез
slice6 := []int{}            // пустой срез, но не nil


Добавление элементов (append)
Функция append возвращает новый слайс, который будет содержать новый элемент.
При этом слайс, который передается в качестве аргумента, остается неизменным.

slice := make([]int, 5, 5)
currentLenght, currentCapacity = len(slice), cap(slice)
slice = append(slice, newValue, currentCapacity+1)


*/

func main() {
	slice := make([]int, 5)
	fmt.Println(slice)
	slice = append(slice, 1)
	fmt.Println(slice)

	// Инициализация с ненулевой длиной
	slice = make([]int, 3)
	slice = append(slice, 1) // [0, 0, 0, 1]

	// then append by index
	slice = make([]int, 3)
	slice[0] = 1 // [1, 0, 0]

	// or initialise empty slice
	slice = make([]int, 0, 3)
	slice = append(slice, 1) // [1, 0, 0]
}
