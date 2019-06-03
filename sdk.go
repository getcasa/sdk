package sdk

import "fmt"

// PrintLogTypes define PrintLog function
type PrintLogTypes func(string)

// PrintLog print log to casa gateway
func PrintLog(str string) {
	fmt.Println(str)
}
