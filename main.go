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
	"DesignPatterns/memento"
	"DesignPatterns/observer"
	"DesignPatterns/prototype"
	"DesignPatterns/proxy"
	"DesignPatterns/singleton"
	"DesignPatterns/state"
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
	memento.ShowcaseMemento()
	observer.ShowcaseObserver()
	prototype.ShowcasePrototype()
	proxy.ShowcaseProxy()
	state.ShowcaseState()
	singleton.ShowcaseSingleton()
}
