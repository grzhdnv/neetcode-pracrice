func twoSum(nums []int, target int) []int {
    digits := make(map[int]int)
    for i, v := range nums {
        difference := target - v
        if idx, ok := digits[difference]; ok {
            return []int{idx, i}
        }
        digits[v] = i
    }
    return nil
}
