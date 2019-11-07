package sdk

// Field
type Field struct {
	Name          string
	Direct        bool
	Type          string
	Possibilities []string
	Config        bool
}

// Trigger struct define trigger necessary datas
type Trigger struct {
	Name        string
	FieldID     string
	Description string
	Fields      []Field
}

// Action struct define action necessary datas
type Action struct {
	Name        string
	Description string
	Fields      []Field
}

// Configuration struct define configuration of plugin
type Configuration struct {
	Name        string
	Version     string
	Author      string
	Description string
	Main        string
	FuncData    string
	Discover    bool
	Triggers    []Trigger
	Actions     []Action
}

// OnData struct
type OnData struct {
	Plugin      string
	TriggerName string
	Data        interface{}
}

type Device struct {
	Name         string `db:"name" json:"name"`
	PhysicalID   string `db:"physical_id" json:"physicalId"`
	PhysicalName string `db:"physical_name" json:"physicalName"`
	Plugin       string `db:"plugin" json:"plugin"`
}
