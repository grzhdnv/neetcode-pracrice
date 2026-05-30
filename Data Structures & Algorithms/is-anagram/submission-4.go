func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    count := [26]int{}
    t2 := []rune(t)

    for i, r := range s {
        count[r - rune('a')]++
        count[t2[i] - rune('a')]--
    }

    for _, b := range count {
        if b != 0 {
            return false
        }
    }

    return true
}
