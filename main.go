package main

import (
	"fmt"

	"github.com/pascalhuerst/go-light/data"
	"github.com/pascalhuerst/go-light/qlcplus"
	"github.com/pascalhuerst/go-light/server"
)

func main() {
	fmt.Println("Hello, world.")

	qlcFixture, err := qlcplus.ReadFixture("/usr/share/qlcplus/fixtures/Eurolite-LED-TMH-6.qxf")
	if err != nil {
		return
	}

	qlcplus.PrintFixture(qlcFixture)

	fixture := data.NewFixtureFromQlc(qlcFixture)

	err = data.WriteFixture("fixture.json", fixture)
	if err != nil {
		return
	}

	readFixture, err := data.ReadFixture("fixture.json")
	if err != nil {
		return
	}

	fmt.Printf("ReadBack:\n%+v\n", readFixture)

	go server.StartHTTPServer(8123, readFixture)

	fmt.Printf("#### Server Should be running!")

	select {}

}
