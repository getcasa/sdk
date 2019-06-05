package sdk

import "fmt"

type Action struct {
	Name        string
	Description string
}

type Trigger struct {
	Name        string
	Description string
}

type Configuration struct {
	Name        string
	Version     string
	Author      string
	Description string
	Main        string
	FuncData    string
	Triggers    []Trigger
	Actions     []Action
}

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
