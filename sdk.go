package sdk

// Field struct
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
	Description string
	Default     bool
	Fields      []Field
}

// Action struct define action necessary datas
type Action struct {
	Name        string
	Description string
	Default     bool
	Fields      []Field
}

// Configuration struct define configuration of plugin
type Configuration struct {
	Name        string
	Version     string
	Author      string
	Description string
	Triggers    []Trigger
	Actions     []Action
}

// Data struct
type Data struct {
	Plugin       string
	PhysicalName string //Model (ex: switch)
	PhysicalID   string
	Data         interface{}
	Values       []Value
}

// Value struct
type Value struct {
	Name  string
	Value []byte
	Type  string //string, int, bool

}

// Device struct
type Device struct {
	Name         string `db:"name" json:"name"`
	PhysicalID   string `db:"physical_id" json:"physicalId"`
	PhysicalName string `db:"physical_name" json:"physicalName"`
	Plugin       string `db:"plugin" json:"plugin"`
}

// FindTriggerFromName find trigger with name trigger
func FindTriggerFromName(triggers []Trigger, name string) Trigger {
	for _, trigger := range triggers {
		if trigger.Name == name {
			return trigger
		}
	}
	return Trigger{}
}

// FindValueFromName find trigger with name trigger
func FindValueFromName(values []Value, name string) Value {
	for _, value := range values {
		if value.Name == name {
			return value
		}
	}
	return Value{}
}
