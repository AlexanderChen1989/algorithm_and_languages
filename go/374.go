package main 


/** 
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

 func guess(num int) int { return 0}

 func guessNumber(n int) int {
    left, right := 1, n
    for left <= right {
        mid := (left + right) / 2
        n := guess(mid)
        if n == -1 {
            right = mid - 1
            continue
        }
        if n == 1 {
            left = mid + 1
            continue
        }
        
        return mid
    }
    return 0
}