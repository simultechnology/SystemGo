package mysort

import "fmt"

func BubbleSort(numbers []int) {
	counter := 0
	k := 0
	flag := true
	for flag {
		flag = false
		l := len(numbers)
		for i := 0; i < l-1-k; i++ {
			counter++
			if numbers[i] > numbers[i+1] {
				tmp := numbers[i]
				numbers[i] = numbers[i+1]
				numbers[i+1] = tmp
				flag = true
			}
		}
		k++
	}
	fmt.Printf("total calculation times : %d\n", counter)
}
