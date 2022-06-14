package handler

import (
	"problem1/domain/model"
)

func getResUserList(uList []*model.User) []resUser {
	res := make([]resUser, 0)
	for _, u := range uList {
		user := resUser {
			UserID: u.UserID,
			Name:   u.Name,
		}
		res = append(res, user)
	}

	return res
}
