package sdk

type Trigger struct {
	Name          string
	FieldID       string
	Description   string
	Field         string
	Type          string
	Possibilities []string
}

type Action struct {
	Name        string
	Description string
	Fields      []string
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

type OnData struct {
	Plugin      string
	TriggerName string
	Data        interface{}
}
