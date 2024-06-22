func isPalindrome(x int) bool {
    if x < 0 {
        return false
    }
    
    var y int
    temp := x
    
    for temp > 0 {
        mod := temp % 10
        y = (y * 10) + mod
        temp /= 10
    }
    
    return y == x
}