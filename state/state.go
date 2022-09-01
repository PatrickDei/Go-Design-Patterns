package state

import "fmt"

func ShowcaseState() {
	fmt.Println("\nState pattern")

	netInt := NetworkInterface{}

	tcpClosed := TCPClosed{NetworkInterface: netInt}
	tcpEstablished := TCPEstablished{NetworkInterface: netInt}

	netInt.tcpClosed = &tcpClosed
	netInt.tcpEstablished = &tcpEstablished
	netInt.currentState = &tcpClosed

	fmt.Println(netInt.sendInformation("failing info"))
	fmt.Println(netInt.openConnection())
	fmt.Println(netInt.sendInformation("passing info"))
	fmt.Println(netInt.openConnection())
	fmt.Println(netInt.closeConnection())
	fmt.Println(netInt.closeConnection())
}
