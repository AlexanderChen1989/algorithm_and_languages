package main

func main() {
	println(removeKdigits("1234567890", 9))
}

func removeKdigits(num string, k int) string {
    if len(num) <= k {
        return "0"
    }
    
    result := []byte{}
    for i := 0; i < len(num); i++ {
        if k <= 0 {
            result = append(result, num[i:]...)
            break
        }
        if len(result) == 0 {
            result = append(result, num[i])
            continue
        }
        c := num[i]
        for len(result) > 0 && result[len(result)-1] > c && k > 0 {
            result = result[:len(result)-1] 
            k--
        }
        
        result = append(result, c)
    }
    
    for k > 0 && len(result) > 0 {
        result = result[:len(result)-1]
        k--
    }
    
    for len(result) > 0 && result[0] == '0' {
        result = result[1:]
    }
    
    s := string(result)
    if s == "" {
        return "0"
    }
    return s
}

