package model

type User struct {
	ID     int
	UserID int
	Name   string
}

type Link struct {
	ID      int
	User1ID int
	User2ID int
}
