package sdk

import "fmt"

// PrintLogType define PrintLog function
type PrintLogType func(string)

// PrintLog print log to casa gateway
func PrintLog(str string) {
	fmt.Println(str)
}
