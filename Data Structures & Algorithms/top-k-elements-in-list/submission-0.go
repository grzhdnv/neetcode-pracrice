func topKFrequent(nums []int, k int) []int {
    frequencies := make(map[int]int)
    for _, n := range nums {
        frequencies[n]++
    }

    freq := make([][]int, len(nums)+1)
    for k, v := range frequencies {
        freq[v] = append(freq[v], k)
    }

    var kFrequent []int

    for i := len(freq)-1; i >= 0 && len(kFrequent) < k; i-- {
        kFrequent = append(kFrequent, freq[i]...)
    }
    return kFrequent
}
