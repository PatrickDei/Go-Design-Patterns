package state

import "fmt"

type TCPEstablished struct {
	NetworkInterface
}

func (tcp *TCPEstablished) openConnection() error {
	return fmt.Errorf("connection already opened")
}

func (tcp *TCPEstablished) closeConnection() error {
	fmt.Println("open -> closed")
	tcp.NetworkInterface.currentState = tcp.NetworkInterface.tcpClosed
	return nil
}

func (tcp *TCPEstablished) sendInformation(info string) error {
	fmt.Printf("sending \"%v\"", info)
	return nil
}
