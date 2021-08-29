package models

type User struct {
	id        int
	username  string
	password  string
	email     string
	lastLogin string
	isActive  bool
}

type Token struct {
	id          int
	userId      int
	token       string
	isActive    bool
	createdDate string
	isLastUse   bool
}

type Person struct {
	id       int
	userId   int
	name     string
	lastName string
}
