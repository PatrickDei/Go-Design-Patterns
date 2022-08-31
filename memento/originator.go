package memento

type Memento struct {
	state string
}

func (m *Memento) getSavedState() string {
	return m.state
}

type Originator struct {
	state string
}

func (o *Originator) createMemento() *Memento {
	return &Memento{state: o.state}
}

func (o *Originator) restoreMemento(m *Memento) {
	o.state = m.getSavedState()
}

func (o *Originator) setState(state string) {
	o.state = state
}

func (o *Originator) getState() string {
	return o.state
}
