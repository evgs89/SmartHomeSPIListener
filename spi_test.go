// Copyright 2018 The Periph Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

package main

import (
	"fmt"
	"log"
	"periph.io/x/periph/conn/physic"
	"periph.io/x/periph/conn/spi"
	"periph.io/x/periph/conn/spi/spireg"
	"periph.io/x/periph/host"
)


func main() {
	state, err := host.Init()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(state)
	port, err := spireg.Open("")
	if err != nil {
		log.Fatal(err)
	}
	defer port.Close()
	connection, err := port.Connect(physic.MegaHertz, spi.Mode3, 8)
	if err != nil {
		log.Fatal(err)
	}
	write := []byte{0x10, 0x00}
	read := make([]byte, len(write))
	if err := connection.Tx(write, read); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("READ: %v\n", read)
}