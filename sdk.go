package sdk

import "fmt"

type Context interface {
	PrintLog(string)
}

type context struct {
}

// PrintLog print log to casa gateway
func (c *context) PrintLog(str string) {
	fmt.Println(str)
}
