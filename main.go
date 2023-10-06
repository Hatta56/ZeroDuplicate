package main

import "fmt"

func duplicateZeros(arr []int) {
	n := len(arr)
	zerosCount := 0

	// Menghitung jumlah nol dalam array
	for _, num := range arr {
		if num == 0 {
			zerosCount++
		}
	}

	// Indeks di akhir array setelah pemindahan elemen
	newIndex := n + zerosCount - 1
	fmt.Println(newIndex)
	// Memproses array dari belakang
	for i := n - 1; i >= 0; i-- {
		// Salin elemen ke newIndex
		if newIndex < n {
			arr[newIndex] = arr[i]
		}

		// Jika elemen yang disalin adalah nol, salin lagi
		if arr[i] == 0 {
			newIndex--
			if newIndex < n {
				arr[newIndex] = 0
			}
		}

		newIndex--
	}
}

func main() {
	arr1 := []int{1, 0, 2, 3, 0, 4, 5, 0}
	duplicateZeros(arr1)
	fmt.Println(arr1) // Output: [1 0 0 2 3 0 0 4]

	arr2 := []int{1, 2, 3}
	duplicateZeros(arr2)
	fmt.Println(arr2) // Output: [1 2 3]
}
