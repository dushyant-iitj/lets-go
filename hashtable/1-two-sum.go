package hashtable

import "fmt"

func twoSum(numbers []int, target int) []int {
	keyTable := make(map[int]int)

	for index, value := range numbers {
		complement := target - value

		// complement exists
		if existingIndex, exists := keyTable[complement]; exists {
			return []int{existingIndex, index}
		} else {
			// else add to table
			keyTable[value] = index
		}
	}

	// nothing found
	return []int{}
}

func TwoSum() {
	numbers := []int{}
	var target int
	// Example 1:
	// Input:
	// numbers := []int{2, 7, 11, 15}
	// target := 9
	// Output: [0,1]
	numbers = []int{2, 7, 11, 15}
	target = 9
	fmt.Println(twoSum(numbers, target))

	// Example 2:
	// Input:
	// numbers := []int{3, 2, 4}
	// target := 6
	// Output: [1,2]
	numbers = []int{3, 2, 4}
	target = 6
	fmt.Println(twoSum(numbers, target))

	// Example 3:
	// Input:
	// numbers := []int{3, 3}
	// target := 6
	// Output: [0,1]
	numbers = []int{3, 3}
	target = 6
	fmt.Println(twoSum(numbers, target))
}
