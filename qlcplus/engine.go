package qlcplus

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Engine interface {
}

// Root type of this XML
type Root struct {
	InputOutputMap InputOutputMap `xml:"InputOutputMap"`
	Fixtures       []Fixture      `xml:"Fixture"`
	FixtureGroups  []FixtureGroup `xml:"FixtureGroup"`
	ChannelGroups  []ChannelGroup `xml:"ChannelsGroup"`
}

// ChannelGroup node in XML
type ChannelGroup struct {
	ID       int    `xml:"ID,attr"`
	Name     string `xml:"Name,attr"`
	Value    int    `xml:"Value,attr"`
	Channels string `xml:",chardata"`
}

// FixtureGroup node in XML
type FixtureGroup struct {
	ID     int    `xml:"ID,attr"`
	Name   string `xml:"Name"`
	Width  int    `xml:"X,attr"`
	Height int    `xml:"Y,attr"`
	Heads  []Head `xml:"Head"`
}

// Head node in XML
type Head struct {
	X       int `xml:"X,attr"`
	Y       int `xml:"Y,attr"`
	Fixture int `xml:"Fixture,attr"`
	Value   int `xml:",chardata"`
}

// Fixture node in XML
type Fixture struct {
	Manufacturer string `xml:"Manufacturer"`
	Model        string `xml:"Model"`
	Mode         string `xml:"Mode"`
	ID           string `xml:"ID"`
	Name         string `xml:"Name"`
	Universe     string `xml:"Universe"`
	Address      string `xml:"Address"`
	Channels     string `xml:"Channels"`
	ExcludeFade  string `xml:"ExcludeFade"`
}

// InputOutputMap node in XML
type InputOutputMap struct {
	Universes []Universe `xml:"Universe"`
}

// Universe node in XML
type Universe struct {
	Name   string `xml:"Name,attr"`
	ID     int    `xml:"ID,attr"`
	Output Output `xml:"Output"`
}

// Output node in XML
type Output struct {
	Plugin          string           `xml:"Plugin,attr"`
	Line            int              `xml:"Line,attr"`
	PluginParameter PluginParameters `xml:"PluginParameters"`
}

// PluginParameters node in XML
type PluginParameters struct {
	OutputIP     string `xml:"outputIP,attr"`
	OutputUni    string `xml:"outputUni,attr"`
	TransmitMode string `xml:"transmitMode,attr"`
}

// ReadEngine does that
func ReadEngine(path string) error {
	xmlFile, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)

	var root Root
	err = xml.Unmarshal(byteValue, &root)
	if err != nil {
		fmt.Println(err)
		return err
	}

	//spew.Dump(fixture)
	println("-----------------------------------------------")

	for _, universe := range root.InputOutputMap.Universes {

		fmt.Printf("Universe: Name=%v ID=%v\n", universe.Name, universe.ID)
		fmt.Printf("  Output: Plugin=%v Line=%v\n", universe.Output.Plugin, universe.Output.Line)
		fmt.Printf("    PluginParameters: outputIP=%v outputUni=%v transmitMode=%v\n",
			universe.Output.PluginParameter.OutputIP,
			universe.Output.PluginParameter.OutputUni,
			universe.Output.PluginParameter.TransmitMode,
		)
	}

	fmt.Printf("\n\nTotal Universes: %v\n\n", len(root.InputOutputMap.Universes))
	println("-----------------------------------------------")

	for _, fixture := range root.Fixtures {
		fmt.Printf("Fixture\n")
		fmt.Printf("  Manufacturer=%v\n", fixture.Manufacturer)
		fmt.Printf("  Model=%v\n", fixture.Model)
		fmt.Printf("  Mode=%v\n", fixture.Mode)
		fmt.Printf("  ID=%v\n", fixture.ID)
		fmt.Printf("  Name=%v\n", fixture.Name)
		fmt.Printf("  Universe=%v\n", fixture.Universe)
		fmt.Printf("  Address=%v\n", fixture.Address)
		fmt.Printf("  Channels=%v\n", fixture.Channels)
		fmt.Printf("  ExcludeFade=%v\n", fixture.ExcludeFade)
	}

	fmt.Printf("\n\nTotal Fixtures: %v\n\n", len(root.Fixtures))
	println("-----------------------------------------------")

	for _, fixtureGroup := range root.FixtureGroups {
		fmt.Printf("FixtureGroup ID=%v Name=%v w=%v h=%v\n",
			fixtureGroup.ID,
			fixtureGroup.Name,
			fixtureGroup.Width,
			fixtureGroup.Height,
		)
		for _, head := range fixtureGroup.Heads {
			fmt.Printf("  Head X=%v Y=%v Fixture=%v Value=%v\n",
				head.X,
				head.Y,
				head.Fixture,
				head.Value,
			)
		}
	}

	fmt.Printf("\n\nTotal FixtureGroups: %v\n\n", len(root.FixtureGroups))
	println("-----------------------------------------------")

	for _, channelGroup := range root.ChannelGroups {
		fmt.Printf("ChannelGroup ID=%v Name=%v Value=%v Channels=%v\n",
			channelGroup.ID,
			channelGroup.Name,
			channelGroup.Value,
			channelGroup.Channels,
		)
	}

	fmt.Printf("\n\nTotal ChannelGroups: %v\n\n", len(root.ChannelGroups))
	println("-----------------------------------------------")

	return nil
}
