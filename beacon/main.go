package main

import (
	"fmt"
	"math/rand"
	"net"
	"os"
	"time"
)

func main() {
	beacons := [4]string{
		"8a99597f-3d56-47d3-8eba-27f6128858e5",
		"9e2aff95-9091-454a-a29b-04a38aa3bf28",
		"dc3f9b99-d50f-4f1a-b38a-0e46c99d433b",
		"f06cd69b-274a-46b7-a8a8-fa9fbdf06edb",
	}

	ports := [4]string{"49164", "41936", "62172", "24442"}

	for i := 0; i < 1; {
		sendBeacon(beacons[rand.Intn(4)], ports[rand.Intn(4)])
		// time.Sleep(5 * time.Second)
		time.Sleep(5 * time.Minute)
	}
}

func sendBeacon(beacon, port string) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", fmt.Sprint("10.10.10.10:", port))
	if err != nil {
		println("ResolveTCPAddr failed:", err.Error())
		os.Exit(1)
	}

	// Can't connect to server or port
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		println("Dial failed:", err.Error())
		os.Exit(1)
	}

	// Can't write the beacon
	_, err = conn.Write([]byte(beacon))
	if err != nil {
		println("Write to server failed:", err.Error())
		os.Exit(1)
	}

	// Print that we're finished sending the beacon
	println(time.Now().Format(time.Kitchen), "->", beacon, "on", port)
	conn.Close()
}
