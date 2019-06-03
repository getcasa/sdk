package sdk

import "fmt"

// Context .
type Context interface {
	PrintLog(string)
}

// Plugin .
type Plugin struct {
}

// PrintLog print log to casa gateway
func (p Plugin) PrintLog(str string) {
	fmt.Println(str)
}
