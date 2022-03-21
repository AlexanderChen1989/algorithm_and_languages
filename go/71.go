package main

import "strings"

func simplifyPath(path string) string {
	parts := strings.Split(path, "/")
	stack := []string{}

	for _, part := range parts {
		if part == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
			continue
		}
		if part != "" && part != "." {
			stack = append(stack, part)
		}
	}

	return "/" + strings.Join(stack, "/")
}