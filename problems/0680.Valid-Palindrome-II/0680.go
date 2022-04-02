package p0680

func validPalindrome(s string) bool {
    i, j := 0, len(s)-1
    for i < j {
        if s[i] != s[j] {
            if ifPalindrome(s, i+1, j) || ifPalindrome(s, i, j-1) {
                return true
            }
            return false
        }
        i++
        j--
    }
    return true
}

func ifPalindrome(s string, i int, j int) bool {
    for i < j {
        if s[i] != s[j] {
            return false
        }
        i++
        j--
    }
    return true
}