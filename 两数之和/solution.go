package main

func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	for key, value := range nums {
		tmp := target - value
		if i, ok := m[tmp]; ok {
			return []int{i, key}
		}
		m[value] = key
	}
	return nil
}
