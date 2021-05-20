package main

func arrayWithoutRepeatNumbers(slice []int) []int {
	mapSlice := make(map[int]int) // створення мапи  для подальшого заповннея її з елементів вхідного слайсу
	var resultSlice []int
	for _, val := range slice {
		if _, ok := mapSlice[val]; ok { //перевірка умови повторення елементів
			mapSlice[val] += 1
		} else {
			mapSlice[val] = 0
			resultSlice = append(resultSlice, val)
		}
	}
	return resultSlice
}
func foo() {

}
