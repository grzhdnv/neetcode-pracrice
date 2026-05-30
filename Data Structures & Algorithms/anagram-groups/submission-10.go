func groupAnagrams(strs []string) [][]string {
    angs := make(map[string][]string)

    for _, s := range strs {
        letters := make([]byte, 26)
        for _, r := range s {
            letters[r-'a']++
        }
        angs[string(letters)] = append(angs[string(letters)], s)
    }

    answer := make([][]string, 0, len(angs))

    for _, v := range angs {
        answer = append(answer, v)
    }

    return answer
}
