package state

import "fmt"

type TCPClosed struct {
	NetworkInterface
}

func (tcp *TCPClosed) openConnection() error {
	fmt.Println("Closed -> Open")
	tcp.NetworkInterface.currentState = tcp.NetworkInterface.tcpEstablished
	return nil
}

func (tcp *TCPClosed) closeConnection() error {
	return fmt.Errorf("connection already closed")
}

func (tcp *TCPClosed) sendInformation(info string) error {
	return fmt.Errorf("cannot send \"%v\" over a closed connection", info)
}
