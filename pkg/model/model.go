package model

import "errors"

type User struct {
	Id		int 	`json:"id"`
	Name	string	`json:"name"`
}

type UserList struct {
	Users		[]User
	LastIndex	int
}

func (ul *UserList) AddLastIndex(idx int) {
	ul.LastIndex = idx
}

func (ul *UserList) GetUserIndex(id int) (int, error) {
	for index, value := range ul.Users{
		if value.Id == id {
			return index, nil
		}
	}
	return 0, errors.New("no this id")
}




func (u *User) AddId(id int) {
	u.Id = id + 1
}

func (u *User) GetId() int {
	return u.Id
}