package usecase

import (
	"errors"
	"problem1/domain/model"
)

func (uu *userUsecaseS) getFriendList(id int) ([]*model.User, error) {
	flList, err := uu.userRepo.FindFriendLinkList(id)
	if err != nil {
		return nil, err
	}

	userID, err := uu.userRepo.FindUserID(id)
	if err != nil {
		return nil, err
	}

	idList := make([]int, 0)
	for _, fl := range flList {
		if fl.User1ID != userID {
			idList = append(idList, fl.User1ID)
		} else {
			idList = append(idList, fl.User2ID)
		}
	}

	uList := make([]*model.User, 0)
	for _, id := range idList {
		user, err := uu.userRepo.FindUser(id)
		if err != nil {
			return nil, err
		}
		uList = append(uList, user)
	}

	return uList, nil
}

func (uu *userUsecaseS) getFriendListExceptBlock(id int, fList []*model.User) ([]*model.User, error) {
	bList, err := uu.userRepo.FindBlockList(id)
	if err != nil {
		return nil, err
	}

	return uu.rmBlockUser(fList, bList, id)
}

func (uu *userUsecaseS) getFriendOfFriendList(fList []*model.User) ([]*model.User, error) {
	ffList := make([]*model.User, 0)
	for _, f := range fList {
		nfList, err := uu.GetFriendList(f.UserID)
		if err != nil {
			return nil, err
		}
		ffList = append(ffList, nfList...)
	}

	return uu.rm1HopFriend(ffList, fList), nil
}

func (uu *userUsecaseS) getUniqueList(fList []*model.User) []*model.User {
	nfList := make([]*model.User, 0)
	for _, f := range fList {
		isUnique := true
		for _, nf := range nfList {
			if f.UserID == nf.UserID {
				isUnique = false
			}
		}

		if isUnique {
			nfList = append(nfList, f)
		}
	}

	return nfList
}

func (uu *userUsecaseS) getPagingList(uList []*model.User, limit, page int) ([]*model.User, error) {
	if limit <= 0 {
		return nil, errors.New("limit is invalid")
	}
	if page <= 0 {
		return nil, errors.New("page is invalid")
	}

	end := limit * page
	start := end - (limit - 1)

	nuList := make([]*model.User, 0)
	for i, u := range uList {
		var nu model.User
		if start <= i+1 && i+1 <= end {
			nu.UserID, nu.Name = u.UserID, u.Name
			nuList = append(nuList, &nu)
		}
	}

	return nuList, nil
}
