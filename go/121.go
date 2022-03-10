package main

func maxProfit(prices []int) int {
    res := 0 
    l := 0 
    for r := range prices {
        for prices[r] - prices[l] < 0 {
            l++
        }
        if res < prices[r] - prices[l] {
            res = prices[r] - prices[l]
        }
    }
    
    return res
}