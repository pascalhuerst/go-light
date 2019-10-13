package main

import (
	"fmt"
	"net"
	"time"

	"github.com/jsimonetti/go-artnet/packet"
)

type ArtnetPacket struct {
}

func SendFrame() error {

	dst := fmt.Sprintf("%s:%d", "192.168.1.3", packet.ArtNetPort)
	node, _ := net.ResolveUDPAddr("udp", dst)
	src := fmt.Sprintf("%s:%d", "192.168.1.11", packet.ArtNetPort)
	localAddr, _ := net.ResolveUDPAddr("udp", src)

	conn, err := net.ListenUDP("udp", localAddr)
	if err != nil {
		fmt.Printf("error opening udp: %s\n", err)
		return err
	}

	// set channels 1 and 4 to FL, 2, 3 and 5 to FD
	// on my colorBeam this sets output 1 to fullbright red with zero strobing

	//for unviverse := 0; unviverse < 10; unviverse++ {
	for i := 0; i < 200; i++ {

		p := &packet.ArtDMXPacket{
			Sequence: uint8(i),
			SubUni:   3,
			Net:      0,
			Data:     [512]byte{0x00, uint8(i), 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, uint8(i), 0x00},
		}
		/*
			for i := 0; i < len(p.Data); i++ {
				p.Data[i] = 0xff
			}
		*/
		b, err := p.MarshalBinary()

		n, err := conn.WriteTo(b, node)
		if err != nil {
			fmt.Printf("error writing packet: %s\n", err)
			return nil
		}
		fmt.Printf("packet sent, wrote %d bytes\n", n)

		time.Sleep(5 * time.Millisecond)
	}
	//}
	return nil
}
