package main

import "fmt"

func decode(msg string) int {
	n := len(msg)
	arr := make([]int, n+1)
	arr[0] = 1

	for i := 1; i <= n; i++ {
		if msg[i-1] != '0' {
			arr[i] += arr[i-1]
		}
		// 10 to 26 cases
		if i > 1 && (msg[i-2] == '1' || (msg[i-2] == '2' && msg[i-1] <= '6')) {
			arr[i] += arr[i-2]
		}
	}
	return arr[n]
}

// Output:
// 2
// 3
// 0
func main() {
	fmt.Println(decode("12"))
	fmt.Println(decode("226"))
	fmt.Println(decode("0"))
}
