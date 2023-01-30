package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := [5]int{2, 4, 6, 8, 10} //Массив

	var wg sync.WaitGroup
	wg.Add(len(arr))

	for idx, val := range arr {
		go func(idx, val int) { //Для каждого значения запускается горутина, сигнализирующая об окончании работы через wg.Done()
			arr[idx] = val * val
			wg.Done()
		}(idx, val)
	}
	wg.Wait()        //Ожидание завершения работы всех горутин
	fmt.Println(arr) //Вывод квадратов
}
