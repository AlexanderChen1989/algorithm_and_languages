package main 

func isPerfectSquare(num int) bool {
    left, right := 1, num
    
    for left <= right {
        mid := (left + right) / 2 
        n := mid * mid
        if n > num {
            right = mid - 1
            continue
        }
        if n < num {
            left = mid + 1 
            continue
        }
        
        return true
    }
    
    return false
}