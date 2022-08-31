package iterator

type User struct {
	name string
	age  int
}

type Iterator interface {
	hasNext() bool
	getNext() *User
}

type Collection interface {
	createIterator() Iterator
}

type UserCollection struct {
	users []*User
}

func (uc *UserCollection) createIterator() Iterator {
	return &UserIterator{
		users: uc.users,
	}
}

type UserIterator struct {
	index int
	users []*User
}

func (ui *UserIterator) hasNext() bool {
	if ui.index < len(ui.users) {
		return true
	}
	return false

}
func (ui *UserIterator) getNext() *User {
	if ui.hasNext() {
		user := ui.users[ui.index]
		ui.index++
		return user
	}
	return nil
}
