func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    letters := make(map[rune]int)
    
    for _, r := range s {
        letters[r]++
    }

    for _, r := range t {
        letters[r]--
        if letters[r] < 0 {
            return false
        }
    }
    return true
}
