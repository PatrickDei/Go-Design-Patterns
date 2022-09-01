package state

type NetworkInterface struct {
	tcpEstablished TCPState
	tcpClosed      TCPState

	currentState TCPState
}

func (ni *NetworkInterface) openConnection() error {
	return ni.currentState.openConnection()
}

func (ni *NetworkInterface) closeConnection() error {
	return ni.currentState.closeConnection()
}

func (ni *NetworkInterface) sendInformation(info string) error {
	return ni.currentState.sendInformation(info)
}
