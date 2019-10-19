package main

import (
	"fmt"
	"qlcplus"
)

func main() {
	fmt.Println("Hello, world.")

	qlcFixture, err := qlcplus.ReadFixture("/usr/share/qlcplus/fixtures/Eurolite-LED-TMH-6.qxf")
	if err != nil {
		return
	}

	qlcplus.Print(qlcFixture)

	//fixture, err := data.NewFixtureFromQlc(qlcFixture)
	//err = data.

	/*
		err := ReadEngine("/home/paso/go/src/github.com/holoplot/go-light/qlcplus/engine.xml")
		if err != nil {
			fmt.Printf("%v\n", err.Error())
			return
		}
	*/

	/*
		err := ReadFixture("/usr/share/qlcplus/fixtures/Eurolite-LED-TMH-6.qxf")
		if err != nil {
			return
		}

		err := SendFrame()
		if err != nil {
			return
		}
	*/

}
