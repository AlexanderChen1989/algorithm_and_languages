package main

func isValid(s string) bool {
    var stack []byte 
    for i := 0; i < len(s); i++ {
        c := s[i]
        if c == '(' || c == '{' || c == '[' {
            stack = append(stack, c)
            continue
        }
        
        if len(stack) == 0 {
            return false 
        }
        
        top := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        
        if (c == ')' && top != '(') ||
            (c == '}' && top != '{') || 
            (c == ']' && top != '[') {
            return false
        }
    }
    
    return len(stack) == 0
}
