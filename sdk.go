package sdk

import "fmt"

// Types define all available functions and variables
type Types struct {
	PrintLog func(string)
}

// PrintLog print log to casa gateway
func PrintLog(str string) {
	fmt.Println(str)
}
