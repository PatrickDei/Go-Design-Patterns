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
	"DesignPatterns/factorymethod"
	"DesignPatterns/flyweight"
	"DesignPatterns/iterator"
	"DesignPatterns/mediator"
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
	factorymethod.ShowcaseFactoryMethod()
	flyweight.ShowcaseFlyweight()
	iterator.ShowcaseIterator()
	mediator.ShowcaseMediator()
}
