package main

import (
	"fmt"

	"github.com/pascalhuerst/go-light/data"
	"github.com/pascalhuerst/go-light/qlcplus"
)

func main() {
	fmt.Println("Hello, world.")

	qlcFixture, err := qlcplus.ReadFixture("/usr/share/qlcplus/fixtures/Eurolite-LED-TMH-6.qxf")
	if err != nil {
		return
	}

	qlcplus.Print(qlcFixture)

	fixture := data.NewFixtureFromQlc(qlcFixture)

	fmt.Printf("\n\n\n%+v", fixture)

	err = data.WriteFixture(fixture)
	if err != nil {
		return
	}

}
