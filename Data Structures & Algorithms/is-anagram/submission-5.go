func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    letters := make(map[rune]int)
    sr := []rune(s)
    st := []rune(t)
    
    for i := range len(s) {
        letters[sr[i]]++
        letters[st[i]]--
    }

    for _, v := range letters {
        if v != 0 {
            return false
        }
    }
    return true
}
