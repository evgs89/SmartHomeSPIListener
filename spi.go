package main

//#include <data_struct.h>
import "C"

import (
	"fmt"
	"log"
	"periph.io/x/periph/conn/physic"
	"periph.io/x/periph/conn/spi"
	"periph.io/x/periph/conn/spi/spireg"
	"periph.io/x/periph/host"
)

func main() {
	port, connection := makeSpiConnection()
	defer port.Close()
	for {
		read := send_data(connection, []byte{0x10, 0x00})
		fmt.Printf("READ: %v\n", read)
	}
}

func send_data(connection spi.Conn, data []byte) []byte {
	write := data
	read := make([]byte, len(write))
	if err := connection.Tx(write, read); err != nil {
		log.Fatal(err)
	}
	return read
}

func makeSpiConnection() (spi.PortCloser, spi.Conn) {
	state, err := host.Init()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(state)
	port, err := spireg.Open("")
	if err != nil {
		log.Fatal(err)
	}
	connection, err := port.Connect(physic.MegaHertz, spi.Mode3, 8)
	if err != nil {
		log.Fatal(err)
	}
	return port, connection
}