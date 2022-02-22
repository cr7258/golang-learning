package main

import (
	"fmt"
	"sort"
)

func sortString(s []string, less func(i, j int) bool) {
	if less == nil {
		less = func(i, j int) bool {
			return s[i] < s[j]
		}
	}
	sort.Slice(s, less)
}

func main() {
	food := []string{"onion", "carrot", "celery"}
	sortString(food, nil)
	fmt.Println(food)
}
