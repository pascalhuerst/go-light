package data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/pascalhuerst/go-light/qlcplus"
)

// FixtureDefinition data model for a fixture
type FixtureDefinition struct {
	Manufacturer string   `json:"manufacturer"`
	Name         string   `json:"name"` // Model
	LampType     LampType `json:"type"`

	Channels []Channel `json:"channels"`
	Modes    []Mode    `json:"modes"`
}

// Channel channel
type Channel struct {
	Name         string       `json:"name"`
	Group        Group        `json:"group"`
	Capabilities []Capability `json:"capabilities"`
}

// ChannelMapping node in json
type ChannelMapping struct {
	Number      int    `xml:"Number,attr"`
	ChannelName string `xml:",chardata"`
}

// Group group
type Group struct {
	Value string `json:"value"`
	Byte  string `json:"byte"`
}

// Capability capability
type Capability struct {
	Value string `json:"value"`
	Min   int    `json:"min"`
	Max   int    `json:"max"`
	Color string `json:"color"`
	Res   string `json:"res"`
}

// LampType is a enum representation for the type
type LampType int

const (
	// ColorChanger type
	ColorChanger LampType = iota
	// Dimmer type
	Dimmer
	// Effect Type
	Effect
	// Flower type
	Flower
	//Hazer type
	Hazer
	// Laser type
	Laser
	// MovingHead type
	MovingHead
	// Other type
	Other
	// Scanner type
	Scanner
	//Smoke type
	Smoke
	// Strobe type
	Strobe
)

var stringForLampType = map[string]LampType{
	"ColorChanger": ColorChanger,
	"Dimmer":       Dimmer,
	"Effect":       Effect,
	"Flower":       Flower,
	"Hazer":        Hazer,
	"Laser":        Laser,
	"MovingHead":   MovingHead,
	"Other":        Other,
	"Scanner":      Scanner,
	"Smoke":        Smoke,
	"Strobe":       Strobe,
}

var lampTypeForString = map[LampType]string{
	ColorChanger: "ColorChanger",
	Dimmer:       "Dimmer",
	Effect:       "Effect",
	Flower:       "Flower",
	Hazer:        "Hazer",
	Laser:        "Laser",
	MovingHead:   "MovingHead",
	Other:        "Other",
	Scanner:      "Scanner",
	Smoke:        "Smoke",
	Strobe:       "Strobe",
}

// Mode mode
type Mode struct {
	Name string `json:"name"`
	// Physical        Physical         `json:"physical"`
	ChannelMappings []ChannelMapping `json:"channel"`
}

// Physical node in JSON
type Physical struct {
	Bulp      Bulp      `json:"bulp"`
	Dimension Dimension `json:"dimensions"`
	Lens      Lens      `json:"lens"`
	Focus     Focus     `json:"focus"`
}

// Bulp node in JSON
type Bulp struct {
	ColorTemperature int    `json:"colour_temperature"`
	Type             string `json:"type"`
	Lumens           int    `json:"lumens"`
}

// Dimension node in JSON
type Dimension struct {
	Width  float32 `json:"width"`
	Depth  float32 `json:"depth"`
	Height float32 `json:"height"`
	Weight float32 `json:"weight"`
}

// Lens node in JSON
type Lens struct {
	DegreesMin int    `json:"degrees_min"`
	DegreesMax int    `json:"degrees_max"`
	Name       string `json:"name"`
}

// Focus node in JSON
type Focus struct {
	Type    string `json:"type"`
	TiltMax int    `json:"tilt_max"`
	PanMax  int    `json:"pan_max"`
}

// Technical node in JSON
type Technical struct {
	PowerConsumption string `json:"power_consumption"`
	DmxConnector     string `json:"dmx_connector"`
}

func extractCapabilities(capabilities []qlcplus.Capability) []Capability {
	var ret []Capability
	for _, capability := range capabilities {
		newCapability := Capability{
			Value: capability.Value,
			Min:   capability.Min,
			Max:   capability.Max,
			Color: capability.Color,
			Res:   capability.Res,
		}
		ret = append(ret, newCapability)
	}
	return ret
}

func extractChannelMappings(channelMappings []qlcplus.ChannelMapping) []ChannelMapping {
	var ret []ChannelMapping
	for _, channelMapping := range channelMappings {
		newChannelMapping := ChannelMapping{
			Number:      channelMapping.Number,
			ChannelName: channelMapping.ChannelName,
		}
		ret = append(ret, newChannelMapping)
	}
	return ret
}

// NewFixtureFromQlc creates a new ficture from a qlcpro fixture
func NewFixtureFromQlc(source *qlcplus.FixtureDefinition) FixtureDefinition {

	var newChannels []Channel
	for _, channel := range source.Channels {
		var newChannel = Channel{
			Name: channel.Name,
			Group: Group{
				Value: channel.Group.Value,
				Byte:  channel.Group.Byte,
			},
			Capabilities: extractCapabilities(channel.Capabilities),
		}
		newChannels = append(newChannels, newChannel)
	}

	var newModes []Mode
	for _, mode := range source.Modes {
		var newMode = Mode{
			Name:            mode.Name,
			ChannelMappings: extractChannelMappings(mode.ChannelMappings),
		}
		newModes = append(newModes, newMode)
	}

	ret := FixtureDefinition{
		Manufacturer: "Manufacturer",
		Name:         "Name",
		LampType:     ColorChanger,
		Channels:     newChannels,
		Modes:        newModes,
	}

	return ret
}

// WriteFixture saves a JSON of a FixtureDefinition
func WriteFixture(fileName string, f FixtureDefinition) error {

	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Printf("Cannot open file: %v\n", err.Error())
		return err
	}
	defer file.Close()

	var byteBuffer []byte
	byteBuffer, err = json.MarshalIndent(f, "", "  ")
	if err != nil {
		fmt.Printf("Cannot marshall JSON: %v\n", err.Error())
	}

	_, err = file.Write(byteBuffer)
	if err != nil {
		fmt.Printf("Cannot write JSON file: %v\n", err.Error())
	}

	return nil
}

// ReadFixture reads a fixture from a JSON file
func ReadFixture(path string) (*FixtureDefinition, error) {

	jsonfile, err := os.Open(path)
	if err != nil {
		fmt.Printf("Cannot open file: %v\n", err.Error())
		return nil, err
	}
	defer jsonfile.Close()

	var byteValue []byte
	byteValue, err = ioutil.ReadAll(jsonfile)
	if err != nil {
		fmt.Printf("Cannot read file: %v\n", err.Error())
		return nil, err
	}

	var fixture FixtureDefinition
	err = json.Unmarshal(byteValue, &fixture)
	if err != nil {
		fmt.Printf("Cannot unmarshal json data: %v\n", err.Error())
	}

	return &fixture, nil
}
