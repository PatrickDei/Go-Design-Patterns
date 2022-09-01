package proxy

type PrivateDataServer interface {
	serve() string
}

type PrivateDataServerImpl struct{}

func (PrivateDataServerImpl) serve() string {
	return "Some private data that should be protected"
}

type PrivateDataServerProxy struct {
	PrivateDataServer
}

func (pds *PrivateDataServerProxy) serve() string {
	if pds.isAuthenticated() {
		return pds.PrivateDataServer.serve()
	}
	return "Access denied!"
}

func (pds *PrivateDataServerProxy) isAuthenticated() bool {
	return true
}
