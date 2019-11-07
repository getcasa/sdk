# Casa software development kit

## Examples
- [Xiaomi](https://github.com/getcasa/plugin-xiaomi)
- [Yeelight](https://github.com/getcasa/plugin-yeelight)
- [Philips Hue](https://github.com/getcasa/plugin-philipshue)
- [Request](https://github.com/getcasa/plugin-request)

## Template
```go
package main

import (
	"github.com/getcasa/sdk"
)

func main() {}

var Config = sdk.Configuration{
	Name:        "Example",
	Version:     "1.0.0",
	Author:      "John Doe",
	Description: "A little example",
	Main:        "",
	FuncData:    "OnData",
	Discover:    true,
	Triggers:    []sdk.Trigger{
    sdk.Trigger{
			Name:    "temperature",
			FieldID: "ID",
			Fields: []sdk.Field{
				sdk.Field{
					Name:          "temp",
					Direct:        false,
					Type:          "int"
				},
				sdk.Field{
					Name:   "unit",
					Direct: false,
					Type:   "string",
					Possibilities: []string{"celsius", "fahrenheit"},
				},
			},
		},
  },
	Actions: []sdk.Action{
		sdk.Action{
			Name: "turnOn",
			Fields: []sdk.Field{
				sdk.Field{
					Name:   "timeout",
					Type:   "int",
					Config: true,
				},
			},
		},
	},
}

// Init function is called only once for the first initialization of casa gateway
func Init() []byte {
  // Create a configuration and return it to casa server
	return []byte("test")
}

// OnStart is called on start and restart of casa gateway
func OnStart(config []byte) {
	// config is the Init config previously created
}

// Params define actions parameters available
type Params struct {
	Timeout  int
}

// OnData get data from source
func OnData() interface{} {
  // wait data from ws, udp or whatever
  
  // switch in case of multiple devices
	switch device {
    case "temperature":
      // get temperature from device
      return temp
    default:
      return nil
	}
}

// CallAction call functions from actions
func CallAction(physicalID string, name string, params []byte, config []byte) {
	if string(params) == "" {
		fmt.Println("Params must be provided")
		return
	}

	// declare parameters
	var req Params

	// unmarshal parameters to use in actions
	err := json.Unmarshal(params, &req)
	if err != nil {
		fmt.Println(err)
	}

	// use name to call actions
	switch name {
	case "turnOn":
		// TODO: add call
	default:
		return
	}
}

// OnStop is called on stop
func OnStop() {
}
```

## Build
To work with casa gateway, a casa plugin must be compiled as a go plugin.
```
go build -buildmode=plugin -o example.so *.go
```
