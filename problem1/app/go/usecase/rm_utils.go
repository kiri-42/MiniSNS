package usecase

import (
	"problem1/domain/model"
)

func (uu *userUsecaseS) rmBlockUser(fList []*model.User, bList []*model.Link, id int) ([]*model.User, error) {
	nuList := make([]*model.User, 0)

	userID, err := uu.userRepo.FindUserID(id)
	if err != nil {
		return nil, err
	}

	for _, f := range fList {
		isBlock := false

		for _, b := range bList {
			if b.User1ID == f.UserID && b.User2ID == userID || b.User1ID == userID && b.User2ID == f.UserID {
				isBlock = true
				break
			}
		}

		if !isBlock {
			nuList = append(nuList, f)
		}
	}

	return nuList, nil
}

func (uu *userUsecaseS) rm1HopFriend(ffList []*model.User, fList []*model.User) []*model.User {
	nffList := make([]*model.User, 0)

	for _, ff := range ffList {
		isFriend := false

		for _, f := range fList {
			if f.UserID == ff.UserID {
				isFriend = true
				break
			}
		}

		if !isFriend {
			nffList = append(nffList, ff)
		}
	}

	return nffList
}
