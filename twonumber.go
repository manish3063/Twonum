package main

import "fmt"

func main() {

	a := twoSum([]int{2, 7, 11, 15}, 26)

	fmt.Println(a)

	// num := []int{2, 7, 11, 15}

	// target := 26

}

func twoSum(num []int, target int) []int {

	result := []int{}

	for i := 0; i < len(num); i++ {

		for j := 0; j < len(num)-1; j++ {

			if num[i]+num[j] == target && i != j {

				fmt.Println(i, j)
				//fmt.Println(j)
				result = append(result, i, j)

			}

		}

	}
	return result

}
