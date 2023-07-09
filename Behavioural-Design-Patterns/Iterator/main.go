package main

import "fmt"

///////////// Collection Item ////////////////////

// we want to iterate over users
type user struct {
	name string
	age  int
}

//////////// Interfaces //////////////////////////

type iterator interface {
	hasNext() bool
	getNext() *user
}

type collection interface {
	createIterator() iterator
}

/////////// Iterator Implementation //////////////

type userIterator struct {
	index int
	users []*user
}

func (u *userIterator) hasNext() bool {
	if u.index < len(u.users) {
		return true
	} else {
		return false
	}
}

func (u *userIterator) getNext() (user *user) {
	if u.hasNext() {
		user = u.users[u.index]
		u.index++
	}
	return
}

///////////// Collection Implementation //////////

type userCollection struct {
	users []*user
}

func (u *userCollection) createIterator() iterator {
	return &userIterator{
		users: u.users,
	}
}

/////////////// Main /////////////////////////////

func main() {
	// create users
	user1 := &user{
		name: "Michael",
		age:  30,
	}

	user2 := &user{
		name: "Thomas",
		age:  20,
	}

	// creat user collection
	userCollection := &userCollection{
		users: []*user{user1, user2},
	}

	// iterate over user collection
	iterator := userCollection.createIterator()
	for iterator.hasNext() {
		user := iterator.getNext()
		fmt.Printf("User is %+v.\n", user)
	}
}
