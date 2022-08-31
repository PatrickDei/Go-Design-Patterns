package main

import (
	"DesignPatterns/abstractfactory"
	"DesignPatterns/adapter"
	"DesignPatterns/bridge"
	"DesignPatterns/chainofresponsibility"
	"DesignPatterns/command"
	"DesignPatterns/composite"
	"DesignPatterns/decorator"
	"DesignPatterns/facade"
)

func main() {
	abstractfactory.ShowcaseAbstractFactory()
	adapter.ShowcaseAdapter()
	bridge.ShowcaseBridge()
	chainofresponsibility.ShowcaseChainOfResponsibility()
	command.ShowcaseCommand()
	composite.ShowcaseComposite()
	decorator.ShowcaseDecorator()
	facade.ShowcaseFacade()
}
