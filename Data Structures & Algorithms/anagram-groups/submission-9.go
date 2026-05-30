func groupAnagrams(strs []string) [][]string {
    angs := make(map[[26]int][]string)

    for _, s := range strs {
        var letters [26]int
        for _, r := range s {
            letters[r-'a']++
        }
        angs[letters] = append(angs[letters], s)
    }

    answer := make([][]string, 0, len(angs))

    for _, v := range angs {
        answer = append(answer, v)
    }
    
    return answer
}
