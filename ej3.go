package main

import "fmt"

func removeZeroSumSublists(arr []int) []int {
	if arr == nil {
		return []int{}
	}

	prefixSumMap := make(map[int]int)
	var result []int
	prefixSum := 0

	// Iterar sobre el arreglo para identificar subarreglos que suman a cero
	for i := 0; i < len(arr); i++ {
		prefixSum += arr[i]

		// Si la suma es cero, resetear el resultado
		if prefixSum == 0 {
			result = []int{}
			prefixSumMap = make(map[int]int)
		} else if index, found := prefixSumMap[prefixSum]; found {
			// Si ya hemos visto este prefixSum, eliminar la subsecuencia desde la última aparición
			result = result[:index]
		} else {
			// Añadir el valor actual al resultado y registrar el prefixSum en el mapa
			result = append(result, arr[i])
			prefixSumMap[prefixSum] = len(result)
		}
	}

	return result
}

func main() {
	// Ejemplo de uso
	arr := []int{3, 4, -7, 5, -6, 2, 5, -1, 8}
	result := removeZeroSumSublists(arr)
	fmt.Println(result)
}
