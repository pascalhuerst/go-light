package main

import "fmt"

func main() {
	fmt.Println("Hello, world.")

	err := ReadEngine("/home/paso/go/src/github.com/holoplot/go-light/qlcplus_engine.xml")
	if err != nil {
		fmt.Printf("%v\n", err.Error())
		return
	}

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
