package main

import (
	"fmt"
	"log"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

var foundDevName = false

func main() {
	// Find the network devices that you can use to capture packets
	var devices []pcap.Interface
	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Panicln("Cannot fetch network interfaces")
	}

	// Ask for a device a user would like to capture packets from
	var devNameFromUser string
	fmt.Print(
		"Please Enter the name of the device you want to analyz its packets: ",
	)

	fmt.Scanln(&devNameFromUser)

	// Ask for a filter a user would like to filter the packets
	var filter string
	fmt.Print("Please Enter your filter: ")

	fmt.Scanln(&filter)

	// Check whether there exists the device with the given name
	for _, device := range devices {
		if device.Name == devNameFromUser {
			fmt.Println("Device found")
			foundDevName = true
		}
	}

	if !foundDevName {
		log.Panicln("Cannot find the given device")
	}

	// Open live device
	handle, err := pcap.OpenLive(
		devNameFromUser,
		int32(65535),
		false,
		-1*time.Second,
	)

	if err != nil {
		fmt.Print(err)
		log.Panicln("Cannot open handle on the given device")
	}
	defer handle.Close()

	// Filter the packets that you want to monitor
	fmt.Println("Handle open")
	if err := handle.SetBPFFilter(filter); err != nil {
		log.Panicln(err)
	}

	packetSource := gopacket.NewPacketSource(
		handle,
		handle.LinkType(),
	)

	// Writing pcap file
	for packet := range packetSource.Packets() {
		fmt.Println(packet)
	}
}
