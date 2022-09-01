package proxy

import "fmt"

func ShowcaseProxy() {
	fmt.Println("\nProxy pattern")

	pds := PrivateDataServerProxy{PrivateDataServer: PrivateDataServerImpl{}}
	fmt.Println(pds.serve())
}
