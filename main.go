package main

import (
	"fmt"
	"math"
)

func firstBadCode(ar []int, apiCall func(s int) bool) string {
	a_index := 0
	b_index := len(ar) - 1
	a_pointer := ar[a_index]
	b_pointer := ar[b_index]
	step := 0
	m_index := int(math.Floor((float64(a_index) + float64(b_index)) * 0.5))
	m_pointer := ar[m_index]
	for {
		m_index = int(math.Floor((float64(a_index) + float64(b_index)) * 0.5))
		m_pointer = ar[m_index]
		if a_pointer >= b_pointer {
			break
		}
		if apiCall(m_pointer) == true {
			step++
			b_index = m_index
			b_pointer = ar[b_index]
		} else {
			step++
			a_index = m_index + 1
			a_pointer = ar[a_index]
		}
	}
	return fmt.Sprintf("step: %v firstbad : %v", step, m_pointer)
}

func apiCall(s int) bool {
	if s >= 17 {
		return true
	} else {
		return false
	}
}

func main() {
	question := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24}
	answer := firstBadCode(question, apiCall)
	fmt.Println(answer)
}
