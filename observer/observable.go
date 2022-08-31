package observer

type Observable interface {
	subscribe(o UserObserver)
	notifyObservers()
}

type UserSubject struct {
	name     string
	surname  string
	username string

	observers []UserObserver
}

func (us *UserSubject) subscribe(o UserObserver) {
	us.observers = append(us.observers, o)
	o.onChange(*us)
}

func (us *UserSubject) notifyObservers() {
	for _, o := range us.observers {
		o.onChange(*us)
	}
}

func (us *UserSubject) setName(name string) {
	us.name = name
	us.notifyObservers()
}

func (us *UserSubject) setSurname(surname string) {
	us.surname = surname
	us.notifyObservers()
}

func (us *UserSubject) setUsername(username string) {
	us.username = username
	us.notifyObservers()
}
