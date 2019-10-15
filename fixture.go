package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/davecgh/go-spew/spew"
)

type FixtureDefinition struct {
	Channels []Channel `xml:"Channel"`
	Modes    []Mode    `xml:"Mode"`
}

// Mode mode
type Mode struct {
	Name            string           `xml:"Name,attr"`
	Physical        Physical         `xml:"Physical"`
	ChannelMappings []ChannelMapping `xml:"Channel"`
}

// Physical node in XML
type Physical struct {
	Bulp      Bulp      `xml:"Bulp"`
	Dimension Dimension `xml:"Dimensions"`
	Lens      Lens      `xml:"Lens"`
	Focus     Focus     `xml:"Focus"`
}

// Bulp node in XML
type Bulp struct {
	ColorTemperature int    `xml:"ColourTemperature,attr"`
	Type             string `xml:"Type,attr"`
	Lumens           int    `xml:"Lumens,attr"`
}

// Dimension node in XML
type Dimension struct {
	Width  float32 `xml:"Width,attr"`
	Depth  float32 `xml:"Depth,attr"`
	Height float32 `xml:"Height,attr"`
	Weight float32 `xml:"Weight,attr"`
}

// Lens node in XML
type Lens struct {
	DegreesMin int    `xml:"DegreesMin,attr"`
	DegreesMax int    `xml:"DegreesMax,attr"`
	Name       string `xml:"Name,attr"`
}

// Focus node in XML
type Focus struct {
	Type    string `xml:"Type,attr"`
	TiltMax int    `xml:"TiltMax,attr"`
	PanMax  int    `xml:"PanMax,attr"`
}

// Technical node in XML
type Technical struct {
	PowerConsumption string `xml:"PowerConsumption,attr"`
	DmxConnector     string `xml:"DmxConnector,attr"`
}

// ChannelMapping node in XML
type ChannelMapping struct {
	Number      int    `xml:"Number,attr"`
	ChannelName string `xml:",chardata"`
}

// Channel channel
type Channel struct {
	Name         string       `xml:"Name,attr"`
	Group        Group        `xml:"Group"`
	Capabilities []Capability `xml:"Capability"`
}

// Group group
type Group struct {
	Value string `xml:",chardata"`
	Byte  string `xml:"Byte,attr"`
}

// Capability capability
type Capability struct {
	Value string `xml:",chardata"`
	Min   int    `xml:"Min,attr"`
	Max   int    `xml:"Max,attr"`
	Color string `xml:"Color,attr"`
	Res   string `xml:"Res,attr"`
}

// ReadFixture does that
func ReadFixture(path string) error {
	xmlFile, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println("Successfully Opened File")
	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)

	var fixture FixtureDefinition
	err = xml.Unmarshal(byteValue, &fixture)
	if err != nil {
		fmt.Println(err)
		return err
	}

	spew.Dump(fixture)
	println("-----------------------------------------------")

	for _, channel := range fixture.Channels {

		fmt.Printf("[%v]\n", channel.Name)
		fmt.Printf("  Group: %v, Byte: %v\n", channel.Group.Value, channel.Group.Byte)

		for _, cap := range channel.Capabilities {
			fmt.Printf("  Capability: %v\n", cap.Value)
			fmt.Printf("    Max: %v, Min: %v, Color: %v, Res: %v\n", cap.Max, cap.Min, cap.Color, cap.Res)
		}
	}

	println("-----------------------------------------------")

	for _, mode := range fixture.Modes {

		fmt.Printf("[%v]\n", mode.Name)

		for _, channelMapping := range mode.ChannelMappings {
			fmt.Printf("  %v:  %v\n", channelMapping.Number, channelMapping.ChannelName)
		}

		fmt.Printf("  Physical:\n")
		fmt.Printf("    Bulp     : colorTemp: %v, lumens: %v, type: %v\n",
			mode.Physical.Bulp.ColorTemperature,
			mode.Physical.Bulp.Lumens,
			mode.Physical.Bulp.Type)
		fmt.Printf("    Dimension: Depth: %v, Height: %v, Weight: %v, Width: %v\n",
			mode.Physical.Dimension.Depth,
			mode.Physical.Dimension.Height,
			mode.Physical.Dimension.Weight,
			mode.Physical.Dimension.Width)
		fmt.Printf("    Lens     : Name: %v, DegreesMin: %v, DegreesMax: %v\n",
			mode.Physical.Lens.Name,
			mode.Physical.Lens.DegreesMin,
			mode.Physical.Lens.DegreesMax)
		fmt.Printf("    Focus    : Type: %v, TiltMax: %v, PanMax: %v\n",
			mode.Physical.Focus.Type,
			mode.Physical.Focus.TiltMax,
			mode.Physical.Focus.PanMax)

		fmt.Printf("    Channel Mappings:\n")
		for _, channelMapping := range mode.ChannelMappings {
			fmt.Printf("      Number: %v, Name: %v\n", channelMapping.ChannelName, channelMapping.Number)
		}
	}

	return nil
}
