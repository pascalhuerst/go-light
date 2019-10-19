package data

import (
	"encoding/json"
	"fmt"
	"os"

	"qlcplus"
)

// NewFixtureFromQlc creates a new ficture from a qlcpro fixture
func NewFixtureFromQlc(source qlcplus.FixtureDefinition) (FixtureDefinition, error) {
	return FixtureDefinition{
		manufacturer: "Manufacturer",
		name:         source.Name,
		lampType:     ColorChanger,
		Channels:     source.Channels,
		Modes:        source.Modes,
	}, nil
}

func writeFicture(f FixtureDefinition) error {

	fileName := fmt.Sprintf("%v", f.name)

	file, err := os.OpenFile(fileName, os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Printf("Cannot open file: %v", err)
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(f)
	if err != nil {
		fmt.Printf("Cannot encode json file: %v", err)
		return err
	}

	return nil
}

// FixtureDefinition data model for a fixture
type FixtureDefinition struct {
	manufacturer string   `json:"manufacturer"`
	name         string   `json:"name"` // Model
	lampType     LampType `json:"type"`

	Channels []Channel `json:"channel"`
	Modes    []Mode    `json:"mode"`
}

// Channel channel
type Channel struct {
	Name         string       `json:"name"`
	Group        Group        `json:"group"`
	Capabilities []Capability `json:"capability"`
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
	Name            string           `json:"name"`
	Physical        Physical         `json:"physical"`
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
