package sdk

// Field struct
type Field struct {
	Name          string
	Direct        bool
	Type          string
	Possibilities []string
	Config        bool
}

// Trigger struct
type Trigger struct {
	Name          string
	Direct        bool
	Type          string
	Possibilities []string
}

// Device struct define Device necessary datas
type Device struct {
	Name           string
	Description    string
	DefaultTrigger string
	DefaultAction  string
	Triggers       []Trigger
	Actions        []string
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
	Devices     []Device
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

// DiscoveredDevice struct define discovered device
type DiscoveredDevice struct {
	Name         string `db:"name" json:"name"`
	GatewayID    string `db:"gateway_id" json:"gatewayId"`
	PhysicalID   string `db:"physical_id" json:"physicalId"`
	PhysicalName string `db:"physical_name" json:"physicalName"`
	Plugin       string `db:"plugin" json:"plugin"`
}

// FindDevicesFromName find trigger with name trigger
func FindDevicesFromName(devices []Device, name string) Device {
	for _, device := range devices {
		if device.Name == name {
			return device
		}
	}
	return Device{}
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
