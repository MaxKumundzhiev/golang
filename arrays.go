package main

import "fmt"

/*
Массив - фундаментальная структура данных, представляющая собой последовательность
		 элементов одного типа фиксированной длины.

При объявлении массива сразу указывается размер и это влияет на тип переменной
array := [3]int{1,2,3}
fmt.Println(reflect.ValueOf(array).Type().String()) // [3]int

Проинициализировать массив можно разными способами:
var array1 [3]int // [0 0 0]
array5 := [3]int{} // [0,0,0]
array2 := [3]int{1, 2, 3} // [1 2 3]
array3 := [3]int{1: 2} // [0 2 0]
array4 := [...]int{1: 2} // [0 2]


Массив может принимать любой тип в качестве типа элементов, даже другой массив.
Так мы можем объявить многомерный массив, где каждым элементом массива будет другой массив:

matrix := [2][2]{
	{1, 2},
	{3, 4},
}

for rowIndex, row := range matrix {
	for columnIndex, value := range row {
		fmt.Println("rowIndex %d, columnIndex %d, value %d\n", rowIndex, columnIndex, value)
	}
}


Практические примеры использования
- Буферизация данных (array of bytes)
var buffer [256]bytes
buffer := [256]bytes

- Фиксированные наборы значений
var daysOfWeek [7]string = [7]string {
	"Monday",
	"Tuesday",
	...
}

- Массивы как ключи в map
// Только если размеры массивов одинаковые
positions := map[[2]int]string{
    {0, 0}: "origin",
    {1, 2}: "target",
}

to chnage array inside of function
we need to pass the pointer to the function

func ChangeFirstElement(arr *[3]int) {
	arr[0] = 100
}

func main() {
	array := [3]int{1, 2, 3}
	ChangeArray(&array)
	fmt.Println(array)
}



practical example
func RotateLeft(values [5]int, rotateFactor int) [5]int {
	arrayLength = len(values)
	shiftedArray := [5]int

	for currIndex, value := values {
		shiftedIndex = (currIndex - rotateFactor + lenght) % lenght
		shiftedArray[shiftedIndex] = value
	}

	return shiftedArray
}
*/

func main() {
	// arrays
	var ArrayA [3]int // [0, 0, 0]

	const size = 10
	/*
		can not be used VAR size while inialisation
		due to mutability of "var" size
	*/
	var ArrayB [size]int // [0, 0, ..., 0]

	ArrayC := []int {1, 2, 3}

	fmt.Println(ArrayA)
	fmt.Println(ArrayB)
	fmt.Println(ArrayC)

	// slices (has lenght and capacity properties)
	var SliceA []int  // size=0, capacity=0	
	fmt.Println(len(SliceA), cap(SliceA))

	SliceB := make([]int, 5, 10)  // size=0, capacity=0
	fmt.Println(SliceB)
	fmt.Println(len(SliceB), cap(SliceB))

	SliceB = append(SliceB, )
	

	array := [3]int{}
	fmt.Println(array)
}

