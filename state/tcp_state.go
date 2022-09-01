package state

type TCPState interface {
	openConnection() error
	closeConnection() error
	sendInformation(string) error
}
