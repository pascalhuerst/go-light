package main

import "fmt"

func main() {
	fmt.Println("Hello, world.")

	/*
		err := ReadFixture("/usr/share/qlcplus/fixtures/Eurolite-LED-TMH-6.qxf")
		if err != nil {
			return
		}
	*/

	err := SendFrame()
	if err != nil {
		return
	}

}
