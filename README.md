# Network Analyzer
> This project is mainly for fun and to explore about cybersecurity and network traffic.

### What is a network analyzer?
It is a tool used to monitor network traffic by analyzing streams of packets pass between any computers on a network.

### What is Packet Capture (PCAP)?
It is simply when you are trying to intercept (capture) of data packets travelling over a network.

## In the context of cybersecurity
Not only this can be useful for a variety of purposes, including troubleshooting network issues, monitoring network performance, and analyzing network traffic patterns. It can also be used for both defensive and offensive use.

- Detecting intrusion by monitoring the traffic for suspicious activities.
- Analyzing the captured traffic to understand how the attack happened.
- Analyzing the traffic to identify vulnerabilities and potential entry pints for an attack.
- Injecting a custom packet to exploit vulnerabilities.

## How to play with this tool
Since this project is built in Go, make sure that you can run Go on your machine. You can check your Go version by using the command `go version`, or download and install [here](https://go.dev/doc/install).

To use it, you have to `cd` into this directory and then run:

	go get github.com/google/gopacket
	go run main.go

Then input the name of the device you would like to inspect, follow by the filter that you would like to use for the captured packets.

If you don't know what are your available devices, you can use `ifconfig` to check your network interface. And [here](https://pkg.go.dev/github.com/google/gopacket/pcap#Handle.SetBPFFilter) is the format of what you can enter as a filter.

### Ideas for more extensions
- Improve CLI to be more user friendly.
- Make use of the captured packets to analyze the traffic.
- Make use of [GoPacket](https://pkg.go.dev/github.com/google/gopacket#hdr-Implementing_Your_Own_Decoder) to implement your own decoder to handle different kinds of packets.
