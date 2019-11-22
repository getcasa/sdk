package sdk

// Field struct
type Field struct {
	Name          string   `json:"name"`
	Description   string   `json:"description"`
	Direct        bool     `json:"direct"`
	Type          string   `json:"type"`
	Possibilities []string `json:"possibilities"`
	Min           int      `json:"min"`
	Max           int      `json:"max"`
	Config        bool     `json:"config"`
}

// Trigger struct
type Trigger struct {
	Name          string   `json:"name"`
	Description   string   `json:"description"`
	Direct        bool     `json:"direct"`
	Type          string   `json:"type"`
	Possibilities []string `json:"possibilities"`
}

// Device struct define Device necessary datas
type Device struct {
	Name           string         `json:"name"`
	Description    string         `json:"description"`
	DefaultTrigger string         `json:"defaultTrigger"`
	DefaultAction  string         `json:"defaultAction"`
	Config         []DeviceConfig `json:"config"`
	Triggers       []Trigger      `json:"triggers"`
	Actions        []string       `json:"actions"`
}

// DeviceConfig struct define device's configuration
type DeviceConfig struct {
	Name      string
	Descriton string
	Type      string
}

// Action struct define action necessary datas
type Action struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Fields      []Field `json:"fields"`
}

// Configuration struct define configuration of plugin
type Configuration struct {
	Name        string   `json:"name"`
	Version     string   `json:"version"`
	Author      string   `json:"author"`
	Description string   `json:"description"`
	Discover    bool     `json:"discover"`
	Devices     []Device `json:"devices"`
	Actions     []Action `json:"actions"`
}

// Data struct
type Data struct {
	Plugin       string      `json:"plugin"`
	PhysicalName string      `json:"physicalName"` //Model (ex: switch)
	PhysicalID   string      `json:"physicalId"`
	Data         interface{} `json:"data"`
	Values       []Value     `json:"values"`
}

// Value struct
type Value struct {
	Name  string `json:"name"`
	Value []byte `json:"value"`
	Type  string `json:"type"` //string, int, bool

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
