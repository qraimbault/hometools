package main

import (
	"github.com/qraimbault/hometools/vpn"
	"log"
)

func main() {
	err := vpn.UpdateVPNRecords()
	if err != nil {
		log.Fatal(err)
	}
}
