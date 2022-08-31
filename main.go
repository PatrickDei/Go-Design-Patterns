package main

import (
	"DesignPatterns/abstractfactory"
	"DesignPatterns/adapter"
	"DesignPatterns/bridge"
	"DesignPatterns/chainofresponsibility"
)

func main() {
	abstractfactory.ShowcaseAbstractFactory()
	adapter.ShowcaseAdapter()
	bridge.ShowcaseBridge()
	chainofresponsibility.ShowcaseChainOfResponsibility()
}
